<!--角色基础信息查询-->
<template>
  <div ref="main">
    <div class="panel panel-default">
      <div class="panel panel-default">
        <div class="panel-body">
          <form class="form-inline" role="form">
            <div class="form-group">
              <input
                type="text"
                class="form-control"
                onkeypress="if(event.keyCode == 13) return false;"
                v-model="paging.role_name"
                placeholder="请输入角色名"
              />
            </div>
            <div class="form-group">
              <select v-model="paging.status" class="form-control visible-md visible-lg">
                <option selected="selected" value="-1">---请选择状态---</option>
                <option value="0">启用</option>
                <option value="2">禁用</option>
              </select>
            </div>
            <a class="btn btn-success" @click="searchClick">查询</a>
            <a class="btn btn-primary" @click="showModal(1,{})">添加角色</a>
          </form>
        </div>
      </div>

      <el-scrollbar style="height:100%">
        <el-table :data="datalist.items" stripe style="width: 100%">
          <el-table-column align="center" width="200" prop="role_name" label="角色名"></el-table-column>

          <el-table-column align="center" width="180" prop="status" label="状态">
            <template slot-scope="scope">
              <!-- <el-tag type="success" v-if="scope.row.status == 0">{{scope.row.status_label}}</el-tag>
              <el-tag type="info" v-if="scope.row.status == 2">{{scope.row.status_label}}</el-tag> -->
              <el-tag :type="scope.row.status == '0' ?'success':'info'">{{scope.row.status | fltrEnum("role_status")}}</el-tag>
            </template>
          </el-table-column>

          <el-table-column align="center" prop="create_time" label="创建时间">
            <template slot-scope="scope">
              <i class="el-icon-time"></i>
              <span style="margin-left: 10px">{{ scope.row.create_time }}</span>
            </template>
          </el-table-column>

          <el-table-column align="center" label="操作">
            <template slot-scope="scope">
              <el-button plain type="primary" size="mini" @click="showModal(0,scope.row)">编辑</el-button>
              <el-button
                plain
                type="success"
                size="mini"
                @click="roleChange(0,scope.row.role_id)"
                v-if="scope.row.status == 2"
              >启用</el-button>
              <el-button
                plain
                type="info"
                size="mini"
                @click="roleChange(2,scope.row.role_id)"
                v-if="scope.row.status == 0"
              >禁用</el-button>
              <el-button plain type="danger" size="mini" @click="roleDel(scope.row.role_id)">删除</el-button>
              <el-button
                plain
                type="success"
                size="mini"
                @click="menuAuth(scope.row.role_id, scope.row.role_name)"
              >菜单授权</el-button>
              <el-button
                plain
                type="success"
                size="mini"
                @click="dataAuth(scope.row.role_id, scope.row.role_name)"
              >数据授权</el-button>
            </template>
          </el-table-column>
        </el-table>
      </el-scrollbar>

      <el-dialog width="30%" :visible.sync="dialogFormVisible">
        <div slot="title">{{isAdd == 1 ? "添加角色" : "编辑角色"}}</div>
        <form role="form" class="ng-pristine ng-valid ng-submitted height-min">
          <div class="form-group">
            <label>角色名</label>
            <input
              name="rolename1"
              type="text"
              class="form-control"
              v-model="roleInfo.role_name"
              v-validate="'required'"
              placeholder="请输入角色名"
              required
              maxlength="32"
            />
            <div class="form-heigit">
              <span v-show="errors.first('rolename1')" class="text-danger">角色名不能为空！</span>
            </div>
          </div>
          <div class="form-group" v-if="isAdd == 1">
            <label class="checkbox-inline">
              <input id="statuscheck" type="checkbox" />是否启用
            </label>
          </div>
        </form>
        <div slot="footer" class="dialog-footer">
          <el-button size="small" @click="dialogFormVisible = false">取消</el-button>
          <el-button type="success" size="small" @click="submitRole">提交</el-button>
        </div>
      </el-dialog>

      <div class="page-pagination">
        <el-pagination
          @size-change="handleSizeChange"
          @current-change="pageChange"
          :current-page="paging.pi"
          :page-size="paging.ps"
          :page-sizes="pageSizes"
          layout="total, sizes, prev, pager, jumper"
          :total="totalpage"
        ></el-pagination>
      </div>
    </div>
  </div>
</template>

<script>
import pager from "vue-simple-pager";
import PullTo from "vue-pull-to";
export default {
  components: {
    "bootstrap-modal": require("vue2-bootstrap-modal"),
    pager: pager,
    PullTo
  },
  data() {
    return {
      title: () => {
        if (this.isAdd == 1) {
          return "添加";
        } else {
          return "编辑";
        }
      },
      editOrAddData: {},
      dialogFormVisible: false,
      paging: {
        ps: 10,
        pi: 1,
        role_name: "",
        status: -1
      },
      pageSizes: [5, 10, 20, 50], //可选显示数据条数
      datalist: {
        count: 0,
        items: []
      },
      totalpage: 0,
      isShowLog: false,
      isAdd: 1,
      selectRole: {},
      roleInfo: {
        role_name: "",
        role_id: -1,
        status: 0,
        is_add: 2
      }
    };
  },
  mounted() {
    this.$refs.main.style.height = document.documentElement.clientHeight + "px";
    this.queryData();
  },
  methods: {
    loadmore: function(loaded) {
      return new Promise(function(resolve, reject) {
        setTimeout(function() {
          loaded("done");
        }, 1000);
      });
    },
  
    queryData: function() {
      if (this.paging.pi == 0) {
        this.paging.pi = 1;
      }
      this.$http
        .get("/role/index/getall", this.paging)
        .then(res => {
          this.datalist.items = res.list;
          this.datalist.count = new Number(res.count);
          this.totalpage = res.count;
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
    },
    pageChange: function(data) {
      this.paging.pi = data;
      this.queryData();
    },
    handleSizeChange(val) {
      this.paging.ps = val;
      this.queryData();
    },
    searchClick: function() {
      this.paging.pi = 1;
      this.paging.ps = 10;
      this.queryData();
    },

    showModal: function(i, j) {
      if (i == 1) {
        this.isAdd = 1;
        this.roleInfo.role_name = "";
        this.roleInfo.role_id = -1;
        this.roleInfo.status = 0;
        this.roleInfo.is_add = 1;
      } else {
        this.isAdd = 0;
        this.roleInfo.role_name = j.role_name;
        this.roleInfo.role_id = j.role_id;
        this.roleInfo.status = j.status;
        this.roleInfo.is_add = 2;
      }
      this.dialogFormVisible = true;
      //this.$refs.editModal.open();
    },
    roleChange: function(ests, roleid) {
      var role = {
        status: ests,
        role_id: roleid
      };
      this.$confirm("确认执行该操作吗? 是否继续?", "提示", {
        confirmButtonText: "确定",
        cancelButtonText: "取消",
        type: "warning"
      }).then(() => {
        this.$http
          .post("/role/index/changestatus", role)
          .then(res => {
            this.queryData();
            this.$notify({
              title: "成功",
              message: "状态修改成功",
              type: "success",
              offset: 50,
              duration: 2000
            });
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
      });
    },
    roleDel: function(roleid) {
      this.$confirm("此操作将永久删除该数据, 是否继续?", "提示", {
        confirmButtonText: "确定",
        cancelButtonText: "取消",
        type: "warning"
      }).then(() => {
        this.$http
          .post("/role/index/del", { role_id: roleid })
          .then(res => {
            this.queryData();
            this.$notify({
              title: "成功",
              message: "删除成功",
              type: "success",
              offset: 50,
              duration: 2000
            });
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
      });
    },
    submitRole() {
      this.$validator.validate().then(result => {
        if (!result) {
          return false;
        } else {
          if (this.isAdd == 1) {
            this.roleInfo.is_add = 1;
            this.roleInfo.status = 2;
            if (document.getElementById("statuscheck").checked) {
              this.roleInfo.status = 0;
            }
          }
          this.$http
            .post("/role/index/save", this.roleInfo)
            .then(res => {
              this.dialogFormVisible = false;
              this.queryData();
            })
            .catch(err => {
              if (err.response) {
                switch (err.response.status) {
                  case 910:
                    this.$notify({
                      title: "错误",
                      message: "角色名称已被使用",
                      type: "error",
                      offset: 50,
                      duration: 2000
                    });
                    break;
                  default:
                    this.$notify({
                      title: "错误",
                      message: "网络错误,请稍后再试",
                      type: "error",
                      offset: 50,
                      duration: 2000
                    });
                }
              }
            });
        }
      });
    },
    menuAuth(id, role_name) {
      this.$emit("addTab", "菜单授权(" + role_name + ")", "/pages/role/auth?id=" + id);
    },
    dataAuth(id, role_name) {
      this.$emit(
        "addTab",
        "数据授权(" + role_name + ")",
        "/pages/role/dataauth?id=" + id
      );
     }
  }
};
</script>

<style scoped>
.page-pagination {
  padding: 10px 15px;
  text-align: right;
  margin-bottom: 50px;
}
</style>
