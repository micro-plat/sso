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
import navMenu from "nav-menu"; // 引入
export default {
  name: "app",
  data() {
    return {
      logo: "",
      copyright:
        (this.$env.conf.copyright.company || "") +
        "Copyright©" +
        new Date().getFullYear() +
        "版权所有",
      copyrightcode: this.$env.conf.copyright.code,
      themes: "", //顶部左侧背景颜色,顶部右侧背景颜色,右边菜单背景颜色
      menus: [{}], //菜单数据
      systemName: "", //系统名称
      userinfo: {},
      items: [],
    };
  },
  components: {
    //注册插件
    navMenu,
  },
  created() {
  },
  mounted() {
    console.log("----------", this.$route.query);
    this.getMenu();
    this.getSystemInfo();
    this.userinfo = this.$sys.getUserInfo();
  },
  methods: {
    pwd() {
        this.$sys.changePwd();
    },
    signOutM() {
      this.$sys.logout();
    },
    getMenu() {
      this.$sys.getMenus().then((res) => {
        this.menus = res;
        this.getUserOtherSys();
        var cur = this.$sys.findMenuItem(res)
        this.$refs.NewTap.open(cur.name, cur.path);
      });
    },
    //获取系统的相关数据
    getSystemInfo() {
      this.$sys.getSystemInfo().then((res) => {
        this.themes = res.theme;
        this.systemName = res.name;
        this.logo = res.logo;
      });
    },
    //用户可用的其他系统
    getUserOtherSys() {
      this.$sys.getSystemList().then((res) => {
        this.items = res;
      });
    }
  },
};
</script>
