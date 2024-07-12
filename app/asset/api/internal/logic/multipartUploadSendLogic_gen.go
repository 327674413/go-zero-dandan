// Code generated by goctl. DO NOT EDIT.
package logic

import (
	"context"

	"go-zero-dandan/app/asset/api/internal/svc"
	"go-zero-dandan/app/asset/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
	"go-zero-dandan/app/user/rpc/user"
	"go-zero-dandan/common/resd"
	"go-zero-dandan/common/utild"
)

type MultipartUploadSendLogicGen struct {
	logx.Logger
	ctx           context.Context
	svc           *svc.ServiceContext
	resd          *resd.Resp
	lang          string
	userMainInfo  *user.UserMainInfo
	platId        string
	platClasEm    int64
	hasUserInfo   bool
	mustUserInfo  bool
	ReqUploadID   string `form:"uploadId"`
	ReqChunkIndex string `form:"chunkIndex"`
	HasReq        struct {
		UploadID   bool
		ChunkIndex bool
	}
}

func NewMultipartUploadSendLogicGen(ctx context.Context, svc *svc.ServiceContext) *MultipartUploadSendLogicGen {
	lang, _ := ctx.Value("lang").(string)
	return &MultipartUploadSendLogicGen{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svc:    svc,
		lang:   lang,
		resd:   resd.NewResd(ctx, resd.I18n.NewLang(lang)),
	}
}

func (l *MultipartUploadSendLogicGen) initReq(req *types.MultipartUploadSendReq) error {
	var err error
	if err = l.initPlat(); err != nil {
		return resd.ErrorCtx(l.ctx, err)
	}

	if req.UploadID != nil {
		l.ReqUploadID = *req.UploadID
		l.HasReq.UploadID = true
	} else {
		l.HasReq.UploadID = false
	}

	if req.ChunkIndex != nil {
		l.ReqChunkIndex = *req.ChunkIndex
		l.HasReq.ChunkIndex = true
	} else {
		l.HasReq.ChunkIndex = false
	}
	l.hasUserInfo = true
	l.mustUserInfo = true

	return nil
}

func (l *MultipartUploadSendLogicGen) initUser() (err error) {
	userMainInfo, ok := l.ctx.Value("userMainInfo").(*user.UserMainInfo)
	if !ok {
		return resd.NewErrCtx(l.ctx, "未配置userInfo中间件", resd.UserMainInfoErr)
	}
	l.userMainInfo = userMainInfo
	return nil
}

func (l *MultipartUploadSendLogicGen) initPlat() (err error) {
	platClasEm := utild.AnyToInt64(l.ctx.Value("platClasEm"))
	if platClasEm == 0 {
		return resd.NewErrCtx(l.ctx, "token中未获取到platClasEm", resd.PlatClasErr)
	}
	platId, _ := l.ctx.Value("platId").(string)
	if platId == "" {
		return resd.NewErrCtx(l.ctx, "token中未获取到platId", resd.PlatIdErr)
	}
	l.platId = platId
	l.platClasEm = platClasEm
	return nil
}