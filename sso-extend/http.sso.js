import axios from 'axios';
import {changeUrl, setRouteBeforeLogin} from './sso.login.js'

const Qs = require('qs');
axios.defaults.timeout = 5000;
axios.defaults.withCredentials = true;
axios.defaults.baseURL = ""; //process.env.service.url;
var storagePlace = "";

let GetTocken = (function () {
    if (!storagePlace){
        return "";
    }
    var jwt = window.localStorage.getItem("__jwt__");
    if (storagePlace == "sessionStorage") {
        jwt = window.sessionStorage.getItem("__jwt__");
    }
    return jwt;
});

function SetToken(response) {
    if (!storagePlace) {
        return;
    }
    if (storagePlace == "sessionStorage") {
        window.sessionStorage.setItem("__jwt__", response.headers.__jwt__);
    } else {
        window.localStorage.setItem("__jwt__", response.headers.__jwt__);
    }
}

/**
 * http初始话
 * @param {前端对应后台的api地址 apiHost} apiUrl
 * @param {jwt存储地方, sessionStorage:就是存储在sessionStorage, localStorage:存储在localStorage} storagePlace}
 */
export function httpConfig(apiBaseUrl, storagePlace) {
    axios.defaults.baseURL = apiBaseUrl;
    storagePlace = storagePlace;
    return {
        get: fetch,
        post:post,
        patch:patch,
        put:put,
        del:del
    }
}

//http request 拦截器
axios.interceptors.request.use(
    config => {

        config.headers = {
          'Content-Type': 'application/x-www-form-urlencoded',
          '__jwt__': GetTocken()
        };

        return config;
    },
    error => {
        return Promise.reject(err);
    }
);

//http response 拦截器
axios.interceptors.response.use(
    response => {
        if (response.headers.__jwt__) {
            SetToken(response);
        }
        if (response.status == 200){
            changeUrl(); //刷新sso token
        }
        return response;
    },
    error => {

        if (error.response.status == 403) {
            setRouteBeforeLogin();
        }
        return Promise.reject(error)
    }
)


/**
 * 封装get方法
 * @param url
 * @param data
 * @returns {Promise}
 */

function fetch(url, params = {}) {
    return new Promise((resolve, reject) => {
        axios.get(url, {
                params: params
            })
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

 function post(url, data = {}) {
    data = Qs.stringify(data)
    return new Promise((resolve, reject) => {
        axios.post(url, data)
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

function patch(url, data = {}) {
    return new Promise((resolve, reject) => {
        axios.patch(url, data)
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

function put(url, data = {}) {
    data = Qs.stringify(data)
    return new Promise((resolve, reject) => {
        axios.put(url, data)
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
 * 封装delete请求
 * @param url
 * @param data
 * @returns {Promise}
 */

function del(url, data = {}) {
    return new Promise((resolve, reject) => {
        axios.delete(url, data)
            .then(response => {
                if (response.status == 200) {
                    resolve(response.data);
                }
            }, err => {
                reject(err)
            })
    })
}