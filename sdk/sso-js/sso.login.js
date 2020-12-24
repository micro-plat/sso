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
        errPage:errPage,
        transformSysInfo:transformSysInfo
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
    VueCookies.remove("__sso_jwt__");
    localStorage.removeItem("__sso_jwt__");
    sessionStorage.removeItem("__sso_jwt__");

    if (window.location.pathname != "/" && window.location.pathname != "/login" && window.location.pathname != "/login/") {
        window.localStorage.setItem("beforeLoginUrl", window.location.pathname);
    }

    // var queryInfo = window.location.search;
    // var sourceUrl = ""
    // if (queryInfo && queryInfo.indexOf("source=") >= 0) {
    //     //是外部系统跳转过来(反正不是我们用golang开发的系统)
    //     var queryArray = queryInfo.split("&")
    //     var sourceInfo = getSourceInfo(queryArray);
    //     if (sourceInfo) {
    //         sourceUrl = "?source=" + sourceInfo.source + "&sessionid=" + sourceInfo.sessionId;
    //     }
    // }
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
    VueCookies.remove("__sso_jwt__");
    localStorage.removeItem("__sso_jwt__");
    sessionStorage.removeItem("__sso_jwt__");

    window.location.href = window.ssoconfig.loginWebHost + "/" + window.ssoconfig.ident + "/login";
}

/**
 * 子系统用户修改密码
 * @param {sso修改密码地址,请带上http} changePwdUrl 
 */
function changePwd() {
    VueCookies.remove("__sso_jwt__");
    localStorage.removeItem("__sso_jwt__");
    sessionStorage.removeItem("__sso_jwt__");

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

/**
 * 转换系统地址
 * @param {*} systems 
 */
function transformSysInfo(systems) {
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
}


/**
 * 返回source及sessionid等信息(时间紧没有使用工厂的方式)
 * @param {*} queryArray 
 */
// function getSourceInfo(queryArray) {
//     if (!queryArray || !queryArray.length) {
//         return ""
//     }
//     var source = "", sessionId = "";
//     for (var index = 0; index < queryArray.length; index++) { 
//         if (queryArray[index].toLowerCase().includes("source")) {
//             source = queryArray[index].split("=")[1]
//         }
//         if (queryArray[index].toLowerCase().includes("sessionid")) {
//             sessionId = queryArray[index].split("=")[1]
//         }
//         if (source && sessionId) {
//             return {
//                 source: source,
//                 sessionid:sessionId
//             }
//         }
//      }
//     return null
// }