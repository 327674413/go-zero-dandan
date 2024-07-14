package logic

import (
	"context"
	"github.com/zeromicro/go-zero/core/mr"
	"go-zero-dandan/app/user/model"
	"go-zero-dandan/app/user/rpc/internal/svc"
	"go-zero-dandan/app/user/rpc/types/userRpc"
	"go-zero-dandan/common/resd"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetUserNormalInfoLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewGetUserNormalInfoLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetUserNormalInfoLogic {
	return &GetUserNormalInfoLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *GetUserNormalInfoLogic) GetUserNormalInfo(in *userRpc.GetUserInfoReq) (*userRpc.GetUserNormalInfoResp, error) {
	if err := l.checkReqParams(in); err != nil {
		return nil, err
	}
	userMainModel := model.NewUserMainModel(l.ctx, l.svcCtx.SqlConn, in.PlatId)
	userInfoMap, err := l.getNormalInfoByIds(in.Ids, userMainModel)
	data := make(map[string]*userRpc.UserNormalInfo)
	for _, id := range in.Ids {
		if v, ok := userInfoMap[id]; ok {
			data[id] = &userRpc.UserNormalInfo{MainInfo: &userRpc.UserMainInfo{
				Id:        v.Id,
				UnionId:   v.UnionId,
				StateEm:   v.StateEm,
				Account:   v.Account,
				Nickname:  v.Nickname,
				Phone:     v.Phone,
				PhoneArea: v.PhoneArea,
				SexEm:     v.SexEm,
				Email:     v.Email,
				AvatarImg: v.AvatarImg,
				PlatId:    v.PlatId,
				Signature: v.Signature,
			}}
		}

	}
	if err != nil {
		return nil, resd.ErrorCtx(l.ctx, err)
	}

	return &userRpc.GetUserNormalInfoResp{Users: data}, nil
}

func (l *GetUserNormalInfoLogic) getNormalInfoByIds(ids []string, userMainModel model.UserMainModel) (map[string]*model.UserMain, error) {
	//通过并行获取数据
	userMainInfos, err := mr.MapReduce[string, *model.UserMain, map[string]*model.UserMain](func(source chan<- string) {
		//生成要处理的数据
		for _, id := range ids {
			source <- id
		}
	}, func(id string, writer mr.Writer[*model.UserMain], cancel func(error)) {
		//处理数据
		userMain, err := userMainModel.CacheFindById(l.svcCtx.Redis, id)
		if err != nil {
			cancel(err)
			return
		}
		writer.Write(userMain)
	}, func(pipe <-chan *model.UserMain, writer mr.Writer[map[string]*model.UserMain], cancel func(error)) {
		//聚合
		ds := make(map[string]*model.UserMain)
		for item := range pipe {
			ds[item.Id] = item
		}
		writer.Write(ds)
	})
	if err != nil {
		return nil, resd.ErrorCtx(l.ctx, err)
	}

	return userMainInfos, nil
}
func (l *GetUserNormalInfoLogic) checkReqParams(in *userRpc.GetUserInfoReq) error {
	if in.PlatId == "" {
		return resd.NewRpcErrWithTempCtx(l.ctx, "参数缺少platId", resd.ErrReqFieldRequired1, "PlatId")
	}
	if len(in.Ids) == 0 {
		return resd.NewRpcErrWithTempCtx(l.ctx, "参数缺少platId", resd.ErrReqFieldRequired1, "Ids")
	}
	return nil
}
