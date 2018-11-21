<template>
  <div class="wrap">
    <!-- 确认 -->
    <div v-if="status==1">
    <div class="confirm-login-img"><img src="../../assets/confirm_login_icon.png"></div>

    <label class="confirm-login-title">正试图登录{{sys.name}}</label>
    <div class="confirm-btn">
    <input type="button" class="btn btn-primary confirm-login-btn" @click="allow" value="确认登录">
    </div>
    <div class="confirm-btn">
    <input type="button" class="cancel-login-btn" @click="deny" value="取消登录">
    </div>
    </div>
    <!-- 成功 -->
    <div v-if="status==0"> 
     <div class="confirm-login-img success"><img src="../../assets/login_success.png"></div>
     <div class="result color-zihong">登录成功</div>
     <label class="confirm-login-tips">{{sys.name}}</label>
     </div>
    <!-- 失败 -->
     <div v-if="status==2"> 
     <div class="confirm-login-img failure"><img src="../../assets/login_failure.png"></div>
     <div class="result">登录失败</div>
     <label class="confirm-login-tips">{{sys.name}}登录失败了，请稍后再试{{errorMsg}}</label>
     </div>
     
   

  </div>
</template>

<script>
//document.title = "确认登录";
export default {
  data() {
    return {
      errorMsg: "",
      status: 1,
      ident: "sso",
      sys: {}
    };
  },
  computed() {
   // document.title = "确认登录";
  },
  mounted() {
    this.ident = this.$route.query.ident;
    this.uid = this.$route.query.uid;
    this.code = this.$route.query.code;
    this.$fetch("/sso/sys/get", { ident: this.ident })
      .then(res => {
        this.sys = res;
      })
      .catch(err => {});
    this.$put("/qrcode/login", { ident: this.ident, uid: this.uid })
      .then(res => {})
      .catch(err => {
        this.errorMsg = err;
        //跳到失败页面
      });
  },
  methods: {
    allow: function() {
      this.$post("/qrcode/login", {
        ident: this.ident,
        uid: this.uid,
        code: this.code
      })
        .then(res => {
          this.status = 0;
          //跳到成功页面
        })
        .catch(err => {
          this.status = 2;
          //跳到成功页面
        });
    },
    deny: function() {
      this.status = 2;
      //跳到成功页面
    }
  }
};
</script>

