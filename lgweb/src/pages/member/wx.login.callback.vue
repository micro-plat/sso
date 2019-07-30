<template>
    <div>{{notice}}</div>
</template>
<script>
   import {JoinUrlParams} from '@/services/common'
  export default {
    name: 'wxcallback',
    data () {
      return {
          notice: "",
          code: "",
          state: ""
      }
    },

    created() {
        document.title = "微信登录";
        this.code = this.$route.query.code;
        this.state = this.$route.query.state;
        this.check()
    },

    methods:{
        check() {
            this.notice = "登录中...";
            this.$post("lg/login/wxcheck",{code:this.code,state: this.state})
            .then(res =>{
                this.notice = "登录成功...";
            }).catch(err => {
                switch (err.response.status) {
                    case 406:
                    case 408:
                    case 510:
                        var message = error.response.data.data; 
                        if (message && message.length > 6 && message.indexOf("error:",0) == 0) {
                            message = message.substr(6); //error:用户名或密码错误 //框架多还回一些东西
                        }
                        this.notice = message;
                        break;
                    default:
                        this.notice = "登录失败";
                }
            });
        },

        checkAndJumpLogin2() {
            var containkey = 0;
            var callbackinfo = sessionStorage.getItem("sso-bssyscallbackinfo");
            if (callbackinfo && callbackinfo.callback && callbackinfo.ident) {
              containkey = 1;
            }

            this.$post("lg/login/wxcheck",{
                    containkey:containkey, 
                    ident:(callbackinfo && callbackinfo.ident) ? callbackinfo.ident : "", 
                    code:this.code,
                    state: this.state})
            .then(res =>{
                this.notice = "登录中...";
                if (callbackinfo && callbackinfo.callback && callbackinfo.ident) {
                    window.location.href = JoinUrlParams(decodeURIComponent(callbackinfo.callback),{code:res.data})
                    return;
                }
                this.$router.push({ path: '/chose'});   
            }).catch(err => {
                var type = 0;
                switch (err.response.status) {
                    case 400:
                        type = 3;
                        break;
                    case 406:
                        type = 4;
                        break;
                    case 408:
                        type = 5;
                        break;
                    case 510:
                        type = 6;
                        break;
                    case 401:
                        type = 7;
                        break;
                    case 415:
                        type = 1;
                        break;
                    default:
                        type = 0
// 400: 用户被锁定或被禁用，暂时无法登录 //3
// 406: 微信登录过程中有些参数丢失,请正常登录 //4
// 408: 微信登录标识过期,请重新登录 //5
// 510: 调用微信失败，稍后再登录 // 6
// 500: 系统出错，等会在试 //0
// 401: 没有关注公众号 // 7
// 415: 没有相应权限，请联系管理员 //1
                }
                this.$router.push({ path: '/errpage', query: {type: type}});
            });
        }
    }
  }
</script>