'use strict'
const merge = require('webpack-merge')
const prodEnv = require('./prod.env')

module.exports = merge(prodEnv, {
    NODE_ENV: '"development"',
    service: {
        //url: '"//192.168.7.188:9091"',
        //url: '"http://api.sso2.100bm.cn:6688"',     //线下
     //url: '"http://api.sso.sinopecscsy.com:6688"',   //线上
        url: '"http://192.168.106.226:6687"',     //线下
        ws: `"ws://192.168.7.188:8099/ws"`,
        webHost:'"http://coupon2.100bm.cn:6060"',
    }
});
