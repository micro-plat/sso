import axios from 'axios';
import router from '../router'
const Qs = require('qs');
axios.defaults.timeout = 5000;
axios.defaults.withCredentials = true;
axios.defaults.baseURL = process.env.VUE_APP_API_URL;

let GetTocken = (function () {
    return window.localStorage.getItem("__jwt__");
});

function SetToken(response) {
    window.localStorage.setItem("__jwt__", response.headers.__jwt__);
}

/**
 * http初始话
 * @param {前端对应后台的api地址 apiHost} apiUrl
 * @param {jwt存储地方, sessionStorage:就是存储在sessionStorage, localStorage:存储在localStorage} storagePlace}
 */
export function httpConfig() {
    return {
        get: get,
        post:post,
        patch:patch,
        put:put,
        del:del
    }
}

//http request 拦截器
axios.interceptors.request.use(
    config => {
        var userName = '';
        var userInfo = localStorage.getItem("userinfo")
        if (userInfo) {
            userName = JSON.parse(userInfo).name || '';
        }
        
        config.headers = {
          'X-Request-Id':userName + '-' + guid(),
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
        return response;
    },
    error => {
        if (error.response.status == 403) {
            return router.push({ path: '/login'});
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

function get(url, params = {}, config={}) {
    return new Promise((resolve, reject) => {
        axios.get(url, {params: params}, config)
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

 function post(url, data = {}, config={}) {
    data = Qs.stringify(data)
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

function patch(url, data = {}, config={}) {
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

function put(url, data = {}, config={}) {
    data = Qs.stringify(data)
    return new Promise((resolve, reject) => {
        axios.put(url, data, config)
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

function del(url, data = {}, config={}) {
    return new Promise((resolve, reject) => {
        axios.delete(url, {data:data}, config)
            .then(response => {
                if (response.status == 200) {
                    resolve(response.data);
                }
            }, err => {
                reject(err)
            })
    })
}

function guid() {
    return 'xxxxxxxx-xxxx-4xxx'.replace(/[x]/g, function(c) {
        var r = Math.random()*16|0, v = c == 'x' ? r : (r&0x3|0x8);
        return v.toString(16);
    });
}