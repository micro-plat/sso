package notify

import (
	"github.com/micro-plat/hydra/component"
	"github.com/micro-plat/lib4go/db"
)

type INotify interface {
	Query(input *UserNotifyInput) (data db.QueryRows,count int,err error)
	Get(userID,sysID string,pi,ps int) (data db.QueryRows,count int,err error)
	Add(input *SettingsInput) (err error)
	DeleteSettingsByID(id string) (err error)
	DeleteNotifyByID(id string) (err error)
	Edit(input *EditSettingsInput) (err error)
	InsertNotify(input *InsertNotifyInput) (err error)
	SendMsg() (err error)
}

type Notify struct {
	c     component.IContainer
	cache ICacheNotify
	db    IDbNotify
}

func NewNotify(c component.IContainer) *Notify {
	return &Notify{
		c:     c,
		cache: NewCacheNotify(c),
		db:    NewDbNotify(c),
	}
}
//获取消息列表
func (n *Notify) Query(input *UserNotifyInput) (data db.QueryRows,count int,err error) {
	//从缓存获取数据
	data, count, err = n.cache.QueryNotify(input.Title, input.UserID, input.SysID, input.Pi, input.Ps)
	if err != nil || data == nil {
		//从数据库取数据
		data, count, err = n.db.Query(input)
		if err != nil {
			return nil, 0, err
		}
		if err = n.cache.SaveNotify(input.Title, input.UserID, input.SysID, input.Pi, input.Ps, data, count); err != nil {
			return nil, 0, err
		}
	}
	
	return
}

//获取消息设置
func(n *Notify) Get(userID, sysID string, pi, ps int) (data db.QueryRows, count int, err error){
	data, count, err = n.cache.QueryNotifySet(userID, sysID, pi, ps)
	if err != nil || data == nil {
		data, count, err = n.db.Get(userID, sysID, pi, ps)
		if err != nil || data == nil{
			return nil, 0, err
		}
		if err = n.cache.SaveNotifySet(userID, sysID, pi, ps, data, count); err != nil {
			return nil, 0, err
		}
	}
	return 
}
//添加配置
func (n *Notify) Add(input *SettingsInput) (err error){
	err = n.db.Add(input)
	if err != nil {
		return err
	}
	return n.cache.FreshNotifySet()
}
//删除配置
func (n *Notify) DeleteSettingsByID(id string) (err error) {
	if err = n.db.DeleteSettingsByID(id); err != nil {
		return err
	} 
	return n.cache.FreshNotifySet();
}
//删除消息
func (n *Notify) DeleteNotifyByID(id string) (err error){
	if err = n.db.DeleteNotifyByID(id); err != nil {
		return err
	} 
	return n.cache.FreshNotify();
}
//编辑配置
func (n *Notify) Edit(input *EditSettingsInput) (err error) {
	if err = n.db.Edit(input); err != nil {
		return err
	}
	return n.cache.FreshNotifySet();
}
//接收系统消息
func (n *Notify) InsertNotify(input *InsertNotifyInput) (err error) {
	if err = n.db.InsertNotify(input); err != nil {
		return err
	}
	//刷新用户消息缓存列表
	return n.cache.FreshNotify()
}
//发送消息给用户
func (n *Notify) SendMsg() (err  error){
	if err = n.db.SendMsg(); err != nil {
		return err
	}
	//发送数据结束，没有错误则刷新用户消息缓存列表
	return n.cache.FreshNotify()
}