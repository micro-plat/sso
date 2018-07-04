
package subsystem

import "github.com/micro-plat/hydra/component"

type IDbSystem interface {
}

type DbSystem struct {
	c component.IContainer
}


func NewDbSystem(c component.IContainer) *DbSystem {
	return &DbSystem{
		c: c,
	}
}



