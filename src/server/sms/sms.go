package sms

import (
	"fmt"
	openapi "github.com/alibabacloud-go/darabonba-openapi/v2/client"
	"github.com/alibabacloud-go/dysmsapi-20170525/v4/client"
	"github.com/alibabacloud-go/tea-utils/v2/service"
	"github.com/alibabacloud-go/tea/tea"
	"github.com/ihangsen/common/src/catch"
	"github.com/ihangsen/common/src/collection/dict"
	"github.com/ihangsen/common/src/i18n"
	"github.com/ihangsen/common/src/log"
	"github.com/ihangsen/common/src/res"
	"github.com/ihangsen/common/src/server/oss"
	"net/http"
	"sync"
)

type SmsConf struct {
	AliSecret      *oss.AliSecret
	SignName       string //签名
	Endpoint       string
	SmsTemplateMap *dict.Dict[uint8, SmsTemplate]
}

type SmsTemplate struct {
	ValidTime     int64
	IntervalTime  int64  //限流间隔时间
	TemplateCode  string //模板code
	TemplateParam string //模板参数
}

var (
	smsClient = new(client.Client)
	smsConf   = new(SmsConf)
	once      sync.Once
)

func Init(sms *SmsConf) {
	once.Do(func() {
		smsConf = sms
		config := &openapi.Config{
			AccessKeyId:     &sms.AliSecret.AccessKey,
			AccessKeySecret: &sms.AliSecret.SecretKey,
			Endpoint:        &sms.Endpoint,
		}
		smsClient = catch.Try1(client.NewClient(config))
	})
}

func Send(smsType uint8, smsTemplate *SmsTemplate, phoneNumber string, params ...any) {
	request := &client.SendSmsRequest{
		PhoneNumbers:  &phoneNumber,
		SignName:      &smsConf.SignName,
		TemplateCode:  &smsTemplate.TemplateCode,
		TemplateParam: tea.String(fmt.Sprintf(smsTemplate.TemplateParam, params...)),
	}
	runtime := &service.RuntimeOptions{}
	response := catch.Try1(smsClient.SendSmsWithOptions(request, runtime))
	body := response.Body
	if *response.StatusCode != http.StatusOK || *body.Code != "OK" {
		log.Zap.Error("短信发送失败:", response)
		res.Msg(i18n.Get.SendFailed)
	}
	log.Zap.Infof("发送短信成功:smsType:%d phoneNumber:%s code:%s", smsType, phoneNumber, params)
}
