package menu

import (
	"fmt"

	"github.com/micro-plat/hydra/component"
	"github.com/micro-plat/hydra/context"
	"github.com/micro-plat/lib4go/db"
	"github.com/micro-plat/sso/modules/sql"
)

type IGet interface {
	Query(uid int64, sysid int) (db.QueryRows, error)
	Verify(uid int64, sysid int, menuURL string) error
}

type Get struct {
	c component.IContainer
}

func NewGet(c component.IContainer) *Get {
	return &Get{
		c: c,
	}
}

//Query 获取用户指定系统的菜单信息
func (l *Get) Query(uid int64, sysid int) (db.QueryRows, error) {
	db := l.c.GetRegularDB()
	//根据用户名密码，查询用户信息
	data, _, _, err := db.Query(sql.QueryUserMenus, map[string]interface{}{
		"user_id": uid,
		"sys_id":  sysid,
	})
	if err != nil {
		return nil, err
	}
	return data, nil
}

//Verify 获取用户指定系统的菜单信息
func (l *Get) Verify(uid int64, sysid int, menuURL string) error {
	db := l.c.GetRegularDB()
	//根据用户名密码，查询用户信息
	data, _, _, err := db.Scalar(sql.QueryUserMenu, map[string]interface{}{
		"user_id": uid,
		"sys_id":  sysid,
		"path":    menuURL,
	})
	if err != nil {
		return err
	}
	if fmt.Sprint(data) == "1" {
		return nil
	}
	return context.NewError(context.ERR_FORBIDDEN, fmt.Errorf("未查找到菜单"))
}
