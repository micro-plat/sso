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

//从本地获取用户的数据权限信息
func getUserDataPermissionFromLocal(userID int64, bussinessType string) (string, error) {
	userDataPermissionKey := fmt.Sprintf("data_permission_cache_%s_%s", types.GetString(userID), bussinessType)
	value, found := localCache.Get(userDataPermissionKey)
	if !found {
		data, err := getUserDataPermissionFromAPI(userID, bussinessType)
		if err != nil {
			return "", err
		}
		localCache.Set(userDataPermissionKey, data, cache.DefaultExpiration)
		return data, nil
	}
	return types.GetString(value), nil
}

//GetUserDataPermissionFromAPI 获取【数据权限】数据
func getUserDataPermissionFromAPI(userID int64, businessType string) (r string, err error) {
	cfg := ssoClient.cfg
	values := net.NewValues()
	values.Set("type", businessType)
	values.Set("user_id", types.GetString(userID))
	values.Set("ident", cfg.ident)
	values.Set("timestamp", types.GetString(time.Now().Unix()))

	values = values.Sort()
	raw := values.Join("", "") + cfg.secret
	values.Set("sign", md5.Encrypt(raw))

	result := make(map[string]string)
	a, err := remoteRequest(cfg.host, getDataPermission, values.Join("=", "&"), &result)
	if err != nil {
		return "", err
	}
	fmt.Println(a)
	return result["data"], nil
}
