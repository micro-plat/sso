'use strict'
const merge = require('webpack-merge')
 
module.exports = merge({
    NODE_ENV: '"development"',
    service: {
        wxcallbackhost: '"http://ssov4.100bm.cn"',
        wxcallbackurl: '"/wxcallback"',
        url: '',   
        codeLabel: '"短信验证码"',
        codeHolder: '"请输入短信验证码"',
        sendBtnLabel:'"获取短信验证码"',
        showText:'"短信验证码发送成功"'
    }
});
