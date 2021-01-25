var __user_info__="__user_info__"
export function Auth(Vue) {
    Auth.prototype.Vue = Vue
}

//checkAuthCode 向服务器发送请求，验证auth code
Auth.prototype.checkAuthCode = function (that ,url){
    //检查请求参数中是否有code
    if (!that.$route.query.code){
       return;
    }

    //检查verify地址
    var verifyURL = url || that.$env.conf.system.verifyURL;
    if(!verifyURL){
        throw new Error("system.verifyURL未配置")
    }

    //从服务器拉取数据
    var userInfo = that.$http.xget(verifyURL,{ code: that.$route.query.code});
    if (!userInfo){
        throw new Error("userInfo数据为空");
    }
    //保存用户信息
    window.localStorage.setItem(__user_info__, JSON.stringify(userInfo)) ;
}

//lognout 退出登录
Auth.prototype.loginout = function(url, logoutURL){

    //清除用户认证信息
    let that = Auth.prototype.Vue.prototype;

    //清除http认证头信息
    that.$http.clearAuthorization();
   
    //清除cookie 
    logoutURL = logoutURL || that.$env.conf.system.logoutURL;
    if (logoutURL){
        that.$http.xget(logoutURL);
    }
    
    if ((!that.$env.conf.sso || !that.$env.conf.sso.host) && !url){
        throw new Error("sso节点或sso.host未配置且退出跳转url为空");
    }
    var redirctURL= "?returnurl=" + encodeURIComponent(window.location.href);
    if(url){
        redirctURL = "?logouturl=" + encodeURIComponent(url);
    }
    //检查loginOutURL是否配置
    window.location = url || that.$env.conf.sso.host + "/" + that.$env.conf.sso.ident + "/login" + redirctURL;    
}

//getUserInfo 获取用户信息
Auth.prototype.getUserInfo = function(){
   let userInfo = window.localStorage.getItem(__user_info__)  
   if (!userInfo){
       return {}
   }
   return JSON.parse(userInfo)
}

//getMenus获取菜单数据
Auth.prototype.getMenus = function(_that, url){
    let that = Auth.prototype.Vue.prototype  
    let menuURL = url || "/sso/member/menus/get";
    return new Promise((resolve, reject) => {
        that.$http.get(menuURL)
        .then(res => { 
            //根据路径查找名称    
            var cur = getMenuItem(res, window.location.pathname);
            _that.$refs.NewTap.open(cur.name, cur.path); //this用menu的this
            resolve(res);
        })
        .catch(err => {
            reject(err)
        })
    });
}

//getSystemInfo获取系统信息
Auth.prototype.getSystemInfo = function(url ){   
    let that = Auth.prototype.Vue.prototype 
    let systemInfoURL = url || "/sso/system/info/get"
    return new Promise((resolve, reject) => {       
        that.$http.get(systemInfoURL)
        .then(res => {
            resolve(res);
        })
        .catch(err => {
            reject(err)
        })
    });
}

//getSystemList获取用户系统列表
Auth.prototype.getSystemList = function(url ){    
    let that = Auth.prototype.Vue.prototype  
    let systemsListURL = url || "/sso/member/systems/get"
    return new Promise((resolve, reject) => {
        that.$http.get(systemsListURL)
        .then(res => {
            resolve(res);
        })
        .catch(err => {
            reject(err)
        })
    });
}

//获取路由name
function getMenuItem(menus, path){    
    for (var i in menus){
        var cur = menus[i];
        if(cur.path == path){
            return cur;
        }
        if(path == "/" && cur.path != "-"){
            return cur;
        }
        if(cur.children){
            var res = getMenuItem(cur.children, path);
            if(res){
                return res;
            }
        }
    }
    return null;
}