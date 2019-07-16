<template>
</template>

<script>

  import VueCookies from 'vue-cookies'
  export default {
    name: 'app',
    data () {
      return {
        //config : process.env.service
      }
    },
    mounted(){
      VueCookies.remove("__jwt__");
      this.jumpToSso();
    },
    methods:{
      jumpToSso(){
        //console.log(this.config);
        /*window.location.href = 
            service.ssoHost + service.jumpUrl + "?callback=" + service.callbackUrl + "&sysid=0";
        */
        window.location.href = 
            "http://192.168.5.78:8081" + "/jump" + "?callback=" + encodeURIComponent("http://192.168.5.78:8080/ssocallback") + "&sysid=0";
      }
    }
  }
</script>
