import "jquery"
import "bootstrap"

import Vue from 'vue'
import Vuex from 'vuex'
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
//var conf = process.env.service;
var conf = process.env
var ssocfg =  ssoHttpConfig(conf.VUE_APP_API_URL, "localStorage", conf.SSO_WEB_HOST, conf.IDENT);

Vue.prototype.$sso = ssocfg.sso;
Vue.prototype.$http = ssocfg.http;

Vue.config.productionTip = false;
Vue.prototype.EnumUtility = new EnumUtility() // 枚举字典

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