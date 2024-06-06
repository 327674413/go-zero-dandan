package logic

import (
	"context"
	"fmt"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
	"go-zero-dandan/app/asset/api/internal/biz"
	"go-zero-dandan/app/asset/api/internal/svc"
	"go-zero-dandan/app/asset/api/internal/types"
	"go-zero-dandan/app/asset/model"
	"go-zero-dandan/app/user/rpc/user"
	"go-zero-dandan/common/constd"
	"math"
	"path/filepath"

	"github.com/nicksnyder/go-i18n/v2/i18n"
	"github.com/zeromicro/go-zero/core/logx"
	"go-zero-dandan/common/resd"
	"go-zero-dandan/common/utild"
)

type MultipartUploadInitLogic struct {
	logx.Logger
	ctx          context.Context
	svcCtx       *svc.ServiceContext
	lang         *i18n.Localizer
	userMainInfo *user.UserMainInfo
	platId       int64
	platClasEm   int64
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
const (
	uploadInitStateNew      = 0 //新建待上传任务
	uploadInitStateContinue = 1 //已存在任务，继续
	uploadInitStateFinish   = 2 //秒传
)

func (l *MultipartUploadInitLogic) MultipartUploadInit(req *types.MultipartUploadInitReq) (resp *types.MultipartUploadInitRes, err error) {
	if err = l.initPlat(); err != nil {
		return nil, resd.ErrorCtx(l.ctx, err)
	}
	if err = l.initUser(); err != nil {
		return nil, resd.ErrorCtx(l.ctx, err)
	}
	netdiskModel := model.NewAssetNetdiskFileModel(l.svcCtx.SqlConn, l.platId)
	whereStr := fmt.Sprintf("sha1 = ? AND mode_em=%d", l.svcCtx.Config.AssetMode)
	findFile, err := netdiskModel.Where(whereStr, req.FileSha1).Find()
	//查询报错
	if err != nil && err != sqlx.ErrNotFound {
		return nil, resd.ErrorCtx(l.ctx, err)
	}
	//有找到历史任务，并且状态为未完成
	if err != sqlx.ErrNotFound && findFile.StateEm < constd.AssetStateEmFinish {
		return l.getTask(findFile, req)
	}
	//没找到 或 历史任务已完成，找是否存在该文件
	assetMainModel := model.NewAssetMainModel(l.svcCtx.SqlConn)
	findAsset, err := assetMainModel.Where(fmt.Sprintf("sha1= ? AND state_em=%d", constd.AssetStateEmFinish), req.FileSha1).Find()
	if err != nil && err != sqlx.ErrNotFound {
		return nil, resd.Error(err)
	}
	if err == sqlx.ErrNotFound {
		return l.addTask(req)
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
		assetData.Name = req.FileName
		assetData.OriginalName = req.FileName
		assetData.FinishAt = utild.GetStamp()
		assetData.Id = utild.MakeId()
		assetData.UserId = l.userMainInfo.Id
		assetData.AssetId = findAsset.Id
		_, err = netdiskModel.Insert(assetData)
		if err != nil {
			return nil, resd.ErrorCtx(l.ctx, err)
		}
		return &types.MultipartUploadInitRes{State: uploadInitStateFinish}, nil
	}

}

func (l *MultipartUploadInitLogic) getTask(findFile *model.AssetNetdiskFile, req *types.MultipartUploadInitReq) (resp *types.MultipartUploadInitRes, err error) {
	assetBiz := biz.NewAssetBiz(l.ctx, l.svcCtx, l.userMainInfo)
	total, completeChunks, err := assetBiz.GetUploading(findFile.Id)
	if err != nil {
		return nil, resd.ErrorCtx(l.ctx, err)
	}
	if total == 0 {
		return l.addTask(req)
	}
	return &types.MultipartUploadInitRes{
		UserId:        l.userMainInfo.Id,
		FileSha1:      findFile.Sha1,
		FileSize:      findFile.SizeNum,
		UploadId:      findFile.Id,
		ChunkSize:     ChunkSize,
		ChunkCount:    total,
		ChunkComplete: completeChunks,
		State:         uploadInitStateContinue,
	}, nil

}
func (l *MultipartUploadInitLogic) addTask(req *types.MultipartUploadInitReq) (resp *types.MultipartUploadInitRes, err error) {
	netdiskModel := model.NewAssetNetdiskFileModel(l.svcCtx.SqlConn, l.platId)
	fileInfo := &model.AssetNetdiskFile{
		Id:           utild.MakeId(),
		Name:         req.FileName,
		OriginalName: req.FileName,
		ModeEm:       l.svcCtx.Config.AssetMode,
		SizeNum:      req.FileSize,
		SizeText:     utild.FormatFileSize(req.FileSize),
		Ext:          filepath.Ext(req.FileName),
		UserId:       l.userMainInfo.Id,
		Sha1:         req.FileSha1,
	}
	if err != nil {
		return nil, resd.ErrorCtx(l.ctx, err)
	}
	_, err = netdiskModel.Insert(fileInfo)
	if err != nil {
		return nil, resd.ErrorCtx(l.ctx, err)
	}
	resp = &types.MultipartUploadInitRes{
		UserId:        l.userMainInfo.Id,
		FileSha1:      req.FileSha1,
		FileSize:      req.FileSize,
		UploadId:      fileInfo.Id,
		ChunkSize:     ChunkSize,
		ChunkCount:    int64(math.Ceil(float64(req.FileSize) / ChunkSize)),
		ChunkComplete: []int64{},
		State:         uploadInitStateNew,
	}
	redisFieldKey := fmt.Sprintf("multipart:%d", resp.UploadId)
	// 将分块的信息写入redis
	if err = l.svcCtx.Redis.HsetCtx(l.ctx, redisFieldKey, "fileSha1", resp.FileSha1); err != nil {
		return nil, resd.ErrorCtx(l.ctx, err)
	}
	if err = l.svcCtx.Redis.HsetCtx(l.ctx, redisFieldKey, "fileSize", fmt.Sprintf("%d", resp.FileSize)); err != nil {
		return nil, resd.ErrorCtx(l.ctx, err)
	}
	if err = l.svcCtx.Redis.HsetCtx(l.ctx, redisFieldKey, "chunkCount", fmt.Sprintf("%d", resp.ChunkCount)); err != nil {
		return nil, resd.ErrorCtx(l.ctx, err)
	}
	return resp, nil
}

func (l *MultipartUploadInitLogic) initUser() (err error) {
	userMainInfo, ok := l.ctx.Value("userMainInfo").(*user.UserMainInfo)
	if !ok {
		return resd.NewErr("未配置userInfo中间件", resd.UserMainInfoErr)
	}
	l.userMainInfo = userMainInfo
	return nil
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
