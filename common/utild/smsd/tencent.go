package smsd

import (
	"encoding/json"
	"fmt"
	"github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common"
	"github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common/errors"
	"github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common/profile"
	sms "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/sms/v20210111"
)

type Tencent struct {
	SecretId  string
	SecretKey string
}

func NewSmsTencent(SecretId string, SecretKey string) *Tencent {
	return &Tencent{SecretId: SecretId, SecretKey: SecretKey}
}
func (t *Tencent) Send(phoneNumer string, smsSdkAppId string, signName string, templateId string, tempData []string) error {
	// 实例化一个认证对象，入参需要传入腾讯云账户 SecretId 和 SecretKey，此处还需注意密钥对的保密
	// 代码泄露可能会导致 SecretId 和 SecretKey 泄露，并威胁账号下所有资源的安全性。以下代码示例仅供参考，建议采用更安全的方式来使用密钥，请参见：https://cloud.tencent.com/document/product/1278/85305
	// 密钥可前往官网控制台 https://console.cloud.tencent.com/cam/capi 进行获取
	credential := common.NewCredential(
		t.SecretId,
		t.SecretKey,
	)
	// 实例化一个client选项，可选的，没有特殊需求可以跳过
	cpf := profile.NewClientProfile()
	cpf.HttpProfile.Endpoint = "sms.tencentcloudapi.com"
	// 实例化要请求产品的client对象,clientProfile是可选的
	client, _ := sms.NewClient(credential, "ap-beijing", cpf)

	// 实例化一个请求对象,每个接口都会对应一个request对象
	request := sms.NewSendSmsRequest()
	request.PhoneNumberSet = common.StringPtrs([]string{phoneNumer})
	request.SmsSdkAppId = common.StringPtr(smsSdkAppId) //template_code
	request.SignName = common.StringPtr(signName)
	request.TemplateId = common.StringPtr(templateId)
	request.TemplateParamSet = common.StringPtrs(tempData)
	fmt.Println("tempData", tempData)
	// 返回的resp是一个SendSmsResponse的实例，与请求对象对应
	response, err := client.SendSms(request)
	if _, ok := err.(*errors.TencentCloudSDKError); ok {
		fmt.Printf("An API error has returned: %s", err)
		return nil
	}
	if err != nil {
		return err
	}
	// 输出json格式的字符串回包
	str := response.ToJsonString()
	resp := &smsResp{}
	err = json.Unmarshal([]byte(str), resp)
	if err != nil {
		return err
	}
	if resp.Response.SendStatusSet[0].Code == "Ok" {
		return nil
	} else {
		return NewSmsError(resp.Response.SendStatusSet[0].Message)
	}

}

type SmsError struct {
	Msg string
}

func (t *SmsError) Error() string {
	return t.Msg
}
func NewSmsError(msg string) *SmsError {
	return &SmsError{Msg: msg}
}

type smsSendStatus struct {
	SerialNo       string `json:"SerialNo"`
	PhoneNumber    string `json:"PhoneNumber"`
	Fee            int    `json:"Fee"`
	SessionContext string `json:"SessionContext"`
	Code           string `json:"Code"`
	Message        string `json:"Message"`
	IsoCode        string `json:"IsoCode"`
}

type smsResponse struct {
	SendStatusSet []smsSendStatus `json:"SendStatusSet"`
	RequestId     string          `json:"RequestId"`
}

type smsResp struct {
	Response smsResponse `json:"Response"`
}
