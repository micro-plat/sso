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

import utility from './services';
 
Vue.use(utility,false);

Vue.config.productionTip = false;
Vue.use(ElementUI);
Vue.use(VueCookies);

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
