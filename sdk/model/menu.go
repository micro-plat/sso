package model

type Menu struct {
	ID       string `json:"id"`
	Name     string `json:"name"`
	Level    string `json:"level_id"`
	IsOpen   string `json:"is_open"`
	Icon     string `json:"icon"`
	SystemID string `json:"sys_id"`
	Parent   string `json:"parent"`
	Path     string `json:"path"`
	Sortrank string `json:"sortrank"`
	Children []Menu `json:"children,omitempty"`
}
