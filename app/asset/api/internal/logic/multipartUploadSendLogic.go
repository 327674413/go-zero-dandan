package logic

import (
	"context"
	"fmt"
	"go-zero-dandan/app/asset/api/internal/svc"
	"go-zero-dandan/app/asset/api/internal/types"
	"go-zero-dandan/common/storaged"
	"net/http"

	"go-zero-dandan/common/resd"
)

type MultipartUploadSendLogic struct {
	*MultipartUploadSendLogicGen
}

func NewMultipartUploadSendLogic(ctx context.Context, svc *svc.ServiceContext) *MultipartUploadSendLogic {
	return &MultipartUploadSendLogic{
		MultipartUploadSendLogicGen: NewMultipartUploadSendLogicGen(ctx, svc),
	}
}

func (l *MultipartUploadSendLogic) MultipartUploadSend(r *http.Request, in *types.MultipartUploadSendReq) (*types.SuccessResp, error) {
	if err := l.initReq(in); err != nil {
		return nil, l.resd.Error(err)
	}
	redisFieldKey := fmt.Sprintf("multipart:%d", l.req.UploadID)
	hasUpload, err := l.svc.Redis.ExistsCtx(l.ctx, redisFieldKey)
	if err != nil {
		return nil, resd.ErrorCtx(l.ctx, err)
	}
	if hasUpload == false {
		return nil, l.resd.NewErrWithTemp(resd.ErrNotFound1, resd.VarUpoladTask)
	}
	fileSha1, err := l.svc.Redis.HgetCtx(l.ctx, redisFieldKey, "fileSha1")
	if err != nil {
		return nil, l.resd.Error(err)
	}
	uploader, err := l.svc.Storage.CreateUploader(&storaged.UploaderConfig{FileType: storaged.FileTypeFile, Bucket: "netdisk"})
	if err != nil {
		return nil, l.resd.Error(err)
	}
	_, err = uploader.MultipartUpload(r, &storaged.UploadConfig{IsMultipart: true, FileSha1: fileSha1, ChunkIndex: l.req.ChunkIndex})
	if err != nil {
		return nil, l.resd.Error(err)
	}
	l.svc.Redis.HsetCtx(l.ctx, redisFieldKey, fmt.Sprintf("chunkIndex_%d", l.req.ChunkIndex), "ok")
	return &types.SuccessResp{Msg: ""}, nil
}
