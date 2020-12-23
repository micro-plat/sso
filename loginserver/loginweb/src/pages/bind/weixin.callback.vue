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
        this.bind()
    },

    methods:{
        bind() {
            this.notice = "绑定中...";
            this.$post("/member/bind/save",{code:this.code, state: this.state})
            .then(res =>{
                this.$router.push({path:"/bindnotice", query :{ type: 1, errorcode:0 }});
            }).catch(err => {
                console.log(err);
                if (err.response.status) {
                    this.$router.push({path:"/bindnotice", query :{ type: 0, errorcode:err.response.status }});
                }
            });
        }
    }
  }
</script>
<style scoped>

</style>