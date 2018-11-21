<!-- 用户信息列表，查询，启动，禁用，解锁 -->
<template>
  <div ref="main">

      <div class="panel panel-default">
        <div class="panel panel-default">
          <div class="panel-body">
            <form class="form-inline" role="form">
              <div class="form-group">
                <label class="sr-only" for="exampleInputEmail2">用户名</label>
                <input type="text" class="form-control" v-model="paging.username"  placeholder="请输入用户名">
              </div>
              <div class="form-group ">
                <label class="col-sm-2 control-label sr-only">角色</label>
                <select v-model="paging.role_id" name="role_id" class="form-control visible-md  visible-lg">
                  <option value="" selected="selected" >---请选择角色---</option>
                  <option v-for="(r, index) in roleList" :key="index" :value="r.role_id">{{r.name}}</option>
                </select>
              </div>
              <a class="btn btn-success" @click="searchClick" >查询</a>
              <a class="btn btn-primary" @click="showModal(1,{})" >添加用户</a>
            </form>
          </div>
        </div>
        <bootstrap-modal ref="editModal" :need-header="true" :need-footer="true" :closed="resetSys">
          <div slot="title">
            {{isAdd == 1 ? "添加用户" : "编辑用户信息"}}
          </div>
          <div slot="body">
            <div class="panel panel-default">
              <div class="panel-body">
                <form role="form" class="ng-pristine ng-valid ng-submitted height-min">
                  <div class="form-group">
                    <label>用户名</label>
                    <input name="username1" type="text" class="form-control" v-validate="'required'" v-model="userInfo.user_name" placeholder="请输入用户名" required >
                    <div class="form-heigit"> <span v-show="errors.first('username1')" class="text-danger">用户名不能为空！</span> </div>
                  </div>
                  <div class="form-group">
                    <label>电话</label>
                    <input name="mobile1" type="text" class="form-control" v-validate="'required|numeric'" v-model="userInfo.mobile" placeholder="请输入电话" maxlength="11" required>
                    <div class="form-heigit"><span v-show="errors.first('mobile1')" class="text-danger">请输入正确的11位手机号！</span></div>
                  </div>
                  <div class="form-group">
                    <label>邮箱</label>
                    <input name="email1" type="text" class="form-control" v-validate="'required|email'" v-model="userInfo.email" placeholder="请输入邮箱" required>
                    <div class="form-heigit"><span v-show="errors.first('email1')" class="text-danger">请输入正确的邮箱！</span></div>
                  </div>
                  <div class="form-group sel-col-5">
                    <div class="form-inline">
                      <label>系统与角色</label>
                      <span class="add-btn"><a class="btn m-b-xs btn-xs btn-success" @click="add"><i class="fa fa-plus"></i></a></span>
                    </div>
                    <form class="form-inline pull-in clearfix"
                          v-for='(list,index) in userInfo.lists'
                          v-bind:key='list.id'>
                      <div class="form-group col-sm-5" >
                        <select name="select1" class="form-control" v-validate="'required'" v-model="list.sys_id" @change="sysStatus(list.sys_id,index)" required >
                          <option value="" selected="selected" >---请选择系统---</option>
                          <option v-for="(r, index) in sysList" :key="index" :value="r.id" :disabled="r.disabled">{{r.name}}</option>
                        </select>
                      </div>
                      <div class="form-group col-sm-5">
                        <select name="select2" class="form-control" v-validate="'required'" v-model="list.role_id" required >
                          <option value="" selected="selected" >---请选择角色---</option>
                          <option v-for="(r, index) in roleList" :key="index" :value="r.role_id">{{r.name}}</option>
                        </select>
                      </div>
                      <div class="form-group del-btn" v-if="userInfo.lists.length > 1">
                        <a class="btn m-b-xs btn-xs btn-danger" @click="del(index)"><i class="fa fa-minus"></i></a>
                      </div>
                    </form>
                    <div class="">
                      <div class="form-group form-heigit col-sm-5"><span v-show="errors.has('select1')" class="text-danger">必须选择系统！</span></div>
                      <div class="form-group form-heigit col-sm-5"><span v-show="errors.has('select2')" class="text-danger">必须选择用户角色！</span></div>
                    </div>
                  </div>
                  <div class="form-group" v-if="isAdd == 1">
                    <label class="checkbox-inline">
                      <input id="statuscheck" type="checkbox">是否启用
                    </label>
                  </div>
                </form>
              </div>
            </div>
          </div>
          <div slot="footer">
            <a class="btn btn-default" @click="onClose">取消</a>
            <a class="btn btn-success" @click="submitUser">提交</a>
          </div>
        </bootstrap-modal>

        <div class="table-responsive">

          <table class="table table-striped m-b-none">
            <thead>
            <tr>
              <th>用户名</th>
              <th  class="visible-md-block  visible-lg-block">系统/角色</th>
              <th class="border-no">联系电话</th>
              <th class="visible-md-block  visible-lg-block border-no">邮箱</th>
              <th>状态</th>
              <th class="visible-md-block visible-lg-block border-no">创建时间</th>
              <th>操作</th>
            </tr>
            </thead>

            <tbody class="table-border">
            <tr v-for="(item, index) in datalist.items" :key="index">
              <td>{{item.user_name}}</td>
              <td class="visible-md-block visible-lg-block over-text" :title="item.rolestr">{{item.rolestr}}</td>
              <td >{{item.mobile}}</td>
              <td class="visible-md-block visible-lg-block">{{item.email}}</td>
              <td v-if="item.status==2" class="text-danger">{{item.status_label}}</td>
              <td v-if="item.status==1" class="text-warning">{{item.status_label}}</td>
              <td v-if="item.status==0" class="text-success">{{item.status_label}}</td>
              <td class="visible-md-block  visible-lg-block">{{item.create_time}}</td>
              <td>
                <div class="form-inline">
                  <div class="form-group">
                    <button class="btn btn-xs btn-primary" @click="showModal(2,item)">编辑</button>
                  </div>
                  <div class="form-group">
                    <button class="btn btn-xs btn-warning" @click="userChange(2,item.user_id)" v-if="item.status==0" >禁用</button>
                    <button class="btn btn-xs btn-warning" @click="userChange(0,item.user_id)" v-if="item.status==2" >启用</button>
                    <button class="btn btn-xs btn-warning" @click="userChange(11,item.user_id)" v-if="item.status==1" >解锁</button>
                  </div>
                  <div class="form-group">
                    <button class="btn btn-xs btn-danger visible-md visible-lg" @click="userDel(item.user_id)">删除</button>
                  </div>
                </div>
              </td>
            </tr>
            </tbody>

          </table>
          <div class="form-group form-inline paging visible-lg-block visible-md-block">
            <div class="form-group">

              <div class="list-number">
                共 {{datalist.count}} 条记录 | 每页显示:
                <select id="ddlps" v-model="paging.ps" @change="searchClick">
                  <option v-for="psl in pageSizeList" :key="psl.id" :value="psl">{{psl}}</option>
                </select>
                条
              </div>

              <div class="list-page">
                <pager class="visible-md visible-lg"
                       :total-page="totalpage"
                       :init-page="paging.pi"
                       :showItems="5"
                       @go-page="pageChange">
                </pager>
              </div>

            </div>

          </div>
        </div>


      </div>

  </div>
</template>
<script>
import pager from "vue-simple-pager";
import PullTo from 'vue-pull-to'
export default {
  components: {
    "bootstrap-modal": require("vue2-bootstrap-modal"),
    pager: pager,
    PullTo
  },
  data() {
    return {
      paging: { ps: 10, pi: 1, username: "", role_id: "" },
      pageSizeList: [5, 10, 20, 50], //可选显示数据条数
      datalist: { count: 0, items: [] },
      userInfo: {
        user_name: "",
        user_id: -1,
        lists: [],
        mobile: null,
        email: null,
        status: 0,
        is_add: 2,
      },
      totalpage: 0,
      sysList: [],
      roleList: [], //角色列表
      selectSys: [],
      isAdd: 1,
    };
  },
  created(){

  },
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
      this.$post("/sso/user", this.paging)
        .then(res => {
          this.datalist.items = res.list;
          this.datalist.count = new Number(res.count);
          this.totalpage = Math.ceil(this.datalist.count / this.paging.ps);
        })
        .catch(err => {
          if (err.response.status == 403) {
            this.$router.push("/login");
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
    },
    pageChange: function(data) {
      this.paging.pi = data.page;
      this.queryData();
    },

    searchClick: function() {
      this.paging.pi = 1;
      this.queryData();
    },
    next(){
      let pi = this.paging.pi
      this.paging.pi = pi + 1;
      this.$post("/sso/user", this.paging)
        .then(res => {
          if(res.list.length <= 0) {
            this.paging.pi = pi
            return false
          }
          this.datalist.items = this.datalist.items.concat(res.list);
          this.datalist.count = new Number(res.count);
          this.totalpage = Math.ceil(this.datalist.count / this.paging.ps);
        })
        .catch(err => {
          if (err.response.status == 403) {
            this.$router.push("/login");
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
    },
    stateChange(e){
      console.log(e)
    },
    showModal: function(i, j) {
      if (j.status == 1) {
          this.$notify({
            title: '警告',
            message: '请先进行解锁操作',
            type: 'warning',
            offset: 50,
            duration:2000
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
        this.userInfo.lists = [{ sys_id: "", role_id: "" }];
        this.userInfo.email = "";
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
        for (var s = 0; s < j.roles.length; s++) {
          this.selectSys.push(j.roles[s].sys_id);
        }
        this.setSys();
      }
      this.$refs.editModal.open();
    },
    userChange: function(ests, userid) {
      var r = confirm("确认执行该操作吗？");
      if (r == true) {
        var user = { status: ests, user_id: userid };
        this.$put("/sso/user", user)
          .then(res => {
            this.queryData();
            this.$notify({
              title: '成功',
              message: '修改状态成功',
              type: 'success',
              offset: 50,
              duration: 2000
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
                onClose: () => {
                  this.$router.push("/login");
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
      } else {
        return false;
      }
    },
    userDel: function(userid) {
      var r = confirm("警告！确认删除该用户吗？");
      if (!r) {
        return false;
      }
      var user = { user_id: userid };
      this.$post("/sso/user/delete", user)
        .then(res => {
          this.queryData();
          this.$notify({
            title: '成功',
            message: '成功删除用户',
            type: 'success',
            offset: 50,
            duration:2000
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
              onClose:() => {
                this.$router.push("/login");
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
          title: '警告',
          message: '至少需要添加一个系统角色！',
          type: 'warning',
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
            this.$post("/sso/user/save", this.userInfo)
              .then(res => {
                this.queryData();
                  this.$notify({
                    title: '成功',
                    message: '添加成功',
                    type: 'success',
                    offset: 50,
                    duration:2000
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
                    onClose: () => {
                      this.$router.push("/login");
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
          }else if (this.isAdd == 2){
            this.$post("/sso/user/edit", this.userInfo)
              .then(res => {
                this.queryData();
                  this.$notify({
                    title: '成功',
                    message: '编辑成功',
                    type: 'success',
                    offset: 50,
                    duration:2000
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
                    onClose: () => {
                      this.$router.push("/login");
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

          this.onClose();
        }
      });
    },
    querySys() {
      this.$fetch("/sso/base", {})
        .then(res => {
          this.roleList = res.rolelist;
        })
        .catch(err => {
        });

      this.$post("/sso/base", {})
        .then(res => {
          this.sysList = res.list;
          for (s in this.sysList) {
            s.disabled = false;
          }
        })
        .catch(err => {
        });


    }
  }
};
</script>

<style scoped>
  .list-number{display: inline-block;padding: 15px 0 3px 15px;}
  .list-page{display: inline-block;position: absolute;right: 15px;margin-top: -43px;}
</style>

