//系统管理,查询，添加，编辑，禁用
<template>
  <div ref="main">

    <div class="panel panel-default">
      <div class="panel panel-default">
        <div class="panel-body">
          <form class="form-inline">
            <div class="form-group">
              <label class="sr-only">标题</label>
              <input type="text" class="form-control" v-model="title"  placeholder="请输入标题">
            </div>
            <a class="visible-xs-inline visible-sm-inline visible-md-inline  visible-lg-inline btn btn btn-success" @click="query" >查询</a>
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
            <th>标题</th>
            <th>等级</th>
            <th>内容</th>
            <th class="visible-md  visible-lg">状态</th>
            <th class="visible-md  visible-lg">已发送次数</th>
            <th class="visible-md  visible-lg">创建时间</th>
            <th class="visible-md  visible-lg">完成时间</th>
            <th>操作</th>
          </tr>
        </thead>
        <tbody>
            <tr v-for="(v,k) in dataList" :key="k">
              <th class="font-thin">{{v.id}}</th>
              <th class="font-thin">{{v.title}}</th>
              <th class="font-thin">{{v.level_id}}</th>
              <th class="font-thin">{{v.content}}</th>
              <th class="font-thin" v-if="v.status == 1">等待</th>
              <th class="font-thin" v-if="v.status == 2">正在</th>
              <th v-if="v.status == 0" class="text-success font-thin">成功</th>
              <th v-if="v.status == 9" class="text-danger font-thin">失败</th>
              <th v-if="v.send_count <= 3" class="text-success text-center-xs font-thin">{{v.send_count}}</th>
              <th v-if="v.send_count >= 4" class="text-danger text-center-xs font-thin">{{v.send_count}}</th>
              <th class="font-thin">{{v.create_times}}</th>
              <th class="font-thin">{{v.finish_times}}</th>
              <th class="font-thin">
                <div class="form-group form-inline">
                  <div class="form-group">
                    <!--<button class="btn btn-xs btn-primary visible-md visible-lg" @click="edit(item.id)">编辑</button>-->
                  </div>
                  <div class="form-group">
                    <a class="btn btn-xs btn-danger visible-md visible-lg" @click="deleteById(v.id)">删除</a>
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
      title: ""
    };
  },
  mounted() {
    this.goPage({page:1})
  },
  methods: {
    goPage(data){
      this.pi = data.page;
      this.$fetch("/sso/notify/info", {
        title:this.title,
        pi: data.page ,
        ps:this.ps})
        .then(res => {
          console.log(res)
          this.dataList = res.list;
          this.datacount = res.count;
          this.totalPage = Math.ceil(res.count / 10);
        })
        .catch(err => {
          if (err.response) {
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
          }
        });
    },
    query() {

      this.goPage({page:this.pi})
    },
    deleteById(id){
      this.$confirm('此操作将永久删除该消息, 是否继续?', '提示', {
        confirmButtonText: '确定',
        cancelButtonText: '取消',
        type: 'warning'
      }).then(() => {
        this.$del("/sso/notify/info", {data:{id:id}})
          .then(res => {
            this.goPage({page:this.pi})
            this.$notify({
              title: '成功',
              message: '删除成功',
              type: 'success',
              offset: 50,
              duration:2000,
            });
          })
          .catch(err => {
            if (err.response) {
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
            }
          });
      }).catch(() => {

      });

    }
  }
};
</script>

<style scoped>
  .list-number{display: inline-block;padding: 15px 0 3px 15px;}
  .list-page{display: inline-block;position: absolute;right: 15px;margin-top: -43px;}
</style>

