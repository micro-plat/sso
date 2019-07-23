<template>
  <div class="main">
    <div class="content">
      <i><img src="../../img/wrong.png"></i>
      <div class="message">{{msg}}</div>
      <div class="btn"><button @click="signOut">登录</button></div>
    </div>
  </div>
</template>

<script>
   import {JoinUrlParams} from '@/services/common'
  export default {
    name: 'errpage',
    data () {
      return {
          msg: "",
          errorType: 0
      }
    },

    mounted(){
      document.title = "登录出错-能源业务中心运营管理系统";
      if (this.$route.query.type) {
          this.errorType = this.$route.query.type
      }
      this.showMsg();
    },

    methods:{
        showMsg() {
           switch (this.errorType) {
               case 0 :
                   this.msg = '登录出错了';
                   break;
                case 1:
                    this.msg = '没有相关的权限,请联系管理员';
                    break;
                case 2:
                    this.msg = '用户被禁用或者被锁,请联系管理员';
                    break;
           }
        },
        signOut() {
            this.$router.push({ path: '/login'});
        }
    }
  }
</script>

<style>
.main {
  margin-top:200px;
}
.content {
    width: 300px;
    margin: 0 auto;
}
.message {
  font-size:30px;
  text-align: center;
  margin-top: 30px;
  color :#AAAAAA;
  font-weight: bold;
}
.btn {
  width:300px;
  margin:0 auto;
  text-align: center;
}
.btn button{
	  width: 140px;
    padding: 14px 0;
    font-size: 20px;
    color: #fff;
    background-color: #f4286e;
    border-radius: 10px;
    border: 0;
    margin-top: 40px;
}
</style>
