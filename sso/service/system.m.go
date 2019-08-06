package service

//System 系统信息 (由于接口返回全是string,因此...)
type System struct {
	Ident string `json:"ident"`
	//Login_url string
	Name  string `json:"name"`
	Theme string `json:"theme"`
	//Wechat_status int
	ID       string `json:"id"`
	IndexUrl string `json:"index_url"`
	Layout   string `json:"layout"`
	Logo     string `json:"logo"`
}
