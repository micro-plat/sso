<template>
    <div class="main">
        <div class="form-group">
            <label>请输入原密码</label>
            <input type="password" class="form-control" name="expassword" v-model="expassword" >
        </div>
        <div class="form-group">
            <label>请输入新密码</label>
            <input type="password" name="password1" class="form-control" v-model="password1" >
        </div>
        <div class="form-group">
            <label>请再次输入新密码</label>
            <input type="password" name="password2" class="form-control" v-model="password2" >
        </div>
        <div>
            <button class="btn btn-sm btn-primary" @click="signOut">取 消</button>
            <button class="btn btn-sm btn-danger"  @click="changePwd">确 定</button>
        </div>
    </div>
</template>

<script>
  import VueCookies from 'vue-cookies'
  export default {
    name: 'pwd',
    data () {
      return {
        expassword: "",
        password1: "",
        password2: "",
        errorMsg: ""
      }
    },
    created() {
        var isExists = VueCookies.isKey("__jwt__");
        if(!isExists) {
            this.$router.push({path:"/login", query :{ changepwd: 1 }});
        }
    },

    mounted(){
      document.title = "修改密码-能源业务中心运营管理系统";
    },

    methods:{

      signOut() {
          this.$router.push({path:"/login"});
      },

      check() {
            if (this.expassword == '') {
                this.errorMsg = '旧密码不能为空';
                return;
            }
            if (this.password1 == '') {
                this.errorMsg = '新密码不能为空';
                return;
            }
            if (this.password1 != this.password2) {
                this.errorMsg = '两个新密码不一致';
                return;
            }

            if (this.password1.length > 20 || this.password2.length > 20) {
                this.errorMsg = '密码长度不能超过20个字符';
            }
        },
        changePwd(){
            this.check();
            if (this.errorMsg) {
                 this.$alert(this.errorMsg, '提示', {
                    confirmButtonText: '确定'
                });
                return;
            }
            this.$post("lg/user/changepwd", {expassword:this.expassword, newpassword:this.password1})
                .then(res => {
                    this.$alert("密码修改成功", '提示', {confirmButtonText: '确定'});
                    this.$router.push("/login");
                }).catch(err => {
                    if (err.response) {
                        switch (err.response.status) {
                            case 403:
                                this.$router.push({path:"/login", query :{ changepwd: 1 }});
                                break;
                            case 406,400:
                                console.log(err.response.data.data);
                                this.$alert(err.response.data.data, '提示', {confirmButtonText: '确定'});
                                break;
                        }
                    } else{
                        this.$alert("网络错误,请稍后再试", '提示', {confirmButtonText: '确定'});
                    }
                })
        }
    }
  }
</script>

<style>
    .main {
        width: 1000px;
        margin:0 auto;
    }
</style>
