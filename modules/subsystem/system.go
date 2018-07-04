
package subsystem

import "github.com/micro-plat/hydra/component"

type ISystem interface {
}

type System struct {
	c component.IContainer
}


func NewSystem(c component.IContainer) *System {
	return &System{
		c: c,
	}
}



