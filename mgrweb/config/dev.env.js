'use strict'
const merge = require('webpack-merge')
const prodEnv = require('./prod.env')

module.exports = merge(prodEnv, {
    NODE_ENV: '"development"',
    service: {
        jumpUrl:"/jump",
        callbackUrl:"ssocallback",
        ssoHost:"http://192.168.5.78:8081",
        url: '"http://192.168.5.78:6688"',     
        
        ws: `"ws://192.168.7.188:8099/ws"`,
        webHost:'"http://coupon2.100bm.cn:6060"',
    }
});
