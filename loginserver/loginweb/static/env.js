//注入初始化
export default {
    install: function(Vue){
        Vue.prototype.$env = new Env()
    }
}

/*
* Env对象使用时须通过引用并进行初始化
* import evn from './evn'
* Vue.use(evn);
*/
function Env() {
    // require("./env.conf.js") 获取到需要加载的配置文件 './env.conf.js'配置文件路径
    let conf = {}
    conf = require("../../static/env.conf.js") || require("../../public/env.conf.js")
    Env.prototype.Conf = conf
}

//配置数据加载
Env.prototype.load = function (f) {
    if (typeof f !== "function"){
        return
    }
    let conf = f() || {}
    return Object.assign(Env.prototype.Conf, conf)   
}