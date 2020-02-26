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
                v-model="paging.mobile"
                onkeypress="if(event.keyCode == 13) return false;"
                placeholder="请输入电话号码"
              />
            </div>
            <div class="form-group">
              <select v-model="paging.status" class="form-control visible-md visible-lg">
                <option selected="selected" value="-1">---请选择状态---</option>
                <option value="0">启用</option>
                <option value="1">锁定</option>
                <option value="2">禁用</option>
              </select>
            </div>
            <a class="btn btn-success" @click="searchClick">查询</a>
            <a class="btn btn-primary" @click="showModal(1,{})">添加用户</a>
          </form>
        </div>
      </div>
      <bootstrap-modal ref="editModal" :need-header="true" :need-footer="true"  no-close-on-backdrop="true">
        <div slot="title">{{isAdd == 1 ? "添加用户" : "编辑用户信息"}}</div>
        <div slot="body">
          <div class="panel panel-default">
            <div class="panel-body">
              <form role="form" class="ng-pristine ng-valid ng-submitted height-min">
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
                    required  />
                  <div class="form-heigit">
                    <span v-show="errors.first('mobile1')" class="text-danger">请输入正确的11位手机号！</span>
                  </div>
                </div>

                <div class="form-group">
                  <label>姓名</label>
                  <input
                    name="fullname"
                    type="text"
                    class="form-control"
                    v-validate="'required'"
                    v-model="userInfo.full_name"
                    placeholder="请输入姓名"
                    required
                    maxlength="10"
                  />
                  <div class="form-heigit">
                    <span v-show="errors.first('fullname')" class="text-danger">姓名不能为空！</span>
                  </div>
                </div>
                <div class="form-group">
                  <div><label>邮箱</label></div>
                  <input
                    name="email"
                    type="text"
                    class="form-control"
                    v-validate="'required'"
                    v-model="userInfo.email"
                    placeholder="请输入邮箱"
                    required
                  />
                </div>
                <div class="form-group">
                  <div><label>选择角色</label></div>
                  <select  v-model="userInfo.role_id" class="form-control visible-md visible-lg">
                    <option value="">---请选择角色---</option>
                    <option v-for="item in roleList" :key="item.role_id" :value="item.role_id" >{{item.name}}</option>
                  </select>
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
          <el-table-column align="center" width="130" prop="mobile" label="电话号码"></el-table-column>
          <el-table-column align="center" width="100" prop="full_name" label="姓名"></el-table-column>
          <el-table-column align="center" width="350" prop="rolestr" label="角色"></el-table-column>
          <el-table-column align="center" width="170" prop="email" label="邮箱"></el-table-column>
          <el-table-column align="center" width="80" prop="status" label="状态">
            <template slot-scope="scope">
              <el-tag :type="scope.row.status == '0' ? 'success' : 'info'" >{{scope.row.status_label}}</el-tag>
            </template>
          </el-table-column>
          <el-table-column align="center" width="180" prop="create_time" label="创建时间">
            <template slot-scope="scope">
              <i class="el-icon-time"></i>
              <span style="margin-left: 10px">{{ scope.row.create_time }}</span>
            </template>
          </el-table-column>
          <el-table-column align="center" label="操作">
            <template slot-scope="scope">
              <el-button plain type="primary" size="mini" @click="showModal(2,scope.row)">编辑</el-button>
              <el-button plain type="success" size="mini" @click="userChange(0,scope.row.user_id, scope.row.mobile)" v-if="scope.row.status == 2">启用</el-button>
              <el-button plain type="info" size="mini" @click="userChange(2,scope.row.user_id,scope.row.mobile)" v-if="scope.row.status == 0">禁用</el-button>
              <el-button plain type="success" size="mini" @click="userChange(0,scope.row.user_id,scope.row.mobile)" v-if="scope.row.status == 1">解锁</el-button>
              <el-button plain type="danger" size="mini" @click="userDel(scope.row.user_id)">删除</el-button>
              <el-button plain type="danger" size="mini" @click="setDefaultPwd(scope.row.user_id)">重置密码</el-button>
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
//import "@/services/qrcode.min.js"

export default {
  components: {
    "bootstrap-modal": require("vue2-bootstrap-modal"),
    pager: pager,
    PullTo
  },
  data() {
    return {
      errorTemplate:{
        902: "用户被锁定",
        903: "用户被禁用",
        905: "用户不存在",
        //909: "此登录名已被使用",
        918: "此姓名名已被使用",
        921: "此电话号码已被使用"
      },
      paging: {
        ps: 10,
        pi: 1,
        mobile: "",
        role_id: "",
        status: -1
      },
      pageSizes: [5, 10, 20, 50], //可选显示数据条数

      datalist: {
        count: 0,
        items: []
      },
      userInfo: {
        full_name:"",
        user_name: "",
        role_id :"",
        user_id: -1,
        lists: [],
        mobile: null,
        email: null,
        status: 0,
        is_add: 2,
        ext_params: ""
      },

      totalpage: 0,
      roleList: [], //角色列表
      isAdd: 1
    };
  },
  created() {},
  mounted() {
    this.queryRoles();
    this.queryData();
  },
  methods: {
    queryRoles() {
      this.$http.get("/base/getroles", {})
        .then(res => {
          this.roleList = res;
        })
        .catch(err => {
          console.log(err)
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
        this.userInfo.full_name = "";
        this.userInfo.role_id = "";
        this.userInfo.mobile = null;
        this.userInfo.status = 0;
        this.userInfo.user_id = -1;
        this.userInfo.is_add = 1;
        this.userInfo.email = null;
      } else {
        // 编辑用户
        this.isAdd = 2;
        this.userInfo.full_name = j.full_name;
        this.userInfo.role_id = j.role_id;
        this.userInfo.mobile = j.mobile;
        this.userInfo.status = j.status;
        this.userInfo.user_id = j.user_id;
        this.userInfo.is_add = 2;
        this.userInfo.email = j.email;
      }
      console.log(this.userInfo);
      this.$refs.editModal.open();
    },
    userChange(status, userid,mobile) {
      var r;
      this.$confirm("确认执行该操作吗?", "提示", {
        confirmButtonText: "确定",
        cancelButtonText: "取消",
        type: "warning"
      })
      .then(() => {
        this.$http.post("/user/changestatus", {user_id: userid, status: status, mobile:mobile})
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

    userDel(userid) {
      this.$confirm("此操作将永久删除该数据, 是否继续?", "提示", {
        confirmButtonText: "确定",
        cancelButtonText: "取消",
        type: "warning"
      })
      .then(() => {
        this.$http.post("/user/del", {user_id: userid})
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
    //重置密码
    setDefaultPwd(userid) {
      this.$confirm("是否要重置用户密码?", "提示", {
        confirmButtonText: "确定",
        cancelButtonText: "取消",
        type: "warning"
      })
      .then(() => {
        this.$http.post("/user/setpwd", {user_id: userid})
          .then(res => {
            this.$notify({
              title: "成功",
              message: "重置成功",
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
    onClose() {
      this.$refs.editModal.close();
    },
    submitUser() { 
      if (!(/^1[3456789]\d{9}$/.test(this.userInfo.mobile))) {
        this.$notify({
          title: "警告",
          message: "请输入正确的电话号码",
          type: "warning",
          offset: 50
        });
        return false;
      }

      if (!/^[\u4E00-\u9FA5]+[1-9]?$/.test(this.userInfo.full_name)) {
        this.$notify({
          title: "警告",
          message: "姓名只能为中文或者中文加一个数字",
          type: "warning",
          offset: 50
        });
        return false;
      }

      if (!this.userInfo.role_id){
        this.$notify({
          title: "警告",
          message: "必须选择一个角色",
          type: "warning",
          offset: 50
        });
        return;
      }

      this.$validator.validate().then(result => {
        if (!result) {
          return false;
        } else {
          //note 检查电话号码
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
                this.onClose();
              })
              .catch(err => {
                if (err.response) {
                  var msg = this.errorTemplate[err.response.status] || "出现错误,请稍后再试"
                  this.$notify({
                        title: "错误",
                        message: msg,
                        type: "error",
                        offset: 50,
                        duration: 2000
                      });
                }
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
                this.onClose();
              })
              .catch(err => {
                if (err.response) {
                   var msg = this.errorTemplate[err.response.status] || "出现错误,请稍后再试"
                   this.$notify({
                        title: "错误",
                        message: msg,
                        type: "error",
                        offset: 50,
                        duration: 2000
                    });
                }
              });
          }
        }
      });
    },
  }

};
</script>

<style>
.page-pagination {
  padding: 10px 15px;
  text-align: right;
}
 </style>
