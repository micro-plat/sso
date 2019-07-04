package rpc

import (
	"fmt"
	"path"
	"path/filepath"
	"strings"
	"time"

	"github.com/micro-plat/lib4go/jsons"
)

//publish 将当前服务器的节点信息发布到注册中心
func (w *RpcResponsiveServer) publish() (err error) {

	addr := w.server.GetAddress()
	ipPort := strings.Split(addr, "://")[1]
	pubPath := filepath.Join(w.currentConf.GetServerPubRootPath(), ipPort)
	data := map[string]string{
		"service": addr,
	}
	jsonData, _ := jsons.Marshal(data)
	nodeData := string(jsonData)
	err = w.engine.GetRegistry().CreateTempNode(pubPath, nodeData)
	if err != nil {
		err = fmt.Errorf("服务发布失败:(%s)[%v]", pubPath, err)
		return
	}
	w.pubs = []string{pubPath}

	names := w.currentConf.GetStrings("host")
	if len(names) == 0 {
		names = append(names, w.currentConf.GetSysName())
	}
	srvs := w.GetServices()
	for _, host := range names {
		for _, srv := range srvs {
			servicePath := path.Join(w.currentConf.GetServicePubRootPath(filepath.Join(host, srv)), ipPort)
			err := w.engine.GetRegistry().CreateTempNode(servicePath, nodeData)
			if err != nil {
				err = fmt.Errorf("服务发布失败:(%s)[%v]", servicePath, err)
				return err
			}
			w.pubs = append(w.pubs, servicePath)
		}

	}
	go w.publishCheck(nodeData)
	return
}

//publishCheck 定时检查节点数据是否存在
func (w *RpcResponsiveServer) publishCheck(data string) {
LOOP:
	for {
		select {
		case <-w.closeChan:
			break LOOP
		case <-time.After(time.Second * 10):
			if w.done {
				break LOOP
			}
			w.checkPubPath(data)
		}
	}
}

//checkPubPath 检查已发布的节点，不存在则创建
func (w *RpcResponsiveServer) checkPubPath(data string) {
	w.pubLock.Lock()
	defer w.pubLock.Unlock()
	for _, path := range w.pubs {
		if w.done {
			break
		}
		ok, err := w.engine.GetRegistry().Exists(path)
		if err != nil {
			break
		}
		if !ok {
			err := w.engine.GetRegistry().CreateTempNode(path, data)
			if err != nil {
				break
			}
			w.Logger.Infof("节点(%s)已恢复", path)
		}
	}
}

//unpublish 删除已发布的节点
func (w *RpcResponsiveServer) unpublish() {
	w.pubLock.Lock()
	defer w.pubLock.Unlock()
	for _, path := range w.pubs {
		w.engine.GetRegistry().Delete(path)
	}
	w.pubs = make([]string, 0, 0)
}
