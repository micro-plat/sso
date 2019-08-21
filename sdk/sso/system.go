package sso

import (
	"time"

	"github.com/micro-plat/lib4go/net"
	"github.com/micro-plat/lib4go/security/md5"
	"github.com/micro-plat/lib4go/types"
)

type SystemLogic struct {
	conf *Config
}

// NewSystem xx
func NewSystem(conf *Config) *SystemLogic {
	return &SystemLogic{
		conf: conf,
	}
}

//GetSystemInfo 获取系统信息
func (s *SystemLogic) getSystemInfo() (data *System, err error) {
	values := net.NewValues()
	values.Set("ident", s.conf.ident)
	values.Set("timestamp", types.GetString(time.Now().Unix()))

	values = values.Sort()
	raw := values.Join("", "") + s.conf.secret
	values.Set("sign", md5.Encrypt(raw))

	sys := &System{}
	result, err := remoteRequest(s.conf.host, systemInfoUrl, values.Join("=", "&"), sys)
	if err != nil {
		return nil, err
	}
	return result.(*System), nil
}
