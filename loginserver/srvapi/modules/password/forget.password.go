package password

import (
	"fmt"

	"github.com/micro-plat/hydra/components"
	"github.com/micro-plat/sso/loginserver/srvapi/modules/const/sqls"
)

// IForgetPassword xx
type IForgetPassword interface {
	ForgetPassword(source, userName, newPassword string) (err error)
}

// ForgetPassword 操作日志
type ForgetPassword struct {
}

// NewForgetPassword xx
func NewForgetPassword() *ForgetPassword {
	return &ForgetPassword{}
}

// ForgetPassword 忘记密码
func (o *ForgetPassword) ForgetPassword(source, userName, newPassword string) (err error) {
	db := components.Def.DB().GetRegularDB()
	lastInsertID, affectedRow, err := db.Executes(sqls.ForgetPassword, map[string]interface{}{
		"source":    source,
		"user_name": userName,
		"password":  newPassword,
	})
	if err != nil {
		return fmt.Errorf("忘记密码数据发生错误(err:%v)lastInsertID:%v,受影响的行数：%v", err, lastInsertID, affectedRow)
	}
	return nil
}
