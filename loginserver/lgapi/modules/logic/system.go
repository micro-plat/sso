package logic

import (
	"strings"

	"github.com/micro-plat/hydra/component"
	"github.com/micro-plat/lib4go/db"
	commodel "github.com/micro-plat/sso/common/module/model"
	"github.com/micro-plat/sso/loginserver/lgapi/modules/access/system"
)

// ISystemLogic xx
type ISystemLogic interface {
	QueryUserSystem(userId int64) (db.QueryRows, error)
	QuerySysInfoByIdent(ident string) (db.QueryRow, error)
	GetSystemConfig(ident string) (map[string]interface{}, error)
}

// SystemLogic 操作日志
type SystemLogic struct {
	c     component.IContainer
	dbSys system.IDbSystem
}

// NewSystemLogic xx
func NewSystemLogic(c component.IContainer) *SystemLogic {
	return &SystemLogic{
		c:     c,
		dbSys: system.NewDbSystem(c),
	}
}

//GetSystemConfig 获取系统配置信息
func (s *SystemLogic) GetSystemConfig(ident string) (map[string]interface{}, error) {
	result := map[string]interface{}{"system_name": "用户登录", "require_wx_code": commodel.GetConf(s.c).RequireWxCode}
	if strings.EqualFold(ident, "") {
		return result, nil
	}
	data, err := s.dbSys.QuerySysInfoByIdent(ident)
	if err != nil {
		return nil, err
	}
	result["system_name"] = data.GetString("name")
	return result, nil
}

// QueryUserSystem 返回用户可用的子系统信息
func (s *SystemLogic) QueryUserSystem(userId int64) (db.QueryRows, error) {
	return s.dbSys.QueryUserSystem(userId)
}

// QuerySysInfoByIdent 根据ident查询系统信息
func (s *SystemLogic) QuerySysInfoByIdent(ident string) (db.QueryRow, error) {
	return s.dbSys.QuerySysInfoByIdent(ident)
}
