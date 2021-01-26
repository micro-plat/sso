
// 拼接url参数
export function JoinUrlParams(url, params) {

  var ops = getQueryParams(url);
  var idx = url.indexOf('?');
  if(idx>0){
    url = url.substring(0,idx);
  }
  
  Object.assign(ops, params);
  var pvs = [];
  for (var item in ops) {
    pvs.push(item + "=" + ops[item]);
  }

  if (pvs.length > 0) {
    var char = "?";
    url = url + char + pvs.join("&");
  }
  return url;
}

function getQueryParams(url) {
  var result = {};
  var idx = url.indexOf("?");
  if(idx<=0){
    return result;
  }
  var query = url.substring(idx + 1);
  //var query = window.location.search.substring(1);
  var kv = query.split("&");
  for (var i = 0; i < kv.length; i++) {
    var pair = kv[i].split("=");
    result[pair[0]] = pair[1];
  }
  return result;
}


/**
 * 返回登录的地址
 * @param {*子系统标识} ident 
 */
export function jumpLogin(ident, callback) {
  var pathT = '/' + ident + '/login';
  if (!ident) {
    pathT = '/login'
  }
  return pathT;
}

export function guid() {
  return 'xxxxxxxx-xxxx-4xxx'.replace(/[x]/g, function (c) {
    var r = Math.random() * 16 | 0, v = c == 'x' ? r : (r & 0x3 | 0x8);
    return v.toString(16);
  });
}