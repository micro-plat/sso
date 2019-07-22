package model

// 需要编辑/添加的用户数据
type UserInputNew struct {
	UserName  string `form:"user_name" json:"user_name" valid:"ascii,required"`
	UserID    int64  `form:"user_id" json:"user_id"`
	RoleID    int64  `form:"role_id" json:"role_id" `
	Mobile    int64  `form:"mobile" json:"mobile" valid:"length(11|11),required"`
	Status    int    `form:"status" json:"status"`
	ExtParams string `form:"ext_params" json:"ext_params"`
	Auth      string `form:"auth" json:"auth" valid:"required"`
	Email     string `form:"email" json:"email" valid:"email,required"`
}

//QueryUserInput 查询用户列表输入参数
type QueryUserInput struct {
	PageIndex int    `form:"pi" json:"pi" valid:"required"`
	PageSize  int    `form:"ps" json:"ps" valid:"required"`
	UserName  string `form:"username" json:"username"`
	RoleID    string `form:"role_id" json:"role_id"`
}
