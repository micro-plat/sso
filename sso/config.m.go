package sso

import (
	"fmt"
	"time"

	"github.com/asaskevich/govalidator"
)

//ConfigData 配置信息
type ConfigData struct {
	//ApiHost ssoApi地址(不是跳转地址)
	host string `valid:"required"`

	/*系统标识*/
	ident string `valid:"required"`

	/*系统秘钥*/
	secret string `valid:"required"`
}

//Valid 验证传入参数
func (c *ConfigData) Valid() error {
	if b, err := govalidator.ValidateStruct(c); !b {
		return fmt.Errorf("sso 调用配置有误:%v", err)
	}
	return nil
}

var httpConfigName string = "http"
var cacheExpireTime time.Duration = 5 * time.Minute
var cacheClearupTime time.Duration = 10 * time.Second
var authorityIgnoreURLs []string

//Option Option
type Option func()

//WithHTTPConfig WithHTTPConfig
func WithHTTPConfig(name string) Option {
	return func() {
		httpConfigName = name
	}
}

//WithCahce WithCahce
func WithCahce(expireTime, clearupTime time.Duration) Option {
	return func() {
		cacheExpireTime = expireTime
		cacheClearupTime = clearupTime
	}
}

//WithAuthorityIgnore 忽略授权检查地址(可通配/*, /** 等)
func WithAuthorityIgnore(urls ...string) Option {
	return func() {
		authorityIgnoreURLs = append(urls, "/sso/**")
	}
}

func getHTTPConfig() []string {
	if httpConfigName == "" {
		return []string{}
	}
	return []string{httpConfigName}

}
