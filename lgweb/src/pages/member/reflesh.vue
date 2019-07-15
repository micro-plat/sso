<template>
</template>

<script>
   import {JoinUrlParams} from '@/services/common'
  export default {
    name: 'reflesh',
    data () {
      return {
      }
    },

    mounted(){
      this.refleshSsoLogin();
    },

    methods:{
        refleshSsoLogin() {
            this.$post("lg/login/refresh")
            .then(res =>{
                console.log("reflesh ok")
            })
        }
    }
  }
</script>
