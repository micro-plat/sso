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
          this.$http.post("/login/user",{code: this.$route.query.code})
            .then(res =>{
                this.$sso.changeRouteAfterLogin(this.$router, res.user_name, res.role_name);
            }).catch(err => {
             if (err.response) {
                if (err.response.status == 406) {
                  this.$sso.errPage(0)
                }
              }
              console.log(err);
            });
      }
    }
  }
</script>
