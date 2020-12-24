package system

import (
	"errors"

	"github.com/micro-plat/hydra/components"
	"github.com/micro-plat/lib4go/db"
	"github.com/micro-plat/sso/loginserver/webserver/modules/const/sqls"
)

type IDbSystem interface {
	QueryUserSystem(userId int64) (s db.QueryRows, err error)
	QuerySysInfoByIdent(ident string) (db.QueryRow, error)
}

//DbSystem 系统信息
type DbSystem struct {
}

//NewDbSystem 系统信息
func NewDbSystem() *DbSystem {
	return &DbSystem{}
}

//QueryUserSystem 还回用户可用的子系统
func (l *DbSystem) QueryUserSystem(userId int64) (s db.QueryRows, err error) {
	db := components.Def.DB().GetRegularDB()
	data, _, _, err := db.Query(
		sqls.SearchUserSystemInfo, map[string]interface{}{
			"user_id": userId,
		})
	return data, err
}

//QuerySysInfoByIdent
func (l *DbSystem) QuerySysInfoByIdent(ident string) (db.QueryRow, error) {
	db := components.Def.DB().GetRegularDB()
	data, _, _, err := db.Query(
		sqls.QuerySysInfoByIdent, map[string]interface{}{
			"ident": ident,
		})
	if err != nil {
		return nil, err
	}
	if data.IsEmpty() {
		return nil, errors.New("系统不存在")
	}
	return data.Get(0), nil
}
