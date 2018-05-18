
package member

import "github.com/micro-plat/hydra/component"

type ILogin interface {
}

type Login struct {
	c component.IContainer
}


func NewLogin(c component.IContainer) *Login {
	return &Login{
		c: c,
	}
}



