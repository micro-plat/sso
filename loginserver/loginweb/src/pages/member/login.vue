<template>
  <div id="app">
    <login-with-up
      :requireOper="false"
      :bgImageUrl="bgImageUrl"
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
      :copyrightcode="copyrightcode"
      :copyRightCallBack="copyRight"
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
        systemName: "",
        bgImageUrl: "http://images.yxtx888.net/sso/background.jpg",
        copyright:  "北京卓易豪斯科技有限公司Copyright©" + new Date().getFullYear() +"版权所有" ,
        copyrightcode:"蜀ICP备20003360号",
        callback:"",
        changePwd:0,
        ident: "",

        loginTitle:"用户登录",
        loginNameLabel:"用户名",
        loginNameHolder:"请输入用户名",
        loginPwdLabel:"密码",
        loginPwdHolder:"请输入用户密码",
        // codeLabel:"短信验证码",
        // codeHolder:"请输入短信验证码",
        // sendBtnLabel:"获取短信验证码",
        codeLabel:process.env.service.codeLabel,
        codeHolder:process.env.service.codeHolder,
        sendBtnLabel:process.env.service.sendBtnLabel,
        requireCode:true,

        errorTemplate:{
            901: "系统被锁定,不能登录",
            902: "用户被锁定,不能登录",
            903: "用户被禁用,不能登录",
            905: "用户不存在",
            906: "没有相关系统权限,不能登录",
            907: "用户名或密码错误",
            912: "请先绑定手机号",
            913: "验证码不能为空",
            914: "验证码过期或不存在,重新发送验证码",
            915: "验证码错误",
            922: "密码错误,还有5次机会",
            923: "密码错误,还有4次机会",
            924: "密码错误,还有3次机会",
            925: "密码错误,还有2次机会",
            926: "密码错误,还有1次机会"
          }
      }
    },
    components:{ 
      loginWithUp
    },

    mounted(){
      
      window.localStorage.removeItem("__sso_jwt__");

      document.title = "登录";
      this.callback = this.$route.query.callback;
      this.changePwd = this.$route.query.changepwd;
      this.ident = this.$route.params.ident ? this.$route.params.ident : "";

      this.controlLoginType();
    },

    methods:{
      controlLoginType() {
        this.$post("/mgrweb/system/config/get", {ident: this.ident})
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
        console.log(process.env.service,"process.env.service")
         e.ident = this.ident;
        //  e.ident = "sso";
         this.$refs.LoginUp.showError("发送验证码中...");
         this.$post("/mgrweb/member/sendcode", e)
          .then(res=>{
            this.$refs.LoginUp.showError(process.env.service.showText);
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
        this.$post("/mgrweb/member/login", req)
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
      },
      copyRight(){
        window.open("http://www.beian.miit.gov.cn")
      }

    }
  }
</script>
<style scoped>

</style>
