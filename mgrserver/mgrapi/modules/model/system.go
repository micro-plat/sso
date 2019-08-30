package model

type SystemEditInput struct {
	Enable        string `form:"enable" json:"enable" valid:"required"`
	Id            string `form:"id" json:"id" valid:"required"`
	Logo          string `form:"logo" json:"logo" valid:"required"`
	Name          string `form:"name" json:"name" valid:"required"`
	Theme         string `form:"theme" json:"theme"`
	Layout        string `form:"layout" json:"layout"`
	Ident         string `form:"ident" json:"ident"`
	Wechat_status string `form:"wechat_status" json:"wechat_status"`
	Secret        string `form:"secret" json:"secret"`
	CallBackUrl   string `form:"callbackurl" json:"callbackurl"`
}

type AddSystemInput struct {
	Name          string `form:"name" json:"name" valid:"required"`
	Logo          string `form:"logo" json:"logo" valid:"required"`
	Style         string `form:"style" json:"style" valid:"required"`
	Theme         string `form:"theme" json:"theme"`
	Ident         string `form:"ident" json:"ident" vaild:"required"`
	Wechat_status string `form:"wechat_status" json:"wechat_status"`
	Secret        string `form:"secret" json:"secret" valid:"required"`
	CallBackUrl   string `form:"callbackurl" json:"callbackurl"`
}
