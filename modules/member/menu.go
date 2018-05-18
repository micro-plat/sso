
package member

import "github.com/micro-plat/hydra/component"

type IMenu interface {
}

type Menu struct {
	c component.IContainer
}


func NewMenu(c component.IContainer) *Menu {
	return &Menu{
		c: c,
	}
}



