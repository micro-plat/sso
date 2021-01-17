<template></template>

<script>
export default {
  data() {
    return {};
  },
  mounted() {
    this.validSsoLogin();
  },
  methods: {
    validSsoLogin() {
      var returnURL = this.$route.query.returnurl;
      this.$http
        .post("/sso/login/verify", { code: this.$route.query.code })
        .then(res => {
          console.log("a.xxxxxxxxxxxx")
          localStorage.setItem(
            "userinfo",
            JSON.stringify({ name: res.user_name, role: res.role_name })
          );

          if (returnURL) {
            window.location = returnURL; 
            return;
          }
          this.$router.push("/");
        })
        .catch(err => {
            console.log(err);
        });
    }
  }
};
</script>
