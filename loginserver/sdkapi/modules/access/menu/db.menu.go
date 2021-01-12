package menu

import (
	"fmt"

	"github.com/micro-plat/hydra/components"
	"github.com/micro-plat/lib4go/types"
	"github.com/micro-plat/sso/loginserver/sdkapi/modules/const/sqls"
)

type IMenu interface {
	Query(uid int64, ident string) ([]map[string]interface{}, error)
	QueryUserMenuTags(uid int64, ident string) (types.XMaps, error)
}

type Menu struct {
}

func NewMenu() *Menu {
	return &Menu{}
}

//Query 获取用户指定系统的菜单信息
func (l *Menu) Query(uid int64, ident string) ([]map[string]interface{}, error) {
	db := components.Def.DB().GetRegularDB()
	data, err := db.Query(sqls.QueryUserMenus, map[string]interface{}{
		"user_id": uid,
		"ident":   ident,
	})
	if err != nil {
		return nil, err
	}
	result := make([]map[string]interface{}, 0, 4)
	for _, row1 := range data {
		if row1.GetInt("parent") == 0 && row1.GetInt("level_id") == 1 {
			children1 := make([]map[string]interface{}, 0, 4)
			for _, row2 := range data {
				if row2.GetInt("parent") == row1.GetInt("id") && row2.GetInt("level_id") == 2 {
					children2 := make([]map[string]interface{}, 0, 8)
					for _, row3 := range data {
						if row3.GetInt("parent") == row2.GetInt("id") && row3.GetInt("level_id") == 3 {
							children2 = append(children2, row3)
						}
					}
					children1 = append(children1, row2)
					row2["children"] = children2
				}
			}
			row1["children"] = children1
			result = append(result, row1)
		}
	}
	return result, nil
}

//QueryUserMenuTags 获取用户有权限的tags
func (l *Menu) QueryUserMenuTags(uid int64, ident string) (types.XMaps, error) {
	db := components.Def.DB().GetRegularDB()
	data, err := db.Query(sqls.QueryUserMenuTags, map[string]interface{}{
		"user_id": uid,
		"ident":   ident,
	})
	if err != nil {
		return nil, fmt.Errorf("QueryUserMenuTags出错: err:%+v", err)
	}
	return data, nil
}
