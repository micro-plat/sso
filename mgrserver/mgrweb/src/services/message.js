//提示成功模板信息
var tmplSuccess = { title: "成功", message: "操作成功", offset: 50, duration: 2000 }

//提示失败模板信息
var tmplFailed = { title: "错误", message: "网络错误,请稍后再试", offset: 50, duration: 2000 }

var __Vue__ = {}

//提示组件初始化
export function Message(Vue) {
    __Vue__ = Vue.prototype
}

//成功提示调用
Message.prototype.success = function (data) {
    show(data, "success")
}

//失败提示调用
Message.prototype.fail = function (data) {
    show(data, "fail")
}

//警告提示调用
Message.prototype.warning = function (data) {
    show(data, "warning")
}

//消息提示调用
Message.prototype.info = function (data) {
    show(data, "info")
}

//设置组件默认模板
Message.prototype.setTmpl = function (sucTmpl, failTmpl) {
    if(!sucTmpl){
        return
    }
    if(typeof sucTmpl == "object"){
        Object.assign(tmplSuccess, sucTmpl)
    }
    if(!failTmpl){
        return
    }
    if(typeof failTmpl == "object"){
        Object.assign(tmplFailed, failTmpl)
    }
}

function show(data, type){
    var tmpl = type == "success" ? tmplSuccess : tmplFailed //初始值
    if(typeof data == "string"){
        tmpl.message = data
    }

    if(typeof data == "object"){
        Object.assign(tmpl, data)
    }

    tmpl.type = type
    if(typeof __Vue__.$toast == "function"){
        __Vue__.$toast(tmpl);
        return
    }

    if(typeof __Vue__.$notify == "function"){
        tmpl.type = type == "fail" ? "error" : type
        __Vue__.$notify(tmpl);
        return
    }

    throw new Error("还未安装vant组件或element-ui组件，请先安装相关组件");
}
