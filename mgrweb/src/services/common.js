//获取UserIndex页面所需信息
function base4UserIndex(obj) {
  var formData = new FormData();
  obj.$post("/sso/base", formData)
      .then(res => {
        obj.roleList = res.rolelist;
      })
      .catch(err => {
        if (err.response) {
          // this.$router.push("/member/login");
        }
      });
}

export{base4UserIndex}
