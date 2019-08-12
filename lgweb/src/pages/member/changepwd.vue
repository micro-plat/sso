<template>
    <div class="swipercontiner">
        <h1 class="lg-title">能源业务中心运营管理系统</h1>
        <div>
            <div class="sing">
                <div class="title_text">修改密码</div>
                <div class="tips_text">
                    <i><img class="tips_image" src="../../img/password.png"></i>
                    <span>原密码</span>
                </div>
                <input class="tips-input" type="password" v-model="expassword"  placeholder="请输入原密码">
                <div class="tips_text">
                    <i><img class="tips_image" src="../../img/password.png"></i>
                    <span>新密码</span>
                </div>
                <input class="tips-input" type="password" v-model="password1" placeholder="请输入新密码">
                <div class="tips_text">
                    <i><img class="tips_image" src="../../img/password.png"></i>
                    <span>确认新密码</span>
                </div>
                <input class="tips-input" @keyup.enter="changePwd" type="password" v-model="password2" placeholder="请再次输入新密码">
                
                <div style="font-size:14px;color:#F7296F;">
                    {{errorMsg}}
                </div>
                <div class="but">
                    <span><button type="button" @click="changePwd"  class="btn blue-btn">确定</button></span>
                    <span><button style="background: #b4b4b4;" type="button" @click="signOut" class="btn blue-btn">取消</button></span>
                </div>
            </div>
        </div>
        <div class="footer"><p>四川千行你我科技有限公司Copyright© 2018 版权所有</p></div>
    </div>
</template>

<script>
  import {trimError} from "@/services/utils"
  import VueCookies from 'vue-cookies'
  import {jumpLogin} from '@/services/utils'

  export default {
    name: 'pwd',
    data () {
      return {
        ident:"",
        expassword: "",
        password1: "",
        password2: "",
        errorMsg: ""
      }
    },
    created() {
        this.ident = this.$route.params.ident ? this.$route.params.ident : "";
        var isExists = VueCookies.isKey("__sso_jwt__");
        if(!isExists) {
            this.$router.push({path:jumpLogin(this.ident), query :{ changepwd: 1 }});
        }
    },

    mounted(){
      document.title = "修改密码-能源业务中心运营管理系统";
    },

    methods:{
      signOut() {
          this.$router.push({path:jumpLogin(this.ident)});
      },

      check() {
            if (!this.expassword) {
                this.errorMsg = '旧密码不能为空';
                return;
            }

            if (!this.password1 || !this.password2) {
                this.errorMsg = '新密码不能为空';
                return;
            }

            if (this.password1.trim() != this.password2.trim()) {
                this.errorMsg = '两个新密码不一致';
                return;
            }

            if (this.password1.length > 20 || this.password2.length > 20) {
                this.errorMsg = '密码长度不能超过20个字符';
                return;
            }
            this.errorMsg = '';
        },
        changePwd(){
            this.check();
            if (this.errorMsg) {
                return;
            }
            this.errorMsg = '';
            this.$post("lg/user/changepwd", {expassword:this.expassword.trim(), newpassword:this.password1.trim()})
                .then(res => {
                    this.errorMsg = "密码修改成功";
                     setTimeout(() => {
                         this.$router.push(jumpLogin(this.ident));
                     }, 1000);
                }).catch(err => {
                    switch (err.response.status) {
                        case 403:
                            this.$router.push({path:jumpLogin(this.ident), query :{ changepwd: 1 }});
                            break;
                        case 406:
                        case 400:
                            this.errorMsg = trimError(err);
                            break;
                        default:
                            this.errorMsg = "网络错误,请稍后再试";
                    }
                })
        }
    }
  }
</script>

<style>

.swipercontiner{ height:100%;}
body{font-family: "\9ED1\4F53";background:url(../../img/background.png); background-size: cover; font-size:12px; margin:0;padding:0;}
li{	list-style:none;}
.input{ border:none;}
.input{ border:none;font-family: "\9ED1\4F53"; width:100%;	}

.title{
	font-size: 60px;
    padding: 80px 0;
    font-weight: 500;
    text-align: center;
    font-weight: bold;
}
.list{
	width: 900px;
    margin: 0 auto;
}
.everyone{
    width: 45%;
    text-align: center;
    display: inline-grid;
    margin: 20px;
    float: left;
}
.list .icon{
	background-color: #fff;
	padding: 60px 0;
	border-top-left-radius: 10px;
    border-top-right-radius: 10px
}

.list .text{
	font-size: 30px;
	color: #fff;
    padding: 30px 0;
     background-color: rgba(0,14,13,0.5);
    border-bottom-left-radius: 10px;
    border-bottom-right-radius: 10px
}

.sing{
    width: 440px;
    margin:0 auto;
	margin-top:7%;
    background-color: rgba(0,29,59,0.7);
    border-radius: 40px;
    padding: 60px 100px;
    
}
.title_text{
	font-size: 22px;
    text-align: center;
    color: #fff;
    padding-bottom: 40px;
}
.tips_text{
	font-size: 14px;
    font-weight: 500;
	color: #FFFFFF;
	padding-bottom: 10px;
}

.tips_image {
    width: 11px;
    height: 11px;
    margin-right: 4px;
}

.sing input{
	width: 92%;
    padding: 15px;
    font-size: 14px;
    color: #333;
    margin-bottom: 24px;
    border:0
}
.but button{
	width: 36%;
    padding: 16px 0;
    font-size: 16px;
    color: #fff;
    background-color: #F7296F;
    border: none;
    margin: 0 10px;
    border-radius: 10px;
}
.but{
	text-align: center;
	margin-top: 16px;
}

.footer{
    margin: 4vw .3vw 2vw;
}

.footer p {
    font-size: 14px;
    color: #fff;
    letter-spacing: 2px;
    text-align: center;
    line-height: 1.8;
}

.lg-title {
    font-size: 2.8vw;
    color: #fff;
    letter-spacing: 3px;
    text-align: center;
    margin: 3vw 1vw;
    font-family: Josefin Sans,sans-serif;
}

 ::-webkit-input-placeholder {
    color: #333;
  }

  :-moz-placeholder {
    /* Firefox 18- */
    color: #333;
  }

  ::-moz-placeholder {
    /* Firefox 19+ */
    color: #333;
  }

  :-ms-input-placeholder {
    color: #333;
  }

</style>
