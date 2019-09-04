<template>
  <div>
      <div>
        <span>绑定用户：</span>{{userName}}
      </div>
      <div><input type="button" value="绑定" @click="checkUserInfo" v-bind:disabled="disabled" /></div>
      <div><span>{{notice}}</span></div>
  </div>
</template>

<script>
  export default {
    name: 'bindwx',
    data () {
      return {
        notice: "",
        userName:"",
        userId:0,
        sign:"",
        timestamp:0,
        disabled:false
      }
    },
    mounted(){
      document.title = "绑定微信账号";
      this.userId = this.$route.query.userid;
      this.sign = this.$route.query.sign;
      this.userName = this.$route.query.name;
      this.timestamp = this.$route.query.timestamp;
    },

    methods:{
        checkUserInfo() {
            this.$post("/member/bind/check",{user_id:this.userId, sign:this.sign, timestamp:this.timestamp})
            .then(res =>{
                var url = res.wxlogin_url + "?" + "appid=" + res.appid + "&state=" + res.state + "&redirect_uri=" +
                        encodeURIComponent(process.env.service.wxcallbackhost + process.env.service.wxcallbackurl + "/bind") +
                        "&response_type=code&scope=snsapi_base#wechat_redirect"; 
                console.log(url);
                window.location.href = url;  

            }).catch(err => {
                switch (err.response.status) {
                    case 902:
                      this.notice = "用户被锁定"
                      break;
                    case 903:
                      this.notice = "用户被禁用"
                      break;
                    case 909:
                      this.notice = "绑定信息错误,请重新去用户系统扫码"
                      break;
                    case 910:
                      this.notice = "用户已绑定微信"
                      break;
                    default:
                      this.notice = "系统错误,稍后在试"
                }
                this.disabled = true;
            });
        }
    }
  }
</script>
