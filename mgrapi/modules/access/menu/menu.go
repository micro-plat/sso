package menu

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"regexp"
	"strings"
	"time"

	"github.com/micro-plat/hydra/component"
	"github.com/micro-plat/hydra/context"
	"github.com/micro-plat/lib4go/db"
	"github.com/micro-plat/lib4go/types"
	"github.com/micro-plat/sso/mgrapi/modules/const/sqls"
	"github.com/micro-plat/sso/mgrapi/modules/model"
	"github.com/micro-plat/sso/mgrapi/modules/util"
)

type IMenu interface {
	Query(uid int64, sysid int) ([]map[string]interface{}, error)
	Verify(uid int64, sysid int, menuURL string, method string) error
	QueryMenuFromSso(uid int64, sysid int) ([]*model.Menu, error)
}

type Menu struct {
	c    component.IContainer
	http *http.Client
}

func NewMenu(c component.IContainer) *Menu {
	return &Menu{
		c:    c,
		http: &http.Client{},
	}
}

//Query 获取用户指定系统的菜单信息
func (l *Menu) Query(uid int64, sysid int) ([]map[string]interface{}, error) {
	db := l.c.GetRegularDB()
	data, _, _, err := db.Query(sqls.QueryUserMenus, map[string]interface{}{
		"user_id": uid,
		"sys_id":  sysid,
	})
	if err != nil {
		return nil, err
	}
	result := make([]map[string]interface{}, 0, 4)
	for _, row1 := range data {
		if row1.GetInt("parent") == 0 && row1.GetInt("level_id") == 1 {
			children1 := make([]map[string]interface{}, 0, 4)
			for _, row2 := range data {
				if row2.GetInt("parent") == row1.GetInt("id") && row2.GetInt("level_id") == 2 {
					children2 := make([]map[string]interface{}, 0, 8)
					for _, row3 := range data {
						if row3.GetInt("parent") == row2.GetInt("id") && row3.GetInt("level_id") == 3 {
							children2 = append(children2, row3)
						}
					}
					children1 = append(children1, row2)
					row2["children"] = children2
				}
			}
			row1["children"] = children1
			result = append(result, row1)
		}
	}

	return result, nil
}

// QueryMenuFromSso xx
func (l *Menu) QueryMenuFromSso(uid int64, sysid int) ([]*model.Menu, error) {
	config := model.GetConf(l.c)
	m := map[string]interface{}{
		"user_id":   uid,
		"system_id": sysid,
		"ident":     "sso",
		"timestamp": time.Now().Unix(),
	}
	_, sign := util.MakeSign(m, config.Secret)
	m["sign"] = sign

	return l.remoteMenuQuery(config, m)
}

//Verify 获取用户指定系统的菜单信息
func (l *Menu) Verify(uid int64, sysid int, menuURL string, method string) error {
	db := l.c.GetRegularDB()
	//根据用户名密码，查询用户信息

	url, funcs, err := getFuncs(menuURL, method)
	data, _, _, err := db.Scalar(sqls.QueryUserMenu, map[string]interface{}{
		"user_id": uid,
		"sys_id":  sysid,
		"path":    "'" + url + "'",
	})
	if err != nil || types.GetInt(data) != 1 {
		return context.NewError(context.ERR_FORBIDDEN, fmt.Errorf("未查找到菜单 %v", err))
	}
	if len(funcs) == 0 {
		return nil
	}
	data, _, _, err = db.Scalar(sqls.QueryUserMenu, map[string]interface{}{
		"user_id": uid,
		"sys_id":  sysid,
		"path":    "'" + strings.Join(funcs, "','") + "'",
	})
	if err != nil || types.GetInt(data) != len(funcs) {
		return context.NewError(context.ERR_FORBIDDEN, fmt.Errorf("未查找到菜单 %v", err))
	}
	return nil
}

func getFuncs(urlParams string, method string) (string, []string, error) {
	funcs := make([]string, 0, 2)
	items := strings.Split(urlParams, "#")
	if len(items) == 0 {
		return "", nil, fmt.Errorf("传入的页面地址不能为空")
	}
	url := items[0]
	if len(items) == 1 {
		return url, nil, nil
	}
	word := regexp.MustCompile(`\[?\w*[\:]?\w+\]?`)
	names := word.FindAllString(strings.Join(items[1:], "#"), -1)
	for _, n := range names {
		text := strings.Split(strings.Trim(strings.Trim(n, "]"), "["), ":")
		if len(text) == 1 {
			funcs = append(funcs, text[0])
		} else {
			if strings.ToUpper(text[0]) == strings.ToUpper(method) {
				funcs = append(funcs, text[1])
			}
		}
	}
	funcs = append(funcs, items[0])
	return url, funcs, nil
}

func (l *Menu) remoteMenuQuery(config *model.Conf, m db.QueryRow) ([]*model.Menu, error) {
	url := fmt.Sprintf(
		"%s?user_id=%s&system_id=%s&ident=%s&timestamp=%s&sign=%s",
		config.GetMenuURL(),
		m.GetString("user_id"),
		m.GetString("system_id"),
		m.GetString("ident"),
		m.GetString("timestamp"),
		m.GetString("sign"))

	request, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return nil, err
	}

	resp, err := l.http.Do(request)
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
