'use strict'
module.exports = {
    NODE_ENV: '"production"',
    service: {
        // freshJwt:`"/lg/login/refresh"`,
        // changePwd:`"/changepwd"`,
        // jumpUrl:`"/jump"`,
        // loginUrl:`"/login"`,
        // errPage: `"/errpage"`,
        // callbackUrl:`"/ssocallback"`, //本地
        //callbackUrl:`"http://192.168.106.152:8081/ssocallback"`,

        ssoApiHost:`"http://192.168.106.226:6687"`,
        ssoWebHost:`"http://192.168.106.226:8091"`,
        url: '"http://192.168.106.152:6677"', //测试环境
        //url: '"http://192.168.5.78:6688"', //本地     


        //url: '"http://api.sso.sinopecscsy.com:6688"',   //线上
        //ws: `"ws://api.sso.sinopecscsy.com:6689/ws"`,
        //webHost:'"http://web.coupon.sinopecscsy.com"',
        // url: '"http://192.168.106.152:6688"',     //线下
        // ws: `"ws://api.sso2.100bm.cn:6689/ws"`,
        // webHost:'"http://coupon2.100bm.cn:6060"',
    }
}
