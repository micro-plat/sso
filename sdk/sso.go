package sdk

import (
	"errors"

	"github.com/micro-plat/sso/sdk/model"
)

var errApiHost error = errors.New("api host 不能为空")
var errIdent error = errors.New("ident 不能为空")
var errSecret error = errors.New("secret 不能为空")

//SetConfig 设置配置信息
func SetConfig(cfg *model.Config) error {
	if cfg.ApiHost == "" {
		return errApiHost
	}

	if cfg.Ident == "" {
		return errIdent
	}

	if cfg.Secret == "" {
		return errSecret
	}

	model.SysInfoConfig = cfg
	return nil
}
