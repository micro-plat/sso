<template>
  <div id="app">
    <login-with-up
      :copyright="copyright"
      :systemName="systemName"
      :wxlg="wxLogin"
      :requireWxLogin="requireWxLogin"
      :requireCode="requireCode"
      :sendCode="getCodeCall"
      :loginCallBack="loginsubmit"

      :loginTitle="loginTitle"
      :loginNameLabel="loginNameLabel"
      :loginNameHolder="loginNameHolder"
      :loginPwdLabel="loginPwdLabel"
      :loginPwdHolder="loginPwdHolder"
      :codeLabel="codeLabel"
      :codeHolder="codeHolder"
      :sendBtnLabel="sendBtnLabel"
      ref="LoginUp">
    </login-with-up>
    <div v-if="requireWxLogin">
      <div id="qrcodeTable"></div>
      <input type="button" @click="generateQrCode" value="二维码" />
    </div>
  </div>
  
</template>

<script>
  import VueCookies from 'vue-cookies'
  import loginWithUp from 'login-with-up';
  import {JoinUrlParams} from '@/services/common.js'

  import "@/services/qrcode.min.js"
  import "@/services/md5.js"
  import {trimError} from "@/services/utils"

  export default {
    name: 'app',
    data () {
      return {
        systemName: "能源业务中心运营管理系统",
        copyright:"四川千行你我科技有限公司Copyright© 2018 版权所有",
        callback:"",
        changepwd:0,
        ident: "",

        loginTitle:"用户登录",
        loginNameLabel:"用户名",
        loginNameHolder:"请输入用户名",
        loginPwdLabel:"密码",
        loginPwdHolder:"请输入用户密码",
        codeLabel:"微信验证码",
        codeHolder:"请输入微信验证码",
        sendBtnLabel:"获取微信验证码",

      }
    },
    components:{ 
      loginWithUp
    },

    mounted(){
      VueCookies.remove("__sso_jwt__");

      document.title = "登录-能源业务中心运营管理系统";
      this.callback = this.$route.query.callback;
      this.changepwd = this.$route.query.changepwd;
      this.ident = this.$route.params.ident ? this.$route.params.ident : "";

      this.controlLoginType();
    },

    methods:{
      controlLoginType() {
        if (! this.ident){
          return;
        }
        this.$post("/system", {ident: this.ident})
        .then(res => {
          if (res.sysname) {
            this.loginTitle = "登录到【" + res.sysname + "】";
          }
        })
      },

      //用户名密码登录
      loginsubmit(e){
        var req = {
          ident: this.ident,
          password: $.md5(e.password),
          username:e.username,
          validatecode:e.validatecode,
        }

        this.$post("/login", req)
          .then(res => {
            setTimeout(() => {

              if (this.changepwd == 1) {
                this.$router.push({ path: '/changepwd'});   
                return;
              }
              if (this.callback) {
                window.location.href = JoinUrlParams(decodeURIComponent(this.callback),{code:res.code});
                return;
              }
              if (this.ident && res.callback) {
                window.location.href = JoinUrlParams(decodeURIComponent(res.callback),{code:res.code});
                return;
              }
              this.$router.push({ path: '/chose'});
            }, 300);
          })
          .catch(err => {
              switch (err.response.status) {
                case 400:
                case 401:
                case 423:
                case 405:
                case 415:
                  this.$refs.LoginUp.showError(trimError(err))
                  break;
                default:
                  this.$refs.LoginUp.showError("登录失败")
              }
          });
      }

    }
  }
</script>
