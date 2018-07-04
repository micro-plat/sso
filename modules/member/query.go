
package member

import "github.com/micro-plat/hydra/component"

type IQuery interface {
}

type Query struct {
	c component.IContainer
}


func NewQuery(c component.IContainer) *Query {
	return &Query{
		c: c,
	}
}



