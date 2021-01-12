package member

import (
	"fmt"
	"github.com/micro-plat/hydra"
	"github.com/micro-plat/sso/loginserver/srvapi/modules/logic"
)

//TagHandler Tag对象(按钮级)
type TagHandler struct {
	m logic.IMenuLogic
}

//NewTagHandler 创建Tag查询对象
func NewTagHandler() (u *TagHandler) {
	return &TagHandler{
		m: logic.NewMenuLogic(),
	}
}

/*
* Handle: 查询用户在某个系统下的tag数据
* ident:子系统标识
* user_id:用户标识
 */
func (u *TagHandler) Handle(ctx hydra.IContext) (r interface{}) {
	ctx.Log().Info("-------获取用户有权限的tag列表------")

	if err := ctx.Request().Check("user_id"); err != nil {
		return fmt.Errorf("参数错误：%v", err)
	}

	ctx.Log().Info("1. 获取用户在指定系统的tag列表数据")
	data, err := u.m.GetTags(ctx.Request().GetInt64("user_id"), ctx.Request().GetString("ident"))
	if err != nil {
		return err
	}

	ctx.Log().Info("2. 返回tag数据")
	return data
}
