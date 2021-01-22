<template>
  <div id="app">
    <nav-menu
      :menus="menus"
      :copyright="copyright"
      :copyrightcode="copyrightcode"
      :themes="themes"
      :logo="logo"
      :systemName="systemName"
      :userinfo="userinfo"
      :items="items"
      :pwd="pwd"
      :signOut="signOutM"
      ref="NewTap"
    >
    </nav-menu>
  </div>
</template>

<script>
  import navMenu from 'nav-menu'; // 引入
  export default {
    name: 'app',
    data () {
      return {
        logo: "",
        copyright: (this.$env.Conf.companyRight||"") + "Copyright©" + new Date().getFullYear() +"版权所有",
        copyrightcode: this.$env.Conf.companyRightCode ,
        themes: "", //顶部左侧背景颜色,顶部右侧背景颜色,右边菜单背景颜色
        menus: [{}],  //菜单数据
        systemName: "",  //系统名称
        userinfo: {name:'',role:"管理员"},
        indexUrl: "/user/index",
        items:[]
      }
    },
    components:{ //注册插件
      navMenu
    },
    created(){
      this.getMenu();
      this.getSystemInfo();
    },
    mounted(){
      this.setDocmentTitle();
      var userinfo = localStorage.getItem("userinfo")
      if(userinfo){
        this.userinfo = JSON.parse(userinfo);
      }
    },
    methods:{
      pwd(){
        this.$http.clearAuthorization();
        if(this.$env.Conf.cookieName){
          VueCookies.remove(this.$env.Conf.cookieName);
        }
        window.location.href = this.$env.Conf.loginWebHost + "/" + this.$env.Conf.ident + "/changepwd";
      },
      signOutM() {
        this.$http.clearAuthorization();
        var logouturl="";//如果想退出后跳转的地址，请设置值
        var returnURL = window.location.href;
        var redirectURL = "?returnurl="+returnURL;
        if (logouturl){
          redirectURL = "?logouturl="+logouturl;
        }
        window.location  = this.$env.Conf.loginWebHost+"/"+this.$env.Conf.ident+"/login"+redirectURL;
      },
      getMenu(){
        this.$http.get("/sso/member/menus/get")
          .then(res => {
            this.menus = res;
            this.$refs.NewTap.open("用户管理", this.indexUrl);
            this.getUserOtherSys();
          })
          .catch(err => {
            console.log(err)
          });
      },
      //获取系统的相关数据
      getSystemInfo() {
        this.$http.get("/sso/system/info/get")
        .then(res => {
          this.themes = res.theme;
          this.systemName = res.name;
          this.logo = res.logo;
          this.setDocmentTitle();
          
        }).catch(err => {
          console.log(err);
        })
      },
      //用户可用的其他系统
      getUserOtherSys() {
        this.$http.get("/sso/member/systems/get")
        .then(res => {
            this.items = (function (systems) {
              if (!systems || !systems.length) {
                  return []
              }
              var items = [];
              systems.forEach(element => {
                  items.push({
                    name: element.name,
                    path: element.index_url.substr(0, element.index_url.lastIndexOf("/")),
                    type: "blank"
                  })
              });
              return items;
          })(res);
        })
        .catch(err => {
          console.log(err);
        })
      },
      setDocmentTitle() {
        document.title = this.systemName;
      }
    
    }
  }
</script>

<style scoped>

</style>
