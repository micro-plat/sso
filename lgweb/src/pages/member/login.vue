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
  export default {
    name: 'app',
    data () {
      return {
        systemName: "能源业务中心运营管理系统",
        copyright:"四川千行你我科技有限公司Copyright© 2018 版权所有",
        conf:{loginNameType:"输入用户名",pwd:"输入密码"},
        callback:"",
        changepwd:0
      }
    },
    components:{ 
      loginWithUp
    },

    mounted(){
      document.title = "登录-能源业务中心运营管理系统";
      this.callback = this.$route.query.callback;
      this.changepwd = this.$route.query.changepwd;
      
      VueCookies.remove("__jwt__")
    },

    methods:{
      loginsubmit(e){
        this.$post("lg/login/post", e)
          .then(res => {
            this.$refs.loginItem.showMsg("登录中.....");
            setTimeout(() => {
              if (this.callback) {
                window.location.href = JoinUrlParams(decodeURIComponent(this.callback),{code:res.data})
                return;
              }
              if (this.changepwd == 1) {
                this.$router.push({ path: '/changepwd'});   
                return;
              }
              this.$router.push({ path: '/chose',query: { code: res.data }});   
            }, 300);
          })
          .catch(err => {
            if (err.response) {
              switch (err.response.status) {
                case 401:
                case 423:
                case 403:
                  this.$refs.loginItem.showMsg(err.response.data.data);
                  break;
                case 415:
                  this.$refs.loginItem.showMsg("不允许登录系统");
                  break;
              }
            }else{
              this.$refs.loginItem.showMsg("不允许登录系统");
            }
          });
      }
    }
  }
</script>
