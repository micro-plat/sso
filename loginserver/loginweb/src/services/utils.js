export function trimError(err) {
    var message = "系统错误,请稍后再试";
    if (!err.response.data) {
        return message;
    }
    var messageT = err.response.data; 
    if (messageT && messageT.length > 6 && messageT.indexOf("error:",0) == 0) {
        message = messageT.substr(6);
    }
    return message;
}

/**
 * 返回登录的地址
 * @param {*子系统标识} ident 
 */
export function jumpLogin(ident) {
    var pathT = '/'+ ident +'/login';
    if (!ident) {
        pathT = '/login'
    }
    return pathT;
}