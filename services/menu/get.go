package menu

import (
	"github.com/micro-plat/sso/modules/member"

	"github.com/micro-plat/hydra/component"
	"github.com/micro-plat/hydra/context"
	"github.com/micro-plat/sso/modules/menu"
)

//GetHandler 菜单查询对象
type GetHandler struct {
	c component.IContainer
	m menu.IGet
}

//NewGetHandler 创建菜单查询对象
func NewGetHandler(container component.IContainer) (u *GetHandler) {
	return &GetHandler{
		c: container,
		m: menu.NewGet(container),
	}
}

//Handle 查询指定用户在指定系统的菜单列表
func (u *GetHandler) Handle(ctx *context.Context) (r interface{}) {
	uid := member.Get(ctx).UserID
	sysid := member.Get(ctx).SystemID
	data, err := u.m.Query(uid, sysid)
	if err != nil {
		return err
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

	return result
}
