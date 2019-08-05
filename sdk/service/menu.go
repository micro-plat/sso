package service

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"time"

	"github.com/micro-plat/lib4go/db"
	"github.com/micro-plat/sso/sdk/model"
	"github.com/micro-plat/sso/sdk/util"
	"github.com/micro-plat/sso/sdk/util/logger"
)

//GetUserMenu 查询用户在某个系统下的菜单数据
func GetUserMenu(userID, systemID int) ([]*model.Menu, error) {
	params := make(db.QueryRow)
	params["user_id"] = userID
	params["system_id"] = systemID
	params["ident"] = model.SysInfoConfig.Ident
	params["timestamp"] = time.Now().Unix()

	_, sign := util.MakeSign(params, model.SysInfoConfig.Secret)
	params["sign"] = sign

	return remoteMenuQuery(params)
}

func remoteMenuQuery(m db.QueryRow) ([]*model.Menu, error) {
	url := fmt.Sprintf("%s?user_id=%s&system_id=%s&ident=%s&timestamp=%s&sign=%s",
		model.SysInfoConfig.ApiHost+"/subsys/menu", m.GetString("user_id"),
		m.GetString("system_id"), m.GetString("ident"), m.GetString("timestamp"),
		m.GetString("sign"))

	logger.Infof("菜单请求url: %s", url)
	request, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return nil, err
	}

	client := http.Client{
		Timeout: 5 * time.Second,
	}
	resp, err := client.Do(request)
	if err != nil {
		return nil, fmt.Errorf("请求:%v失败(%v)", url, err)
	}

	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("读取远程数据失败 %s %v", url, err)
	}
	var menus []*model.Menu
	err = json.Unmarshal(body, &menus)
	if err != nil {
		return nil, fmt.Errorf("解析返回结果失败 %s：%v(%s)", url, err, string(body))
	}
	return menus, nil
}

// r.Micro("/subsys/menu", user.NewMenuHandler)     //子系统获取菜单数据
// 	r.Micro("/subsys/user", user.NewUserInfoHandler) //子系统,获取用户信息
// 	// r.Micro("/subsys/pwd", user.NewPwdHandler)       //子系统,修改密码
// 	r.Micro("/subsys/info", system.NewInfoHandler) //子系统,获取系统信息
