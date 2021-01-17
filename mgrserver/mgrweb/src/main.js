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


import DateConvert from './services/date'
import DateFilter from './services/filter';
import env from './services/env'
import http from './services/http'
import senum from './services/enum'

Vue.use(http);
Vue.use(env);
Vue.use(senum);

window.globalConfig = {companyRight:""}

Vue.prototype.$enum.callback(async function(type){
  var url =  (window.globalConfig.apiURL || "") + "/dds/dictionary/get";
  var data = await Vue.prototype.$http.get(url,{ dic_type: type });
  console.log("dictionary.data:",type,data);
  return data;
})

Vue.prototype.$env.load(async function(){
  var data = await Vue.prototype.$http.get("/system/webconfig");
  console.log("webconfig.data:",data);
  window.globalConfig = data;
  return data;
});

Vue.prototype.$http.setEnableHeader(true);
Vue.prototype.$http.addStatusCodeHandle(res=>{
  console.log("addStatusCodeHandle:403",res)
  var url = (res.headers||{}).location;

  console.log("redirect:url",url)
 // window.location = url + encodeURIComponent(document.URL); 
  
  //return new Error("请补充注册中心auth/jwt的AuthURL配置");
},403);

import router from './router'
Vue.config.productionTip = false;
Vue.prototype.DateConvert = DateConvert //日期格式转换
// Vue.prototype.DateFilter = DateFilter
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
