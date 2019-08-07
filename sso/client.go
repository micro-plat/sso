package sso

//Client sso client
type Client struct {
	cfg *Config
}

//New SSOClient
func New(apiHost, ident, secret string) (*Client, error) {
	cfg := &Config{
		host:   apiHost,
		ident:  ident,
		secret: secret,
	}
	if err := cfg.Valid(); err != nil {
		return nil, err
	}
	return &Client{cfg: cfg}, nil
}

//CheckCodeLogin 验证回传code并获取登录用户信息
func (client *Client) CheckCodeLogin(code string) (res *LoginState, err error) {
	u := newUser(client.cfg)
	return u.checkCodeLogin(code)
}

//GetUserInfoByName 根据用户名获取用户信息
func (client *Client) GetUserInfoByName(userName string) (info *User, err error) {
	u := newUser(client.cfg)
	return u.getUserInfoByName(userName)
}

//GetUserMenu 获取用户菜单信息
func (client *Client) GetUserMenu(userID int) (*[]*Menu, error) {
	u := newUser(client.cfg)
	return u.getUserMenu(userID)
}

//GetSystemInfo 获取系统信息
func (client *Client) GetSystemInfo() (data *System, err error) {
	s := NewSystem(client.cfg)
	return s.getSystemInfo()
}
