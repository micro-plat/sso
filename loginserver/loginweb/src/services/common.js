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
  
  var pvs = []
  for (var item in params) {
    pvs.push(item + "=" + params[item] );
  }
  
  if(pvs.length>0){
    var char = "?";
    var index = url.indexOf("?");
    if (index > 0) {
      char = "&";
    }
    url = url + char + pvs.join("&");
  }
  return url ;
}

//获取url的host
export function GetUrlHosts(url) {
  return url.substr(0, url.lastIndexOf("/"));
}


/**
 * 返回登录的地址
 * @param {*子系统标识} ident 
 */
export function jumpLogin(ident,callback) {
  var pathT = '/'+ ident +'/login';
  if (!ident) {
      pathT = '/login'
  }
  return pathT;
}