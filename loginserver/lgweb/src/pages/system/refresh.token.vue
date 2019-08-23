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
            this.$post("/member/refresh", {})
        }
        
    }
  }
</script>
