// The Vue build version to load with the `import` command
// (runtime-only or standalone) has been set in webpack.base.conf with an alias.
import "jquery"
import "bootstrap"
 
import Vue from 'vue'
import App from './App'
import router from './router'
import store from './store'
import VueCookies from 'vue-cookies'
import ElementUI from 'element-ui';
import 'element-ui/lib/theme-chalk/index.css';

import env from './services/env'
import http from './services/http'
import senum from './services/enum'



Vue.use(http);
Vue.use(env);
Vue.use(senum);

Vue.config.productionTip = false;
Vue.use(ElementUI);
Vue.use(VueCookies);

Vue.prototype.$enum.callback(async function(type){
    var url =  (window.globalConfig.apiURL || "") + "/dds/dictionary/get";
    var data = await Vue.prototype.$http.get(url,{ dic_type: type });
    return data;
})

Vue.prototype.$env.load(async function(){
    var data = await Vue.prototype.$http.get("/system/webconfig");
    window.globalConfig = data;
    return data;
});


  /* eslint-disable no-new */
new Vue({
    el: '#vapp',
    store,
    router,
    components: {
        App
    },

    template: '<App/>'
});
