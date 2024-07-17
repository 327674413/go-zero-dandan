package logic

import (
	"context"
	"go-zero-dandan/app/asset/api/internal/svc"
	"go-zero-dandan/app/asset/api/internal/types"
	"go-zero-dandan/app/asset/model"
	"net/http"

	"go-zero-dandan/common/resd"
)

type DownloadLogic struct {
	*DownloadLogicGen
}

func NewDownloadLogic(ctx context.Context, svc *svc.ServiceContext) *DownloadLogic {
	return &DownloadLogic{
		DownloadLogicGen: NewDownloadLogicGen(ctx, svc),
	}
}

func (l *DownloadLogic) Download(w http.ResponseWriter, in *types.DownloadReq, r *http.Request) error {
	if err := l.initReq(in); err != nil {
		return l.resd.Error(err)
	}
	assetModel := model.NewAssetMainModel(l.ctx, l.svc.SqlConn)
	asset, err := assetModel.FindById(l.req.Id)
	if err != nil {
		return resd.ErrorCtx(l.ctx, err)
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
		return resd.ErrorCtx(l.ctx, err)
	}
	downloader, _ := l.svc.Storage.CreateDownloader(nil)
	return downloader.Download(w, asset.Path, asset.Name)
}
