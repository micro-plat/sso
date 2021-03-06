import "jquery"
import "bootstrap"
 
import Vue from 'vue'
import App from './App'
import router from './router'

import VeeValidate, {Validator} from 'vee-validate';
import store from './store'
import 'vue-tree-halower/dist/halower-tree.min.css'
import VTree from 'vue-tree-halower'
import uploader from 'vue-simple-uploader'
Vue.use(uploader);
import ElementUI from 'element-ui';
import 'element-ui/lib/theme-chalk/index.css';
Vue.use(ElementUI);
import {EnumUtility,EnumFilter} from 'qxnw-enum';

const config = {fieldsBagName: 'vee-fields'}

import VueCookies from 'vue-cookies'
Vue.use(VueCookies);
Vue.use(VTree);
Vue.use(VeeValidate,config);


import {ssoHttpConfig} from 'qxnw-sso';
var conf = window.globalConfig
var ssocfg =  ssoHttpConfig(conf.apiURL ||"", "localStorage", conf.loginWebHost, conf.ident);

Vue.prototype.$sso = ssocfg.sso;
Vue.prototype.$http = ssocfg.http;

Vue.config.productionTip = false;
Vue.prototype.EnumUtility = new EnumUtility(); // 枚举字典
Vue.prototype.EnumUtility.defaultCallback(function(type){
  debugger
  var result = []
  $.ajax({
    url: (window.globalConfig.apiURL ||"") + "/dds/dictionary/get",
    data:  { dic_type: type },
    async: false,
    // beforeSend: function (request) {
    //   request.setRequestHeader("__sso_jwt__", GetTocken())
    //   request.setRequestHeader("Content-Type", "application/x-www-form-urlencoded")
    // },
    type: "GET",
    success: function (data) {
      console.log("/dds/dictionary/get:",type,data);
      result = data;
    }
  });

  return result; 
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
  