<template>
  <div class="app" id="app" :class="[layoutClass?layoutClass:'app-header-fixed']">
       
       <!-- / navbar 头部-->
    <div class="app-header navbar">
      <!-- navbar header -->
      <div class="navbar-header" v-bind:class="[navbarHeaderClass?navbarHeaderClass:'bg-info']">
        <button class="pull-right visible-xs dk" data-toggle="class:show" data-target=".navbar-collapse">
          <i class="glyphicon glyphicon-cog"></i>
        </button>
        <button class="pull-right visible-xs" data-toggle="class:off-screen" data-target=".app-aside" ui-scroll="app">
          <i class="glyphicon glyphicon-align-justify"></i>
        </button>
        <!-- brand -->
        <a href="#/" class="navbar-brand text-lt">
          <!-- <i class="fa fa-btc"></i> -->
          <img src="/static/img/lt.png" alt="." class="hide"/>
          <span class="hidden-folded m-l-xs">
                        <img v-bind:src="`/static/img/`+sys.logo"/>
                        </span>
        </a>
        <!-- / brand -->
      </div>
      <!-- / navbar header -->

      <!-- navbar collapse -->
      <div class="collapse pos-rlt navbar-collapse box-shadow"
           v-bind:class="[navbarCollapseClass?navbarCollapseClass:'bg-info']">
        <!-- buttons -->
        <div class="nav navbar-nav hidden-xs">
          <a href="#" class="btn no-shadow navbar-btn" data-toggle="class:app-aside-folded" data-target=".app">
            <i class="fa fa-outdent fa-fw text"></i>
            <i class="fa fa-indent fa-fw text-active"></i>
          </a>
          <!-- <a href class="btn no-shadow navbar-btn" data-toggle="class:show" data-target="#aside-user">
              <i class="fa fa-user-o fa-fw"></i>
          </a> -->
        </div>

        <div class="title">{{sys.name}}</div>

        <!-- / buttons -->
        <!-- nabar right -->

        <ul class="nav navbar-nav navbar-right">
          <li class="dropdown">
            <a href="#" data-toggle="dropdown" class="dropdown-toggle" @click="showAlert()">
              <i class="iconfont icon-tixing fa-fw" aria-hidden="true"></i>
              <span class="visible-xs-inline">通知</span>
              <span class="badge badge-sm up bg-danger pull-right-xs" v-if="notifies.count>0">{{notifies.count}} </span>

            </a>
            <!-- dropdown -->
            <ul class="dropdown-menu animated fadeInRight s" :style="IsShowAlert ? styleShow:styleNone">
              <li>
                <a @click="goto('/notify_records')">
                  <span>历史消息</span>
                </a>
              </li>
              <li class="divider"></li>
              <li>
                <a @click="goto('/notify_settings')">
                  <span>消息配置</span>
                </a>
              </li>
            </ul>
          </li>
          <li class="dropdown">
            <a href="#" class="dropdown-toggle clear" @click="showSys()" >
                            <span class="thumb-sm avatar pull-right m-t-n-sm m-b-n-sm m-l-sm">
                <img src="/static/img/a0.jpg" alt="...">
                <i class="on md b-white bottom"></i>
              </span>
            </a>
            <!-- dropdown -->
            <ul class="dropdown-menu animated fadeInRight" :style="IsShowSys ? styleShow:styleNone">
              <li class="wrapper b-b m-b-sm bg-light m-t-n-xs">
                <div>
                  {{user.user_name}} --- {{user.role_name}}
                </div>
              </li>
              <li>
                <a @click="goto('/change_info')">
                  <i class="fa fa-user-circle fa-fw fa-lg text-primary"></i>
                  <!--<span class="badge bg-danger pull-right "> {{user.profile_percent}}%</span>-->
                  <span>个人资料</span>
                </a>
              </li>
              <li>
                <a ui-sref="app.page.profile" @click="goto('/change_passwd')">
                  <i class="fa fa-cog fa-fw fa-lg text-primary"></i>
                  修改密码
                </a>
              </li>
              <li v-if="isDev()">
                <a ui-sref="app.page.profile" @click="test()">
                  <i class="fa fa-cog fa-fw fa-lg text-primary"></i>
                  开发者专用
                </a>
              </li>
              <li class="divider"></li>

              <li class="divider"></li>
              <!--系统-->
              <li v-for="(v,k) in sysList" :key="k" v-if="v.ident != ''">
                <a @click="changeSys(v)">
                  <i class="fa fa-check fa-fw fa-lg text-success" v-if="sysActive(v.ident)"></i>
                  <i class="fa fa-fw fa-lg" v-if="!sysActive(v.ident)"></i>
                  {{v.name}}
                </a>
              </li>
            </ul>
            <!-- / dropdown -->
          </li>
        </ul>
        <ul class="nav navbar-nav navbar-right">
          <li class="dropdown visible-sm visible-md">
            <a href="#" data-toggle="dropdown" class="dropdown-toggle clear">
              <span class="hidden-sm hidden-md">菜单</span> <b class="caret"></b>
            </a>
            <!-- dropdown -->
            <ul class="dropdown-menu animated fadeInRight w">
              <li v-for="(v,k) in menus" :key="k">
                <span class="badge bg-danger pull-right">1</span>
                <a ui-sref="app.page.profile" @click="selectMenu(v)">{{v.name}}</a>
              </li>
            </ul>
            <!-- / dropdown -->
          </li>
        </ul>


        <!-- search form -->
        <!--<form class="navbar-form navbar-form-sm navbar-right shift" ui-shift="prependTo" data-target=".navbar-collapse" role="search">-->
        <!--<div class="form-group">-->
        <!--<div class="input-group">-->
        <!--<input type="text" typeahead="state for state in states | filter:$viewValue | limitTo:8" class="form-control input-sm bg-light no-border rounded padder" placeholder="搜索..." v-model="searchWord" v-on:keyup="searchMenu" @focus="searchMenu" >-->
        <!--<ul class="dropdown-menu ng-isolate-scope" style="display: block; top: 30px; left: 0px;" v-if="resMenu.length > 0">-->

        <!--<li class="ng-scope" @click="open(v.path)" v-for="(v,k) in resMenu" :key="k">-->
        <!--<a class="ng-scope ng-binding" >{{v.name}}</a>-->
        <!--</li>-->
        <!--</ul>-->
        <!--<span class="input-group-btn">-->
        <!--<button type="submit" class="btn btn-sm bg-light rounded"><i class="fa fa-search"></i></button>-->
        <!--</span>-->
        <!--</div>-->
        <!--</div>-->
        <!--</form>-->
        <!-- / search form -->
        <ul class="nav navbar-nav navbar-right" v-for="(v,k) in reverseMenus" :key="k" v-show="reverseMenusNum > 1">
          <li class="dropdown visible-lg">
            <a class="dropdown-toggle clear" @click="selectMenu(v)">
              <span class="hidden-sm hidden-md">{{v.name}}</span>
            </a>
          </li>
        </ul>
        <!-- / navbar right -->

      </div>
      <!-- / navbar collapse -->
    </div>


    <!-- / navbar 头部-->

    <!-- menu -->
    <div class="app-aside hidden-xs" v-bind:class="[appAsideClass?appAsideClass:'']">
      <div class="aside-wrap">
        <div class="navi-wrap">
          <!-- user -->
          <div class="clearfix hidden-xs text-center hide" id="aside-user">
            <div class="dropdown wrapper">
              <a ui-sref="app.page.profile">
                                <span class="thumb-lg w-auto-folded avatar m-t-sm">
                  <img src="/static/img/a0.jpg" class="img-full" alt="...">
                </span>
              </a>
              <a href="#" data-toggle="dropdown" class="dropdown-toggle hidden-folded">
                                <span class="clear">
                  <span class="block m-t-sm">
                    <strong class="font-bold text-lt">John.Smith</strong>
                    <b class="caret"></b>
                  </span>
                                <span class="text-muted text-xs block">Art Director</span>
                                </span>
              </a>
              <!-- dropdown -->
              <ul class="dropdown-menu animated fadeInRight w hidden-folded">
                <li class="wrapper b-b m-b-sm bg-info m-t-n-xs">
                  <span class="arrow top hidden-folded arrow-info"></span>
                  <div>
                    <p>300mb of 500mb used</p>
                  </div>

                </li>
                <li>
                  <a href>Settings</a>
                </li>
                <li>
                  <a ui-sref="app.page.profile">Profile</a>
                </li>
                <li>
                  <a href>
                    <span class="badge bg-danger pull-right">3</span> Notifications
                  </a>
                </li>
                <li class="divider"></li>
                <li>
                  <a ui-sref="access.signin">Logout</a>
                </li>
              </ul>
              <!-- / dropdown -->
            </div>
            <div class="line dk hidden-folded"></div>
          </div>
          <!-- / user -->




          <!-- nav -->
          <!-- 左边导航 -->
          <nav ui-nav class="navi">
            <ul class="nav visible-sm">
              <li class="hidden-folded padder m-t m-b-sm text-muted text-xs">
                <span translate="aside.nav.your_stuff.YOUR_STUFF">快捷菜单</span>
              </li>
              <li v-for="(short,index) in shortcuts" :key="short.id" :class="index==active?'active':''">
                <a @click="open(short.path)">
                                       <span class="pull-right text-muted">
                                          <i class="fa fa-times fa-angle-right text" @click="del(short)"></i>
                                        </span>
                  <!-- <i :class="index==active?'fa fa-angle-right text-info-lter':'fa fa-angle-right'"></i> -->
                  <i :class="index==active? short.icon+' text-info-lter':short.icon"></i>
                  <span :class="index==active?'text-info-lter':''">{{short.name}}</span>
                </a>

              </li>
            </ul>
            <ul class="nav" v-for="l1 in menus" :key="l1.id" v-if="l1.id == showMenu">
              <!--<li class="line dk hidden-folded" ></li>-->

              <li class="hidden-folded padder m-t m-b-sm text-muted text-xs qx-top-name">
                <span>{{l1.name}}</span>
              </li>
              <li v-for="l2 in l1.children" :key="l2.id" :class="l2.is_open==1 ? 'active' : ''">
                <a href class="auto">
                                  <span class="pull-right text-muted">
                                    <i class="fa fa-fw fa-angle-right text"></i>
                                    <i class="fa fa-fw fa-angle-down text-active"></i>
                                  </span>
                  <i :class="l2.icon"></i>
                  <span class="font-bold">{{l2.name}}</span>
                </a>
                <ul class="nav nav-sub dk">
                  <li class="nav-sub-header">
                    <a href>
                      <span>{{l2.name}}</span>
                    </a>
                  </li>

                  <li v-for="l3 in l2.children" :key="l3.id"
                      :class="{'select' : l3func(l3.id),'selectcoupon':l3coupon(l3.id)}" @click="link(l3)">
                    <a :class="{'qx-third-hover':l3defaultfunc(),'qxcoupon-third-hover':l3couponfunc()}">
                      <span>{{l3.name}}</span>
                    </a>
                  </li>

                </ul>
              </li>
            </ul>
          </nav>
          <!-- nav -->

          <!-- aside footer -->
          <!-- / aside footer -->
        </div>
      </div>
    </div>
    <!-- / menu -->

    <!-- content -->
    <div class="app-content">
      <!--<div ui-butterbar class="butterbar" :class="isload?'active':'hide'"><span class="bar"></span></div>-->
      <a href class="off-screen-toggle hide" data-toggle="class:off-screen" data-target=".app-aside"></a>
      <div class="app-content-body fade-in-up">
        <!-- COPY the content from "tpl/" -->
        <!-- <div class="bg-light lter b-b wrapper-md" v-show="iframeState">
        <h1 class="m-n font-thin h3">{{activeMenu.name}}</h1>
        </div> -->
        <!-- qx-add-tabbox -->
        <div class="visible-md visible-lg">
          <el-tabs v-model="editableTabsValue2" type="card" closable @tab-remove="removeTab" @tab-click="changeTab">
            <el-tab-pane
              v-for="item in editableTabs2"
              :key="item.content.path"
              :label="item.title"
              :name="item.name"
              :parent="item.content.parent"
            >
            </el-tab-pane>
          </el-tabs>
        </div>
        <iframe v-if="iframeState" id="show-iframe" :path="path" :src="currentPageURL" frameborder="no" border="0"
                marginwidth="0" marginheight="0" style="width: 100%;overflow-x: auto;"></iframe>

        <router-view v-if="!iframeState" :path="path"></router-view>

        <!-- PASTE above -->
      </div>
    </div>
    <!-- /content -->

    <!-- aside right -->
    <div class="app-aside-right pos-fix no-padder w-md w-auto-xs bg-white b-l animated fadeInRight hide">
      <div class="vbox">
        <div class="wrapper b-b b-t b-light m-b">
          <a href class="pull-right text-muted text-md" data-toggle="class:show" data-target=".app-aside-right"><i
            class="icon-close"></i></a> Chat
        </div>
        <div class="row-row">
          <div class="cell">
            <div class="cell-inner padder">
              <!-- chat list -->
              <div class="m-b">
                <a href class="pull-left thumb-xs avatar"><img src="/static/img/a0.jpg" alt="..."></a>
                <div class="clear">
                  <div class="pos-rlt wrapper-sm b b-light r m-l-sm">
                    <span class="arrow left pull-up"></span>
                    <p class="m-b-none">Hi John, What's up...</p>
                  </div>
                  <small class="text-muted m-l-sm"><i class="fa fa-ok text-success"></i> 2 minutes ago</small>
                </div>
              </div>
              <div class="m-b">
                <a href class="pull-right thumb-xs avatar"><img src="/static/img/a0.jpg" class="img-circle"
                                                                alt="..."></a>
                <div class="clear">
                  <div class="pos-rlt wrapper-sm bg-light r m-r-sm">
                    <span class="arrow right pull-up arrow-light"></span>
                    <p class="m-b-none">Lorem ipsum dolor :)</p>
                  </div>
                  <small class="text-muted">1 minutes ago</small>
                </div>
              </div>
              <div class="m-b">
                <a href class="pull-left thumb-xs avatar"><img src="/static/img/a0.jpg" alt="..."></a>
                <div class="clear">
                  <div class="pos-rlt wrapper-sm b b-light r m-l-sm">
                    <span class="arrow left pull-up"></span>
                    <p class="m-b-none">Great!</p>
                  </div>
                  <small class="text-muted m-l-sm"><i class="fa fa-ok text-success"></i>Just Now</small>
                </div>
              </div>
              <!-- / chat list -->
            </div>
          </div>
        </div>
        <div class="wrapper m-t b-t b-light">
          <form class="m-b-none">
            <div class="input-group">
              <input type="text" class="form-control" placeholder="Say something">
              <span class="input-group-btn">
                <button class="btn btn-default" type="button">SEND</button>
              </span>
            </div>
          </form>
        </div>
      </div>
    </div>
    <!-- / aside right -->

    <!-- footer -->
    <div class="app-footer wrapper b-t bg-light">
      <span class="pull-right">1.0.0 <a href="#app" class="m-l-sm text-muted"><i
        class="fa fa-long-arrow-up"></i></a></span> &copy; 2018 All Copyright Reserved. <a
      href="http://www.qianxingniwo.com/" target="_blank" title="qxnw">中国石化四川石油</a>
    </div>
    <!-- / footer -->
  </div>
</template>

<script>
  import {loadNavbar, resizeIframe} from "@/services/nav";
  import {changeTheme} from "@/services/bg";
  import {Messenger} from "@/services/messenger";
  import VueCookies from 'vue-cookies';

  loadNavbar();

  export default {
    data() {
      return {
        IsShowSys:false,
        IsShowAlert:false,
        styleShow: {
          display:'block'
        },
        styleNone:{
          display:'none'
        },
        isDev:function(){
          return process.env.NODE_ENV == "development"
        },
        editableTabsValue2: '0',
        editableTabs2: [],
        tabIndex: 0,
        isload: true,
        navbarHeaderClass: "bg-primary",
        navbarCollapseClass: "bg-primary",
        appAsideClass: "bg-dark",
        layoutClass: "",
        loginPage: "/member/login",
        ident: "sso",
        sys: {
          name: "",
          logo: "logo.png"
        },
        currentPageURL: "",
        iframeState: true,
        notifies: {count: 0, details: []}, //{count:1,details:[{id:1,title:"下游卡单59笔"}]}
        user: {},
        shortcuts: [],
        populars: [],
        active: 0,
        activeMenu: {},
        menus: [],
        searchWord: '',
        resMenu: [],
        path: "",
        code: "",
        sysList: {},
        showMenu: "",
        a: "",
        l3Active: null,
        l3func: function (id) {
         // console.log(this.ident)
          if (this.ident != "coupon") {
            return this.l3Active == id
          } else {
            return false
          }

        },
        l3coupon(id) {

          if (this.ident == "coupon") {
            return this.l3Active == id
          } else {
            return false
          }
        },
        l3defaultfunc() {
          if (this.ident == "coupon") {
            return false
          } else {
            return true
          }
        },
        l3couponfunc() {
          if (this.ident == "coupon") {
            return true
          } else {
            return false
          }
        },
        sysActive: function (ident) {

          if (ident == this.sys.ident) {
            return true
          }
          return false
        }
      };
    },
    computed: {
      reverseMenus() {
        return this.menus.reverse();
      },
      reverseMenusNum() {
        return this.menus.length;
      }
    },
    mounted() {
      // resizeIframe();
      this.ident = this.$route.query.ident ? this.$route.query.ident : "sso";
      this.code = sessionStorage.getItem("code");
      this.loadMember();
      this.loadSysInfo();
      this.initMessenger();
      this.loadSysList();
      let self = this;
      setTimeout(function () {
        self.isload = false;
      }, 1000);
    },
    methods: {
      showSys(){
        //console.log(this.IsShowSys)
        this.IsShowSys = !this.IsShowSys
      },
      showAlert(){
        this.IsShowAlert =!this.IsShowAlert
      },
      addMsg(i) {

        i.path = process.env.service.webHost + i.path;
        i.parent = process.env.service.webHost + i.parent;
        // console.log("i",i)
        this.addTab(i);
        this.open(i.path)
        this.lightMenu(i.parent)
      },
      changeSys(v) {
        let user = JSON.parse(sessionStorage.getItem("user"));
        this.$post("/sso/login/code", {ident: v.ident, code: sessionStorage.getItem("code"), username: user.user_name})
          .then(res => {
            if (sessionStorage.getItem("__jwt__") != null) {
              sessionStorage.setItem("code", res.code);
              let host = "http://" + window.location.host;
              window.location.href = host + "/?ident=" + res.ident

              // this.$router.push('/?ident=' + this.login.ident);
            }
          })
          .catch(err => {
            if (err.response.status == 403) {
              this.$router.push('/member/login')
            } else {
              this.$notify({
                title: '错误',
                message: '网络错误,请稍后再试',
                type: 'error',
                offset: 50,
                duration: 2000,
              });
            }
          });
      },
      changeTab(tab, event) {
        this.open(tab.$vnode.data.key);
        // console.log("parent: ",typeof tab.$vnode.data.attrs.parent)
        if(tab.$vnode.data.attrs.parent){
          this.lightMenu(tab.$vnode.data.attrs.parent)
        }else{
          this.lightMenu(tab.$vnode.data.key)
        }

      },
      addTab(item) {
        let newTabName = ++this.tabIndex + '';
        let isadd = true;
        let activeName;

        this.editableTabs2.forEach((tab, index) => {
          if (tab.title == item.name) {
            isadd = false;
            activeName = tab
          }
        })
        if (isadd != true) {
          this.editableTabsValue2 = activeName.name;
          return
        }
        let p
        if (item.parent){
          p = item.parent
        }else{
          p = ""
        }
        this.editableTabs2.push({
          title: item.name,
          name: newTabName,
          content: item,
        });
        this.editableTabsValue2 = newTabName;
      },
      removeTab(targetName) {
        let tabs = this.editableTabs2;
        let activeName = this.editableTabsValue2;

        if (activeName === targetName) {
          tabs.forEach((tab, index) => {
            if (tab.name === targetName) {
              let nextTab = tabs[index + 1] || tabs[index - 1];
              let msg = {type: "close", data: this.substrDomain(tab.content.path)}
              this.notifyMessenger(JSON.stringify(msg));
              if (nextTab) {
                this.open(nextTab.content.path);
                activeName = nextTab.name;
              }
            }
          });
        }

        this.editableTabsValue2 = activeName;
        this.editableTabs2 = tabs.filter(tab => tab.name !== targetName);
      },
      selectMenu(item) {
        this.showMenu = item.id
      },
      loadSysList() {
        this.$fetch("/sso/user/query")
          .then(res => {
            this.sysList = res
            //console.log("系统列表",res)
          })
          .catch(err => {
            if (err.response.status == 403) {
              this.$router.push('/member/login')
            } else {
              this.$notify({
                title: '错误',
                message: '网络错误,请稍后再试',
                type: 'error',
                offset: 50,
                duration: 2000,
              });
            }
          });
      },
      searchMenu() {
        this.resMenu = [];
        this.menus.forEach((item, index) => {
          item.children.forEach((v, k) => {
            v.children.forEach((val, key) => {

              if (this.searchWord != "" && val.name.indexOf(this.searchWord) >= 0) {
                this.resMenu.push(val)
              }
            })
          })
        });
      },
      loadMember: function () {
        this.$fetch("/sso/member/query")
          .then(res => {
            this.user = res;
            sessionStorage.setItem("user", JSON.stringify(res))
          })
          .catch(err => {
            if (err.response.status == 403) {
              this.$router.push('/member/login')
            } else {
              this.$notify({
                title: '错误',
                message: '网络错误,请稍后再试',
                type: 'error',
                offset: 50,
                duration: 2000,
              });
            }
          });
      },
      loadSysInfo: function () {
        this.$fetch("/sso/sys/get", {ident: this.ident})
          .then(res => {
            this.sys = res;
            document.title = res.name;
            this.currentPageURL = res.login_url + "?code=" + this.code;
            if (this.currentPageURL.startsWith("http")) {
              this.iframeState = true;
            }
            if (res.id !== 0) {
              this.loginPage = "/" + this.ident + "/member/login/";
            }

            if (this.sys.theme) {
              let themes = this.sys.theme.split("|");
              if (themes.length == 3) {
                this.navbarHeaderClass = themes[0];
                this.navbarCollapseClass = themes[1];
                this.appAsideClass = themes[2];
              }
            }
            if (this.sys.layout) {
              this.layoutClass = this.sys.layout;
            }
            this.loadMenu();
          })
          .catch(err => {
          });
      },
      loadMenu: function () {
        this.$fetch("/sso/menu/get")
          .then(res => {
            //console.log("菜单： ",res)
            this.menus = res;
            for (var l1 of res) {
              for (var l2 of l1.children) {
                for (var l3 of l2.children) {
                  if (this.sys.index_url == l3.path) {
                    this.link(l3);
                  }
                }
              }
            }
            this.lightMenu("");
          })
          .catch(err => {
            if (err.response.status == 403) {

            } else {
              this.$notify({
                title: '错误',
                message: '网络错误,请稍后再试',
                type: 'error',
                offset: 50,
                duration: 2000,
              });
            }
          });
        this.$fetch("/sso/popular")
          .then(res => {
            this.populars = res;
          })
          .catch(err => {
            if (err.response.status == 403) {

            } else {
              this.$notify({
                title: '错误',
                message: '网络错误,请稍后再试',
                type: 'error',
                offset: 50,
                duration: 2000,
              });
            }
          });
      },
      goto: function (i) {
        this.iframeState = false;
        this.$router.push(i);
      },
      del: function (i) {
        this.shortcuts.splice(this.shortcuts.indexOf(i), 1);
      },
      open: function (path) {
        if (path.startsWith("http")) {
          //this.currentPageURL = path;
          let msg = {type: "path", data: this.substrDomain(path)}
          this.notifyMessenger(JSON.stringify(msg));
          this.iframeState = true;
          this.resMenu = [];
         // console.log("lightpath: ",path)
          this.lightMenu(path)  //切换左边选中
          return;
        }
        this.iframeState = false;
        this.$router.push(path);
        this.resMenu = [];
      },
      findMenu: function (i) {
        if (this.shortcuts.length == 0) {
          return null;
        }
        for (var one of this.shortcuts) {
          if (one != null && one.id == i.id) {
            return one;
          }
        }
        return null;
      },
      link: function (item) {
        if (item.id == this.l3Active) {
          return
        }
        var i = this.findMenu(item);
        if (!i) {
          if (this.shortcuts.length > 7) {
            this.shortcuts.splice(0, 1);
          }
          this.shortcuts.push(item);
          i = item;
        }

        this.active = this.shortcuts.indexOf(i);
        this.activeMenu = i;
        // console.log("当前链接："+i.path);
        this.lightMenu(i.path)
        this.path = i.path;
        this.addTab(i);
        //console.log("item",i)
        this.open(i.path);
      },
      lightMenu(path) {
        //console.log("lightMenu: ",path)
        if (path == "") {
          path = this.sys.index_url
        }
        for (var l1 of this.menus) {
          for (var l2 of l1.children) {
            for (var l3 of l2.children) {
              if (path == l3.path) {
                this.showMenu = l1.id;
                l2.is_open = 1;
                this.l3Active = l3.id;
              }
            }
          }
        }
      },
      initMessenger() {
        self = this;
        if (document.getElementById("show-iframe")) {
          this.messenger = new Messenger("UMS");
          this.messenger.addTarget(
            document.getElementById("show-iframe").contentWindow,
            "SUBSYSTEM"
          );
        }
        this.messenger.listen(function (msg) {
          let data = JSON.parse(msg);
         // console.log(data)
          if (data.func == "height"){
            var res = parseInt(data.value, 10);

            if (res < document.documentElement.clientHeight) {
              res = document.documentElement.clientHeight
            }
            document.getElementById("show-iframe").style.height = res + 'px'
          }else if (data.func == "page"){
            self.addMsg(data.value)
          }else if (data.func == "http") {
            self.$router.push("/member/login");
          }
        })
      },
      notifyMessenger: function (path) {
        this.messenger.targets["SUBSYSTEM"].send(path);
      },
      substrDomain: function (url) {
        var oldIndex = -1;
        for (var i = 0; i < 3; i++) {
          if (url.indexOf("/", oldIndex + 1) > -1) {
            oldIndex = url.indexOf("/", oldIndex + 1);
          }
        }
        // console.log(url.substr(oldIndex));
        return url.substr(oldIndex);
      },
      //退出登录,要跳转到sso登录页面(两边的jwt cookie都要清理)
      signOut() {
        console.log("ddddd");
        //VueCookies.remove("__jwt__");
        //window.location.href =  process.env.service.ssoWebHost + process.env.service.loginUrl;
      }
    },
    destroyed: function () {
      let menu_ids = [];
      let pids = [];
      this.shortcuts.forEach(function (v, k) {
        menu_ids.push(v.id);
        pids.push(v.parent);
      });
      if (menu_ids.length == 0) {
        return;
      }
      this.$post("/sso/popular", {
        menu_ids: menu_ids.join(","),
        pids: pids.join(",")
      })
        .then(res => {
        })
        .catch(err => {
        });
    }
  };
</script>


