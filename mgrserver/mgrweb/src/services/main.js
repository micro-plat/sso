import { Enum } from './enum'
import { Http } from './http'
import { Env } from './env'
import { Utility } from './utility'
import { Sys } from './sys'
import { Message } from './message'

import packageData from '../../package.json'

/*
* 初始化注入
* import utility from './utility'
* Vue.use(utility);
* 或传入加载配置文件路径
* Vue.use(utility, "../static/env.conf.json");
*/
export default {
    install: function(Vue, inject403Code = true, path){
        Vue.prototype.$msg = new Message(Vue);
        Vue.prototype.$enum = new Enum();
        Vue.prototype.$http = new Http(Vue);
        Vue.prototype.$env = new Env(getConf(Vue, path))    
        Vue.prototype.$sys = new Sys(Vue);
        Vue.prototype.$utility = new Utility();

        let that = Vue.prototype

        //设置http请求的服务器地址
        if (that.$env.conf.api && that.$env.conf.api.host){
            that.$http.setBaseURL(that.$env.conf.api.host);
        }

        //处理接口返回403时自动跳转到指定的地址
        if(inject403Code){ //注入时可配置是否默认处理403
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
                window.location = conf.sso.host + "/" + conf.sso.ident + "/jump?returnurl=" + encodeURIComponent(document.URL);
                return;
            }, 403);

            inject405CodeHandle(that) //405权限处理
        }

        //拉到服务器配置信息
        if (that.$env.conf.api.confURL){
            that.$env.load(function(){
                return that.$http.xget(that.$env.conf.api.confURL);  
            });
        }

        //拉取enum数据
        if (that.$env.conf.api.enumURL){
            that.$enum.callback(function(type){
                return that.$http.xget(that.$env.conf.api.enumURL, { dic_type: type || "" }, "") 
            })
        }

        //保存初始数据
        if (that.$env.conf.enums){
            that.$enum.set(that.$env.conf.enums)
        }
    }
}

//获取配置数据
function getConf(Vue, path){
    if(path)
        return Vue.prototype.$http.xget(path) || {};
    
    if(!packageData)
        return
    
    path = packageData.scripts.serve ? "/env.conf.json" : "/static/env.conf.json"   
    return Vue.prototype.$http.xget(path) || {}
}

function inject405CodeHandle(that){
    that.$http.addStatusCodeHandle(res => {
        that.$msg.fail("请求的接口与页面不匹配或未配置权限")
        return
    }, 405);
}