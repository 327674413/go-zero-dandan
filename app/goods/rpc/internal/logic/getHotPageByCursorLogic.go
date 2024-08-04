package logic

import (
	"context"
	"fmt"
	"github.com/zeromicro/go-zero/core/logc"
	"github.com/zeromicro/go-zero/core/mr"
	"github.com/zeromicro/go-zero/core/threading"
	"go-zero-dandan/app/goods/model"
	"go-zero-dandan/app/goods/rpc/internal/svc"
	"go-zero-dandan/app/goods/rpc/types/goodsRpc"
	"go-zero-dandan/common/resd"
	"math"
	"sort"
)

type GetHotPageByCursorLogic struct {
	*GetHotPageByCursorLogicGen
}

const (
	defaultCacheGoodsNum   = 5
	defaultPageSize        = 10
	maxPageSize            = 100
	redisKeyHotViewGoodses = "HotViewByCursor"
	defaultCacheExpireSec  = 3600
	defaultCursor          = math.MaxInt32 //倒序找，默认的最大值
	endFlagId              = "-1"          //倒序找，最后一条用-1来判断
)

func NewGetHotPageByCursorLogic(ctx context.Context, svc *svc.ServiceContext) *GetHotPageByCursorLogic {
	return &GetHotPageByCursorLogic{
		GetHotPageByCursorLogicGen: NewGetHotPageByCursorLogicGen(ctx, svc),
	}
}

// 这个方式缺陷还挺多：在边界时只能返回缓存剩余数据，导致数量不满足分页要求 和 关联引发的系列问题；当然在高并发和加载更多时可能可以忽略

func (l *GetHotPageByCursorLogic) GetHotPageByCursor(in *goodsRpc.GetHotPageByCursorReq) (*goodsRpc.GetPageByCursorResp, error) {
	if err := l.initReq(in); err != nil {
		return nil, l.resd.Error(err)
	}
	var size int64
	if l.req.Size > 0 {
		size = l.req.Size
	} else {
		size = defaultPageSize
	}
	if size > maxPageSize {
		size = maxPageSize
	}
	if l.req.Page <= 0 {
		l.req.Page = 1
	}
	if l.req.Cursor == 0 {
		l.req.Cursor = defaultCursor
	}
	//先尝试从缓存中获取数据
	ids, _ := l.getCacheIds(l.req.Page, size)
	var (
		currCacheData  []*model.GoodsMain
		lastId         string
		respCursor     int64
		isCache, isEnd bool
	)
	currPageGoodses := make([]*goodsRpc.GoodsInfo, 0)
	goodsModel := model.NewGoodsMainModel(l.ctx, l.svc.SqlConn, l.meta.PlatId)
	if len(ids) > 0 {
		//存在缓存
		isCache = true
		if ids[len(ids)-1] == endFlagId {
			//没有数据，最后一页
			isEnd = true
		}
		// 根据id获取数据
		list, err := l.getGoodsByIds(ids)
		if err != nil {
			return nil, resd.ErrorCtx(l.ctx, err)
		}
		//缓存读出来是并发度的，乱序，重新排序
		sort.Slice(list, func(i, j int) bool {
			return list[i].ViewNum > list[j].ViewNum
		})
		for _, item := range list {
			currPageGoodses = append(currPageGoodses, &goodsRpc.GoodsInfo{
				Id:        item.Id,
				Name:      item.Name,
				Spec:      item.Spec,
				Cover:     item.Cover,
				SellPrice: item.SellPrice,
				StoreQty:  item.StoreQty,
				State:     item.State,
				IsSpecial: item.IsSpecial,
				UnitId:    item.UnitId,
				UnitName:  item.UnitName,
				PlatId:    item.PlatId,
				ViewNum:   item.ViewNum,
			})
		}
	} else {
		//不存在缓存，通过单通道查数据，其他请求等待查完后一并返回方式查询
		v, err, _ := l.svc.SingleFlightGroup.Do(redisKeyHotViewGoodses, func() (interface{}, error) {
			//这里要先判断，是一条缓存都没有的初始化，还是按需加载，如果是初始化应该用默认defaultCursor，如果是按需加载就用用户传的cursor
			isExistCache, err := l.svc.Redis.ExistsCtx(l.ctx, redisKeyHotViewGoodses)
			if err != nil {
				return nil, resd.ErrorCtx(l.ctx, err)
			}
			curr := l.req.Cursor
			if !isExistCache || curr == 0 {
				//不存在key，代表从无到有初始化，应该走默认的第一页数据的cursor
				//存在key，只能是按需加载，用请求的cursor
				curr = defaultCursor
			}
			return goodsModel.Order("view_num DESC").Where("view_num <= ?", curr).Limit(defaultCacheGoodsNum).Select()
		})
		if err != nil || v == nil {
			return nil, resd.ErrorCtx(l.ctx, err)
		}
		currCacheData = v.([]*model.GoodsMain)
		var firstPageList []*model.GoodsMain
		if len(currCacheData) > int(size) {
			//当前获取的全部数据 大于 分页数，则按分页数截取当前页数据
			firstPageList = currCacheData[:int(size)]
		} else {
			//小于分页数则代表实际数据只有一页，到最后了
			firstPageList = currCacheData
			isEnd = true
		}
		for _, item := range firstPageList {
			currPageGoodses = append(currPageGoodses, &goodsRpc.GoodsInfo{
				Id:        item.Id,
				Name:      item.Name,
				Spec:      item.Spec,
				Cover:     item.Cover,
				SellPrice: item.SellPrice,
				StoreQty:  item.StoreQty,
				State:     item.State,
				IsSpecial: item.IsSpecial,
				UnitId:    item.UnitId,
				UnitName:  item.UnitName,
				PlatId:    item.PlatId,
				ViewNum:   item.ViewNum,
			})
		}
	}
	if len(currPageGoodses) > 0 {
		//如果有数据，获取最后一条数据id，给前端下次请求时作为cursor带上
		lastData := currPageGoodses[len(currPageGoodses)-1]
		respCursor = lastData.ViewNum
		lastId = lastData.Id
	}
	if !isCache {
		//不是缓存，则异步写入缓存
		threading.GoSafe(func() {
			if len(currCacheData) < defaultCacheGoodsNum && len(currCacheData) > 0 {
				//有数据 且小于默认缓存数量，代表所有数据已经全部缓存了，写入一个-1作为缓存，识别已经最后一条
				currCacheData = append(currCacheData, &model.GoodsMain{Id: endFlagId})
			}
			err := l.addListCache(currCacheData)
			if err != nil {
				logc.Error(l.ctx, err)
			}
		})
	}
	return &goodsRpc.GetPageByCursorResp{
		Size:    size,
		List:    currPageGoodses,
		IsCache: isCache,
		IsEnd:   isEnd,
		Cursor:  respCursor,
		LastId:  lastId,
	}, nil
}
func (l *GetHotPageByCursorLogic) getCacheIds(page, size int64) ([]string, error) {
	isExist, err := l.svc.Redis.Exists(redisKeyHotViewGoodses)
	if err != nil {
		//保证能查到，缓存异常不处理
	}
	if isExist {
		// 如果缓存是定期失效的，防止缓存穿透，这里应该执行一次续期
		//err = l.svc.Redis.ExpireCtx(l.ctx, "hotView", "", 7200)
	}
	// zrevrange 是高到低，所以分页时，cursor要放到max里，也就是这里的stop， page是计算偏移量的，游标分页不需要，从cursor开始获取目标行数即可， 如果低到高应该是放到start里，确认，待定
	pair, err := l.svc.Redis.ZpageCtx(l.ctx, redisKeyHotViewGoodses, "", int(page), int(size), true)
	ids := make([]string, 0)
	for _, item := range pair {
		ids = append(ids, item.Key)
	}
	return ids, nil
}
func (l *GetHotPageByCursorLogic) getGoodsByIds(ids []string) ([]*model.GoodsMain, error) {
	//通过并行获取数据
	goodses, err := mr.MapReduce[string, *model.GoodsMain, []*model.GoodsMain](func(source chan<- string) {
		//生成要处理的数据
		for _, id := range ids {
			if id == "-1" {
				continue
			}
			source <- id
		}
	}, func(id string, writer mr.Writer[*model.GoodsMain], cancel func(error)) {
		//处理数据
		goodsModel := model.NewGoodsMainModel(l.ctx, l.svc.SqlConn, l.meta.PlatId)
		goods, err := goodsModel.CacheFindById(l.svc.Redis, id)
		if err != nil {
			cancel(err)
			return
		}
		writer.Write(goods)
	}, func(pipe <-chan *model.GoodsMain, writer mr.Writer[[]*model.GoodsMain], cancel func(error)) {
		//聚合
		var ds []*model.GoodsMain
		for item := range pipe {
			ds = append(ds, item)
		}
		writer.Write(ds)
	})
	if err != nil {
		return nil, resd.ErrorCtx(l.ctx, err)
	}

	return goodses, nil
}

func (l *GetHotPageByCursorLogic) addListCache(list []*model.GoodsMain) error {
	if len(list) == 0 {
		return nil
	}
	for _, goods := range list {
		score := goods.ViewNum
		_, err := l.svc.Redis.Zadd(redisKeyHotViewGoodses, "", score, fmt.Sprintf("%d", goods.Id))
		if err != nil {
			return resd.ErrorCtx(l.ctx, err)
		}
	}
	return l.svc.Redis.ExpireCtx(l.ctx, redisKeyHotViewGoodses, "", defaultCacheExpireSec)
}
