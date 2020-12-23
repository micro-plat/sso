export function trimError(err) {
    var message = "系统错误,请稍后再试";
    if (!err.response.data.data) {
        return message;
    }
    var messageT = err.response.data.data; 
    if (messageT && messageT.length > 6 && messageT.indexOf("error:",0) == 0) {
        message = messageT.substr(6);
    }
    return message;
}