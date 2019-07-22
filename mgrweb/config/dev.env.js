'use strict'
const merge = require('webpack-merge')
const prodEnv = require('./prod.env')

module.exports = merge(prodEnv, {
    NODE_ENV: '"development"',
    service: {
        freshJwt:`"/lg/login/refresh"`,
        changePwd:`"/changepwd"`,
        jumpUrl:`"/jump"`,
        loginUrl:`"/login"`,
        errPage: `"/errpage"`,
        //callbackUrl:`"http://192.168.5.78:8081/ssocallback"`, //本地
        callbackUrl:`"http://192.168.106.152:8081/ssocallback"`,

        ssoApiHost:`"http://192.168.106.226:6687"`,
        ssoWebHost:`"http://192.168.106.226:8091"`,
        url: '"http://192.168.106.152:6677"', //测试环境
        //url: '"http://192.168.5.78:6688"', //本地     
        
        ws: `"ws://192.168.7.188:8099/ws"`,
        webHost:'"http://coupon2.100bm.cn:6060"',
    }
});
