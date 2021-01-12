## sso sdk使用说明

```
author:liujinyin
date: 2021-1-6
desction: sso 集成使用说明
```

SSO系统集成分"前端项目"集成和"后端服务"接口集成以及部分配置修改.

### 1. 前端项目集成

#### 1.1 npm包引入

1. 在项目的package.json的dependencies节点添加 

```javascript
    "nav-menu": "^1.3.49",
    "qxnw-enum": "^1.0.14",
    "qxnw-sso": "^1.0.20",

//注：版本可能存在更新，以最新版本为准
```

#### 1.2 index.html更新

1. index.html(vue版本不同,文件在跟目录或public目录)，在index.html的head节点内中添加如下代码：

``` jsavascript

 <% if(process.env.NODE_ENV=="production") { %>
      <script type="text/javascript" src="/config/vue?t=<%= new Date().getTime() %>" ></script>
 <% } else{  %>
      <script type="text/javascript" >
          window.globalConfig= <%= JSON.stringify( process.env)%>
          for(var k in window.globalConfig){
            if(k.indexOf("VUE_APP_")>=0){
              var nk = k.replace("VUE_APP_","");
              window.globalConfig[nk] = window.globalConfig[k];
              delete window.globalConfig[k]
            }
          }

      </script>
<% }  %>
```

#### 1.3 env.dev 配置

1. 目前有的配置项

```jsavascript
NODE_ENV：标识当前的执行环境，可选值：development,production
VUE_APP_apiURL: 标识后端服务地址:http://api.100bm0.com
VUE_APP_loginWebHost：标识登陆的SSO跳转的地址，如：http://ssov4.100bm0.com:6687
VUE_APP_ident：标识系统编号，与sso中的sso_system_info.ident相同

```

2. 查看配置项

```
可在浏览器中输入window.globalConfig 查看所有配置内容


```


#### 1.4 env.prod 配置

1. 当前prod 的配置都已由后端注册中心获取，统一由接口/config/vue提供，前端配置不会生效

```jsavascript

NODE_ENV：标识当前的执行环境，可选值：development,production

/*
注： 
所有服务端配置不能配置有“VUE_APP_”这样的前缀.
配置内容在/platname/systemname/serverType/cluster/conf/vueconf
如：/sso_v4/mgrserver/web/prod/conf/vueconf
*/
```

2. 查看配置项

```go
可在浏览器中输入window.globalConfig 查看所有配置内容

```

#### 1.5 src/main.js 引入 qxnw-sso 


```javascript

import {ssoHttpConfig} from 'qxnw-sso';
var conf = window.globalConfig
var ssocfg =  ssoHttpConfig(conf.apiURL ||"", "localStorage", conf.loginWebHost, conf.ident);

Vue.prototype.$sso = ssocfg.sso;
Vue.prototype.$http = ssocfg.http;

/*
注：
1. 去掉原有的引入“ utility/http.js”，并删除该文件，避免引起误解
2. 如原先有使用到 Vue.prototype.$get,Vue.prototype.$post,Vue.prototype.$fetch 等方式。请添加如下代码

Vue.prototype.$get=ssocfg.http.get;
Vue.prototype.$post=ssocfg.http.post;
Vue.prototype.$fetch=ssocfg.http.fetch;
Vue.prototype.$del=ssocfg.http.del;
Vue.prototype.$patch=ssocfg.http.patch;

*/

```

#### 1.5 ssocallback路由处理 

创建 sso.callback.vue 文件，复制添加如下内容。添加前段路由： /ssocallback 到路由表
```javascript

<template>
</template>

<script>
  export default {
    data () {
      return {
      }
    },
    mounted(){
      this.validSsoLogin();
    },
    methods:{
      validSsoLogin(){
          this.$http.post("/sso/login/verify",{code: this.$route.query.code})
            .then(res =>{
                this.$sso.changeRouteAfterLogin(this.$router, res.user_name, res.role_name);
            }).catch(err => {
             if (err.response) {
                if (err.response.status == 406) {
                  this.$sso.errPage(0)
                }
              }
              console.log(err);
            });
      }
    }
  }
</script>

```

### 2. 后端项目集成

1. 引入sdk包
```go

import "github.com/micro-plat/sso/sso"
 
```

2. 添加 OnHandleExecuting 钩子函数处理

```go 

	App.OnHandleExecuting(func(ctx hydra.IContext) (rt interface{}) {
 		//验证jwt并缓存登录用户信息
		if err := sso.CheckAndSetMember(ctx); err != nil {
			return err
		}
		return nil
	})
	
```

3. 初始化SDK的接口数据(地址，系统标识，系统密钥)

```go

App.OnStarting(func(appConf app.IAPPConf) error {
    //检查配置信息
    var appcfg model.AppConfig
    if _, err := appConf.GetServerConf().GetSubObject("app", &appcfg);err != nil {
        return fmt.Errorf("获取appconf配置失败,err:%v", err)
    }
        
    model.SaveAppConfig(&appcfg);
    //初始化sso必须数据
    if err := sso.Config(appcfg.SsoApiHost, appcfg.Ident,appcfg.SsoSecret); err != nil {
        return fmt.Errorf("sso-配置失败,err:%v", err)
    }
    return nil
}
/*
SsoApiHost:
线下:http://ssov4.100bm0.com:6689
*/

```

4. jwt忽略配置处理

```go
//jwt 的配置忽略中增加 "/sso/login/verify"

hydra.OnReady(func() error {
    hydra.Conf.Web("8181"). //端口根据业务自定定义
    Jwt(jwt.WithHeader(),
        jwt.WithExcludes("/sso/login/verify"),
    )
})
/*
注：
如果WithExcludes包含其他的地址，将 “/sso/login/verify” 放在其中即可
*/
```

5. vueConf 配置（可选）

```go

hydra.OnReady(func() error {

    hydra.Conf.Web("8181").
		Sub("webconf", model.VueConfig{
			Ident:        "sas",
			LoginWebHost: "//ssov4.100bm0.com:6687",
        })
})



/*
注：
1. model.VueConfig 是自己定义的Struct 
2. Ident,LoginWebHost必填字段

*/

/*
VueConfig供参考，可修改

type VueConfig struct {
	APIURL string `json:"apiURL"`
	//Ident .
	Ident string `json:"ident" valid:"required"`
	//登录系统地址
	LoginWebHost  string `json:"loginWebHost" valid:"required"`
}

*/

```


6. 配置读取服务：/config/vue 

```go
package config

import (
	"encoding/json"
	"fmt"

	"github.com/micro-plat/hydra"
)

//VueHandler VueConfig
func VueHandler(ctx hydra.IContext) interface{} {
	configData := map[string]interface{}{}
	_, err := ctx.APPConf().GetServerConf().GetSubObject("webconf", &configData)
	if err != nil {
		return fmt.Errorf("GetSubObject:vueconf:%v", err)
	}
	ctx.Response().ContentType("text/plain")
	bytes, _ := json.Marshal(configData)
	return fmt.Sprintf("window.globalConfig=%s", string(bytes))
}

//路由注册
App.Micro("/config/vue", config.VueHandler)

```

### 3. sdk接口说明
 
#### 3.1  根据⽤⼾名获取⽤⼾信息
``` go
GetUserInfoByName(userName string)
```
输⼊参数

|参数 |类型|说明|
| -------------|:--------------:|:--------------:|
|userName| string| ⽤⼾名称|

输出

|参数 |类型|说明|
| -------------|:--------------:|:--------------:|
|userName |string |⽤⼾名称|
|WxOpID |string |微信openID|
|ExtParams |string| ⽤⼾扩展参数(⼀个json对象)|
|UserID |string| ⽤⼾标识|

#### 3.2 获取⽤⼾在某个⼦系统下的菜单数据
```go
GetUserMenu(userID int)
```
输⼊参数

|参数 |类型|说明|
| -------------|:--------------:|:--------------:|
|userID| int |⽤⼾标识|

输出

|参数 |类型|说明|
| -------------|:--------------:|:--------------:|
|ID |string |菜单标识
|Name |string |菜单名称
|Level |string |级次
|IsOpen |string |是否展开
|Icon |string |图标
|SystemID |string |系统标识
|Parent |string |⽗级编号
|Path |string |路由地址
|Sortrank |string |排序编号
|Children |对象数组 |⼦菜单

#### 3.3获取⼦系统信息
``` go
GetSystemInfo()
```
输⼊参数(无)

输出

|参数 |类型|说明|
| -------------|:--------------:|:--------------:|
|ID |string |系统标识
|Ident |string |系统ident(英⽂名称)
|Name |string |系统名称
|Theme |string |主题样式
|Layout |string |⻚⾯布局样式
|Logo |string |系统图标地址


#### 3.4 获取当前用户可访问的其他子系统
``` go
GetUserOtherSystems()
```
输⼊参数(无)

输出

|参数 |类型|说明|
| -------------|:--------------:|:--------------:|
|ID |string |系统标识
|Ident |string |系统ident(英⽂名称)
|Name |string |系统名称
|IndexUrl |string |⼦系统地址 host
