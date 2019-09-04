<template>
    <div>{{notice}}</div>
</template>
<script>
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
        document.title = "微信绑定";
        this.code = this.$route.query.code;
        this.state = this.$route.query.state;
        console.log(this.code);
        console.log(this.state);
        this.bind()
    },

    methods:{
        bind() {
            this.notice = "绑定中...";
            this.$post("/member/bind/save",{code:this.code, state: this.state})
            .then(res =>{
                this.notice = "绑定成功...";
            }).catch(err => {
                switch (err.response.status) {
                    case 911:
                        this.notice = "绑定超时,请重新扫码绑定";
                        break;
                    default:
                        this.notice = "绑定失败,稍后再试";
                }
            });
        }
    }
  }
</script>