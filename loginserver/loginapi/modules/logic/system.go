package logic

import (
	"strings"

	"github.com/micro-plat/lib4go/types"
	"github.com/micro-plat/sso/loginserver/loginapi/modules/access/system"
	commodel "github.com/micro-plat/sso/loginserver/loginapi/modules/model"
)

// ISystemLogic xx
type ISystemLogic interface {
	QueryUserSystem(userId int64) (types.XMaps, error)
	QuerySysInfoByIdent(ident string) (types.IXMap, error)
	GetSystemConfig(ident string) (map[string]interface{}, error)
}

// SystemLogic 操作日志
type SystemLogic struct {
	dbSys system.IDbSystem
}

// NewSystemLogic xx
func NewSystemLogic() *SystemLogic {
	return &SystemLogic{
		dbSys: system.NewDbSystem(),
	}
}

//GetSystemConfig 获取系统配置信息
func (s *SystemLogic) GetSystemConfig(ident string) (map[string]interface{}, error) {
	result := map[string]interface{}{"system_name": "用户登录", "require_valid_code": commodel.GetLoginConf().RequireValidCode}
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
func (s *SystemLogic) QueryUserSystem(userId int64) (types.XMaps, error) {
	return s.dbSys.QueryUserSystem(userId)
}

// QuerySysInfoByIdent 根据ident查询系统信息
func (s *SystemLogic) QuerySysInfoByIdent(ident string) (types.IXMap, error) {
	return s.dbSys.QuerySysInfoByIdent(ident)
}
