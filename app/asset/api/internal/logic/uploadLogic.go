package logic

import (
	"context"

	"go-zero-dandan/app/asset/api/internal/svc"
	"go-zero-dandan/app/asset/api/internal/types"
)

type UploadLogic struct {
	*UploadLogicGen
}

func NewUploadLogic(ctx context.Context, svc *svc.ServiceContext) *UploadLogic {
	return &UploadLogic{
		UploadLogicGen: NewUploadLogicGen(ctx, svc),
	}
}

func (l *UploadLogic) Upload() (resp *types.UploadResp, err error) {

	return

}
