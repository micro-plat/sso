//系统管理,查询，添加，编辑，禁用
<template>
  <div ref="main">

    <div class="panel panel-default">
      <div class="panel panel-default">
        <div class="panel-body">
          <form class="form-inline">

            <a class="visible-xs-inline visible-sm-inline visible-md-inline  visible-lg-inline btn btn btn-success" @click="openAdd" >添加配置</a>
            <span ng-controller="ModalDemoCtrl">
              <script type="text/ng-template" id="myModalContent.html">
                <div ng-include="'src/pages/user/index/add.vue'"></div>
              </script>
            </span>
          </form>
        </div>
      </div>
      <div class="table-responsive">

         <table class="table table-striped m-b-none">
        <thead>
          <tr>
            <th>#</th>
            <th>关键字</th>
            <th>等级</th>
            <th class="visible-md  visible-lg">创建时间</th>
            <th>操作</th>
          </tr>
        </thead>
        <tbody>
            <tr v-for="(item,k) in dataList" :key="k">
              <th class="font-thin">{{item.id}}</th>
              <th class="font-thin">{{item.keywords}}</th>
              <th class="font-thin">{{item.level_id}}</th>
              <th class="font-thin">{{item.create_times}}</th>
              <th class="font-thin">
                <div class="form-group form-inline">
                  <div class="form-group">
                    <button class="btn btn-xs btn-primary visible-md visible-lg" @click="edit(item.id)">编辑</button>
                  </div>
                  <div class="form-group">
                    <a class="btn btn-xs btn-danger visible-md visible-lg" @click="deleteById(item.id)">删除</a>
                  </div>
                </div>
              </th>
            </tr>
        </tbody>
      </table>
        <div  style="float:right">

        </div>
      </div>
      <div class="form-group form-inline paging visible-lg-block visible-md-block">
        <div class="form-group">
          <div class="list-number">
            共 {{datacount}} 条记录 | 每页显示:
            <select id="ddlps" v-model="ps" @change="goPage({page:pi})">
              <option v-for="(psl,k) in pageSizeList" :key="k" :value="psl">{{psl}}</option>
            </select>
            条
          </div>
          <div class="list-page">
            <pager
              :total-page="totalPage"
              :init-page="pi"
              @go-page="goPage"></pager>
          </div>
        </div>

      </div>
    </div>
    <bootstrap-modal ref="addModal" :need-header="true" :need-footer="true">
      <div slot="title">
        {{type == "add" ? "添加配置" : "编辑配置"}}
      </div>
      <div slot="body">
        <div class="panel panel-default">
          <div class="panel-body">
            <form role="form" class="ng-pristine ng-valid ng-submitted">
              <div class="form-group">
                <label>关键字</label>
                <select  class="form-control visible-md  visible-lg" v-model="addData.keywords" v-validate="'required'" name="keywords">
                  <option value="" selected="selected" >---请选择关键字---</option>
                  <option value="数据库">数据库</option>
                  <option value="网络">网络</option>
                  <option value="参数">参数</option>
                </select>
                <div class="form-height text-danger"><span v-show="errors.first('keywords')">请选择关键字</span></div>
                <input class="form-control" placeholder=""  type="hidden">
              </div>
              <div class="form-group">
                <label>等级</label>
                <select  class="form-control visible-md  visible-lg" v-model="addData.level_id" v-validate="'required'" name="level">
                  <option value="" selected="selected" >---请选择等级---</option>
                  <option value="1">1</option>
                  <option value="2">2</option>
                  <option value="3">3</option>
                </select>
                <div class="form-height text-danger"><span v-show="errors.first('level')">请选择等级</span></div>
                <input class="form-control" placeholder=""    type="hidden">
              </div>
            </form>
          </div>
        </div>
      </div>
      <div slot="footer">
        <a class="btn btn-sm btn-danger" @click="post">提交</a>
        <a class="btn btn-sm btn-primary" @click="close">取消</a>
      </div>
    </bootstrap-modal>

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
      pageSizeList: [10, 20, 50, 100], //可选显示数据条数
      pi: 1,
      ps:10,
      totalPage: 2,
      datacount: 0,
      dataList: [],
      title: "",
      addData:{user_id:"",sys_id:"",keywords:"",level_id:""},
      type:"add",
      editID:null
    };
  },
  mounted() {
    this.goPage({page:1})
  },
  methods: {
    goPage(data){
      this.pi = data.page;
      this.$fetch("/sso/notify/settings", {
        title:this.title,
        pi: data.page ,
        ps:this.ps})
        .then(res => {
          this.dataList = res.list;
          this.datacount = res.count;
          this.totalPage = Math.ceil(res.count / 10);
        })
        .catch(err => {
          if (err.response) {
            if (err.response.status == 403) {
              this.$router.push("/member/login");
            }else{
              this.$notify({
                title: '错误',
                message: '网络错误,请稍后再试',
                type: 'error',
                offset: 50,
                duration:2000,
              });
            }
          }
        });
    },
    openAdd() {
      this.type = "add"
      this.addData = {user_id:"",sys_id:"",keywords:"",level_id:""}
      this.$refs.addModal.open()
    },
    post(){
      if (this.type == "add") {
        this.$put("/sso/notify/settings",this.addData)
          .then(res=>{
            this.close();
            this.$notify({
              title: '成功',
              message: '添加配置成功',
              type: 'success',
              offset: 50,
              duration:2000,
            });
            this.goPage({page:this.pi})
          }).catch(err=>{
          this.$notify({
            title: '错误',
            message: '请检查填写内容',
            type: 'error',
            offset: 50,
            duration:2000,
          });
        })
      }else if (this.type == "edit") {
        this.addData.id  = this.editID
        this.$post("/sso/notify/settings",this.addData)
          .then(res=>{
            this.close();
            this.$notify({
              title: '成功',
              message: '编辑配置成功',
              type: 'success',
              offset: 50,
              duration:2000,
            });
            this.goPage({page:this.pi})
          }).catch(err=>{
            console.log(err)
          this.$notify({
            title: '错误',
            message: '请检查填写内容',
            type: 'error',
            offset: 50,
            duration:2000,
          });
        })
      }

    },
    edit(id){
      this.type = "edit";
      this.editID = id;
      this.dataList.forEach((item, index) => {
        if (item.id == id) {
          this.addData = item;
          this.$refs.addModal.open();
        }
      });
    },
    deleteById(id){
      this.$confirm('此操作将永久删除该配置, 是否继续?', '提示', {
        confirmButtonText: '确定',
        cancelButtonText: '取消',
        type: 'warning'
      }).then(() => {
        this.$del("/sso/notify/settings",{ data: { id: id } })
          .then(res=>{
            this.$notify({
              title: '成功',
              message: '删除成功',
              type: 'success',
              offset: 50,
              duration:2000,
            });
            this.goPage({page:this.pi})
          }).catch(err=>{
          this.$notify({
            title: '错误',
            message: '网络错误',
            type: 'error',
            offset: 50,
            duration:2000,
          });
      }).catch(() => {
      });

      })
    },
    close(){
      this.$refs.addModal.close();
    }
  }
};
</script>

<style scoped>
  .list-number{display: inline-block;padding: 15px 0 3px 15px;}
  .list-page{display: inline-block;position: absolute;right: 15px;margin-top: -43px;}
</style>

