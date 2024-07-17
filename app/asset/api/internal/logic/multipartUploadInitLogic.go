package logic

import (
	"context"
	"fmt"
	"go-zero-dandan/app/asset/api/internal/biz"
	"go-zero-dandan/app/asset/api/internal/svc"
	"go-zero-dandan/app/asset/api/internal/types"
	"go-zero-dandan/app/asset/model"
	"go-zero-dandan/common/constd"
	"math"
	"path/filepath"

	"go-zero-dandan/common/resd"
	"go-zero-dandan/common/utild"
)

type MultipartUploadInitLogic struct {
	*MultipartUploadInitLogicGen
}

func NewMultipartUploadInitLogic(ctx context.Context, svc *svc.ServiceContext) *MultipartUploadInitLogic {
	return &MultipartUploadInitLogic{
		MultipartUploadInitLogicGen: NewMultipartUploadInitLogicGen(ctx, svc),
	}
}

const ChunkSize = 5 << 20 //暂停写死5M 1个片
const (
	uploadInitStateNew      = 0 //新建待上传任务
	uploadInitStateContinue = 1 //已存在任务，继续
	uploadInitStateFinish   = 2 //秒传
)

func (l *MultipartUploadInitLogic) MultipartUploadInit(in *types.MultipartUploadInitReq) (resp *types.MultipartUploadInitRes, err error) {
	if err = l.initReq(in); err != nil {
		return nil, l.resd.Error(err)
	}
	netdiskModel := model.NewAssetNetdiskFileModel(l.ctx, l.svc.SqlConn, l.meta.PlatId)
	whereStr := fmt.Sprintf("sha1 = ? AND mode_em=%d", l.svc.Config.AssetMode)
	findFile, err := netdiskModel.Where(whereStr, l.req.FileSha1).Find()
	//查询报错
	if err != nil {
		return nil, l.resd.Error(err)
	}
	//有找到历史任务，并且状态为未完成
	if findFile != nil && findFile.StateEm < constd.AssetStateEmFinish {
		return l.getTask(findFile)
	}
	//没找到 或 历史任务已完成，找是否存在该文件
	assetMainModel := model.NewAssetMainModel(l.ctx, l.svc.SqlConn)
	findAsset, err := assetMainModel.Where(fmt.Sprintf("sha1= ? AND state_em=%d", constd.AssetStateEmFinish), l.req.FileSha1).Find()
	if err != nil {
		return nil, resd.Error(err)
	}
	if findAsset == nil {
		return l.addTask()
	} else {
		assetData := &model.AssetNetdiskFile{}
		assetData.Sha1 = findAsset.Sha1
		assetData.ModeEm = findAsset.ModeEm
		assetData.SizeNum = findAsset.SizeNum
		assetData.SizeText = findAsset.SizeText
		assetData.StateEm = findAsset.StateEm
		assetData.Mime = findAsset.Mime
		assetData.Ext = findAsset.Ext
		assetData.Url = findAsset.Url
		assetData.Path = findAsset.Path
		assetData.Name = l.req.FileName
		assetData.OriginalName = l.req.FileName
		assetData.FinishAt = utild.GetStamp()
		assetData.Id = utild.MakeId()
		assetData.UserId = l.meta.UserId
		assetData.AssetId = findAsset.Id
		_, err = netdiskModel.Insert(assetData)
		if err != nil {
			return nil, l.resd.Error(err)
		}
		return &types.MultipartUploadInitRes{State: uploadInitStateFinish}, nil
	}

}

func (l *MultipartUploadInitLogic) getTask(findFile *model.AssetNetdiskFile) (resp *types.MultipartUploadInitRes, err error) {
	assetBiz := biz.NewAssetBiz(l.ctx, l.svc, l.meta)
	total, completeChunks, err := assetBiz.GetUploading(findFile.Id)
	if err != nil {
		return nil, l.resd.Error(err)
	}
	if total == 0 {
		return l.addTask()
	}
	return &types.MultipartUploadInitRes{
		UserId:        l.meta.UserId,
		FileSha1:      findFile.Sha1,
		FileSize:      findFile.SizeNum,
		UploadId:      findFile.Id,
		ChunkSize:     ChunkSize,
		ChunkCount:    total,
		ChunkComplete: completeChunks,
		State:         uploadInitStateContinue,
	}, nil

}
func (l *MultipartUploadInitLogic) addTask() (resp *types.MultipartUploadInitRes, err error) {
	netdiskModel := model.NewAssetNetdiskFileModel(l.ctx, l.svc.SqlConn, l.meta.PlatId)
	fileInfo := &model.AssetNetdiskFile{
		Id:           utild.MakeId(),
		Name:         l.req.FileName,
		OriginalName: l.req.FileName,
		ModeEm:       l.svc.Config.AssetMode,
		SizeNum:      l.req.FileSize,
		SizeText:     utild.FormatFileSize(l.req.FileSize),
		Ext:          filepath.Ext(l.req.FileName),
		UserId:       l.meta.UserId,
		Sha1:         l.req.FileSha1,
	}
	if err != nil {
		return nil, l.resd.Error(err)
	}
	_, err = netdiskModel.Insert(fileInfo)
	if err != nil {
		return nil, l.resd.Error(err)
	}
	resp = &types.MultipartUploadInitRes{
		UserId:        l.meta.UserId,
		FileSha1:      l.req.FileSha1,
		FileSize:      l.req.FileSize,
		UploadId:      fileInfo.Id,
		ChunkSize:     ChunkSize,
		ChunkCount:    int64(math.Ceil(float64(l.req.FileSize) / ChunkSize)),
		ChunkComplete: []int64{},
		State:         uploadInitStateNew,
	}
	redisFieldKey := fmt.Sprintf("multipart:%d", resp.UploadId)
	// 将分块的信息写入redis
	if err = l.svc.Redis.HsetCtx(l.ctx, redisFieldKey, "fileSha1", resp.FileSha1); err != nil {
		return nil, l.resd.Error(err)
	}
	if err = l.svc.Redis.HsetCtx(l.ctx, redisFieldKey, "fileSize", fmt.Sprintf("%d", resp.FileSize)); err != nil {
		return nil, l.resd.Error(err)
	}
	if err = l.svc.Redis.HsetCtx(l.ctx, redisFieldKey, "chunkCount", fmt.Sprintf("%d", resp.ChunkCount)); err != nil {
		return nil, l.resd.Error(err)
	}
	return resp, nil
}
