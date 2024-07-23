package logic

import (
	"context"
	"database/sql"
	"fmt"
	"go-zero-dandan/app/asset/api/internal/svc"
	"go-zero-dandan/app/asset/api/internal/types"
	"go-zero-dandan/app/asset/model"
	"go-zero-dandan/common/dao"
	"go-zero-dandan/common/storaged"
	"strconv"
	"strings"

	"go-zero-dandan/common/resd"
	"go-zero-dandan/common/utild"
)

type MultipartUploadCompleteLogic struct {
	*MultipartUploadCompleteLogicGen
}

func NewMultipartUploadCompleteLogic(ctx context.Context, svc *svc.ServiceContext) *MultipartUploadCompleteLogic {
	return &MultipartUploadCompleteLogic{
		MultipartUploadCompleteLogicGen: NewMultipartUploadCompleteLogicGen(ctx, svc),
	}
}

func (l *MultipartUploadCompleteLogic) MultipartUploadComplete(in *types.MultipartUploadCompleteReq) (*types.MultipartUploadCompleteRes, error) {
	if err := l.initReq(in); err != nil {
		return nil, l.resd.Error(err)
	}
	//先判断是否存在该上传任务
	netdiskFileModel := model.NewAssetNetdiskFileModel(l.ctx, l.svc.SqlConn)
	uploadTask, err := netdiskFileModel.Ctx(l.ctx).FindById(l.req.UploadId)
	if err != nil {
		return nil, l.resd.NewErrWithTemp(resd.ErrNotFound1, resd.VarUpoladTask)
	}
	uploadKey := fmt.Sprintf("multipart:%d", uploadTask.Id)
	chunkCountStr, err := l.svc.Redis.HgetCtx(l.ctx, uploadKey, "chunkCount")
	if err != nil {
		return nil, l.resd.Error(err)
	}
	fileSha1, err := l.svc.Redis.HgetCtx(l.ctx, uploadKey, "fileSha1")
	if err != nil {
		return nil, l.resd.Error(err)
	}
	// 通过 uploadId 查询 Redis 并判断是否所有分块上传完成
	uploadInfoMap, err := l.svc.Redis.HgetallCtx(l.ctx, uploadKey)
	if err != nil {
		return nil, l.resd.Error(err)
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
		return nil, l.resd.Error(err)
	}
	// 所需分片数量不等于redis中查出来已经完成分片的数量，返回无法满足合并条件
	if chunkCount != count {
		return nil, l.resd.NewErr(resd.ErrMultipartUploadNotComplete)
	}
	// 开始合并分块
	uploader, err := l.svc.Storage.CreateUploader(&storaged.UploaderConfig{
		FileType: storaged.FileTypeMultipart,
		Bucket:   "netdisk",
	})
	if err != nil {
		return nil, l.resd.Error(err)
	}
	mergeRes, err := uploader.MultipartMerge(fileSha1, uploadTask.Name, chunkCount)
	if err != nil {
		return nil, l.resd.Error(err)
	}
	assetMainModel := model.NewAssetMainModel(l.ctx, l.svc.SqlConn)
	err = dao.WithTrans(l.ctx, l.svc.SqlConn, func(tx *sql.Tx) error {
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
			return l.resd.Error(err)
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
			return l.resd.Error(err)
		}
		_, err = l.svc.Redis.DelCtx(l.ctx, "multipart", fmt.Sprintf("%d", uploadTask.Id))
		if err != nil {
			return l.resd.Error(err)
		}
		return nil
	})

	if err != nil {
		return nil, l.resd.Error(err)
	}
	return &types.MultipartUploadCompleteRes{
		UploadId: l.req.UploadId,
	}, nil
}
