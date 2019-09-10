'use strict'
module.exports = {
    NODE_ENV: '"production"',
    service: {
        //wxcallbackhost:'"http://ssov3.100bm.cn"',
        wxcallbackhost:'"http://login.sso.18jiayou.com"',
        wxcallbackurl:'"/wxcallback"',

        //url: '"http://loginapi.sso.18jiayou1.com:6687"',   //测试环境有dns(跳转登录的api地址)
        url: '"http://loginapi.sso.18jiayou.com"',   //线上环境(跳转登录的api地址)
    }
}