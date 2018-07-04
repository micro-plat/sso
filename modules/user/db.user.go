
package user

import "github.com/micro-plat/hydra/component"

type IDbUser interface {
}

type DbUser struct {
	c component.IContainer
}


func NewDbUser(c component.IContainer) *DbUser {
	return &DbUser{
		c: c,
	}
}



