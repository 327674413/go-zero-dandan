package logic

import (
	"context"
	"fmt"
	"go-zero-dandan/app/asset/api/internal/svc"
	"go-zero-dandan/app/asset/api/internal/types"
	"go-zero-dandan/common/storaged"
	"net/http"

	"github.com/nicksnyder/go-i18n/v2/i18n"
	"github.com/zeromicro/go-zero/core/logx"
	"go-zero-dandan/common/resd"
	"go-zero-dandan/common/utild"
)

type MultipartUploadSendLogic struct {
	logx.Logger
	ctx        context.Context
	svcCtx     *svc.ServiceContext
	lang       *i18n.Localizer
	platId     string
	platClasEm int64
}

func NewMultipartUploadSendLogic(ctx context.Context, svcCtx *svc.ServiceContext) *MultipartUploadSendLogic {
	localizer := ctx.Value("lang").(*i18n.Localizer)
	return &MultipartUploadSendLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
		lang:   localizer,
	}
}

func (l *MultipartUploadSendLogic) MultipartUploadSend(r *http.Request, req *types.MultipartUploadSendReq) (*types.SuccessResp, error) {
	redisFieldKey := fmt.Sprintf("multipart:%d", req.UploadID)
	hasUpload, err := l.svcCtx.Redis.ExistsCtx(l.ctx, redisFieldKey)
	if err != nil {
		return nil, resd.ErrorCtx(l.ctx, err)
	}
	if hasUpload == false {
		return nil, resd.NewErrWithTempCtx(l.ctx, "该分片上传id不存在", resd.NotFound1, "UpoladTask")
	}
	fileSha1, err := l.svcCtx.Redis.HgetCtx(l.ctx, redisFieldKey, "fileSha1")
	if err != nil {
		return nil, resd.ErrorCtx(l.ctx, err)
	}
	uploader, err := l.svcCtx.Storage.CreateUploader(&storaged.UploaderConfig{FileType: storaged.FileTypeFile, Bucket: "netdisk"})
	if err != nil {
		return nil, resd.ErrorCtx(l.ctx, err)
	}
	_, err = uploader.MultipartUpload(r, &storaged.UploadConfig{IsMultipart: true, FileSha1: fileSha1, ChunkIndex: req.ChunkIndex})
	if err != nil {
		return nil, resd.ErrorCtx(l.ctx, err)
	}
	l.svcCtx.Redis.HsetCtx(l.ctx, redisFieldKey, fmt.Sprintf("chunkIndex_%d", req.ChunkIndex), "ok")
	return &types.SuccessResp{Msg: ""}, nil
}

func (l *MultipartUploadSendLogic) initPlat() (err error) {
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
