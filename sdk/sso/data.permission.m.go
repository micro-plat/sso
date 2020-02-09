package sso

//SyncReq 同步子系统的【数据权限】数据
type SyncReq struct {
	Name     string `form:"name" json:"name" valid:"required"`
	Type     string `form:"type" json:"type" valid:"required"`
	TypeName string `form:"type_name" json:"type_name" valid:"required"` //类型名称
	Value    string `form:"value" json:"value" valid:"required"`
	Remark   string `form:"remark" json:"remark"`
}
