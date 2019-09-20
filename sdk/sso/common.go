package sso

import (
	"fmt"
	"strconv"
	"time"

	"github.com/micro-plat/lib4go/net"
	"github.com/micro-plat/lib4go/security/md5"
	"github.com/micro-plat/lib4go/types"
	"github.com/patrickmn/go-cache"
)

//存储sso client对象
var ssoClient *Client

//本地菜单缓存对象
var localCache *cache.Cache

//从本地获取用户的菜单信息
func getUserMenuFromLocal(userID int) ([]Menu, error) {
	userMenuKey := fmt.Sprintf("%s_cache", strconv.Itoa(userID))
	value, found := localCache.Get(userMenuKey)
	if !found {
		menu, err := getUserMenuFromAPIServer(userID)
		if err != nil {
			return nil, err
		}
		localCache.Set(userMenuKey, menu, cache.DefaultExpiration)
		return menu, nil
	}
	return value.([]Menu), nil
}

//getUserMenuFromAPIServer 从apiserver中获取菜单信息,并缓存在系统中
func getUserMenuFromAPIServer(userID int) ([]Menu, error) {
	cfg := ssoClient.cfg
	values := net.NewValues()
	values.Set("user_id", types.GetString(userID))
	values.Set("ident", cfg.ident)
	values.Set("timestamp", types.GetString(time.Now().Unix()))

	values = values.Sort()
	raw := values.Join("", "") + cfg.secret
	values.Set("sign", md5.Encrypt(raw))

	var other []Menu
	_, err := remoteRequest(cfg.host, userMenuUrl, values.Join("=", "&"), &other)
	if err != nil {
		return nil, err
	}
	return other, nil
}
