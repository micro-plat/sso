package user

import (
	"fmt"
	"strings"

	"github.com/micro-plat/hydra/component"
	"github.com/micro-plat/lib4go/security/md5"
	"github.com/micro-plat/sso/apiserver/modules/const/enum"
	"github.com/micro-plat/sso/apiserver/modules/const/sqls"
)

var _ IDbUser = &DbUser{}

type IDbUser interface {
	ChangePwd(user_id int, expassword string, newpassword string) (err error)
}

type DbUser struct {
	c component.IContainer
}

func NewDbUser(c component.IContainer) *DbUser {
	return &DbUser{
		c: c,
	}
}

// ChangePwd 修改密码
func (u *DbUser) ChangePwd(user_id int, expassword string, newpassword string) (err error) {
	db := u.c.GetRegularDB()

	//获取旧密码
	data, q, a, err := db.Query(sqls.QueryOldPwd, map[string]interface{}{
		"user_id": user_id,
	})
	if err != nil || data.Get(0).GetInt("changepwd_times") >= enum.MaxFailCnt {
		return fmt.Errorf("获取数据错误或密码修改超过限制(err:%v),sql:%s,参数：%v", err, q, a)
	}
	if strings.ToLower(md5.Encrypt(expassword)) != strings.ToLower(data.Get(0).GetString("password")) {
		return fmt.Errorf("原密码错误")
	}
	dbTrans, err := db.Begin()
	if err != nil {
		return fmt.Errorf("开启DB事务出错(err:%v)", err)
	}
	_, q, a, err = dbTrans.Execute(sqls.SetNewPwd, map[string]interface{}{
		"user_id":  user_id,
		"password": md5.Encrypt(newpassword),
	})
	if err != nil {
		dbTrans.Rollback()
		return fmt.Errorf("设置密码错误(err:%v),sql:%s,参数：%v", err, q, a)
	}
	dbTrans.Commit()
	return nil
}
