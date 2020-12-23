package model

type ImportReq struct {
	Id    string     `form:"id" json:"id" valid:"required"`
	Menus []MenuInfo `form:"menus" json:"menus" valid:"required"`
}

type MenuInfo struct {
	ID       string `form:"id" json:"id" valid:"required"`
	Name     string `form:"name" json:"name" valid:"required"`
	Parent   string `form:"parent" json:"parent" valid:"required"`
	LevelID  string `form:"level_id" json:"level_id" valid:"required"`
	Icon     string `form:"icon" json:"icon"`
	Path     string `form:"path" json:"path"`
	Enable   string `form:"enable" json:"enable"`
	Sortrank string `form:"sortrank" json:"sortrank"`
	IsOpen   string `form:"is_open" json:"is_open"`
}
