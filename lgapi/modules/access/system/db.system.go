package system

import (
	"github.com/micro-plat/hydra/component"
	"github.com/micro-plat/lib4go/db"
	"github.com/micro-plat/sso/lgapi/modules/const/sqls"
)

type IDbSystem interface {
	QueryUserSystem(userId int64) (s db.QueryRows, err error)
}

//DbSystem 系统信息
type DbSystem struct {
	c component.IContainer
}

//NewDbSystem 系统信息
func NewDbSystem(c component.IContainer) *DbSystem {
	return &DbSystem{
		c: c,
	}
}

//QueryUserSystem 还回用户可用的子系统
func (l *DbSystem) QueryUserSystem(userId int64) (s db.QueryRows, err error) {
	db := l.c.GetRegularDB()
	data, _, _, err := db.Query(
		sqls.SearchUserSystemInfo, map[string]interface{}{
			"user_id": userId,
		})
	return data, err

}
