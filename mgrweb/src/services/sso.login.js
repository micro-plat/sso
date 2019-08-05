import VueCookies from 'vue-cookies'
import router from '../router'

/**
 *  配置sso相关的host 在main中调用(Vue.prototype.$post = post;前)
 * @param {*sso前端host, 登录跳转用到} loginWebHost 
 * @param {*sso后端host, 刷新token用到} loginApiHost
 * @param {*子系统的ident} ident
 */
export function ssoConfig(loginWebHost, loginApiHost, ident) {
    window.ssoconfig = {
        loginWebHost: loginWebHost,
        loginApiHost: loginApiHost,
        ident:ident,
    }
    var refleshHtml = '<iframe id="ssoreflesh" src="'+ loginApiHost + '/lg/login/refresh" style="display:none"></iframe>';
    $('body').append(refleshHtml);
}

/**
 * 改变url，达到刷新sso token的目的
 */
export function changeUrl() {
    var url = $("#ssoreflesh").attr("src");
    if (url) {
        var index = url.indexOf("?",0);
        if (index > 0) {
            url = url.substr(0, index + 1);
        } else {
            url += "?"
        }
        $("#ssoreflesh").attr("src", url + "random=" + Date.now());
    }
}

/**
 * 跳转登录地址，同时将地址记录下来,回调时要路由到那个页面
 * @param {*sso登录地址,请带上 http / https} ssoJumpUrl 
 * @param {*系统ident} ident 
 * @param {*回调地址, 请带上 http / https} callBackUrl
 */
export function setRouteBeforeLogin() {
    VueCookies.remove("__jwt__");
    localStorage.removeItem("__jwt__");
    sessionStorage.removeItem("__jwt__");

    if (router.currentRoute.fullPath != "/" && router.currentRoute.fullPath != "/login") {
        window.localStorage.setItem("beforeLoginUrl", router.currentRoute.fullPath);
    }

    //加上 ident 标识
    window.location.href= 
        window.ssoconfig.loginWebHost + "/jump/" + window.ssoconfig.ident
        //+"?callback=" + encodeURIComponent(window.location.protocol + "//" + window.location.host + "/ssocallback");
}

/**
 * sso登录回调，并相关验证成功后,运行此代码
 * 主要是为了如果是子系统间的调要，只加载相应的页面
 */
export function changeRouteAfterLogin(vueRouter) {
    var oldPath = window.localStorage.getItem("beforeLoginUrl");
    localStorage.removeItem("beforeLoginUrl");
    if (oldPath && oldPath != "/" && oldPath.indexOf("/external") == 0) {
        vueRouter.push(oldPath);
        return;
    }
    vueRouter.push("/");
}

/**
 * 子系统退出登录,会跳转到sso登录界面
 * @param {sso登录地址，和跳转地址不一样, 请注意,不然退不出去, 请带上http} loginUrl 
 */
export function signOut() {
    VueCookies.remove("__jwt__");
    localStorage.removeItem("__jwt__");
    sessionStorage.removeItem("__jwt__");

    window.location.href = window.ssoconfig.loginWebHost + "/login/" + window.ssoconfig.ident;
}

/**
 * 子系统用户修改密码
 * @param {sso修改密码地址,请带上http} changePwdUrl 
 */
export function changePwd() {
    VueCookies.remove("__jwt__");
    localStorage.removeItem("__jwt__");
    sessionStorage.removeItem("__jwt__");

    window.location.href = window.ssoconfig.loginWebHost + "/changepwd/" + window.ssoconfig.ident;
}

/**
 * sso的错误页面
 */
export function errPage() {
    window.location.href = window.ssoconfig.loginWebHost + "/errpage";
}