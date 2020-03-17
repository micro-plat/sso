package model

//UserInputNew 增加用户实体
type UserInputNew struct {
	UserName    string `form:"user_name" json:"user_name" valid:"required"`
	Mobile      string `form:"mobile" json:"mobile" valid:"required"`
	Email       string `form:"email" json:"email"`
	Ident       string `form:"ident" json:"ident"`
	FullName    string `form:"full_name" json:"full_name" valid:"required"`
	TargetIdent string `form:"target_ident" json:"target_ident" valid:"required"` //目标ident(系统)
	Source      string `form:"source" json:"source"`                              //来源
	SourceID    int    `form:"source_id" json:"source_id"`                        //来源id
}
