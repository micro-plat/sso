'use strict'
const merge = require('webpack-merge')
const prodEnv = require('./prod.env')

module.exports = merge(prodEnv, {
    NODE_ENV: '"development"',
    VUE_APP_API_URL:'"http://192.168.0.103:6678"',
    //IDENT:'"oil_station"'
    IDENT:'"sso"'
});
