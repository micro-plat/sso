<template>
  <div class="app-content-body fade-in-up ng-scope">
      <div class="panel panel-default">
        <div class="bg-light lter b-b wrapper-md hidden-print ng-scope">
          <h1 class="m-n font-thin h3">数据角色授权管理</h1>
        </div>
        <div class="panel-body">
          <form class="form-inline">
            <label>系统：</label>
            <select name="roleid" class="form-control not-100" v-model="sysid"  @change="queryRoleDataPermission">
                <option v-for="(s, index) in datalist" :key="index" :value="s.id">{{s.name}}</option>
            </select>
            <a style="margin-left:20px;" class="btn btn-default head-right" @click="back">返回</a>
            <a class="btn btn-success head-right" @click="saveAuth">保存</a>
          </form>
          <div class="line line-dashed b-b line-lg"></div>
              <el-scrollbar style="height:100%">
                  <el-table :data="list" stripe  style="width: 100%">
                    <el-table-column  width="100" prop="enable" label="选择" >
                      <template slot-scope="scope">
                        <el-checkbox v-model="scope.row.checked"></el-checkbox>
                      </template>
                    </el-table-column>
                    <el-table-column width="300" prop="name" label="名称" ></el-table-column>
                    <el-table-column width="300" prop="table_name" label="表名" ></el-table-column>
                    <el-table-column width="300" prop="operate_action" label="操作动作" ></el-table-column>
                    <el-table-column width="300" prop="rules" label="规则信息" ></el-table-column>
                    <el-table-column width="500"   prop="remark" label="备注" ></el-table-column>
                  </el-table>
            </el-scrollbar>

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
      list: null,
      sysid: null,
      tables: null,
      role_id: null
    };
  },
  mounted() {
    this.role_id = this.$route.query.role_id;
    this.querySys();
  },
  methods: {
    back: function() {
      this.$router.push({path: '/user/role'})
    },
    //保存数据权限
    saveAuth: function() {
      var selectAuth = [];
      this.list.forEach(item => {
        if (item.checked) {
          selectAuth.push(item.permissionId)
        }
      });
      this.$http.post("/auth/savepermission", {
        role_id: this.role_id,
        sys_id: this.sysid,
        select_auth: selectAuth.join(",")
      })
        .then(res => {
          this.$notify({
            title: '成功',
            message: '授权成功',
            type: 'success',
            offset: 50,
            duration:2000,
          });
        })
        .catch(err => {
            this.$notify({
              title: '错误',
              message: '网络错误,请稍后再试',
              type: 'error',
              offset: 50,
              duration:2000,
            });
        });
    },
    querySys: function() {
      this.$http.post("/base/getsystems",{})
        .then(res => {
          this.datalist = res;
          if (this.datalist.length > 0) {
            this.sysid = this.datalist[0].id;
            this.queryRoleDataPermission();
          }
        })
        .catch(err => {
            this.$notify({
              title: '错误',
              message: '网络错误,请稍后再试',
              type: 'error',
              offset: 50,
              duration:2000,
            });
        });
    },

    //查询角色与数据的关联信息
    queryRoleDataPermission() {
      this.list = [];
      this.$http.post("/auth/permissionquery", {
        sys_id: this.sysid,
        role_id: this.role_id
      })
      .then(res => {
        if (res.length > 0) {
          res.forEach(item => {
            this.list.push({
              checked:item.checked == "1",
              name:item.name,
              remark: item.remark,
              operate_action: item.operate_action,
              table_name:item.table_name,
              rules: item.rules,
              permissionId:item.id});
          });
        }
      })
    }
  }
};
</script>
<style scoped>
</style>
