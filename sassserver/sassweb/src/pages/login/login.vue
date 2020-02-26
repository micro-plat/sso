<template>
  <div id="app">
    <login-with-up
      :copyright="copyright"
      :systemName="systemName"
      :requireCode="requireCode"
      :loginCallBack="loginsubmit"

      :loginTitle="loginTitle"
      :loginNameLabel="loginNameLabel"
      :loginNameHolder="loginNameHolder"
      :loginPwdLabel="loginPwdLabel"
      :loginPwdHolder="loginPwdHolder"
      ref="LoginUp">
    </login-with-up>
  </div>
  
</template>

<script>
  import VueCookies from 'vue-cookies'
  import loginWithUp from 'login-with-up';
  import "@/services/md5.js"

  export default {
    name: 'app',
    data () {
      return {
        systemName: "能源业务中心运营管理系统",
        copyright:"四川千行你我科技有限公司Copyright©" + new Date().getFullYear() +" 版权所有",

        loginTitle:"用户登录",
        loginNameLabel:"手机号",
        loginNameHolder:"请输入手机号",
        loginPwdLabel:"密码",
        loginPwdHolder:"请输入密码",
        //codeLabel:"微信验证码",
        //codeHolder:"请输入微信验证码",
        //sendBtnLabel:"获取微信验证码",
        requireCode:false,

        errorTemplate:{
            901: "系统被锁定,不能登录",
            902: "用户被锁定,不能登录",
            903: "用户被禁用,不能登录",
            905: "用户不存在",
            906: "没有相关系统权限,不能登录",
            907: "手机号或密码错误",
            913: "验证码不能为空",
            914: "验证码过期或不存在,重新发送验证码",
            915: "验证码错误或者过期"
          }
      }
    },
    components:{ 
      loginWithUp
    },

    mounted(){
      window.localStorage.removeItem("__jwt__");
      document.title = "登录-能源业务中心运营管理系统";
    },

    methods:{
      //用户名密码登录
      loginsubmit(e){
        var req = {
          password: $.md5(e.password),
          mobile:e.username,
          verify_code: "123",
        }
        this.$http.post("/sso/login/verify", req)
          .then(res => {
            localStorage.setItem("userinfo", JSON.stringify({name:res.user_name, role:res.role_name}));
            return this.$router.push({ path: '/'});
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
<style scoped>

</style>
