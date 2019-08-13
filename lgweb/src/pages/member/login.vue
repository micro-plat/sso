<template>
  <div id="app">
    <login-with-up
      :copyright="copyright"
      :systemName="systemName"
      :wxlg="wxLogin"
      :requireWxLogin="requireWxLogin"
      :requireCode="requireCode"
      :sendCode="getCodeCall"
      :loginCallBack="loginsubmit"

      :loginNameLabel="loginNameLabel"
      :loginNameHolder="loginNameHolder"
      :loginPwdLabel="loginPwdLabel"
      :loginPwdHolder="loginPwdHolder"
      :codeLabel="codeLabel"
      :codeHolder="codeHolder"
      :sendBtnLabel="sendBtnLabel"
      ref="LoginUp">
    </login-with-up>
    <div v-if="requireWxLogin">
      <div id="qrcodeTable"></div>
      <input type="button" @click="generateQrCode" value="二维码" />
    </div>
  </div>
  
</template>

<script>
  import VueCookies from 'vue-cookies'
  import loginWithUp from 'login-with-up';
  import {JoinUrlParams} from '@/services/common.js'

  import "@/services/qrcode.min.js"
  import "@/services/md5.js"
  import {trimError} from "@/services/utils"

  export default {
    name: 'app',
    data () {
      return {
        systemName: "能源业务中心运营管理系统",
        copyright:"四川千行你我科技有限公司Copyright© 2018 版权所有",
        callback:"",
        changepwd:0,
        ident: "",
        requireWxLogin:false, //是否支持跳转登录
        requireCode: false, //是否支持微信验证码登录

        loginNameLabel:"用户名",
        loginNameHolder:"请输入用户名",
        loginPwdLabel:"密码",
        loginPwdHolder:"请输入用户密码",
        codeLabel:"微信验证码",
        codeHolder:"请输入微信验证码",
        sendBtnLabel:"获取微信验证码",

        stateCode : "", //动态为用户生成标识,用于扫码登录 (在table切换时要改这个值，相当与这个值会随着table切换而变化)
        //todo 还要改
      }
    },
    components:{ 
      loginWithUp
    },

    created() {
      this.controlLoginType();
    },

    mounted(){
      document.title = "登录-能源业务中心运营管理系统";
      this.callback = this.$route.query.callback;
      this.changepwd = this.$route.query.changepwd;
      this.ident = this.$route.params.ident ? this.$route.params.ident : "";
      
      VueCookies.remove("__sso_jwt__");
    },

    methods:{
      //取配置，显示验证码登录还是扫码登录
      controlLoginType() {
        this.$post("lg/login/typeconf", {})
        .then(res => {
          this.requireWxLogin = res.requirewxlogin;
          this.requireCode = res.requirecode;
        })
      },

      //用户名密码登录
      loginsubmit(e){
        var req = {
          ident: this.ident,
          password: $.md5(e.password),
          username:e.username,
          validatecode:e.validatecode,
        }

        this.$post("lg/login/post", req)
          .then(res => {
            setTimeout(() => {

              if (this.changepwd == 1) {
                this.$router.push({ path: '/changepwd'});   
                return;
              }
              if (this.callback) {
                window.location.href = JoinUrlParams(decodeURIComponent(this.callback),{code:res.code});
                return;
              }
              if (this.ident && res.callback) {
                window.location.href = JoinUrlParams(decodeURIComponent(res.callback),{code:res.code});
                return;
              }
              this.$router.push({ path: '/chose'});
            }, 300);
          })
          .catch(err => {
              switch (err.response.status) {
                case 400:
                case 401:
                case 423:
                case 405:
                case 415:
                  this.$refs.LoginUp.showError(trimError(err))
                  break;
                default:
                  this.$refs.LoginUp.showError("登录失败")
              }
          });
      },

      //发送微信验证码
      getCodeCall(e){
         e.ident = this.ident ? this.ident : "";
         this.$refs.LoginUp.showError("发送中...");

         this.$post("/lg/login/wxvalidcode", e)
          .then(res=>{
            this.$refs.LoginUp.showError("微信验证码发送成功");
            this.$refs.LoginUp.countDown(this.sendBtnLabel);
          })
          .catch(error=>{
            switch(error.response.status) {
              case 401:
              case 400:
                this.$refs.LoginUp.showError(trimError(error));
                break;
              default:
                this.$refs.LoginUp.showError("系统繁忙");
            }
          })
      },

      //生成二维码
      generateQrCode() {
        this.$post("/lg/login/getwxstate", {})
        .then(res => {
          this.stateCode = res.data;
          console.log(window.location.protocol + "//" + window.location.host + "/qrcodelogin?state=" + this.stateCode);
        
          jQuery('#qrcodeTable').qrcode({
            render:"table",
            width:200,
            height:200,
            text:window.location.protocol + "//" + window.location.host + "/qrcodelogin?state=" + this.stateCode 
          });	

          //这个暂时放这里
          this.wxLogin();
        })
        .catch(err => {
          this.$refs.LoginUp.showError("系统繁忙,请先用其他方式登录");
          return;
        });
      },

      wxLogin(){
        if (!this.stateCode) {
          return ;
        }
        var req = {
          ident: this.ident,
          state: this.stateCode,
        }

        //定时处理(调用api)
        var that = this;
        var timesRun = 0;
        var interval = setInterval(function(){
            timesRun += 1;
            if(timesRun === 60*5){    
                clearInterval(interval);    
            }

            that.$post("/lg/login/wxlogin", req)
            .then(res=>{
                if (res.data  == "success") {
                  return;
                }

                clearInterval(interval);

                if (that.changepwd == 1) {
                  that.$router.push({ path: '/changepwd'});   
                  return;
                }

                if (that.ident && that.callback) {
                  window.location.href = JoinUrlParams(decodeURIComponent(that.callback),{code:res.data})
                  return;
                }
                that.$router.push({ path: '/chose'});
            })
            .catch(err=>{
              var type = 0;
              switch (err.response.status) {
                    case 400:
                        type = 3;
                        break;
                    case 406:
                        type = 4;
                        break;
                    case 408:
                        type = 5;
                        break;
                    case 510:
                        type = 6;
                        break;
                    case 401:
                        type = 7;
                        break;
                    case 415:
                        type = 1;
                        break;
                    default:
                        type = 0
// 400: 用户被锁定或被禁用，暂时无法登录 //3
// 406: 微信登录过程中有些参数丢失,请正常登录 //4
// 408: 微信登录标识过期,请重新登录 //5
// 510: 调用微信失败，稍后再登录 // 6
// 500: 系统出错，等会在试 //0
// 401: 没有关注公众号 // 7
// 415: 没有相应权限，请联系管理员 //1
                }
                that.$router.push({ path: '/errpage', query: {type: type}});
            });

        }, 1500);
      }

    }
  }
</script>
