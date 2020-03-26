'use strict'
module.exports = {
    NODE_ENV: '"production"',
    //这个配置文件是hbs_sso使用, 专门管理sass用户信息的

    //SSO_WEB_HOST:`"http://login.sso.18jiayou0.com:8091"`,
    //VUE_APP_API_URL:'"http://webapi.sso.18jiayou.com"', //内部用户系统用
    //VUE_APP_API_URL:'"http://api.bss.sso.17ebs.18jiayou0.com:6677"', //sass用户系统用 
    //VUE_APP_API_URL:'"http://api.bss.sso.17ebs.18jiayou.com"', //sass用户系统用 正式环境
    SSO_WEB_HOST:`"http://login.sso.18jiayou.com"`,
    VUE_APP_API_URL:'"http://api.bss.sso.hbs.18jiayou.com:6677"', //内部用户系统用
    //VUE_APP_API_URL:'"http://webapi.sso.18jiayou1.com:6677"',
    IDENT:`"hbs_sso"`,
}