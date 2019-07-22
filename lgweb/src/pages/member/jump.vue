<template>
  <div>
      <span style="text-align:center">{{notice}}</span>
  </div>
</template>

<script>
   import {JoinUrlParams} from '@/services/common'
  export default {
    name: 'jump',
    data () {
      return {
        callback:"",
        sysid:0,
        notice: "页面调转中..."
      }
    },

    mounted(){
      document.title = "用户登录";
      this.callback = this.$route.query.callback;
      this.checkAndJumpLogin();
    },

    methods:{
        checkAndJumpLogin() {
            this.$post("lg/login/check",{})
            .then(res =>{
                console.log(res.data);
                this.notice = "已登录,跳转中...";

                setTimeout(() => {
                  if (this.callback) {
                    window.location.href = JoinUrlParams(decodeURIComponent(this.callback),{code:res.data})
                    return;
                  }
                  this.$router.push({ path: '/chose',query: { code: res.data }});   
                }, 300);
            }).catch(err => {
                this.$router.push({ path: '/login', query: { callback: this.callback }});
            });
        }
    }
  }
</script>
