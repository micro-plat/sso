<template>
    <div>
        <div>{{title}}</div>
        <div>{{notice}}</div>
        <div><input type="button" value="微信授权登录" @click="wxConfirm"></div>
    </div>
</template>
<script>
  export default {
    name: 'qrcode',
    data () {
      return {
          title: "能源业务中心运营管理系统",
          notice:"没有关注公众号的请先关注公众号",
          state:"" //动态生成的用户标识
      }
    },
    created() {
        document.title = "微信登录";
        this.state = this.$route.query.state;
    },
    methods:{
        wxConfirm() {
            console.log(this.state);
            
            this.$post("lg/login/wxconf", {})
            .then(res => {
                var url = res.wxlogin_url + "?" + "appid=" + res.appid + "&state=" + this.state + "&redirect_uri=" +
                        encodeURIComponent(process.env.service.wxcallbackhost + process.env.service.wxlogincallbackurl) +
                        "&response_type=code&scope=snsapi_base#wechat_redirect";
                        
                window.location.href = url;
            })
            .catch(err => {
                this.errMsg = {message: "系统繁忙,请稍后在试"};
            });
        }
    }
  }
</script>