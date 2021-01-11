//数据权限管理,查询，添加，编辑
<template>
  <div ref="main">
    <div class="panel panel-default">
      <div class="panel panel-default">
        <div class="panel-body">
          <form class="form-inline">
            <div class="form-group">
              <input
                name="queryName"
                type="text"
                class="form-control"
                v-model="queryName"
                placeholder="请输入规则名"
                maxlength="64"
              />
            </div>
            <div class="form-group">
              <input
                name="queryTableName"
                type="text"
                class="form-control"
                v-model="queryTableName"
                placeholder="请输入表名"
                maxlength="64"
              />
            </div>
            <a
              class="visible-xs-inline visible-sm-inline visible-md-inline visible-lg-inline btn btn btn-success"
              @click="querySearch"
            >查询</a>
            <a
              class="visible-xs-inline visible-sm-inline visible-md-inline visible-lg-inline btn btn btn-success"
              @click="Add"
            >添加</a>
          </form>
        </div>
      </div>
      <el-scrollbar style="height:100%">
        <el-table :data="datalist" stripe style="width: 100%">
          <el-table-column width="200" prop="name" label="规则名称" align="center"></el-table-column>
          <el-table-column width="180" prop="table_name" label="表名" align="center"></el-table-column>
          <el-table-column width="100" prop="operate_action" label="操作动作" align="center">
            <template slot-scope="scope">
              <span>{{scope.row.operate_action | EnumFilter("operate_action") }}</span>
            </template>
          </el-table-column>
          <el-table-column width="100" prop="status" label="状态" align="center">
            <template slot-scope="scope">
              <el-tag type="info" v-if="scope.row.status == 1">禁用</el-tag>
              <el-tag type="success" v-if="scope.row.status == 0">启用</el-tag>
            </template>
          </el-table-column>
          <el-table-column width="500" prop="rules" label="规则内容">
            <template slot-scope="scope">{{scope.row.rules | subStr}}</template>
          </el-table-column>
          <el-table-column width="250" prop="remark" label="备注" align="center">
            <template slot-scope="scope">
              <span>{{scope.row.remark | EllipsisFilter(20)}}</span>
            </template>
          </el-table-column>
          <el-table-column label="操作">
            <template slot-scope="scope">
              <el-button plain type="primary" size="mini" @click="edit(scope.row.id)">编辑</el-button>
              <el-button
                plain
                type="success"
                size="mini"
                @click="enable(scope.row.id,0)"
                v-if="scope.row.status == 1"
              >启用</el-button>
              <el-button
                plain
                type="info"
                size="mini"
                @click="disable(scope.row.id,1)"
                v-if="scope.row.status == 0"
              >禁用</el-button>
              <el-button plain type="danger" size="mini" @click="deleteById(scope.row.id)">删除</el-button>
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
          :total="datacount"
        ></el-pagination>
      </div>

      <bootstrap-modal
        ref="editModal"
        :need-header="true"
        :need-footer="true"
        no-close-on-backdrop="true"
      >
        <div slot="title">{{isAdd ? "新增" : "编辑"}}</div>
        <div slot="body">
          <div class="panel panel-default">
            <div class="panel-body">
              <form role="form" class="ng-pristine ng-valid ng-submitted height-min">
                <div class="form-group">
                  <label>规则名称(必填)</label>
                  <el-input v-model="permissionData.name" placeholder="请输入名称" maxlength="64"></el-input>
                </div>
                <el-row :span="24">
                  <el-col :span="12">
                    <div class="form-group">
                      <label>表名(必填)</label>
                      <el-input
                        v-model="permissionData.table_name"
                        placeholder="请输入表名"
                        :disabled="!isAdd"
                        maxlength="64"
                      ></el-input>
                    </div>
                  </el-col>
                  <el-col :span="12">
                    <div class="form-group" style="margin-left:10px;">
                      <label>操作动作(非必填)</label>
                      <el-select v-model="permissionData.operate_action" placeholder="请选择操作">
                        <el-option key label="选择操作动作" value></el-option>
                        <el-option
                          v-for="item in operateActionList"
                          :key="item.value"
                          :label="item.name"
                          :value="item.value"
                        ></el-option>
                      </el-select>
                    </div>
                  </el-col>
                </el-row>
                <div class="form-group">
                  <label>备注(非必填)</label>
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
                <div>
                  <div class="form-inline">
                    <label>规则配置</label>
                    <span class="add-btn">
                      <a class="btn m-b-xs btn-xs btn-success" @click="addRole">
                        <i class="fa fa-plus"></i>
                      </a>
                    </span>
                  </div>
                  <div style="max-height:300px;overflow-y:scroll">
                    <div
                      class="form-inline pull-in clearfix"
                      v-for="(item,index) in ruleslist"
                      v-bind:key="item.id"
                    >
                      <el-row>
                        <!-- <el-col :span="3"> -->
                        <div class="form-group" style="width:80px;margin-left:4px;margin-top:4px">
                          <el-select v-model="item.conlink_symbol" :disabled="index == 0">
                            <el-option key label="链接" value></el-option>
                            <el-option key="and" label="并且" value="and"></el-option>
                            <el-option key="or" label="或者" value="or"></el-option>
                          </el-select>
                        </div>
                        <!-- </el-col> -->
                        <!-- <el-col :span="5"> -->
                        <div class="form-group" style="width:110px;margin-left:4px;margin-top:4px">
                          <el-input v-model="item.field_name" placeholder="请输入字段名" maxlength="32"></el-input>
                        </div>
                        <!-- </el-col> -->
                        <!-- <el-col :span="4"> -->
                        <div class="form-group" style="width:80px;margin-left:4px;margin-top:4px">
                          <el-select
                            v-model="item.field_type"
                            @change="fieldTypeChange(item.field_name,item.field_type)"
                          >
                            <el-option
                              v-for="item in fieldTypeList"
                              :key="item.value"
                              :label="item.label"
                              :value="item.value"
                            ></el-option>
                          </el-select>
                        </div>
                        <!-- </el-col> -->
                        <!-- <el-col :span="4"> -->
                        <div class="form-group" style="width:90px;margin-left:4px;margin-top:4px">
                          <el-select
                            v-model="item.compare_symbol"
                            :disabled="item.compareSymbolDisabled"
                          >
                            <el-option
                              v-for="item in item.compareSymbolList"
                              :key="item.value"
                              :label="item.label"
                              :value="item.value"
                              :disabled="item.disabled"
                            ></el-option>
                          </el-select>
                        </div>
                        <!-- </el-col> -->
                        <!-- <el-col :span="5"> -->
                        <div class="form-group" style="width:110px;margin-left:4px;margin-top:4px">
                          <el-input v-model="item.value" placeholder="请输入值"></el-input>
                        </div>
                        <!-- </el-col> -->
                        <!-- <el-col :span="2"> -->
                        <div
                          class="form-group del-btn"
                          v-if="ruleslist.length > 1"
                          style="margin-left:4px; margin-top:4px"
                        >
                          <a class="btn m-b-xs btn-xs btn-danger" @click="del(index)">
                            <i class="fa fa-minus"></i>
                          </a>
                        </div>
                        <!-- </el-col > -->
                      </el-row>
                    </div>
                  </div>
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
import pager from "vue-simple-pager";
import PullTo from "vue-pull-to";
export default {
  components: {
    "bootstrap-modal": require("vue2-bootstrap-modal"),
    pager,
    PullTo
  },
  data() {
    return {
      queryName: null,
      queryTableName: null,
      datalist: null,
      pageSizeList: [5, 10, 20, 50], //可选显示数据条数
      datacount: 0,
      sysId: 0,
      tablelist: null, //表下拉数据
      table_type: "",
      isAdd: true,
      sysname: "",
      permissionData: {
        name: "",
        table_name: "",
        operate_action: "",
        rules: "",
        remark: ""
      },
      fieldTypeList: [
        { label: "类型", value: "" },
        { label: "数字", value: "number" },
        { label: "字符", value: "string" }
      ],
      compareSymbolList: [
        { label: "比较符", value: "" },
        { label: "=", value: "=" },
        { label: ">", value: ">" },
        { label: ">=", value: ">=" },
        { label: "<", value: "<" },
        { label: "<=", value: "<=" },
        { label: "<>", value: "<>" },
        { label: "in", value: "in" },
        { label: "notin", value: "not in" },
        { label: "like", value: "like" }
      ],
      ruleslist: [],
      pi: 1,
      ps: 10,
      totalPage: 0,
      compareSymbolDisabled: true,
      operateActionList: this.EnumUtility.Get("operate_action")
    };
  },
  props: ["path"],
  mounted() {
    this.sysId = this.$route.params.id;
    this.$refs.main.style.height = document.documentElement.clientHeight + "px";
    //this.getDictionInfo("operate_action");
    this.query();
  },
  methods: {
    Add() {
      this.isAdd = true;
      this.permissionData = {};
      this.ruleslist = [];
      this.$refs.editModal.open();
    },
    handleSizeChange(val) {
      this.ps = val;
      this.query();
    },
    goPage(val) {
      this.pi = val;
      this.query();
    },
    onClose() {
      this.$refs.editModal.close();
    },
    //点查询事件
    querySearch() {
      this.pi = 1;
      this.ps = 10;
      this.query();
    },

    query() {
      this.$http
        .post("/system/permission/getall", {
          pi: this.pi,
          ps: this.ps,
          sys_id: this.sysId,
          name: this.queryName,
          table_name: this.queryTableName
        })
        .then(res => {
          this.datalist = res.list;
          this.datacount = res.count;
          this.totalPage = Math.ceil(res.count / 10);
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

    savePermissionData() {
      if (!this.checkBeforSave()) {
        return;
      }
      var reallyRules = this.constructRules();
      if (reallyRules.length == 0) {
        this.$notify({
          title: "提示",
          message: "至少配置一条完整的规则数据",
          type: "error",
          offset: 50,
          duration: 2000
        });
        return;
      }
      this.permissionData.sys_id = this.sysId;
      this.permissionData.rules = JSON.stringify(reallyRules);
      if (this.isAdd) {
        this.$http
          .post("/system/permission/add", this.permissionData)
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
        this.$http
          .post("/system/permission/edit", this.permissionData)
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
    },

    checkBeforSave() {
      if (!this.permissionData.name || !this.permissionData.table_name) {
        this.$notify({
          title: "提示",
          message: "必填字段不能为空",
          type: "error",
          offset: 50,
          duration: 2000
        });
        return false;
      }
      if (this.ruleslist.length == 0) {
        this.$notify({
          title: "提示",
          message: "请配置规则数据",
          type: "error",
          offset: 50,
          duration: 2000
        });
        return false;
      }
      for (var i = 0; i < this.ruleslist.length; i++) {
        var temp = this.ruleslist[i];
        if (
          !temp.field_name ||
          !temp.compare_symbol ||
          !temp.field_type ||
          !temp.value ||
          (i != 0 && !temp.conlink_symbol)
        ) {
          this.$notify({
            title: "提示",
            message: "第(" + (i + 1) + ")行规则数据请填完整",
            type: "error",
            offset: 50,
            duration: 2000
          });
          return false;
        }
      }
      return true;
    },

    // field_name => f
    // compare_symbol => c
    // field_type => t
    // conlink_symbol => s
    // value => v

    constructRules() {
      var result = [];
      for (var i = 0; i < this.ruleslist.length; i++) {
        var temp = this.ruleslist[i];
        if (
          temp.field_name &&
          temp.value &&
          temp.field_type &&
          temp.compare_symbol
        ) {
          result.push({
            id: temp.id,
            f: temp.field_name,
            c: temp.compare_symbol,
            t: temp.field_type,
            s: temp.conlink_symbol,
            v: temp.value
          });
        }
      }
      return result;
    },

    deleteById(id) {
      this.$confirm("确定执行此操作?是否继续?", "提示", {
        confirmButtonText: "确定",
        cancelButtonText: "取消",
        type: "warning"
      }).then(() => {
        this.$http
          .post("/system/permission/del", { id: id })
          .then(res => {
            this.goPage(this.pi);
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

    edit(id) {
      for (var index = 0; index < this.datalist.length; index++) {
        if (this.datalist[index].id == id) {
          var temp = this.datalist[index];
          this.permissionData = temp;
          var tempRules = JSON.parse(temp.rules);
          this.ruleslist = [];
          tempRules.forEach(item => {
            this.ruleslist.push({
              field_name: item.f,
              compare_symbol: item.c,
              field_type: item.t,
              conlink_symbol: item.s,
              value: item.v
            });
          });
          this.ruleslist.forEach(rule => {
            rule.compareSymbolList = JSON.parse(
              JSON.stringify(this.compareSymbolList)
            );
            rule.compareSymbolList.forEach(item => {
              item.disabled = false;
              if (rule.field_type == "number") {
                item.disabled = item.value == "like";
              } else if (rule.field_type == "string") {
                item.disabled =
                  item.value == ">=" ||
                  item.value == ">" ||
                  item.value == "<=" ||
                  item.value == "<";
              }
            });
          });
          break;
        }
      }
      this.isAdd = false;
      this.$refs.editModal.open();
    },

    addRole() {
      if (this.ruleslist.length > 20) {
        return;
      }

      this.ruleslist.push({
        id: Date.now(),
        name: "",
        field_name: "",
        compare_symbol: "",
        field_type: "",
        conlink_symbol: "and",
        compareSymbolList: JSON.parse(JSON.stringify(this.compareSymbolList)),
        compareSymbolDisabled: true
      });
      var length = this.ruleslist.length;
      if (this.ruleslist.length == 1) {
        this.ruleslist[length - 1].conlink_symbol = "";
      }
    },

    del(index) {
      this.ruleslist.splice(index, 1);
      if (this.ruleslist[0].conlink_symbol != "") {
        this.ruleslist[0].conlink_symbol = "";
      }
    },

    enable(id) {
      this.$http
        .post("/system/permission/enable", {
          id: id
        })
        .then(res => {
          this.query();
        });
    },

    disable(id) {
      this.$http
        .post("/system/permission/disable", {
          id: id
        })
        .then(res => {
          this.query();
        });
    },

    getDictionInfo(type) {
      this.$http.get("/dds/dictionary/get?dic_type=" + type, {}).then(res => {
        this.operateActionList = res;
      });
    },

    fieldTypeChange(field_name, fieldType) {
      this.ruleslist.forEach(rule => {
        if (rule.field_name == field_name) {
          rule.compare_symbol = "";
          rule.compareSymbolDisabled = false;
          if (fieldType != "number" && fieldType != "string") {
            rule.compareSymbolDisabled = true;
          }
          console.log(rule.compareSymbolDisabled);

          rule.compareSymbolList.forEach(item => {
            item.disabled = false;
            if (fieldType == "number") {
              item.disabled = item.value == "like";
            } else if (fieldType == "string") {
              item.disabled =
                item.value == ">=" ||
                item.value == ">" ||
                item.value == "<=" ||
                item.value == "<";
            }
          });
        }
      });
    }
  },
  filters: {
    subStr(value) {
      if (value.length > 350) {
        return value.substring(0, 350) + " ...";
      }
      return value;
    }
  }
};
</script>

<style scoped>
.page-pagination {
  padding: 10px 15px;
  text-align: right;
}

/* id => id, 字段对应
field_name => f
compare_symbol => c
field_type => t
conlink_symbol => s
value => v */
</style>

