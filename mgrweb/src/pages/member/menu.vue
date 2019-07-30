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
      :signOut="signOut"
      ref="NewTap"
      @loadHttpPage=gotoLoadHttp
    >
    <iframe v-show="bdTokenUrl.indexOf('http://') == 0 || bdTokenUrl.indexOf('https://') == 0" 
        ref="bdIframe" id="bdIframe" 
        :src="frameUrl" 
        width="100%" 
        height="100%" 
        frameborder="0" 
        allowtransparency="true" 
        allowfullscreen="true" 
      ></iframe>
      <router-view v-show="bdTokenUrl.indexOf('http://') != 0 && bdTokenUrl.indexOf('https://') != 0" />
    </nav-menu>
  </div>
</template>

<script>
  import navMenu from 'nav-menu'; // 引入
  import VueCookies from 'vue-cookies';
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
        bdTokenUrl: "",
        frameUrl: "",
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
      this.setIfrema()
    },
    methods:{
      setIfrema(){
        var oIframe = this.$refs.bdIframe
        var deviceWidth = document.documentElement.clientWidth;
        var deviceHeight = document.documentElement.clientHeight;
        oIframe.style.width = (Number(deviceWidth)-220) + 'px'; //数字是页面布局宽度差值
        oIframe.style.height = (Number(deviceHeight)-150) + 'px'; //数字是页面布局高度差
      },
      pwd(val){
        VueCookies.remove("__jwt__");
        var config = process.env.service;
        window.location.href = 
            config.ssoWebHost + config.changePwd;
      },
      signOut() {
        VueCookies.remove("__jwt__");
        var config = process.env.service;
        window.location.href = config.ssoWebHost + config.loginUrl;
      },

      getMenu(){
        console.log("调用菜单数据")
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
      close(v){
        this.$refs.NewTap.closeTab(v);
      },
      gotoLoadHttp(url){
        if (this.indexUrl && this.indexUrl == url){
          return;
        }
        if (url.indexOf('http://') == 0 || url.indexOf('https://') == 0) {
          this.frameUrl = url
        }
        this.bdTokenUrl = url; 
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
