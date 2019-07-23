<template>
</template>

<script>
  import VueCookies from 'vue-cookies'
  export default {
    name: 'callback',
    data () {
      return {
        code : ""
      }
    },
    mounted(){
      this.code = this.$route.query.code;
      this.validSsoLogin();
    },
    methods:{
      validSsoLogin(){
          this.$post("sso/login/user",{code: this.code})
            .then(res =>{
                localStorage.setItem("userinfo", JSON.stringify({name:res.user_name, role:res.role_name}));
                this.$router.push("/");
            }).catch(err => {
              var config  = process.env.service;
              window.location.href = config.ssoWebHost + config.errPage;
            });
      }
    }
  }
</script>

<style>
  .main{
    text-align: center;
  }
</style>
