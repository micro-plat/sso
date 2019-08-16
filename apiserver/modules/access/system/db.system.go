package system

import (
	"errors"

	"github.com/micro-plat/hydra/component"
	"github.com/micro-plat/lib4go/db"
	"github.com/micro-plat/sso/apiserver/modules/const/sqls"
)

type IDbSystem interface {
	Get(ident string) (s db.QueryRow, err error)
	QuerySystemStatus(ident string) (s int, err error)
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

	if err != nil {
		return nil, err
	}
	if data.IsEmpty() {
		return nil, errors.New("ident 不存在")
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
