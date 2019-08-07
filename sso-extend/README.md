# 使用说明

#### 1. 引入(导入到本地)

在main.js中注入

``` js


import {ssoConfig} from './services/sso.login'
var serviceConfig = process.env.service;
Vue.prototype.$sso = ssoConfig(serviceConfig.ssoWebHost, serviceConfig.ssoApiHost, "sso"); 

import {httpConfig} from './services/http'
Vue.prototype.$http = httpConfig(serviceConfig.url, "local")

```

#### 2. 所有http的交互

``` js
在vue页面中使用


this.$http.get();
this.$http.post();
this.$http.put();
this.$http.del();
```



