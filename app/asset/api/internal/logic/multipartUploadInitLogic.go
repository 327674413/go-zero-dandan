package logic

import (
	"context"
	"fmt"
	"go-zero-dandan/app/asset/api/internal/svc"
	"go-zero-dandan/app/asset/api/internal/types"
	"math"

	"github.com/nicksnyder/go-i18n/v2/i18n"
	"github.com/zeromicro/go-zero/core/logx"
	"go-zero-dandan/common/resd"
	"go-zero-dandan/common/utild"
)

type MultipartUploadInitLogic struct {
	logx.Logger
	ctx        context.Context
	svcCtx     *svc.ServiceContext
	lang       *i18n.Localizer
	platId     int64
	platClasEm int64
}

func NewMultipartUploadInitLogic(ctx context.Context, svcCtx *svc.ServiceContext) *MultipartUploadInitLogic {
	localizer := ctx.Value("lang").(*i18n.Localizer)
	return &MultipartUploadInitLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
		lang:   localizer,
	}
}

const ChunkSize = 5 << 20 //暂停写死5M 1个片
func (l *MultipartUploadInitLogic) MultipartUploadInit(req *types.MultipartUploadInitReq) (resp *types.MultipartUploadInitRes, err error) {
	//todo::分片上传属于大文件，应该只有网盘之类项目才会用，必然是注册用户才会用到
	resp = &types.MultipartUploadInitRes{
		FileSha1:   req.FileSha1,
		FileSize:   req.FileSize,
		UploadId:   utild.MakeId(),
		ChunkSize:  ChunkSize,
		ChunkCount: int64(math.Ceil(float64(req.FileSize) / ChunkSize)),
	}
	redisFieldKey := fmt.Sprintf("multipart:%d", resp.UploadId)
	// 将分块的信息写入redis
	if err = l.svcCtx.Redis.HsetCtx(l.ctx, redisFieldKey, "fileSha1", resp.FileSha1); err != nil {
		return l.apiFail(resd.Error(err))
	}
	if err = l.svcCtx.Redis.HsetCtx(l.ctx, redisFieldKey, "fileSize", fmt.Sprintf("%d", resp.FileSize)); err != nil {
		return l.apiFail(resd.Error(err))
	}
	if err = l.svcCtx.Redis.HsetCtx(l.ctx, redisFieldKey, "chunkCount", fmt.Sprintf("%d", resp.ChunkCount)); err != nil {
		return l.apiFail(resd.Error(err))
	}
	return resp, nil

}

func (l *MultipartUploadInitLogic) initPlat() (err error) {
	platClasEm := utild.AnyToInt64(l.ctx.Value("platClasEm"))
	if platClasEm == 0 {
		return resd.NewErrCtx(l.ctx, "token中未获取到platClasEm", resd.PlatClasErr)
	}
	platClasId := utild.AnyToInt64(l.ctx.Value("platId"))
	if platClasId == 0 {
		return resd.NewErrCtx(l.ctx, "token中未获取到platId", resd.PlatIdErr)
	}
	l.platId = platClasId
	l.platClasEm = platClasEm
	return nil
}
func (l *MultipartUploadInitLogic) apiFail(err error) (*types.MultipartUploadInitRes, error) {
	return nil, resd.ApiFail(l.lang, err)
}
