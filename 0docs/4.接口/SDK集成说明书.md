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
    "name": "用户系统",
    "copyright": { //版权信息      
        "company": "四川千行你我科技股份有限公司",
        "code": "蜀ICP备20003360号"
    },
    "api": { //接口相关调用
        "host": "http://ssov4.100bm0.com:6677", //后端api接口地址
        "verifyURL": "/sso/login/verify", //sso code验证相关（固定）
        "confURL": "",                    //服务端配置接口
        "enumURL": "/dds/dictionary/get", //枚举获取地址
        "logoutURL": "/sso/logout"        //退出地址
    },
    "sso": {
        "ident": "sso",                    //系统标识
        "host": "http://ssov4.100bm0.com:6687"//sso登录戏台的地址
    }
}

```
#### 1.2 修改main.js 

```javascript
//1. 添加引用（ services 是js包存放的位置）
import utility from './utility'
Vue.use(utility);


//2. 增加路由拦截处理
router.beforeEach((to, from, next) => {
  /* 路由发生变化修改页面title */
  Vue.prototype.$sys.checkAuthCode(to)
  if (to.path != "/") {
      document.title = Vue.prototype.$sys.getTitle(to.path)
  }
  next()
})


```
 
2. 后端提供配置读取服务(如果需要)

```text
	请配置env.conf.json 文件中 api.confURL 为后端配置提供地址
  注：
     1. 接口地址必须返回json 格式数据

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
import navMenu from "nav-menu"; // 引入
export default {
  name: "app",
  data() {
    return {
      logo: "",
      copyright:
        (this.$env.conf.copyright.company || "") +
        "Copyright©" +
        new Date().getFullYear() +
        "版权所有",
      copyrightcode: this.$env.conf.copyright.code,
      themes: "", //顶部左侧背景颜色,顶部右侧背景颜色,右边菜单背景颜色
      menus: [{}], //菜单数据
      systemName: "", //系统名称
      userinfo: {},
      items: [],
    };
  },
  components: {
    //注册插件
    navMenu,
  },
  created() {
  },
  mounted() {
    console.log("----------", this.$route.query);
    this.getMenu();
    this.getSystemInfo();
    this.userinfo = this.$sys.getUserInfo();
  },
  methods: {
    pwd() {
        this.$sys.changePwd();
    },
    signOutM() {
      this.$sys.logout();
    },
    getMenu() {
      this.$sys.getMenus().then((res) => {
        this.menus = res;
        this.getUserOtherSys();
        var cur = this.$sys.findMenuItem(res)
        this.$refs.NewTap.open(cur.name, cur.path);
      });
    },
    //获取系统的相关数据
    getSystemInfo() {
      this.$sys.getSystemInfo().then((res) => {
        this.themes = res.theme;
        this.systemName = res.name;
        this.logo = res.logo;
      });
    },
    //用户可用的其他系统
    getUserOtherSys() {
      this.$sys.getSystemList().then((res) => {
        this.items = res;
      });
    }
  },
};
</script>

```
 

#### 1.6 添加路由地址

```javascript

import menu from '@/pages/member/menu';
 
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
  }
  ]
})

```

### 2. 后端项目集成

1. 引入sdk包
2. 
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
    procObj, err := appConf.GetProcessorConf()
    if err != nil {
      return err
    }

    prefix := procObj.ServicePrefix
  
    if err := ssoSdk.Config(appcfg.SSOApiHost,
      appcfg.Ident,
      appcfg.Secret,
      ssoSdk.WithAuthIgnore(prefix, appcfg.AuthIgnores...)); err != nil {
      return err
    }
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

```go
var Archive="mgr.static.zip"
hydra.OnReady(func() error {
    hydra.Conf.Web("8181"). //端口根据业务自定定义
    Static(static.WithAutoRewrite(),static.WithAssetsPath(Archive))
})
```
