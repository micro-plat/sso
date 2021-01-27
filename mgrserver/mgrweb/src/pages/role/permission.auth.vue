<template>
  <div class="app-content-body fade-in-up ng-scope">
      <div class="panel panel-default">
        <div class="bg-light lter b-b wrapper-md hidden-print ng-scope">
          <h1 class="m-n font-thin h3">数据角色授权管理</h1>
        </div>
        <div class="panel-body">
          <form class="form-inline">
            <label>系统：</label>
            <select name="roleid" class="form-control not-100" v-model="sysid"  @change="query">
                <option v-for="(s, index) in systemlist" :key="index" :value="s.id">{{s.name}}</option>
            </select>
            <a class="btn btn-default head-right" @click="back">返回</a>
            <a class="btn btn-success head-right" @click="saveAuth">保存</a>
          </form>
          <div style="overflow-y:scroll">
              <el-scrollbar style="height:100%">
                  <el-table :data="datalist" stripe  style="width: 100%">
                    <el-table-column  width="100" prop="selected" label="选择" >
                      <template slot-scope="scope">
                        <el-checkbox v-model="scope.row.selected"></el-checkbox>
                      </template>
                    </el-table-column>
                    <el-table-column width="300" prop="name" label="规则名称" ></el-table-column>
                    <el-table-column width="300" prop="table_name" label="表名" ></el-table-column>
                    <el-table-column width="300" prop="operate_action" label="操作动作" ></el-table-column>
                    <!-- <el-table-column width="200" prop="status" label="规则状态">
                      <template slot-scope="scope">
                        <el-tag type="info" v-if="scope.row.status == 1">禁用</el-tag>
                        <el-tag type="success" v-if="scope.row.status == 0">启用</el-tag>
                      </template>
                    </el-table-column> -->
                    <el-table-column width="400" prop="remark" label="规则备注" ></el-table-column>
                    <!-- <el-table-column  label="操作">
                      <template slot-scope="scope">
                        <el-button plain type="primary" size="mini" @click="edit(scope.row.id)">编辑</el-button>
                        <el-button plain type="success" size="mini" @click="enable(scope.row.id,0)" v-if="scope.row.status == 1" >启用</el-button>
                        <el-button plain type="info" size="mini" @click="disable(scope.row.id,1)" v-if="scope.row.status == 0">禁用</el-button>
                        <el-button plain  type="danger" size="mini" @click="del(scope.row.id)">删除</el-button>
                      </template>
                     </el-table-column> -->
                  </el-table>
                </el-scrollbar>
            </div>
            <!-- <div class="page-pagination">
              <el-pagination
                @size-change="handleSizeChange"
                @current-change="goPage"
                :current-page="pi"
                :page-size="ps"
                :page-sizes="pageSizeList"
                layout="total, sizes, prev, pager, jumper"
                :total="datacount">
              </el-pagination>
            </div> -->
        </div>
        <footer class="panel-footer text-right bg-light lter">
          <a class="btn btn-default" @click="back">返回</a>
          <a class="btn btn-success" @click="saveAuth">保存</a>
        </footer>
      </div>

      <!-- <bootstrap-modal ref="editModal" :need-header="true" :need-footer="true" no-close-on-backdrop="true">
        <div slot="title">{{isAdd ? "新增" : "编辑"}}</div>
        <div slot="body">
          <div class="panel panel-default">
            <div class="panel-body">
              <form role="form" class="ng-pristine ng-valid ng-submitted height-min">
                <div class="form-group">
                  <label>名称(必填)</label>
                  <el-input v-model="rolePermission.name" placeholder="请输入名称" maxlength="64" ></el-input>
                </div>
                <div>
                  <label>选择规则(选择多个时,之间的关系是and)</label>
                  <div style="max-height:300px;overflow-y:scroll">
                    <el-table :data="currentPermissions" stripe  style="width: 100%">
                      <el-table-column  width="100" prop="enable" label="选择" >
                        <template slot-scope="scope">
                          <el-checkbox v-model="scope.row.checked"></el-checkbox>
                        </template>
                      </el-table-column>
                      <el-table-column width="100" prop="id" label="标识" ></el-table-column>
                      <el-table-column width="200" prop="name" label="规则名称" ></el-table-column>
                    </el-table>
                  </div>
                </div>
              </form>
            </div>
          </div>
        </div>
        <div slot="footer">
          <el-button size="small" @click="onClose">取消</el-button>
          <el-button type="success" size="small" @click="saveAuth">保存</el-button>
        </div>
      </bootstrap-modal> -->
  
  </div>
</template>
<script>
// import pager from "vue-simple-pager"
import PullTo from 'vue-pull-to'

export default {
  components: {
    // "bootstrap-modal": require("vue2-bootstrap-modal"),
    // pager,
    PullTo
  },
  data() {
    return {
      // pageSizeList: [5, 10, 20, 50], //可选显示数据条数
      // datacount: 0,
      datalist: null,
      systemlist :null,
      sysid: null,
      //tables: null,
      role_id: null,
      // pi: 1,
      // ps:10,
      // totalPage: 0,
      // isAdd:false,
      // rolePermission:{},
      // allPermissionConfig:[],
      // currentPermissions:[],
    };
  },
  mounted() {
    this.role_id = this.$route.query.id;
    this.querySys();
  },
  methods: {
    back: function() {
      this.$router.push({path: '/pages/role/index'})
    }, 
    //保存数据权限
    saveAuth: function() {
      var selectAuth = [];
      this.datalist.forEach(item => {
        if (item.selected) {
          selectAuth.push(item.id)
        }
      });
      this.$http.post("/role/index/permissionsave", {
        //id: this.rolePermission.id,
        role_id: this.role_id,
        sys_id: this.sysid,
        // table_name:this.rolePermission.table_name,
        // operate_action: this.rolePermission.operate_action,
        // name: this.rolePermission.name,
        permissions: selectAuth.join(",")
      })
      .then(res => {
        this.$notify({
          title: '成功',
          message: '授权成功',
          type: 'success',
          offset: 50,
          duration:2000,
        });
        this.query();
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
          this.systemlist = res;
          if (this.systemlist.length > 0) {
            this.sysid = this.systemlist[0].id;
            this.query();
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
    query() {
      this.$http.post("/role/index/permissionquery", {
        sys_id: this.sysid,
        role_id: this.role_id,
        //pi: this.pi,
        //ps: this.ps
      })
      .then(res => {
          // this.datacount = res.count;
          // this.totalPage = Math.ceil(res.count / 10);
          var temp = res.list
          temp.forEach(item => {
            item.selected = false;
            if (item.checked == "1") {
                item.selected = true
            }
            // //var roleConfig = this.generateRuleConfig(item.permissions);
            // item.rules = roleConfig.showStr;
            // //item.editData = roleConfig.editData;
          });
          temp.sort(function(a, b){return b.selected-a.selected });
          console.log(temp);
          this.datalist = temp;
      })
    },
     
  }
};
</script>
<style scoped>
.page-pagination {
    padding: 10px 15px;
    text-align: right;
}
</style>
