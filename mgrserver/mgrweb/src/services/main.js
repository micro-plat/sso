import { Enum } from './enum'
import { Http } from './http'
import { Env } from './env'
import { Utility } from './utility'
import { Auth } from './auth'

import packageData from '../../package.json'

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
        Vue.prototype.$utility = new Utility();

        Vue.prototype.$env = new Env(getConf(Vue, path))    
        Vue.prototype.$auth = new Auth(Vue);

        let that =　Vue.prototype

        //设置http请求的服务器地址
        if (that.$env.conf.system.apiHost){
            that.$http.setBaseURL(that.$env.conf.system.apiHost);
        }

        //处理接口返回４０３时自动跳转到指定的地址
        that.$http.addStatusCodeHandle(res => {
            var url = (res.headers || {}).location || ""; 
            if(url){
                window.location = url + encodeURIComponent(document.URL);
                return
            }
            
            let conf = that.$env.conf
            if (!conf.sso || !conf.sso.host || !conf.sso.ident){
                throw new Error("sso.host或sso.ident未配置");
            }
            window.location = conf.sso.host+"/" + conf.sso.ident+"/login?returnurl=" + encodeURIComponent(document.URL);
            return
            
        }, 403);

        //拉到服务器配置信息
        if (that.$env.conf.system.confURL){
            that.$env.load(function(){
                return that.$http.xget(that.$env.conf.system.confURL);  
            });
        }

        //拉取enum数据
        if (that.$env.conf.system.enumURL){
            that.$enum.callback(function(type){
                return that.$http.xget(that.$env.conf.system.enumURL, { dic_type: type || "" }, "") 
            })
        }
    }
}

function getConf(Vue, path){
    if(path)
        return Vue.prototype.$http.xget(path) || {};
    
    if(!packageData)
        return
    
    var vueVersion =  (packageData.dependencies.vue).charAt(1)
    path = vueVersion > 3 ?"../../public/env.conf.json" : "../../static/env.conf.json"   
    return Vue.prototype.$http.xget(path) || {}
}

