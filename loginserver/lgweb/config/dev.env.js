'use strict'
const merge = require('webpack-merge')
const prodEnv = require('./prod.env')

module.exports = merge(prodEnv, {
    NODE_ENV: '"development"',
    service: {
        //wxcallbackhost:'"http://alipaygrs.100bm.cn"',
        //wxlogincallbackurl:'"/wxlgcallback"',
       // wxbindcallbackurl:'"/wxbindcallback"',
        //url: '"http://192.168.5.78:6687"',     //本地
        url: '"http://loginapi.sso.18jiayou1.com:6687"'
    }
});
