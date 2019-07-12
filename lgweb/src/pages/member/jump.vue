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
      //encodeURIComponent decodeURIComponent

      document.title = "用户登录";
      this.callback = this.$route.query.callback;
      this.sysid = this.$route.query.sysid;

      this.checkAndJumpLogin();
    },

    methods:{
        checkAndJumpLogin() {
            this.$post("lg/login/check",{sysid: this.sysid})
            .then(res =>{
                console.log(res.data);
                this.notice = "已登录...";
                setTimeout(() => {
                     window.location.href = JoinUrlParams(decodeURIComponent(this.callback),{key:res.data})
                }, 300);
            }).catch(err => {
                this.$router.push({ path: '/login', query: { callback: this.callback, sysid: this.sysid }});
            });
        }
    }
  }
</script>
