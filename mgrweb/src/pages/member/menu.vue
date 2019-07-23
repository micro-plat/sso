<template>
  <div id="app">
    <nav-menu
      :menus="menus"
      :copyright="copyright"
      :themes="themes"
      :logo="logo"
      :systemName="systemName"
      :headpic="headpic"
      :userinfo="userinfo"
      :pwd="pwd"
      :signOut="signOut"
      ref="NewTap"
    >
      <keep-alive>
        <router-view v-if="$route.meta.keepAlive" @addTab="addTab" @close="close" @setTab="setTab" ></router-view>
      </keep-alive>
      <router-view v-if="!$route.meta.keepAlive" @addTab="addTab" @close="close" @setTab="setTab"></router-view>
      <!--<router-view  @addTab="addTab" @close="close"></router-view>-->
    </nav-menu>
    <!-- Add Form -->
    <el-dialog title="修改密码" width="30%" :visible.sync="dialogAddVisible">
      <el-form :model="updateInfo" :rules="rules" ref="addForm">

        <el-form-item label="请输入原密码" prop="password_old">
          <el-input type="password" v-model="updateInfo.password_old"  ></el-input>
        </el-form-item>

        <el-form-item label="请输入新密码" prop="password">
          <el-input type="password" v-model="updateInfo.password"  ></el-input>
        </el-form-item>

        <el-form-item label="请确认密码" prop="checkPass">
          <el-input type="password" v-model="updateInfo.checkPass"  ></el-input>
        </el-form-item>

      </el-form>
      <div slot="footer" >
        <button class="btn btn-sm btn-primary" @click="resetForm('addForm')">取 消</button>
        <button class="btn btn-sm btn-danger"  @click="add('addForm')">确 定</button>
      </div>

    </el-dialog>
    <!--Add Form -->
  </div>
</template>

<script>
  import navMenu from 'nav-menu'; // 引入
  import VueCookies from 'vue-cookies';
  export default {
    name: 'app',
    data () {

      var validatePass = (rule, value, callback) => {
        if (value === '') {
          callback(new Error('请输入密码'));
        } else {
          if (this.updateInfo.checkPass !== '') {
            this.$refs.addForm.validateField('checkPass');
          }
          callback();
        }
      };
      var validatePass2 = (rule, value, callback) => {
        if (value === '') {
          callback(new Error('请再次输入密码'));
        } else if (value !== this.updateInfo.password) {
          callback(new Error('两次输入密码不一致!'));
        } else {
          callback();
        }
      };
      return {
        headpic: "http://sso2.100bm.cn:6888/static/img/a0.jpg",
        logo: "http://sso2.100bm.cn:6888/static/img/d663155de6dc8e060415bbcd891cb9d4.png",
        copyright: "2018 admin-web", //版权信息
        themes: "bg-danger|bg-danger|bg-dark light-danger", //顶部左侧背景颜色,顶部右侧背景颜色,右边菜单背景颜色
        menus: [{}],  //菜单数据
        systemName: "用户权限系统",  //系统名称
        userinfo: {name:'wule',role:"管理员"},
        indexUrl: "/user/index",
        dialogAddVisible:false,     //添加表单显示隐藏
        updateInfo:{
          password_old: "",
          password: "",
          checkPass: "",
        },
        rules: {                    //数据验证规则
          password_old: [
            { required: true, message: "请输入原密码", trigger: "blur" }
          ],
          password: [
            { required: true, message: "请输入新密码", trigger: "blur" },
            { validator: validatePass, trigger: 'change' }
          ],
          checkPass: [
            { required: true, message: "请确认密码", trigger: "blur" },
            { validator: validatePass2, trigger: 'change' }
          ],
        },

      }
    },
    components:{ //注册插件
      navMenu
    },
    created(){
      //this.getAllDictionaryData();
      this.getMenu();
    },
    mounted(){
      // this.$get("/member/getsysinfo",{})
      //   .then(res=>{
      //     console.log("系统信息",res);
      //     this.systemName =res.name;
      //     this.logo = res.logo;
      //     this.themes = res.theme;
      //     if (res.index_url){
      //       this.indexUrl  = res.index_url
      //     }
      //     this.$refs.NewTap.add("首页", this.indexUrl ,{});   //设置默认页面
      //   }).catch(err=>{
      //   console.log(err)
      // });
      this.$refs.NewTap.add("首页", this.indexUrl ,{});

      this.userinfo = JSON.parse(sessionStorage.getItem("userinfo"));
      document.title = "用户权限系统";
    },
    methods:{
      pwd(val){
        VueCookies.remove("__jwt__");
        var config = process.env.service;
        window.location.href = 
            config.ssoWebHost + config.changePwd;
      },
      signOut() {
        VueCookies.remove("__jwt__");
        var config = process.env.service;
        window.location.href = 
            config.ssoWebHost + config.loginUrl + "?callback=" + encodeURIComponent(config.callbackUrl);
      },
      resetForm(formName) {
        this.dialogAddVisible = false;
        this.$refs[formName].resetFields();
      },
      add(formName){
        // console.log(this.addData)
        this.$refs[formName].validate((valid) => {
          if (valid) {
            this.$fetch("/sso/user/changepwd",{
              expassword : this.updateInfo.password_old,
              newpassword : this.updateInfo.password,
            }).then(res=>{
              this.$notify({
                title:'成功',
                message:'修改操作完成',
                type:'success'
              });
              this.dialogAddVisible = false;
              this.$refs[formName].resetFields();
            }).catch(errro=>{
              this.$notify({
                title:'失败',
                message:"原密码错误或密码修改次数超过限制",
                type:'error'
              });
              this.$refs[formName].resetFields();
            })
          } else {
            console.log('error submit!!');
            return false;
          }
        });
      },

      getMenu(){
        this.$fetch("/sso/menu")
          .then(res => {
            this.menus = res;
          })
          .catch(err => {
            console.log(err)
          });
      },
      //@name 标签名称
      //@path 路由
      //@obj  路由参数 类型：Object
      addTab(name,path,obj){
        this.$refs.NewTap.add(name,path,obj);   //调用组件方法，添加一个页面
      },
      close(v){
        this.$refs.NewTap.closeTab(v);
      },
      setTab(name,path,obj){
        console.log("outer",name,path,obj);
        this.$refs.NewTap.set(name,path,obj);
      },
      // 获取枚举字典
      getAllDictionaryData(){
        this.$post('/base/dictionary/info/getall', {})
          .then(response=>{
            // console.log(response);
            this.EnumUtility.Set(response.list);

          })
          .catch(error=>{
            console.error("获取数据字典失败,error:",error)

          });
      },

    }
  }
</script>

<style scoped>

</style>
