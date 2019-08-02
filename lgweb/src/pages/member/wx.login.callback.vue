<template>
    <div>{{notice}}</div>
</template>
<script>
   import {JoinUrlParams} from '@/services/common'
   import {trimError} from "@/services/utils"
  export default {
    name: 'wxcallback',
    data () {
      return {
          notice: "",
          code: "",
          state: "",
          type:"",
      }
    },

    created() {
        document.title = "微信登录";
        this.code = this.$route.query.code;
        this.state = this.$route.query.state;
        this.type = this.$route.params.type;
        this.check()
    },

    methods:{
        check() {
            if (this.type != "bind") {
                //1: 登录回调
                this.notice = "登录中...";
                this.$post("lg/login/wxcheck",{code:this.code,state: this.state})
                .then(res =>{
                    this.notice = "登录成功...";
                }).catch(err => {
                    switch (err.response.status) {
                        case 401:
                            this.$router.push({ path: '/wxbind'}); 
                            break;
                        case 415:
                        case 406:
                        case 408:
                        case 510:
                            this.notice = trimError(err);
                            break;
                        default:
                            this.notice = "登录失败";
                    }
                });
            } else {
                //绑定回调

            }
            
        }
    }
  }
</script>