import Vue from 'vue';
import Router from 'vue-router';
import menu from '@/pages/member/menu';
import ssocallback from '@/pages/member/sso.callback.vue';
import sysindex from '@/pages/system/sys.index.new.vue'
import sysfunc from '@/pages/system/func.index.vue'
import userindex from '@/pages/user/index.new.vue'
import roleindex from '@/pages/role/index.new.vue';
import roleauth from '@/pages/role/auth.vue';

Vue.use(Router);

export default new Router({
  mode: "history",
  routes: [
    {
      path: '/external/other',
      name: 'other',
      component: roleindex,
    },
    {
    path: '/',
    name: 'menu',
    component: menu,
    meta:{
      name:"用户权限系统"
    },
    children: [{
      path: 'user/index',
      name: 'userindex',
      component: userindex
    }, {
      path: "sys/index",
      name: "sysindex",
      component: sysindex
    }, {
      path: "sys/func",
      name: "sysfunc",
      component: sysfunc
    }, {
      path: 'user/role',
      name: 'userrole',
      component: roleindex
    }, {
      path: 'role/auth',
      name: 'roleauth',
      component: roleauth
    }
  ]
  },
    {
      path: '/ssocallback',
      name: 'ssocallback',
      component: ssocallback
    }
  ]
})
