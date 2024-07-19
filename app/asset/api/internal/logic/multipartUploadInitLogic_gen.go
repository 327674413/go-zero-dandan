// Code generated by goctl. DO NOT EDIT.
package logic

import (
	"context"

	"go-zero-dandan/app/asset/api/internal/svc"
	"go-zero-dandan/app/asset/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
	"go-zero-dandan/common/resd"
	"go-zero-dandan/common/typed"
	"strings"
)

type MultipartUploadInitLogicGen struct {
	logx.Logger
	ctx          context.Context
	svc          *svc.ServiceContext
	resd         *resd.Resp
	meta         *typed.ReqMeta
	hasUserInfo  bool
	mustUserInfo bool
	req          struct {
		FileName string `json:"fileName"`
		FileSha1 string `json:"fileSha1"`
		FileSize int64  `json:"fileSize"`
	}
	hasReq struct {
		FileName bool
		FileSha1 bool
		FileSize bool
	}
}

func NewMultipartUploadInitLogicGen(ctx context.Context, svc *svc.ServiceContext) *MultipartUploadInitLogicGen {
	meta, _ := ctx.Value("reqMeta").(*typed.ReqMeta)
	if meta == nil {
		meta = &typed.ReqMeta{}
	}
	return &MultipartUploadInitLogicGen{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svc:    svc,
		resd:   resd.NewResp(ctx, meta.Lang),
		meta:   meta,
	}
}

func (l *MultipartUploadInitLogicGen) initReq(req *types.MultipartUploadInitReq) error {

	if req.FileName != nil {
		l.req.FileName = strings.TrimSpace(*req.FileName)
		l.hasReq.FileName = true
	} else {
		l.hasReq.FileName = false
	}

	if req.FileSha1 != nil {
		l.req.FileSha1 = strings.TrimSpace(*req.FileSha1)
		l.hasReq.FileSha1 = true
	} else {
		l.hasReq.FileSha1 = false
	}

	if req.FileSize != nil {
		l.req.FileSize = *req.FileSize
		l.hasReq.FileSize = true
	} else {
		l.hasReq.FileSize = false
	}
	l.hasUserInfo = true
	l.mustUserInfo = true

	return nil
}
