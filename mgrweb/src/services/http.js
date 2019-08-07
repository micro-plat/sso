import axios from 'axios';
import {changeUrl, setRouteBeforeLogin} from './sso.login.js'

const Qs = require('qs');
axios.defaults.timeout = 5000;
axios.defaults.withCredentials = true;
axios.defaults.baseURL = process.env.service.url;

let GetTocken = (function () {
    return window.localStorage["__jwt__"]
});

/**
 * http初始话
 * @param {前端对应apiHost} apiUrl
 */
export function httpConfig(apiBaseUrl) {
    axios.defaults.baseURL = apiBaseUrl;

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
            window.localStorage["__jwt__"] = response.headers.__jwt__
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

export function fetch(url, params = {}) {
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

export function post(url, data = {}) {
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

export function patch(url, data = {}) {
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

export function put(url, data = {}) {
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

export function del(url, data = {}) {
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