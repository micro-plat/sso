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
                placeholder="请输入登录名或者姓名"
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
      <bootstrap-modal id ="qrCodeModal" ref="qrCodeModal" :need-header="true" size="small">
        <div slot="title">绑定微信账号</div>
        <div slot="body">
          <div>
          <div class="panel-body">
            <div id="qrcodeTable"></div>
          </div>
          </div>
        </div>
      </bootstrap-modal>
      <bootstrap-modal ref="editModal" :need-header="true" :need-footer="true" :closed="resetSys" no-close-on-backdrop="true">
        <div slot="title">{{isAdd == 1 ? "添加用户" : "编辑用户信息"}}</div>
        <div slot="body">
          <div class="panel panel-default">
            <div class="panel-body">
              <form role="form" class="ng-pristine ng-valid ng-submitted height-min">
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
                    maxlength="5"
                  />
                  <div class="form-heigit">
                    <span v-show="errors.first('fullname')" class="text-danger">姓名不能为空！</span>
                  </div>
                </div>
                <div class="form-group">
                  <label>登录名</label>
                  <input
                    name="username1"
                    type="text"
                    class="form-control"
                    v-validate="'required'"
                    v-model="userInfo.user_name"
                    placeholder="请输入登录名"
                    required
                    maxlength="32"
                    :disabled="isAdd==1"
                  />
                  <div class="form-heigit">
                    <span v-show="errors.first('username1')" class="text-danger">登录名不能为空！</span>
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
                  <div><label>邮箱</label></div>
                  <input
                    name="email_pre"
                    type="text"
                    class="email-input"
                    v-validate="'required'"
                    v-model="userInfo.email_pre"
                    placeholder="请输入邮箱"
                    required
                  />
                  <select class="email-select" v-model="userInfo.email_suffix">
                     <option selected="selected" value="@100bm.cn">@100bm.cn</option>
                     <option value="@hztx18.com">@hztx18.com</option>
                  </select>
                  <!-- <div class="form-heigit">
                    <span v-show="errors.first('email1')" class="text-danger">请输入正确的邮箱！</span>
                  </div> -->
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
                    <div class="form-group col-sm-5" style="margin-right:30px">
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
                    <div class="form-group col-sm-5">
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
                      class="form-group del-btn col-sm-1"
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
          <el-table-column align="center" width="100" prop="user_name" label="登录名"></el-table-column>
          <el-table-column align="center" width="100" prop="full_name" label="姓名"></el-table-column>
          <el-table-column align="center" width="350" prop="rolestr" label="系统/角色"></el-table-column>
          <el-table-column align="center" width="130" prop="mobile" label="联系电话"></el-table-column>
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
              <el-button plain type="success" size="mini" @click="userChange(0,scope.row.user_id, scope.row.user_name)" v-if="scope.row.status == 2">启用</el-button>
              <el-button plain type="info" size="mini" @click="userChange(2,scope.row.user_id,scope.row.user_name)" v-if="scope.row.status == 0">禁用</el-button>
              <el-button plain type="success" size="mini" @click="userChange(0,scope.row.user_id,scope.row.user_name)" v-if="scope.row.status == 1">解锁</el-button>
              <el-button plain type="danger" size="mini" @click="userDel(scope.row.user_id)">删除</el-button>
              <el-button plain type="danger" size="mini" v-if="!scope.row.wx_openid && scope.row.status == 0" @click="bindWx(scope.row.user_id, scope.row.user_name)">绑定微信</el-button>
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
import {trimError} from '@/services/util'
import "@/services/qrcode.min.js"

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
        909: "此登录名已被使用",
        918: "此姓名名已被使用"
      },
      paging: {
        ps: 10,
        pi: 1,
        username: "",
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
        user_id: -1,
        lists: [],
        mobile: null,
        email: null,
        email_pre:null,
        email_suffix:"@100bm.cn",
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
  watch:{
    'userInfo.full_name': {
      handler(newValue, oldValue) {
        console.log(newValue);
        if (!/^[\u4E00-\u9FA5]+[1-9]?$/.test(newValue)) {
          return
        }

        if (newValue && newValue.length >= 2) {
          this.$http.post("/user/generateusername", {full_name: newValue})
          .then(res => {
            this.userInfo.user_name = res;
            this.userInfo.email_pre = res;
          })
        }
      }
    }
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
        this.userInfo.full_name = "";
        this.userInfo.user_name = "";
        this.userInfo.role_id = "";
        this.userInfo.mobile = null;
        this.userInfo.status = 0;
        this.userInfo.user_id = -1;
        this.userInfo.is_add = 1;
        this.userInfo.lists = [{sys_id: "",role_id: ""}];
        this.userInfo.email = null;
        this.userInfo.ext_params = "",
        this.userInfo.email_pre = "";
        this.userInfo.email_suffix = "@100bm.cn";
        this.selectSys.push("");
      } else {
        // 编辑用户
        this.isAdd = 2;
        this.userInfo.full_name = j.full_name;
        this.userInfo.user_name = j.user_name;
        this.userInfo.role_id = j.role_id;
        this.userInfo.mobile = j.mobile;
        this.userInfo.status = j.status;
        this.userInfo.user_id = j.user_id;
        this.userInfo.lists = j.roles || [];
        this.userInfo.is_add = 2;
        this.userInfo.email = j.email;
        this.userInfo.email_pre = j.email ? j.email.split("@")[0] : "";
        this.userInfo.email_suffix = j.email ? "@" + j.email.split("@")[1] : "";
        this.userInfo.ext_params = j.ext_params;
        (j.roles || []).forEach(item => {
          this.selectSys.push(item.sys_id);
        });
        this.setSys();
      }
      this.$refs.editModal.open();
    },
    userChange(status, userid,user_name) {
      var r;
      this.$confirm("确认执行该操作吗?", "提示", {
        confirmButtonText: "确定",
        cancelButtonText: "取消",
        type: "warning"
      })
      .then(() => {
        this.$http.post("/user/changestatus", {user_id: userid, status: status, user_name:user_name})
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

    bindWx(userid, userName) {
      jQuery('#qrcodeTable canvas').remove();

      this.$http.post("/user/generateqrcode", {user_id: userid})
        .then(res => {
          console.log(window.globalConfig.loginWebHost + "/bindwx?userid=" + res.user_id + "&sign=" + res.sign + "&timestamp=" + res.timestamp + "&name=" + userName);
          jQuery('#qrcodeTable').qrcode(window.globalConfig.loginWebHost+ "/bindwx?userid=" + res.user_id + "&sign=" + res.sign + "&timestamp=" + res.timestamp + "&name=" + userName);

          // console.log("http://192.168.5.78:8091" + "/bindwx?userid=" + res.user_id + "&sign=" + res.sign + "&timestamp=" + res.timestamp + "&name=" + userName);
          // jQuery('#qrcodeTable').qrcode("http://192.168.5.78:8091" + "/bindwx?userid=" + res.user_id + "&sign=" + res.sign + "&timestamp=" + res.timestamp + "&name=" + userName);

          this.$refs.qrCodeModal.open();
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
          console.log(err)
        });
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
    onQrCodeClose() {
      jQuery('#qrcodeTable canvas').remove();
      this.$refs.qrCodeModal.close();
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

      if (!this.userInfo.email_pre || !this.userInfo.email_suffix) {
        this.$notify({
          title: "警告",
          message: "请输入正确的邮箱地址",
          type: "warning",
          offset: 50
        });
        return false;
      }

      var s = "";
      for (var i = 0; i < this.userInfo.lists.length; i++) {
        s = s + this.userInfo.lists[i].sys_id + "," + this.userInfo.lists[i].role_id +"|";
      }
      this.userInfo.auth = s;
      this.userInfo.email = this.userInfo.email_pre + this.userInfo.email_suffix;
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

<style>
.page-pagination {
  padding: 10px 15px;
  text-align: right;
}

#qrCodeModal .modal-footer {
  display: none;
}
#qrCodeModal .modal-body {
  margin-left: -7px;
}

.email-input {
  width: 80%;
  height: 34px;
  padding: 6px 12px;
  background-color: #fff;
  background-image: none;
  border: 1px solid #ccc;
}

.email-select {
  height: 34px;
  background-color: #fff;
  border: 1px solid #ccc;
}
 </style>
