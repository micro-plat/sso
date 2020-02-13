package model

//DataPermissionReq 数据权限
type DataPermissionReq struct {
	ID            string `form:"id" json:"id" `
	Ident         string `form:"ident" json:"ident" `
	SysID         string `form:"sys_id" json:"sys_id" valid:"required"`
	Name          string `form:"name" json:"name" valid:"required"`             //名称
	TableName     string `form:"table_name" json:"table_name" valid:"required"` //表
	OperateAction string `form:"operate_action" json:"operate_action"`          //操作动作
	Rules         string `form:"rules" json:"rules" valid:"required"`           //值
	Remark        string `form:"remark" json:"remark"`
}
