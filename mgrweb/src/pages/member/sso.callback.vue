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
                this.$router.push("/");
            }).catch(err => {
            window.location.href = 
              "http://192.168.5.78:8081" + "/jump" + "?callback=" + encodeURIComponent("http://192.168.5.78:8080/ssocallback") + "&sysid=0";
            });
      }
    }
  }
</script>
