<!-- 用户信息列表，查询，启动，禁用，解锁 -->
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
                v-model="paging.username"
                onkeypress="if(event.keyCode == 13) return false;"
                placeholder="请输入用户名"
              />
            </div>
            <div class="form-group">
              <select
                v-model="paging.role_id"
                name="role_id"
                class="form-control visible-md visible-lg"
              >
                <option value selected="selected">---请选择角色---</option>
                <option v-for="(r, index) in roleList" :key="index" :value="r.role_id">{{r.name}}</option>
              </select>
            </div>
            <a class="btn btn-success" @click="searchClick">查询</a>
            <a class="btn btn-primary" @click="showModal(1,{})">添加用户</a>
          </form>
        </div>
      </div>
      <bootstrap-modal ref="editModal" :need-header="true" :need-footer="true" :closed="resetSys">
        <div slot="title">{{isAdd == 1 ? "添加用户" : "编辑用户信息"}}</div>
        <div slot="body">
          <div class="panel panel-default">
            <div class="panel-body">
              <form role="form" class="ng-pristine ng-valid ng-submitted height-min">
                <div class="form-group">
                  <label>用户名</label>
                  <input
                    name="username1"
                    type="text"
                    class="form-control"
                    v-validate="'required'"
                    v-model="userInfo.user_name"
                    placeholder="请输入用户名"
                    required
                    maxlength="32"
                  />
                  <div class="form-heigit">
                    <span v-show="errors.first('username1')" class="text-danger">用户名不能为空！</span>
                  </div>
                </div>
                <div class="form-group">
                  <label>电话</label>
                  <input
                    name="mobile1"
                    type="text"
                    class="form-control"
                    v-validate="'required|numeric'"
                    v-model="userInfo.mobile"
                    placeholder="请输入电话"
                    maxlength="11"
                    required
                  />
                  <div class="form-heigit">
                    <span v-show="errors.first('mobile1')" class="text-danger">请输入正确的11位手机号！</span>
                  </div>
                </div>
                <div class="form-group">
                  <label>邮箱</label>
                  <input
                    name="email1"
                    type="text"
                    class="form-control"
                    v-validate="'required|email'"
                    v-model="userInfo.email"
                    placeholder="请输入邮箱"
                    required
                  />
                  <div class="form-heigit">
                    <span v-show="errors.first('email1')" class="text-danger">请输入正确的邮箱！</span>
                  </div>
                </div>
                <div class="form-group">
                  <label>扩展参数(非必须)</label>
                  <textarea
                    name="ext_params"
                    style="resize:none"
                    rows="5"
                    type="text"
                    class="form-control"
                    v-model="userInfo.ext_params"
                    placeholder="扩展参数"
                  ></textarea>
                  <!-- <div class="form-heigit">
                                    <span v-show="errors.first('ext')" class="text-danger">请输入正确的邮箱！</span>
                  </div>-->
                </div>
                <div class="form-group sel-col-5">
                  <div class="form-inline">
                    <label>系统与角色</label>
                    <span class="add-btn">
                      <a class="btn m-b-xs btn-xs btn-success" @click="add">
                        <i class="fa fa-plus"></i>
                      </a>
                    </span>
                  </div>
                  <form
                    class="form-inline pull-in clearfix"
                    v-for="(list,index) in userInfo.lists"
                    v-bind:key="list.id"
                  >
                    <div class="form-group col-sm-7">
                      <select
                        name="select1"
                        class="form-control"
                        v-validate="'required'"
                        v-model="list.sys_id"
                        @change="sysStatus(list.sys_id,index)"
                        required
                      >
                        <option value selected="selected">---请选择系统---</option>
                        <option
                          v-for="(r, index) in sysList"
                          :key="index"
                          :value="r.id"
                          :disabled="r.disabled"
                        >{{r.name}}</option>
                      </select>
                    </div>
                    <div class="form-group col-sm-4">
                      <select
                        name="select2"
                        class="form-control"
                        v-validate="'required'"
                        v-model="list.role_id"
                        required
                      >
                        <option value selected="selected">---请选择角色---</option>
                        <option
                          v-for="(r, index) in roleList"
                          :key="index"
                          :value="r.role_id"
                        >{{r.name}}</option>
                      </select>
                    </div>
                    <div
                      class="form-group del-btn"
                      v-if="userInfo.lists.length > 1"
                      style="margin:4px"
                    >
                      <a class="btn m-b-xs btn-xs btn-danger" @click="del(index)">
                        <i class="fa fa-minus"></i>
                      </a>
                    </div>
                  </form>
                  <div class>
                    <div class="form-group form-heigit col-sm-5">
                      <span v-show="errors.has('select1')" class="text-danger">必须选择系统！</span>
                    </div>
                    <div class="form-group form-heigit col-sm-5">
                      <span v-show="errors.has('select2')" class="text-danger">必须选择用户角色！</span>
                    </div>
                  </div>
                </div>
                <div class="form-group" v-if="isAdd == 1">
                  <label class="checkbox-inline">
                    <input id="statuscheck" type="checkbox" />是否启用
                  </label>
                </div>
              </form>
            </div>
          </div>
        </div>
        <div slot="footer">
          <el-button size="small" @click="onClose">取消</el-button>
          <el-button type="success" size="small" @click="submitUser">提交</el-button>
        </div>
      </bootstrap-modal>

      <el-scrollbar style="height:100%">
        <el-table :data="datalist.items" stripe style="width: 100%">
          <el-table-column align="center" width="100" prop="user_name" label="用户名"></el-table-column>

          <el-table-column align="center" width="350" prop="rolestr" label="系统/角色"></el-table-column>

          <el-table-column align="center" width="130" prop="mobile" label="联系电话"></el-table-column>
          <el-table-column align="center" width="170" prop="email" label="邮箱"></el-table-column>

          <el-table-column align="center" width="80" prop="status" label="状态">
            <template slot-scope="scope">
              <el-tag type="success" v-if="scope.row.status == 0">{{scope.row.status_label}}</el-tag>
              <el-tag type="info" v-if="scope.row.status == 2">{{scope.row.status_label}}</el-tag>
            </template>
          </el-table-column>

          <el-table-column align="center" prop="create_time" label="创建时间">
            <template slot-scope="scope">
              <i class="el-icon-time"></i>
              <span style="margin-left: 10px">{{ scope.row.create_time }}</span>
            </template>
          </el-table-column>
          <el-table-column align="center" prop="create_time" label="扩展参数">
            <template slot-scope="scope">
              <span style="margin-left: 10px">{{ scope.row.ext_params | RemarkFilter}}</span>
            </template>
          </el-table-column>

          <el-table-column align="center" label="操作">
            <template slot-scope="scope">
              <el-button plain type="primary" size="mini" @click="showModal(2,scope.row)">编辑</el-button>
              <el-button
                plain
                type="success"
                size="mini"
                @click="userChange(0,scope.row.user_id)"
                v-if="scope.row.status == 2"
              >启用</el-button>
              <el-button
                plain
                type="info"
                size="mini"
                @click="userChange(2,scope.row.user_id)"
                v-if="scope.row.status == 0"
              >禁用</el-button>

              <el-button plain type="danger" size="mini" @click="userDel(scope.row.user_id)">删除</el-button>
            </template>
          </el-table-column>
        </el-table>
      </el-scrollbar>

      <div class="page-pagination">
        <el-pagination
          @size-change="handleSizeChange"
          @current-change="pageChange"
          :current-page="paging.pi"
          :page-size="paging.ps"
          :page-sizes="pageSizes"
          layout="total, sizes, prev, pager, next, jumper"
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
      paging: {
        ps: 5,
        pi: 1,
        username: "",
        role_id: ""
      },
      pageSizes: [5, 10, 20, 50], //可选显示数据条数

      datalist: {
        count: 0,
        items: []
      },
      userInfo: {
        user_name: "",
        user_id: -1,
        lists: [],
        mobile: null,
        email: null,
        status: 0,
        is_add: 2,
        ext_params: ""
      },

      totalpage: 0,
      sysList: [],
      roleList: [], //角色列表
      selectSys: [],
      isAdd: 1
    };
  },
  created() {},
  mounted() {
    this.querySys();
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
      this.$http.post("/user/getall", this.paging)
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
    pageChange(val) {
      this.paging.pi = val;
      this.queryData();
    },
    handleSizeChange(val) {
      this.paging.ps = val;
      this.queryData();
    },

    searchClick: function() {
      this.paging.pi = 1;
      this.queryData();
    },
    next() {
      let pi = this.paging.pi;
      this.paging.pi = pi + 1;
      this.$http.post("/user/getall", this.paging)
        .then(res => {
          if (res.list.length <= 0) {
            this.paging.pi = pi;
            return false;
          }
          this.datalist.items = this.datalist.items.concat(res.list);
          this.datalist.count = new Number(res.count);
          this.totalpage = Math.ceil(this.datalist.count / this.paging.ps);
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
    stateChange(e) {
      console.log(e);
    },
    showModal: function(i, j) {
      if (j.status == 1) {
        this.$notify({
          title: "警告",
          message: "请先进行解锁操作",
          type: "warning",
          offset: 50,
          duration: 2000
        });
        return false;
      }

      if (i == 1) {
        // 添加用户
        this.isAdd = 1;
        this.userInfo.user_name = "";
        this.userInfo.role_id = "";
        this.userInfo.mobile = null;
        this.userInfo.status = 0;
        this.userInfo.user_id = -1;
        this.userInfo.is_add = 1;
        this.userInfo.lists = [
          {
            sys_id: "",
            role_id: ""
          }
        ];
        this.userInfo.email = null;
        this.selectSys.push("");
      } else {
        // 编辑用户
        this.isAdd = 2;
        this.userInfo.user_name = j.user_name;
        this.userInfo.role_id = j.role_id;
        this.userInfo.mobile = j.mobile;
        this.userInfo.status = j.status;
        this.userInfo.user_id = j.user_id;
        this.userInfo.lists = j.roles;
        this.userInfo.is_add = 2;
        this.userInfo.email = j.email;
        this.userInfo.ext_params = j.ext_params;
        for (var s = 0; s < j.roles.length; s++) {
          this.selectSys.push(j.roles[s].sys_id);
        }
        this.setSys();
      }
      this.$refs.editModal.open();
    },
    userChange: function(ests, userid) {
      var r;
      this.$confirm("确认执行该操作吗?", "提示", {
        confirmButtonText: "确定",
        cancelButtonText: "取消",
        type: "warning"
      })
        .then(() => {
          var user = {
            status: ests,
            user_id: userid
          };
          this.$http.post("/user/changestatus", user)
            .then(res => {
              this.queryData();
              this.$notify({
                title: "成功",
                message: "修改状态成功",
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
        })
    },
    userDel: function(userid) {
      var user = {
        user_id: userid
      };

      this.$confirm("此操作将永久删除该数据, 是否继续?", "提示", {
        confirmButtonText: "确定",
        cancelButtonText: "取消",
        type: "warning"
      })
        .then(() => {
          this.$http.post("/user/del", user)
            .then(res => {
              this.queryData();
              this.$notify({
                title: "成功",
                message: "成功删除用户",
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
        })
    },
    add() {
      this.userInfo.lists.push({
        sys_id: "",
        role_id: ""
      });
      this.selectSys.push("");
    },
    del(index) {
      for (var s = 0; s < this.sysList.length; s++) {
        if (this.sysList[s].id == this.userInfo.lists[index].sys_id) {
          this.sysList[s].disabled = false;
        }
      }
      this.userInfo.lists.splice(index, 1);
      this.selectSys.splice(index, 1);
    },
    onClose() {
      this.$refs.editModal.close();
    },
    resetSys() {
      this.selectSys = [];
      for (var s = 0; s < this.sysList.length; s++) {
        this.sysList[s].disabled = false;
      }
    },
    setSys() {
      for (var s = 0; s < this.sysList.length; s++) {
        for (var i = 0; i < this.selectSys.length; i++) {
          if (this.sysList[s].id == this.selectSys[i]) {
            this.sysList[s].disabled = true;
          }
        }
      }
    },
    sysStatus(id, idx) {
      if (this.selectSys[idx] != "") {
        for (var s = 0; s < this.sysList.length; s++) {
          if (this.sysList[s].id == this.selectSys[idx]) {
            this.sysList[s].disabled = false;
          }
        }
      }
      this.selectSys[idx] = id;
      for (var s = 0; s < this.sysList.length; s++) {
        if (this.sysList[s].id == id) {
          this.sysList[s].disabled = true;
        }
      }
    },
    submitUser() {
      if (this.userInfo.lists.length < 1) {
        this.$notify({
          title: "警告",
          message: "至少需要添加一个系统角色！",
          type: "warning",
          offset: 50
        });
        return false;
      }
      var s = "";
      for (var i = 0; i < this.userInfo.lists.length; i++) {
        s =
          s +
          this.userInfo.lists[i].sys_id +
          "," +
          this.userInfo.lists[i].role_id +
          "|";
      }
      this.userInfo.auth = s;
      this.$validator.validate().then(result => {
        if (!result) {
          return false;
        } else {
          if (this.isAdd == 1) {
            // 添加用户
            this.userInfo.is_add = 1;
            var x = document.getElementById("statuscheck").checked;
            if (x) {
              this.userInfo.status = 0;
            } else {
              this.userInfo.status = 2;
            }
            this.$http.post("/user/add", this.userInfo)
              .then(res => {
                this.queryData();
                this.$notify({
                  title: "成功",
                  message: "添加成功",
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
          } else if (this.isAdd == 2) {
            this.$http.post("/user/edit", this.userInfo)
              .then(res => {
                this.queryData();
                this.$notify({
                  title: "成功",
                  message: "编辑成功",
                  type: "success",
                  offset: 50,
                  duration: 2000
                });
              })
              .catch(err => {
                if (err.response.status == 403) {
                  this.$notify({
                    title: "错误",
                    message: "登录超时,请重新登录",
                    type: "error",
                    offset: 50,
                    duration: 2000,
                    onClose: () => {
                      this.$router.push("/login");
                    }
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
              });
          }

          this.onClose();
        }
      });
    },
    querySys() {
      this.$http.get("/base/getroles", {})
        .then(res => {
          this.roleList = res;
        })
        .catch(err => {});

      this.$http.post("/base/getsystems", {})
        .then(res => {
          this.sysList = res;
          for (s in this.sysList) {
            s.disabled = false;
          }
        })
        .catch(err => {});
    }
  },
  filters: {
    RemarkFilter(value) {
      if (value != "") {
        return value;
      } else {
        return "-";
      }
    }
  }
};
</script>

<style scoped>
.page-pagination {
  padding: 10px 15px;
  text-align: right;
}
</style>
