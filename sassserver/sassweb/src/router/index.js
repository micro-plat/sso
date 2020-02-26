import Vue from 'vue';
import Router from 'vue-router';
import menu from '@/pages/member/menu';
import login from '@/pages/login/login.vue';
import changepwd from '@/pages/manage/change.pwd.vue';
import userindex from '@/pages/user/index.new.vue'
import roleindex from '@/pages/role/index.new.vue';
import roleauth from '@/pages/role/auth.vue';

Vue.use(Router);

export default new Router({
  mode: "history",
  routes: [
    {
      path: '/login',
      name: 'login',
      component: login
    },
    {
      path: '/changepwd',
      name: 'changepwd',
      component: changepwd
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
    }, 
    {
      path: 'user/role',
      name: 'userrole',
      component: roleindex
    }, {
      path: 'role/auth/:id',
      name: 'roleauth',
      component: roleauth
    },
  ]
  }
  ]
})
