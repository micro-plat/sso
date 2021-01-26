//角色的授权

<template>
  <div class="app-content-body fade-in-up ng-scope">
    <div class="panel panel-default">
      <div class="bg-light lter b-b wrapper-md hidden-print ng-scope">
        <h1 class="m-n font-thin h3">菜单角色授权管理</h1>
      </div>
      <div class="panel-body">
        <form class="form-inline">
          <select name="roleid" class="form-control not-100" v-model="sysid">
            <option value selected="selected">---请选择系统---</option>
            <option v-for="(s, index) in datalist" :key="index" :value="s.id">{{s.name}}</option>
          </select>
          <a class="btn btn-success" @click="queryTree">切换</a>
          <a class="btn btn-default head-right" @click="back">返回</a>
          <a class="btn btn-success head-right" @click="saveAuth">保存</a>
        </form>
        <div class="line line-dashed b-b line-lg"></div>
        <v-tree ref="tree" :data="ztreeDataSource" :multiple="true" :halfcheck="true" />
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
    this.role_id = this.$route.query.id;
    this.querySys();
  },
  methods: {
    back: function() {
      this.$router.push({ path: "/pages/user/role" });
    },
    saveAuth: function() {
      this.selectAuth = [];
      var array = this.$refs.tree.getCheckedNodes();
      for (var i = 0; i < array.length; i++) {
        this.selectAuth.push(array[i].id);
      }
      var tempStr = this.selectAuth.join(",");
      if (tempStr == "" || tempStr == undefined) {
        this.$notify({
          title: "错误",
          message: "请选择菜单",
          type: "error",
          offset: 50,
          duration: 2000
        });
        return false;
      }

      if (this.sysid == "") {
        this.$notify({
          title: "错误",
          message: "请选择系统",
          type: "error",
          offset: 50,
          duration: 2000
        });
        return false;
      }

      this.$http
        .post("/role/index/authsave", {
          role_id: this.role_id,
          sys_id: this.sysid,
          selectauth: tempStr
        })
        .then(res => {
          this.selectAuth = [];
          this.$notify({
            title: "成功",
            message: "授权成功",
            type: "success",
            offset: 50,
            duration: 2000
          });
        })
        .catch(err => {
          console.log("err", err.response);
          if (err.response.status == 655) {
            this.$message({
              type: "error",
              message: "角色权限修改未成功，请重试"
            });
          } else {
            this.$notify({
              title: "错误",
              message: "网络错误,请稍后再试",
              type: "error",
              offset: 50,
              duration: 2000
            });
          }
          this.selectAuth = [];
        });
    },
    queryTree: function() {
      if (this.sysid == "") {
        this.$notify({
          title: "错误",
          message: "请选择系统",
          type: "error",
          offset: 50,
          duration: 2000
        });
        return false;
      }
      this.$http
        .post("/role/index/authquery", {
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
        });
    },
    querySys: function() {
      this.$http
        .post("/base/getsystems", {})
        .then(res => {
          this.datalist = res;
          if (this.datalist.length > 0) {
            this.sysid = this.datalist[0].id;
            this.queryTree();
          }
        })
        .catch(err => {
          this.$notify({
            title: "错误",
            message: "网络错误,请稍后再试",
            type: "error",
            offset: 50,
            duration: 2000
          });
        });
    }
  }
};
</script>
<style scoped>
</style>
