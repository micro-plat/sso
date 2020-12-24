package permission

import (
	"net/http"

	"github.com/micro-plat/hydra"
	"github.com/micro-plat/lib4go/errs"
	"github.com/micro-plat/sso/loginserver/apiserver/modules/logic"
	"github.com/micro-plat/sso/loginserver/apiserver/modules/model"
)

//DataPerssionHandler 数据权限相关功能
type DataPerssionHandler struct {
	sys logic.IDataPermissionLogic
}

//NewDataPerssionHandler new
func NewDataPerssionHandler() (u *DataPerssionHandler) {
	return &DataPerssionHandler{
		sys: logic.NewDataPermissionLogic(),
	}
}

//ObtainHandle: 获取当前用户数据权限的规则信息
func (u *DataPerssionHandler) ObtainHandle(ctx hydra.IContext) (r interface{}) {
	ctx.Log().Info("-------获取当前用户可用的【数据权限】数据------")

	ctx.Log().Info("-------验证数据的有效性------")
	var req model.DataPermissionGetReq
	if err := ctx.Request().Bind(&req); err != nil {
		return errs.NewError(http.StatusNotAcceptable, err)
	}

	ctx.Log().Info("-------获取数据------")
	result, err := u.sys.GetUserDataPermissionConfigs(req)
	if err != nil {
		return err
	}

	ctx.Log().Info("------返回结果------")
	return result
}
