package model

type SystemFuncAddInput struct {
	Parentid    int    `form:"parentid" json:"parentid"`
	ParentLevel int    `form:"parentlevel" json:"parentlevel"`
	Sysid       int    `form:"sysid" json:"sysid"`
	Name        string `form:"name" json:"name" valid:"required"`
	Icon        string `form:"icon" json:"icon" valid:"required"`
	Path        string `form:"path" json:"path" valid:"required"`
	IsOpen      int    `form:"is_open" json:"is_open"`
}

type SystemFuncEditInput struct {
	Id       string `form:"id" json:"id" valid:"required"`
	Name     string `form:"name" json:"name" valid:"required"`
	Icon     string `form:"icon" json:"icon" valid:"required"`
	Path     string `form:"path" json:"path" valid:"required"`
	IsOpen   int    `form:"is_open" json:"is_open"`
	Sortrank string `form:"sortrank" json:"sortrank" valid:"required"`
}
