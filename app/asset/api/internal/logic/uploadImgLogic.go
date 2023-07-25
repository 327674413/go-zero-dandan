package logic

import (
	"context"
	"fmt"
	"github.com/nicksnyder/go-i18n/v2/i18n"
	"github.com/zeromicro/go-zero/core/logc"
	"github.com/zeromicro/go-zero/core/logx"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
	"go-zero-dandan/app/asset/api/internal/svc"
	"go-zero-dandan/app/asset/api/internal/types"
	"go-zero-dandan/app/asset/model"
	"go-zero-dandan/common/constd"
	"go-zero-dandan/common/dao"
	"go-zero-dandan/common/imgd"
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
	localizer := ctx.Value("lang").(*i18n.Localizer)
	l := &UploadImgLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
		lang:   localizer,
	}
	l.initPlat()
	return l
}

const maxMemorySize = 20 << 20 // 20 MB
const maxImgSize = 5 << 20

func (l *UploadImgLogic) UploadImg(r *http.Request, req *types.UploadImgReq) (resp *types.UploadResp, err error) {
	uploader, err := l.svcCtx.Storage.CreateUploader(&storaged.UploaderConfig{FileType: storaged.FileTypeImage})
	if err != nil {
		return nil, resd.ApiFail(l.lang, err)
	}
	hash, err := uploader.GetHash(r, "img")
	if err != nil {
		return l.apiFail(err)
	}
	//检查是否已经存在
	assetMainModel := model.NewAssetMainModel(l.svcCtx.SqlConn)
	whereStr := fmt.Sprintf("hash='%s' AND state_em > 1 AND mode_em=%d", hash, l.svcCtx.Config.AssetMode)
	find, err := assetMainModel.WhereStr(whereStr).Find()
	//存在，则秒传
	if err == nil {
		return &types.UploadResp{
			Url:      find.Url,
			FileName: find.Name,
		}, nil
	}
	//查询报错
	if err != sqlx.ErrNotFound {
		return l.apiFail(resd.Error(err))
	}
	//不存在，则上传
	res, err := uploader.UploadImg(r, &storaged.UploadImgConfig{
		WatermarkConfig: &imgd.WatermarkConfig{
			Type:     imgd.WatermarkTypeImg,
			Path:     "public/water_kkzhw.png",
			Position: imgd.WatermarkPositionContain,
		},
	})
	if err != nil {
		return l.apiFail(err)
	}
	assetMainData := &model.AssetMain{
		Id:       utild.MakeId(),
		StateEm:  constd.AssetStateEmFinish,
		Hash:     res.Hash,
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
	_, err = assetMainModel.Insert(data)
	if err != nil {
		return l.apiFail(err)
	}
	return &types.UploadResp{
		Url:      res.Url,
		FileName: res.Name,
	}, nil
	return resp, nil

}

func (l *UploadImgLogic) initPlat() (err error) {
	platClasEm := utild.AnyToInt64(l.ctx.Value("platClasEm"))
	if platClasEm == 0 {
		logc.Error(l.ctx, "token not get platClasEm")
		return resd.FailCode(l.lang, resd.PlatClasErr)
	}
	platClasId := utild.AnyToInt64(l.ctx.Value("platId"))
	if platClasId == 0 {
		logc.Error(l.ctx, "token not get platId")
		return resd.FailCode(l.lang, resd.PlatIdErr)
	}
	l.platId = platClasId
	l.platClasEm = platClasEm
	return nil
}

func (l *UploadImgLogic) apiFail(err error) (*types.UploadResp, error) {
	return nil, resd.ApiFail(l.lang, err)
}
