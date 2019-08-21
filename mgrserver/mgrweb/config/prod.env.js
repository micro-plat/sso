'use strict'
module.exports = {
    NODE_ENV: '"production"',
    service: {
        ident:`"sso"`,
        apiHost: '"http://192.168.106.226:6677"', //测试环境 
        ssoWebHost:`"http://192.168.106.226:8091"`, //测试环境
        
        //apiHost: '"http://webapi.sso.18jiayou.com:6677"', //线上环境(用户管理api地址)
        //ssoWebHost:`"http://login.sso.18jiayou.com:8091"`, //线上环境(sso登录跳转地址)
    }
}
