<template>
  <div>
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
        ident:"",
        notice: "页面调转中..."
      }
    },

    mounted(){
      document.title = "用户登录";
      this.callback = this.$route.query.callback;
      this.ident = this.$route.query.ident;
      this.checkAndJumpLogin();
    },

    methods:{
        checkAndJumpLogin() {
            var containkey = 0;
            if (this.callback && this.ident) {
              containkey = 1;
            }
            this.$post("lg/login/check",{containkey:containkey, ident:this.ident})
            .then(res =>{
                console.log(res.data);
                this.notice = "已登录,跳转中...";

                setTimeout(() => {
                  if (this.callback && this.ident) {
                    window.location.href = JoinUrlParams(decodeURIComponent(this.callback),{code:res.data})
                    return;
                  }
                  this.$router.push({ path: '/chose'});   
                }, 300);
            }).catch(err => {
                switch (err.response.status) {
                    case 423:
                      this.$router.push({ path: '/errpage', query: {type: 2}});
                      break;
                    case 415:
                      this.$router.push({ path: '/errpage', query: {type: 1}});
                    default:
                      this.$router.push({ path: '/login', query: { callback: this.callback, ident:this.ident }});
                }
            });
        }
    }
  }
</script>
