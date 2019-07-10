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
            <a class="visible-xs-inline visible-sm-inline visible-md-inline  visible-lg-inline btn btn btn-success" @click="query">查询</a>
            <span ng-controller="ModalDemoCtrl">
              <script type="text/ng-template" id="myModalContent.html">
                <div ng-include="'src/pages/user/index/add.vue'"></div>
              </script>
              <a class="visible-sm-inline visible-md-inline  visible-lg-inline btn btn-primary" ref="addsys" @click="openAdd">添加</a>
            </span>
          </form>
        </div>
      </div>
      <div class="table-responsive">

         <table class="table table-striped m-b-none">
        <thead>
          <tr>
            <th>编号</th>
            <th>系统名称</th>
            <th class="visible-md  visible-lg">首页地址</th>
            <th>状态</th>
            <th class="visible-md  visible-lg">超时时长</th>
            <th class="visible-md  visible-lg">logo</th>
            <th>操作</th>
          </tr>
        </thead>
        <tbody>
          <tr v-for="(item, index) in datalist" :key="index">
            <td>{{item.id}}</td>
            <td>{{item.name}}</td>
            <td class="visible-md  visible-lg">{{item.index_url}}</td>
            <td v-if="item.enable==0" class="text-danger">禁用</td>
            <td v-if="item.enable==1" class="text-success">启用</td>
            <td class="visible-md  visible-lg">{{item.login_timeout}}</td>
            <td class="visible-md  visible-lg ">
              <img v-if="item.theme" :class="item.theme.split('|')[0]"
                   :src="item.logo" :onerror="errorImg" alt="">
            </td>
            <td>
              <div class="form-group form-inline">
              <div class="form-group">
                <button class="btn btn-xs btn-primary visible-md visible-lg" @click="edit(item.id)">编辑</button>
              </div>
              <div class="form-group" >
                <a class="btn btn-xs btn-warning" @click="enable(item.id,1)" v-if="item.enable==0" >启用</a>
                <a class="btn btn-xs btn-warning" @click="enable(item.id,0)" v-if="item.enable==1" >禁用</a>
              </div>
              <div class="form-group">
                <a class="btn btn-xs btn-danger visible-md visible-lg" @click="deleteById(item.id)">删除</a>

              </div>
              <div class="form-group">
                <a class="btn btn-xs btn-default visible-md visible-lg" @click="manage(item.id)">管理</a>
              </div>
              </div>
            </td>
          </tr>
        </tbody>
      </table>
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
    <bootstrap-modal ref="addModal" :need-header="true" :need-footer="true">
      <div slot="title">
        添加系统
      </div>
      <div slot="body">
        <div class="panel panel-default">
          <div class="panel-body">
            <form role="form" class="ng-pristine ng-valid ng-submitted height-min">
              <div class="form-group">
                <label>系统名称</label>
                <input class="form-control" placeholder="请输入系统名称" v-validate="'required'" name="name" v-model="addData.name"  type="text">
                <div class="form-height text-danger"><span v-show="errors.first('name')">系统名称不能为空</span></div>
              </div>
              <div class="form-group">
                <label>系统英文名称</label>
                <input class="form-control" placeholder="请输入系统英文名称" v-validate="'required|alpha'" name="name-e" v-model="addData.ident"  type="text">
                <div class="form-height text-danger"><span v-show="errors.first('name-e')">系统英文名称不能为空且为字母</span></div>
              </div>
              <div class="form-group">
                <label>首页地址</label>
                <input class="form-control" placeholder="请输入首页地址" v-validate="'required'" name="addr" v-model="addData.addr"  type="text">
                <div class="form-height text-danger"><span v-show="errors.first('addr')">首页地址不能为空</span></div>
              </div>
              <div class="form-group">
                <label>超时时常</label>
                <input class="form-control" placeholder="请输入超时时常" v-validate="'required|numeric'" v-model="addData.time_out" name="time_out"  type="text">
                <div class="form-height text-danger"><span v-show="errors.first('time_out')">超时时常不能为空且必须为数字</span> </div>
              </div>
              <div class="form-group">
                <label>微信登录</label>
                <div class="radio">
                  <label class="i-checks">
                    <input type="radio" name="toasts" v-model="addData.wechat_status"  value="1" class="ng-pristine ng-untouched ng-valid">
                    <i></i>
                    启用
                  </label>
                  <label class="i-checks">
                    <input type="radio" name="toasts" v-model="addData.wechat_status" value="0" class="ng-pristine ng-untouched ng-valid">
                    <i></i>
                    禁用
                  </label>
                </div>
              </div>
              <div class="form-group">
                <!--<label>logo</label>-->
                <input class="form-control" placeholder="请输入logo地址" v-validate="'required'" name="logo" v-model="addData.logo"  type="hidden">
                <!--<div class="form-height text-danger"> <span v-show="errors.first('logo')">logo地址不能为空</span> </div>-->
                <uploader :options="options" class="uploader-example" :file-status-text="statusText"   ref="uploader" @file-success="fileSuccess" @file-error="fileError">
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
                      <div class="col col-lg-2-4">
                        <div class="panel-body ng-scope">
                          <div class="m-b-sm">
                            <label class="i-switch bg-info pull-right">
                              <input type="checkbox" v-model="addData.style" class="ng-pristine ng-untouched ng-valid"  value="app-header-fixed">
                              <i></i>
                            </label>
                            Fixed header
                          </div>
                          <div class="m-b-sm">
                            <label class="i-switch bg-info pull-right">
                              <input type="checkbox" v-model="addData.style" class="ng-pristine ng-untouched ng-valid"  value="app-aside-fixed">
                              <i></i>
                            </label>
                            Fixed aside
                          </div>
                          <div class="m-b-sm">
                            <label class="i-switch bg-info pull-right">
                              <input type="checkbox" v-model="addData.style" class="ng-pristine ng-untouched ng-valid"  value="app-aside-folded">
                              <i></i>
                            </label>
                            Folded aside
                          </div>
                          <div class="m-b-sm">
                            <label class="i-switch bg-info pull-right">
                              <input type="checkbox" v-model="addData.style" class="ng-pristine ng-untouched ng-valid"  value="app-aside-dock">
                              <i></i>
                            </label>
                            Dock aside
                          </div>
                          <div>
                            <label class="i-switch bg-info pull-right">
                              <input type="checkbox" v-model="addData.style" class="ng-pristine ng-untouched ng-valid"  value="container">
                              <i></i>
                            </label>
                            Boxed layout
                          </div>
                        </div>
                      </div>
                      <div class="col col-lg-2-4">
                        <div class="wrapper b-t b-light bg-light lter r-b ng-scope">
                          <div class="row row-sm">
                            <div class="col-xs-6">
                              <label class="i-checks block m-b-sm" ng-click="
          app.settings.navbarHeaderColor='bg-black';
          app.settings.navbarCollapseColor='bg-white-only';
          app.settings.asideColor='bg-black';
         " role="button" tabindex="0">
                                <input type="radio" name="a" ng-model="app.settings.themeID" v-model="addData.theme"
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
                                <input type="radio" name="a" ng-model="app.settings.themeID" v-model="addData.theme"
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
                                <input type="radio" ng-model="app.settings.themeID" v-model="addData.theme"
                                       value="bg-white|bg-white|bg-black black-white" class="ng-pristine ng-untouched ng-valid" name="a" aria-checked="false" tabindex="-1" aria-invalid="false">
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
                                <input type="radio" ng-model="app.settings.themeID" v-model="addData.theme"
                                       value="bg-primary|bg-white|bg-dark dark-primary" class="ng-pristine ng-untouched ng-valid" name="a" aria-checked="false" tabindex="-1" aria-invalid="false">
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
                                <input type="radio" ng-model="app.settings.themeID" v-model="addData.theme"
                                       value="bg-info|bg-white|bg-black black-info" class="ng-pristine ng-untouched ng-valid" name="a" aria-checked="false" tabindex="-1" aria-invalid="false">
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
                                <input type="radio" ng-model="app.settings.themeID" v-model="addData.theme"
                                       value="bg-success|bg-white|bg-dark dark-success" class="ng-pristine ng-untouched ng-valid" name="a" aria-checked="false" tabindex="-1" aria-invalid="false">
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
                                <input type="radio" v-model="addData.theme"
                                       value="bg-danger|bg-white|bg-dark dark-danger" class="ng-pristine ng-untouched ng-valid" name="a" aria-checked="false" tabindex="-1" aria-invalid="false">
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
                                <input type="radio" v-model="addData.theme"
                                       value="bg-danger|bg-white|bg-light light-danger" class="ng-pristine ng-untouched ng-valid" name="a" aria-checked="false" tabindex="-1" aria-invalid="false">
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
                                <input type="radio" ng-model="app.settings.themeID" v-model="addData.theme"
                                       value="bg-black|bg-black|bg-white white-black" class="ng-pristine ng-untouched ng-valid" name="a" aria-checked="false" tabindex="-1" aria-invalid="false">
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
                                <input type="radio" name="a" ng-model="app.settings.themeID" v-model="addData.theme"
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
                                <input type="radio" ng-model="app.settings.themeID" v-model="addData.theme"
                                       value="bg-info|bg-info|bg-light light-info" class="ng-pristine ng-untouched ng-valid" name="a" aria-checked="false" tabindex="-1" aria-invalid="false">
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
                                <input type="radio" ng-model="app.settings.themeID" v-model="addData.theme"
                                       value="bg-primary|bg-primary|bg-dark dark-primary" class="ng-pristine ng-untouched ng-valid" name="a" aria-checked="false" tabindex="-1" aria-invalid="false">
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
                                <input type="radio" ng-model="app.settings.themeID" v-model="addData.theme"
                                       value="bg-info|bg-info|bg-black black-info" class="ng-pristine ng-untouched ng-valid" name="a" aria-checked="false" tabindex="-1" aria-invalid="false">
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
                                <input type="radio" ng-model="app.settings.themeID" v-model="addData.theme"
                                       value="bg-success|bg-success|bg-dark dark-success" class="ng-pristine ng-untouched ng-valid" name="a" aria-checked="false" tabindex="-1" aria-invalid="false">
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
                                <input type="radio" ng-model="app.settings.themeID" v-model="addData.theme"
                                       value="bg-danger|bg-danger|bg-dark dark-danger" class="ng-pristine ng-untouched ng-valid" name="a" aria-checked="false" tabindex="-1" aria-invalid="false">
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
                                <input type="radio" v-model="addData.theme"
                                       value="bg-danger|bg-danger|bg-light light-danger" class="ng-pristine ng-untouched ng-valid" name="a" aria-checked="false" tabindex="-1" aria-invalid="false">
                                <span class="block bg-light clearfix pos-rlt">
            <span class="active pos-abt w-full h-full bg-black-opacity text-center">
              <i class="glyphicon glyphicon-ok text-white m-t-xs"></i>
            </span>
            <b class="bg-danger header"></b>
            <b class="bg-danger header"></b>
            <b class="bg-light light-danger"></b>
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
        <a  class="btn btn-sm btn-danger" @click="addNew">提交</a>
        <a  class="btn btn-sm btn-primary" @click="cancel">取消</a>
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
              <div class="form-group">
                <label>系统名称</label>
                <input class="form-control" placeholder="" v-validate="'required'" name="name2" v-model="editData.name"  type="text">
                <div class="form-height text-danger"><span v-show="errors.first('name2')">系统名称不能为空</span></div>
                <input class="form-control" placeholder=""  v-model="editData.id"  type="hidden">
              </div>
              <div class="form-group">
                <label>系统英文名称</label>
                <input class="form-control" placeholder="" v-validate="'required|alpha'" name="name-e" v-model="editData.ident"  type="text">
                <div class="form-height text-danger"><span v-show="errors.first('name-e')">系统名称不能为空</span></div>
                <input class="form-control" placeholder=""  v-model="editData.id"  type="hidden">
              </div>
              <div class="form-group">
                <label>首页地址</label>
                <input class="form-control" placeholder="" v-validate="'required'" name="url2" v-model="editData.index_url"  type="text">
                <div class="form-height text-danger"><span v-show="errors.first('url2')">首页地址不能为空</span></div>
              </div>
              <div class="form-group">
                <label>超时时常</label>
                <input class="form-control" placeholder="" v-validate="'required|numeric'" name="time_out2" v-model="editData.login_timeout"  type="text">
                <div class="form-height text-danger"><span v-show="errors.first('time_out2')">超时时常不能为空且必须为数字</span></div>
              </div>
              <div class="form-group">
                <label>微信登录</label>
                <div class="radio">
                  <label class="i-checks">
                    <input type="radio" name="toasts" v-model="editData.wechat_status"  value="1" class="ng-pristine ng-untouched ng-valid">
                    <i></i>
                    启用
                  </label>
                  <label class="i-checks">
                    <input type="radio" name="toasts" v-model="editData.wechat_status" value="0" class="ng-pristine ng-untouched ng-valid">
                    <i></i>
                    禁用
                  </label>
                </div>
              </div>
              <div class="form-group">
                <label>{{editData.logo}}</label>
                <input class="form-control" placeholder="" v-validate="'required'" name="logo2" v-model="editData.logo"  type="hidden">
                <!--<div class="form-height text-danger"><span v-show="errors.first('logo2')">logo地址不能为空</span></div>-->
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
      options: {
        // 可通过 https://github.com/simple-uploader/Uploader/tree/develop/samples/Node.js 示例启动服务
        target: process.env.service.url+'/sso/img/upload',   //上传地址
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
        ident: ""
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
    this.goPage({ page: 1,ps:this.ps });
    //this.queryTags(this.path);
  },
  computed: {
    editLayout() {
      return this.editData.layout.split(" ");
    }
  },
  methods: {
    // queryTags(path){
    //   if (path == '') return;
    //   this.$fetch("/sso/role/auth",{sys_id:0,role_id:0,path:path})
    //     .then(res => {
    //       this.pageAuth = res;
    //       if (this.pageAuth.length <= 0){
    //         return false
    //       }
    //       this.pageAuth.forEach((item,index)=>{
    //         let path = item.path;
    //        this.auth[path] = true
    //
    //       })
    //     })
    //     .catch(err => {
    //       this.$notify({
    //         title: '错误',
    //         message: '网络错误,请稍后再试',
    //         type: 'error',
    //         offset: 50,
    //         duration:2000,
    //       });
    //     });
    // },
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
    next(){
      let pi =this.pi;
      this.pi = pi + 1;
      this.$fetch("/sso/sys/manage", { pi: this.pi ,ps:this.ps,name: this.sysname,status: this.selected})
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
    goPage(data) {
      this.pi = data.page;
      this.$fetch("/sso/sys/manage", { pi: data.page ,ps:this.ps,name: this.sysname,status: this.selected})
        .then(res => {
          this.datalist = res.list;
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
    query() {
      this.$fetch("/sso/sys/manage", {
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
      this.$post("/sso/sys/manage", {
        name: this.addData.name,
        addr: this.addData.addr,
        time_out: this.addData.time_out,
        logo: this.addData.logo,
        style: str,
        theme: this.addData.theme,
        ident: this.addData.ident,
        wechat_status: this.addData.wechat_status,
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
      this.$refs.addModal.close();
      this.$refs.msg2Modal.close();
    },
    ok() {
      this.$del("/sso/sys/manage", { data: { id: this.sysid } })
        .then(res => {
          this.cancel();
          this.goPage({ page: this.pi });
            this.$notify({
              title: '成功',
              message: '删除成功',
              type: 'success',
              offset: 50,
              duration:2000
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
    enable(id, status) {
      this.$refs.msg2Modal.open();
      this.enableData.id = id;
      this.enableData.status = status;
    },
    enableOk() {
      this.$put("/sso/sys/manage", {
        id: this.enableData.id,
        status: this.enableData.status
      })
        .then(res => {
          this.goPage({ page: this.pi });
          this.cancel();
            this.$notify({
              title: '成功',
              message: '状态修改成功',
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
    edit(id) {
      this.datalist.forEach((item, index) => {
        if (item.id == id) {
          let a = typeof item.layout;
          if (a == "string") {
            item.layout = item.layout.split(" ");
          }
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
      this.$post("/sso/sys/manage/edit", edit)
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
  .list-number{display: inline-block;padding: 15px 0 3px 15px;}
  .list-page{display: inline-block;position: absolute;right: 15px;margin-top: -43px;}
</style>
