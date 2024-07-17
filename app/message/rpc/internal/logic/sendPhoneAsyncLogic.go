package logic

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/zeromicro/go-zero/core/logc"
	"github.com/zeromicro/go-zero/core/threading"
	"go-zero-dandan/app/message/rpc/internal/svc"
	"go-zero-dandan/app/message/rpc/types/messageRpc"
	"go-zero-dandan/common/constd"
	"go-zero-dandan/common/resd"
	"go-zero-dandan/common/utild"
)

type SendPhoneAsyncLogic struct {
	*SendPhoneAsyncLogicGen
}

func NewSendPhoneAsyncLogic(ctx context.Context, svc *svc.ServiceContext) *SendPhoneAsyncLogic {
	return &SendPhoneAsyncLogic{
		SendPhoneAsyncLogicGen: NewSendPhoneAsyncLogicGen(ctx, svc),
	}
}

func (l *SendPhoneAsyncLogic) SendPhoneAsync(in *messageRpc.SendPhoneReq) (*messageRpc.ResultResp, error) {
	if err := l.initReq(in); err != nil {
		return nil, l.resd.Error(err)
	}
	if err := l.checkReq(); err != nil {
		return nil, l.resd.Error(err)
	}
	threading.GoSafe(func() {
		data, err := json.Marshal(l.req)
		if err != nil {
			logc.Error(l.ctx, "[SendMmsAsync]json.Marshal req error：%v", err)
			return
		}
		_, _, err = l.svc.SmsPusher.PushCtx(l.ctx, l.svc.Config.KqSmsPusher.Topic, string(data))
		if err != nil {
			logc.Errorf(l.ctx, "[SendMmsAsync]kqPusher push error：%v", err)
			return
		}
		fmt.Println("推送成功")
		return
	})
	return &messageRpc.ResultResp{Code: constd.ResultFinish}, nil

}
func (l *SendPhoneAsyncLogic) checkReq() error {
	//校验模版id
	if l.req.TempId == "" {
		return l.resd.NewErrWithTemp(resd.ErrReqFieldRequired1, resd.VarSmsTemp)
	}
	//校验手机号
	if utild.CheckIsPhone(l.req.Phone) == false {
		return l.resd.NewErr(resd.ErrReqPhone)
	}
	//校验区号
	if l.req.PhoneArea != "" && l.req.PhoneArea != constd.PhoneAreaEmChina {
		return l.resd.NewErr(resd.ErrNotSupportPhoneArea)
	}
	return nil
}
