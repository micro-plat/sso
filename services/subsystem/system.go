
package subsystem

import (
	"fmt"
	"github.com/micro-plat/hydra/component"
	"github.com/micro-plat/hydra/context"
	sub "github.com/micro-plat/sso/modules/subsystem"
)

type SystemHandler struct {
	container component.IContainer
	subLib sub.ISystem
}

type UpdateSystemInput struct {
	Name string `form:"name"`
	Addr string `form:"addr"`
	Time_out string `form:"time_out"`
	Logo string `form:"logo"`
	Style string `form:"style"`
	Theme string `form:"theme"`
}

func NewSystemHandler(container component.IContainer) (u *SystemHandler) {
	return &SystemHandler{
		container: container,
		subLib:   sub.NewSystem(container),
	}
}


func (u *SystemHandler) Handle(ctx *context.Context) (r interface{}) {
	return "default"
}

//获取系统管理列表
func (u *SystemHandler) GetHandle(ctx *context.Context) (r interface{}){
	ctx.Log.Info("--------查询系统管理数据--------")
	ctx.Log.Info("1.从数据库查询数据--------")
	page := ctx.Request.GetInt("page",1)
	rows, count, err := u.subLib.Query(page)
	if err != nil {
		return context.NewError(context.ERR_NOT_IMPLEMENTED, err)
	}

	ctx.Log.Info("2.返回数据。")
	return map[string]interface{}{
		"count": count.(string),
		"list":  rows,
	}
}

//添加系统
func (u *SystemHandler) PostHandle(ctx *context.Context) (r interface{}){
	ctx.Log.Info("------添加系统管理数据------")
	ctx.Log.Info("1. 参数检查")
	var input UpdateSystemInput
	if err := ctx.Request.Bind(&input); err != nil{
		return nil
	}
	dbInput := map[string]interface{}{
		"name":    input.Name,
		"addr": input.Addr,
		"time_out":   input.Time_out,
		"logo":  input.Logo,
		"style": input.Style,
		"theme": input.Theme,
	}

	ctx.Log.Info("2.添加数据库查询--------")
	err := u.subLib.Add(dbInput)
	if err != nil {
		return context.NewError(context.ERR_NOT_IMPLEMENTED, err)
	}
	ctx.Log.Info("3.返回数据。")
	return map[string]interface{}{
		"msg":  "success",
	}
}
//删除系统管理ByID

func (u *SystemHandler) DeleteHandle(ctx *context.Context)(r interface{}){
	ctx.Log.Info("------删除系统管理数-----")
	ctx.Log.Info("1.参数检查")
	Id := ctx.Request.GetInt("id")

	ctx.Log.Info("2.从数据库删除数据")
	ctx.Log.Info("请求参数 id：",Id)
	if Id == 0 {
		return context.NewError(context.ERR_NOT_IMPLEMENTED,fmt.Errorf("不能删除当前系统，系统编号：%v",Id))
	}
	err := u.subLib.DeleteByID(Id)
	if err != nil {
		return context.NewError(context.ERR_NOT_IMPLEMENTED, err)
	}

	return map[string]interface{}{
		"msg":  "success",
	}
}







