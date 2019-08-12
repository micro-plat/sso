import $ from "jquery"; 
import VueCookies from 'vue-cookies';

/**
 *  配置sso相关的host)
 * @param {*sso前端host, 登录跳转用到} loginWebHost 
 * @param {*子系统的ident} ident
 */
export function ssoConfig(loginWebHost, ident) {
    window.ssoconfig = {
        loginWebHost: loginWebHost,
        ident:ident,
    }
    var refleshHtml = '<iframe id="ssoreflesh" src="'+ loginWebHost + '/refresh" style="display:none"></iframe>';
    $('body').append(refleshHtml);

    var sso = {
        changeRouteAfterLogin:changeRouteAfterLogin,
        signOut:signOut,
        changePwd:changePwd,
        errPage:errPage
    };
    return sso;
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

    if (window.location.pathname != "/" && window.location.pathname != "/login" && window.location.pathname != "/login/") {
        window.localStorage.setItem("beforeLoginUrl", window.location.pathname);
    }

    var url = window.ssoconfig.loginWebHost + "/" + window.ssoconfig.ident + "/jump";
    if (process.env.NODE_ENV == "development") {
        url += "?callback=" + encodeURIComponent(window.location.protocol + "//" + window.location.host + "/ssocallback");
    }
    window.location.href= url;
}

/**
 * sso登录回调，并相关验证成功后,运行此代码
 * 主要是为了如果是子系统间的调要，只加载相应的页面
 */
function changeRouteAfterLogin(vueRouter,userName, userRole) {
    //保存登录用户的用户名和角色名称(主要是菜单组件要用到)
    localStorage.setItem("userinfo", JSON.stringify({name:userName, role:userRole}));

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
function signOut() {
    VueCookies.remove("__jwt__");
    localStorage.removeItem("__jwt__");
    sessionStorage.removeItem("__jwt__");

    window.location.href = window.ssoconfig.loginWebHost + "/" + window.ssoconfig.ident + "/login";
}

/**
 * 子系统用户修改密码
 * @param {sso修改密码地址,请带上http} changePwdUrl 
 */
function changePwd() {
    VueCookies.remove("__jwt__");
    localStorage.removeItem("__jwt__");
    sessionStorage.removeItem("__jwt__");

    window.location.href = window.ssoconfig.loginWebHost + "/" + window.ssoconfig.ident + "/changepwd";
}

/**
 * sso的错误页面
 */
function errPage(errType) {
    if (!errType) {
        errType = 0
    }
    window.location.href = window.ssoconfig.loginWebHost + "/errpage?type=" + errType;
}