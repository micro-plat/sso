package sso

//import cache "github.com/patrickmn/go-cache"

//SaveSSOClient  保存sso client
func saveSSOClient(ssoAPIHost, ident, secret string) error {
	client, err := New(ssoAPIHost, ident, secret)
	if err != nil {
		return err
	}
	ssoClient = client
	//localCache = cache.New(2*time.Minute, 10*time.Second)
	return nil
}

//GetSSOClient  获取sso client
func GetSSOClient() *Client {
	return ssoClient
}
