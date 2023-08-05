package logic

import (
	"context"
	"github.com/nicksnyder/go-i18n/v2/i18n"
	"github.com/zeromicro/go-zero/core/logx"
	"go-zero-dandan/app/asset/api/internal/svc"
	"go-zero-dandan/app/asset/api/internal/types"
	"go-zero-dandan/app/asset/model"
	"go-zero-dandan/common/constd"
	"go-zero-dandan/common/dao"
	"go-zero-dandan/common/resd"
	"go-zero-dandan/common/storaged"
	"go-zero-dandan/common/utild"
	"net/http"
)

type UploadImgLogic struct {
	logx.Logger
	ctx        context.Context
	svcCtx     *svc.ServiceContext
	lang       *i18n.Localizer
	platId     int64
	platClasEm int64
	fileData   *model.AssetMain
}

func NewUploadImgLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UploadImgLogic {
	return &UploadImgLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *UploadImgLogic) UploadImg(r *http.Request, req *types.UploadImgReq) (resp *types.UploadResp, err error) {
	if err = l.initPlat(); err != nil {
		return nil, resd.ErrorCtx(l.ctx, err)
	}
	uploader, err := l.svcCtx.Storage.CreateUploader(&storaged.UploaderConfig{FileType: storaged.FileTypeImage})
	if err != nil {
		return nil, resd.ErrorCtx(l.ctx, err)
	}
	res, err := uploader.UploadImg(r, &storaged.UploadImgConfig{
		/*WatermarkConfig: &imgd.WatermarkConfig{
			Type:     imgd.WatermarkTypeImg,
			Path:     "public/water_kkzhw.png",
			Position: imgd.WatermarkPositionContain,
		},*/
	})
	if err != nil {
		return nil, resd.ErrorCtx(l.ctx, err)
	}
	assetMainData := &model.AssetMain{
		Id:       utild.MakeId(),
		StateEm:  constd.AssetStateEmFinish,
		Sha1:     res.Hash,
		Name:     res.Name,
		ModeEm:   l.svcCtx.Config.AssetMode,
		Mime:     res.Mime,
		SizeNum:  res.SizeByte,
		SizeText: res.SizeText,
		Ext:      res.Ext,
		Url:      res.Url,
		Path:     res.Path,
	}
	data, err := dao.PrepareData(assetMainData)
	if err != nil {
		return nil, resd.Error(err)
	}
	assetMainModel := model.NewAssetMainModel(l.svcCtx.SqlConn)
	_, err = assetMainModel.Insert(data)
	if err != nil {
		return nil, resd.Error(err)
	}
	return &types.UploadResp{
		Url:      res.Url,
		FileName: res.Name,
	}, nil
}

func (l *UploadImgLogic) initPlat() (err error) {
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
