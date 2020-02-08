package sso

import (
	"encoding/json"
	"fmt"
	"time"

	"github.com/micro-plat/hydra/context"
	"github.com/micro-plat/lib4go/net/http"
)

func remoteRequest(host, path, content string, data interface{}) (interface{}, error) {
	url := host + path
	client, err := http.NewHTTPClient(http.WithRequestTimeout(5 * time.Second))
	if err != nil {
		return nil, err
	}
	body, statusCode, err := client.Post(url, content)
	if err != nil {
		return nil, err
	}
	if statusCode != 200 {
		return nil, context.NewErrorf(statusCode, "获取apiserver信息失败, url: %s,HttpStatus:%d, body:%s", url, statusCode, body)
	}
	listByte := []byte(body)
	//err = json.Unmarshal(listByte, &data)
	err = json.Unmarshal(listByte, data)
	if err != nil {
		return nil, fmt.Errorf("字符串转json发生错误，err：%v", err)
	}

	return data, nil
}
