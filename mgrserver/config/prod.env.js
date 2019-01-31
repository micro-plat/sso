'use strict'
module.exports = {
    NODE_ENV: '"production"',
    service: {
        // url: '"http://api.sso.sinopecscsy.com:6688"',   //线上
        // ws: `"ws://api.sso.sinopecscsy.com:6689/ws"`,
        // webHost:'"http://web.coupon.sinopecscsy.com"',
        //url: '"http://api.sso2.100bm.cn:6688"',     //线下
        url: '"http://192.168.106.152:6688"',     //线下
        ws: `"ws://api.sso2.100bm.cn:6689/ws"`,
        webHost:'"http://coupon2.100bm.cn:6060"',
    }
}
