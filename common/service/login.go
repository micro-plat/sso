package service

import (
	"github.com/micro-plat/lib4go/logger"
	l "github.com/micro-plat/sso/common/module/login"
	"github.com/micro-plat/sso/common/module/model"
)

//Login 用户名密码登录
func Login(log logger.ILogger, req model.LoginReq) (*model.LoginState, error) {
	server := l.NewLoginLogic()

	log.Info("1: 判断系统是否被禁用")
	ident := req.Ident
	if err := server.CheckSystemStatus(ident); err != nil {
		return nil, err
	}

	log.Info("2: 判断用户是否被锁定, 锁定时间过期后要解锁")
	userName := req.UserName
	if err := server.CheckUserIsLocked(userName); err != nil {
		return nil, err
	}

	log.Info("3: 判断用户输入的验证码")
	if err := server.CheckWxValidCode(userName, req.Wxcode); err != nil {
		return nil, err
	}

	log.Info("4:处理用户账号登录")
	member, err := server.Login(userName, req.Password, ident)
	if err != nil {
		return nil, err
	}

	log.Info("5:记录用户登录时间")
	if err = server.UpdateUserLoginTime(member.UserID); err != nil {
		log.Errorf("记录用户登录时间出错: %+v", err)
	}

	return member, nil
}
