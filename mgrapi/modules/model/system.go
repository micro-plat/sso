package model

type SystemEditInput struct {
	Enable        string `form:"enable" json:"enable" valid:"required"`
	Id            string `form:"id" json:"id" valid:"required"`
	Index_url     string `form:"index_url" json:"index_url" valid:"required"`
	Login_timeout string `form:"login_timeout" json:"login_timeout" valid:"required"`
	//Logo          string `form:"logo" json:"logo" valid:"required"`
	Logo   string `form:"logo" json:"logo"`
	Name   string `form:"name" json:"name" valid:"required"`
	Theme  string `form:"theme" json:"theme"`
	Layout string `form:"layout" json:"layout"`
	Ident  string `form:"ident" json:"ident"`
	//Wechat_status string `form:"wechat_status" json:"wechat_status" valid:"required"`
	Wechat_status string `form:"wechat_status" json:"wechat_status"`
	//Secret        string `form:"secret" json:"secret" valid:"required"`
	Secret string `form:"secret" json:"secret"`
}

type AddSystemInput struct {
	Name     string `form:"name" json:"name" valid:"required"`
	Addr     string `form:"addr" json:"addr" valid:"required"`
	Time_out string `form:"time_out" json:"time_out" valid:"required"`
	//Logo     string `form:"logo" json:"logo" valid:"required"`
	Logo  string `form:"logo" json:"logo"`
	Style string `form:"style" json:"style" valid:"required"`
	Theme string `form:"theme" json:"theme"`
	Ident string `form:"ident" json:"ident" vaild:"required"`
	//Wechat_status string `form:"wechat_status" json:"wechat_status" valid:"required"`
	Wechat_status string `form:"wechat_status" json:"wechat_status"`
	Secret        string `form:"secret" json:"secret" valid:"required"`
}
