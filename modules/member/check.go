
package member

import "github.com/micro-plat/hydra/component"

type ICheck interface {
}

type Check struct {
	c component.IContainer
}


func NewCheck(c component.IContainer) *Check {
	return &Check{
		c: c,
	}
}



