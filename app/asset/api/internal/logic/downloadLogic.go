package logic

import (
	"context"
	"go-zero-dandan/app/asset/api/internal/svc"
	"go-zero-dandan/app/asset/api/internal/types"
	"go-zero-dandan/app/asset/model"
	"net/http"

	"github.com/nicksnyder/go-i18n/v2/i18n"
	"github.com/zeromicro/go-zero/core/logx"
	"go-zero-dandan/common/resd"
	"go-zero-dandan/common/utild"
)

type DownloadLogic struct {
	logx.Logger
	ctx        context.Context
	svcCtx     *svc.ServiceContext
	lang       *i18n.Localizer
	platId     int64
	platClasEm int64
}

func NewDownloadLogic(ctx context.Context, svcCtx *svc.ServiceContext) *DownloadLogic {
	localizer := ctx.Value("lang").(*i18n.Localizer)
	return &DownloadLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
		lang:   localizer,
	}
}

func (l *DownloadLogic) Download(w http.ResponseWriter, req *types.DownloadReq, r *http.Request) (err error) {
	assetModel := model.NewAssetMainModel(l.svcCtx.SqlConn)
	asset, err := assetModel.FindById(req.Id)
	if err != nil {
		return l.apiFail(err)
	}
	/*domain := utild.GetRequestDomain(r)
	objectName := ""
	if l.svcCtx.Config.AssetMode == constd.AssetModeLocal {
		objectName = strings.Replace(asset.Url, domain+"/", "", 1)
	} else if l.svcCtx.Config.AssetMode == constd.AssetModeMinio {
		objectName = strings.Replace(asset.Url, "http://"+l.svcCtx.Config.Minio.Address+"/"+l.svcCtx.Config.Minio.Bucket, "", 1)
	} else if l.svcCtx.Config.AssetMode == constd.AssetModeTxCos {
		objectName = strings.Replace(asset.Url, l.svcCtx.Config.TxCos.PublicBucketAddr, "", 1)
	} else if l.svcCtx.Config.AssetMode == constd.AssetModeAliOss {
		objectName = strings.Replace(asset.Url, "https://danapp."+l.svcCtx.Config.AliOss.PublicBucketAddr+"/", "", 1)
	}*/
	if err != nil {
		return l.apiFail(err)
	}
	downloader, _ := l.svcCtx.Storage.CreateDownloader(nil)
	return downloader.Download(w, asset.Path, asset.Name)
}

func (l *DownloadLogic) initPlat() (err error) {
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
func (l *DownloadLogic) apiFail(err error) error {
	return resd.ApiFail(l.lang, err)
}
