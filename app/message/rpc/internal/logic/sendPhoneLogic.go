package logic

import (
	"context"
	"encoding/json"
	"fmt"
	"go-zero-dandan/app/message/model"
	"go-zero-dandan/app/message/rpc/internal/svc"
	"go-zero-dandan/app/message/rpc/types/messageRpc"
	"go-zero-dandan/common/constd"
	"go-zero-dandan/common/resd"
	"go-zero-dandan/common/utild"
	"go-zero-dandan/common/utild/smsd"
)

type SendPhoneLogic struct {
	*SendPhoneLogicGen
}

func NewSendPhoneLogic(ctx context.Context, svc *svc.ServiceContext) *SendPhoneLogic {
	return &SendPhoneLogic{
		SendPhoneLogicGen: NewSendPhoneLogicGen(ctx, svc),
	}
}

func (l *SendPhoneLogic) SendPhone(in *messageRpc.SendPhoneReq) (*messageRpc.ResultResp, error) {
	if err := l.initReq(in); err != nil {
		return nil, l.resd.Error(err)
	}
	if err := l.checkReq(); err != nil {
		return nil, l.resd.Error(err)
	}
	messageSmsSendModel := model.NewMessageSmsSendModel(l.ctx, l.svc.SqlConn)
	if err := l.checkSmsLimit(l.req.Phone, messageSmsSendModel); err != nil {
		return nil, l.resd.Error(err)
	}
	messageSmsTempModel := model.NewMessageSmsTempModel(l.ctx, l.svc.SqlConn)
	smsTemp, err := messageSmsTempModel.WhereId(l.req.TempId).CacheFind(l.svc.Redis)
	if err != nil {
		return nil, l.resd.Error(err)
	}
	if smsTemp.Id == "" {
		return nil, l.resd.NewErrWithTemp(resd.ErrConfigNotInit1, resd.VarSmsTemp)
	}
	content, _ := json.Marshal(in.TempData)
	smsSendData := &model.MessageSmsSend{
		Id:      utild.MakeId(),
		Phone:   l.req.Phone,
		TempId:  smsTemp.Id,
		PlatId:  smsTemp.PlatId,
		Content: string(content),
	}
	sms := smsd.NewSmsTencent(smsTemp.SecretId, smsTemp.SecretKey)
	err = sms.Send(l.req.Phone, smsTemp.SmsSdkAppid, smsTemp.SignName, smsTemp.TemplateId, in.TempData)
	resp := &messageRpc.ResultResp{}
	if err != nil {
		smsSendData.Err = err.Error()
		smsSendData.StateEm = -1

	} else {
		smsSendData.StateEm = 1
		resp.Code = constd.ResultFinish
	}
	_, err = messageSmsSendModel.Insert(smsSendData)
	return resp, nil

}

func (l *SendPhoneLogic) checkReq() error {
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
func (l *SendPhoneLogic) checkSmsLimit(phone string, messageSmsSendModel model.MessageSmsSendModel) error {
	//校验是否获取太频繁
	preGet, err := l.svc.Redis.Get("verifyCodeGetAt", phone)
	if err != nil {
		return l.resd.Error(err)
	}
	if preGet != "" {
		return l.resd.NewErr(resd.ErrReqGetPhoneVerifyCodeWait)
	}
	//获取系统短信配置
	messageSysConfigModel := model.NewMessageSysConfigModel(l.ctx, l.svc.SqlConn)
	sysConfig, err := messageSysConfigModel.WhereId("1").CacheFind(l.svc.Redis)
	if err != nil {
		return l.resd.Error(err)
	}
	if sysConfig.Id == "" {
		return l.resd.NewErrWithTemp(resd.ErrConfigNotInit1, resd.VarMessageSysConfig)
	}

	//校验每日上限
	if sysConfig.SmsLimitDayNum > 0 {
		todayStr := utild.Date("Y-m-d")
		todayAt := utild.StrToStamp(todayStr)
		whereStr := fmt.Sprintf("create_at > %d", todayAt)
		messageSendList, err := messageSmsSendModel.Ctx(l.ctx).Where(whereStr).Select()
		if err != nil {
			return l.resd.Error(err)
		}
		dayNum := int64(len(messageSendList))
		if dayNum >= sysConfig.SmsLimitDayNum {
			return l.resd.NewErr(resd.ErrReqGetPhoneVerifyCodeDayLimit)
		}

	}
	//校验小时上限
	if sysConfig.SmsLimitHourNum > 0 {
		hourAt := utild.GetStamp() - 3600
		whereStr := fmt.Sprintf("create_at > %d", hourAt)
		messageSendList, err := messageSmsSendModel.Ctx(l.ctx).Where(whereStr).Select()
		if err != nil {
			return l.resd.Error(err)
		}
		dayNum := int64(len(messageSendList))
		if dayNum >= sysConfig.SmsLimitHourNum {
			return l.resd.NewErr(resd.ErrReqGetPhoneVerifyCodeHourLimit)
		}
	}
	return nil
}
