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
  import {JoinUrlParams,guid} from '@/services/common.js'

  import "@/services/qrcode.min.js"
  import "@/services/md5.js"

  export default {
    name: 'app',
    data () {
      return {
        systemName: "",
        bgImageUrl: this.$env.conf.system.staticImageUrl ,//"http://images.yxtx888.net/sso/background.jpg",
        copyright: (this.$env.conf.system.companyRight||"") + "Copyright©" + new Date().getFullYear() +"版权所有",//"北京卓易豪斯科技有限公司Copyright©" + new Date().getFullYear() +"版权所有" ,
        copyrightcode: this.$env.conf.system.companyRightCode ,//"蜀ICP备20003360号",
        requireCode: this.$env.conf.system.requireCode||true,

        returnURL:"",
        ident: "",

        loginTitle:"用户登录",
        loginNameLabel:"用户名",
        loginNameHolder:"请输入用户名",
        loginPwdLabel:"密码",
        loginPwdHolder:"请输入用户密码",

        codeLabel:this.$env.conf.system.codeLabel,
        codeHolder:this.$env.conf.system.codeHolder,
        sendBtnLabel:this.$env.conf.system.sendBtnLabel,


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
      this.$http.clearAuthorization();
      var keys = this.$cookies.keys();
      for(var i in keys){
          this.$cookies.remove(keys[i]);
      }
            
      var logoutURL = this.$route.query.logouturl;
      if(logoutURL){
         window.location.href = logoutURL
         return
      }
      document.title = "登录";
      this.returnURL = this.$route.query.returnurl;
      this.ident = this.$route.params.ident || "";

      this.controlLoginType();
    },

    methods:{
      controlLoginType() {
        this.$http.post("/loginweb/system/config/get", {ident: this.ident})
        .then(res => {
            console.log("---:",res)
            this.loginTitle = "登录到【" + res.system_name + "】";
            this.requireCode = res.require_valid_code;
        })
        .catch(err => {
            this.$refs.LoginUp.showError("获取系统信息失败");
        }); 
      },

      //发送微信验证码
      getCodeCall(e){
         this.$refs.LoginUp.showError("发送验证码中...");
         this.$http.post("/loginweb/member/sendcode", {ident: this.ident, username: e.username,guid:guid()})
          .then(res=>{
            this.$refs.LoginUp.showError(this.$env.conf.system.showText);
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
          validcode:e.wxcode
        }
        var that = this;
        this.$http.post("/loginweb/member/login", req)
          .then(res => {
            setTimeout(() => { 
              var parmscode={code:res.code} 
              if(that.returnURL){
                window.location.href = JoinUrlParams(decodeURIComponent(that.returnURL),parmscode);
                return 
              }
              if (that.ident && res.callback) {
                var url = JoinUrlParams(decodeURIComponent(res.callback),parmscode);
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
