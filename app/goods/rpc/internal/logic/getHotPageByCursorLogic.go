package logic

import (
	"context"
	"fmt"
	"github.com/zeromicro/go-zero/core/logc"
	"github.com/zeromicro/go-zero/core/mr"
	"github.com/zeromicro/go-zero/core/threading"
	"go-zero-dandan/app/goods/model"
	"go-zero-dandan/common/resd"
	"math"
	"strconv"

	"go-zero-dandan/app/goods/rpc/internal/svc"
	"go-zero-dandan/app/goods/rpc/types/pb"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetHotPageByCursorLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

const (
	defaultCacheGoodsNum   = 200
	defaultCursor          = math.MaxInt32
	redisKeyHotViewGoodses = "HotViewByCursor"
	defaultCacheExpireSec  = 3600
)

func NewGetHotPageByCursorLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetHotPageByCursorLogic {
	return &GetHotPageByCursorLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *GetHotPageByCursorLogic) GetHotPage(in *pb.GetHotPageByCursorReq) (*pb.GetPageResp, error) {
	size := in.Size
	cursor := in.Cursor
	//先尝试从缓存中获取数据
	ids, _ := l.getCacheIds(cursor, size)
	fmt.Println("ids:", ids)
	var (
		currPageGoodses []*pb.GoodsInfo
		currCacheData   []*model.GoodsMain
		lastId          int64
		isCache, isEnd  bool
	)
	if cursor == 0 {
		//如果没传则代表从第一页开始查
		cursor = defaultCursor
	}
	goodsModel := model.NewGoodsMainModel(l.svcCtx.SqlConn, in.PlatId)
	if len(ids) > 0 {
		//存在缓存
		isCache = true
		if ids[len(ids)-1] == -1 {
			//没有数据，最后一页
			isEnd = true
		}
		// 根据id获取数据
		list, err := l.getGoodsByIds(ids, &goodsModel)
		if err != nil {
			for _, item := range list {
				currPageGoodses = append(currPageGoodses, &pb.GoodsInfo{
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
				})
			}
		}
	} else {
		//不存在缓存，通过单通道查数据，其他请求等待查完后一并返回方式查询
		v, err, _ := l.svcCtx.SingleFlightGroup.Do(redisKeyHotViewGoodses, func() (interface{}, error) {
			//这里要先判断，是一条缓存都没有的初始化，还是按需加载，如果是初始化应该用默认defaultCursor，如果是按需加载就用用户传的cursor
			isExistCache, err := l.svcCtx.Redis.ExistsCtx(l.ctx, redisKeyHotViewGoodses)
			if err != nil {
				return nil, resd.ErrorCtx(l.ctx, err)
			}
			var curr int64
			if isExistCache == true {
				//存在key，只能是按需加载，用请求的cursor
				curr = in.Cursor
			} else {
				//不存在key，代表从无到有初始化，应该走默认的第一页数据的cursor
				curr = defaultCursor
			}
			return goodsModel.Order("view_num DESC").Where("view_num < ?", curr).Limit(defaultCacheGoodsNum).Select()
		})
		if err != nil || v == nil {
			return nil, resd.ErrorCtx(l.ctx, err)
		}
		currCacheData = v.([]*model.GoodsMain)
		var firstPageList []*model.GoodsMain
		if len(currCacheData) > int(in.Size) {
			//当前获取的全部数据 大于 分页数，则按分页数截取当前页数据
			firstPageList = currCacheData[:int(in.Size)]
		} else {
			//小于分页数则代表实际数据只有一页，到最后了
			firstPageList = currCacheData
			isEnd = true
		}
		for _, item := range firstPageList {
			currPageGoodses = append(currPageGoodses, &pb.GoodsInfo{
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
			})
		}
	}

	if len(currPageGoodses) > 0 {
		//如果有数据，获取最后一条数据id，给前端下次请求时作为cursor带上
		lastData := currPageGoodses[len(currPageGoodses)-1]
		lastId = lastData.Id
		cursor = lastData.ViewNum
	}
	if !isCache {
		//不是缓存，则异步写入缓存
		threading.GoSafe(func() {
			if len(currCacheData) < defaultCacheGoodsNum && len(currCacheData) > 0 {
				//有数据 且小于默认缓存数量，代表所有数据已经全部缓存了，写入一个-1缓存，识别已经最后一条
				currCacheData = append(currCacheData, &model.GoodsMain{Id: -1})
			}
			err := l.addListCache(currCacheData)
			if err != nil {
				logc.Error(l.ctx, err)
			}
		})
	}
	return &pb.GetPageResp{
		Size:    size,
		List:    currPageGoodses,
		IsCache: isCache,
		IsEnd:   &isEnd,
		Cursor:  &cursor,
		LastId:  &lastId,
	}, nil
}
func (l *GetHotPageByCursorLogic) getCacheIds(cursor, size int64) ([]int64, error) {
	isExist, err := l.svcCtx.Redis.ExistsCtx(l.ctx, redisKeyHotViewGoodses)
	if err != nil {
		//保证能查到，缓存异常不处理
	}
	if isExist {
		// 如果缓存是定期失效的，防止缓存穿透，这里应该执行一次续期
		//err = l.svcCtx.Redis.ExpireCtx(l.ctx, "hotView", "", 7200)
	}
	// zrevrange 是高到低，所以分页时，cursor要放到max里，也就是这里的stop， page是计算偏移量的，游标分页不需要，从cursor开始获取目标行数即可， 如果低到高应该是放到start里，确认，待定
	pair, err := l.svcCtx.Redis.Client().ZrevrangebyscoreWithScoresAndLimitCtx(l.ctx, l.svcCtx.Redis.FieldKey("hotView", ""), 0, cursor, 0, int(size))
	ids := make([]int64, 0)
	for _, item := range pair {
		id, err := strconv.ParseInt(item.Key, 10, 64)
		if err != nil {
			return nil, resd.ErrorCtx(l.ctx, err)
		}
		ids = append(ids, id)
	}
	return ids, nil
}
func (l *GetHotPageByCursorLogic) getGoodsByIds(ids []int64, goodsModel *model.GoodsMainModel) ([]*model.GoodsMain, error) {
	//通过并行获取数据
	goodses, err := mr.MapReduce[int64, *model.GoodsMain, []*model.GoodsMain](func(source chan<- int64) {
		//生成要处理的数据
		for _, id := range ids {
			if id == -1 {
				continue
			}
			source <- id
		}
	}, func(id int64, writer mr.Writer[*model.GoodsMain], cancel func(error)) {
		//处理数据
		goods, err := (*goodsModel).CacheFindById(l.svcCtx.Redis, id)
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
		_, err := l.svcCtx.Redis.ZaddCtx(l.ctx, redisKeyHotViewGoodses, "", score, fmt.Sprintf("%d", goods.Id))
		if err != nil {
			return resd.ErrorCtx(l.ctx, err)
		}
	}
	return l.svcCtx.Redis.ExpireCtx(l.ctx, redisKeyHotViewGoodses, "", defaultCacheExpireSec)
}
