<template>
</template>

<script>
  export default {
    data () {
      return {
      }
    },
    mounted(){
      this.validSsoLogin();
    },
    methods:{
      validSsoLogin(){
          this.$post("sso/login/user",{code: this.$route.query.code})
            .then(res =>{
                // localStorage.setItem("userinfo", JSON.stringify({name:res.user_name, role:res.role_name}));
                this.$sso.changeRouteAfterLogin(this.$router); //将上面的东西放到里面去
            }).catch(err => {
              console.log(err);
            });
      }
    }
  }
</script>
