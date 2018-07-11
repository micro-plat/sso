### 回调请求的一般处理逻辑（一个回调地址处理一个公众号的消息和事件）
```Go
package main

import (
	"log"
	"net/http"

	"github.com/micro-plat/wechat/mp"
	"github.com/micro-plat/wechat/mp/menu"
	"github.com/micro-plat/wechat/mp/message/callback/request"
	"github.com/micro-plat/wechat/mp/message/callback/response"
)

const (
	wxAppId     = "appid"
	wxAppSecret = "appsecret"

	wxOriId         = "oriid"
	wxToken         = "token"
	wxEncodedAESKey = "aeskey"
)

var (
	// 下面两个变量不一定非要作为全局变量, 根据自己的场景来选择.
	msgHandler mp.Handler
	msgServer  *mp.Server
)

func init() {
	mux := mp.NewServeMux()
	mux.DefaultMsgHandleFunc(defaultMsgHandler)
	mux.DefaultEventHandleFunc(defaultEventHandler)
	mux.MsgHandleFunc(request.MsgTypeText, textMsgHandler)
	mux.EventHandleFunc(menu.EventTypeClick, menuClickEventHandler)

	msgHandler = mux
	msgServer = mp.NewServer(wxOriId, wxAppId, wxToken, wxEncodedAESKey, msgHandler, nil)
}

func textMsgHandler(ctx *mp.Context) {
	log.Printf("收到文本消息:\n%s\n", ctx.MsgPlaintext)

	msg := request.GetText(ctx.MixedMsg)
	resp := response.NewText(msg.FromUserName, msg.ToUserName, msg.CreateTime, msg.Content)
	//ctx.RawResponse(resp) // 明文回复
	ctx.AESResponse(resp, 0, "", nil) // aes密文回复
}

func defaultMsgHandler(ctx *mp.Context) {
	log.Printf("收到消息:\n%s\n", ctx.MsgPlaintext)
	ctx.NoneResponse()
}

func menuClickEventHandler(ctx *mp.Context) {
	log.Printf("收到菜单 click 事件:\n%s\n", ctx.MsgPlaintext)

	event := menu.GetClickEvent(ctx.MixedMsg)
	resp := response.NewText(event.FromUserName, event.ToUserName, event.CreateTime, "收到 click 类型的事件")
	//ctx.RawResponse(resp) // 明文回复
	ctx.AESResponse(resp, 0, "", nil) // aes密文回复
}

func defaultEventHandler(ctx *mp.Context) {
	log.Printf("收到事件:\n%s\n", ctx.MsgPlaintext)
	ctx.NoneResponse()
}

func init() {
	http.HandleFunc("/wx_callback", wxCallbackHandler)
}

// wxCallbackHandler 是处理回调请求的 http handler.
//  1. 不同的 web 框架有不同的实现
//  2. 一般一个 handler 处理一个公众号的回调请求(当然也可以处理多个, 这里我只处理一个)
func wxCallbackHandler(w http.ResponseWriter, r *http.Request) {
	msgServer.ServeHTTP(w, r, nil)
}

func main() {
	log.Println(http.ListenAndServe(":80", nil))
}
```

### 公众号api调用的一般处理逻辑
```Go
package main

import (
	"fmt"

	"github.com/micro-plat/wechat/mp/base"
	"github.com/micro-plat/wechat/mp"
)

const (
	wxAppId     = "appid"
	wxAppSecret = "appsecret"

	wxOriId         = "oriid"
	wxToken         = "token"
	wxEncodedAESKey = "aeskey"
)

var (
	accessTokenServer mp.AccessTokenServer = mp.NewDefaultAccessTokenServer(wxAppId, wxAppSecret, nil)
	wechatClient      *mp.Context           = mp.NewClient(accessTokenServer, nil)
)

func main() {
	fmt.Println(base.GetCallbackIP(wechatClient))
}
```