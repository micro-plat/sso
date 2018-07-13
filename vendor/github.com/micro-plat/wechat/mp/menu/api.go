package menu

import (
	"github.com/micro-plat/wechat/mp"
)

// 创建自定义菜单.
func Create(clt *mp.Context, menu *Menu) (err error) {
	const incompleteURL = "https://api.weixin.qq.com/cgi-bin/menu/create?access_token="

	var result mp.Error
	if err = clt.PostJSON(incompleteURL, menu, &result); err != nil {
		return
	}
	if result.ErrCode != mp.ErrCodeOK {
		err = &result
		return
	}
	return
}

// 查询自定义菜单.
func Get(clt *mp.Context) (menu *Menu, conditionalMenus []Menu, err error) {
	const incompleteURL = "https://api.weixin.qq.com/cgi-bin/menu/get?access_token="

	var result struct {
		mp.Error
		Menu             Menu   `json:"menu"`
		ConditionalMenus []Menu `json:"conditionalmenu"`
	}
	if err = clt.GetJSON(incompleteURL, &result); err != nil {
		return
	}
	if result.ErrCode != mp.ErrCodeOK {
		err = &result.Error
		return
	}
	menu = &result.Menu
	conditionalMenus = result.ConditionalMenus
	return
}

// 删除自定义菜单.
func Delete(clt *mp.Context) (err error) {
	const incompleteURL = "https://api.weixin.qq.com/cgi-bin/menu/delete?access_token="

	var result mp.Error
	if err = clt.GetJSON(incompleteURL, &result); err != nil {
		return
	}
	if result.ErrCode != mp.ErrCodeOK {
		err = &result
		return
	}
	return
}
