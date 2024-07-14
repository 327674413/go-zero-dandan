// Code generated by goctl. DO NOT EDIT.
package logic

import (
	"context"

	"go-zero-dandan/app/asset/api/internal/svc"
	"go-zero-dandan/app/asset/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
	"go-zero-dandan/common/resd"
	"go-zero-dandan/common/typed"
)

type MultipartUploadSendLogicGen struct {
	logx.Logger
	ctx          context.Context
	svc          *svc.ServiceContext
	resd         *resd.Resp
	meta         *typed.ReqMeta
	hasUserInfo  bool
	mustUserInfo bool
	req          struct {
		UploadID   string `form:"uploadId"`
		ChunkIndex int64  `form:"chunkIndex"`
	}
	hasReq struct {
		UploadID   bool
		ChunkIndex bool
	}
}

func NewMultipartUploadSendLogicGen(ctx context.Context, svc *svc.ServiceContext) *MultipartUploadSendLogicGen {
	meta, _ := ctx.Value("reqMeta").(*typed.ReqMeta)
	if meta == nil {
		meta = &typed.ReqMeta{}
	}
	return &MultipartUploadSendLogicGen{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svc:    svc,
		resd:   resd.NewResp(ctx, resd.I18n.NewLang(meta.Lang)),
		meta:   meta,
	}
}

func (l *MultipartUploadSendLogicGen) initReq(req *types.MultipartUploadSendReq) error {

	if req.UploadID != nil {
		l.req.UploadID = *req.UploadID
		l.hasReq.UploadID = true
	} else {
		l.hasReq.UploadID = false
	}

	if req.ChunkIndex != nil {
		l.req.ChunkIndex = *req.ChunkIndex
		l.hasReq.ChunkIndex = true
	} else {
		l.hasReq.ChunkIndex = false
	}
	l.hasUserInfo = true
	l.mustUserInfo = true

	return nil
}
