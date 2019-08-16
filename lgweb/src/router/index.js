import Vue from 'vue';
import Router from 'vue-router';
import login from '@/pages/member/login';
import jump from '@/pages/member/jump.vue';
import changepwd from '@/pages/member/change.password.vue';
import choose from '@/pages/system/choose.vue';
import errpage from '@/pages/system/errpage.vue';
import refresh from '@/pages/system/refresh.token.vue'

Vue.use(Router);

export default new Router({
  mode: "history",
  routes: [
    {
      path: '/',
      name: 'home',
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

      path: '/choose',
      name: 'choose',
      component: choose
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
