//系统管理,查询，添加，编辑，禁用
<template>
  <div ref="main">

    <div class="panel panel-default">
      <div class="panel panel-default">
        <div class="panel-body">
          <form class="form-inline">
            <div class="form-group">
              <label class="sr-only">系统名</label>
              <input type="text" class="form-control" onkeypress="if(event.keyCode == 13) return false;" v-model="sysname" placeholder="请输入系统名称">
            </div>
            <div class="form-group">
              <label class="col-sm-2 control-label sr-only">状态</label>
              <select  class="form-control visible-md  visible-lg" v-model="selected">
                <option value="" selected="selected" >---请选择状态---</option>
                <option value="0">禁用</option>
                <option value="1">启用</option>
              </select>
            </div>
            <a class="visible-xs-inline visible-sm-inline visible-md-inline  visible-lg-inline btn btn btn-success" @click="querySearch">查询</a>
            <span ng-controller="ModalDemoCtrl">
              <script type="text/ng-template" id="myModalContent.html">
                <div ng-include="'src/pages/user/index/add.vue'"></div>
              </script>
              <a class="visible-sm-inline visible-md-inline  visible-lg-inline btn btn-primary" ref="addsys" @click="Add">添加</a>
            </span>
          </form>
        </div>
      </div>
      <el-scrollbar style="height:100%">
        <el-table :data="datalist" stripe  style="width: 100%">
          <el-table-column width="130" prop="ident" label="英文名称" ></el-table-column>
          <el-table-column width="250" prop="name" label="系统名称" ></el-table-column>
          <el-table-column  width="150" prop="enable" label="状态" >
            <template slot-scope="scope">
              <el-tag type="info" v-if="scope.row.enable == 0">禁用</el-tag>
              <el-tag type="success" v-if="scope.row.enable == 1">启用</el-tag>
            </template>
          </el-table-column>
          <el-table-column width="200" prop="logo" label="logo" >
            <template slot-scope="scope">
              <img v-if="scope.row.theme" :class="scope.row.theme.split('|')[0]"
                   :src="scope.row.logo" :onerror="errorImg" alt="">
            </template>
          </el-table-column>
          <el-table-column width="300" prop="secret" label="secret" ></el-table-column>
          <el-table-column width="320" prop="callbackurl" label="登录回调地址" ></el-table-column>

          <el-table-column  label="操作">
            <template slot-scope="scope">
              <el-button plain type="primary" size="mini" @click="edit(scope.row.id)">编辑</el-button>
              <el-button plain type="success" size="mini" @click="enable(scope.row.id,1)" v-if="scope.row.enable == 0" >启用</el-button>

              <el-button plain type="info" size="mini" @click="enable(scope.row.id,0)" v-if="scope.row.enable == 1">禁用</el-button>

              <el-button plain  type="danger" size="mini" @click="deleteById(scope.row.id)">删除</el-button>

              <el-button plain  type="warning" size="mini" @click="manage(scope.row.id)">管理</el-button>

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
          layout="total, sizes, prev, pager, next, jumper"
          :total="datacount">
        </el-pagination>
      </div>

    <bootstrap-modal ref="theModal" :need-header="true" :need-footer="true" >
      <div slot="title">
        删除系统
      </div>
      <div slot="body">
        确定删除系统？
      </div>
      <div slot="footer">
        <a class="btn btn-sm btn-danger" @click="ok">确定</a>
        <a class="btn btn-sm btn-primary" @click="cancel">取消</a>
      </div>
    </bootstrap-modal>
    <bootstrap-modal ref="deletOkModal" :need-header="true" :need-footer="false" >
      <div slot="title">
        提示
      </div>
      <div slot="body">
        删除成功
      </div>
      <div slot="footer">
        <a class="btn btn-sm btn-primary" @click="cancel">确定</a>
      </div>
    </bootstrap-modal>
    <bootstrap-modal ref="editModal" :need-header="true" :need-footer="true">
      <div slot="title">
        编辑系统
      </div>
      <div slot="body">
        <div class="panel panel-default">
          <div class="panel-body">
            <form role="form" class="ng-pristine ng-valid ng-submitted">
               <el-row :span="24">
                 <el-col :span="12">
                   <div class="form-group">
                    <label>系统名称</label>
                    <input class="form-control" placeholder="" name="name2" v-model="editData.name"  type="text">
                    <div class="form-height text-danger"><span v-show="errors.first('name2')">系统名称不能为空</span></div>
                    <input class="form-control" placeholder=""  v-model="editData.id"  type="hidden">
                  </div>
                 </el-col>
                 <el-col :span="12">
                   <div class="form-group" style="margin-left:10px;">
                    <label>系统英文名称</label>
                    <input class="form-control" placeholder="" name="name-e" v-model="editData.ident"  type="text">
                    <div class="form-height text-danger"><span v-show="errors.first('name-e')">系统名称不能为空</span></div>
                    <input class="form-control" placeholder=""  v-model="editData.id"  type="hidden">
                  </div>
                 </el-col>
               </el-row>
              <div class="form-group">
                <label>secret</label>
                <input class="form-control" placeholder="系统签名的secret" v-model="editData.secret" name="secret"  type="text">
                <div class="form-height text-danger"><span v-show="errors.first('secret')">secret不能为空</span> </div>
              </div>
              <div class="form-group">
                <label>sso登录后回调子系统的地址(如:http://www.123.com/abc)</label>
                <input class="form-control" placeholder="请输入回调地址"  v-model="editData.callbackurl" name="callbackurl"  type="text">
              </div>
              <div class="form-group">
                <label>{{editData.logo}}</label>
                <input class="form-control" placeholder="" name="logo2" v-model="editData.logo"  type="hidden">
                <div class="form-height text-danger"><span v-show="errors.first('logo2')">logo地址不能为空</span></div>
                <uploader :options="options" class="uploader-example" :file-status-text="statusText" ref="uploader" @file-success="fileEditSuccess" @file-error="fileError">
                  <uploader-unsupport></uploader-unsupport>
                  <uploader-drop>
                    <p>上传logo</p>
                    <uploader-btn :attrs="attrs">选择图片</uploader-btn>
                  </uploader-drop>
                  <uploader-list></uploader-list>
                </uploader>
              </div>
              <div class="form-group">
                <div class="settings">
                  <div class="container">
                    <div class="row">
                      <div class="col col-lg-2-4 col-sm-6">
                        <div class="panel-body ng-scope">
                          <div class="m-b-sm">
                            <label class="i-switch bg-info pull-right">
                              <input type="checkbox" v-model="editData.layout" class="ng-pristine ng-untouched ng-valid"   value="app-header-fixed">
                              <i></i>
                            </label>
                            Fixed header
                          </div>
                          <div class="m-b-sm">
                            <label class="i-switch bg-info pull-right">
                              <input type="checkbox" v-model="editData.layout" class="ng-pristine ng-untouched ng-valid"   value="app-aside-fixed">
                              <i></i>
                            </label>
                            Fixed aside
                          </div>
                          <div class="m-b-sm">
                            <label class="i-switch bg-info pull-right">
                              <input type="checkbox" v-model="editData.layout" class="ng-pristine ng-untouched ng-valid"   value="app-aside-folded">
                              <i></i>
                            </label>
                            Folded aside
                          </div>
                          <div class="m-b-sm">
                            <label class="i-switch bg-info pull-right">
                              <input type="checkbox" v-model="editData.layout" class="ng-pristine ng-untouched ng-valid"   value="app-aside-dock">
                              <i></i>
                            </label>
                            Dock aside
                          </div>
                          <div>
                            <label class="i-switch bg-info pull-right">
                              <input type="checkbox" v-model="editData.layout" class="ng-pristine ng-untouched ng-valid" value="container">
                              <i></i>
                            </label>
                            Boxed layout
                          </div>
                        </div>
                      </div>
                      <div class="col col-lg-2-4 col-sm-6">
                        <div class="wrapper b-t b-light bg-light lter r-b ng-scope">
                          <div class="row row-sm">
                            <div class="col-xs-6 col-sm-6">
                              <label class="i-checks block m-b-sm" ng-click="
          app.settings.navbarHeaderColor='bg-black';
          app.settings.navbarCollapseColor='bg-white-only';
          app.settings.asideColor='bg-black';
         " role="button" tabindex="0">
                                <input type="radio" name="b" ng-model="app.settings.themeID" v-model="editData.theme"
                                       value="bg-black|bg-white|bg-black black-black" class="ng-pristine ng-untouched ng-valid" aria-checked="true" tabindex="0" aria-invalid="false">
                                <span class="block bg-light clearfix pos-rlt">
            <span class="active pos-abt w-full h-full bg-black-opacity text-center">
              <i class="glyphicon glyphicon-ok text-white m-t-xs"></i>
            </span>
            <b class="bg-black header"></b>
            <b class="bg-white header"></b>
            <b class="bg-black black-black"></b>
          </span>
                              </label>

                              <label class="i-checks block m-b-sm" ng-click="
          app.settings.navbarHeaderColor='bg-dark';
          app.settings.navbarCollapseColor='bg-white-only';
          app.settings.asideColor='bg-dark';
         " role="button" tabindex="0">
                                <input type="radio" name="b" ng-model="app.settings.themeID" v-model="editData.theme"
                                       value="bg-dark|bg-white|bg-dark dark-dark" class="ng-pristine ng-untouched ng-valid" aria-checked="false" tabindex="-1" aria-invalid="false">
                                <span class="block bg-light clearfix pos-rlt">
            <span class="active pos-abt w-full h-full bg-black-opacity text-center">
              <i class="glyphicon glyphicon-ok text-white m-t-xs"></i>
            </span>
            <b class="bg-dark header"></b>
            <b class="bg-white header"></b>
            <b class="bg-dark dark-dark"></b>
          </span>
                              </label>

                              <label class="i-checks block m-b-sm" ng-click="
          app.settings.navbarHeaderColor='bg-white-only';
          app.settings.navbarCollapseColor='bg-white-only';
          app.settings.asideColor='bg-black';
         " role="button" tabindex="0">
                                <input type="radio" ng-model="app.settings.themeID" v-model="editData.theme"
                                       value="bg-white|bg-white|bg-black black-white" class="ng-pristine ng-untouched ng-valid" name="b" aria-checked="false" tabindex="-1" aria-invalid="false">
                                <span class="block bg-light clearfix pos-rlt">
            <span class="active pos-abt w-full h-full bg-black-opacity text-center">
              <i class="glyphicon glyphicon-ok text-white m-t-xs"></i>
            </span>
            <b class="bg-white header"></b>
            <b class="bg-white header"></b>
            <b class="bg-black black-white"></b>
          </span>
                              </label>

                              <label class="i-checks block m-b-sm" ng-click="
          app.settings.navbarHeaderColor='bg-primary';
          app.settings.navbarCollapseColor='bg-white-only';
          app.settings.asideColor='bg-dark';
         " role="button" tabindex="0">
                                <input type="radio" ng-model="app.settings.themeID" v-model="editData.theme"
                                       value="bg-primary|bg-white|bg-dark dark-primary" class="ng-pristine ng-untouched ng-valid" name="b" aria-checked="false" tabindex="-1" aria-invalid="false">
                                <span class="block bg-light clearfix pos-rlt">
            <span class="active pos-abt w-full h-full bg-black-opacity text-center">
              <i class="glyphicon glyphicon-ok text-white m-t-xs"></i>
            </span>
            <b class="bg-primary header"></b>
            <b class="bg-white header"></b>
            <b class="bg-dark dark-primary"></b>
          </span>
                              </label>

                              <label class="i-checks block m-b-sm" ng-click="
          app.settings.navbarHeaderColor='bg-info';
          app.settings.navbarCollapseColor='bg-white-only';
          app.settings.asideColor='bg-black';
         " role="button" tabindex="0">
                                <input type="radio" ng-model="app.settings.themeID" v-model="editData.theme"
                                       value="bg-info|bg-white|bg-black black-info" class="ng-pristine ng-untouched ng-valid" name="b" aria-checked="false" tabindex="-1" aria-invalid="false">
                                <span class="block bg-light clearfix pos-rlt">
            <span class="active pos-abt w-full h-full bg-black-opacity text-center">
              <i class="glyphicon glyphicon-ok text-white m-t-xs"></i>
            </span>
            <b class="bg-info header"></b>
            <b class="bg-white header"></b>
            <b class="bg-black black-info"></b>
          </span>
                              </label>

                              <label class="i-checks block m-b-sm" ng-click="
          app.settings.navbarHeaderColor='bg-success';
          app.settings.navbarCollapseColor='bg-white-only';
          app.settings.asideColor='bg-dark';
         " role="button" tabindex="0">
                                <input type="radio" ng-model="app.settings.themeID" v-model="editData.theme"
                                       value="bg-success|bg-white|bg-dark dark-success" class="ng-pristine ng-untouched ng-valid" name="b" aria-checked="false" tabindex="-1" aria-invalid="false">
                                <span class="block bg-light clearfix pos-rlt">
            <span class="active pos-abt w-full h-full bg-black-opacity text-center">
              <i class="glyphicon glyphicon-ok text-white m-t-xs"></i>
            </span>
            <b class="bg-success header"></b>
            <b class="bg-white header"></b>
            <b class="bg-dark dark-success"></b>
          </span>
                              </label>

                              <label class="i-checks block">
                                <input type="radio" ng-model="app.settings.themeID" v-model="editData.theme"
                                       value="bg-danger|bg-white|bg-dark dark-danger" class="ng-pristine ng-untouched ng-valid" name="b" aria-checked="false" tabindex="-1" aria-invalid="false">
                                <span class="block bg-light clearfix pos-rlt">
            <span class="active pos-abt w-full h-full bg-black-opacity text-center">
              <i class="glyphicon glyphicon-ok text-white m-t-xs"></i>
            </span>
            <b class="bg-danger header"></b>
            <b class="bg-white header"></b>
            <b class="bg-dark dark-danger"></b>
          </span>
                              </label>
                              <label class="i-checks block">
                                <input type="radio" ng-model="app.settings.themeID" v-model="editData.theme"
                                       value="bg-danger|bg-white|bg-light light-danger" class="ng-pristine ng-untouched ng-valid" name="b" aria-checked="false" tabindex="-1" aria-invalid="false">
                                <span class="block bg-light clearfix pos-rlt">
            <span class="active pos-abt w-full h-full bg-black-opacity text-center">
              <i class="glyphicon glyphicon-ok text-white m-t-xs"></i>
            </span>
            <b class="bg-danger header"></b>
            <b class="bg-white header"></b>
            <b class="bg-light light-danger"></b>
          </span>
                              </label>
                            </div>
                            <div class="col-xs-6">
                              <label class="i-checks block m-b-sm" ng-click="
          app.settings.navbarHeaderColor='bg-black';
          app.settings.navbarCollapseColor='bg-black';
          app.settings.asideColor='bg-white b-r';
         " role="button" tabindex="0">
                                <input type="radio" ng-model="app.settings.themeID" v-model="editData.theme"
                                       value="bg-black|bg-black|bg-white white-black" class="ng-pristine ng-untouched ng-valid" name="b" aria-checked="false" tabindex="-1" aria-invalid="false">
                                <span class="block bg-light clearfix pos-rlt">
            <span class="active pos-abt w-full h-full bg-black-opacity text-center">
              <i class="glyphicon glyphicon-ok text-white m-t-xs"></i>
            </span>
            <b class="bg-black header"></b>
            <b class="bg-black header"></b>
            <b class="bg-white white-black"></b>
          </span>
                              </label>

                              <label class="i-checks block m-b-sm" ng-click="
          app.settings.navbarHeaderColor='bg-dark';
          app.settings.navbarCollapseColor='bg-dark';
          app.settings.asideColor='bg-light';
         " role="button" tabindex="0">
                                <input type="radio" name="b" ng-model="app.settings.themeID" v-model="editData.theme"
                                       value="bg-dark|bg-dark|bg-light light-dark" class="ng-pristine ng-untouched ng-valid" aria-checked="false" tabindex="-1" aria-invalid="false">
                                <span class="block bg-light clearfix pos-rlt">
            <span class="active pos-abt w-full h-full bg-black-opacity text-center">
              <i class="glyphicon glyphicon-ok text-white m-t-xs"></i>
            </span>
            <b class="bg-dark header"></b>
            <b class="bg-dark header"></b>
            <b class="bg-light light-dark"></b>
          </span>
                              </label>

                              <label class="i-checks block m-b-sm" ng-click="
          app.settings.navbarHeaderColor='bg-info dker';
          app.settings.navbarCollapseColor='bg-info dker';
          app.settings.asideColor='bg-light dker b-r';
         " role="button" tabindex="0">
                                <input type="radio" ng-model="app.settings.themeID" v-model="editData.theme"
                                       value="bg-info|bg-info|bg-light light-info" class="ng-pristine ng-untouched ng-valid" name="b" aria-checked="false" tabindex="-1" aria-invalid="false">
                                <span class="block bg-light clearfix pos-rlt">
            <span class="active pos-abt w-full h-full bg-black-opacity text-center">
              <i class="glyphicon glyphicon-ok text-white m-t-xs"></i>
            </span>
            <b class="bg-info header"></b>
            <b class="bg-info header"></b>
            <b class="bg-light light-info"></b>
          </span>
                              </label>

                              <label class="i-checks block m-b-sm" ng-click="
          app.settings.navbarHeaderColor='bg-primary';
          app.settings.navbarCollapseColor='bg-primary';
          app.settings.asideColor='bg-dark';
         " role="button" tabindex="0">
                                <input type="radio" ng-model="app.settings.themeID" v-model="editData.theme"
                                       value="bg-primary|bg-primary|bg-dark dark-primary" class="ng-pristine ng-untouched ng-valid" name="b" aria-checked="false" tabindex="-1" aria-invalid="false">
                                <span class="block bg-light clearfix pos-rlt">
            <span class="active pos-abt w-full h-full bg-black-opacity text-center">
              <i class="glyphicon glyphicon-ok text-white m-t-xs"></i>
            </span>
            <b class="bg-primary header"></b>
            <b class="bg-primary header"></b>
            <b class="bg-dark dark-primary"></b>
          </span>
                              </label>

                              <label class="i-checks block m-b-sm" ng-click="
          app.settings.navbarHeaderColor='bg-info dker';
          app.settings.navbarCollapseColor='bg-info dk';
          app.settings.asideColor='bg-black';
         " role="button" tabindex="0">
                                <input type="radio" ng-model="app.settings.themeID" v-model="editData.theme"
                                       value="bg-info|bg-info|bg-black black-info" class="ng-pristine ng-untouched ng-valid" name="b" aria-checked="false" tabindex="-1" aria-invalid="false">
                                <span class="block bg-light clearfix pos-rlt">
            <span class="active pos-abt w-full h-full bg-black-opacity text-center">
              <i class="glyphicon glyphicon-ok text-white m-t-xs"></i>
            </span>
            <b class="bg-info header"></b>
            <b class="bg-info header"></b>
            <b class="bg-black black-info"></b>
          </span>
                              </label>

                              <label class="i-checks block m-b-sm" ng-click="
          app.settings.navbarHeaderColor='bg-success';
          app.settings.navbarCollapseColor='bg-success';
          app.settings.asideColor='bg-dark';
          " role="button" tabindex="0">
                                <input type="radio" ng-model="app.settings.themeID" v-model="editData.theme"
                                       value="bg-success|bg-success|bg-dark dark-success" class="ng-pristine ng-untouched ng-valid" name="b" aria-checked="false" tabindex="-1" aria-invalid="false">
                                <span class="block bg-light clearfix pos-rlt">
            <span class="active pos-abt w-full h-full bg-black-opacity text-center">
              <i class="glyphicon glyphicon-ok text-white m-t-xs"></i>
            </span>
            <b class="bg-success header"></b>
            <b class="bg-success header"></b>
            <b class="bg-dark dark-success"></b>
          </span>
                              </label>

                              <label class="i-checks block" >
                                <input type="radio"  v-model="editData.theme"
                                       value="bg-danger|bg-danger|bg-dark dark-danger" class="ng-pristine ng-untouched ng-valid" name="b" aria-checked="false" tabindex="-1" aria-invalid="false">
                                <span class="block bg-light clearfix pos-rlt">
            <span class="active pos-abt w-full h-full bg-black-opacity text-center">
              <i class="glyphicon glyphicon-ok text-white m-t-xs"></i>
            </span>
            <b class="bg-danger header"></b>
            <b class="bg-danger header"></b>
            <b class="bg-dark dark-danger"></b>
          </span>
                              </label>

                              <label class="i-checks block" >
                                <input type="radio"  v-model="editData.theme"
                                       value="bg-danger|bg-danger|bg-light light-danger" class="ng-pristine ng-untouched ng-valid" name="b" aria-checked="false" tabindex="-1" aria-invalid="false">
                                <span class="block bg-light clearfix pos-rlt">
            <span class="active pos-abt w-full h-full bg-black-opacity text-center">
              <i class="glyphicon glyphicon-ok text-white m-t-xs"></i>
            </span>
            <b class="bg-danger header"></b>
            <b class="bg-danger header"></b>
            <b class="bg-light light-danger"></b>
          </span>
                              </label>
                              <label class="i-checks block" >
                                <input type="radio"  v-model="editData.theme"
                                       value="bg-success|bg-success|bg-light light-success" class="ng-pristine ng-untouched ng-valid" name="b" aria-checked="false" tabindex="-1" aria-invalid="false">
                                <span class="block bg-light clearfix pos-rlt">
            <span class="active pos-abt w-full h-full bg-black-opacity text-center">
              <i class="glyphicon glyphicon-ok text-white m-t-xs"></i>
            </span>
            <b class="bg-success header"></b>
            <b class="bg-success header"></b>
            <b class="bg-light light-success"></b>
          </span>
                              </label>
                            </div>
                          </div>
                        </div>
                      </div>
                    </div>
                  </div>
                </div>
              </div>
            </form>
          </div>
        </div>
      </div>
      <div slot="footer">
        <a class="btn btn-sm btn-danger" @click="editPost">提交</a>
        <a class="btn btn-sm btn-primary" @click="cancel">取消</a>
      </div>
    </bootstrap-modal>
    <bootstrap-modal ref="msgModal" :need-header="true" :need-footer="true">
      <div slot="title">
        提示
      </div>
      <div slot="body">
        不能删除该系统
      </div>
      <div slot="footer">
        <button class="btn btn-sm btn-primary" @click="cancel">确定</button>
      </div>
    </bootstrap-modal>
    <bootstrap-modal ref="msg2Modal" :need-header="true" :need-footer="true">
      <div slot="title">
        提示
      </div>
      <div slot="body">
        确定修改当前状态？
      </div>
      <div slot="footer">
        <button class="btn btn-sm btn-danger" @click="enableOk">确定</button>
        <button class="btn btn-sm btn-primary" @click="cancel">取消</button>
      </div>
    </bootstrap-modal>
    <add-modal ref="addModal" v-on:refresh-data="querySearch"></add-modal>
  </div>

  </div>
</template>

<script>
  import pager from "vue-simple-pager"
  import PullTo from 'vue-pull-to'
  import AddModal from './addModal.vue'
export default {
  components: {
    "add-modal": AddModal,
    "bootstrap-modal": require("vue2-bootstrap-modal"),
    pager,
    PullTo
  },
  data() {
    return {
      options: {
        // 可通过 https://github.com/simple-uploader/Uploader/tree/develop/samples/Node.js 示例启动服务
        target: process.env.service.apiHost+'/img/upload',   //上传地址
        testChunks: false,
        withCredentials:true,   //携带jwt
        singleFile:true,        //单文件上传
      },
      attrs: {
        accept: 'image/*'
      },
      errorImg: '',
      statusText: {
        success: '上传成功',
        error: '出错了',
        uploading: '上传中',
        paused: '暂停中',
        waiting: '等待中'
      },
      datalist: null,
      pageSizeList: [5, 10, 20, 50], //可选显示数据条数
      datacount: 0,
      sysid: 0,
      id: null,
      sysname: "",
      selected: "",
      addData: {
        name: "",
        addr: "",
        time_out: "",
        logo: "",
        theme: "",
        style: [],
        ident: "",
        secret:"",
        wechat_status:"1",
        callbackurl:""
      },
      editData: {},
      enableData: { id: null, status: null },
      pi: 1,
      ps:10,
      totalPage: 0,
      pageAuth: [],
      auth:{addsys:false}
    };
  },
  props:["path"],
  mounted() {
    this.$refs.main.style.height = document.documentElement.clientHeight + 'px';
    this.query()
  },
  computed: {
    editLayout() {
      return this.editData.layout.split(" ");
    }
  },
  methods: {
    //上传成功的事件
    fileSuccess (rootFile, file, message, chunk) {
      let data =JSON.parse(message);
      this.addData.logo=data.url;
        this.$notify({
          title: '成功',
          message: '上传成功',
          type: 'success',
          offset: 50,
          duration:2000,
        });
    },
    fileEditSuccess (rootFile, file, message, chunk) {
      let data =JSON.parse(message);
      this.editData.logo=data.url
    },
    //上传失败事件
    fileError(rootFile, file, message, chunk){
        this.$notify({
          title: '错误',
          message: '上传失败，请稍后再试',
          type: 'error',
          offset: 50,
          duration:2000,
        });
    },
    Add() {
            this.$refs.addModal.setModal()
        },
    next(){
      let pi =this.pi;
      this.pi = pi + 1;
      this.$http.get("/sys/manage", { pi: this.pi ,ps:this.ps,name: this.sysname,status: this.selected})
        .then(res => {
          if(res.list.length <= 0){
            this.pi=pi;
            return false
          }
          this.datalist = this.datalist.concat(res.list);
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
    handleSizeChange(val){
      this.ps =val;
      this.query()
    },
    goPage(val) {
      this.pi = val;
      this.query()
    },

    //点查询事件
    querySearch() {
      this.pi = 1;
      this.query();
    },

    query() {
      this.$http.get("/sys/manage", {
        pi: this.pi,
        ps:this.ps,
        name: this.sysname,
        status: this.selected
      })
        .then(res => {
          this.datalist = res.list;
          this.datacount = res.count;
          this.totalPage = Math.ceil(res.count / 10);
        })
        .catch(err => {
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
        });
    },
    openAdd() {
      this.$refs.addModal.open();
    },
    addNew() {
      let str = "";
      this.addData.style.forEach((item, index) => {
        str += item + " ";
      });
      this.$http.post("/sys/manage", {
        name: this.addData.name,
        addr: this.addData.addr,
        time_out: this.addData.time_out,
        logo: this.addData.logo,
        style: str,
        theme: this.addData.theme,
        ident: this.addData.ident,
        secret:this.addData.secret,
        wechat_status: this.addData.wechat_status,
        callbackurl: this.addData.callbackurl
      })
        .then(res => {
          this.$refs.addModal.close();
          this.goPage({ page: this.pi });
            this.$notify({
              title: '成功',
              message: '添加成功',
              type: 'success',
              offset: 50,
              duration:2000,
            });
        })
        .catch(err => {
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
        });
    },
    deleteById(id) {
      this.sysid = id;
      if (this.sysid == 0 || this.sysid == 100) {
        this.$refs.msgModal.open();
        return false;
      }
      this.$refs.theModal.open();
    },
    cancel() {
      this.$refs.theModal.close();
      this.$refs.msgModal.close();
      this.$refs.editModal.close();
      this.$refs.msg2Modal.close();
    },
    ok() {
      this.$http.del("/sys/manage", { data: { id: this.sysid } })
        .then(res => {
          this.cancel();
          this.querySearch();
            this.$notify({
              title: '成功',
              message: '删除成功',
              type: 'success',
              offset: 50,
              duration:2000
            });
        })
        .catch(err => {
            console.log(err);
            this.$notify({
              title: '错误',
              message: '网络错误,请稍后再试',
              type: 'error',
              offset: 50,
              duration:2000,
            });
        });
    },
    enable(id, status) {

      this.enableData.id = id;
      this.enableData.status = status;

      this.$confirm("确定执行此操作?, 是否继续?", "提示", {
        confirmButtonText: "确定",
        cancelButtonText: "取消",
        type: "warning"
      }).then(() => {
        this.$http.put("/sys/manage", {
          id: this.enableData.id,
          status: this.enableData.status
        })
        .then(res => {
          this.goPage(this.pi);
          this.$notify({
            title: '成功',
            message: '状态修改成功',
            type: 'success',
            offset: 50,
            duration:2000,
          });
        })
        .catch(err => {
          console.log(err);
          this.$notify({
            title: '错误',
            message: '网络错误,请稍后再试',
            type: 'error',
            offset: 50,
            duration:2000,
          });

        });

      }).catch(() => {
        this.$message({
          type: "info",
          message: "已取消删除"
        });
      });


    },
    enableOk() {

    },
    edit(id) {
      this.datalist.forEach((item, index) => {
        if (item.id == id) {
          let a = typeof item.layout;
          if (a == "string") {
            item.layout = item.layout.split(" ");
          }
          console.log(item);
          this.editData = item;
          this.$refs.editModal.open();
        }
      });
    },
    editPost() {
      let str = "";
      let edit = this.editData;
      this.editData.layout.forEach((item, index) => {
        str += item + " ";
      });
      edit.layout = str.trim();

      var msg = this.checkBeforeSave(edit);
      if (msg) {
        this.$notify({
              title: '错误',
              message: msg,
              type: 'error',
              offset: 50,
              duration:2000
            });
        return;
      }
      this.$http.post("/sys/manage/edit", edit)
        .then(res => {
          this.$refs.editModal.close();
          this.goPage({ page: this.pi });
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
              message: '登录超时，请重新登录',
              type: 'error',
              offset: 50,
              duration:2000,
              onClose: ()=> {
                this.$router.push("/member/login");
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
    manage(id) {
      this.$router.push({
        name: "sysfunc",
        query: {
          id: id
        }
      });
    },
    checkBeforeSave(editData) {
      if (!editData.name) {
        return "系统名称不能为空";
      }
      if(!editData.ident) {
        return "系统英文名称不能为空";
      }
      if(!editData.secret) {
        return "secret不能为空";
      }
      if (!editData.logo) {
        return "logo图片必须上传";
      }
      if(!editData.theme) {
        return "请选择主题样式";
      }
      if(!editData.layout) {
        return "请选择页面布局样式";
      }
      return ""
    }

  }
};
</script>

<style>
  .uploader-example {
    font-size: 12px;
  }
  .uploader-example .uploader-btn {
    margin-right: 4px;
  }
  .uploader-example .uploader-list {
    max-height: 440px;
    overflow: auto;
    overflow-x: hidden;
    overflow-y: auto;
  }
  .page-pagination{padding: 10px 15px;text-align: right;}
</style>

