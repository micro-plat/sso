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
        copyright: (this.$env.conf.copyright.company||"") + "Copyright©" + new Date().getFullYear() +"版权所有",
        copyrightcode: this.$env.conf.copyright.code ,
        themes: "", //顶部左侧背景颜色,顶部右侧背景颜色,右边菜单背景颜色
        menus: [{}],  //菜单数据
        systemName: "",  //系统名称
        userinfo:{},
        items:[]
      }
    },
    components:{ //注册插件
      navMenu
    },
    created(){
     
    },
    mounted(){
      console.log("----------",this.$route.query)
      this.$auth.checkAuthCode(this)
      this.getMenu();
      this.getSystemInfo();

      this.setDocmentTitle();
      this.userinfo = this.$auth.getUserInfo()
    },
    methods:{
      pwd(){
        this.$http.clearAuthorization();

        var keys = this.$cookies.keys();
        for(var i in keys){
            this.$cookies.remove(keys[i]);
        }  
        var url = this.$env.conf.sso.host + "/"+ this.$env.conf.sso.ident + "/changepwd"
        window.location.href = url;
      },
      signOutM() {
        this.$auth.loginout();
      },
      getMenu(){
          this.$auth.getMenus(this).then(res=>{
            this.menus =res ;
            this.getUserOtherSys();
          });
      },
      //获取系统的相关数据
      getSystemInfo() { 
         this.$auth.getSystemInfo().then(res=>{
            this.themes = res.theme;
            this.systemName = res.name;
            this.logo = res.logo;
         })
      },
      //用户可用的其他系统
      getUserOtherSys() {
        this.$auth.getSystemList().then(res=>{
          this.items = res;
        }) 
      },
      setDocmentTitle() {
        document.title = this.$env.conf.name;
      }
    
    }
  }
</script>
