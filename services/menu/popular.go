package menu

import (
	"github.com/micro-plat/hydra/component"
	"github.com/micro-plat/hydra/context"
	"github.com/micro-plat/sso/modules/member"
	"github.com/micro-plat/sso/modules/menu"
)

type PopularHandler struct {
	container component.IContainer
	m         menu.IGet
}

func NewPopularHandler(container component.IContainer) (u *PopularHandler) {
	return &PopularHandler{
		container: container,
		m:         menu.NewGet(container),
	}
}

func (u *PopularHandler) Handle(ctx *context.Context) (r interface{}) {
	uid := member.Get(ctx).UserID
	sysid := member.Get(ctx).SystemID
	data, err := u.m.QueryPopular(uid, sysid)
	if err != nil {
		return err
	}
	result := make([]map[string]interface{}, 0, 4)
	for _, row1 := range data {
		if row1.GetInt("parent") == 0 && row1.GetInt("level_id") == 2 {
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
	return result
}
