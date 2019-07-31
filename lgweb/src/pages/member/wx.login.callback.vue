<template>
    <div>{{notice}}</div>
</template>
<script>
   import {JoinUrlParams} from '@/services/common'
  export default {
    name: 'wxcallback',
    data () {
      return {
          notice: "",
          code: "",
          state: ""
      }
    },

    created() {
        document.title = "微信登录";
        this.code = this.$route.query.code;
        this.state = this.$route.query.state;
        this.check()
    },

    methods:{
        check() {
            this.notice = "登录中...";
            this.$post("lg/login/wxcheck",{code:this.code,state: this.state})
            .then(res =>{
                this.notice = "登录成功...";
            }).catch(err => {
                switch (err.response.status) {
                    case 415:
                    case 406:
                    case 408:
                    case 510:
                        var message = err.response.data.data; 
                        if (message && message.length > 6 && message.indexOf("error:",0) == 0) {
                            message = message.substr(6); //error:用户名或密码错误 //框架多还回一些东西
                        }
                        this.notice = message;
                        break;
                    default:
                        this.notice = "登录失败";
                }
            });
        }
    }
  }
</script>