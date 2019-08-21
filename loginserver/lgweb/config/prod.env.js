'use strict'
module.exports = {
    NODE_ENV: '"production"',
    service: {
        wxcallbackhost:'"http://alipaygrs.100bm.cn"',
        wxlogincallbackurl:'"/wxlgcallback"',
        wxbindcallbackurl:'"/wxbindcallback"',

        url: '"http://192.168.106.226:6687"',   //测试环境
        //url: '"http://loginapi.sso.18jiayou.com:6687"',   //线上环境(跳转登录的api地址)
    }
}