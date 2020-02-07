package permission

import (
	"github.com/micro-plat/hydra/component"
	"github.com/micro-plat/hydra/context"
	"github.com/micro-plat/sso/apiserver/apiserver/modules/logic"
	"github.com/micro-plat/sso/apiserver/apiserver/modules/model"
)

//DataPerssionHandler 数据权限相关功能
type DataPerssionHandler struct {
	c   component.IContainer
	sys logic.IDataPermissionLogic
}

//NewDataPerssionHandler new
func NewDataPerssionHandler(container component.IContainer) (u *DataPerssionHandler) {
	return &DataPerssionHandler{
		c:   container,
		sys: logic.NewDataPermissionLogic(container),
	}
}

/*
* SyncHandle: 同步子系统的[数据权限] 数据 (如商品分类等)
 */
func (u *DataPerssionHandler) SyncHandle(ctx *context.Context) (r interface{}) {
	ctx.Log.Info("-------同步子系统的[数据权限] 数据------")

	ctx.Log.Info("-------验证数据的有效性------")
	var req model.DataPermissionSyncReq
	if err := ctx.Request.Bind(&req); err != nil {
		return context.NewError(context.ERR_NOT_ACCEPTABLE, err)
	}

	ctx.Log.Info("-------保存同步数据------")
	err := u.sys.SyncDataPermission(req)
	if err != nil {
		return err
	}

	ctx.Log.Info("------返回结果------")
	return "success"
}

/*
* GetHandle: 获取当前用户可用的【数据权限】数据(1,2,3,4)
 */
func (u *DataPerssionHandler) GetHandle(ctx *context.Context) (r interface{}) {
	ctx.Log.Info("-------获取当前用户可用的【数据权限】数据------")

	ctx.Log.Info("-------验证数据的有效性------")
	var req model.DataPermissionGetReq
	if err := ctx.Request.Bind(&req); err != nil {
		return context.NewError(context.ERR_NOT_ACCEPTABLE, err)
	}

	ctx.Log.Info("-------获取数据------")
	result, err := u.sys.GetUserDataPermission(req)
	if err != nil {
		return err
	}

	ctx.Log.Info("------返回结果------")
	return result
}
