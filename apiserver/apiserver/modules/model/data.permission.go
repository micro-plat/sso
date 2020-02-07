package model

//DataPermissionSyncReq 同步子系统的【数据权限】数据
type DataPermissionSyncReq struct {
	Ident  string `form:"ident" json:"ident" valid:"required"`
	Name   string `form:"name" json:"name" valid:"required"`
	Type   string `form:"type" json:"type" valid:"required"`
	Value  string `form:"value" json:"value" valid:"required"`
	Remark string `form:"remark" json:"remark" valid:"required"`
}

//DataPermissionGetReq 获取用户有权限的　[数据权限]　数据
type DataPermissionGetReq struct {
	UserID int    `form:"user_id" json:"user_id" valid:"required"`
	Ident  string `form:"ident" json:"ident" valid:"required"`
	Type   string `form:"type" json:"type" valid:"required"`
}
