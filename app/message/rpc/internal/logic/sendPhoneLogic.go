package logic

import (
	"context"
	"github.com/zeromicro/go-zero/core/logx"
	"go-zero-dandan/app/message/model"
	"go-zero-dandan/app/message/rpc/internal/svc"
	"go-zero-dandan/app/message/rpc/types/pb"
	"go-zero-dandan/common/respd"
	"go-zero-dandan/common/utild"
	"go-zero-dandan/common/utild/smsd"
)

type SendPhoneLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewSendPhoneLogic(ctx context.Context, svcCtx *svc.ServiceContext) *SendPhoneLogic {
	return &SendPhoneLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *SendPhoneLogic) SendPhone(in *pb.SendPhoneReq) (*pb.SendPhoneResp, error) {
	resp := &pb.SendPhoneResp{}
	if in.TempId == 0 {
		return nil, respd.RpcEncodeTempErr(respd.ReqFieldRequired, []string{"TempId"})
	}
	if !utild.CheckIsPhone(in.Phone) {
		return nil, respd.RpcEncodeTempErr(respd.ReqPhoneErr, []string{})
	}
	messageSmsTempModel := model.NewMessageSmsTempModel(l.svcCtx.SqlConn)
	smsTemp, err := messageSmsTempModel.CacheFind(l.ctx, l.svcCtx.Redis, in.TempId)
	if err != nil {
		return nil, respd.RpcEncodeSysErr(err.Error())
	}
	if smsTemp.Id == 0 {
		return nil, respd.RpcEncodeTempErr(respd.PlatConfigNotInit, []string{"SmsTemp"})
	}
	sms := smsd.NewSmsTencent(smsTemp.SecretId, smsTemp.SecretKey)
	err = sms.Send(in.Phone, smsTemp.SmsSdkAppid, smsTemp.SignName, smsTemp.TemplateId, in.TempData)
	if err != nil {
		return nil, respd.RpcEncodeMsgErr(err.Error(), respd.TrdSmsSendErr)
	}
	resp.Code = 200
	resp.Trade = "1111111"
	return resp, nil
}
