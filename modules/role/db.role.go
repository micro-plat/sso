
package role

import "github.com/micro-plat/hydra/component"

type IDbRole interface {
}

type DbRole struct {
	c component.IContainer
}


func NewDbRole(c component.IContainer) *DbRole {
	return &DbRole{
		c: c,
	}
}



