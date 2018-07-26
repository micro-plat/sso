// +build prod

package main

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"

	"github.com/micro-plat/hydra/component"
)

//bindConf 绑定启动配置， 启动时检查注册中心配置是否存在，不存在则引导用户输入配置参数并自动创建到注册中心
func (s *SSO) install() {
	s.IsDebug = false
	s.Conf.SetInput("email", "邮箱地址", "接收账户确认邮件时使用", func(v string) (string, error) {
		if !strings.Contains(v, "@") {
			return "", fmt.Errorf("请输入正确的邮箱地址")
		}
		return strings.Replace(v, "@", "\\@", -1), nil
	})
	s.Conf.SetInput("#wx_host_name", "服务器域名", "以http开头")

	s.Conf.API.SetMainConf(`{"address":":9091"}`)
	s.Conf.API.SetSubConf("app", `
			{
				"qrlogin-check-url":"#wx_host_name/member/wxlogin",
				"wx-login-url":"#wx_host_name/member/wxlogin",
				"appid":"#appid",
				"secret":"#app_secret",
				"wechat-url":"http://#wx_token_server_host/#appid/wechat/token/get"
			}			
			`)
	s.Conf.API.SetSubConf("header", `
				{
					"Access-Control-Allow-Origin": "*", 
					"Access-Control-Allow-Methods": "GET,POST,PUT,DELETE,PATCH,OPTIONS", 
					"Access-Control-Allow-Headers": "sso", 
					"Access-Control-Allow-Credentials": "true"
				}
			`)

	s.Conf.API.SetSubConf("auth", `
		{
			"jwt": {
				"exclude": ["/sso/login","/sso/login/code","/sso/wxcode/get","/sso/sys/get","/qrcode/login","/qrcode/login/put"],
				"expireAt": 36000,
				"mode": "HS512",
				"name": "sso",
				"secret": "ef1a8839cb511780903ff6d5d79cf8f8"
			}
		}
		`)

	s.Conf.WS.SetSubConf("app", `
			{
				"qrlogin-check-url":"http://#wx_host_name/member/wxlogin",
				"wx-login-url":"http://#wx_host_name/member/wxlogin",
				"appid":"#appid",
				"secret":"#app_secret",
				"wechat-url":"http://#wx_token_server_host/#appid/wechat/token/get"
			}		
			`)
	s.Conf.Plat.SetVarConf("db", "db", `{			
			"provider":"ora",
			"connString":"#db_connection_string",
			"maxOpen":100,
			"maxIdle":10,
			"lifeTime":100		
	}`)

	s.Conf.Plat.SetVarConf("cache", "cache", `
		{
			"proto":"redis",
			"addrs":[
					#redis_server
			],
			"db":1,
			"dial_timeout":10,
			"read_timeout":10,
			"write_timeout":10,
			"pool_size":10
	}
		`)

	//自定义安装程序
	s.Conf.API.Installer(func(c component.IContainer) error {
		if !s.Conf.Confirm("创建数据库表结构,添加基础数据?") {
			return nil
		}
		path, err := getSQLPath()
		if err != nil {
			return err
		}
		sqls, err := s.Conf.GetSQL(path)
		if err != nil {
			return err
		}
		db, err := c.GetDB()
		if err != nil {
			return err
		}
		for _, sql := range sqls {
			if sql != "" {
				if _, q, _, err := db.Execute(sql, map[string]interface{}{}); err != nil {
					if !strings.Contains(err.Error(), "ORA-00942") {
						s.Conf.Log.Errorf("执行SQL失败： %v %s\n", err, q)
					}
				}
			}
		}
		return nil
	})

}

//getSQLPath 获取getSQLPath
func getSQLPath() (string, error) {
	gopath := os.Getenv("GOPATH")
	if gopath == "" {
		return "", fmt.Errorf("未配置环境变量GOPATH")
	}
	path := strings.Split(gopath, ";")
	if len(path) == 0 {
		return "", fmt.Errorf("环境变量GOPATH配置的路径为空")
	}
	return filepath.Join(path[0], "src/github.com/micro-plat/sso/modules/const/sql"), nil
}
