package logic

import (
	"github.com/micro-plat/lib4go/errs"
	"github.com/micro-plat/lib4go/types"
	"github.com/micro-plat/sso/mgrserver/mgrapi/modules/access/menu"
	"github.com/micro-plat/sso/mgrserver/mgrapi/modules/model"
	"github.com/micro-plat/sso/sso/errorcode"
)

//ISystemMenuLogic interface
type ISystemMenuLogic interface {
	Export(sysID int) (s types.XMaps, err error)
	Import(req *model.ImportReq) error
}

//SystemMenuLogic 系统菜单相关操作
type SystemMenuLogic struct {
	db menu.IDbSystemMenu
}

//NewSystemMenuLogic new
func NewSystemMenuLogic() *SystemMenuLogic {
	return &SystemMenuLogic{
		db: menu.NewDbSystemMenu(),
	}
}

//Export 导出系统菜单
func (u *SystemMenuLogic) Export(sysID int) (s types.XMaps, err error) {
	if s, err = u.db.Export(sysID); err != nil {
		return nil, err
	}
	return s, nil
}

//Import 导入系统菜单
func (u *SystemMenuLogic) Import(req *model.ImportReq) error {
	flag, err := u.db.Exists(req.Id)
	if err != nil {
		return err
	}
	if flag {
		return errs.NewError(errorcode.ERR_SYSTEM_HASMENUS, "当前系统下面已存在菜单数据,不能导入")
	}

	return u.db.Import(req)
}
