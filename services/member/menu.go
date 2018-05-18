package member

import (
	"github.com/micro-plat/hydra/component"
	"github.com/micro-plat/hydra/context"
)

type MenuHandler struct {
	container component.IContainer
}

func NewMenuHandler(container component.IContainer) (u *MenuHandler) {
	return &MenuHandler{container: container}
}

//GetHandle 查询菜单列表
func (u *MenuHandler) GetHandle(ctx *context.Context) (r interface{}) {
	return "success"
}
