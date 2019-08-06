package service

import (
	"strings"
	"time"

	"github.com/micro-plat/lib4go/net"
	"github.com/micro-plat/lib4go/security/md5"
	"github.com/micro-plat/lib4go/types"
	"github.com/micro-plat/sso/sso/model"
)

//GetSystemInfo 获取系统信息
func GetSystemInfo(conf *model.Config) (data *System, err error) {
	values := net.NewValues()
	values.Set("ident", conf.Ident)
	values.Set("timestamp", types.GetString(time.Now().Unix()))

	values = values.Sort()
	raw := values.Join("=", "&") + "&key=" + conf.Secret
	values.Set("sign", strings.ToUpper(md5.Encrypt(raw)))

	sys := &System{}
	result, err := remoteRequest(conf.Host, model.SystemInfoUrl, values.Join("=", "&"), sys)
	if err != nil {
		return nil, err
	}
	return result.(*System), nil
}
