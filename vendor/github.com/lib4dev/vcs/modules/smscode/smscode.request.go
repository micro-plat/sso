package smscode

import (
	"encoding/json"
	"fmt"

	"github.com/lib4dev/vcs/modules/const/conf"
	"github.com/micro-plat/hydra/components"
)

func (s *Code) SendRequest(info *SendRequest) (r *SendResult, err error) {

	b, err := json.Marshal(info)
	if err != nil {
		return
	}

	client, err := components.Def.HTTP().GetClient()
	if err != nil {
		return
	}

	url := conf.SmsCodeSetting.SmsCodeSendRequestURL
	resultVal, status, err := client.Post(url, string(b))
	if err != nil || status != 200 {
		return nil, fmt.Errorf("发送短信请求错误,status:%d,url:%s,params:%+v,err:%+v", status, url, string(b), err)
	}

	r = &SendResult{}
	if err = json.Unmarshal([]byte(resultVal), r); err != nil {
		return nil, fmt.Errorf("发送短信返回错误,resultVal:%s,err:%+v", resultVal, err)
	}

	return
}
