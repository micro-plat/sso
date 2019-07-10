package system

import (
	"github.com/micro-plat/hydra/component"
	"github.com/micro-plat/lib4go/db"
	"github.com/micro-plat/sso/apiserver/modules/const/sqls"
)

type IDbSystem interface {
	Get(ident string) (s db.QueryRow, err error)
}

// DbSystem  db 系统信息
type DbSystem struct {
	c component.IContainer
}

func NewDbSystem(c component.IContainer) *DbSystem {
	return &DbSystem{
		c: c,
	}
}

//Get 从数据库中获取系统信息
func (l *DbSystem) Get(ident string) (s db.QueryRow, err error) {
	db := l.c.GetRegularDB()
	data, _, _, err := db.Query(sqls.QuerySystemInfo, map[string]interface{}{
		"ident": ident,
	})
	return data.Get(0), err
}
