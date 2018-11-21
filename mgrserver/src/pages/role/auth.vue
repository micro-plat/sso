//角色的授权

<template>
  <div  class="app-content-body fade-in-up ng-scope">
      <div class="panel panel-default">
        <div class="bg-light lter b-b wrapper-md hidden-print ng-scope">
          <h1 class="m-n font-thin h3">角色授权管理</h1>
        </div>
        <div class="panel-body">
          <form class="form-inline">
            <select name="roleid" class="form-control not-100" v-model="sysid">
                <option value="" selected="selected" >---请选择系统---</option>
                <option v-for="(s, index) in datalist" :key="index" :value="s.id">{{s.name}}</option>
            </select>
            <a class="btn btn-success" @click="queryTree" >切换</a>
            <a class="btn btn-default head-right" @click="back">返回</a>
            <a class="btn btn-success head-right" @click="saveAuth">保存</a>
          </form>
          <div class="line line-dashed b-b line-lg"></div>
          <v-tree ref="tree" :data='ztreeDataSource' :multiple='true' :halfcheck='true'/>
        </div>
        <footer class="panel-footer text-right bg-light lter">
          <a class="btn btn-success" @click="saveAuth">保存</a>
          <a class="btn btn-default" @click="back">返回</a>
        </footer>
      </div>
  </div>
</template>
<script>
export default {
  data() {
    return {
      datalist: null,
      sysid: null,
      currentData: {},
      role_id: null,
      ztreeDataSource: [],
      selectAuth: []
    };
  },
  mounted() {
    this.role_id = this.$route.query.role_id;
    this.querySys();
  },
  methods: {
    back: function() {
      this.$router.go(-1);
    },
    saveAuth: function() {
      var array = this.$refs.tree.getCheckedNodes();
      for (var i = 0; i < array.length; i++) {
        this.selectAuth.push(array[i].id);
      }
      this.$post("/sso/auth", {
        role_id: this.role_id,
        sys_id: this.sysid,
        selectauth: this.selectAuth.join(",")
      })
        .then(res => {
          this.selectAuth = [];
          this.$notify({
            title: '成功',
            message: '授权成功',
            type: 'success',
            offset: 50,
            duration:2000,
          });
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
            this.selectAuth = [];
          }

        });
    },
    queryTree: function() {
      this.$put("sso/auth", {
        sys_id: this.sysid,
        role_id: this.role_id
      })
        .then(res => {
          if (res.length > 0) {
            this.ztreeDataSource = res;
            return;
          }
          this.ztreeDataSource = [
            {
              title: "新节点",
              children: [],
              path: "",
              icon: "",
              isNew: true,
              parentId: 0,
              parentLevel: 0
            }
          ];
        })
        .catch(err => {
          // this.$router.push("/member/login");
        });
    },
    querySys: function() {
      this.$post("/sso/base",{})
        .then(res => {
          this.datalist = res.list;
          if (this.datalist.length > 0) {
            this.sysid = this.datalist[0].id;
            this.queryTree();
          }
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
  }
};
</script>
<style scoped>
/*.not-100{width: auto !important;}*/
/*.head-right{float: right;margin-left: 4px;}*/
</style>
>
