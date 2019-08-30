<template>
  <div id="app">
    <login-with-up
      :copyright="copyright"
      :systemName="systemName"
      :requireWxLogin="requireWxLogin"
      :requireCode="requireCode"
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
        copyright:"四川千行你我科技有限公司Copyright©" + new Date().getFullYear() +" 版权所有",
        callback:"",
        changePwd:0,
        ident: "",

        loginTitle:"用户登录",
        loginNameLabel:"用户名",
        loginNameHolder:"请输入用户名",
        loginPwdLabel:"密码",
        loginPwdHolder:"请输入用户密码",
        codeLabel:"微信验证码",
        codeHolder:"请输入微信验证码",
        sendBtnLabel:"获取微信验证码",
        requireWxLogin:false,
        requireCode:false
      }
    },
    components:{ 
      loginWithUp
    },

    mounted(){
      //VueCookies.remove("__sso_jwt__");
      window.localStorage.removeItem("__sso_jwt__");

      document.title = "登录-能源业务中心运营管理系统";
      this.callback = this.$route.query.callback;
      this.changePwd = this.$route.query.changepwd;
      this.ident = this.$route.params.ident ? this.$route.params.ident : "";

      this.controlLoginType();
    },

    methods:{
      controlLoginType() {
        if (! this.ident){
          return;
        }
        this.$post("/system/get", {ident: this.ident})
        .then(res => {
            this.loginTitle = "登录到【" + res.name + "】";
        })
        .catch(err => {
            this.$refs.LoginUp.showError("获取系统信息失败");
        }); 
      },

      //用户名密码登录
      loginsubmit(e){
        var req = {
          ident: this.ident,
          password: $.md5(e.password),
          username:e.username
        }

        this.$post("/member/login", req)
          .then(res => {
            setTimeout(() => {

              if (this.changePwd == 1) {
                this.$router.push({ path: '/changepwd'});   
                return;
              }
              if (this.callback) {
                window.location.href = JoinUrlParams(decodeURIComponent(this.callback),{code:res.code});
                return;
              }
              if (this.ident && res.callback) {
                var url = JoinUrlParams(decodeURIComponent(res.callback),{code:res.code});
                console.log("denglutiaozhuan: url: ", url)
                window.location.href = url;
                return;
              }
              this.$router.push({ path: '/choose'});
            }, 300);
          })
          .catch(err => {
              var msg = "登录失败";
              switch (err.response.status) {
                case 901:
                  msg = "系统被锁定,不能登录"
                  break;
                case 902:
                  msg = "用户被锁定,不能登录"
                  break;
                case 903:
                  msg = "用户被禁用,不能登录";
                  break;
                case 906:
                  msg = "没有相关系统权限,不能登录";
                  break;
                case 907:
                  msg = "用户名或密码错误";
                  break;    
                default:
                  msg = "登录失败,稍后再试";
              }
              this.$refs.LoginUp.showError(msg)
          });
      }

    }
  }
</script>
