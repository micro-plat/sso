import axios from 'axios';
const Qs = require('qs');
axios.defaults.timeout = 5000;
axios.defaults.withCredentials = true;
axios.defaults.baseURL = window.globalConfig.url;


let GetTocken = (function () {
    return window.localStorage.getItem("__sso_jwt__");
});

function SetToken(response) {
    window.localStorage.setItem("__sso_jwt__", response.headers.__sso_jwt__);
}

//http request 拦截器
axios.interceptors.request.use(
    config => {

        config.headers = {
          'Content-Type': 'application/x-www-form-urlencoded',
          '__sso_jwt__': GetTocken()
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
        if (response.headers.__sso_jwt__) {
            SetToken(response);
        }
        return response;
    },
    error => {
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
                    console.log("--fetch-",response.data)
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
                    console.log("--post-",response.data)
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
                    console.log("--post-",response.data)
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
                    console.log("--post-",response.data)
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
                    console.log("--post-",response.data)
                    resolve(response.data);
                }
            }, err => {
                reject(err)
            })
    })
}
