package service

import (
	"fmt"
	"strings"
	"time"

	"github.com/micro-plat/lib4go/net"
	"github.com/micro-plat/lib4go/security/md5"
	"github.com/micro-plat/lib4go/types"
	"github.com/micro-plat/sso/sso/model"
)

//GetUserMenu 查询用户在某个系统下的菜单数据
func GetUserMenu(conf *model.Config, userID int) (*[]*Menu, error) {
	values := net.NewValues()
	values.Set("user_id", types.GetString(userID))
	values.Set("ident", conf.Ident)
	values.Set("timestamp", types.GetString(time.Now().Unix()))

	values = values.Sort()
	raw := values.Join("=", "&") + "&key=" + conf.Secret
	fmt.Println(raw)
	values.Set("sign", strings.ToUpper(md5.Encrypt(raw)))

	menu := &[]*Menu{}
	result, err := remoteRequest(conf.Host, model.UserMenuUrl, values.Join("=", "&"), menu)
	if err != nil {
		return nil, err
	}
	return result.(*[]*Menu), nil
}
