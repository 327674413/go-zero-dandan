package logic

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/zeromicro/go-zero/core/logx"
	"go-zero-dandan/app/message/model"
	"go-zero-dandan/app/message/rpc/internal/svc"
	"go-zero-dandan/app/message/rpc/types/messageRpc"
	"go-zero-dandan/app/user/rpc/user"
	"go-zero-dandan/common/constd"
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

func (l *SendPhoneLogic) SendPhone(in *messageRpc.SendPhoneReq) (*messageRpc.ResultResp, error) {

	if err := l.checkReq(in); err != nil {
		return nil, err
	}
	messageSmsSendModel := model.NewMessageSmsSendModel(l.ctx, l.svcCtx.SqlConn)
	if err := l.checkSmsLimit(in.Phone, messageSmsSendModel); err != nil {
		return nil, err
	}
	messageSmsTempModel := model.NewMessageSmsTempModel(l.ctx, l.svcCtx.SqlConn)
	smsTemp, err := messageSmsTempModel.WhereId(in.TempId).CacheFind(l.svcCtx.Redis)
	if err != nil {
		l.rpcFail(err)
	}
	if smsTemp.Id == "" {
		return nil, resd.NewErrWithTempCtx(l.ctx, "未配置短信模版", resd.ConfigNotInit1, "SmsTemp")
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

func (l *SendPhoneLogic) checkReq(in *messageRpc.SendPhoneReq) error {
	//校验模版id
	if in.TempId == "" {
		return resd.NewErrWithTempCtx(l.ctx, "未配置Temp Id", resd.ErrReqFieldRequired1, "TempId")
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
func (l *SendPhoneLogic) checkSmsLimit(phone string, messageSmsSendModel model.MessageSmsSendModel) error {
	//校验是否获取太频繁
	preGet, err := l.svcCtx.Redis.Get("verifyCodeGetAt", phone)
	if err != nil {
		return resd.ErrorCtx(l.ctx, err, resd.RedisGetErr)
	}
	if preGet != "" {
		return resd.NewErrCtx(l.ctx, "获取验证码太频繁", resd.ReqGetPhoneVerifyCodeWait)
	}
	//获取系统短信配置
	messageSysConfigModel := model.NewMessageSysConfigModel(l.ctx, l.svcCtx.SqlConn)
	sysConfig, err := messageSysConfigModel.WhereId("1").CacheFind(l.svcCtx.Redis)
	if err != nil {
		return resd.ErrorCtx(l.ctx, err)
	}
	if sysConfig.Id == "" {
		return resd.NewErrWithTempCtx(l.ctx, "参数错误", resd.ConfigNotInit1, "MessageSysConfig")
	}

	//校验每日上限
	if sysConfig.SmsLimitDayNum > 0 {
		todayStr := utild.Date("Y-m-d")
		todayAt := utild.StrToStamp(todayStr)
		whereStr := fmt.Sprintf("create_at > %d", todayAt)
		messageSendList, err := messageSmsSendModel.Ctx(l.ctx).Where(whereStr).Select()
		if err != nil {
			return resd.ErrorCtx(l.ctx, err, resd.MysqlSelectErr)
		}
		dayNum := int64(len(messageSendList))
		if dayNum >= sysConfig.SmsLimitDayNum {
			return resd.NewErrCtx(l.ctx, "短信获取超日上限", resd.ReqGetPhoneVerifyCodeDayLimit)
		}

	}
	//校验小时上限
	if sysConfig.SmsLimitHourNum > 0 {
		hourAt := utild.GetStamp() - 3600
		whereStr := fmt.Sprintf("create_at > %d", hourAt)
		messageSendList, err := messageSmsSendModel.Ctx(l.ctx).Where(whereStr).Select()
		if err != nil {
			return resd.ErrorCtx(l.ctx, err, resd.MysqlSelectErr)
		}
		dayNum := int64(len(messageSendList))
		if dayNum >= sysConfig.SmsLimitHourNum {
			return resd.NewErrCtx(l.ctx, "短信获取超小时上限", resd.ReqGetPhoneVerifyCodeHourLimit)
		}
	}
	return nil
}
func (l *SendPhoneLogic) rpcFail(err error) (*messageRpc.ResultResp, error) {
	return nil, resd.RpcErrEncode(err)
}
func (l *SendPhoneLogic) initUser() (err error) {
	userMainInfo, ok := l.ctx.Value("userMainInfo").(*user.UserMainInfo)
	if !ok {
		return resd.NewErr("未配置userInfo中间件", resd.ErrUserMainInfo)
	}
	l.userMainInfo = userMainInfo
	return nil
}
func (l *SendPhoneLogic) initPlat() (err error) {
	platClasEm := utild.AnyToInt64(l.ctx.Value("platClasEm"))
	if platClasEm == 0 {
		return resd.NewErrCtx(l.ctx, "token中未获取到platClasEm", resd.ErrPlatClas)
	}
	platClasId := utild.AnyToInt64(l.ctx.Value("platId"))
	if platClasId == 0 {
		return resd.NewErrCtx(l.ctx, "token中未获取到platId", resd.ErrPlatId)
	}
	l.platId = platClasId
	l.platClasEm = platClasEm
	return nil
}
