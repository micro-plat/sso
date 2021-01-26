package sso

import (
	"fmt"
	"time"

	"github.com/micro-plat/lib4go/net"
	"github.com/micro-plat/lib4go/security/md5"
	"github.com/micro-plat/lib4go/types"
	cache "github.com/patrickmn/go-cache"
)

//从本地获取用户的Tag信息
func getRoleTagFromLocal(roleID int) (map[string]*SystemRoleAuthority, error) {
	cfg := ssoClient.cfg
	ident := cfg.ident

	userTagKey := fmt.Sprintf("_sso_menus_cache_%s_%d", ident, roleID)
	value, found := localCache.Get(userTagKey)
	if !found {
		tags, err := getRoleAuthorityFromReomteServer(roleID)
		if err != nil {
			return nil, err
		}
		localCache.Set(userTagKey, tags, cache.DefaultExpiration)
		return tags, nil
	}
	return value.(map[string]*SystemRoleAuthority), nil
}

//getRoleAuthorityFromReomteServer 从apiserver获取用户有权限的Tag
func getRoleAuthorityFromReomteServer(roleID int) (map[string]*SystemRoleAuthority, error) {
	cfg := ssoClient.cfg
	values := net.NewValues()
	values.Set("role_id", types.GetString(roleID))
	values.Set("ident", cfg.ident)
	values.Set("timestamp", types.GetString(time.Now().Unix()))

	values = values.Sort()
	raw := values.Join("", "") + cfg.secret
	values.Set("sign", md5.Encrypt(raw))

	var authorities types.XMaps
	_, err := remoteRequest(cfg.host, getRoleTags, values.Join("=", "&"), &authorities)
	if err != nil {
		return nil, err
	}
	result := buildAuthorityData(authorities)
	return result, nil
}

func buildAuthorityData(authorities types.XMaps) map[string]*SystemRoleAuthority {
	result := map[string]*SystemRoleAuthority{}
	for i := range authorities {
		cur := authorities[i]
		level := cur.GetString("level_id")
		path := cur.GetString("path")
		_, ok := result[path]
		if !ok {
			result[path] = &SystemRoleAuthority{
				MenuID:   cur.GetString("id"),
				MenuName: cur.GetString("name"),
				Parent:   cur.GetString("parent"),
				SystemID: cur.GetString("sys_id"),
				Level:    level,
				Path:     cur.GetString("path"),
				FuncTags: map[string]bool{},
			}
		}
		if cur.GetString("tag") != "" {
			result[path].FuncTags[cur.GetString("tag")] = true
		}
	}
	return result
}
