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
            this.config.ssoWebHost + this.config.jumpUrl + "?ident=sso&callback=" +
            encodeURIComponent(window.location.protocol + "//" + window.location.host + this.config.callbackUrl);
      }
    }
  }
</script>
