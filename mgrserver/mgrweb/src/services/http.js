import Vue from "vue"
import axios from 'axios';

axios.defaults.timeout = 5000;
axios.defaults.withCredentials = true;
axios.defaults.baseURL = process.env.VUE_APP_API_URL || "";

//获取vue实例
var __vue__ = {}
//header请求头key
var headerKey = "__http_response_header__";
var authorizationKey = "authorization";

//提示成功模板信息
var tmplSuccess = { title: "成功", message: "操作成功", type: "success", offset: 50, duration: 2000 }

//提示失败模板信息
var tmplFailed = { title: "错误", message: "网络错误,请稍后再试", type: "error", offset: 50, duration: 2000 }

//根据状态码回调
var statusCodeHandles = { "403":function(response){} } 

/*
* http对象使用时须通过引用并进行初始化
* import http from './http'
* Vue.use(http);
*/
export function Http() { 
    __vue__ = Vue.prototype;
}

//http request 拦截器
axios.interceptors.request.use(
    config => {              
        var headers = {};  
        var defaults = {};
        Object.assign(defaults,Http.prototype.defaults.headers);
        for(var k in defaults){
            headers[k.toLowerCase()] = defaults[k];
        }
        headers[authorizationKey] = localStorage.getItem(authorizationKey);
        for(var k in config.headers||{}){
            headers[k.toLowerCase()] = config.headers[k];
        }

        config.headers = headers;
        return config;
    },
    error => {
        return Promise.reject(error);
    }
);

//http response 拦截器
axios.interceptors.response.use(
    response => {
        console.log("response.headers:",response.headers);
        if(response.headers && response.headers[authorizationKey]){
            if (!Http.prototype.defaults.headers){
                Http.prototype.defaults.headers={};
            }
            Http.prototype.setAuthorization(response.headers[authorizationKey]);
        }        
        return response;
    },
    error => {
        if (error.response) {
            let handle = statusCodeHandles[error.response.status];
            if (handle){
                handle(error.response);
            }
        }
        return Promise.reject(error);
    }
)

Http.prototype.defaults = {
    headers:{"content-type": "application/json; charset=UTF-8"}
};

//设置http请求地址
Http.prototype.setBaseURL = function (apiBaseUrl) {
    axios.defaults.baseURL = apiBaseUrl;
}

//设置根据状态码回调
Http.prototype.addStatusCodeHandle = function (f, code = "*") {
    let vcode = code || "*";
    if(typeof f != "function"){
        return;
    }
    statusCodeHandles[vcode] = f;
}

/**
 * 封装get方法
 * @param url
 * @param data
 * @returns {Promise}
 */
Http.prototype.get = function (url, params = {}, config = {}) {
    return new Promise((resolve, reject) => {
        axios.get(url, { params: params }, config)
            .then(response => {
                if (response.status == 200) {
                    resolve(response.data);
                }
            })
            .catch(err => {
                reject(err)
            })
    })
}

/**
 * 封装post请求
 * @param url
 * @param data
 * @returns {Promise}
 */
Http.prototype.post = function (url, data = {}, config = {}) {
    return new Promise((resolve, reject) => {
        axios.post(url, data, config)
            .then(response => {
                if (response.status == 200) {
                    resolve(response.data);
                }
            }, err => {
                reject(err)
            })
    })
}

/**
 * 封装patch请求
 * @param url
 * @param data
 * @returns {Promise}
 */
Http.prototype.patch = function (url, data = {}, config = {}) {
    return new Promise((resolve, reject) => {
        axios.patch(url, data, config)
            .then(response => {
                if (response.status == 200) {
                    resolve(response.data);
                }
            }, err => {
                reject(err)
            })
    })
}

/**
 * 封装put请求
 * @param url
 * @param data
 * @returns {Promise}
 */
Http.prototype.put = function (url, data = {}, config = {}) {
    return new Promise((resolve, reject) => {
        axios.put(url, data, config)
            .then(response => {
                if (response.status == 200) {
                    resolve(response.data);
                }
            }, err => {
                reject(err);
            });
    });
}


/**
 * 封装delete请求
 * @param url
 * @param data
 * @returns {Promise}
 */

Http.prototype.del = function (url, data = {}, config = {}) {
    return new Promise((resolve, reject) => {
        axios.delete(url, { data: data }, config)
            .then(response => {
                if (response.status == 200) {
                    resolve(response.data);
                }
            }, err => {
                reject(err)
            })
    })
}

/**
 * 二次封装get方法
 * @param url
 * @param data
 * @returns {Promise}
 */

Http.prototype.xget = function (url, params = {}, config = {}, success = "", fail = "") {
    return new Promise((resolve, reject) => {
        Http.prototype.get(url, params, config)
        .then((res)=>{
            showSuccessNotify(success)
            resolve(res)
        }).catch(err => {
            showFailedNotify(fail, err)
            reject(err)
        })
    }).catch(err=>{})
}

/**
 * 二次封装post方法
 * @param url
 * @param data
 * @returns {Promise}
 */

Http.prototype.xpost = function (url, params = {}, config = {}, success = "", fail = "") {
    return new Promise((resolve, reject) => {
        Http.prototype.post(url, params, config)
        .then((res)=>{
            showSuccessNotify(success)
            resolve(res)
        }).catch(err => {
            showFailedNotify(fail, err)
            reject(err)
        })
    }).catch(err=>{})
}

/**
 * 二次封装patch方法
 * @param url
 * @param data
 * @returns {Promise}
 */

Http.prototype.xpatch = function (url, params = {}, config = {}, success = "", fail = "") {
    return new Promise((resolve, reject) => {
        Http.prototype.patch(url, params, config)
        .then((res)=>{
            showSuccessNotify(success)
            resolve(res)
        }).catch(err => {
            showFailedNotify(fail, err)
            reject(err)
        })
    }).catch(err=>{})
}

/**
 * 二次封装put方法
 * @param url
 * @param data
 * @returns {Promise}
 */

Http.prototype.xput = function (url, params = {}, config = {}, success = "", fail = "") {
    return new Promise((resolve, reject) => {
        Http.prototype.put(url, params, config)
        .then((res)=>{
            showSuccessNotify(success)
            resolve(res)
        }).catch(err => {
            showFailedNotify(fail, err)
            reject(err)
        })
    }).catch(err=>{})
}

/**
 * 二次封装del方法
 * @param url
 * @param data
 * @returns {Promise}
 */

Http.prototype.xdel = function (url, params = {}, config = {}, success = "", fail = "") {
    return new Promise((resolve, reject) => {
        Http.prototype.del(url, params, config)
        .then((res)=>{
            showSuccessNotify(success)
            resolve(res)
        }).catch(err => {
            showFailedNotify(fail, err)
            reject(err)
        })
    }).catch(err=>{})
}

//初始化成功模板
Http.prototype.setSuccessTmplt = function (data){
    setTmplt(tmplSuccess,data)
}

//初始化失败模板
Http.prototype.setFailedTmplt = function (data){
    setTmplt(tmplFailed,data)
}

//清空header
Http.prototype.clear = function (){
    localStorage.removeItem(headerKey);
    return
}
Http.prototype.clearAuthorization =function(){
    window.localStorage.removeItem(authorizationKey);    
}
Http.prototype.setAuthorization = function(token){
    if(token){
       window.localStorage.setItem(authorizationKey,token);
    }
}

//显示成功提示
function showSuccessNotify(msg){
    let tmpl = tmplSuccess
    if (msg) {
        if (typeof msg == "string"){
            tmpl.message = msg
        }
    }
    if(typeof msg == "boolean" && !msg){
        return
    }

    if(typeof __vue__.$notify != "function"){
        console.error("未找到提示组件'this.$notify'方法，可能未安装element-ui组件，请先安装相关组件");
        return
    }
    __vue__.$notify(tmpl);
}

//显示失败提示
function showFailedNotify(msg, err){
    let tmpl = tmplFailed
    if (msg) {
        if (typeof msg == "string"){
            tmpl.message = msg + err
        }
        if (typeof msg == "boolean"){
            tmpl.message = tmpl.message + err
        }       
    }
    if(typeof msg == "boolean" && !msg){
        return
    }
    if(typeof __vue__.$notify != "function"){
        console.info("未找到提示组件'this.$notify'方法，可能未安装element-ui组件，请先安装相关组件");
        return
    }
    __vue__.$notify(tmpl);
}

//初始化模板数据
function setTmplt(tmpl, data){
    if(!data){
        return
    }
    if (typeof data == "string"){ 
        tmpl.message = data
        return
    }
    if(typeof data != "object"){
        return
    }
    if (data.hasOwnProperty("title")){
        tmpl.title = data.title
    }
    if (data.hasOwnProperty("message")){
        tmpl.message = data.message
    }
    if (data.hasOwnProperty("offset")){
        tmpl.offset = data.offset
    }
    if (data.hasOwnProperty("duration")){
        tmpl.duration = data.duration
    }
}
