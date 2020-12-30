package conf

import (
	"fmt"

	"github.com/asaskevich/govalidator"
)

var SmsCodeSetting *SmsCodeConf

type SmsCodeConf struct {
	SmsCodeSendRequestURL    string `json:"smscode_request_url" valid:"required"`
	SmsCodeCacheTimeout      int    `json:"smscode_cache_timeout" valid:"required"`
	SmsCodeErrorLimit        int    `json:"smscode_err_limit" valid:"required"`
	SmsCodeErrorLimitTimeout int    `json:"smscode_err_limit_timeout" valid:"required"`
}

func (c *SmsCodeConf) Valid() error {
	if b, err := govalidator.ValidateStruct(c); !b {
		return fmt.Errorf("SmsCodeConf 配置有误:%v", err)
	}
	return nil
}

func NewSmsCodeConf() *SmsCodeConf {
	return &SmsCodeConf{
		SmsCodeCacheTimeout:      300,
		SmsCodeErrorLimit:        5,
		SmsCodeErrorLimitTimeout: 300,
	}
}

func init() {
	SmsCodeSetting = NewSmsCodeConf()
}
