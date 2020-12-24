'use strict'
module.exports = {
    NODE_ENV: '"production"',
    service: {
        wxcallbackhost:'"http://ssov3.100bm.cn"',
        // wxcallbackhost:'"http://login.sso.18jiayou.com"',
        wxcallbackurl:'"/wxcallback"',
        // url: '"http://121.196.17.172:6687"',   //预上线(跳转登录的api地址)
        
        // url: '"http://192.168.106.226:6687"'     //本地/ "http://192.168.0.103:6687
        // url: '"http://loginapi.sso.jzjy6.com:6687"',   //测试环境有dns(跳转登录的api地址)
        // url: '"http://47.97.1.98:6687"', //亿翔
        // url: '"http://47.96.91.131:6687"', //北京卓易豪斯
        //url: '"http://loginapi.sso.18jiayou.com"'
        // url: '"//login.sso.18jiayou.com/loginapi"',   //线上环境(跳转登录的api地址)(北京机房)
        url: '"http://192.168.5.94:6687"', //
        codeLabel: '"短信验证码"',
        codeHolder: '"请输入短信验证码"',
        sendBtnLabel:'"获取短信验证码"',
        showText:'"短信验证码发送成功"'
    //     codeLabel: '"微信验证码"',
    //     codeHolder: '"请输入微信验证码"',
    //     sendBtnLabel:'"获取微信验证码"'
    // showText:'"微信验证码发送成功,【运维云管家】中查看"'
     }
}