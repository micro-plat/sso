package conf

import (
	"fmt"

	"github.com/asaskevich/govalidator"
)

var ImgCodeSetting *ImgCodeConf

type ImgCodeConf struct {
	ImgCodeCacheTimeout      int `json:"imgcode_cache_timeout" valid:"required"`
	ImgCodeErrorLimit        int `json:"imgcode_err_limit" valid:"required"`
	ImgCodeErrorLimitTimeout int `json:"imgcode_err_limit_timeout" valid:"required"`
}

func (c *ImgCodeConf) Valid() error {
	if b, err := govalidator.ValidateStruct(c); !b {
		return fmt.Errorf("ImgCodeConf 配置有误:%v", err)
	}
	return nil
}

func NewImgConf() *ImgCodeConf {
	return &ImgCodeConf{
		ImgCodeCacheTimeout:      1200,
		ImgCodeErrorLimit:        3,
		ImgCodeErrorLimitTimeout: 1800,
	}
}

func init() {
	ImgCodeSetting = NewImgConf()
}
