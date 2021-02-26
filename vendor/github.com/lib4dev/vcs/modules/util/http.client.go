package util

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/lib4dev/vcs/modules/const/conf"
	"github.com/micro-plat/hydra"
	"github.com/micro-plat/lib4go/types"
)

//HttpRequest HttpRequest
func HttpRequest(url string, info types.XMap) (content types.XMap, err error) {
	client := hydra.C.HTTP().GetRegularClient(conf.RemoteName)
	bytes, _ := json.Marshal(info)
	dataInfo := string(bytes)
	hdr := http.Header{}
	hdr["Content-Type"] = []string{"application/json"}
	contentVal, status, err := client.Request("POST", url, dataInfo, "utf-8", hdr)
	if err != nil {
		err = fmt.Errorf("sms.HttpRequest:%s;error:%v", dataInfo, err)
		return
	}
	if status != http.StatusOK {
		err = fmt.Errorf("sms.HttpRequest,status:%d;error:%v", status, err)
		return
	}
	err = json.Unmarshal(contentVal, &content)
	if err != nil {
		err = fmt.Errorf("sms.HttpRequest,Unmarshal:%s, error;%v", contentVal, err)
	}
	return
}
