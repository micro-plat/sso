'use strict'
module.exports = {
    NODE_ENV: '"production"',
    service: {
        wxcallbackhost:'"http://alipaygrs.100bm.cn"',
        wxlogincallbackurl:'"/wxlgcallback"',
        wxbindcallbackurl:'"/wxbindcallback"',

        //url: '"http://192.168.106.226:6687"',   //测试环境
        //url: '"http://loginapi.sso.18jiayou1.com:6687"',   //测试环境有dns(跳转登录的api地址)
        url: '"http://loginapi.sso.18jiayou.com"',   //线上环境(跳转登录的api地址)
    }
}