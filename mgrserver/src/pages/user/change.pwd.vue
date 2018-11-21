<!-- 修改密码 -->
<template>
  <div class="bg-light lter b-b wrapper-md panel panel-default">
    <bootstrap-modal ref="msg2Modal" :need-header="true" :need-footer="true">
      <div slot="title">
        修改密码
      </div>
      <div slot="body">
        <div class="panel panel-default">
          <div class="panel-body">
            <form role="form" class="ng-pristine ng-valid ng-submitted">
              <div class="form-group">

                  <label>请输入原密码</label>
                  <input type="password" class="form-control" name="expassword" v-model="expassword" v-validate="'required|min:6'" required >
                  <span v-show="errors.has('expassword')" style="color:red">密码不少于6位</span>

              </div>
              <div class="form-group">
                  <label>请输入新密码</label>
                  <input type="password" name="password1" class="form-control" v-model="password1" v-validate="'required|min:6'" required >
                  <span v-show="errors.has('password1')" style="color:red">密码不少于6位</span>
              </div>
              <div class="form-group">
                  <label>请再次输入新密码</label>
                  <input type="password" name="password2" class="form-control" v-model="password2" v-validate="{'required': 'true', 'is': password1}" required >
                  <span v-show="errors.has('password2')" style="color:red">两次密码不匹配</span>

              </div>
            </form>
          </div>
        </div>
      </div>
      <div slot="footer">
        <button class="btn btn-sm btn-danger" @click="submitPswd">确定</button>
        <button class="btn btn-sm btn-primary" @click="$router.back(-1)">取消</button>
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
      notMatch: false
    };
  },
  mounted(){
   this.showModal()
  },
  methods: {
    showModal(){
      this.$refs.msg2Modal.open()
    },
    check: function() {
      if (this.password1 != this.password2) {
        this.notMatch = true;
      } else {
        this.notMatch = false;
      }
    },
    submitPswd() {
      this.$validator.validateAll().then(msg => {
        if (!msg) {

          return false;
        } else {
          this.$fetch("/sso/user/changepwd",{expassword:this.expassword,newpassword:this.password1})
            .then(res => {
              this.$notify({
                title: '成功',
                message: '密码修改成功',
                type: 'success',
                offset: 50,
                duration: 2000
              });
              this.$router.back(-1)
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

