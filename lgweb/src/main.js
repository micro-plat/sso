// The Vue build version to load with the `import` command
// (runtime-only or standalone) has been set in webpack.base.conf with an alias.
import "jquery"
import "bootstrap"

import Vue from 'vue'
import Vuex from 'vuex'
import App from './App'
import router from './router'

//import VeeValidate from 'vee-validate';
import store from './store'
//import 'vue-tree-halower/dist/halower-tree.min.css'
//import VTree from 'vue-tree-halower'
//import uploader from 'vue-simple-uploader'
//Vue.use(uploader);

import {
    post,
    fetch,
    patch,
    put,
    del
} from './services/http'
import VueCookies from 'vue-cookies'
Vue.use(VueCookies);

import ElementUI from 'element-ui';
import 'element-ui/lib/theme-chalk/index.css';
Vue.use(ElementUI);



//定义全局变量
Vue.prototype.$post = post;
Vue.prototype.$fetch = fetch;
Vue.prototype.$patch = patch;
Vue.prototype.$put = put;
//Vue.prototype.$del = del;


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
