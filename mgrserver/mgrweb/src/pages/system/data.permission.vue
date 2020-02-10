//系统管理,查询，添加，编辑，禁用
<template>
  <div ref="main">

    <div class="panel panel-default">
      <div class="panel panel-default">
        <div class="panel-body">
          <form class="form-inline">
            <div class="form-group">
              <label class="col-sm-2 control-label sr-only">状态</label>
              <select name="roleid" class="form-control not-100" v-model="data_type" @change="query">
                <option value="" selected="selected" >---请选择类型---</option>
                <option v-for="(s, index) in typelist" :key="index" :value="s.type">{{s.type_name}}</option>
            </select>
            </div>
            <a class="visible-xs-inline visible-sm-inline visible-md-inline  visible-lg-inline btn btn btn-success" @click="Add">添加</a>
          </form>
        </div>
      </div>
      <el-scrollbar style="height:100%">
        <el-table :data="datalist" stripe  style="width: 100%">
          <!-- <el-table-column width="100" prop="ident" label="系统名称" ></el-table-column> -->
          <el-table-column width="200" prop="type_name" label="类型名称" ></el-table-column>
          <el-table-column width="200" prop="type" label="类型" ></el-table-column>
          <el-table-column width="200" prop="name" label="名称" ></el-table-column>
          <el-table-column width="200" prop="value" label="值" ></el-table-column>
          <el-table-column width="300" prop="remark" label="备注" ></el-table-column>
          <el-table-column  label="操作">
            <template slot-scope="scope">
              <el-button plain type="primary" size="mini" @click="edit(scope.row.id)">编辑</el-button>
              <el-button plain  type="danger" size="mini" @click="deleteById(scope.row.id)">删除</el-button>
            </template>
          </el-table-column>
        </el-table>

      </el-scrollbar>
      <div class="page-pagination">
        <el-pagination
          @size-change="handleSizeChange"
          @current-change="goPage"
          :current-page="pi"
          :page-size="ps"
          :page-sizes="pageSizeList"
          layout="total, sizes, prev, pager, jumper"
          :total="datacount">
        </el-pagination>
      </div>

    <bootstrap-modal ref="editModal" :need-header="true" :need-footer="true" no-close-on-backdrop="true">
        <div slot="title">{{isAdd ? "新增" : "编辑"}}</div>
        <div slot="body">
          <div class="panel panel-default">
            <div class="panel-body">
              <form role="form" class="ng-pristine ng-valid ng-submitted height-min">
                <div class="form-group">
                  <label>类型</label>
                  <input
                    name="type"
                    type="text"
                    class="form-control"
                    v-validate="'required'"
                    v-model="permissionData.type"
                    placeholder="请输入类型"
                    required
                    maxlength="32"
                    :disabled="!isAdd" />
                  <div class="form-heigit">
                    <span v-show="errors.first('type')" class="text-danger">类型不能为空！</span>
                  </div>
                </div>
                <div class="form-group">
                  <label>类型名称</label>
                  <input
                    name="type_name"
                    type="text"
                    class="form-control"
                    v-validate="'required'"
                    v-model="permissionData.type_name"
                    placeholder="请输入类型名称"
                    maxlength="64"
                    required
                    :disabled="!isAdd"
                  />
                  <div class="form-heigit">
                    <span v-show="errors.first('type_name')" class="text-danger">请输入类型名称！</span>
                  </div>
                </div>
                <div class="form-group">
                  <label>名称</label>
                  <input
                    name="name"
                    type="text"
                    class="form-control"
                    v-validate="'required'"
                    v-model="permissionData.name"
                    placeholder="请输入名称"
                    required
                    maxlength="32"/>
                  <div class="form-heigit">
                    <span v-show="errors.first('name')" class="text-danger">名称不能为空！</span>
                  </div>
                </div>
                <div class="form-group">
                  <div><label>值</label></div>
                  <input
                    name="value"
                    type="text"
                    class="form-control"
                    v-validate="'required'"
                    v-model="permissionData.value"
                    placeholder="请输入值"
                    required />
                    <div class="form-heigit">
                    <span v-show="errors.first('value')" class="text-danger">值不能为空！</span>
                  </div>
                </div>
                <div class="form-group">
                  <label>备注</label>
                  <textarea
                    name="ext_params"
                    style="resize:none"
                    rows="5"
                    type="text"
                    class="form-control"
                    v-model="permissionData.remark"
                    placeholder="备注"
                  ></textarea>
                </div>
              </form>
            </div>
          </div>
        </div>
        <div slot="footer">
          <el-button size="small" @click="onClose">取消</el-button>
          <el-button type="success" size="small" @click="savePermissionData">保存</el-button>
        </div>
      </bootstrap-modal>
  </div>
  </div>
</template>

<script>
  import pager from "vue-simple-pager"
  import PullTo from 'vue-pull-to'
export default {
  components: {
    "bootstrap-modal": require("vue2-bootstrap-modal"),
    pager,
    PullTo
  },
  data() {
    return {
      datalist: null,
      pageSizeList: [5, 10, 20, 50], //可选显示数据条数
      datacount: 0,
      sysId: 0,
      typelist: null,
      data_type: "",
      isAdd: true,
      sysname: "",
      permissionData: {
        id: 0,
        name: "",
        type: "",
        type_name: "",
        Value: "",
        remark: ""
      },
      pi: 1,
      ps:10,
      totalPage: 0
    };
  },
  props:["path"],
  mounted() {
    this.sysId = this.$route.query.id;
    this.$refs.main.style.height = document.documentElement.clientHeight + 'px';
    this.querySystemTypeInfo();
    this.query()
  },
  methods: {
    querySystemTypeInfo() {
        this.$http.post("/base/getpermisstypes",{sys_id: this.sysId})
            .then(res => {
            this.typelist = res;
            if (this.typelist.length > 0) {
                this.query();
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

    Add() {
      this.isAdd = true;
      this.permissionData = {};
      this.$refs.editModal.open();
    },
    handleSizeChange(val){
      this.ps =val;
      this.query()
    },
    goPage(val) {
      this.pi = val;
      this.query()
    },
    onClose() {
      this.$refs.editModal.close();
    },
    //点查询事件
    querySearch() {
      this.pi = 1;
      this.query();
    },

    query() {
      this.$http.post("/system/permission/getall", {
        pi: this.pi,
        ps:this.ps,
        sys_id: this.sysId,
        data_type: this.data_type
      })
        .then(res => {
          this.datalist = res.list;
          this.datacount = res.count;
          this.totalPage = Math.ceil(res.count / 10);
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

    savePermissionData() {
      this.permissionData.sys_id = this.sysId;
      this.$validator.validate().then(result => {
        if (!result) {
          return false;
        } else {
          if (this.isAdd) {
            this.$http.post("/system/permission/add", this.permissionData)
              .then(res => {
                this.query();
                this.$notify({
                  title: "成功",
                  message: "添加成功",
                  type: "success",
                  offset: 50,
                  duration: 2000
                });
                this.onClose();
                this.querySystemTypeInfo();
              })
              .catch(err => {
                if (err.response) {
                  this.$notify({
                        title: "错误",
                        message: "出现错误,请稍后再试",
                        type: "error",
                        offset: 50,
                        duration: 2000
                      });
                }
              });
          } else {
            this.$http.post("/system/permission/edit", this.permissionData)
              .then(res => {
                this.query();
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
                   this.$notify({
                        title: "错误",
                        message: "出现错误,请稍后再试",
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

    deleteById(id) {
      this.$confirm("确定执行此操作?, 是否继续?", "提示", {
        confirmButtonText: "确定",
        cancelButtonText: "取消",
        type: "warning"
      }).then(() => {
        this.$http.post("/system/permission/del", { id: id } )
        .then(res => {
          this.querySystemTypeInfo();
          this.goPage(this.pi);
          this.$notify({
            title: '成功',
            message: '删除成功',
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
      });
    },

    edit(id) {
      for (var index = 0; index < this.datalist.length; index++) {
        if (this.datalist[index].id == id) {
          this.permissionData = this.datalist[index];
          break;
        }
      }
      this.isAdd = false;
      this.$refs.editModal.open();
    }
    
  }
};
</script>

<style scoped>
  .page-pagination{padding: 10px 15px;text-align: right;}
</style>