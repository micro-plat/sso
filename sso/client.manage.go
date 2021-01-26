package sso

import (
	cache "github.com/patrickmn/go-cache"
)

//存储sso client对象
var ssoClient *Client

//本地Tag缓存对象
var localCache *cache.Cache

//SaveSSOClient  保存sso client
func saveSSOClient(ssoAPIHost, ident, secret string) error {
	client, err := New(ssoAPIHost, ident, secret)
	if err != nil {
		return err
	}
	ssoClient = client
	localCache = cache.New(cacheExpireTime, cacheClearupTime)
	return nil
}

//GetSSOClient  获取sso client
func GetSSOClient() *Client {
	return ssoClient
}
