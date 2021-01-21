import "jquery"
import "bootstrap"

import Vue from 'vue'
import App from './App'


import VeeValidate, { Validator } from 'vee-validate';
import store from './store'
import 'vue-tree-halower/dist/halower-tree.min.css'
import VTree from 'vue-tree-halower'

import uploader from 'vue-simple-uploader'
Vue.use(uploader);
import ElementUI from 'element-ui';
import 'element-ui/lib/theme-chalk/index.css';
Vue.use(ElementUI);
const config = { fieldsBagName: 'vee-fields' }

import VueCookies from 'vue-cookies'
Vue.use(VueCookies);
Vue.use(VTree);
Vue.use(VeeValidate, config);


import utility from './services'
Vue.use(utility, "../static/env.conf.json");


Vue.prototype.$enum.callback(async function(type){
  var url = "/dds/dictionary/get";
  var data = await Vue.prototype.$http.get(url, { dic_type: type });
  console.log("dictionary.data:", type, data);
  return data;
});

Vue.prototype.$http.setBaseURL(Vue.prototype.$env.Conf.apiURL);

//Vue.prototype.$http.setEnableHeader(true);

Vue.prototype.$http.addStatusCodeHandle(res => {
  console.log("addStatusCodeHandle:403", res);
  var url = (res.headers || {}).location ||""; 
  if(!url){
    url = this.$env.Conf.loginWebHost + "/sso/jump?returnurl=";
  }

  url =url + encodeURIComponent(document.URL);
  console.log("redirect:url", url);
  window.location = url ;

  //return new Error("请补充注册中心auth/jwt的AuthURL配置");
}, 403);

import router from './router';
Vue.config.productionTip = false;
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

 
