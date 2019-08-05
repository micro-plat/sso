package service

import (
	"github.com/micro-plat/lib4go/db"
	"github.com/micro-plat/sso/sdk/model"
)

//GetUserInfoByName 根据用户名获取用户信息
func GetUserInfoByName(userName string) (info db.QueryRow, err error) {
	return nil, nil
}

//CheckLoginCodeAndReturnUserInfo 验证用户登录的code
func CheckLoginCodeAndReturnUserInfo(code string) (res *model.LoginState, err error) {
	return nil, nil
}
