<template>
<div class="bind-container">
  <div class="top-title">绑定微信账户</div>
  <div class="user-info">
    <span class="user-photo"><img src="../../img/photo.png" /></span>
    <span class="user-name"><a class="notice-info">登录账号</a><a class="name">{{userName}}</a></span>
  </div>
  <div class="notice">您正在绑定微信账户,绑定后请关注微信公众号【运维云管家】,登录验证码会通过此公众号发送</div>
  <div class="button"><button class="btn" @click="checkUserInfo">确认</button></div>
</div>
</template>

<script>
  export default {
    name: 'bindwx',
    data () {
      return {
        userName:"",
        userId:0,
        sign:"",
        timestamp:0
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
                console.log(err);
                if (err.response.status) {
                  this.$router.push({path:"/bindnotice", query :{ type: 0, errorcode:err.response.status }});
                }
            });
        }
    }
  }
</script>

<style>
*{margin:0;padding:0;}
html,body{height:100%; }
li{	list-style:none;}
body{font-family:"黑体";	background:#f5f5f5; font-size:14px; }
a:visited {	text-decoration: none;}
a:hover {	text-decoration: none;}
a:active {	text-decoration: none;}

input[type="button"],button,input[type="text"]{-webkit-appearance: none;border-radius: 5px;outline: none;}
input:disabled,input[disabled],button:disabled,button[disabled]{ border: 1px solid #DDD;background-color: #d9d9d9;color:#999;}

.bind-container{ height:100%; background:#fff;}
.top-title{ text-align:center; font-size:18px; color:#333; font-weight:600; padding:10% 10px 10px 10px;}
.user-info{ margin:10px 10px; padding:10px;line-height:60px; border-top:1px solid #efefef;border-bottom:1px solid #efefef; background:#fff;}
.user-photo{ width:50px; height:50px; display:inline-block; float:left;}
.user-photo img{ width:100%; height:100%; border-radius:5px;}
.user-name{ display:inline-block; padding-left:10px;}
.user-name .name,.user-name .notice-info{ display:block; height:25px; line-height:25px;}
.name{ font-size:14px; color:#333;}
.notice-info{ color:#a8a8a8;}
.notice{ padding:0 20px; color:#acabab; font-size:14px; line-height:20px; }
.button{ position:absolute; width:100%; bottom:40px;text-align:center;}
.btn{ width:90%; border:0px; background:#4bc065; padding:15px 0; font-size:16px; color:#fff; font-weight:bold;}
</style>
