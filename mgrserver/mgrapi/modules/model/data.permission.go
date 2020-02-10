package model

//DataPermissionReq 数据权限
type DataPermissionReq struct {
	ID       string `form:"id" json:"id" `
	Ident    string `form:"ident" json:"ident" `
	SysID    string `form:"sys_id" json:"sys_id" valid:"required"`
	Name     string `form:"name" json:"name" valid:"required"`   //名称
	Type     string `form:"type" json:"type"`                    //类型(good_category)
	TypeName string `form:"type_name" json:"type_name"`          //类型名称
	Value    string `form:"value" json:"value" valid:"required"` //值
	Remark   string `form:"remark" json:"remark"`
}
