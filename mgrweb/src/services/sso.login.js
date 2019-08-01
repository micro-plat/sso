
import router from '../router'
import VueCookies from 'vue-cookies'

/**
 * 跳转登录地址，同时将地址记录下来,回调时要路由到那个页面
 * @param {*sso登录地址,请带上 http / https} ssoJumpUrl 
 * @param {*系统ident} ident 
 * @param {*回调地址, 请带上 http / https} callBackUrl
 */
export function setLoginInfo(ssoJumpUrl, ident, callBackUrl) {
    //由于不知道子系统用的哪种方式在保存jwt,只有全delete
    VueCookies.remove("__jwt__");
    localStorage.removeItem("__jwt__");
    sessionStorage.removeItem("__jwt__");

    if (router.currentRoute.fullPath != "/" && router.currentRoute.fullPath != "/login") {
        window.localStorage.setItem("beforeLoginUrl", router.currentRoute.fullPath);
    }
    window.location.href= ssoJumpUrl + "?ident="+ ident +"&callback=" + encodeURIComponent(callBackUrl);
}

/**
 * sso登录回调，并相关验证成功后,运行此代码
 * 主要是为了如果是子系统间的调要，只加载相应的页面
 */
export function setAfterLogin(vueRouter) {
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
export function signOut(loginUrl) {
    VueCookies.remove("__jwt__");
    localStorage.removeItem("__jwt__");
    sessionStorage.removeItem("__jwt__");

    window.location.href = loginUrl;
}

/**
 * 子系统用户修改密码
 * @param {sso修改密码地址,请带上http} changePwdUrl 
 */
export function changePwd(changePwdUrl) {
    VueCookies.remove("__jwt__");
    localStorage.removeItem("__jwt__");
    sessionStorage.removeItem("__jwt__");

    window.location.href = changePwdUrl;
}