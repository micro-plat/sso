import Vue from 'vue';
import Router from 'vue-router';
import menu from '@/pages/member/menu';
//import login from '@/pages/member/login';
import login from '@/pages/member/login.jump.vue';
import ssocallback from '@/pages/member/sso.callback.vue';
import check from '@/pages/member/check';
import sysindex from '@/pages/system/sys.index.new.vue'
import sysfunc from '@/pages/system/func.index.vue'
import userindex from '@/pages/user/index.new.vue'
// import userindex from '@/pages/user/index';
import changepassword from '@/pages/user/change.pwd';
import changeInfo from '@/pages/user/change.info';
import userBind from '@/pages/user/user.bind';
import roleindex from '@/pages/role/index.new';
import roleauth from '@/pages/role/auth';
import notifyRecords from '@/pages/notify/notify.records.vue';
import notifySettings from '@/pages/notify/notify.settings.vue';

Vue.use(Router);


export default new Router({
  mode: "history",
  routes: [{
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
      path: 'password',
      name: 'password',
      component: changepassword
    }, {
      path: 'userinfo',
      name: 'userinfo',
      component: changeInfo
    }, {
      path: 'user/role',
      name: 'userrole',
      component: roleindex
    }, {
      path: 'role/auth',
      name: 'roleauth',
      component: roleauth
    },{
      path: '/notify_records',
      name: 'notify_records',
      component: notifyRecords,
    },{
      path: '/notify_settings',
      name: 'notify_settings',
      component: notifySettings
    }
  ]
  },
    {
      path: '/member/login',
      name: 'loginss',
      component: login
    },
    {
      path: '/:ident/member/login',
      name: 'logins',
      component: login
    },
    {
      path: '/login',
      name: 'login',
      component: login
    },
    {
      path: '/ssocallback',
      name: 'ssocallback',
      component: ssocallback
    },
    {
      path: '/member/check',
      name: 'check',
      component: check
    },{
      path: '/user/bind',
      name: 'bind',
      component: userBind,
    }
  ]
})
