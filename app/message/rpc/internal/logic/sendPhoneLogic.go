package logic

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/zeromicro/go-zero/core/logx"
	"go-zero-dandan/app/message/model"
	"go-zero-dandan/app/message/rpc/internal/svc"
	"go-zero-dandan/app/message/rpc/types/pb"
	"go-zero-dandan/app/user/rpc/user"
	"go-zero-dandan/common/constd"
	"go-zero-dandan/common/dao"
	"go-zero-dandan/common/resd"
	"go-zero-dandan/common/utild"
	"go-zero-dandan/common/utild/smsd"
)

type SendPhoneLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
	platId       int64
	platClasEm   int64
	userMainInfo *user.UserMainInfo
}

func NewSendPhoneLogic(ctx context.Context, svcCtx *svc.ServiceContext) *SendPhoneLogic {
	return &SendPhoneLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *SendPhoneLogic) SendPhone(in *pb.SendPhoneReq) (*pb.SuccResp, error) {

	if err := l.checkReq(in); err != nil {
		return nil, err
	}
	messageSmsSendModel := model.NewMessageSmsSendModel(l.svcCtx.SqlConn)
	if err := l.checkSmsLimit(in.Phone, messageSmsSendModel); err != nil {
		return nil, err
	}
	messageSmsTempModel := model.NewMessageSmsTempModel(l.svcCtx.SqlConn)
	smsTemp, err := messageSmsTempModel.Ctx(l.ctx).WhereId(in.TempId).CacheFind(l.svcCtx.Redis)
	if err != nil {
		return nil, resd.RpcEncodeSysErr(err.Error())
	}
	if smsTemp.Id == 0 {
		return nil, resd.RpcEncodeTempErr(resd.ConfigNotInit, []string{"SmsTemp"})
	}
	content, _ := json.Marshal(in.TempData)
	smsSendData := &model.MessageSmsSend{
		Id:      utild.MakeId(),
		Phone:   in.Phone,
		TempId:  smsTemp.Id,
		PlatId:  smsTemp.PlatId,
		Content: string(content),
	}
	sms := smsd.NewSmsTencent(smsTemp.SecretId, smsTemp.SecretKey)
	err = sms.Send(in.Phone, smsTemp.SmsSdkAppid, smsTemp.SignName, smsTemp.TemplateId, in.TempData)
	resp := &pb.SuccResp{}
	if err != nil {
		err = resd.RpcEncodeMsgErr(err.Error(), resd.TrdSmsSendErr)
		smsSendData.StateEm = -1
		smsSendData.Err = err.Error()
	} else {
		smsSendData.StateEm = 1
		resp.Code = 200
	}
	data, err := dao.PrepareData(smsSendData)
	if err != nil {

	}
	_, err = messageSmsSendModel.Insert(data)
	return resp, nil

}

func (l *SendPhoneLogic) checkReq(in *pb.SendPhoneReq) error {
	//校验模版id
	if in.TempId == 0 {
		return resd.RpcEncodeTempErr(resd.ReqFieldRequired, []string{"TempId"})
	}
	//校验手机号
	if utild.CheckIsPhone(in.Phone) == false {
		return resd.RpcEncodeTempErr(resd.ReqPhoneErr)
	}
	//校验区号
	if in.PhoneArea != "" && in.PhoneArea != constd.PhoneAreaEmChina {
		return resd.RpcEncodeTempErr(resd.NotSupportPhoneArea)
	}
	return nil
}
func (l *SendPhoneLogic) checkSmsLimit(phone string, messageSmsSendModel model.MessageSmsSendModel) error {
	//校验是否获取太频繁
	preGet, err := l.svcCtx.Redis.Get("verifyCodeGetAt", phone)
	if err != nil {
		return resd.RpcEncodeTempErr(resd.RedisGetErr)
	}
	if preGet != "" {
		return resd.RpcEncodeTempErr(resd.ReqGetPhoneVerifyCodeWait)
	}
	//获取系统短信配置
	messageSysConfigModel := model.NewMessageSysConfigModel(l.svcCtx.SqlConn)
	sysConfig, err := messageSysConfigModel.Ctx(l.ctx).WhereId(1).CacheFind(l.svcCtx.Redis)
	if err != nil {
		return resd.RpcEncodeSysErr(err.Error())
	}
	if sysConfig.Id == 0 {
		return resd.RpcEncodeTempErr(resd.ConfigNotInit, []string{"MessageSysConfig"})
	}

	//校验每日上限
	if sysConfig.SmsLimitDayNum > 0 {
		todayStr := utild.Date("Y-m-d")
		todayAt := utild.StrToStamp(todayStr)
		whereStr := fmt.Sprintf("create_at > %d", todayAt)
		messageSendList, err := messageSmsSendModel.Ctx(l.ctx).Where(whereStr).Select()
		if err != nil {
			return resd.RpcEncodeTempErr(resd.MysqlErr)
		}
		dayNum := int64(len(messageSendList))
		if dayNum >= sysConfig.SmsLimitDayNum {
			return resd.RpcEncodeTempErr(resd.ReqGetPhoneVerifyCodeDayLimit)
		}

	}
	//校验小时上限
	if sysConfig.SmsLimitHourNum > 0 {
		fmt.Println("触发了小时查询")
		hourAt := utild.GetStamp() - 3600
		whereStr := fmt.Sprintf("create_at > %d", hourAt)
		messageSendList, err := messageSmsSendModel.Ctx(l.ctx).Where(whereStr).Select()
		if err != nil {
			return resd.RpcEncodeTempErr(resd.MysqlErr)
		}
		dayNum := int64(len(messageSendList))
		if dayNum >= sysConfig.SmsLimitHourNum {
			return resd.RpcEncodeTempErr(resd.ReqGetPhoneVerifyCodeHourLimit)
		}
	}
	return nil
}
func (l *SendPhoneLogic) initUser() (err error) {
	userMainInfo, ok := l.ctx.Value("userMainInfo").(*user.UserMainInfo)
	if !ok {
		return resd.NewErr("未配置userInfo中间件", resd.UserMainInfoErr)
	}
	l.userMainInfo = userMainInfo
	return nil
}
func (l *SendPhoneLogic) initPlat() (err error) {
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
