## sso sdk使用说明

```
author:liujinyin
date: 2021-1-6
desction: sso 集成使用说明
```

SSO系统集成分"前端项目"集成和"后端服务"接口集成以及部分配置修改.

### 1. 前端项目集成

#### 1.1 公共js包引入

1. 拷贝 gitlab.100bm.cn/micro-plat/jspkg/utility 中的js放在项目的 src/utility

```javascript
//注：
//   1.此内容会通过工具直接生成到里面
//   2.npm包方式

```

2. 拷贝 gitlab.100bm.cn/micro-plat/jspkg/utility/env.conf.json 中的js放在项目的 根目录/public （vue2.0 放在 根目录/static)

```javascript
//注：此内容会通过工具直接生成到里面

// 增加文件内配置项
{
    "ident":"sso",    // 项目的IDENT
    "apiURL":"http://ssov4.100bm0.com:6677",       //服务请求地址，为空为与项目相同地址
    "loginWebHost":"http://ssov4.100bm0.com:6687", //sso 登录系统地址
    "cookieName":"Authorization",                  //如果使用cookie保存token时填写
    "companyRight":"四川千行你我科技股份有限公司",      //根据实际业务填写
    "companyRightCode":"蜀ICP备20003360号"          //根据实际业务填写
}

```
#### 1.2 修改main.js 

```javascript
//1. 添加引用
import utility from './services'
Vue.use(utility, "../static/env.conf.json");

//2. 增加枚举的回调函数
Vue.prototype.$enum.callback(async function(type){
  var url = "/dds/dictionary/get";
  var data = await Vue.prototype.$http.get(url, { dic_type: type });
  console.log("dictionary.data:", type, data);
  return data;
});

//3. 设置默认的http服务请求地址
Vue.prototype.$http.setBaseURL(Vue.prototype.$env.Conf.apiURL);

//4. 增加对403 状态码的拦截处理
Vue.prototype.$http.addStatusCodeHandle(res => {
  var url = (res.headers || {}).location ||""; 
  if(!url){
    url = this.$env.Conf.loginWebHost + "/sso/jump?returnurl=";
  }

  url =url + encodeURIComponent(document.URL);
  console.log("redirect:url", url);
  window.location = url ;
  //return new Error("请补充注册中心auth/jwt的AuthURL配置");
}, 403);


```

#### 1.3  src/App.vue 文件调整（存在需读取服务端配置才调整）

1.  直接复制下面内容到App.vue文件

```vue

<template>
  <div id="vapp">
    <router-view v-if="hasLoaded"/>
  </div>
</template>

<script>
export default {
  name: 'App',
  data(){
    return{
      hasLoaded: false
    }
  },
  created(){ 
    this.getWebconfig()
  },
  methods:{
    async getWebconfig(){
      var that = this;
      await this.$env.load(async function(){
        var data = await that.$http.get("/system/webconfig");
        that.hasLoaded = true
        return data;
      });
    }
  }
}
</script>


```
2. 后端提供配置读取服务：/system/webconfig

```go
	App.Micro("/system/webconfig", system.WebConfigHandler)    

//WebConfigHandler WebConfigHandler
func WebConfigHandler(ctx hydra.IContext) interface{} {
  //读取服务端配置
  configData := map[string]interface{}{}
	 
	return configData
}

```


#### 1.4 Menu.vue 菜单组件引入

1. 直接将下面内容复制到项目内(src/pages/member/menu.vue)

```vue

<template>
  <div id="app">
    <nav-menu
      :menus="menus"
      :copyright="copyright"
      :copyrightcode="copyrightcode"
      :themes="themes"
      :logo="logo"
      :systemName="systemName"
      :userinfo="userinfo"
      :items="items"
      :pwd="pwd"
      :signOut="signOutM"
      ref="NewTap"
    >
    </nav-menu>
  </div>
</template>

<script>
  import navMenu from 'nav-menu'; // 引入
  export default {
    name: 'app',
    data () {
      return {
        logo: "",
        copyright: (this.$env.Conf.companyRight||"") + "Copyright©" + new Date().getFullYear() +"版权所有",
        copyrightcode: this.$env.Conf.companyRightCode ,
        themes: "", //顶部左侧背景颜色,顶部右侧背景颜色,右边菜单背景颜色
        menus: [{}],  //菜单数据
        systemName: "",  //系统名称
        userinfo: {name:'',role:"管理员"},
        indexUrl: "/user/index",
        items:[]
      }
    },
    components:{ //注册插件
      navMenu
    },
    created(){
      this.getMenu();
      this.getSystemInfo();
    },
    mounted(){
      this.setDocmentTitle();
      var userinfo = localStorage.getItem("userinfo")
      if(userinfo){
        this.userinfo = JSON.parse(userinfo);
      }
    },
    methods:{
      pwd(){
        this.$http.clearAuthorization();
        if(this.$env.Conf.cookieName){
          VueCookies.remove(this.$env.Conf.cookieName);
        }
        window.location.href = this.$env.Conf.loginWebHost + "/" + this.$env.Conf.ident + "/changepwd";
      },
      signOutM() {
        this.$http.clearAuthorization();
        var logouturl="";//如果想退出后跳转的地址，请设置值
        var returnURL = window.location.href;
        var redirectURL = "?returnurl="+returnURL;
        if (logouturl){
          redirectURL = "?logouturl="+logouturl;
        }
        window.location  = this.$env.Conf.loginWebHost+"/"+this.$env.Conf.ident+"/login"+redirectURL;
      },
      getMenu(){
        this.$http.get("/sso/member/menus/get")
          .then(res => {
            this.menus = res;
            this.$refs.NewTap.open("用户管理", this.indexUrl); //修改此处的菜单名与地址
            this.getUserOtherSys();
          })
          .catch(err => {
            console.log(err)
          });
      },
      //获取系统的相关数据
      getSystemInfo() {
        this.$http.get("/sso/system/info/get")
        .then(res => {
          this.themes = res.theme;
          this.systemName = res.name;
          this.logo = res.logo;
          this.setDocmentTitle();
          
        }).catch(err => {
          console.log(err);
        })
      },
      //用户可用的其他系统
      getUserOtherSys() {
        this.$http.get("/sso/member/systems/get")
        .then(res => {
            this.items = (function (systems) {
              if (!systems || !systems.length) {
                  return []
              }
              var items = [];
              systems.forEach(element => {
                  items.push({
                    name: element.name,
                    path: element.index_url.substr(0, element.index_url.lastIndexOf("/")),
                    type: "blank"
                  })
              });
              return items;
          })(res);
        })
        .catch(err => {
          console.log(err);
        })
      },
      setDocmentTitle() {
        document.title = this.systemName;
      }
    
    }
  }
</script>

```

2. 修改内容

```javascript
//1. indexUrl值
//2. 修改默认地址

``` 
#### 1.5 ssocallback路由处理 

1. 创建  src/pages/member/sso.callback.vue 文件，复制添加如下内容。
```vue
<template></template>

<script>
export default {
  data() {
    return {};
  },
  mounted() {
    this.validSsoLogin();
  },
  methods: {
    validSsoLogin() {
      var returnURL = this.$route.query.returnurl;
      this.$http
        .post("/sso/login/verify", { code: this.$route.query.code })
        .then(res => {
          localStorage.setItem(
            "userinfo",
            JSON.stringify({ name: res.user_name, role: res.role_name })
          );
          if (returnURL) {
            window.location = returnURL; 
            return;
          }
          this.$router.push("/");
        })
        .catch(err => {
            console.log(err);
        });
    }
  }
};
</script>

```

#### 1.6 添加路由地址

```javascript

import menu from '@/pages/member/menu';
import ssocallback from '@/pages/member/sso.callback.vue';

export default new Router({
  mode: "history",
  routes: [
  {
    path: '/',
    name: 'menu',
    component: menu,
    meta:{
      name:"****" //系统名字
    },
    children: [
      //此处添加业务处理路由地址
    ]
  },
    {
      path: '/ssocallback',
      name: 'ssocallback',
      component: ssocallback
    }
  ]
})



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
    Jwt(jwt.WithExcludes("/sso/login/verify"))
})
/*
注：
如果WithExcludes包含其他的地址，将 “/sso/login/verify” 放在其中即可
*/
```

4. static rewriters 配置处理

```go
//WithRewriters 增加:"/ssocallback"

hydra.OnReady(func() error {
    hydra.Conf.Web("8181"). //端口根据业务自定定义
    Static(static.WithRewriters("/", "/index.htm", "/ssocallback"))
})
```

### 3. sdk接口说明
 
#### 3.1  根据⽤⼾名获取⽤⼾信息
``` go
GetUserInfoByName(userName string)
```
输⼊参数

| 参数     |  类型  |  说明  |
| -------- | :----: | :----: |
| userName | string | ⽤⼾名称 |

输出

| 参数      |  类型  |          说明           |
| --------- | :----: | :---------------------: |
| userName  | string |         ⽤⼾名称          |
| WxOpID    | string |       微信openID        |
| ExtParams | string | ⽤⼾扩展参数(⼀个json对象) |
| UserID    | string |         ⽤⼾标识          |

#### 3.2 获取⽤⼾在某个⼦系统下的菜单数据
```go
GetUserMenu(userID int)
```
输⼊参数

| 参数   | 类型  |  说明  |
| ------ | :---: | :----: |
| userID |  int  | ⽤⼾标识 |

输出

| 参数     |   类型   |   说明   |
| -------- | :------: | :------: |
| ID       |  string  | 菜单标识 |
| Name     |  string  | 菜单名称 |
| Level    |  string  |   级次   |
| IsOpen   |  string  | 是否展开 |
| Icon     |  string  |   图标   |
| SystemID |  string  | 系统标识 |
| Parent   |  string  | ⽗级编号  |
| Path     |  string  | 路由地址 |
| Sortrank |  string  | 排序编号 |
| Children | 对象数组 |  ⼦菜单   |

#### 3.3获取⼦系统信息
``` go
GetSystemInfo()
```
输⼊参数(无)

输出

| 参数   |  类型  |        说明        |
| ------ | :----: | :----------------: |
| ID     | string |      系统标识      |
| Ident  | string | 系统ident(英⽂名称) |
| Name   | string |      系统名称      |
| Theme  | string |      主题样式      |
| Layout | string |     ⻚⾯布局样式     |
| Logo   | string |    系统图标地址    |


#### 3.4 获取当前用户可访问的其他子系统
``` go
GetUserOtherSystems()
```
输⼊参数(无)

输出

| 参数     |  类型  |        说明        |
| -------- | :----: | :----------------: |
| ID       | string |      系统标识      |
| Ident    | string | 系统ident(英⽂名称) |
| Name     | string |      系统名称      |
| IndexUrl | string |   ⼦系统地址 host   |
