<template>
  <div>
      <span>{{notice}}</span>
  </div>
</template>

<script>
   import {JoinUrlParams} from '@/services/common'
   import {jumpLogin} from '@/services/utils'
  export default {
    name: 'jump',
    data () {
      return {
        callback:"", //这个主要是为了本地测试用到，线上用数据库的配置地址
        ident:"",
        notice: "页面调转中..."
      }
    },

    mounted(){
      document.title = "用户登录";
      this.callback = this.$route.query.callback;
      this.ident = this.$route.params.ident ? this.$route.params.ident : "";
      this.checkAndJumpLogin();
    },

    methods:{
        checkAndJumpLogin() {
            this.$post("/mgrweb/login/check",{ident:this.ident})
            .then(res =>{
                this.notice = "已登录,跳转中..."; 
                if (this.callback) { //本地测试走这条线
                  window.location.href = JoinUrlParams(decodeURIComponent(this.callback),{code:res.code})
                  return;
                }
                if (this.ident && res.callback) {
                    window.location.href = JoinUrlParams(decodeURIComponent(res.callback),{code:res.code})
                    return;
                }
                this.$router.push({ path: '/choose'});   
            }).catch(err => {
                if (err.response.status) {
                  if (err.response.status == 403){
                    this.$router.push({ path: jumpLogin(this.ident), query:{callback: this.callback}});
                    return;
                  }
                  this.$router.push({ path: '/errpage', query: {type: err.response.status}});
                }
            });
        }
    }
  }
</script>
<style>
  body{
    margin: 0;
  }
</style>
