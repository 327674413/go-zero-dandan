package middled

import (
	"context"
	"errors"
	"github.com/zeromicro/go-zero/core/limit"
	"github.com/zeromicro/go-zero/core/logx"
	"go-zero-dandan/app/user/rpc/user"
	"go-zero-dandan/common/ctxd"
	"go-zero-dandan/common/fmtd"
	"go-zero-dandan/common/resd"
	"go-zero-dandan/common/typed"
	"go-zero-dandan/common/utild"
	"net/http"
)

// SetCtxMeta 获取与组装基础元数据
func SetCtxMeta(r *http.Request) context.Context {
	ctx := r.Context()
	meta := &typed.ReqMeta{}
	platClasEm := utild.AnyToInt64(ctx.Value("platClasEm"))
	platId, _ := ctx.Value("platId").(string)
	meta.Lang = r.FormValue("lang")
	meta.PlatClasIm = platClasEm
	meta.PlatId = platId
	ctx = context.WithValue(ctx, ctxd.KeyReqMeta, meta)
	return ctx
}

// SetCtxUser 设置用户信息
func SetCtxUser(r *http.Request, userRpc user.User) context.Context {
	userToken := r.Header.Get("Token")
	ctx := r.Context()
	meta, _ := ctx.Value(ctxd.KeyReqMeta).(*typed.ReqMeta)
	if meta == nil {
		meta = &typed.ReqMeta{}
	}
	var err error
	if userToken != "" && userRpc != nil {
		//传了token
		userInfo := &user.UserMainInfo{}
		userInfo, err = userRpc.GetUserByToken(r.Context(), &user.TokenReq{Token: &userToken})
		// 存在报错
		if err != nil && !resd.IsUserNotLoginErr(err) {
			userInfo = &user.UserMainInfo{}
			meta.UserErr = err.Error()
		}
		meta.UserId = userInfo.Id

	}
	ctx = context.WithValue(ctx, ctxd.KeyReqMeta, meta)
	return ctx
}

func ReqRateLimit(r *http.Request, limiter *limit.PeriodLimit) error {
	userToken := r.Header.Get("Token")
	// 如果没有用户 ID，不进行限流
	if userToken == "" {
		return nil
	}

	// 限流
	v, err := limiter.Take(userToken)
	fmtd.Info(v)
	if err != nil {
		logx.Errorf("take limit failed: %v", err)
		return err
	}
	// 如果超过限流次数，返回 429
	if v != limit.Allowed {
		logx.Errorf("rate limit exceeded")
		return errors.New("rate limit exceeded")
	}
	return nil
}