<template>
  <div id="app">
    <nav-menu
      :menus="menus"
      :copyright="copyright"
      :themes="themes"
      :logo="logo"
      :systemName="systemName"
      :headpic="headpic"
      :userinfo="userinfo"
      :pwd="pwd"
      :signOut="signOutM"
      ref="NewTap"
    >
    </nav-menu>
  </div>
</template>

<script>
  import navMenu from 'nav-menu'; // 引入
  import VueCookies from 'vue-cookies';
  import {signOut,changePwd} from '@/services/http.js'
  export default {
    name: 'app',
    data () {
      return {
        headpic: "http://sso2.100bm.cn:6888/static/img/a0.jpg",
        logo: "http://sso2.100bm.cn:6888/static/img/d663155de6dc8e060415bbcd891cb9d4.png",
        copyright: "2018 admin-web", //版权信息
        themes: "bg-danger|bg-danger|bg-dark light-danger", //顶部左侧背景颜色,顶部右侧背景颜色,右边菜单背景颜色
        menus: [{}],  //菜单数据
        systemName: "用户权限系统",  //系统名称
        userinfo: {name:'wule',role:"管理员"},
        indexUrl: "/user/index",
        dialogAddVisible:false,     //添加表单显示隐藏
      }
    },
    components:{ //注册插件
      navMenu
    },
    created(){
      this.getMenu();
    },
    mounted(){
      document.title = "用户权限系统";
      this.userinfo = JSON.parse(localStorage.getItem("userinfo"));
    },
    methods:{
      pwd(){
        changePwd();
      },
      signOutM() {
        signOut();
      },
      getMenu(){
        this.$fetch("/sso/menu")
          .then(res => {
            this.menus = res;

            /*
            //这是处理登录后的回调
            var oldPath = localStorage.getItem("beforeLoginUrl");
            localStorage.removeItem("beforeLoginUrl");

            //如果登录过的(不会跳转到sso登录,直接就会加载了)就要在当前地址中找
            var loginedPath = window.location.pathname;
            if (!oldPath) {
              oldPath = loginedPath;
            }

            if (oldPath  && oldPath != "/") {    
              var name = this.getOneMenuName(oldPath, res);
              if (name == "") {
                name = "未知";
              }
              this.$refs.NewTap.add(name, oldPath ,{});
            } else {
              this.$refs.NewTap.add("首页", this.indexUrl ,{});
            }
            */
            this.$refs.NewTap.add("首页", this.indexUrl ,{});
          })
          .catch(err => {
            console.log(err)
          });
      },

      //查询某个url对应的菜单
      getOneMenuName(url, menus) {
        for (var i = 0; i < menus.length; i++) { 
          for (var j = 0; j < menus[i].children.length; j++) {
            for (var k = 0; k < menus[i].children[j].children.length; k++) {
                if (menus[i].children[j].children[k].path == url) {
                  return menus[i].children[j].children[k].name;
                }
            }
          }
        }
      }

    }
  }
</script>

<style scoped>

</style>
