'use strict'
const merge = require('webpack-merge')
const prodEnv = require('./prod.env')

module.exports = merge(prodEnv, {
    NODE_ENV: '"development"',
    service: {
        wxcallbackhost:'"http://ssov3.100bm.cn"',
        wxcallbackurl:'"/wxcallback"',
        //url: '"http://192.168.5.78:6687"',     //本地
        //url: '"http://loginapi.sso.18jiayou1.com:6687"'
        url : '"http://192.168.5.78:6687"'
    }
});
