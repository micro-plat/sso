package notify

import (
	"encoding/json"
	"fmt"

	"github.com/micro-plat/lib4go/db"
	"github.com/micro-plat/hydra/component"
	"github.com/micro-plat/lib4go/transform"
	"github.com/micro-plat/hydra/context"
	"github.com/micro-plat/lib4go/types"
	
	
)

const (
	cacheFormatNotify    	= "{sso}:notify:info:{@title}-{@userID}-{@sysID}-{@pi}-{@ps}"
	cacheFormatNotifyDel 	= "{sso}:notify:info:*"
	cacheFormatNotifyCount  = "{sso}:notify:info:{@title}-{@userID}-{@sysID}"

	cacheFormatNotifySet	= "{sso}:notify:setlist:{@userID}-{sysID}-{@pi}-{@ps}"
	cacheFormatNotifySetCount = "{sso}:notify:setlist:{@userID}-{sysID}"
	cacheFormatNotifySetDel  = "{sso}:notify:setlist:*"
)

type ICacheNotify interface {
	SaveNotify(title, userID, sysID, pi, ps string,data db.QueryRows, count int) (err error)
	QueryNotify(title, userID, sysID, pi, ps string) (data db.QueryRows, count int, err error)
	FreshNotify() (err error)

	SaveNotifySet(userID,sysID string,pi,ps int,data db.QueryRows,count int) (err error)
	QueryNotifySet(userID,sysID string,pi,ps int) (data db.QueryRows,count int,err error)
	FreshNotifySet() (err error)
}

type CacheNotify struct {
	c         component.IContainer
	cacheTime int
}

//NewCacheSystem 创建缓存对象
func NewCacheNotify(c component.IContainer) *CacheNotify {
	return &CacheNotify{
		c:         c,
		cacheTime: 3600 * 24,
	}
}

func (l *CacheNotify) SaveNotify(title,userID,sysID,pi,ps string,data db.QueryRows,count int) (err error) {
	buff, err := json.Marshal(data)
	if err != nil {
		return err
	}
	cache := l.c.GetRegularCache()
	keyData := transform.Translate(cacheFormatNotify, "title", title, "userID", userID, "sysID",sysID, "pi", pi, "ps", ps)
	keyCount := transform.Translate(cacheFormatNotifyCount,"title",title,"userID",userID,"sysID",sysID)
	if err := cache.Set(keyData, string(buff), l.cacheTime); err != nil {
		return err
	}
	return cache.Set(keyCount, fmt.Sprint(count), l.cacheTime)
}

func (l *CacheNotify) QueryNotify(title, userID, sysID, pi, ps string) (data db.QueryRows, count int, err error){
	
	cache := l.c.GetRegularCache()
	keyData := transform.Translate(cacheFormatNotify, "title", title, "userID", userID, "sysID",sysID, "pi", pi, "ps", ps)
	keyCount := transform.Translate(cacheFormatNotifyCount,"title",title,"userID",userID,"sysID",sysID)
	v, err := cache.Get(keyData)
	if err != nil {
		return nil,0,err
	}
	if v == "" {
		return nil, 0,context.NewError(context.ERR_FORBIDDEN, "无数据")
	}
	var nmap db.QueryRows
	if err = json.Unmarshal([]byte(v), &nmap); err != nil {
		return nil, 0,err
	}
	c, err := cache.Get(keyCount)
	if err != nil {
		return nil, 0, err
	}
	return nmap, types.ToInt(c, 0), err
}

func (l *CacheNotify) FreshNotify() (err error){
	return l.c.GetRegularCache().Delete(cacheFormatNotifyDel)
}

func (l *CacheNotify)SaveNotifySet(userID,sysID string,pi,ps int,data db.QueryRows,count int) (err error){
	buff, err := json.Marshal(data)
	if err != nil {
		return err
	}
	cache := l.c.GetRegularCache()
	keyData := transform.Translate(cacheFormatNotifySet, "userID", userID, "sysID",sysID, "pi", pi, "ps", ps)
	keyCount := transform.Translate(cacheFormatNotifySetCount,"userID",userID,"sysID",sysID)
	if err := cache.Set(keyData, string(buff), l.cacheTime); err != nil {
		return err
	}
	return cache.Set(keyCount, fmt.Sprint(count), l.cacheTime)
}

func (l *CacheNotify)QueryNotifySet(userID,sysID string,pi,ps int) (data db.QueryRows,count int,err error){
	cache := l.c.GetRegularCache()
	keyData := transform.Translate(cacheFormatNotifySet, "userID", userID, "sysID",sysID, "pi", pi, "ps", ps)
	keyCount := transform.Translate(cacheFormatNotifySetCount,"userID",userID,"sysID",sysID)
	v, err := cache.Get(keyData)
	if err != nil {
		return nil,0,err
	}
	if v == "" {
		return nil, 0,context.NewError(context.ERR_FORBIDDEN, "无数据")
	}
	var nmap db.QueryRows
	if err = json.Unmarshal([]byte(v), &nmap); err != nil {
		return nil, 0,err
	}
	c, err := cache.Get(keyCount)
	if err != nil {
		return nil, 0, err
	}
	return nmap, types.ToInt(c, 0), err
}

func (l *CacheNotify)FreshNotifySet() (err error){
	return l.c.GetRegularCache().Delete(cacheFormatNotifySetDel)
}

