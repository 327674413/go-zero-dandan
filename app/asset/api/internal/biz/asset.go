package biz

import (
	"context"
	"fmt"
	"github.com/zeromicro/go-zero/core/stores/redis"
	"go-zero-dandan/app/asset/api/internal/svc"
	"go-zero-dandan/common/resd"
	"go-zero-dandan/common/typed"
	"go-zero-dandan/common/utild"
	"strings"
)

type AssetBiz struct {
	ctx  context.Context
	svc  *svc.ServiceContext
	meta *typed.ReqMeta
}

func NewAssetBiz(ctx context.Context, svc *svc.ServiceContext, meta *typed.ReqMeta) *AssetBiz {
	return &AssetBiz{
		ctx:  ctx,
		svc:  svc,
		meta: meta,
	}
}
func (t *AssetBiz) GetUploading(uploadId string) (int64, []int64, error) {
	//先判断是否存在该上传任务
	uploadKey := getUploadRedisKey(uploadId)

	chunkCountStr, err := t.svc.Redis.HgetCtx(t.ctx, uploadKey, "chunkCount")
	if err != nil && err != redis.Nil {
		return 0, nil, resd.Error(err)
	}
	if chunkCountStr == "" {
		return 0, nil, nil
	}
	_, err = t.svc.Redis.HgetCtx(t.ctx, uploadKey, "fileSha1")
	if err != nil {
		return 0, nil, resd.Error(err)
	}
	// 通过 uploadId 查询 Redis 并判断是否所有分块上传完成
	uploadInfoMap, err := t.svc.Redis.HgetallCtx(t.ctx, uploadKey)
	if err != nil {
		return 0, nil, resd.Error(err)
	}
	completeIndexs := make([]int64, 0)
	// 遍历map
	for k, v := range uploadInfoMap {
		// 检测k是否以"chunkIndex_"为前缀,存在则完成，值为索引
		if strings.HasPrefix(k, "chunkIndex_") {
			completeIndexs = append(completeIndexs, utild.AnyToInt64(v))
		}
	}
	if len(completeIndexs) == 0 {
		completeIndexs = []int64{}
	}
	return utild.AnyToInt64(chunkCountStr), completeIndexs, nil
}
func getUploadRedisKey(uploadId string) string {
	return fmt.Sprintf("multipart:%s", uploadId)
}
