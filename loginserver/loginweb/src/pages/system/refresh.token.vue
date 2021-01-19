<template>
</template>

<script>
  export default {
    name: 'refresh',
    data () {
      return {
      }
    },

    mounted(){
        this.refreshToken();
    },
    methods:{
        refreshToken() {
            this.$http.post("/loginweb/member/refresh", {})
        }
        
    }
  }
</script>
<style scoped>

</style>