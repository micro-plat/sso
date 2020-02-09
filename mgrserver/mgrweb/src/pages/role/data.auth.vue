<template>
  <div class="app-content-body fade-in-up ng-scope">
      <div class="panel panel-default">
        <div class="bg-light lter b-b wrapper-md hidden-print ng-scope">
          <h1 class="m-n font-thin h3">数据角色授权管理</h1>
        </div>
        <div class="panel-body">
          <form class="form-inline">
            <label>选择系统：</label>
            <select name="roleid" class="form-control not-100" v-model="sysid"  @change="querySystemTypeInfo">
                <option v-for="(s, index) in datalist" :key="index" :value="s.id">{{s.name}}</option>
            </select>
            <label style="margin-left:20px;">选择类型：</label>
            <select name="roleid" class="form-control not-100" v-model="data_type" @change="queryRoleDataPermission">
                <option v-for="(s, index) in typelist" :key="index" :value="s.type">{{s.type_name}}</option>
            </select>
            <a style="margin-left:20px;" class="btn btn-default head-right" @click="back">返回</a>
            <a class="btn btn-success head-right" @click="saveAuth">保存</a>
          </form>
          <div class="line line-dashed b-b line-lg"></div>
              <el-scrollbar style="height:100%">
                  <el-table :data="list" stripe  style="width: 100%">
                    <el-table-column  width="100" prop="enable" label="选择" >
                      <template slot-scope="scope">
                        <el-checkbox v-model="scope.row.checked" @change="checked=>choseChange(checked, scope.row.isall)"></el-checkbox>
                      </template>
                    </el-table-column>
                    <el-table-column width="300" prop="name" label="名称" ></el-table-column>
                    <el-table-column width="300" prop="type_name" label="类型名称" ></el-table-column>
                    <el-table-column width="300" prop="value" label="值" ></el-table-column>
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
      typelist:null,
      list: null,
      sysid: null,
      data_type: null,
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
      if (!this.data_type) {
        return
      }

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
            this.querySystemTypeInfo();
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

    //查询某个系统下面的type信息
    querySystemTypeInfo() {
      this.$http.post("/base/getpermisstypes",{sys_id: this.sysid})
        .then(res => {
          this.typelist = res;
          this.data_type = null;
          if (this.typelist.length > 0) {
            this.data_type = this.typelist[0].type;
            this.queryRoleDataPermission();
            return;
          }
          this.list = [];
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
      if (!this.data_type) {
        console.log("data_type empty");
        return;
      }

      this.$http.post("/auth/permissionquery", {
        sys_id: this.sysid,
        role_id: this.role_id,
        data_type: this.data_type
      })
      .then(res => {
        if (res.length > 0) {
          res.forEach(item => {
            this.list.push({
              checked:item.checked == "1",
              isall:item.isall == "1",
              name:item.name,
              remark: item.remark,
              type: item.type,
              type_name:item.type_name,
              value: item.value,
              permissionId:item.id});
          });
        }
      })
    },

    //checked 是否选中，isAll是否表示全部数据
    choseChange(checked, isAll){
      console.log("checked:",checked)
      console.log("是否代表全部:", isAll)

      if (!checked) {
        return;
      }

      this.list.forEach(item => {
        if (isAll) {
          if (!item.isall) {
            item.checked = false;
          }
        } else {
         if (item.isall) {
           item.checked = false;
         } 
        }
      });
    }
  }
};
</script>
<style scoped>
</style>
