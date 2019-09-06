<template>
  <div class="main">
    <div class="content">
      <i><img style="margin: auto; width:300px" src="../../img/wrong.png"></i>
      <div class="message">{{msg}}</div>
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
          errorType: 0,
          errorTemplate:{
            0: "登录出错了,等会再试",
            901: "系统被禁用,暂时不能登录",
            902: "用户被锁定，暂时无法登录",
            903: "用户被禁用，暂时无法登录",
            906: "没有相应的系统权限,请联系管理员"          
          }
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
          this.msg = this.errorTemplate[this.errorType] || "登录出错,等会再试"
        }
    }
  }
</script>

<style>
.main {
  margin-top:200px;
}
.content {
    width: 60%;
    text-align: center;
    margin: 0 auto;
}
.message {
  font-size:20px;
  text-align: center;
  margin-top: 20px;
  color :#999;
  font-weight: bold;
}
.btn {
  width:300px;
  margin:0 auto;
  text-align: center;
}
.btn button{
	  width: 160px;
    padding: 14px 0;
    font-size: 16px;
    color: #fff;
    background-color: #f7296f;
    border: 0;
    margin-top: 24px;
    border-radius: 10px;
}
</style>
