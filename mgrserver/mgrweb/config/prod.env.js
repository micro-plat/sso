'use strict'
module.exports = {
    NODE_ENV: '"production"',
    SSO_WEB_HOST:`"http://login.sso.18jiayou.com"`,
    VUE_APP_API_URL:'"http://webapi.sso.18jiayou.com"',
    //SSO_WEB_HOST:`"http://login.sso.18jiayou1.com"`,
    //VUE_APP_API_URL:'"http://webapi.sso.18jiayou1.com:6677"',
    
    IDENT:`"sso"`,

    // service: {
    //     ident:`"sso"`,
    //     //apiHost: '"http://192.168.106.226:6677"', //测试环境 
    //     //ssoWebHost:`"http://192.168.106.226:8091"`, //测试环境
        
    //     apiHost: '"http://webapi.sso.18jiayou1.com:6677"', //测试环境有dns(用户管理api地址)
    //     ssoWebHost:`"http://login.sso.18jiayou1.com:80"`, //测试环境有dns(sso登录跳转地址) 8091

    //     //  apiHost: '"http://webapi.sso.18jiayou.com"', //线上环境(用户管理api地址)
    //     //  ssoWebHost:`"http://login.sso.18jiayou.com"`, //线上环境(sso登录跳转地址)
    // }
}
