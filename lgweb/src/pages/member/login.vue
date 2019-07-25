<template>
  <div id="app">
    <login-with-up
      :copyright="copyright"
      :systemName="systemName"
      :conf="conf"
      :wxlg="wxLogin"
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
        conf:{loginNameType:"输入用户名",pwd:"输入密码"},
        callback:"",
        changepwd:0,
        ident: "",
        errMsg:{message:""}
      }
    },
    components:{ 
      loginWithUp
    },

    mounted(){
      document.title = "登录-能源业务中心运营管理系统";
      this.callback = this.$route.query.callback;
      this.changepwd = this.$route.query.changepwd;
      this.ident = this.$route.query.ident;
      
      VueCookies.remove("__jwt__")
    },

    methods:{
      
      wxLogin(){
        this.$post("lg/login/wxconf", {})
        .then(res => {
            var url = res.wxlogin_url + "?" + "appid=" + res.appid + "&state=" + res.state + "&redirect_uri=" + "https%3A%2F%2Fpassport.yhd.com%2Fwechat%2Fcallback.do" + "&response_type=code&scope=snsapi_login#wechat_redirect"
            console.log(url);
            sessionStorage.setItem("sso-bssyscallbackinfo", JSON.stringify({callback: this.callback, changepwd: this.changepwd, ident:this.ident}));
            window.location.href = url;
        })
        .catch(err => {
          this.errMsg = {message: "系统繁忙,请稍后在试"};
        });
      },

      loginsubmit(e){
        var req = {
          containkey: 0,
          ident: this.ident,
          password: $.md5(e.password),
          username:e.username
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
      }
    }
  }
</script>
