'use strict'
const merge = require('webpack-merge')
const prodEnv = require('./prod.env')

module.exports = merge(prodEnv, {
    NODE_ENV: '"development"',
    service: {
        wxcallbackhost: '"http://ssov3.100bm.cn"',
        wxcallbackurl: '"/wxcallback"',
        url: '"http://192.168.5.94:6687"',   
        codeLabel: '"短信验证码"',
        codeHolder: '"请输入短信验证码"',
        sendBtnLabel:'"获取短信验证码"',
        showText:'"短信验证码发送成功"'
        //url: '"http://loginapi.sso.18jiayou1.com:6687"'
    }
});
