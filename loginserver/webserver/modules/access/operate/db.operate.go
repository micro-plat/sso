package operate

import (
	"fmt"

	"github.com/micro-plat/hydra/components"
	"github.com/micro-plat/sso/loginserver/webserver/modules/const/sqls"
	"github.com/micro-plat/sso/loginserver/webserver/modules/model"
)

type IDbOperate interface {
	// 登录行为
	LoginOperate(m *model.LoginState) (err error)
}

type DbOperate struct {
}

func NewDbOperate() *DbOperate {
	return &DbOperate{}
}

//LoginOperate 登录行为
func (d *DbOperate) LoginOperate(m *model.LoginState) (err error) {
	db := components.Def.DB().GetRegularDB()
	params := map[string]interface{}{
		"type":    10,
		"sys_id":  m.SystemID,
		"user_id": m.UserID,
		"content": fmt.Sprintf(`{"desc":"%s","data":"%v"}`, "用户登录", m.UserName),
	}
	_, q, a, err := db.Execute(sqls.AddOperate, params)
	if err != nil {
		return fmt.Errorf("添加登录行为数据发生错误(err:%v),sql:%s,输入参数:%v,", err, q, a)
	}
	return nil

}
