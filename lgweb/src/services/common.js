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

// 拼接url参数
export function JoinUrlParams(url,params) {
  var index = url.indexOf("?");
  if (index < 0) {
    url += "?"
  }
  for (var item in params) {
    url += item + "=" + params[item] + "&";
  }
  return url;
}

//获取url的host
export function GetUrlHost(url) {
  return url.substr(0, url.lastIndexOf("/"));
}
