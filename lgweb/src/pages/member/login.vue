<template>
  <div id="app">
    <login-with-up
      :copyright="copyright"
      :systemName="systemName"
      :conf="conf"
      :wxlg="wxLogin"
      :requireWxLogin="requireWxLogin"
      :requireCode="requireCode"
      :getCodeCall="getCodeCall"
      :call="loginsubmit"
      :err-msg.sync="errMsg"
      ref="loginItem">
    </login-with-up>
  </div>
</template>

<script>
  import VueCookies from 'vue-cookies'
  import loginWithUp from 'login-with-up';
  import {JoinUrlParams} from '@/services/common.js'
  import "@/services/md5.js"
  export default {
    name: 'app',
    data () {
      return {
        systemName: "能源业务中心运营管理系统",
        copyright:"四川千行你我科技有限公司Copyright© 2018 版权所有",
        conf:{loginNameType:"输入用户名",pwd:"输入密码",validateCode:"请输入微信验证码"},
        callback:"",
        changepwd:0,
        ident: "",
        errMsg:{message:""},
        requireWxLogin:false, //是否支持跳转登录
        requireCode: false //是否支持微信验证码登录
      }
    },
    components:{ 
      loginWithUp
    },

    created() {
      this.controlLoginType();
    },

    mounted(){
      document.title = "登录-能源业务中心运营管理系统";
      this.callback = this.$route.query.callback;
      this.changepwd = this.$route.query.changepwd;
      this.ident = this.$route.query.ident;
      
      VueCookies.remove("__jwt__")
    },

    methods:{

      //取配置，显示验证码登录还是扫码登录
      controlLoginType() {
        this.$post("lg/login/typeconf", {})
        .then(res => {
          this.requireWxLogin = res.requirewxlogin;
          this.requireCode = res.requirecode;
        })
        .catch(err => {

        })
      },

      // 微信调转登录
      wxLogin(){
        this.$post("lg/login/wxconf", {})
        .then(res => {
            var url = res.wxlogin_url + "?" + "appid=" + res.appid + "&state=" + res.state + "&redirect_uri=" +
                      encodeURIComponent(process.env.service.wxcallbackhost + process.env.service.wxcallbackurl) +
                      "&response_type=code&scope=snsapi_login#wechat_redirect";
                      
                      
            console.log(url);
            sessionStorage.setItem("sso-bssyscallbackinfo", JSON.stringify({callback: this.callback, changepwd: this.changepwd, ident:this.ident}));
            window.location.href = url;
        })
        .catch(err => {
          this.errMsg = {message: "系统繁忙,请稍后在试"};
        });
      },

      //用户名密码登录
      loginsubmit(e){
        var req = {
          containkey: 0,
          ident: this.ident,
          password: $.md5(e.password),
          username:e.username,
          validatecode:e.validatecode,
        }

        if (this.callback && this.ident) {
          req.containkey = 1
        }

        this.$post("lg/login/post", req)
          .then(res => {
            setTimeout(() => {

              if (this.changepwd == 1) {
                this.$router.push({ path: '/changepwd'});   
                return;
              }

              if (this.ident && this.callback) {
                window.location.href = JoinUrlParams(decodeURIComponent(this.callback),{code:res.data})
                return;
              }
              this.$router.push({ path: '/chose'});
            }, 300);
          })
          .catch(err => {
              var message = err.response.data.data; 
              if (message && message.length > 6 && message.indexOf("error:",0) == 0) {
                message = message.substr(6); //error:用户名或密码错误 //框架多还回一些东西
              }
              switch (err.response.status) {
                case 400:
                case 401:
                case 423:
                case 405:
                case 415:
                  this.errMsg = {message: message}; 
                  break;
                default:
                  this.errMsg = {message: "登录失败"};
              }
          });
      },

      //发送微信验证码
      getCodeCall(e){
         e.ident = this.ident ? this.ident : "";
         this.errMsg = {message: "发送中....."}; 

         this.$post("/lg/login/wxvalidcode", e)
          .then(res=>{
            this.errMsg = {message: "微信验证码发送成功"}; 
          })
          .catch(error=>{
            var message = err.response.data.data; 
            if (message && message.length > 6 && message.indexOf("error:",0) == 0) {
              message = message.substr(6); //error:用户名或密码错误 //框架多还回一些东西
            }
            switch(err.response.status) {
              case 401:
                this.errMsg = {message: message}; 
                break;
              default:
                this.errMsg = {message: "系统繁忙"}; 
            }
          })
      }

    }
  }
</script>
