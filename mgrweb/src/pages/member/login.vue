<template>
  <div id="app">
    <login-with-up
      :copyright="copyright"
      :systemName="systemName"
      :conf="conf"
      @loginCall="call"
      ref="loginItem"
    >

    </login-with-up>
  </div>
</template>

<script>
  import VueCookies from 'vue-cookies'
  import loginWithUp from 'login-with-up'; // 引入
  export default {
    name: 'app',
    data () {
      return {
        systemName: "用户权限系统",
        copyright:"四川千行你我科技有限公司Copyright© 2018 版权所有",
        conf:{loginNameType:"请输入邮箱或用户名",pwd:"输入密码"},   //输入框提示信息配置
      }
    },
    components:{ //注册插件
      loginWithUp
    },
    mounted(){
      document.title = "用户登录";
      VueCookies.remove("__jwt__")
    },
    methods:{
      call(e){
        //在这里获取数据进行登录
        e.ident = "sso";
        this.$fetch("/sso/login", e)
          .then(res => {
            this.$refs.loginItem.showMsg("登录中.....");

            setTimeout(()=>{

                this.$router.push("/")

            }, 500);
          })
          .catch(err => {
            if (err.response) {
              switch (err.response.status) {
                case 415:
                  this.$refs.loginItem.showMsg("不允许登录系统");
                  break;
                case 403:
                  this.$refs.loginItem.showMsg("用户名或密码错误");
                  break;
                case 423:
                  this.$refs.loginItem.showMsg("用户被锁定暂时无法登录");
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
