package logic

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/zeromicro/go-zero/core/logc"
	"github.com/zeromicro/go-zero/core/threading"
	"go-zero-dandan/app/message/rpc/internal/svc"
	"go-zero-dandan/app/message/rpc/types/pb"
	"go-zero-dandan/common/constd"
	"go-zero-dandan/common/resd"
	"go-zero-dandan/common/utild"

	"github.com/zeromicro/go-zero/core/logx"
)

type SendSMSAsyncLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewSendSMSAsyncLogic(ctx context.Context, svcCtx *svc.ServiceContext) *SendSMSAsyncLogic {
	return &SendSMSAsyncLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *SendSMSAsyncLogic) SendSMSAsync(in *pb.SendPhoneReq) (*pb.SuccResp, error) {
	if err := l.checkReq(in); err != nil {
		return nil, err
	}
	threading.GoSafe(func() {
		data, err := json.Marshal(in)
		if err != nil {
			logc.Error(l.ctx, "[SendMmsAsync]json.Marshal req error：%v", err)
			return
		}
		_, _, err = l.svcCtx.Pusher.PushCtx(l.ctx, l.svcCtx.Config.KafkaPusherConf.Topic, string(data))
		if err != nil {
			logc.Errorf(l.ctx, "[SendMmsAsync]kqPusher push error：%v", err)
			return
		}
		fmt.Println("推送成功")
		return
	})
	return &pb.SuccResp{Code: 200}, nil

}
func (l *SendSMSAsyncLogic) checkReq(in *pb.SendPhoneReq) error {
	//校验模版id
	if in.TempId == 0 {
		return resd.NewErrWithTempCtx(l.ctx, "未配置Temp Id", resd.ReqFieldRequired, "TempId")
	}
	//校验手机号
	if utild.CheckIsPhone(in.Phone) == false {
		return resd.NewErrCtx(l.ctx, "手机号格式错误", resd.ReqPhoneErr)
	}
	//校验区号
	if in.PhoneArea != "" && in.PhoneArea != constd.PhoneAreaEmChina {
		return resd.NewErrCtx(l.ctx, "不支持的手机区号", resd.NotSupportPhoneArea)
	}
	return nil
}
