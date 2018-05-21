
package menu

import "github.com/micro-plat/hydra/component"

type IGet interface {
}

type Get struct {
	c component.IContainer
}


func NewGet(c component.IContainer) *Get {
	return &Get{
		c: c,
	}
}



