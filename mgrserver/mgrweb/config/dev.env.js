'use strict'
const merge = require('webpack-merge')
const prodEnv = require('./prod.env')

module.exports = merge(prodEnv, {
    NODE_ENV: '"development"',
    ident:'"sso"',
    apiURL:'"//localhost:6677"',
    loginWebHost:'"http://ssonew.login0.com:6687"'
});



