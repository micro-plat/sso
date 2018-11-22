package menu

import (
	"fmt"

	"github.com/micro-plat/hydra/component"
	"github.com/micro-plat/sso/flowserver/modules/const/sql"
)

type IPopular interface {
	Query(uid int64, sysid int) ([]map[string]interface{}, error)
	Save(uid int64, sysid int, pid []string, mid []string) error
}

type Popular struct {
	c component.IContainer
}

func NewPopular(c component.IContainer) *Popular {
	return &Popular{
		c: c,
	}
}

//Query 获取用户指定系统的菜单信息
func (l *Popular) Query(uid int64, sysid int) ([]map[string]interface{}, error) {
	db := l.c.GetRegularDB()
	data, _, _, err := db.Query(sql.QueryUserPopularMenus, map[string]interface{}{
		"user_id": uid,
		"sys_id":  sysid,
	})
	if err != nil {
		return nil, err
	}
	level2 := make(map[string]int)
	result := make([]map[string]interface{}, 0, 4)
	for _, row1 := range data {
		if _, ok := level2[row1.GetString("id")]; !ok && row1.GetInt("level_id") == 2 {
			level2[row1.GetString("id")] = 0
			children1 := make([]map[string]interface{}, 0, 4)
			for _, row2 := range data {
				if row2.GetInt("parent") == row1.GetInt("id") && row2.GetInt("level_id") == 3 {
					children1 = append(children1, row2)
				}
			}
			row1["children"] = children1
			result = append(result, row1)
		}
	}
	return result, nil
}

//Save 保存用户常用菜单列表
func (l *Popular) Save(uid int64, sysid int, pid []string, mid []string) error {
	db := l.c.GetRegularDB()
	trans, err := db.Begin()
	if err != nil {
		return err
	}
	for i, p := range pid {
		d, _, _, err := trans.Scalar(sql.CheckUserPopularMenu, map[string]interface{}{
			"user_id": uid,
			"sys_id":  sysid,
			"menu_id": mid[i],
		})
		if err != nil {
			trans.Rollback()
			return err
		}
		if fmt.Sprint(d) == "1" {
			_, _, _, err := trans.Execute(sql.UpdateUserPopularMenu, map[string]interface{}{
				"user_id": uid,
				"sys_id":  sysid,
				"menu_id": mid[i],
			})
			if err != nil {
				trans.Rollback()
				return err
			}
			continue
		}
		_, _, _, err = trans.Execute(sql.SaveUserPopularMenu, map[string]interface{}{
			"user_id":   uid,
			"sys_id":    sysid,
			"parent_id": p,
			"menu_id":   mid[i],
			"used_cnt":  1,
		})
		if err != nil {
			trans.Rollback()
			return err
		}
	}
	trans.Commit()
	return nil
}
