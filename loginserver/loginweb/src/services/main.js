import { Enum } from './enum'
import { Http } from './http'
import { Env } from './env'
import { Utility } from './utility'

/*
* 初始化注入
* import utility from './utility'
* Vue.use(utility);
* 或传入加载配置文件路径
* Vue.use(utility, "../static/env.conf.json");
*/
export default {
    install: function(Vue, path){
        Vue.prototype.$enum = new Enum();
        Vue.prototype.$http = new Http();
        Vue.prototype.$env = new Env(path)
        Vue.prototype.$utility = new Utility();
    }
}