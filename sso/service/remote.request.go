package service

import (
	"encoding/json"
	"fmt"
	"time"

	"github.com/micro-plat/lib4go/net/http"
)

func remoteRequest(host, path, content string, data interface{}) (interface{}, error) {
	url := host + path
	client := http.NewHTTPClient(http.WithRequestTimeout(5 * time.Second))
	body, statusCode, err := client.Post(url, content)
	if err != nil {
		return nil, err
	}
	if statusCode != 200 {
		return nil, fmt.Errorf("读取系统信息失败,HttpStatus:%d, body:%s", statusCode, body)
	}

	listByte := []byte(body)
	err = json.Unmarshal(listByte, &data)
	if err != nil {
		return nil, fmt.Errorf("字符串转json发生错误，err：%v", err)
	}

	return data, nil
}
