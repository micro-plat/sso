package vcs

import (
	"github.com/lib4dev/vcs/modules/smscode"
	"github.com/lib4dev/vcs/structs"
)

//SendSmsCode 发送消息验证码 *使用前,请先配置消息发送的rpc地址 SetConfig(WithSmsSendUrl(""))
//req-->消息验证码获取实体
//实体req_id:请求id,ident:系统标识,user_account:手机号,template_id:消息模板编号,keywords:发送内容,delivery_time:定时发送时间(可空),格式：yyyy-mm-dd hh:mm:ss
//返回值result {"record_id":"xxxx"}
func SendSmsCode(req *structs.SendRequest, platName string) (result *structs.SendResult, err error) {

	obj, err := smscode.NewCode()
	if err != nil {
		return nil, err
	}
	r, err := obj.Send(req, platName)
	if err != nil {
		return nil, err
	}

	result = &structs.SendResult{}
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
