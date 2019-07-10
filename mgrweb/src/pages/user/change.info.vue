<!-- 修改密码 -->
<template>
  <div class="bg-light lter b-b wrapper-md panel panel-default">
    <bootstrap-modal ref="msg2Modal" :need-header="true" :need-footer="true">
      <div slot="title">
        个人资料修改
      </div>
      <div slot="body">
        <div class="panel panel-default">

          <div class="panel-body">
            <form role="form" class="ng-pristine ng-valid ng-submitted">
              <div class="form-group">
                <label>用户名</label>
                <input type="text" class="form-control" v-model="userData.username" readonly="readonly">
              </div>
              <div class="form-group">
                <label>电话</label>
                <input type="text" class="form-control" v-model="userData.tel" v-validate="'required|numeric|digits:11'" name="tel">
                <span v-show="errors.has('tel')" style="color:red">电话号码为11位数字</span>
              </div>
              <div class="form-group">
                <label>邮箱</label>
                <input type="email" class="form-control" v-model="userData.email" name="email" v-validate="'required|email'">
                <span v-show="errors.has('email')" style="color:red">邮箱格式不对</span>
              </div>
            </form>
          </div>
        </div>
      </div>
      <div slot="footer">

        <a class="btn btn-sm btn-danger" @click="submitInfo">确定</a>
        <a class="btn btn-sm btn-primary" @click="$router.back(-1)">取消</a>
      </div>
    </bootstrap-modal>
  </div>

</template>
<script>
  export default {
    components: {
      "bootstrap-modal": require("vue2-bootstrap-modal"),
    },
    data() {
      return {
        expassword: "",
        password1: "",
        password2: "",
        notMatch: false,
        userData: {username:localStorage.getItem("username"),tel:null,email:null},
      };
    },
    mounted(){
      this.getUserInfo()
      this.showModal()
    },
    methods: {
      showModal(){
        this.$refs.msg2Modal.open()
      },
      getUserInfo(){
          this.$fetch("/sso/user/info",{})
            .then(res => {
              this.userData.tel = res.userinfo.mobile
              this.userData.email = res.userinfo.email
          })
            .catch(err => {
              if (err.response.status == 403) {
                this.$router.push("/member/login");
              }else{
                this.$notify({
                  title: '错误',
                  message: '网络错误,请稍后再试',
                  type: 'error',
                  offset: 50,
                  duration:2000,
                });
              }
            });
      },
      submitInfo() {
        this.$validator.validateAll().then(msg => {
          if (!msg) {
            this.$notify({
              title: '提示',
              message: '请检查填写的内容',
              type: 'warning',
              offset: 50,
              duration: 2000
            });
            return false;
          } else {
            this.$fetch("/sso/user/edit",this.userData)
              .then(res => {
                this.getUserInfo()
                this.$notify({
                  title: '成功',
                  message: '个人信息修改成功',
                  type: 'success',
                  offset: 50,
                  duration: 2000
                });
                this.$router.back(-1);
              })
              .catch(err => {
                if (err.response.status == 403) {
                  this.$notify({
                    title: '错误',
                    message: '登录超时,请重新登录',
                    type: 'error',
                    offset: 50,
                    duration:2000,
                    onClose: function () {
                      this.$router.push("/member/login");
                    }
                  });
                }else{
                  this.$notify({
                    title: '错误',
                    message: '网络错误,请稍后再试',
                    type: 'error',
                    offset: 50,
                    duration:2000,
                  });
                }
              });
          }
        });
      }
    }
  };
</script>
<style scoped>
  .line {
    width: 100%;
    height: 20px;
    margin: 10px 0;
    overflow: hidden;
    font-size: 0;
  }
  .line-xs {
    margin: 0;
  }

  .line-lg {
    margin-top: 15px;
    margin-bottom: 15px;
  }

  .line-dashed {
    background-color: transparent;
    border-style: dashed !important;
    border-width: 0;
  }
  .b-b {
    border-bottom: 1px solid #dee5e7;
  }
</style>
