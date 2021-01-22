
/*
* Env对象使用时须通过引用并进行初始化
* import evn from './evn'
* Vue.use(evn);
* 或 可配置加载文件地址(需json格式的文件)
* Vue.use(evn,"../static/env.conf.json")
*/
export function Env(data) {
    Env.prototype.conf = {}
    if(typeof data != "object"){
        throw new Error("无效参数，类型:object 返回一个对象数据");
    }
    Object.assign(Env.prototype.conf, data)
}

/*
*配置数据加载
*var that = this
*this.$env.load(function(){
*   var ress = that.$http.xpost("/dds/dictionary/get", { dic_type: "operate_action" }, "", false) || {}
*   return ress[0]
*})
*/
Env.prototype.load = function (f) {
    if (typeof f !== "function"){
        return
    }

    let conf = f() || {}
    return Object.assign(Env.prototype.conf, conf)   
}