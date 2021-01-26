package sso

//Menu 菜单数据
type Menu struct {
	ID   string `json:"id"`
	Name string `json:"name"`
	//Level    string `json:"level_id"`
	IsOpen   string `json:"is_open"`
	Icon     string `json:"icon"`
	SystemID string `json:"sys_id"`
	//Parent string `json:"parent"`
	Path string `json:"path"`
	//Sortrank string `json:"sortrank"`
	Children []Menu `json:"children,omitempty"`
}

const (
	MenuLevelMenu = "3"
	MenuLevelFunc = "4"
)

//SystemRoleAuthority SystemRoleAuthority
type SystemRoleAuthority struct {
	MenuID   string `json:"id"`
	MenuName string `json:"name"`
	Parent   string `json:"parent"`
	SystemID string `json:"sys_id"`
	Level    string `json:"level_id"`
	Path     string `json:"path"`
	FuncTags map[string]bool
}
