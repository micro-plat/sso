package vcs

import (
	"github.com/lib4dev/vcs/modules/smscode"
	"github.com/micro-plat/lib4go/types"
)

type SendRequest struct {
	ReqID        string     `json:"req_id" m2s:"req_id" valid:"required"`
	Ident        string     `json:"ident" m2s:"ident" valid:"required"`
	UserAccount  string     `json:"user_account" m2s:"user_account" valid:"required"`
	TemplateID   string     `json:"template_id" m2s:"template_id" valid:"required"`
	Keywords     string     `json:"keywords" m2s:"keywords" valid:"required"`
	DeliveryTime string     `json:"delivery_time,omitempty" m2s:"delivery_time"` //格式：yyyy-mm-dd hh:mm:ss
	ExtParams    types.XMap `json:"ext_params,omitempty" m2s:"ext_params,omitempty"  `
}

type SendResult struct {
	RecordID string `json:"record_id"`
}

//SendSmsCode 发送消息验证码 *使用前,请先配置消息发送的rpc地址 SetConfig(WithSmsSendUrl(""))
//req-->消息验证码获取实体
//实体req_id:请求id,ident:系统标识,user_account:手机号,template_id:消息模板编号,keywords:发送内容,delivery_time:定时发送时间(可空),格式：yyyy-mm-dd hh:mm:ss
//返回值result {"record_id":"xxxx"}
func SendSmsCode(req *SendRequest, platName string) (result *SendResult, err error) {

	info := &smscode.SendRequest{
		ReqID:        req.ReqID,
		Ident:        req.Ident,
		UserAccount:  req.UserAccount,
		TemplateID:   req.TemplateID,
		Keywords:     req.Keywords,
		DeliveryTime: req.DeliveryTime,
		ExtParams:    req.ExtParams,
	}

	obj, err := smscode.NewCode()
	if err != nil {
		return nil, err
	}
	r, err := obj.Send(info, platName)
	if err != nil {
		return nil, err
	}

	result = &SendResult{}
	err = r.ToStruct(result)
	if err != nil {
		return nil, err
	}

	return
}

//VerifySmsCode 验证消息验证码
//ident-->系统标识,userAccount-->手机号,code-->验证码
func VerifySmsCode(ident, userAccount, code, platName string) (err error) {

	obj, err := smscode.NewCode()
	if err != nil {
		return err
	}

	return obj.Validate(ident, userAccount, code, platName)
}
