package system

import (
	"fmt"

	"github.com/micro-plat/hydra/components"
	"github.com/micro-plat/lib4go/types"
	commsqls "github.com/micro-plat/sso/loginserver/loginapi/modules/const/sqls"
)

type IDbSystem interface {
	QueryUserSystem(userId int64) (s types.XMaps, err error)
	QuerySysInfoByIdent(ident string) (types.IXMap, error)
}

//DbSystem 系统信息
type DbSystem struct {
}

//NewDbSystem 系统信息
func NewDbSystem() *DbSystem {
	return &DbSystem{}
}

//QueryUserSystem 还回用户可用的子系统
func (l *DbSystem) QueryUserSystem(userId int64) (s types.XMaps, err error) {
	db := components.Def.DB().GetRegularDB()
	data, err := db.Query(
		commsqls.SearchUserSystemInfo, map[string]interface{}{
			"user_id": userId,
		})
	return data, err
}

//QuerySysInfoByIdent
func (l *DbSystem) QuerySysInfoByIdent(ident string) (types.IXMap, error) {
	db := components.Def.DB().GetRegularDB()
	data, err := db.Query(
		commsqls.QuerySysInfoByIdent, map[string]interface{}{
			"ident": ident,
		})
	if err != nil {
		return nil, err
	}
	if data.IsEmpty() {
		return nil, fmt.Errorf("系统不存在:%s", ident)
	}
	return data.Get(0), nil
}
