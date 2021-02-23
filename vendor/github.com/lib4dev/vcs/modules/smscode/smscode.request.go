package smscode

import (
	"encoding/json"
	"fmt"
	"net/http"
	"github.com/lib4dev/vcs/modules/const/conf"
	"github.com/micro-plat/hydra/components"
)

func (s *Code) SendRequest(info *SendRequest) (r *SendResult, err error) {

	b, err := json.Marshal(info)
	if err != nil {
		err = fmt.Errorf("Smscode.SendRequest.Marshal:%v;%v", info, err)
		return
	}

	client, err := components.Def.HTTP().GetClient(conf.HTTPName)
	if err != nil {
		err = fmt.Errorf("Smscode.SendRequest.GetClient:%s;%v", conf.HTTPName, err)
		return
	}

	url := conf.SmsCodeSetting.SmsCodeSendRequestURL
	resultVal, status, err := client.Request(http.MethodPost, url, string(b), "UTF-8", http.Header{
		"Content-Type": []string{"application/json"},
	})
	if err != nil || status != 200 {
		return nil, fmt.Errorf("发送消息请求错误,status:%d,url:%s,params:%+v,err:%+v", status, url, string(b), err)
	}

	r = &SendResult{}
	if err = json.Unmarshal([]byte(resultVal), r); err != nil {
		return nil, fmt.Errorf("发送消息返回错误,resultVal:%s,err:%+v", resultVal, err)
	}

	return
}
