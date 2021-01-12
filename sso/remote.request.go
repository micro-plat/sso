package sso

import (
	"encoding/json"
	"fmt"

	"github.com/micro-plat/hydra/components"
	"github.com/micro-plat/lib4go/errs"
)

func remoteRequest(host, path, content string, data interface{}) (interface{}, error) {
	url := host + path
	client := components.Def.HTTP().GetRegularClient()
	body, statusCode, err := client.Post(url, content)
	if err != nil {
		return nil, err
	}
	if statusCode != 200 {
		return nil, errs.NewErrorf(statusCode, "获取apiserver信息失败, url: %s,HttpStatus:%d, body:%s", url, statusCode, body)
	}
	listByte := []byte(body)
	err = json.Unmarshal(listByte, data)
	if err != nil {
		return nil, fmt.Errorf("字符串转json发生错误，err：%v;org:%s;%s", err, string(listByte), url)
	}

	return data, nil
}
