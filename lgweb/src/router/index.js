import Vue from 'vue';
import Router from 'vue-router';
import login from '@/pages/member/login';
import jump from '@/pages/member/jump.vue';
import changepwd from '@/pages/member/changepwd.vue';
import chose from '@/pages/system/choose.vue';
import errpage from '@/pages/system/errpage.vue';
import refresh from '@/pages/system/refresh.token.vue'

Vue.use(Router);

export default new Router({
  mode: "history",
  routes: [
    {
      path: '/',
      name: 'first',
      component: login
    },
    {
      path: '/:ident?/login',
      name: 'login',
      component: login
    },
    {

      path: '/:ident?/jump',
      name: 'jump',
      component: jump
    },
    {

      path: '/chose',
      name: 'chose',
      component: chose
    },
    {
      path: '/:ident?/changepwd',
      name: 'changepwd',
      component: changepwd
    },
    {
      path: '/errpage',
      name: 'errpage',
      component: errpage
    },
    {
      path: '/refresh',
      name: 'refresh',
      component: refresh
    }
  ]
})
