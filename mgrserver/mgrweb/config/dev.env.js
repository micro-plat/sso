'use strict'
const merge = require('webpack-merge')
const prodEnv = require('./prod.env')

module.exports = merge(prodEnv, {
    NODE_ENV: '"development"',
    VUE_APP_API_URL:'"http://192.168.0.103:6677"',
    //SSO_WEB_HOST:`"http://login.sso.18jiayou1.com:80"`,
    SSO_WEB_HOST:`"http://login.sso.18jiayou1.com:8026"`,
    IDENT:`"sso"`,
});



