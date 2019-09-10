<template>
  <div id="app">
    <login-with-up
      :copyright="copyright"
      :systemName="systemName"
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
        requireCode:false,

        errorTemplate:{
            901: "系统被锁定,不能登录",
            902: "用户被锁定,不能登录",
            903: "用户被禁用,不能登录",
            905: "用户不存在",
            906: "没有相关系统权限,不能登录",
            907: "用户名或密码错误",
            912: "请先绑定微信账户,并且关注【运维云管家】",
            913: "验证码不能为空",
            914: "验证码过期或不存在,重新发送验证码",
            915: "验证码错误"
          }
      }
    },
    components:{ 
      loginWithUp
    },

    mounted(){
      window.localStorage.removeItem("__sso_jwt__");

      document.title = "登录-能源业务中心运营管理系统";
      this.callback = this.$route.query.callback;
      this.changePwd = this.$route.query.changepwd;
      this.ident = this.$route.params.ident ? this.$route.params.ident : "";

      this.controlLoginType();
    },

    methods:{
      controlLoginType() {
        this.$post("/system/config/get", {ident: this.ident})
        .then(res => {
            this.loginTitle = "登录到【" + res.system_name + "】";
            this.requireCode = res.require_wx_code;
        })
        .catch(err => {
            this.$refs.LoginUp.showError("获取系统信息失败");
        }); 
      },

      //发送微信验证码
      getCodeCall(e){
         e.ident = this.ident;
         this.$refs.LoginUp.showError("发送验证码中...");
         this.$post("/member/sendcode", e)
          .then(res=>{
            this.$refs.LoginUp.showError("微信验证码发送成功,【运维云管家】中查看");
            this.$refs.LoginUp.countDown();
          })
          .catch(err=>{
              var msg = "登录失败,稍后再试";
              if (err.response) {
                msg = this.errorTemplate[err.response.status] || msg
              }
              this.$refs.LoginUp.showError(msg);
          })
      },

      //用户名密码登录
      loginsubmit(e){
        var req = {
          ident: this.ident,
          password: $.md5(e.password),
          username:e.username,
          wxcode:e.wxcode
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
                window.location.href = url;
                return;
              }
              this.$router.push({ path: '/choose'});
            }, 300);
          })
          .catch(err => {
              var msg = "登录失败,稍后再试";
              if (err.response) {
                msg = this.errorTemplate[err.response.status] || msg
                if (err.response.status == 907) {
                  this.$refs.LoginUp.reSendCode();
                }
              }
              this.$refs.LoginUp.showError(msg)
          });
      }

    }
  }
</script>
