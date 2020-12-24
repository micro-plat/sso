package system

import (
	"errors"

	"github.com/micro-plat/hydra/components"
	"github.com/micro-plat/lib4go/db"
	"github.com/micro-plat/sso/loginserver/apiserver/modules/const/sqls"
)

type IDbSystem interface {
	Get(ident string) (s db.QueryRow, err error)
	QuerySystemStatus(ident string) (s int, err error)
}

// DbSystem  db 系统信息
type DbSystem struct {
}

func NewDbSystem() *DbSystem {
	return &DbSystem{}
}

//Get 从数据库中获取系统信息
func (l *DbSystem) Get(ident string) (s db.QueryRow, err error) {
	db := components.Def.DB().GetRegularDB()
	data, _, _, err := db.Query(sqls.QuerySystemInfo, map[string]interface{}{
		"ident": ident,
	})

	if err != nil {
		return nil, err
	}
	if data.IsEmpty() {
		return nil, errors.New("系统不存在或则系统被禁用")
	}
	return data.Get(0), err
}

//QuerySystemStatus 查询某个系统的状态
func (l *DbSystem) QuerySystemStatus(ident string) (s int, err error) {
	data, err := l.Get(ident)
	if err != nil {
		return 0, err
	}
	return data.GetInt("enable"), nil
}
