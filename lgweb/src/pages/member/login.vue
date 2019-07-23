<template>
  <div id="app">
    <login-with-up
      :copyright="copyright"
      :systemName="systemName"
      :conf="conf"
      @loginCall="loginsubmit"
      ref="loginItem">
    </login-with-up>
  </div>
</template>

<script>
  import VueCookies from 'vue-cookies'
  import loginWithUp from 'login-with-up';
  import {JoinUrlParams} from '@/services/common.js'
  import "@/services/md5.js"
  export default {
    name: 'app',
    data () {
      return {
        systemName: "能源业务中心运营管理系统",
        copyright:"四川千行你我科技有限公司Copyright© 2018 版权所有",
        conf:{loginNameType:"输入用户名",pwd:"输入密码"},
        callback:"",
        changepwd:0,
        ident: ""
      }
    },
    components:{ 
      loginWithUp
    },

    mounted(){
      document.title = "登录-能源业务中心运营管理系统";
      this.callback = this.$route.query.callback;
      this.changepwd = this.$route.query.changepwd;
      this.ident = this.$route.query.ident;
      
      VueCookies.remove("__jwt__")
    },

    methods:{
      loginsubmit(e){
        var req = {
          containkey: 0,
          ident: this.ident,
          password: $.md5(e.password),
          username:e.username
        }

        if (this.callback && this.ident) {
          req.containkey = 1
        }

        this.$post("lg/login/post", req)
          .then(res => {
            setTimeout(() => {

              if (this.changepwd == 1) {
                this.$router.push({ path: '/changepwd'});   
                return;
              }

              if (this.ident && this.callback) {
                window.location.href = JoinUrlParams(decodeURIComponent(this.callback),{code:res.data})
                return;
              }
              this.$router.push({ path: '/chose'});
            }, 300);
          })
          .catch(err => {
              switch (err.response.status) {
                case 400:
                case 401:
                case 423:
                case 405:
                case 415:
                  this.$refs.loginItem.showMsg(err.response.data.data);
                  break;
                default:
                  this.$refs.loginItem.showMsg("登录失败");
              }
          });
      }
    }
  }
</script>
