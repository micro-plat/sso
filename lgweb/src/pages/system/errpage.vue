<template>
  <div class="main">
    <div>{{msg}}</div>
    <div><button class="btn btn-sm btn-primary" @click="signOut">登录</button></div>
  </div>
</template>

<script>
   import {JoinUrlParams} from '@/services/common'
  export default {
    name: 'errpage',
    data () {
      return {
          msg: "",
          errorType: 0
      }
    },

    mounted(){
      document.title = "登录出错-能源业务中心运营管理系统";
      if (this.$route.query.type) {
          this.errorType = this.$route.query.type
      }
      this.showMsg();
    },

    methods:{
        showMsg() {
           switch (this.errorType) {
               case 0 :
                   this.msg = '登录出错了';
                   break;
                case 1:
                    this.msg = '没有相关的权限,请联系管理员';
                    break;
           }
        },
        signOut() {
            this.$router.push({ path: '/login'});
        }
    }
  }
</script>

<style>

</style>
