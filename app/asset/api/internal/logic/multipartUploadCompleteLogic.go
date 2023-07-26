package logic

import (
	"context"
	"fmt"
	"go-zero-dandan/app/asset/api/internal/svc"
	"go-zero-dandan/app/asset/api/internal/types"
	"go-zero-dandan/common/storaged"
	"os"
	"path"
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
	uploadKey := fmt.Sprintf("multipart:%d", req.UploadId)
	chunkCountStr, err := l.svcCtx.Redis.HgetCtx(l.ctx, uploadKey, "chunkCount")
	if err != nil {
		return l.apiFail(resd.Error(err))
	}
	fileHash, err := l.svcCtx.Redis.HgetCtx(l.ctx, uploadKey, "fileSha1")
	if err != nil {
		return l.apiFail(resd.Error(err))
	}
	// 通过 uploadId 查询 Redis 并判断是否所有分块上传完成
	uploadInfoMap, err := l.svcCtx.Redis.HgetallCtx(l.ctx, uploadKey)
	if err != nil {
		return l.apiFail(resd.Error(err))
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
		return l.apiFail(resd.Error(err))
	}
	// 所需分片数量不等于redis中查出来已经完成分片的数量，返回无法满足合并条件
	if chunkCount != count {
		return l.apiFail(resd.NewErr("文件未完全上传", resd.MultipartUploadFileHashRequired))
	}
	// 开始合并分块
	// 合并后的文件路径
	mergedFilePath := l.svcCtx.Config.LocalPath + "/file/" + storaged.GetDateDir() + "/" + fileHash + ".png"
	err = os.MkdirAll(path.Dir(mergedFilePath), 0744)
	if err != nil {
		return l.apiFail(resd.Error(err))
	}

	mergedFile, err := os.Create(mergedFilePath)
	if err != nil {
		return l.apiFail(resd.Error(err))
	}
	defer mergedFile.Close()
	// 读取每个分块文件数据并加入到合并文件中
	for i := 0; i < chunkCount; i++ {
		chunkFilePath := l.svcCtx.Config.LocalPath + "/multipart/" + fileHash[:2] + "/" + fileHash + "_" + strconv.Itoa(i) // 分块文件路径
		chunkData, err := os.ReadFile(chunkFilePath)
		if err != nil {
			return l.apiFail(resd.Error(err))

		}

		_, err = mergedFile.Write(chunkData)
		if err != nil {
			return l.apiFail(resd.Error(err))
		}

		// 删除已合并的分块文件
		err = os.Remove(chunkFilePath)
		if err != nil {
			return l.apiFail(resd.Error(err))
		}
	}

	return &types.MultipartUploadCompleteRes{
		AssetId: req.UploadId,
	}, nil
}

func (l *MultipartUploadCompleteLogic) initPlat() (err error) {
	platClasEm := utild.AnyToInt64(l.ctx.Value("platClasEm"))
	if platClasEm == 0 {
		return resd.FailCode(l.lang, resd.PlatClasErr)
	}
	platClasId := utild.AnyToInt64(l.ctx.Value("platId"))
	if platClasId == 0 {
		return resd.FailCode(l.lang, resd.PlatIdErr)
	}
	l.platId = platClasId
	l.platClasEm = platClasEm
	return nil
}
func (l *MultipartUploadCompleteLogic) apiFail(err error) (*types.MultipartUploadCompleteRes, error) {
	return nil, resd.ApiFail(l.lang, err)
}
