<!--角色基础信息查询-->
<template>
  <div ref="main">
    <pull-to @infinite-scroll="next">
    <div class="panel panel-default">
      <div class="panel panel-default">
        <div class="panel-body">
          <form class="form-inline" role="form">
            <div class="form-group">
              <label class="sr-only" for="exampleInputEmail2">角色名</label>
              <input type="text" class="form-control" v-model="paging.role_name"  placeholder="请输入角色名">
            </div>
            <a class="btn btn-success" @click="searchClick" >查询</a>
            <a class="btn btn-primary" @click="showModal(1,{})" >添加角色</a>
          </form>
        </div>
      </div>
      <bootstrap-modal ref="editModal" :need-header="true" :need-footer="true">
        <div slot="title">
          {{isAdd == 1 ? "添加用户" : "编辑用户信息"}}
        </div>
        <div slot="body">
          <div class="panel panel-default">
            <div class="panel-body">
              <form role="form" class="ng-pristine ng-valid ng-submitted height-min">
                <div class="form-group">
                  <label>角色名</label>
                  <input name="ename" type="text" class="form-control" v-validate="'required'" v-model="roleInfo.role_name" placeholder="请输入角色名" required >
                  <div class="form-heigit"><span v-show="errors.first('ename')" class="text-danger">角色名不能为空！</span></div>
                </div>
                <div class="form-group"  v-if="isAdd == 1">
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
            <th>角色名</th>
            <th>状态</th>
            <th class="visible-md-block visible-lg-block border-no">创建时间</th>
            <th>操作</th>
          </tr>
          </thead>
          <tbody class="table-border">
          <tr v-for="(item, index) in datalist.items" :key="index">
            <td>{{item.role_name}}</td>
            <td v-if="item.status==2" class="text-danger">{{item.status_label}}</td>
            <td v-if="item.status==0" class="text-success">{{item.status_label}}</td>
            <td class="visible-md-block  visible-lg-block">{{item.create_time}}</td>
            <td>
              <div class="form-inline">
                <div class="form-group">
                  <button class="btn btn-xs btn-primary visible-md visible-lg" @click="showModal(0,item)">编辑</button>
                </div>
                <div class="form-group">
                  <button class="btn btn-xs btn-warning" @click="roleChange(2,item.role_id)" v-if="item.status==0" >禁用</button>
                  <button class="btn btn-xs btn-warning" @click="roleChange(0,item.role_id)" v-if="item.status==2" >启用</button>
                </div>
                <div class="form-group">
                  <button class="btn btn-xs btn-danger visible-md visible-lg" @click="roleDel(item.role_id)">删除</button>
                </div>
                <div class="form-group">
                  <button class="btn btn-xs btn-default" @click="auth(item.role_id)">授权</button>
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
    </pull-to>
  </div>
</template>
<script>
import pager from "vue-simple-pager";
import PullTo from 'vue-pull-to';
export default {
  components: {
    "bootstrap-modal": require("vue2-bootstrap-modal"),
    pager: pager,
    PullTo
  },
  data() {
    return {
      paging: { ps: 10, pi: 1, role_name: "" },
      pageSizeList: [5, 10, 20, 50], //可选显示数据条数
      datalist: { count: 0, items: [] },
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
    this.$refs.main.style.height = document.documentElement.clientHeight + 'px';
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
    next(){
      let pi = this.paging.pi
      this.paging.pi = pi + 1;
      this.$post("/sso/role", this.paging)
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
    queryData: function() {
      if (this.paging.pi == 0) {
        this.paging.pi = 1;
      }
      this.$fetch("/sso/role", this.paging)
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
      this.$refs.editModal.open();
    },
    roleChange: function(ests, roleid) {
      var r = confirm("确认执行该操作吗？");
      if (!r) {
        return false;
      }
      var role = { status: ests, role_id: roleid };
      this.$put("/sso/role", role)
        .then(res => {
          this.queryData();
            this.$notify({
              title: '成功',
              message: '状态修改成功',
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
              onClose: function () {
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
    roleDel: function(roleid) {
      var r = confirm("警告！确认删除该用户吗？");
      if (!r) {
        return false;
      }
      var role = {data:{ role_id: roleid }};
      this.$del("/sso/role", role)
        .then(res => {
          this.queryData();
            this.$notify({
              title: '成功',
              message: '删除成功',
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
    onClose() {
      this.$refs.editModal.close();
    },
    submitUser() {
      this.$validator.validate().then(result => {
        if (!result) {
          return false;
        } else {
          if (this.isAdd == 1) {
            this.roleInfo.is_add = 1;
            var x = document.getElementById("statuscheck").checked;
            if (x == true) {
              this.roleInfo.status = 0;
            } else if (x == false) {
              this.roleInfo.status = 2;
            }
          }
          this.$post("/sso/role", this.roleInfo)
            .then(res => {
              this.queryData();
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
          this.onClose();
        }
      });
    },
    auth(id) {
      this.$router.push({
        name: "roleauth",
        query: {
          role_id: id
        }
      });
    }
  }
};
</script>

<style scoped>
  .list-number{display: inline-block;padding: 15px 0 3px 15px;}
  .list-page{display: inline-block;position: absolute;right: 15px;margin-top: -43px;}
</style>
