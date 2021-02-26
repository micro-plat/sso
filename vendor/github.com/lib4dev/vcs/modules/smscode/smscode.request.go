package smscode

import (
	"encoding/json"
	"fmt"

	"github.com/lib4dev/vcs/modules/const/conf"
	"github.com/lib4dev/vcs/modules/util"
	"github.com/lib4dev/vcs/structs"
	"github.com/micro-plat/lib4go/types"
)

//SendRequest SendRequest
func (s *Code) SendRequest(info *structs.SendRequest) (r *structs.SendResult, err error) {

	b, err := json.Marshal(info)
	if err != nil {
		err = fmt.Errorf("Smscode.SendRequest.Marshal:%v;%v", info, err)
		return
	}
	params, err := types.NewXMapByJSON(string(b))
	if err != nil {
		err = fmt.Errorf("Smscode.NewXMapByJSON:%s;%v", string(b), err)
		return
	}
	url := conf.SmsCodeSetting.SmsCodeSendRequestURL
	response, err := s.invokeRemote(url, params)
	if err != nil {
		err = fmt.Errorf("Smscode.invokeRemote:%s;%v;%v", url, params, err)
		return
	}
	r = &structs.SendResult{}
	if err = response.ToAnyStruct(r); err != nil {
		return nil, fmt.Errorf("发送消息返回错误,resultVal:%s,err:%+v", response, err)
	}

	return
}

func (s *Code) invokeRemote(url string, data types.XMap) (response types.XMap, err error) {
	if conf.SmsCodeSetting.IsHTTP() {
		response, err = util.HttpRequest(url, data)
		return
	}
	response, err = util.RpcRequest(url, data)
	return
}
