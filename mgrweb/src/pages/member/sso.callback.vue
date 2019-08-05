<template>
</template>

<script>
  import VueCookies from 'vue-cookies'
  import {changeRouteAfterLogin} from '@/services/sso.login.js'
  
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
                this.$sso.changeRouteAfterLogin(this.$router);
            }).catch(err => {
              console.log(err);
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
