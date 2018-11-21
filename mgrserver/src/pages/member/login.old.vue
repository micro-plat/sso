<template>
<div>

  <div id="bg" style=" background:#2196f3">
    <canvas></canvas>
    <canvas></canvas>
    <canvas></canvas>
  </div>

  <h1 class="visible-lg-block visible-md-block">{{sys.name}}</h1>
  <div class="sub-main-w3">
    <form  v-show="!upLogin" @submit.prevent="toggleQRCode">
      <h2>扫码登录
        <i class="fas fa-level-down-alt"></i>
      </h2>
      <div class="form-style-agile center">
        <label>
          <i class="fas fa-qrcode"></i>
            微信扫码，安全登录
        </label>

        <div id="qrcode" class="code-center"></div>

      </div>

      <div class="form-style-agile center">

         <label>打开 微信 扫一扫登录</label>
        <p class="wthree-text"> <span v-if="qrcode_msg">
<i class="fa fa-exclamation-triangle"></i> {{qrcode_msg}}
         </span></p>

      </div>

      <input type="submit" class="submit" value="返回">
    </form>
  </div>
  <div class="sub-main-w3">
    <form action="#" method="post" @submit.prevent="loginNow" v-show="upLogin">
      <div class="logo-img" v-if="login.ident == 'coupon'">
         <img src="../../img/zsh-logo.png" alt="">
      </div>
      <h2>
        用户登录
        <i class="fas fa-level-down-alt"></i>
      </h2>

      <div v-if="sys.wechat_status == 1" class="scan_code"><img class="scan_code_icon_text" src="../../assets/scan_code_icon_text.png"><a @click="toggleQRCode" class="scan_code_icon"><img src="../../assets/scan_code_icon.png"></a></div>
      <div class="form-style-agile">
        <label>
                                            		<i class="fas fa-user"></i>
                                            		用户名
                                            		</label>
        <input placeholder="请输入用户名" maxlength="32" v-model="login.username" type="text" required oninvalid="setCustomValidity('请输入用户名')" oninput="setCustomValidity('')">
      </div>
      <div class="form-style-agile">
        <label>
                                            		<i class="fas fa-unlock-alt"></i>
                                            		密码
                                            		</label>
        <input placeholder="请输入密码" maxlength="32" v-model="login.password" type="password" required oninvalid="setCustomValidity('请输入密码')" oninput="setCustomValidity('')">
      </div>

      <div class="form-style-agile" v-if="sys.wechat_status == 1">
        <label>
                                            		<i class="fas fa-check"></i>
                                            		 验证码
                                            		</label>
        <input placeholder="请输入微信验证码" maxlength="5" v-model="login.wxcode" type="text" required oninvalid="setCustomValidity('请输入微信验证码')" oninput="setCustomValidity('')">
      </div>

      <!-- checkbox -->
      <div class="wthree-text" v-if="sys.wechat_status == 1">
        <ul>
          <li>
            <a @click="getValidCode" v-show="valid_show">获取验证码</a>
            <label v-show="!valid_show">微信验证码已发送到手机:{{valid_left_second}}秒</label>
          </li>
         <li>
            <label v-if="errmsg">
              <span class="">
                <i class="fas fa-exclamation-triangle"></i>
              {{errmsg}}
              </span>

             </label>
          </li>
        </ul>
      </div>
      <input type="submit" class="submit" value="立即登录">
    </form>
  </div>
  <div class="footer">
    <p>中国石化四川分公司Copyright&copy; 2018 版权所有 蜀ICP备 17001788号-2</p>
  </div>

</div>
</template>

<script>
import { drawDynamicsBG } from "../../services/bg";
import { setInterval, clearInterval, clearTimeout, setTimeout } from "timers";
import QRCode from "qrcodejs2";
// document.title = "用户登录";
import Cookies from 'js-cookie'
export default {
  data() {
    return {
      sys: {
        name: "用户系统",
        logo: "logo.png"
      },
      upLogin: true,
      errmsg: "",
      qrcode_msg: "",
      login: {
        username: localStorage.getItem("username"),
        user_id:null,
        password: "",
        wxcode: "",
        ident: "sso"
      },
      valid_show: true,
      valid_left_second: 10,
      code: null,
    };
  },
  mounted() {
    sessionStorage.removeItem("__jwt__");
    drawDynamicsBG("bg");
    this.code = this.$route.query.code;
    this.login.ident = this.$route.params.ident
      ? this.$route.params.ident
      : "sso";
    // let ua = navigator.userAgent.toLowerCase();
    // //android终端
    // let isAndroid = ua.indexOf('Android') > -1 || ua.indexOf('Adr') > -1;
    // //ios终端
    // let isiOS = !!ua.match(/\(i[^;]+;( U;)? CPU.+Mac OS X/);
    // //是否微信环境
    // if (isWeixinBrowser()) {
    //   this.wxLogin();
    // } else {
    //   if (/(iPhone|iPad|iPod|iOS)/i.test(navigator.userAgent)) {
    //     //ios
    //      alert("请在微信中打开");
    //      return false;
    //   } else if (/(Android)/i.test(navigator.userAgent)) {
    //     //android
    //     alert("请在微信中打开");
    //     return false;
    //   }else{
    //     this.pcLogin();
    //   }
    // }
    this.pcLogin();
    // function isWeixinBrowser() {
    //   return (/micromessenger/.test(ua)) ? true : false;
    // }
  },
  watch: {
    valid_left_second: function(n, o) {
      if (n == 0) {
        clearTimeout(this.timer);
        this.valid_left_second = 20;
        this.valid_show = true;
      }
    }
  },
  methods: {
    pcLogin(){
      this.$fetch("/sso/sys/get", { ident: this.login.ident })
        .then(res => {
          this.sys = res;
          console.log(this.sys);
          document.title = "用户登录 - " + res.name;
        })
        .catch(err => {
          // console.log("err:", err);
          this.errmsg = err.response ? err.response : "无法请求接口服务器";
        });

      let self = this;
      this.wsconn = this.$ws();
      this.wsconn.onclose = function() {
        self.qrcode_msg = "暂时无法使用二维码登录";
        self.wsconn.reconnect();
      };
      this.wsconn.onopen = function() {
        self.$wssend(
          self.wsconn,
          JSON.stringify({
            service: "/qrcode/login/get",
            ident: self.login.ident
          })
        );
        self.qrcode_msg = "";
      };
      this.wsconn.onmessage = function(evt) {
        if (!evt.data) {
          return;
        }
        let res = JSON.parse(evt.data);
        if (res.code != 200) {
          console.log("暂时无法使用二维码登录:", res);
          self.qrcode_msg = "暂时无法使用二维码登录";
          return;
        }
        switch (res.service) {
          case "/qrcode/login/success": //登录成功,调用/sso/login/code设置jwt并转跳到网站首页
            self
              .$post("/sso/login/code", res.data)
              .then(res => {
                self.$router.push("/?ident=" + self.login.ident);
              })
              .catch(err => {
                console.log("登录失败:", err);
              });
            break;
          case "/qrcode/login/check": //已扫码等待用户确认登录
            self.qrcode_msg = "请在手机上确认登录";
            break;
          case "/qrcode/login/get": //已查询到二维码信息，显示到页面
            let url = decodeURIComponent(res.data.url);
            if (self.qrcode) {
              self.qrcode.makeCode(url);
            } else {
              self.qrcode = new QRCode("qrcode", {
                width: 200,
                height: 200, // 高度
                colorDark: "#000000",
                colorLight: "#ffffff",
                text: url
              });
            }
        }
      };
      this.$wssend(
        this.wsconn,
        JSON.stringify({
          service: "/qrcode/login/get",
          ident: this.login.ident
        })
      );
    },
    wxLogin(){
      this.$fetch("/sso/sys/get", { ident: this.login.ident })
        .then(res => {
          this.sys = res;
          document.title = "用户登录 - " + res.name;
        })
        .catch(err => {
          this.errmsg = err.response ? err.response : "无法请求接口服务器";
        });
      if (this.code == null) {
        let r = "http://sso.sinopecscsy.com/"+ this.login.ident +"/member/login?ident=";
        window.location = "https://open.weixin.qq.com/connect/oauth2/authorize?appid=wx9e02ddcc88e13fd4&redirect_uri="+ r +"&response_type=code&scope=snsapi_userinfo&state=STATE#wechat_redirect";
      }else{
        this.$post("/wx/login", {code:this.code,ident:this.login.ident})
          .then(res => {
            window.location.href = "http://sso.sinopecscsy.com/?ident=" + this.login.ident
          })
          .catch(err => {
            alert(err);
            this.$notify({
              title: '错误',
              message: '自动登录失败，请联系管理员',
              type: 'error',
              offset: 50,
              duration:4000,
            });
          });
      }
    },
    loginNow: function() {
      this.$post("/sso/login", this.login)
        .then(res => {
          sessionStorage.setItem("username", this.login.username);
          if(sessionStorage.getItem("__jwt__") != null){
              sessionStorage.setItem("code",res.code);
              let host = "http://" + window.location.host;
              window.location.href = host + "/?ident=" + this.login.ident
          }
        })
        .catch(err => {
          if (err.response) {
            switch (err.response.status) {
              case 415:
                this.$notify({
                  title: '错误',
                  message: '不允许登录系统',
                  type: 'error',
                  offset: 50,
                  duration:3000,
                });
                break;
              case 403:
                this.$notify({
                  title: '错误',
                  message: '用户名或密码错误',
                  type: 'error',
                  offset: 50,
                  duration:3000,
                });
                break;
              case 423:
                this.$notify({
                  title: '错误',
                  message: '用户被锁定暂时无法登录',
                  type: 'error',
                  offset: 50,
                  duration:3000,
                });
                break;
              case 901:
                this.$notify({
                  title: '错误',
                  message: '微信验证码错误',
                  type: 'error',
                  offset: 50,
                  duration:3000,
                });
                break;
            }
          }else{
            this.$notify({
                  title: '错误',
                  message: '网络错误，请稍后再试',
                  type: 'error',
                  offset: 50,
                  duration:3000,
                });
          }
          setTimeout(this.clearErrMsg, 5000);
        });
    },
    getValidCode: function() {
      this.errmsg = "";
      if (this.login.username == null || this.login.username == "") {
        return false;
      }
      this.$post("/sso/wxcode/get", this.login)
        .then(response => {
          this.valid_show = false;
          this.timer = setInterval(this.leftSecond, 1000);
        })
        .catch(err => {
          if (err.response) {
            switch (err.response.status) {
              case 406:
                this.errmsg = "用户名或系统不存在或未开通微信账户";
                break;
            }
          }
          if (this.errmsg == "") {
            this.errmsg = "暂时无法获取请稍后再试";
          }

          setTimeout(this.clearErrMsg, 5000);
        });
    },
    clearErrMsg: function() {
      this.errmsg = "";
    },
    toggleQRCode: function() {
      this.upLogin = !this.upLogin;
    },
    leftSecond: function() {
      this.valid_left_second = this.valid_left_second - 1;
    }
  }
};
</script>


<style scoped>
.wthree-text li {
  height: 18px;
  line-height: 18px;
}
.wthree-text span {
  color: #f7296f;
  font-weight: 600;
}

.scan_code {
  position: absolute;
  float: right;
  right: 5%;
  top: 5%;
}

.scan_code_icon_text {
  display: none;
}

.scan_code_icon img {
  opacity: 0.5;
  filter: alpha(opacity=50);
}

.scan_code:hover .scan_code_icon_text {
  display: inline-block;
}

.scan_code a:hover img {
  opacity: 1;
  filter: alpha(opacity=100);
}

.code-center {
  display: flex;
  justify-content: center;
}

ol,
ul {
  list-style: none;
  margin: 0px;
  padding: 0px;
}

blockquote,
q {
  quotes: none;
}

blockquote:before,
blockquote:after,
q:before,
q:after {
  content: "";
  content: none;
}

table {
  border-collapse: collapse;
  border-spacing: 0;
}

/* start editing from here */

a {
  text-decoration: none;
}

.txt-rt {
  text-align: right;
}

/* text align right */

.txt-lt {
  text-align: left;
}

/* text align left */

.txt-center {
  text-align: center;
}

/* text align center */

.float-rt {
  float: right;
}

/* float right */

.float-lt {
  float: left;
}

/* float left */

.clear {
  clear: both;
}

/* clear float */

.pos-relative {
  position: relative;
}

/* Position Relative */

.pos-absolute {
  position: absolute;
}

/* Position Absolute */

.vertical-base {
  vertical-align: baseline;
}

/* vertical align baseline */

.vertical-top {
  vertical-align: top;
}

/* vertical align top */

nav.vertical ul li {
  display: block;
}

/* vertical menu */

nav.horizontal ul li {
  display: inline-block;
}

/* horizontal menu */

img {
  max-width: 100%;
}

/*end reset*/

/* body {
  font-family: "Open Sans", sans-serif;
  font-size: 100%;
  background: #2196f3;
} */

/*--header--*/

h1 {
  font-size: 3.5em;
  color: #fff;
  letter-spacing: 3px;
  text-align: center;
  margin: 3vw 1vw;
  font-family: "Josefin Sans", sans-serif;
}

h1 span,
h2 i {
  color: #f7296f;
}

h2 {
  color: #fff;
  font-size: 22px;
  font-weight: 500;
  letter-spacing: 1px;
  text-transform: capitalize;
  margin-bottom: 1em;
}

/*--//header--*/

/*-- content --*/

.sub-main-w3 {
  display: -webkit-flex;
  display: -webkit-box;
  display: -moz-flex;
  display: -moz-box;
  display: -ms-flexbox;
  display: flex;
  align-items: center;
  -webkit-box-pack: center;
  -moz-box-pack: center;
  -ms-flex-pack: center;
  -webkit-justify-content: center;
  justify-content: center;
}

.sub-main-w3 form {
  max-width: 600px;
  margin: 0 5vw;
  background: rgba(10, 10, 10, 0.17);
  padding: 3.5vw;
  box-sizing: border-box;
  display: -webkit-flex;
  display: flex;
  flex-wrap: wrap;
  justify-content: center;
  border-bottom: 8px solid #f7296f;
  border-radius: 30px 30px 50px 50px;
  position: relative;
}

.form-style-agile {
  margin-bottom: 1.5em;
  flex-basis: 100%;
  -webkit-flex-basis: 100%;
}

.center {
  text-align: center;
}

.sub-main-w3 label {
  font-size: 14px;
  color: #fff;
  display: inline-block;
  font-weight: 500;
  margin-bottom: 12px;
  text-transform: capitalize;
  letter-spacing: 1px;
}

.sub-main-w3 label i {
  font-size: 15px;
  margin-left: 5px;
  color: #f7296f;
  border-radius: 50%;
  line-height: 1.9;
  text-align: center;
}

.form-style-agile input[type="text"],
.form-style-agile input[type="password"] {
  width: 100%;
  color: #000;
  outline: none;
  font-size: 14px;
  letter-spacing: 1px;
  padding: 15px 15px;
  box-sizing: border-box;
  border: none;
  border: 1px solid #000;
  background: #fff;
}

.sub-main-w3 input[type="submit"] {
  color: #fff;
  background: #f7296f;
  border: none;
  padding: 13px 0;
  margin-top: 30px;
  outline: none;
  width: 36%;
  font-size: 16px;
  cursor: pointer;
  letter-spacing: 2px;
  -webkit-transition: 0.5s all;
  -o-transition: 0.5s all;
  -moz-transition: 0.5s all;
  -ms-transition: 0.5s all;
  transition: 0.5s all;
  box-shadow: 2px 2px 6px rgba(0, 0, 0, 0.49);
}

.sub-main-w3 input[type="submit"]:hover {
  background: #000;
  -webkit-transition: 0.5s all;
  -o-transition: 0.5s all;
  -moz-transition: 0.5s all;
  -ms-transition: 0.5s all;
  transition: 0.5s all;
}

.wthree-text {
  width: 100%;
}

.wthree-text ul li:nth-child(1) {
  float: right;
}

.wthree-text ul li:nth-child(2) {
  float: left;
}

.wthree-text ul li {
  display: inline-block;
}

.wthree-text ul li a {
  color: #fff;
  font-size: 14px;
  letter-spacing: 1px;
  font-weight: 500;
}

/*-- checkbox --*/

.wthree-text label {
  font-size: 14px;
  color: #fff;
  cursor: pointer;
  position: relative;
}

.wthree-text {
  text-align: center;
}

input.invalid {
}

input.checkbox {
  background: #f7296f;
  cursor: pointer;
  width: 1.2em;
  height: 1.2em;
  vertical-align: text-bottom;
}

input.checkbox:before {
  content: "";
  position: absolute;
  width: 1.2em;
  height: 1.2em;
  background: inherit;
  cursor: pointer;
}

input.checkbox:after {
  content: "";
  position: absolute;
  top: 4px;
  left: 4px;
  z-index: 1;
  width: 1em;
  height: 1em;
  border: 1px solid #fff;
  -webkit-transition: 0.4s ease-in-out;
  -moz-transition: 0.4s ease-in-out;
  -o-transition: 0.4s ease-in-out;
  transition: 0.4s ease-in-out;
}

input.checkbox:checked:after {
  -webkit-transform: rotate(-45deg);
  -moz-transform: rotate(-45deg);
  -o-transform: rotate(-45deg);
  -ms-transform: rotate(-45deg);
  transform: rotate(-45deg);
  height: 0.5rem;
  border-color: #fff;
  border-top-color: transparent;
  border-right-color: transparent;
}

.anim input.checkbox:checked:after {
  -webkit-transform: rotate(-45deg);
  -moz-transform: rotate(-45deg);
  -o-transform: rotate(-45deg);
  -ms-transform: rotate(-45deg);
  transform: rotate(-45deg);
  height: 0.5rem;
  border-color: transparent;
  border-right-color: transparent;
  animation: 0.4s rippling 0.4s ease;
  animation-fill-mode: forwards;
}

@keyframes rippling {
  50% {
    border-left-color: #fff;
  }
  .clear {
    clear: both;
  }
  body {
    font-family: "Open Sans", sans-serif;
    font-size: 100%;
    background: #2196f3;
  }
  /*--header--*/
  h1 {
    font-size: 3.5em;
    color: #fff;
    letter-spacing: 3px;
    text-align: center;
    margin: 3vw 1vw;
    font-family: "Josefin Sans", sans-serif;
  }
  h1 span,
  h2 i {
    color: #f7296f;
  }
  h2 {
    color: #fff;
    font-size: 22px;
    font-weight: 500;
    letter-spacing: 1px;
    text-transform: capitalize;
    margin-bottom: 1em;
  }
  /*--//header--*/
  /*-- content --*/
  .sub-main-w3 {
    display: -webkit-flex;
    display: -webkit-box;
    display: -moz-flex;
    display: -moz-box;
    display: -ms-flexbox;
    display: flex;
    align-items: center;
    -webkit-box-pack: center;
    -moz-box-pack: center;
    -ms-flex-pack: center;
    -webkit-justify-content: center;
    justify-content: center;
  }
  .sub-main-w3 form {
    max-width: 600px;
    margin: 0 5vw;
    background: rgba(10, 10, 10, 0.17);
    padding: 3.5vw;
    box-sizing: border-box;
    display: -webkit-flex;
    display: flex;
    flex-wrap: wrap;
    justify-content: center;
    border-bottom: 8px solid #f7296f;
    border-radius: 30px 30px 50px 50px;
  }
  .form-style-agile {
    margin-bottom: 1.5em;
    flex-basis: 100%;
    -webkit-flex-basis: 100%;
  }
  .sub-main-w3 label {
    font-size: 14px;
    color: #fff;
    display: inline-block;
    font-weight: 500;
    margin-bottom: 12px;
    text-transform: capitalize;
    letter-spacing: 1px;
  }
  .sub-main-w3 label i {
    font-size: 15px;
    margin-left: 5px;
    color: #f7296f;
    border-radius: 50%;
    line-height: 1.9;
    text-align: center;
  }
  .form-style-agile input[type="text"],
  .form-style-agile input[type="password"] {
    width: 100%;
    color: #000;
    outline: none;
    font-size: 14px;
    letter-spacing: 1px;
    padding: 15px 15px;
    box-sizing: border-box;
    border: none;
    border: 1px solid #000;
    background: #fff;
  }
  .sub-main-w3 input[type="submit"] {
    color: #fff;
    background: #f7296f;
    border: none;
    padding: 13px 0;
    margin-top: 30px;
    outline: none;
    width: 36%;
    font-size: 16px;
    cursor: pointer;
    letter-spacing: 2px;
    -webkit-transition: 0.5s all;
    -o-transition: 0.5s all;
    -moz-transition: 0.5s all;
    -ms-transition: 0.5s all;
    transition: 0.5s all;
    box-shadow: 2px 2px 6px rgba(0, 0, 0, 0.49);
  }
  .sub-main-w3 input[type="submit"]:hover {
    background: #000;
    -webkit-transition: 0.5s all;
    -o-transition: 0.5s all;
    -moz-transition: 0.5s all;
    -ms-transition: 0.5s all;
    transition: 0.5s all;
  }
  .wthree-text {
    width: 100%;
  }
  .wthree-text ul li:nth-child(1) {
    float: left;
  }
  .wthree-text ul li:nth-child(2) {
    float: right;
  }
  .wthree-text ul li {
    display: inline-block;
  }
  .wthree-text ul li a {
    color: #fff;
    font-size: 14px;
    letter-spacing: 1px;
    font-weight: 500;
  }
  /*-- checkbox --*/
  .wthree-text label {
    font-size: 15px;
    color: #fff;
    cursor: pointer;
    position: relative;
  }
  .wthree-text {
    text-align: center;
  }
  .anim input.checkbox:checked:after {
    -webkit-transform: rotate(-45deg);
    -moz-transform: rotate(-45deg);
    -o-transform: rotate(-45deg);
    -ms-transform: rotate(-45deg);
    transform: rotate(-45deg);
    height: 0.5rem;
    border-color: transparent;
    border-right-color: transparent;
    animation: 0.4s rippling 0.4s ease;
    animation-fill-mode: forwards;
  }
  @keyframes rippling {
    50% {
      border-left-color: #fff;
    }
    100% {
      border-bottom-color: #fff;
      border-left-color: #fff;
    }
  }
  ::-webkit-input-placeholder {
    color: #000;
  }
  :-moz-placeholder {
    /* Firefox 18- */
    color: #000;
  }
  ::-moz-placeholder {
    /* Firefox 19+ */
    color: #000;
  }
  :-ms-input-placeholder {
    color: #000;
  }
  .footer {
    margin: 4vw 0.3vw 2vw;
  }
  .footer p {
    font-size: 14px;
    color: #fff;
    letter-spacing: 2px;
    text-align: center;
    line-height: 1.8;
  }
  .footer p a {
    color: #fff;
    -webkit-transition: 0.5s all;
    -o-transition: 0.5s all;
    -moz-transition: 0.5s all;
    -ms-transition: 0.5s all;
    transition: 0.5s all;
  }
  .footer p a:hover {
    color: #f7296f;
    -webkit-transition: 0.5s all;
    -o-transition: 0.5s all;
    -moz-transition: 0.5s all;
    -ms-transition: 0.5s all;
    transition: 0.5s all;
  }
  #bg {
    position: fixed;
    top: 0;
    left: 0;
    width: 100%;
    height: 100%;
    z-index: -1;
  }
  #bg canvas {
    position: absolute;
    top: 0;
    left: 0;
    width: 100%;
    height: 100%;
  }
  @media (max-width: 1920px) {
    h1 {
      font-size: 2.6vw;
    }
  }
  @media (max-width: 1024px) {
    h1 {
      font-size: 4.5vw;
    }
  }
  @media (max-width: 800px) {
    h1 {
      font-size: 2.8em;
    }
  }
  @media (max-width: 480px) {
    h1 {
      font-size: 2.5em;
      letter-spacing: 1px;
    }
    .sub-main-w3 form {
      padding: 7.5vw;
    }
    .footer p {
      letter-spacing: 1px;
      font-size: 12px;
      padding: 0 20px;
    }
  }
  @media (max-width: 414px) {
    .form-style-agile input[type="text"],
    .form-style-agile input[type="password"] {
      font-size: 13px;
      padding: 13px 15px;
    }
    .wthree-text ul li:nth-child(1),
    .wthree-text ul li:nth-child(2) {
      float: none;
    }
    .wthree-text ul li:nth-child(2) {
      margin-top: 10px;
    }
    .sub-main-w3 input[type="submit"] {
      width: 56%;
    }
    .wthree-text ul li {
      display: block;
    }
  }
  @media (max-width: 320px) {
    h1 {
      font-size: 2em;
      margin: 5vw 1vw;
    }
    .sub-main-w3 form {
      padding: 25px 14px;
    }
  }
}

/*-- //checkbox --*/

/*--placeholder-color--*/

::-webkit-input-placeholder {
  color: #000;
}

:-moz-placeholder {
  /* Firefox 18- */
  color: #000;
}

::-moz-placeholder {
  /* Firefox 19+ */
  color: #000;
}

:-ms-input-placeholder {
  color: #000;
}

/*-- //placeholder-color --*/

/*-- //content --*/

/*-- copyright --*/

.footer {
  margin: 4vw 0.3vw 2vw;
}

.footer p {
  font-size: 14px;
  color: #fff;
  letter-spacing: 2px;
  text-align: center;
  line-height: 1.8;
}

.footer p a {
  color: #fff;
  -webkit-transition: 0.5s all;
  -o-transition: 0.5s all;
  -moz-transition: 0.5s all;
  -ms-transition: 0.5s all;
  transition: 0.5s all;
}

.footer p a:hover {
  color: #f7296f;
  -webkit-transition: 0.5s all;
  -o-transition: 0.5s all;
  -moz-transition: 0.5s all;
  -ms-transition: 0.5s all;
  transition: 0.5s all;
}

/*-- //copyright --*/

/*-- bg effect --*/

#bg {
  position: fixed;
  top: 0;
  left: 0;
  width: 100%;
  height: 100%;
  z-index: -1;
}

#bg canvas {
  position: absolute;
  top: 0;
  left: 0;
  width: 100%;
  height: 100%;
}

/*-- //bg effect --*/

/*--responsive--*/

@media (max-width: 1920px) {
  h1 {
    font-size: 2.6vw;
  }
}

@media (max-width: 1024px) {
  h1 {
    font-size: 4.5vw;
  }
}

@media (max-width: 800px) {
  h1 {
    font-size: 2.8em;
  }
}

@media (max-width: 480px) {
  h1 {
    font-size: 2.5em;
    letter-spacing: 1px;
  }
  .sub-main-w3 form {
    padding: 7.5vw;
  }
  .footer p {
    letter-spacing: 1px;
    font-size: 12px;
    padding: 0 20px;
  }
}

@media (max-width: 414px) {
  .form-style-agile input[type="text"],
  .form-style-agile input[type="password"] {
    font-size: 13px;
    padding: 13px 15px;
  }
  .wthree-text ul li:nth-child(1),
  .wthree-text ul li:nth-child(2) {
    float: none;
  }
  .wthree-text ul li:nth-child(2) {
    margin-top: 10px;
  }
  .sub-main-w3 input[type="submit"] {
    width: 56%;
  }
  .wthree-text ul li {
    display: block;
  }
}

@media (max-width: 320px) {
  h1 {
    font-size: 2em;
    margin: 5vw 1vw;
  }
  .sub-main-w3 form {
    padding: 25px 14px;
  }
}
.logo-img{
    text-align: center;
    position: absolute;
    top: -12px;
    }
/*--//responsive--*/
</style>
