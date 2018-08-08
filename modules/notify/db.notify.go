package notify

import (
	"fmt"
	"github.com/micro-plat/hydra/component"
	"github.com/micro-plat/lib4go/db"
	"github.com/micro-plat/sso/modules/const/sql"
	"github.com/micro-plat/lib4go/types"
)

type UserNotifyInput struct {
	Title string `form:"title" json:"title"`
	UserID string `form:"user_id" json:"user_id" valid:"required"`
	SysID string `form:"sys_id" json:"sys_id" valid:"required"`
	Pi string `form:"pi" json:"pi" valid:"required"`
	Ps string `form:"ps" json:"ps" valid:"required"`
}

type SettingsInput struct {
	Keywords string `form:"keywords" json:"keywords" valid:"required"`
	Level string `form:"level_id" json:"level_id" valid:"required"`
	Status string `form:"status" json:"status" valid:"required"`
	UserID string `form:"user_id" json:"user_id" valid:"required"`
	SysID string `form:"sys_id" json:"sys_id" valid:"required"`
}

type EditSettingsInput struct {
	ID string `form:"id" json:"id" valid:"required"`
	Keywords string `form:"keywords" json:"keywords" valid:"required"`
	Level string `form:"level_id" json:"level_id" valid:"required"`
	Status string `form:"status" json:"status" valid:"required"`
}


type IDbNotify interface {
	Query(input *UserNotifyInput) (data db.QueryRows, count int, err error)
	Get(userID, sysID string, pi, ps int) (data db.QueryRows, count int, err error)
	Add(input *SettingsInput) (err error)
	DeleteSettingsByID(id string) (err error)
	DeleteNotifyByID(id string) (err error)
	Edit(input *EditSettingsInput) (err error)
}

type DbNotify struct {
	c component.IContainer
}

func NewDbNotify(c component.IContainer) *DbNotify {
	return &DbNotify{
		c: c,
	}
}

func(d *DbNotify) Query(input *UserNotifyInput) (data db.QueryRows, count int,err error) {
	db := d.c.GetRegularDB()
	c, q, a, err := db.Scalar(sql.QueryUserNotifyCount, map[string]interface{}{
		"title":   input.Title,
		"user_id": input.UserID,
		"sys_id": input.SysID,
	})
	if err != nil {
		return nil, 0,fmt.Errorf("获取消息列表条数发生错误(err:%v),sql:(%s),输入参数:%v,", err, q, a)
	}
	data, q, a, err = db.Query(sql.QueryUserNotifyPageList, map[string]interface{}{
		"title": input.Title,
		"user_id": input.UserID,
		"sys_id": input.SysID,
		"pi": input.Pi,
		"ps": input.Ps,
	})
	if err != nil {
		return nil, 0,fmt.Errorf("获取消息列表发生错误(err:%v),sql:%s,输入参数:%v,", err, q, a)
	}
	return data, types.ToInt(c),nil
}

func (d *DbNotify) Get(userID, sysID string, pi, ps int) (data db.QueryRows,count int,err error){
	db := d.c.GetRegularDB()
	c, q, a, err := db.Scalar(sql.QueryUserNotifySetCount, map[string]interface{}{
		"user_id": userID,
		"sys_id": sysID,
	})
	if err != nil {
		return nil, 0,fmt.Errorf("获取消息设置列表条数发生错误(err:%v),sql:(%s),输入参数:%v,", err, q, a)
	}
	data, q, a, err = db.Query(sql.QueryUserNotifySetPageList, map[string]interface{}{
		"user_id": userID,
		"sys_id": sysID,
		"pi": pi,
		"ps": ps,
	})
	if err != nil {
		return nil, 0,fmt.Errorf("获取消息设置列表发生错误(err:%v),sql:%s,输入参数:%v,", err, q, a)
	}
	return data, types.ToInt(c),nil
}

func (d *DbNotify) Add(input *SettingsInput) (err error) {
	db := d.c.GetRegularDB()
	_, q, a, err := db.Execute(sql.AddNotifySettings, map[string]interface{}{
		"user_id":      input.UserID,
		"sys_id":     	input.SysID,
		"keywords": 	input.Keywords,
		"level_id":     input.Level,
		"status":    	input.Status,
	})
	if err != nil {
		return fmt.Errorf("添加系统管理数据发生错误(err:%v),sql:%s,输入参数:%v,", err, q, a)
	}
	return nil
}

func (d *DbNotify) DeleteSettingsByID(id string) (err error) {
	db := d.c.GetRegularDB()
	_, q, a, err := db.Execute(sql.DelNotifySettings, map[string]interface{}{
		"id": id,
	})
	if err != nil {
		return fmt.Errorf("删除消息设置数据发生错误(err:%v),sql:%s,输入参数:%v,", err, q, a)
	}
	return nil
}

func (d *DbNotify) Edit(input *EditSettingsInput) (err error) {
	db := d.c.GetRegularDB()
	_, q, a, err := db.Execute(sql.EditNotifySettings, map[string]interface{}{
		"id": 			input.ID,
		"keywords": 	input.Keywords,
		"level_id":     input.Level,
		"status":    	input.Status,
	})
	if err != nil {
		return fmt.Errorf("编辑消息设置数据发生错误(err:%v),sql:%s,输入参数:%v,", err, q, a)
	}
	return nil
}

func (d *DbNotify) DeleteNotifyByID(id string) (err error) {
	db := d.c.GetRegularDB()
	_, q, a, err := db.Execute(sql.DelNotify, map[string]interface{}{
		"id": id,
	})
	if err != nil {
		return fmt.Errorf("删除消息设置数据发生错误(err:%v),sql:%s,输入参数:%v,", err, q, a)
	}
	return nil
}