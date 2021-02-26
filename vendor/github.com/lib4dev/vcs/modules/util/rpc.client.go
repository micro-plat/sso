package util

import (
	"fmt"
	"net/http"

	"github.com/lib4dev/vcs/modules/const/conf"
	"github.com/micro-plat/hydra"
	"github.com/micro-plat/lib4go/types"
)

//RpcRequest RpcRequest
func RpcRequest(url string, info types.XMap) (content types.XMap, err error) {
	client := hydra.C.RPC().GetRegularRPC(conf.RemoteName)

	resp, err := client.Request(url, info)
	if err != nil {
		err = fmt.Errorf("sms.RpcRequest:%s;error:%v", info, err)
		return
	}
	status := resp.GetStatus()
	if status != http.StatusOK {
		err = fmt.Errorf("sms.RpcRequest,status:%d;error:%v", status, err)
	}
	content = resp.GetMap()
	return
}
