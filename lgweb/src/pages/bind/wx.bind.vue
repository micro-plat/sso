<template>
    <div>
        <h2>绑定微信</h2>
        <div>
            <label >用户名</label> 
            <input  placeholder="输入用户名" maxlength="32" type="text" v-model="userName" />
        </div>
        <div>
            <label>密码</label>    
            <input  placeholder="输入密码" maxlength="32" type="password" v-model="password" />
        </div> 
        <div><input type="button" value="绑定" @click="bind" /></div>
        <div>{{message}}</div>
    </div>
</template>
<script>
  import {trimError} from '@/services/utils.js'
  export default {
    name: 'qrcode',
    data () {
      return {
          message: "",
          userName:"",
          password:"",
          flag : false,
      }
    },
    created() {
        document.title = "绑定微信用户";
    },
    methods:{
        bind() {
            //1:先验证用户名密码
            this.check();
            if (this.flag) {
                return;
            }
            //2:成功后跳转weixin地址
            this.validUser(this.userName, this.password);
        },

        check() {
            if(!this.userName || !this.password) {
                this.message = "用户名、密码不能为空"
                this.flag =  true;
                return;
            }
            this.flag = false;
            this.message = "";
        },

        validUser(userName, password) {
            this.$post("lg/user/check",{username:userName, password: password})
            .then(res => {
                var url = res.wxlogin_url + "?" + "appid=" + res.appid + "&state=" + res.state + "&redirect_uri=" +
                        //encodeURIComponent(process.env.service.wxcallbackhost + process.env.service.wxbindcallbackurl) +
                        encodeURIComponent(process.env.service.wxcallbackhost + process.env.service.wxlogincallbackurl + "/bind") +
                        "&response_type=code&scope=snsapi_base#wechat_redirect";            
                window.location.href = url;
            })
            .catch(err => {
                switch(err.response.status) {
                    case 406:
                        this.message = trimError(err);
                        break;
                    default:
                        this.message = "系统繁忙,请稍后再绑定";
                }
            });
        }
    }
  }
</script>