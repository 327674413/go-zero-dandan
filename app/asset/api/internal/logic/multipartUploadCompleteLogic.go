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
	platId     string
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
	netdiskFileModel := model.NewAssetNetdiskFileModel(l.ctx, l.svcCtx.SqlConn)
	uploadTask, err := netdiskFileModel.Ctx(l.ctx).FindById(req.UploadId)
	if err != nil {
		return nil, resd.NewErrWithTempCtx(l.ctx, "该分片上传id不存在", resd.NotFound1, "UpoladTask")
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
	assetMainModel := model.NewAssetMainModel(l.ctx, l.svcCtx.SqlConn)
	tx, err := dao.StartTrans(l.svcCtx.SqlConn)
	if err != nil {
		return nil, resd.Error(err)
	}
	assetId := utild.MakeId()
	assetMainData := &model.AssetMain{
		Id:       assetId,
		StateEm:  2,
		Sha1:     uploadTask.Sha1,
		Name:     uploadTask.Name,
		ModeEm:   uploadTask.ModeEm,
		Mime:     mergeRes.Mime,
		SizeNum:  uploadTask.SizeNum,
		SizeText: utild.FormatFileSize(uploadTask.SizeNum),
		Ext:      uploadTask.Ext,
		Url:      mergeRes.Url,
		Path:     mergeRes.Path,
		PlatId:   "",
		CreateAt: 0,
		UpdateAt: 0,
		DeleteAt: 0,
	}
	_, err = assetMainModel.TxInsert(tx, assetMainData)
	if err != nil {
		return nil, resd.ErrorCtx(l.ctx, err)
	}
	_, err = netdiskFileModel.TxUpdate(tx, map[dao.TableField]any{
		model.AssetNetdiskFile_Id:       uploadTask.Id,
		model.AssetNetdiskFile_StateEm:  2,
		model.AssetNetdiskFile_Mime:     uploadTask.Mime,
		model.AssetNetdiskFile_Ext:      uploadTask.Ext,
		model.AssetNetdiskFile_Url:      mergeRes.Url,
		model.AssetNetdiskFile_Path:     mergeRes.Path,
		model.AssetNetdiskFile_AssetId:  uploadTask.AssetId,
		model.AssetNetdiskFile_FinishAt: utild.GetStamp(),
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
	platId, _ := l.ctx.Value("platId").(string)
	if platId == "" {
		return resd.NewErrCtx(l.ctx, "token中未获取到platId", resd.PlatIdErr)
	}
	l.platId = platId
	l.platClasEm = platClasEm
	return nil
}
