import {
    post,
    fetch,
    patch,
    put,
    del
} from './http'
   

export function init(Vue){
    Vue.prototype.$fetch = function (url, params = {}) {
        var result = fetch(url,params);
        console.log("ins..fetch....end")
        changeUrl();
        return result
    }

    Vue.prototype.$post = function (url, params = {}) {
        var result = post(url,params)
        console.log("ins..post....end")
        changeUrl();
        return result
    }

    Vue.prototype.$patch = function (url, params = {}) {
        var result = patch(url,params)
        console.log("ins..patch....end")
        changeUrl();
        return result
    }

    Vue.prototype.$put = function (url, params = {}) {
        var result = put(url,params)
        console.log("ins..put....end")
        changeUrl();
        return result
    }

    Vue.prototype.$del = function (url, params = {}) {
        var result = del(url,params)
        console.log("ins..del....end")
        changeUrl();
        return result
    }
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