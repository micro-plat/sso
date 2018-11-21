<template>
  <div class="wrap">
    <!-- 确认 -->
    <div v-if="status==1">
      <div class="confirm-login-img"><img src="../../assets/confirm_login_icon.png"></div>

      <label class="confirm-login-title">正在绑定帐号.....</label>

    </div>
    <!-- 成功 -->
    <div v-if="status==0">
      <div class="confirm-login-img success"><img src="../../assets/login_success.png"></div>
      <div class="result color-zihong">绑定成功</div>
      <label class="confirm-login-tips">绑定成功，你现在可以使用微信登录系统</label>
    </div>
    <!-- 失败 -->
    <div v-if="status==2">
      <div class="confirm-login-img failure"><img src="../../assets/login_failure.png"></div>
      <div class="result">绑定失败</div>
      <label class="confirm-login-tips">绑定失败了,请稍后再试</label>
    </div>

    <div v-if="status==3">
      <div class="confirm-login-img failure"><img src="../../assets/login_failure.png"></div>
      <div class="result">发生错误</div>
      <label class="confirm-login-tips"></label>
    </div>
  </div>
</template>

<script>
  export default {
    data() {
      return {
        errorMsg: "",
        status: 1,
        guid: null,
        code: null,
      };
    },
    mounted() {
      this.guid = this.$route.query.guid;
      this.code = this.$route.query.code;
      this.$fetch("/sso/user/bind", { guid: this.guid,code: this.code })
        .then(res => {
          this.status = 0
        })
        .catch(err => {
          this.status = 2
        });
    },
  };
</script>
