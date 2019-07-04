package notify

import (
	"github.com/micro-plat/hydra/component"
	"github.com/micro-plat/lib4go/db"
)

type INotify interface {
	Query(input *UserNotifyInput) (data db.QueryRows, count int, err error)
	Get(uid, sysID, pi, ps int64) (data db.QueryRows, count int, err error)
	AddSettings(input *SettingsInput) (err error)
	DeleteSettings(id, uid int64) (err error)
	Delete(id, uid int64) (err error)
	EditSettings(input *EditSettingsInput) (err error)
	Add(input *InsertNotifyInput) (err error)
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
func (n *Notify) Query(input *UserNotifyInput) (data db.QueryRows, count int, err error) {
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
func (n *Notify) Get(userID, sysID, pi, ps int64) (data db.QueryRows, count int, err error) {
	data, count, err = n.cache.QueryNotifySet(userID, sysID, pi, ps)
	if err != nil || data == nil {
		data, count, err = n.db.Get(userID, sysID, pi, ps)
		if err != nil || data == nil {
			return nil, 0, err
		}
		if err = n.cache.SaveNotifySet(userID, sysID, pi, ps, data, count); err != nil {
			return nil, 0, err
		}
	}
	return
}

//添加配置
func (n *Notify) AddSettings(input *SettingsInput) (err error) {
	err = n.db.AddSettings(input)
	if err != nil {
		return err
	}
	return n.cache.FreshNotifySet()
}

//删除配置
func (n *Notify) DeleteSettings(id, uid int64) (err error) {
	if err = n.db.DeleteSettings(id, uid); err != nil {
		return err
	}
	return n.cache.FreshNotifySet()
}

//删除消息
func (n *Notify) Delete(id, uid int64) (err error) {
	if err = n.db.Delete(id, uid); err != nil {
		return err
	}
	return n.cache.FreshNotify()
}

//编辑配置
func (n *Notify) EditSettings(input *EditSettingsInput) (err error) {
	if err = n.db.EditSettings(input); err != nil {
		return err
	}
	return n.cache.FreshNotifySet()
}

//接收系统消息
func (n *Notify) Add(input *InsertNotifyInput) (err error) {
	if err = n.db.Add(input); err != nil {
		return err
	}
	//刷新用户消息缓存列表
	return n.cache.FreshNotify()
}

//发送消息给用户
// func (n *Notify) SendMsg() (err error) {
// 	data, err := n.db.QueryToUserNotify()
// 	if err != nil {
// 		return err
// 	}
// 	//循环发送消息
// 	for _, v := range data {
// 		//使用微信发送模板消息，发送失败则进入下轮继续发送
// 		err = n.wxMsg.Send(&TpMsg{
// 			Openid:  v.GetString("wx_openid"),
// 			Name:    v.GetString("name"),
// 			Content: v.GetString("content"),
// 			Time:    v.GetString("create_times"),
// 		})
// 		//发送成功，修改状态
// 		if err == nil {
// 			err = n.db.ChangeStatus(v.GetString("id"))
// 			fmt.Println(err)
// 		}
// 	}
// 	return n.cache.FreshNotify()
// }
