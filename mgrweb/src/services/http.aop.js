import {
    post,
    fetch,
    patch,
    put,
    del
} from './http'
   

export function init(Vue, freshUrl){
    Vue.prototype.$fetch = function (url, params = {}) {
        var result = fetch(url,params);
        changeUrl();
        return result
    }

    Vue.prototype.$post = function (url, params = {}) {
        var result = post(url,params)
        changeUrl();
        return result
    }

    Vue.prototype.$patch = function (url, params = {}) {
        var result = patch(url,params)
        changeUrl();
        return result
    }

    Vue.prototype.$put = function (url, params = {}) {
        var result = put(url,params)
        changeUrl();
        return result
    }

    Vue.prototype.$del = function (url, params = {}) {
        var result = del(url,params)
        changeUrl();
        return result
    }

    if (!freshUrl) {
        console.log("freshUrl不能为空, 这个是为刷新sso的jwt信息")
    }

    var refleshHtml = '<iframe id="ssoreflesh" src="'+ freshUrl + '" style="display:none"></iframe>';
    $('body').append(refleshHtml);
}

function changeUrl() {
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