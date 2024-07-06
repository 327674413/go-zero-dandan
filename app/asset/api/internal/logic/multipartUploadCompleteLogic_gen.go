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

type MultipartUploadCompleteLogicGen struct {
	logx.Logger
	ctx          context.Context
	svc          *svc.ServiceContext
	resd         *resd.Resp
	lang         string
	userMainInfo *user.UserMainInfo
	platId       string
	platClasEm   int64
	hasUserInfo  bool
	mustUserInfo bool
	ReqFileSha1  string `json:"fileSha1"`
	ReqUploadId  string `json:"uploadId"`
	HasReq       struct {
		FileSha1 bool
		UploadId bool
	}
}

func NewMultipartUploadCompleteLogicGen(ctx context.Context, svc *svc.ServiceContext) *MultipartUploadCompleteLogicGen {
	lang, _ := ctx.Value("lang").(string)
	return &MultipartUploadCompleteLogicGen{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svc:    svc,
		lang:   lang,
		resd:   resd.NewResd(ctx, resd.I18n.NewLang(lang)),
	}
}

func (l *MultipartUploadCompleteLogicGen) initReq(req *types.MultipartUploadCompleteReq) error {
	var err error
	if err = l.initPlat(); err != nil {
		return resd.ErrorCtx(l.ctx, err)
	}

	if req.FileSha1 != nil {
		l.ReqFileSha1 = *req.FileSha1
		l.HasReq.FileSha1 = true
	} else {
		l.HasReq.FileSha1 = false
	}

	if req.UploadId != nil {
		l.ReqUploadId = *req.UploadId
		l.HasReq.UploadId = true
	} else {
		l.HasReq.UploadId = false
	}
	l.hasUserInfo = true
	l.mustUserInfo = true

	return nil
}

func (l *MultipartUploadCompleteLogicGen) initUser() (err error) {
	userMainInfo, ok := l.ctx.Value("userMainInfo").(*user.UserMainInfo)
	if !ok {
		return resd.NewErrCtx(l.ctx, "未配置userInfo中间件", resd.UserMainInfoErr)
	}
	l.userMainInfo = userMainInfo
	return nil
}

func (l *MultipartUploadCompleteLogicGen) initPlat() (err error) {
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
