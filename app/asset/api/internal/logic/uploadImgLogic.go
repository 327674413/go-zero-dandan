package logic

import (
	"context"
	"go-zero-dandan/app/asset/api/internal/svc"
	"go-zero-dandan/app/asset/api/internal/types"
	"go-zero-dandan/app/asset/model"
	"go-zero-dandan/common/constd"
	"go-zero-dandan/common/storaged"
	"go-zero-dandan/common/utild"
	"net/http"
)

type UploadImgLogic struct {
	*UploadImgLogicGen
}

func NewUploadImgLogic(ctx context.Context, svc *svc.ServiceContext) *UploadImgLogic {
	return &UploadImgLogic{
		UploadImgLogicGen: NewUploadImgLogicGen(ctx, svc),
	}
}

func (l *UploadImgLogic) UploadImg(r *http.Request, in *types.UploadImgReq) (resp *types.UploadResp, err error) {
	if err = l.initReq(in); err != nil {
		return nil, l.resd.Error(err)
	}
	uploader, err := l.svc.Storage.CreateUploader(&storaged.UploaderConfig{FileType: storaged.FileTypeImage})
	if err != nil {
		return nil, l.resd.Error(err)
	}
	res, err := uploader.UploadImg(r, &storaged.UploadImgConfig{
		/*WatermarkConfig: &imgd.WatermarkConfig{
			Type:     imgd.WatermarkTypeImg,
			Path:     "public/water_kkzhw.png",
			Position: imgd.WatermarkPositionContain,
		},*/
	})
	if err != nil {
		return nil, l.resd.Error(err)
	}
	assetMainData := &model.AssetMain{
		Id:       utild.MakeId(),
		StateEm:  constd.AssetStateEmFinish,
		Sha1:     res.Hash,
		Name:     res.Name,
		ModeEm:   l.svc.Config.AssetMode,
		Mime:     res.Mime,
		SizeNum:  res.SizeByte,
		SizeText: res.SizeText,
		Ext:      res.Ext,
		Url:      res.Url,
		Path:     res.Path,
	}
	assetMainModel := model.NewAssetMainModel(l.ctx, l.svc.SqlConn)
	_, err = assetMainModel.Insert(assetMainData)
	if err != nil {
		return nil, l.resd.Error(err)
	}
	return &types.UploadResp{
		Url:      res.Url,
		FileName: res.Name,
	}, nil
}
