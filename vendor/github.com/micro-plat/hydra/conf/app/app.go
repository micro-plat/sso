package app

import (
	"fmt"

	"github.com/micro-plat/hydra/conf"
	"github.com/micro-plat/hydra/conf/server"
	"github.com/micro-plat/hydra/conf/server/acl/blacklist"
	"github.com/micro-plat/hydra/conf/server/acl/limiter"
	"github.com/micro-plat/hydra/conf/server/acl/proxy"
	"github.com/micro-plat/hydra/conf/server/acl/whitelist"
	"github.com/micro-plat/hydra/conf/server/apm"
	"github.com/micro-plat/hydra/conf/server/auth/apikey"
	"github.com/micro-plat/hydra/conf/server/auth/basic"
	"github.com/micro-plat/hydra/conf/server/auth/jwt"
	"github.com/micro-plat/hydra/conf/server/auth/ras"
	"github.com/micro-plat/hydra/conf/server/header"
	"github.com/micro-plat/hydra/conf/server/metric"
	"github.com/micro-plat/hydra/conf/server/mqc"
	"github.com/micro-plat/hydra/conf/server/processor"
	"github.com/micro-plat/hydra/conf/server/queue"
	"github.com/micro-plat/hydra/conf/server/render"
	"github.com/micro-plat/hydra/conf/server/static"
	"github.com/micro-plat/hydra/conf/server/task"
	"github.com/micro-plat/hydra/conf/vars"
	"github.com/micro-plat/hydra/conf/vars/rlog"
	"github.com/micro-plat/hydra/global"
	"github.com/micro-plat/hydra/registry"
)

//IAPPConf 服务器配置信息
type IAPPConf interface {
	GetServerConf() conf.IServerConf
	GetVarConf() conf.IVarConf

	GetMQCMainConf() (*mqc.Server, error)
	GetMQCQueueConf() (*queue.Queues, error)

	GetCRONTaskConf() (*task.Tasks, error)

	GetJWTConf() (*jwt.JWTAuth, error)
	GetHeaderConf() (header.Headers, error)
	GetMetricConf() (*metric.Metric, error)
	GetStaticConf() (*static.Static, error)
	GetAPIKeyConf() (*apikey.APIKeyAuth, error)
	GetRASConf() (*ras.RASAuth, error)
	GetBasicConf() (*basic.BasicAuth, error)
	GetRenderConf() (*render.Render, error)
	GetWhiteListConf() (*whitelist.WhiteList, error)
	GetBlackListConf() (*blacklist.BlackList, error)
	GetLimiterConf() (*limiter.Limiter, error)
	GetProxyConf() (*proxy.Proxy, error)
	GetAPMConf() (*apm.APM, error)
	GetProcessorConf() (*processor.Processor, error)

	//获取远程日志配置
	GetRLogConf() (*rlog.Layout, error)

	Close() error
}

var _ IAPPConf = &APPConf{}

//APPConf 应用配置信息
type APPConf struct {
	serverConf conf.IServerConf
	varConf    conf.IVarConf
	*server.HttpSub
	*server.CronSub
	*server.MQCSub
	*vars.VarSub
}

//NewAPPConfBy 构建服务器配置缓存
func NewAPPConfBy(platName, sysName, serverType, clusterName string, rgst registry.IRegistry) (s *APPConf, err error) {
	s = &APPConf{}

	//构建server配置
	s.serverConf, err = server.NewServerConf(platName, sysName, serverType, clusterName, rgst)
	if err != nil {
		return nil, err
	}

	//构建var配置
	s.varConf, err = vars.NewVarConf(platName, rgst)
	if err != nil {
		return nil, err
	}

	//构建server的组件配置(todo:移到server配置内)
	s.HttpSub = server.NewHttpSub(s.serverConf)
	s.CronSub = server.NewCronSub(s.serverConf)
	s.MQCSub = server.NewMQCSub(s.serverConf)
	s.VarSub = vars.NewVarSub(s.varConf)
	return s, nil

}

// NewAPPConf 构建服务器配置
func NewAPPConf(mainConfpath string, rgst registry.IRegistry) (s *APPConf, err error) {

	//处理平台名、系统名包含多段问题
	//获取服务器类型
	list := registry.Split(registry.Trim(mainConfpath))
	tp := list[len(list)-3]

	//无法准确获得平台、系统名，只能通过当前应用配置获得，再比较
	pub := server.NewServerPub(global.Def.PlatName, global.Def.SysName, tp, global.Def.ClusterName)
	if pub.GetServerPath() != mainConfpath {
		return nil, fmt.Errorf("非当前平台、系统、集群的服务不支持获取APPConf")
	}
	return NewAPPConfBy(global.Def.PlatName, global.Def.SysName, tp, global.Def.ClusterName, rgst)

}

//Close 关闭清理资源
func (s *APPConf) Close() error {
	return s.serverConf.Close()
}

//GetServerConf 获取server配置
func (s *APPConf) GetServerConf() conf.IServerConf {
	return s.serverConf
}

//GetVarConf 获取var配置
func (s *APPConf) GetVarConf() conf.IVarConf {
	return s.varConf
}
