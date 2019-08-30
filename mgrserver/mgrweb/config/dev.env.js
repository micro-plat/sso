'use strict'
const merge = require('webpack-merge')
const prodEnv = require('./prod.env')

module.exports = merge(prodEnv, {
    NODE_ENV: '"development"',
    service: {
        ident:`"sso"`,
        apiHost: '"http://webapi.sso.18jiayou1.com:6677"', //测试环境有dns(用户管理api地址)
        ssoWebHost:`"http://login.sso.18jiayou1.com:8091"`, //测试环境有dns(sso登录跳转地址)
         //apiHost: '"http://192.168.5.78:6677"', //本地
        // ssoWebHost:`"http://192.168.106.226:8091"`, //测试环境
             
        
        // ws: `"ws://192.168.7.188:8099/ws"`,
        // webHost:'"http://coupon2.100bm.cn:6060"',
    }
});
