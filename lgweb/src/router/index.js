import Vue from 'vue';
import Router from 'vue-router';
import login from '@/pages/member/login';
import jump from '@/pages/member/jump.vue';
import changepwd from '@/pages/member/changepwd.vue';
import chose from '@/pages/system/chose.vue';
import errpage from '@/pages/system/errpage.vue';
import wxlgcallback from '@/pages/member/wx.login.callback.vue';
import qrcodelogin from '@/pages/member/qrcode.login.vue';
import refresh from '@/pages/system/refresh.token.vue'

import wxbind from '@/pages/bind/wx.bind.vue';
import wxbindcallback from '@/pages/bind/wx.bind.callback.vue';


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
      path: '/login/:ident?',
      name: 'login',
      component: login
    },
    {

      path: '/jump/:ident?',
      name: 'jump',
      component: jump
    },
    {

      path: '/chose',
      name: 'chose',
      component: chose
    },
    {
      path: '/changepwd/:ident?',
      name: 'changepwd',
      component: changepwd
    },
    {
      path: '/qrcodelogin',
      name: 'qrcodelogin',
      component: qrcodelogin
    },
    {
      path: '/wxlgcallback/:type',
      name: 'wxlgcallback',
      component:  wxlgcallback 
    },
    {
      path: '/wxbind',
      name: 'wxbind',
      component: wxbind
    },
    {
      path: '/wxbindcallback',
      name: 'wxbindcallback',
      component: wxbindcallback
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
