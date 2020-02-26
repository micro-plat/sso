package model

//RoleEditInput 编辑角色参数
type RoleEditInput struct {
	RoleName   string `form:"role_name" json:"role_name" valid:"required"`
	RoleID     int64  `form:"role_id" json:"role_id"`
	Status     int    `form:"status" json:"status"`
	IsAdd      int    `form:"is_add" json:"is_add" valid:"required"`
	BelongID   int    `form:"belong_id" json:"belong_id"`     //所属id(如加油站id,或公司id)
	BelongType int    `form:"belong_type" json:"belong_type"` //系统标识
}

//RoleAuthInput 角色授权输入参数
type RoleAuthInput struct {
	RoleID     string `form:"role_id" json:"role_id" valid:"required"`
	SysID      string `form:"sys_id" json:"sys_id" valid:"required"`
	SelectAuth string `form:"selectauth" json:"selectauth" valid:"ascii, required"`
}

//QueryRoleInput 查询角色信息所需参数
type QueryRoleInput struct {
	PageIndex int    `form:"pi" json:"pi" valid:"required"`
	PageSize  int    `form:"ps" json:"ps" valid:"required"`
	RoleName  string `form:"role_name" json:"role_name"`
	Status    int    `form:"status" json:"status"`
	BelongID  int    `form:"belong_id" json:"belong_id"` //所属id(如加油站id,或公司id)
}
