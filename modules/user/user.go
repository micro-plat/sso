
package user

import "github.com/micro-plat/hydra/component"

type IUser interface {
}

type User struct {
	c component.IContainer
}


func NewUser(c component.IContainer) *User {
	return &User{
		c: c,
	}
}



