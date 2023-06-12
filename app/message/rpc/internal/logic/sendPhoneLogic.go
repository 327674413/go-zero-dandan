package logic

import (
	"context"
	"go-zero-dandan/common/util/smsd"

	"go-zero-dandan/app/message/rpc/internal/svc"
	"go-zero-dandan/app/message/rpc/pb"

	"github.com/zeromicro/go-zero/core/logx"
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
	sms := smsd.Tencent{
		SecretKey: "qPineEl1LOBiuEMEPQDmGEm4nHchndvx",
		SecretId:  "AKIDypeYlcDK3buxkG8XKuEIjuoSqytZmJTR",
	}
	err := sms.Send("AKIDypeYlcDK3buxkG8XKuEIjuoSqytZmJTR", "qPineEl1LOBiuEMEPQDmGEm4nHchndvx", "15267877096", []string{"5210"})
	if err != nil {
		return nil, err
	}
	resp.Trade = "1111111"
	return resp, nil
}
