<template>
  <div>
      <span>{{notice}}</span>
  </div>
</template>

<script>
   import {JoinUrlParams,jumpLogin} from '@/services/common'
  export default {
    name: 'jump',
    data () {
      return {
        returnurl:"", 
        ident:"",
        notice: "页面调转中..."
      }
    },

    mounted(){
      document.title = "用户登录";
      this.returnurl = this.$route.query.returnurl;
      this.ident = this.$route.params.ident || "";
      this.checkAndJumpLogin();
    },

    methods:{
        checkAndJumpLogin() {
            this.$http.post("/loginweb/login/check",{ident:this.ident})
            .then(res =>{
                this.notice = "已登录,跳转中..."; 
                if (this.ident && res.callback) {
                  console.log("login/check:",this.ident,res)
                  window.location.href = JoinUrlParams(decodeURIComponent(this.returnurl),{code:res.code})
                  return;
                }
                this.$router.push({ path: '/choose'});   
            }).catch(err => {
                if (err.response) {
                  if (err.response.status == 403){
                    console.log("login/check.catch:",this.ident,err)
                    this.$router.push({ path: jumpLogin(this.ident), query:{returnurl: this.returnurl}});
                    return;
                  }
                  this.$router.push({ path: '/errpage', query: {type: err.response.status||500}});
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
