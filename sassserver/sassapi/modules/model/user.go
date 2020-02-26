package model

//UserInputNew 需要编辑/添加的用户数据
type UserInputNew struct {
	FullName   string `form:"full_name" json:"full_name" valid:"required"`
	UserName   string `form:"user_name" json:"user_name"`
	UserID     int64  `form:"user_id" json:"user_id"`
	RoleID     int64  `form:"role_id" json:"role_id" `
	Mobile     string `form:"mobile" json:"mobile" valid:"length(11|11),required"`
	Status     int    `form:"status" json:"status"`
	ExtParams  string `form:"ext_params" json:"ext_params"`
	Email      string `form:"email" json:"email"`
	SystemID   int    `form:"sys_id" json:"sys_id"`
	BelongID   int    `form:"belong_id" json:"belong_id"`     //这个不用前端传
	BelongType int    `form:"belong_type" json:"belong_type"` //这个不用前端传
}

//QueryUserInput 查询用户列表输入参数
type QueryUserInput struct {
	PageIndex  int    `form:"pi" json:"pi" valid:"required"`
	PageSize   int    `form:"ps" json:"ps" valid:"required"`
	Mobile     string `form:"mobile" json:"mobile"`
	RoleID     string `form:"role_id" json:"role_id"`
	Status     string `form:"status" json:"status"`
	BelongID   int    `form:"belong_id" json:"belong_id"`     //这个不用前端传
	BelongType int    `form:"belong_type" json:"belong_type"` //这个不用前端传
}
