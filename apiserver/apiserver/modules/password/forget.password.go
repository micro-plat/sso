package password

import (
	"fmt"

	"github.com/micro-plat/hydra/component"
	"github.com/micro-plat/sso/apiserver/apiserver/modules/const/sqls"
)

// IForgetPassword xx
type IForgetPassword interface {
	ForgetPassword(source, userName, newPassword string) (err error)
}

// ForgetPassword 操作日志
type ForgetPassword struct {
	c component.IContainer
}

// NewForgetPassword xx
func NewForgetPassword(c component.IContainer) *ForgetPassword {
	return &ForgetPassword{
		c: c,
	}
}

// ForgetPassword 忘记密码
func (o *ForgetPassword) ForgetPassword(source, userName, newPassword string) (err error) {
	db := o.c.GetRegularDB()
	lastInsertID, affectedRow, q, a, err := db.Executes(sqls.ForgetPassword, map[string]interface{}{
		"source":    source,
		"user_name": userName,
		"password":  newPassword,
	})
	if err != nil {
		return fmt.Errorf("忘记密码数据发生错误(err:%v),sql:%s,参数：%v,lastInsertID:%v,受影响的行数：%v", err, q, a, lastInsertID, affectedRow)
	}
	return nil
}
