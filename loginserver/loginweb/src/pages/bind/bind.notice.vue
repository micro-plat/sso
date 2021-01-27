<template>
<div class="notice-container">
    <div class="bind-info">
        <div class="info-ico">
            <i :class="type==1? 'iconfont icon-success':'iconfont icon-fail'"></i>
            <div v-if="type==1" class="info-text">绑定成功</div>
            <div v-if="type==0" class="info-text">绑定失败</div>
            <div :class="type==1? 'info-notice': 'info-notice color-black'">{{msg}}</div>
        </div>
    </div>
    <div v-if="type==1 && wechatOfficalAccoutImage!=''"  class="cloud-manager">
      <img :src="wechatOfficalAccoutImage" />
    </div>
</div>
</template>
<script>
  export default {
    name: 'bindnotice',
    data () {
      return {
          type:1, //1表示成功, 0表示失败
          errorCode: 0, //错误码 
          msg: "",
          wechatOfficalAccoutImage:"",
          errorTemplate:{
            902: "用户被锁定",
            903: "用户被禁用",
            909: "绑定信息错误,请重新去用户系统扫码",
            910: "账号已绑定微信",
            911: "绑定超时,请重新扫码绑定",
            916: "二维码过期,请联系管理员重新生成",
            917: "一个微信只能绑定一个账号"
          }
      }
    },
    created() {
        document.title = "微信绑定";
        this.type = this.$route.query.type;
        this.errorCode = this.$route.query.errorcode;
        this.showMsg();
    },
    methods:{
        showMsg(){
            this.wechatOfficalAccoutImage = this.$env.conf.system.wechatOfficalAccoutImage
            if (this.type == 1) {
                this.msg = "已绑定微信账户,请关注【运维云管家】";
                return
            }
            console.log(this.msg);
            this.msg = this.errorTemplate[this.errorCode] || "系统错误,稍后再试";
        }
    }
  }
</script>
<style scoped>
@font-face {font-family: "iconfont";
  src: url('//at.alicdn.com/t/font_1387944_tqpve63qm3.eot?t=1567650441639'); /* IE9 */
  src: url('//at.alicdn.com/t/font_1387944_tqpve63qm3.eot?t=1567650441639#iefix') format('embedded-opentype'), /* IE6-IE8 */
  url('data:application/x-font-woff2;charset=utf-8;base64,d09GMgABAAAAAAMcAAsAAAAAByAAAALNAAEAAAAAAAAAAAAAAAAAAAAAAAAAAAAAHEIGVgCDBgqCBIF3ATYCJAMMCwgABCAFhG0HOhtBBsgekiSCQN1RGAdXZC/Jogii/djP3h0mksSSJ5JoIpJohEKEBJHWKM1Csf9LeP+7pimRGuj/fFkOWRXpmgNscxmScGzUhJpwmyOwCI7+55pXVLmgcQWOPAp4+M+a1lzNowlc+sJGh3sz4HL6K/F5luWy5rCNYwImGNDeGAUUaIGkgXDD2E0ED/A2gUajvAyHA2MxZFeARYG4E7gysheiiuIU6oXahp1FvLCoT0/pAPHMPx//0UrYSWoysPL4tr+Kur6j2zbTUgUoRwmy5dxw9pGxDCnETaP3mklml5k0RmpzXIvUKkJaKvFfwaWxvvagf7xE1ALFzWBezyW+Iywg+Gn8UAIZ1KLgCHK/wCuQpOZhd/PskRvzvHgQJH0/P4GdqV7GNsyDlrhKbGlHS7lIblt7kM9nj54x7xUPPkn/ne5tkBu9Xb5E38wdrfRMIJO79omfn3kfiqiYKLMqW3Q3PZlezEFzYgQCeemLACDOoRmAKLPzsstE39Dzcb/e9PTjd558tZ04ZqP/ekoXhOqRTX3XT/+O7ff6U5Nr+7a7Cviajgw+pwIzyg+Q+lvyBb9YNrCrxApgyyWUYmnqz+YnuhO62A2oEBDsdDNd3zO4mFCvJ0NSZwRZvWmysJZRo8kaatXbQaMlg/ubdCFMojRg0bqG0G4PSas3ZO0OycK6Ro1ez6jVHiY0uoyGE5vMhuDsJbhMcRUq0yAwmsoiB3WK0gjmxuUyySs4O4GJLiQgVCyXC4NYxWSJNfoEF6aUBZZoCgxo92FZ1qBGNBEz1M9TWmsJlAZNT/IzmoKcDhBYGYVVgYppQMDQqFg6mHJW3h+BccbJykhH3Ug4ASN0wvAgxC84gBzUqoPqHuUR3QROGEWxAIvQKMCA5sZkmNWAWvM0Ecag/PgJ6ZoWAaMfO1Tr316v/ECJ0wS4PEeKHEXtbVmQQ6xNH2cYrOsIAQAAAA==') format('woff2'),
  url('//at.alicdn.com/t/font_1387944_tqpve63qm3.woff?t=1567650441639') format('woff'),
  url('//at.alicdn.com/t/font_1387944_tqpve63qm3.ttf?t=1567650441639') format('truetype'), /* chrome, firefox, opera, Safari, Android, iOS 4.2+ */
  url('//at.alicdn.com/t/font_1387944_tqpve63qm3.svg?t=1567650441639#iconfont') format('svg'); /* iOS 4.1- */
}

.iconfont {
  font-family: "iconfont" !important;
  font-size: 16px;
  font-style: normal;
  -webkit-font-smoothing: antialiased;
  -moz-osx-font-smoothing: grayscale;
}

.icon-fail:before {
  content: "\e6d9";
}

.icon-success:before {
  content: "\e73c";
}
.color-black{ color:#000 !important;}

.notice-continer{ height:100%; background:#fff;}
.bind-info{ padding:10% 10px 10px 10px;}
.info-ico{ text-align:center; padding:10px;}
.info-ico i{ font-size:54px; color:#4bc065;}
.icon-fail{ color:#F00 !important;}
.info-text{ padding:10px; font-size:20px; color:#333; text-align:center;}
.info-notice{ padding:15px 10px 10px 10px; color:#aeaeae; text-align:center; line-height:20px;}
.cloud-manager {text-align: center}
</style>