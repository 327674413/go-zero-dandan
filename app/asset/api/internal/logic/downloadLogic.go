package logic

import (
	"context"
	"fmt"
	"go-zero-dandan/app/asset/model"
	"net/http"
	"os"
	"strings"

	"go-zero-dandan/app/asset/api/internal/svc"
	"go-zero-dandan/app/asset/api/internal/types"

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

func (l *DownloadLogic) Download(req *types.DownloadReq, r *http.Request) (resp *types.DownloadResp, err error) {
	assetModel := model.NewAssetMainModel(l.svcCtx.SqlConn)
	asset, err := assetModel.FindById(req.Id)
	if err != nil {
		return l.apiFail(err)
	}
	domain := utild.GetRequestDomain(r)
	path := strings.Replace(asset.Url, domain+"/", "", 1)
	content, err := os.ReadFile(path)
	if err != nil {
		return l.apiFail(err)
	}
	fmt.Println(path)
	return &types.DownloadResp{
		Content:  content,
		FileName: asset.Name,
	}, nil
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
func (l *DownloadLogic) apiFail(err error) (*types.DownloadResp, error) {
	return nil, resd.ApiFail(l.lang, err)
}
