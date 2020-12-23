package model

type TemplateMsg struct {
	Touser      string        `json:"touser"`      //接收者的OpenID
	TemplateID  string        `json:"template_id"` //模板消息ID
	URL         string        `json:"url"`         //点击后跳转链接
	Miniprogram Miniprogram   `json:"miniprogram"` //点击跳转小程序
	Data        *TemplateData `json:"data"`
}
type Miniprogram struct {
	AppID    string `json:"appid"`
	Pagepath string `json:"pagepath"`
}

type TemplateData struct {
	First    KeyWordData `json:"first,omitempty"`
	Keyword1 KeyWordData `json:"keyword1,omitempty"`
	Keyword2 KeyWordData `json:"keyword2,omitempty"`
	Keyword3 KeyWordData `json:"keyword3,omitempty"`
	Remark   KeyWordData `json:"remark,omitempty"`
}

type KeyWordData struct {
	Value string `json:"value"`
	Color string `json:"color,omitempty"`
}
