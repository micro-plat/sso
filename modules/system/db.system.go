package system

import (
	"github.com/micro-plat/hydra/component"
	"github.com/micro-plat/lib4go/db"
	"github.com/micro-plat/sso/modules/const/sql"
)

type IDbSystem interface {
	Query(sysid int) (s db.QueryRow, err error)
}

type DbSystem struct {
	c component.IContainer
}

func NewDbSystem(c component.IContainer) *DbSystem {
	return &DbSystem{
		c: c,
	}
}

//Query 从数据库中获取系统信息
func (l *DbSystem) Query(sysid int) (s db.QueryRow, err error) {
	db := l.c.GetRegularDB()
	data, _, _, err := db.Query(sql.QuerySystemInfo, map[string]interface{}{
		"sys_id": sysid,
	})
	return data.Get(0), err
}
