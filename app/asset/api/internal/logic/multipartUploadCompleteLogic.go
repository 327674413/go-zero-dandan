package logic

import (
	"context"
	"fmt"
	"go-zero-dandan/app/asset/api/internal/svc"
	"go-zero-dandan/app/asset/api/internal/types"
	"go-zero-dandan/app/asset/model"
	"go-zero-dandan/common/dao"
	"go-zero-dandan/common/storaged"
	"strconv"
	"strings"

	"github.com/nicksnyder/go-i18n/v2/i18n"
	"github.com/zeromicro/go-zero/core/logx"
	"go-zero-dandan/common/resd"
	"go-zero-dandan/common/utild"
)

type MultipartUploadCompleteLogic struct {
	logx.Logger
	ctx        context.Context
	svcCtx     *svc.ServiceContext
	lang       *i18n.Localizer
	platId     int64
	platClasEm int64
}

func NewMultipartUploadCompleteLogic(ctx context.Context, svcCtx *svc.ServiceContext) *MultipartUploadCompleteLogic {
	localizer := ctx.Value("lang").(*i18n.Localizer)
	return &MultipartUploadCompleteLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
		lang:   localizer,
	}
}

func (l *MultipartUploadCompleteLogic) MultipartUploadComplete(req *types.MultipartUploadCompleteReq) (*types.MultipartUploadCompleteRes, error) {
	//先判断是否存在该上传任务
	netdiskFileModel := model.NewAssetNetdiskFileModel(l.svcCtx.SqlConn)
	uploadTask, err := netdiskFileModel.Ctx(l.ctx).FindById(req.UploadId)
	if err != nil {
		return nil, resd.Error(err, resd.NotFound)
	}
	uploadKey := fmt.Sprintf("multipart:%d", uploadTask.Id)
	chunkCountStr, err := l.svcCtx.Redis.HgetCtx(l.ctx, uploadKey, "chunkCount")
	if err != nil {
		return nil, resd.ErrorCtx(l.ctx, err)
	}
	fileSha1, err := l.svcCtx.Redis.HgetCtx(l.ctx, uploadKey, "fileSha1")
	if err != nil {
		return nil, resd.ErrorCtx(l.ctx, err)
	}
	// 通过 uploadId 查询 Redis 并判断是否所有分块上传完成
	uploadInfoMap, err := l.svcCtx.Redis.HgetallCtx(l.ctx, uploadKey)
	if err != nil {
		return nil, resd.ErrorCtx(l.ctx, err)
	}
	count := 0
	// 遍历map
	for k, v := range uploadInfoMap {
		// 检测k是否以"chunk_"为前缀并且v为"1"
		if strings.HasPrefix(k, "chunkIndex_") && v == "ok" {
			count++
		}
	}
	chunkCount, err := strconv.Atoi(chunkCountStr)
	if err != nil {
		return nil, resd.ErrorCtx(l.ctx, err)
	}
	// 所需分片数量不等于redis中查出来已经完成分片的数量，返回无法满足合并条件
	if chunkCount != count {
		return nil, resd.NewErrCtx(l.ctx, "文件未完全上传", resd.MultipartUploadNotComplete)
	}
	// 开始合并分块
	uploader, err := l.svcCtx.Storage.CreateUploader(&storaged.UploaderConfig{
		FileType: storaged.FileTypeMultipart,
		Bucket:   "netdisk",
	})
	if err != nil {
		return nil, resd.ErrorCtx(l.ctx, err)
	}
	mergeRes, err := uploader.MultipartMerge(fileSha1, uploadTask.Name, chunkCount)
	if err != nil {
		return nil, resd.ErrorCtx(l.ctx, err)
	}
	assetMainModel := model.NewAssetMainModel(l.svcCtx.SqlConn)
	tx, err := dao.StartTrans(l.svcCtx.SqlConn)
	if err != nil {
		return nil, resd.Error(err)
	}
	assetId := utild.MakeId()
	_, err = assetMainModel.TxInsert(tx, map[string]string{
		"id":        fmt.Sprintf("%d", assetId),
		"sha1":      uploadTask.Sha1,
		"name":      uploadTask.Name,
		"mode_em":   fmt.Sprintf("%d", uploadTask.ModeEm),
		"size_num":  fmt.Sprintf("%d", uploadTask.SizeNum),
		"size_text": utild.FormatFileSize(uploadTask.SizeNum),
		"state_em":  "2",
		"mime":      mergeRes.Mime,
		"ext":       uploadTask.Ext,
		"url":       mergeRes.Url,
		"path":      mergeRes.Path,
	})
	if err != nil {
		return nil, resd.ErrorCtx(l.ctx, err)
	}
	_, err = netdiskFileModel.TxUpdate(tx, map[string]string{
		"id":        fmt.Sprintf("%d", uploadTask.Id),
		"state_em":  "2",
		"mime":      uploadTask.Mime,
		"ext":       uploadTask.Ext,
		"url":       mergeRes.Url,
		"path":      mergeRes.Path,
		"asset_id":  fmt.Sprintf("%d", assetId),
		"finish_at": fmt.Sprintf("%d", utild.GetStamp()),
	})
	if err != nil {
		return nil, resd.ErrorCtx(l.ctx, err)
	}
	_, err = l.svcCtx.Redis.DelCtx(l.ctx, "multipart", fmt.Sprintf("%d", uploadTask.Id))
	if err != nil {
		return nil, resd.ErrorCtx(l.ctx, err)
	}
	err = dao.Commit(tx)
	if err != nil {
		return nil, resd.ErrorCtx(l.ctx, err)
	}
	return &types.MultipartUploadCompleteRes{
		UploadId: req.UploadId,
	}, nil
}

func (l *MultipartUploadCompleteLogic) initPlat() (err error) {
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
