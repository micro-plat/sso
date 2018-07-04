
package role

import "github.com/micro-plat/hydra/component"

type IRole interface {
}

type Role struct {
	c component.IContainer
}


func NewRole(c component.IContainer) *Role {
	return &Role{
		c: c,
	}
}



