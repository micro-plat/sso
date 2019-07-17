<template>
  <div class="main">{{msg}}</div>
</template>

<script>
  import VueCookies from 'vue-cookies'
  export default {
    name: 'callback',
    data () {
      return {
        code : "",
        msg: ""
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
                this.$router.push("/");
            }).catch(err => {
              this.msg = "登录出错";
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
