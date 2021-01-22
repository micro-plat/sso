import axios from 'axios';
import Vue from "vue";
import $ from 'jquery';

axios.defaults.timeout = 5000;
axios.defaults.withCredentials = true;
axios.defaults.baseURL = process.env.VUE_APP_API_URL || "";

var __vue__ = {}
//header请求头key
var authorizationKey = "authorization";

//提示成功模板信息
var tmplSuccess = { title: "成功", message: "操作成功", type: "success", offset: 50, duration: 2000 }

//提示失败模板信息
var tmplFailed = { title: "错误", message: "网络错误,请稍后再试", type: "error", offset: 50, duration: 2000 }

//根据状态码回调
var statusCodeHandles = { "403":function(){} } 

/*
* http对象使用时须通过引用并进行初始化
* import { Http } from 'qxnw-http';
* Vue.prototype.$http = new Http(); // http初始化
* 或者 
* import { InitHttp } from 'qxnw-http';
* InitHttp(Vue.prototype)
*
*使用方法
*async queryAsix(){    通过await同步获取接口数据
*    var data = await this.$http.post("/dds/dictionary/get", {}, {}, true)
*    console.log("daaa", data);
*},
*
*query(){ 通过回调获取数据
*   this.$http.post("/dds/dictionary/get", {}, {}, true)
*   .then(res=>{ 
*       //成功回调
*   }).catch(error=>{
*       //失败回调
*   })
*},
*
*getConf(){   其中xget,xpost,xput,xdel都是同步请求，直接获取返回数据
*    var data = this.$http.xpost("/dds/dictionary/get", { dic_type: "operate_action" }, "", false) || {}
*},
*
*/
export function Http() {
    __vue__ = Vue.prototype
}

//http request 拦截器
axios.interceptors.request.use(
    config => {              
        config.headers = getHeaders(config);
        return config;
    },
    error => {
        return Promise.reject(error);
    }
);

//http response 拦截器
axios.interceptors.response.use(
    response => {
        if(response.headers && response.headers[authorizationKey]){
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

//默认 content-type
Http.prototype.defaults = {
    headers: {"content-type": "application/json; charset=UTF-8"}
};

//设置http请求地址
Http.prototype.setBaseURL = function (apiBaseUrl) {
    axios.defaults.baseURL = apiBaseUrl
}

//设置根据状态码回调
Http.prototype.addStatusCodeHandle = function (f, code = "*") {
    let vcode = code || "*"
    if(typeof f != "function"){
        return
    }
    statusCodeHandles[vcode] = f
}

/**
 * 封装get方法
 * @param url
 * @param data
 * @returns {Promise}
 */
Http.prototype.get = function (url, params = {}, config = {}, success = "", fail = "") {
    return new Promise((resolve, reject) => {
        axios.get(url, { params: params }, config)
            .then(response => {
                if (response.status == 200) {
                    showSuccessNotify(success)                    
                    resolve(response.data);
                }
            })
            .catch(err => {
                showFailedNotify(fail, err)
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
Http.prototype.post = function (url, data = {}, config = {}, success = "", fail = "") {
    return new Promise((resolve, reject) => {
        axios.post(url, data, config)
            .then(response => {
                if (response.status == 200) {
                    showSuccessNotify(success)
                    resolve(response.data);
                }
            }, err => {
                showFailedNotify(fail, err)
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
Http.prototype.put = function (url, data = {}, config = {}, success = "", fail = "") {
    return new Promise((resolve, reject) => {
        axios.put(url, data, config)
            .then(response => {
                if (response.status == 200) {
                    showSuccessNotify(success)
                    resolve(response.data);
                }
            }, err => {
                showFailedNotify(fail, err)
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

Http.prototype.del = function (url, data = {}, config = {}, success = "", fail = "") {
    return new Promise((resolve, reject) => {
        axios.delete(url, { data: data }, config)
            .then(response => {
                if (response.status == 200) {
                    showSuccessNotify(success)
                    resolve(response.data);
                }
            }, err => {
                showFailedNotify(fail, err)
                reject(err)
            })
    })
}

/**
 * ajax同步封装get方法
 * @param url
 * @param data
 * @returns {Promise}
 */
Http.prototype.xget = function (url, params = {}, config = {}, success = "", fail = "") {
    return Http.prototype.ajax("GET", url, params, config, success, fail)
}

/**
 * ajax同步封装post方法
 * @param url
 * @param data
 * @returns {Promise}
 */
Http.prototype.xpost = function (url, params = {}, config = {}, success = "", fail = "") {
    return Http.prototype.ajax("POST", url, params, config, success, fail)
}

/**
 * ajax同步封装put方法
 * @param url
 * @param data
 * @returns {Promise}
 */
Http.prototype.xput = function (url, params = {}, config = {}, success = "", fail = "") {
    return Http.prototype.ajax("PUT", url, params, config, success, fail)
}

/**
 * ajax同步封装del方法
 * @param url
 * @param data
 * @returns {Promise}
 */
Http.prototype.xdel = function (url, params = {}, config = {}, success = "", fail = "") {
    return Http.prototype.ajax("DELETE", url, params, config, success, fail)
}

Http.prototype.ajax = function (method, url, params, config, success, fail){
    var result = {}
    $.ajax({
        type: method, //请求方式
        async: false, // fasle表示同步请求，true表示异步请求
        xhrFields: { withCredentials: true },        
        headers: getHeaders(config||{}), 
        beforeSend:function(jqXHR,settings){
            settings.contentType = settings.headers["content-type"]
            if((settings.contentType||"").indexOf("/json")>0){
                settings.data = JSON.stringify(params);
            }
        },  
        url: getURL(url),//请求地址
        data: params, //请求参数
        success: function(res) { //请求成功   
            showSuccessNotify(success)   
            result = res
        },
        error : function(err){  //请求失败，包含具体的错误信息  
            showFailedNotify(fail, err)       
        },
        complete(xhr){
            var token = xhr.getResponseHeader(authorizationKey);
            if(token){
                Http.prototype.setAuthorization(token)
            }
        }
    })
    return result
}

//初始化成功模板
Http.prototype.setSuccessTmplt = function (data){
    setTmplt(tmplSuccess,data)
}

//初始化失败模板
Http.prototype.setFailedTmplt = function (data){
    setTmplt(tmplFailed,data)
}

//消除header
Http.prototype.clearAuthorization =function(){
    window.localStorage.removeItem(authorizationKey); 
}

//设置需要的header
Http.prototype.setAuthorization = function(token){
    if(token){
       window.localStorage.setItem(authorizationKey,token);
    }
}

//显示成功提示
function showSuccessNotify(msg){   
    if (!msg) {
        return
    }

    let tmpl = tmplSuccess
    if (typeof msg == "string"){
        tmpl.message = msg
    }  
    if(typeof __vue__.$notify != "function"){
        console.info("还未安装element-ui组件，请先安装相关组件");
        return
    }
    __vue__.$notify(tmpl);
}

//显示失败提示
function showFailedNotify(msg, err){   
    if (!msg) {    
        return
    }

    let tmpl = tmplFailed
    if (typeof msg == "string"){
        tmpl.message = msg + err
    }  
    if(typeof msg == "boolean"){
        tmpl.message = tmpl.message + err
    }
    if(typeof __vue__.$notify != "function"){
        console.info("还未安装element-ui组件，请先安装相关组件");
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
    if (Object.prototype.hasOwnProperty.call(data, "title")){
        tmpl.title = data.title
    }
    if (Object.prototype.hasOwnProperty.call(data, "message")){
        tmpl.message = data.message
    }
    if (Object.prototype.hasOwnProperty.call(data, "offset")){
        tmpl.offset = data.offset
    }
    if (Object.prototype.hasOwnProperty.call(data, "duration")){
        tmpl.duration = data.duration
    }
} 

//获取response返回的header  
function getHeaders(config){ 
    var headers = {};  
    var defaults = {};
    Object.assign(defaults,Http.prototype.defaults.headers);
    for(var k in defaults){
        headers[k.toLowerCase()] = defaults[k];
    }
    var token = localStorage.getItem(authorizationKey);
    if(token){
        headers[authorizationKey] = token;
    }
    for(var k in config.headers||{}){
        headers[k.toLowerCase()] = config.headers[k];
    }
    return headers
}

function getURL(url){
    if(url.indexOf(".") == 0) {
        return url
    }
    return axios.defaults.baseURL + url  //请求地址
}