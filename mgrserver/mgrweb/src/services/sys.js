var __user_info__="__user_info__"
export function Sys(Vue) {
    Sys.prototype.Vue = Vue
}

//checkAuthCode 向服务器发送请求，验证auth code
Sys.prototype.checkAuthCode = function (router, url){
    //检查请求参数中是否有code
    if (!router.query.code){
        return
    }

    //检查verify地址
    var verifyURL = url || "/sso/login/verify";
    //从服务器拉取数据
    let that = Sys.prototype.Vue.prototype;
    var userInfo = that.$http.xget(verifyURL, {code: router.query.code})
    if (!userInfo){
        throw new Error("userInfo数据为空");
    }
    //保存用户信息
    window.localStorage.setItem(__user_info__, JSON.stringify(userInfo)) ;
}

//lognout 退出登录
Sys.prototype.logout = function(url, logoutURL){
    //清除http认证头信息及cookie
    clear(logoutURL)

    let that = Sys.prototype.Vue.prototype;
    if ((!that.$env.conf.sso || !that.$env.conf.sso.host) && !url){
        throw new Error("sso节点或sso.host未配置且退出跳转url为空");
    }
    var redirctURL= "?returnurl=" + encodeURIComponent(window.location.href);
    if(url){
        redirctURL = "?logouturl=" + encodeURIComponent(url);
    }
    //检查logoutURL是否配置
    window.location = url || that.$env.conf.sso.host + "/" + that.$env.conf.sso.ident + "/login" + redirctURL;    
}

//changePwd 修改密码
Sys.prototype.changePwd = function(url, logoutURL){
    //清除http认证头信息
    clear(logoutURL)

    let that = Sys.prototype.Vue.prototype;
    //跳转到修改密码页面
    window.location.href = url || that.$env.conf.sso.host + "/" + that.$env.conf.sso.ident + "/changepwd";
}

//getUserInfo 获取用户信息
Sys.prototype.getUserInfo = function(){
   let userInfo = window.localStorage.getItem(__user_info__)  
   if (!userInfo){
       return {}
   }
   return JSON.parse(userInfo)
}

//根据路由获取标题
Sys.prototype.getTitle = function(path){
    //获取菜单
    Sys.prototype.getMenus()
    
    //获取系统信息
    Sys.prototype.getSystemInfo()
    
    //获取本地配置的菜单
    let that = Sys.prototype.Vue.prototype  
    var menus = that.$env.conf.menus
   
    //根据路径查找名称
    var cur = Sys.prototype.findMenuItem(menus, path)
    return cur ? cur.name + " - " + that.$env.conf.system.name : "";
}

//递归查找父级菜单
Sys.prototype.findMenuItem = function(menus, path){   
    path = path || window.location.pathname;
    if(path != "/" && path[path.length-1] == '/'){
        path = path.substring(0,path.length-1);
    }
    //递归查找父级菜单
    var cur = getMenuItem(menus, path);
    if (!cur){                
        path = path.substring(0, path.lastIndexOf('/'));
        if (!path){
            return {}
        }
        cur = getMenuItem(menus,path);
    }
    if(!cur){
        cur = getMenuItem(menus,'/');
    }
    return cur;
}

//getMenus获取菜单数据
Sys.prototype.getMenus = function(url){  
    return new Promise((resolve, reject) => {
        let that = Sys.prototype.Vue.prototype  
        //获取本地配置的菜单
        let menus = that.$env.conf.menus
        if (menus){
            let data = typeof menus == "string" ? that.$http.xget(menus) || [] : menus
            resolve(data)
            return
        }       

        //远程获取菜单
        let menuURL = url || "/sso/member/menus/get"
        let res = that.$http.xget(menuURL)

        //保存菜单信息
        Object.assign(that.$env.conf, { menus: res || [] }) 
        resolve(res);
    });
}

//getSystemInfo获取系统信息
Sys.prototype.getSystemInfo = function(url ){   
    return new Promise((resolve, reject) => {   
        let that = Sys.prototype.Vue.prototype 
        //获取本地配置的系统信息
        if (that.$env.conf.system){
            resolve(that.$env.conf.system)
            return
        }  

        //获取远程系统信息
        let systemInfoURL = url || "/sso/system/info/get"
        let res = that.$http.xget(systemInfoURL)
        
        //保存系统信息
        Object.assign(that.$env.conf, { system: res || {} }) 
        resolve(res);
    });
}

//getSystemList获取用户系统列表
Sys.prototype.getSystemList = function(url ){      
    return new Promise((resolve, reject) => {
        let that = Sys.prototype.Vue.prototype  
        //获取本地配置的菜单
        if (that.$env.conf.sysList){
            resolve(that.$env.conf.sysList)
            return
        }  

        //远程获取其它系统
        let systemsListURL = url || "/sso/member/systems/get"
        that.$http.get(systemsListURL)
        .then(res => {             
            Object.assign(that.$env.conf, { sysList: res || [] }) 
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
        if(path == "/" && cur.path && cur.path != "-"){
            return cur;
        }        
        var res = getMenuItem(cur.children || [], path);
        if(res){
            return res;
        }
    }
    return null;
}

//清除http认证头信息及cookie
function clear(logoutURL){
    //清除用户认证信息
    let that = Sys.prototype.Vue.prototype;

    //清除http认证头信息
    that.$http.clearAuthorization();
   
    //清除cookie 
    logoutURL = logoutURL || "/sso/logout";
    if (logoutURL){
        that.$http.xget(logoutURL);
    }
}