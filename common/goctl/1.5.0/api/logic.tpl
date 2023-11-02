package {{.pkgName}}

import (
	{{.imports}}
	"go-zero-dandan/common/resd"
    "go-zero-dandan/common/utild"
    "go-zero-dandan/app/user/rpc/user"
)

type {{.logic}} struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
	userMainInfo *user.UserMainInfo
	platId     int64
    platClasEm int64
}

func New{{.logic}}(ctx context.Context, svcCtx *svc.ServiceContext) *{{.logic}} {
	return &{{.logic}}{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *{{.logic}}) {{.function}}({{.request}}) {{.responseType}} {
	if err = l.initPlat(); err != nil {
    	return nil,resd.ErrorCtx(l.ctx,err)
    }

	{{.returnString}}
}

func (l *{{.logic}}) initUser() (err error) {
	userMainInfo, ok := l.ctx.Value("userMainInfo").(*user.UserMainInfo)
	if !ok {
		return resd.NewErrCtx(l.ctx,"未配置userInfo中间件", resd.UserMainInfoErr)
	}
	l.userMainInfo = userMainInfo
	return nil
}

func (l *{{.logic}}) initPlat() (err error) {
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
