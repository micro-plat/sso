package menu

import (
	"fmt"

	"github.com/micro-plat/hydra/components"
	"github.com/micro-plat/lib4go/db"
	"github.com/micro-plat/lib4go/types"
	"github.com/micro-plat/sso/mgrserver/webserver/modules/const/sqls"
	"github.com/micro-plat/sso/mgrserver/webserver/modules/model"
)

//IDbSystemMenu interface
type IDbSystemMenu interface {
	Export(sysID int) (s db.QueryRows, err error)
	Import(req *model.ImportReq) error
	Exists(sysID string) (bool, error)
}

//DbSystemMenu 系统菜单
type DbSystemMenu struct {
}

//NewDbSystemMenu new
func NewDbSystemMenu() *DbSystemMenu {
	return &DbSystemMenu{}
}

//Export 导出菜单
func (l *DbSystemMenu) Export(sysID int) (s db.QueryRows, err error) {
	db := components.Def.DB().GetRegularDB()
	data, _, _, err := db.Query(sqls.QuerySystemMenuInfo, map[string]interface{}{
		"sys_id": sysID,
	})
	if err != nil {
		return nil, fmt.Errorf("Export menu出错: sys_id:%d, err:%+v", sysID, err)
	}

	return data, nil
}

//Exists 判断一个系统下面是否有菜单数据
func (l *DbSystemMenu) Exists(sysID string) (bool, error) {
	db := components.Def.DB().GetRegularDB()
	count, _, _, err := db.Scalar(sqls.ExistsSystemMenu, map[string]interface{}{
		"sys_id": sysID,
	})
	if err != nil {
		return false, fmt.Errorf("判断系统下面是否有菜单数据出错: sys_id:%s, err:%+v", sysID, err)
	}
	return types.GetInt(count) > 0, nil
}

//Import 导入菜单(现在菜单数据有4级:1大分类.2:小分类,3:菜单,4:权限(tag))
func (l *DbSystemMenu) Import(req *model.ImportReq) error {
	first := l.getLevelData(1, req.Menus)
	oldIDMaping := make(map[string]int64)

	db := components.Def.DB().GetRegularDB()
	trans, err := db.Begin()
	if err != nil {
		return fmt.Errorf("导入菜单,生成事务出错: %+v", err)
	}

	if err = l.insertLevelData(trans, oldIDMaping, req.Id, 1, first); err != nil {
		trans.Rollback()
		return err
	}

	second := l.getLevelData(2, req.Menus)
	if err = l.insertLevelData(trans, oldIDMaping, req.Id, 2, second); err != nil {
		trans.Rollback()
		return err
	}

	third := l.getLevelData(3, req.Menus)
	if err = l.insertLevelData(trans, oldIDMaping, req.Id, 3, third); err != nil {
		trans.Rollback()
		return err
	}

	fourth := l.getLevelData(4, req.Menus)
	if err = l.insertLevelData(trans, oldIDMaping, req.Id, 3, fourth); err != nil {
		trans.Rollback()
		return err
	}

	trans.Commit()
	return nil
}

//GetLevelData 查询某个级次下面的菜单数据
func (l *DbSystemMenu) getLevelData(levelID int, data []model.MenuInfo) (result []model.MenuInfo) {
	result = make([]model.MenuInfo, 0)
	for _, val := range data {
		if val.LevelID == types.GetString(levelID) {
			result = append(result, val)
		}
	}
	return result
}

//insertLevelData 插入菜单数据
func (l *DbSystemMenu) insertLevelData(trans db.IDBTrans, oldIDMaping map[string]int64, sysID string, levelID int, data []model.MenuInfo) error {
	for _, val := range data {
		var newParent int64
		if levelID != 1 {
			flag := true
			newParent, flag = oldIDMaping[val.Parent]
			if !flag {
				continue
			}
		}
		id, _, q, a, err := trans.Executes(sqls.AddSystemMenu, map[string]interface{}{
			"name":     val.Name,
			"parent":   newParent,
			"sys_id":   sysID,
			"level_id": val.LevelID,
			"icon":     val.Icon,
			"path":     val.Path,
			"enable":   val.Enable,
			"sortrank": val.Sortrank,
			"is_open":  val.IsOpen,
		})
		if err != nil {
			trans.Rollback()
			return fmt.Errorf("导入菜单出错: sysID:%s, levelID:%d, q:%s,a:%+v,err:%+v", sysID, levelID, q, a, err)
		}
		if levelID != 4 {
			oldIDMaping[val.ID] = id
		}
	}
	return nil
}
