<template>
</template>

<script>

  import VueCookies from 'vue-cookies'
  export default {
    name: 'app',
    data () {
      return {
        config : process.env.service
      }
    },
    mounted(){
      VueCookies.remove("__jwt__");
      this.jumpToSso();
    },
    methods:{
      jumpToSso(){
        window.location.href = 
            this.config.ssoWebHost + this.config.jumpUrl + "?ident=sso&callback=" + encodeURIComponent(this.config.callbackUrl);
      }
    }
  }
</script>
