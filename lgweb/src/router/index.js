import Vue from 'vue';
import Router from 'vue-router';
import login from '@/pages/member/login';
import jump from '@/pages/member/jump.vue';

Vue.use(Router);

export default new Router({
  mode: "history",
  routes: [{
      path: '/login',
      name: 'login',
      component: login
    },
    {

      path: '/jump',
      name: 'jump',
      component: jump
    }
  ]
})
