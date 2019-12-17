package sso

import (
	"fmt"
	"strconv"
	"time"

	"github.com/micro-plat/lib4go/net"
	"github.com/micro-plat/lib4go/security/md5"
	"github.com/micro-plat/lib4go/types"
	cache "github.com/patrickmn/go-cache"
)

//存储sso client对象
var ssoClient *Client

//本地Tag缓存对象
var localCache *cache.Cache

//从本地获取用户的Tag信息
func getUserTagFromLocal(userID int) ([]Menu, error) {
	userTagKey := fmt.Sprintf("tag_cache_%s", strconv.Itoa(userID))
	value, found := localCache.Get(userTagKey)
	if !found {
		tags, err := getUserTagFromApiserver(userID)
		if err != nil {
			return nil, err
		}
		localCache.Set(userTagKey, tags, cache.DefaultExpiration)
		return tags, nil
	}
	return value.([]Menu), nil
}

//getUserTagFromApiserver 从apiserver获取用户有权限的Tag
func getUserTagFromApiserver(userID int) ([]Menu, error) {
	cfg := ssoClient.cfg
	values := net.NewValues()
	values.Set("user_id", types.GetString(userID))
	values.Set("ident", cfg.ident)
	values.Set("timestamp", types.GetString(time.Now().Unix()))

	values = values.Sort()
	raw := values.Join("", "") + cfg.secret
	values.Set("sign", md5.Encrypt(raw))

	var other []Menu
	_, err := remoteRequest(cfg.host, userAllTag, values.Join("=", "&"), &other)
	if err != nil {
		return nil, err
	}
	return other, nil
}
