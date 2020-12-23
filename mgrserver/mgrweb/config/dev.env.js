'use strict'
const merge = require('webpack-merge')
const prodEnv = require('./prod.env')

module.exports = merge(prodEnv, {
    NODE_ENV: '"development"',
    // VUE_APP_API_URL:'"http://192.168.5.78:6677"',
    VUE_APP_API_URL:'"http://webapi.sso.18jiayou1.com:6677"',
    SSO_WEB_HOST:`"http://login.sso.18jiayou1.com"`,
    // SSO_WEB_HOST:`"http://192.168.5.78:8091"`,
    IDENT:`"sso"`,
});



