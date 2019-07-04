package member

import (
	"encoding/json"

	"github.com/micro-plat/hydra/component"
	"github.com/micro-plat/hydra/context"
	"github.com/micro-plat/lib4go/transform"
	"github.com/micro-plat/lib4go/utility"
	"github.com/micro-plat/sso/mgrapi/modules/model"
)

//ICodeMember 用户数据库缓存
type ICodeMember interface {
	Query(code string) (ls *model.LoginState, err error)
	Save(s *model.LoginState) (string, error)
	ExchangeCode(code string, s *model.LoginState) (newCode string, err error)
}

//CodeMember 控制用户登录
type CodeMember struct {
	c           component.IContainer
	cacheFormat string
	cacheTime   int
}

//NewCodeMember 创建登录对象
func NewCodeMember(c component.IContainer) *CodeMember {
	return &CodeMember{
		c:           c,
		cacheTime:   30000,
		cacheFormat: "sso:login:code:{@code}",
	}
}

//Save 缓存用户信息
func (l *CodeMember) Save(s *model.LoginState) (code string, err error) {
	code = utility.GetGUID()
	buff, err := json.Marshal(s)
	if err != nil {
		return "", err
	}
	cache := l.c.GetRegularCache()
	key := transform.Translate(l.cacheFormat, "code", code)
	return code, cache.Set(key, string(buff), l.cacheTime)
}

//Query 用户登录
func (l *CodeMember) Query(code string) (ls *model.LoginState, err error) {
	//从缓存中查询用户数据
	cache := l.c.GetRegularCache()
	key := transform.Translate(l.cacheFormat, "code", code)
	v, err := cache.Get(key)
	if err != nil {
		return nil, err
	}
	if v == "" {
		return nil, context.NewError(context.ERR_FORBIDDEN, "code无效")
	}
	if err = json.Unmarshal([]byte(v), &ls); err != nil {
		return nil, err
	}
	//cache.Delete(key)
	return ls, err
}

// 删除旧code,生成新code
func (l *CodeMember) ExchangeCode(code string, s *model.LoginState) (newCode string, err error) {
	// 从缓存中查询用户数据
	cache := l.c.GetRegularCache()
	key := transform.Translate(l.cacheFormat, "code", code)
	v, err := cache.Get(key)
	if err != nil {
		return "", err
	}
	if v == "" {
		return "", context.NewError(context.ERR_FORBIDDEN, "code无效")
	}
	// 删除code
	if err = cache.Delete(key); err != nil {
		return "", err
	}
	// 生成新code
	newCode = utility.GetGUID()
	buff, err := json.Marshal(s)
	if err != nil {
		return "", err
	}
	key = transform.Translate(l.cacheFormat, "code", newCode)
	return newCode, cache.Set(key, string(buff), l.cacheTime)
}
