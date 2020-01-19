package model

type ImportReq struct {
	Id    string     `form:"id" json:"id" valid:"required"`
	Menus []MenuInfo `form:"menus" json:"menus" valid:"required"`
}

type MenuInfo struct {
	ID       int    `form:"id" json:"id" valid:"required"`
	Name     string `form:"name" json:"name" valid:"required"`
	Parent   int    `form:"parent" json:"parent" valid:"required"`
	LevelID  int    `form:"level_id" json:"level_id" valid:"required"`
	Icon     string `form:"icon" json:"icon"`
	Path     string `form:"path" json:"path"`
	Enable   int    `form:"enable" json:"enable"`
	Sortrank int    `form:"sortrank" json:"sortrank"`
	IsOpen   int    `form:"is_open" json:"is_open"`
}
