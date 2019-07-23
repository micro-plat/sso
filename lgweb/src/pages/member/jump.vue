<template>
  <div style="text-align:center">
      <span>{{notice}}</span>
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
            var containkey = 0;
            if (this.callback) {
              containkey = 1;
            }
            this.$post("lg/login/check",{containkey:containkey})
            .then(res =>{
                console.log(res.data);
                this.notice = "已登录,跳转中...";

                setTimeout(() => {
                  if (this.callback) {
                    window.location.href = JoinUrlParams(decodeURIComponent(this.callback),{code:res.data})
                    return;
                  }
                  this.$router.push({ path: '/chose'});   
                }, 300);
            }).catch(err => {
                this.$router.push({ path: '/login', query: { callback: this.callback }});
            });
        }
    }
  }
</script>
