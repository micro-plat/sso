package notify

import (
	"fmt"

	"github.com/micro-plat/hydra/component"
	"github.com/micro-plat/hydra/context"
	"github.com/micro-plat/lib4go/db"
	"github.com/micro-plat/lib4go/types"
	"github.com/micro-plat/lib4go/utility"
	"github.com/micro-plat/sso/apiserver/modules/const/sql"
)

type UserNotifyInput struct {
	Title  string `form:"title" json:"title"`
	UserID string `form:"user_id" json:"user_id"`
	SysID  string `form:"sys_id" json:"sys_id"`
	Pi     string `form:"pi" json:"pi" valid:"required"`
	Ps     string `form:"ps" json:"ps" valid:"required"`
}

type SettingsInput struct {
	Keywords string `form:"keywords" json:"keywords" valid:"required"`
	Level    string `form:"level_id" json:"level_id" valid:"required"`
	UserID   string `form:"user_id" json:"user_id"`
	SysID    string `form:"sys_id" json:"sys_id"`
}

type EditSettingsInput struct {
	ID       string `form:"id" json:"id" valid:"required"`
	Keywords string `form:"keywords" json:"keywords" valid:"required"`
	Level    string `form:"level_id" json:"level_id" valid:"required"`
	Status   string `form:"status" json:"status" valid:"required"`
}

type InsertNotifyInput struct {
	SysID    string `form:"sys_id" json:"sys_id" valid:"required"`
	LevelID  string `form:"level_id" json:"level_id" valid:"required"`
	Title    string `form:"title" json:"title" valid:"required"`
	Keywords string `form:"keywords" json:"keywords" valid:"required"`
	Content  string `form:"content" json:"content" valid:"required"`
}

type TpMsg struct {
	Openid  string
	Name    string
	Content string
	Time    string
}

type IDbNotify interface {
	Query(input *UserNotifyInput) (data db.QueryRows, count int, err error)
	Get(userID, sysID, pi, ps int64) (data db.QueryRows, count int, err error)
	AddSettings(input *SettingsInput) (err error)
	DeleteSettings(id, uid int64) (err error)
	Delete(id, uid int64) (err error)
	EditSettings(input *EditSettingsInput) (err error)
	Add(input *InsertNotifyInput) (err error)
	QueryToUserNotify() (data db.QueryRows, err error)
	ChangeStatus(id string) (err error)
}

type DbNotify struct {
	c component.IContainer
}

func NewDbNotify(c component.IContainer) *DbNotify {
	return &DbNotify{
		c: c,
	}
}

func (d *DbNotify) Query(input *UserNotifyInput) (data db.QueryRows, count int, err error) {
	db := d.c.GetRegularDB()
	c, q, a, err := db.Scalar(sql.QueryUserNotifyCount, map[string]interface{}{
		"title":   input.Title,
		"user_id": input.UserID,
		"sys_id":  input.SysID,
	})
	if err != nil {
		return nil, 0, fmt.Errorf("获取消息列表条数发生错误(err:%v),sql:(%s),输入参数:%v,", err, q, a)
	}
	data, q, a, err = db.Query(sql.QueryUserNotifyPageList, map[string]interface{}{
		"title":   input.Title,
		"user_id": input.UserID,
		"sys_id":  input.SysID,
		"pi":      input.Pi,
		"ps":      input.Ps,
	})
	if err != nil {
		return nil, 0, fmt.Errorf("获取消息列表发生错误(err:%v),sql:%s,输入参数:%v,", err, q, a)
	}
	return data, types.GetInt(c), nil
}

func (d *DbNotify) Get(userID, sysID, pi, ps int64) (data db.QueryRows, count int, err error) {
	db := d.c.GetRegularDB()
	c, q, a, err := db.Scalar(sql.QueryUserNotifySetCount, map[string]interface{}{
		"user_id": userID,
		"sys_id":  sysID,
	})
	if err != nil {
		return nil, 0, fmt.Errorf("获取消息设置列表条数发生错误(err:%v),sql:(%s),输入参数:%v,", err, q, a)
	}
	data, q, a, err = db.Query(sql.QueryUserNotifySetPageList, map[string]interface{}{
		"user_id": userID,
		"sys_id":  sysID,
		"pi":      pi,
		"ps":      ps,
	})
	if err != nil {
		return nil, 0, fmt.Errorf("获取消息设置列表发生错误(err:%v),sql:%s,输入参数:%v,", err, q, a)
	}
	return data, types.GetInt(c), nil
}

func (d *DbNotify) AddSettings(input *SettingsInput) (err error) {
	db := d.c.GetRegularDB()
	_, q, a, err := db.Execute(sql.AddNotifySettings, map[string]interface{}{
		"user_id":  input.UserID,
		"sys_id":   input.SysID,
		"keywords": input.Keywords,
		"level_id": input.Level,
	})
	if err != nil {
		return fmt.Errorf("添加添加消息配置数据发生错误(err:%v),sql:%s,输入参数:%v,", err, q, a)
	}
	return nil
}

func (d *DbNotify) DeleteSettings(id, uid int64) (err error) {
	db := d.c.GetRegularDB()
	_, q, a, err := db.Execute(sql.DelNotifySettings, map[string]interface{}{
		"id":  id,
		"uid": uid,
	})
	if err != nil {
		return fmt.Errorf("删除消息设置数据发生错误(err:%v),sql:%s,输入参数:%v,", err, q, a)
	}
	return nil
}

func (d *DbNotify) EditSettings(input *EditSettingsInput) (err error) {
	db := d.c.GetRegularDB()
	_, q, a, err := db.Execute(sql.EditNotifySettings, map[string]interface{}{
		"id":       input.ID,
		"keywords": input.Keywords,
		"level_id": input.Level,
		"status":   input.Status,
	})
	if err != nil {
		return fmt.Errorf("编辑消息设置数据发生错误(err:%v),sql:%s,输入参数:%v,", err, q, a)
	}
	return nil
}

func (d *DbNotify) Delete(id, uid int64) (err error) {
	db := d.c.GetRegularDB()
	_, q, a, err := db.Execute(sql.DelNotify, map[string]interface{}{
		"id":  id,
		"uid": uid,
	})
	if err != nil {
		return fmt.Errorf("删除消数据发生错误(err:%v),sql:%s,输入参数:%v,", err, q, a)
	}
	return nil
}

func (d *DbNotify) Add(input *InsertNotifyInput) (err error) {
	db := d.c.GetRegularDB()
	dbTrans, err := db.Begin()
	if err != nil {
		return fmt.Errorf("开启DB事务出错(err:%v)", err)
	}
	_, q, a, err := dbTrans.Execute(sql.InsertNotify, map[string]interface{}{
		"sys_id":   input.SysID,
		"level_id": input.LevelID,
		"title":    input.Title,
		"keywords": input.Keywords,
		"content":  input.Content,
	})
	if err != nil {
		dbTrans.Rollback()
		return fmt.Errorf("添加系统消息数据发生错误(err:%v),sql:%s,输入参数:%v,", err, q, a)
	}
	_, q, a, err = dbTrans.Execute(sql.InsertNotifyUser, map[string]interface{}{
		"sys_id":   input.SysID,
		"level_id": input.LevelID,
		"title":    input.Title,
		"keywords": input.Keywords,
		"content":  input.Content,
	})
	if err != nil {
		dbTrans.Rollback()
		return fmt.Errorf("添加用户消息数据发生错误(err:%v),sql:%s,输入参数:%v,", err, q, a)
	}
	dbTrans.Commit()
	return nil
}

//扫描消息
func (d *DbNotify) QueryToUserNotify() (data db.QueryRows, err error) {
	db := d.c.GetRegularDB()
	//扫描消息并修改状态
	guid := utility.GetGUID()
	_, q, a, err := db.Execute(sql.UpdateNotifyUser, map[string]interface{}{
		"guid": guid,
	})
	if err != nil {
		return nil, fmt.Errorf("修改消息数据发生错误(err:%v),sql:%s,输入参数:%v,", err, q, a)
	}

	data, q, a, err = db.Query(sql.QueryToUserNotify, map[string]interface{}{
		"guid": guid,
	})
	if len(data) <= 0 {
		return nil, context.NewError(context.ERR_NO_CONTENT, "没有可发送的消息")
	}
	if err != nil {
		return nil, fmt.Errorf("获取消息列表发生错误(err:%v),sql:%s,输入参数:%v,", err, q, a)
	}
	return data, nil
}

//修改消息状态
func (d *DbNotify) ChangeStatus(id string) (err error) {
	db := d.c.GetRegularDB()
	_, _, _, err = db.Execute(sql.SendNotifyUserSucc, map[string]interface{}{
		"id": id,
	})
	return err
}
