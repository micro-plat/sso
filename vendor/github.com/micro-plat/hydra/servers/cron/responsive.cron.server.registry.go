package cron

import (
	"time"

	"github.com/micro-plat/hydra/servers"
	"github.com/micro-plat/hydra/servers/pkg/sharding"
	"github.com/micro-plat/lib4go/types"
)

func (s *CronResponsiveServer) watchMasterChange(root, path string) error {
	cldrs, _, err := s.engine.GetRegistry().GetChildren(root)
	if err != nil {
		return err
	}
	s.master = s.isMaster(path, cldrs)
	servers.Tracef(s.Infof, "%s", types.DecodeString(s.master, true, "master cron server", "slave cron server"))
	if err = s.notifyConsumer(s.master); err != nil {
		return err
	}
	children, err := s.engine.GetRegistry().WatchChildren(root)
	if err != nil {
		return err
	}
	go func() {
		for {
			select {
			case <-s.closeChan:
				return
			case cldWatcher := <-children:
				if cldWatcher.GetError() != nil {
					break
				}
				cldrs, _ = cldWatcher.GetValue()
				master := s.isMaster(path, cldrs)
				if master != s.master {
					servers.Tracef(s.Infof, "%s", types.DecodeString(master, true, "master cron server", "slave cron server"))
					s.notifyConsumer(master)
					s.master = master
				}

			LOOP:
				children, err = s.engine.GetRegistry().WatchChildren(root)
				if err != nil {
					servers.Tracef(s.Errorf, "监控服务节点发生错误:err:%v", err)
					if s.done {
						return
					}
					time.Sleep(time.Second)
					goto LOOP
				}
			}
		}
	}()
	return nil
}

func (s *CronResponsiveServer) isMaster(path string, cldrs []string)(isMaster bool) {
	s.shardingIndex, isMaster = sharding.IsMaster(s.master, s.shardingCount, path, cldrs)
	return isMaster

}
func (s *CronResponsiveServer) notifyConsumer(v bool) error {
	if v {
		return s.server.Resume()
	}
	s.server.Pause()
	return nil
}
func getSharding(index int, count int) int {
	if count <= 0 && index >= 0 {
		return index
	}
	if index < 0 || index >= count {
		return -1
	}
	return index % count
}
