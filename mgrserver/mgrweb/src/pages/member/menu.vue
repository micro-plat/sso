<template>
  <div id="app">
    <nav-menu
      :menus="menus"
      :copyright="copyright"
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
        copyright: new Date().getFullYear() + " admin-web", //版权信息
        themes: "", //顶部左侧背景颜色,顶部右侧背景颜色,右边菜单背景颜色
        menus: [{}],  //菜单数据
        systemName: "用户权限系统",  //系统名称
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
      this.userinfo = JSON.parse(localStorage.getItem("userinfo"));
    },
    methods:{
      pwd(){
        this.$sso.changePwd();
      },
      signOutM() {
        this.$sso.signOut();
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
         this.items = this.$sso.transformSysInfo(res);
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
