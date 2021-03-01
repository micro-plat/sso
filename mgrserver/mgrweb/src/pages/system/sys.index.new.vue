//系统管理,查询，添加，编辑，禁用
<template>
  <div ref="main">
    <div class="panel panel-default">
      <div class="panel panel-default">
        <div class="panel-body">
          <form class="form-inline">
            <div class="form-group">
              <label class="sr-only">系统名</label>
              <input
                type="text"
                class="form-control"
                onkeypress="if(event.keyCode == 13) return false;"
                v-model="sysname"
                placeholder="请输入系统名称"
              />
            </div>
            <div class="form-group">
              <label class="col-sm-2 control-label sr-only">状态</label>
              <select class="form-control visible-md visible-lg" v-model="selected">
                <option value selected="selected">---请选择状态---</option>
                <option value="0">禁用</option>
                <option value="1">启用</option>
              </select>
            </div>
            <a
              class="visible-xs-inline visible-sm-inline visible-md-inline visible-lg-inline btn btn btn-success"
              @click="querySearch"
            >查询</a>
            <span ng-controller="ModalDemoCtrl">
              <script type="text/ng-template" id="myModalContent.html">
  <div ng-include="'src/pages/user/index/add.vue'"></div>
              </script>
              <a
                class="visible-sm-inline visible-md-inline visible-lg-inline btn btn-primary"
                ref="addsys"
                @click="Add"
              >添加</a>
            </span>
          </form>
        </div>
      </div>
      <el-scrollbar style="height:100%">
        <el-table :data="datalist" stripe style="width: 100%">
          <el-table-column width="100" prop="ident" label="英文名称"></el-table-column>
          <el-table-column width="230" prop="name" label="系统名称"></el-table-column>
          <el-table-column width="80" prop="enable" label="状态">
            <template slot-scope="scope">
              <el-tag type="info" v-if="scope.row.enable == 0">禁用</el-tag>
              <el-tag type="success" v-if="scope.row.enable == 1">启用</el-tag>
            </template>
          </el-table-column>
          <el-table-column width="230" prop="logo" label="logo">
            <template slot-scope="scope">
              <img
                v-if="scope.row.theme"
                :class="scope.row.theme.split('|')[0]"
                :src="scope.row.logo"
                :onerror="errorImg"
                alt
              />
            </template>
          </el-table-column>
          <el-table-column width="300" prop="callbackurl" label="登录回调地址"></el-table-column>

          <el-table-column label="操作">
            <template slot-scope="scope">
              <el-button plain type="primary" size="mini" @click="edit(scope.row.id)">编辑</el-button>
              <el-button
                plain
                type="success"
                size="mini"
                @click="enable(scope.row.id,1)"
                v-if="scope.row.enable == 0"
              >启用</el-button>
              <el-button
                plain
                type="info"
                size="mini"
                @click="enable(scope.row.id,0)"
                v-if="scope.row.enable == 1"
              >禁用</el-button>
              <el-button plain type="primary" size="mini" @click="exportMenu(scope.row.id)">导出菜单</el-button>
              <el-button plain type="primary" size="mini" @click="importMenu(scope.row.id)">导入菜单</el-button>
              <el-button plain type="primary" size="mini" @click="setSecret(scope.row.id)">设置秘钥</el-button>
              <el-button plain type="danger" size="mini" @click="deleteById(scope.row.id)">删除</el-button>
              <el-button
                plain
                type="warning"
                size="mini"
                @click="manage(scope.row.id, scope.row.ident)"
              >菜单</el-button>
              <el-button
                plain
                type="warning"
                size="mini"
                @click="managePermission(scope.row.id, scope.row.ident)"
              >数据权限</el-button>
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
      <bootstrap-modal ref="secretModal" :need-header="true" :need-footer="true">
        <div slot="title">密钥设置</div>
        <div slot="body">
          <div class="panel panel-default">
            <div class="panel-body">
              <form role="form" class="ng-pristine ng-valid ng-submitted">
                <div style="display:none">
                  <input maxlength="32" v-model="secrectData.id" name="id" />
                </div>
                <div class="form-group">
                  <label>秘钥</label>
                  <input
                    maxlength="32"
                    class="form-control"
                    placeholder="请输入秘钥"
                    v-model="secrectData.secret"
                    name="secret"
                    type="text"
                  />
                </div>
              </form>
            </div>
          </div>
        </div>
        <div slot="footer">
          <a class="btn btn-sm btn-danger" @click="saveSecret">提交</a>
          <a class="btn btn-sm btn-primary" @click="secretCancel">取消</a>
        </div>
      </bootstrap-modal>
      <bootstrap-modal ref="importModal" :need-header="true" :need-footer="true">
        <div slot="title">导入菜单</div>
        <div slot="body">
          <div class="panel panel-default">
            <div class="panel-body">
              <form role="form" class="ng-pristine ng-valid ng-submitted">
                <div style="display:none">
                  <input v-model="importData.id" name="id" />
                </div>
                <div class="form-group">
                  <label>选择菜单文件</label>
                  <input type="file" id="upload_file" @change="fileChange" accept=".xlsx, .xls" />
                </div>
              </form>
            </div>
          </div>
        </div>
        <div slot="footer">
          <a class="btn btn-sm btn-danger" @click="saveImportMenu">提交</a>
          <a class="btn btn-sm btn-primary" @click="importClose">取消</a>
        </div>
      </bootstrap-modal>
      <bootstrap-modal ref="editModal" :need-header="true" :need-footer="true">
        <div slot="title">编辑系统</div>
        <div slot="body">
          <div class="panel panel-default">
            <div class="panel-body">
              <form role="form" class="ng-pristine ng-valid ng-submitted">
                <el-row :span="24">
                  <el-col :span="12">
                    <div class="form-group">
                      <label>系统名称</label>
                      <input
                        class="form-control"
                        placeholder
                        maxlength="30"
                        name="name2"
                        v-model="editData.name"
                        type="text"
                      />
                      <div class="form-height text-danger">
                        <span v-show="errors.first('name2')">系统名称不能为空</span>
                      </div>
                      <input class="form-control" placeholder v-model="editData.id" type="hidden" />
                    </div>
                  </el-col>
                  <el-col :span="12">
                    <div class="form-group" style="margin-left:10px;">
                      <label>系统英文名称</label>
                      <input
                        class="form-control"
                        placeholder
                        maxlength="30"
                        name="name-e"
                        v-model="editData.ident"
                        type="text"
                      />
                      <div class="form-height text-danger">
                        <span v-show="errors.first('name-e')">系统名称不能为空</span>
                      </div>
                      <input class="form-control" placeholder v-model="editData.id" type="hidden" />
                    </div>
                  </el-col>
                </el-row>
                <!-- <div class="form-group">
                <label>secret</label>
                <input class="form-control" placeholder="系统签名的secret" v-model="editData.secret" name="secret"  type="text">
                <div class="form-height text-danger"><span v-show="errors.first('secret')">secret不能为空</span> </div>
                </div>-->
                <div class="form-group">
                  <label>sso登录后回调子系统的地址(如:http://www.123.com/abc)</label>
                  <input
                    class="form-control"
                    placeholder="请输入回调地址"
                    v-model="editData.callbackurl"
                    name="callbackurl"
                    type="text"
                    maxlength="64"
                  />
                </div>
                <div class="form-group">
                  <label>{{editData.logo}}</label>
                  <input v-show="editData.logo != ''" type="button" value="删除" @click="deletePic" />
                  <input
                    class="form-control"
                    placeholder
                    name="logo2"
                    v-model="editData.logo"
                    type="hidden"
                  />
                  <div class="form-height text-danger">
                    <span v-show="errors.first('logo2')">logo地址不能为空</span>
                  </div>
                  <uploader
                    :options="options"
                    class="uploader-example"
                    :file-status-text="statusText"
                    ref="uploader"
                    @file-success="fileEditSuccess"
                    @file-error="fileError"
                  >
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
                                <input
                                  type="checkbox"
                                  v-model="editData.layout"
                                  class="ng-pristine ng-untouched ng-valid"
                                  value="app-header-fixed"
                                />
                                <i></i>
                              </label>
                              Fixed header
                            </div>
                            <div class="m-b-sm">
                              <label class="i-switch bg-info pull-right">
                                <input
                                  type="checkbox"
                                  v-model="editData.layout"
                                  class="ng-pristine ng-untouched ng-valid"
                                  value="app-aside-fixed"
                                />
                                <i></i>
                              </label>
                              Fixed aside
                            </div>
                            <div class="m-b-sm">
                              <label class="i-switch bg-info pull-right">
                                <input
                                  type="checkbox"
                                  v-model="editData.layout"
                                  class="ng-pristine ng-untouched ng-valid"
                                  value="app-aside-folded"
                                />
                                <i></i>
                              </label>
                              Folded aside
                            </div>
                            <div class="m-b-sm">
                              <label class="i-switch bg-info pull-right">
                                <input
                                  type="checkbox"
                                  v-model="editData.layout"
                                  class="ng-pristine ng-untouched ng-valid"
                                  value="app-aside-dock"
                                />
                                <i></i>
                              </label>
                              Dock aside
                            </div>
                            <div>
                              <label class="i-switch bg-info pull-right">
                                <input
                                  type="checkbox"
                                  v-model="editData.layout"
                                  class="ng-pristine ng-untouched ng-valid"
                                  value="container"
                                />
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
                                <label
                                  class="i-checks block m-b-sm"
                                  ng-click="
          app.settings.navbarHeaderColor='bg-black';
          app.settings.navbarCollapseColor='bg-white-only';
          app.settings.asideColor='bg-black';
         "
                                  role="button"
                                  tabindex="0"
                                >
                                  <input
                                    type="radio"
                                    name="b"
                                    ng-model="app.settings.themeID"
                                    v-model="editData.theme"
                                    value="bg-black|bg-white|bg-black black-black"
                                    class="ng-pristine ng-untouched ng-valid"
                                    aria-checked="true"
                                    tabindex="0"
                                    aria-invalid="false"
                                  />
                                  <span class="block bg-light clearfix pos-rlt">
                                    <span
                                      class="active pos-abt w-full h-full bg-black-opacity text-center"
                                    >
                                      <i class="glyphicon glyphicon-ok text-white m-t-xs"></i>
                                    </span>
                                    <b class="bg-black header"></b>
                                    <b class="bg-white header"></b>
                                    <b class="bg-black black-black"></b>
                                  </span>
                                </label>

                                <label
                                  class="i-checks block m-b-sm"
                                  ng-click="
          app.settings.navbarHeaderColor='bg-dark';
          app.settings.navbarCollapseColor='bg-white-only';
          app.settings.asideColor='bg-dark';
         "
                                  role="button"
                                  tabindex="0"
                                >
                                  <input
                                    type="radio"
                                    name="b"
                                    ng-model="app.settings.themeID"
                                    v-model="editData.theme"
                                    value="bg-dark|bg-white|bg-dark dark-dark"
                                    class="ng-pristine ng-untouched ng-valid"
                                    aria-checked="false"
                                    tabindex="-1"
                                    aria-invalid="false"
                                  />
                                  <span class="block bg-light clearfix pos-rlt">
                                    <span
                                      class="active pos-abt w-full h-full bg-black-opacity text-center"
                                    >
                                      <i class="glyphicon glyphicon-ok text-white m-t-xs"></i>
                                    </span>
                                    <b class="bg-dark header"></b>
                                    <b class="bg-white header"></b>
                                    <b class="bg-dark dark-dark"></b>
                                  </span>
                                </label>

                                <label
                                  class="i-checks block m-b-sm"
                                  ng-click="
          app.settings.navbarHeaderColor='bg-white-only';
          app.settings.navbarCollapseColor='bg-white-only';
          app.settings.asideColor='bg-black';
         "
                                  role="button"
                                  tabindex="0"
                                >
                                  <input
                                    type="radio"
                                    ng-model="app.settings.themeID"
                                    v-model="editData.theme"
                                    value="bg-white|bg-white|bg-black black-white"
                                    class="ng-pristine ng-untouched ng-valid"
                                    name="b"
                                    aria-checked="false"
                                    tabindex="-1"
                                    aria-invalid="false"
                                  />
                                  <span class="block bg-light clearfix pos-rlt">
                                    <span
                                      class="active pos-abt w-full h-full bg-black-opacity text-center"
                                    >
                                      <i class="glyphicon glyphicon-ok text-white m-t-xs"></i>
                                    </span>
                                    <b class="bg-white header"></b>
                                    <b class="bg-white header"></b>
                                    <b class="bg-black black-white"></b>
                                  </span>
                                </label>

                                <label
                                  class="i-checks block m-b-sm"
                                  ng-click="
          app.settings.navbarHeaderColor='bg-primary';
          app.settings.navbarCollapseColor='bg-white-only';
          app.settings.asideColor='bg-dark';
         "
                                  role="button"
                                  tabindex="0"
                                >
                                  <input
                                    type="radio"
                                    ng-model="app.settings.themeID"
                                    v-model="editData.theme"
                                    value="bg-primary|bg-white|bg-dark dark-primary"
                                    class="ng-pristine ng-untouched ng-valid"
                                    name="b"
                                    aria-checked="false"
                                    tabindex="-1"
                                    aria-invalid="false"
                                  />
                                  <span class="block bg-light clearfix pos-rlt">
                                    <span
                                      class="active pos-abt w-full h-full bg-black-opacity text-center"
                                    >
                                      <i class="glyphicon glyphicon-ok text-white m-t-xs"></i>
                                    </span>
                                    <b class="bg-primary header"></b>
                                    <b class="bg-white header"></b>
                                    <b class="bg-dark dark-primary"></b>
                                  </span>
                                </label>

                                <label
                                  class="i-checks block m-b-sm"
                                  ng-click="
          app.settings.navbarHeaderColor='bg-info';
          app.settings.navbarCollapseColor='bg-white-only';
          app.settings.asideColor='bg-black';
         "
                                  role="button"
                                  tabindex="0"
                                >
                                  <input
                                    type="radio"
                                    ng-model="app.settings.themeID"
                                    v-model="editData.theme"
                                    value="bg-info|bg-white|bg-black black-info"
                                    class="ng-pristine ng-untouched ng-valid"
                                    name="b"
                                    aria-checked="false"
                                    tabindex="-1"
                                    aria-invalid="false"
                                  />
                                  <span class="block bg-light clearfix pos-rlt">
                                    <span
                                      class="active pos-abt w-full h-full bg-black-opacity text-center"
                                    >
                                      <i class="glyphicon glyphicon-ok text-white m-t-xs"></i>
                                    </span>
                                    <b class="bg-info header"></b>
                                    <b class="bg-white header"></b>
                                    <b class="bg-black black-info"></b>
                                  </span>
                                </label>

                                <label
                                  class="i-checks block m-b-sm"
                                  ng-click="
          app.settings.navbarHeaderColor='bg-success';
          app.settings.navbarCollapseColor='bg-white-only';
          app.settings.asideColor='bg-dark';
         "
                                  role="button"
                                  tabindex="0"
                                >
                                  <input
                                    type="radio"
                                    ng-model="app.settings.themeID"
                                    v-model="editData.theme"
                                    value="bg-success|bg-white|bg-dark dark-success"
                                    class="ng-pristine ng-untouched ng-valid"
                                    name="b"
                                    aria-checked="false"
                                    tabindex="-1"
                                    aria-invalid="false"
                                  />
                                  <span class="block bg-light clearfix pos-rlt">
                                    <span
                                      class="active pos-abt w-full h-full bg-black-opacity text-center"
                                    >
                                      <i class="glyphicon glyphicon-ok text-white m-t-xs"></i>
                                    </span>
                                    <b class="bg-success header"></b>
                                    <b class="bg-white header"></b>
                                    <b class="bg-dark dark-success"></b>
                                  </span>
                                </label>

                                <label class="i-checks block">
                                  <input
                                    type="radio"
                                    ng-model="app.settings.themeID"
                                    v-model="editData.theme"
                                    value="bg-danger|bg-white|bg-dark dark-danger"
                                    class="ng-pristine ng-untouched ng-valid"
                                    name="b"
                                    aria-checked="false"
                                    tabindex="-1"
                                    aria-invalid="false"
                                  />
                                  <span class="block bg-light clearfix pos-rlt">
                                    <span
                                      class="active pos-abt w-full h-full bg-black-opacity text-center"
                                    >
                                      <i class="glyphicon glyphicon-ok text-white m-t-xs"></i>
                                    </span>
                                    <b class="bg-danger header"></b>
                                    <b class="bg-white header"></b>
                                    <b class="bg-dark dark-danger"></b>
                                  </span>
                                </label>
                                <label class="i-checks block">
                                  <input
                                    type="radio"
                                    ng-model="app.settings.themeID"
                                    v-model="editData.theme"
                                    value="bg-danger|bg-white|bg-light light-danger"
                                    class="ng-pristine ng-untouched ng-valid"
                                    name="b"
                                    aria-checked="false"
                                    tabindex="-1"
                                    aria-invalid="false"
                                  />
                                  <span class="block bg-light clearfix pos-rlt">
                                    <span
                                      class="active pos-abt w-full h-full bg-black-opacity text-center"
                                    >
                                      <i class="glyphicon glyphicon-ok text-white m-t-xs"></i>
                                    </span>
                                    <b class="bg-danger header"></b>
                                    <b class="bg-white header"></b>
                                    <b class="bg-light light-danger"></b>
                                  </span>
                                </label>
                              </div>
                              <div class="col-xs-6">
                                <label
                                  class="i-checks block m-b-sm"
                                  ng-click="
          app.settings.navbarHeaderColor='bg-black';
          app.settings.navbarCollapseColor='bg-black';
          app.settings.asideColor='bg-white b-r';
         "
                                  role="button"
                                  tabindex="0"
                                >
                                  <input
                                    type="radio"
                                    ng-model="app.settings.themeID"
                                    v-model="editData.theme"
                                    value="bg-black|bg-black|bg-white white-black"
                                    class="ng-pristine ng-untouched ng-valid"
                                    name="b"
                                    aria-checked="false"
                                    tabindex="-1"
                                    aria-invalid="false"
                                  />
                                  <span class="block bg-light clearfix pos-rlt">
                                    <span
                                      class="active pos-abt w-full h-full bg-black-opacity text-center"
                                    >
                                      <i class="glyphicon glyphicon-ok text-white m-t-xs"></i>
                                    </span>
                                    <b class="bg-black header"></b>
                                    <b class="bg-black header"></b>
                                    <b class="bg-white white-black"></b>
                                  </span>
                                </label>

                                <label
                                  class="i-checks block m-b-sm"
                                  ng-click="
          app.settings.navbarHeaderColor='bg-dark';
          app.settings.navbarCollapseColor='bg-dark';
          app.settings.asideColor='bg-light';
         "
                                  role="button"
                                  tabindex="0"
                                >
                                  <input
                                    type="radio"
                                    name="b"
                                    ng-model="app.settings.themeID"
                                    v-model="editData.theme"
                                    value="bg-dark|bg-dark|bg-light light-dark"
                                    class="ng-pristine ng-untouched ng-valid"
                                    aria-checked="false"
                                    tabindex="-1"
                                    aria-invalid="false"
                                  />
                                  <span class="block bg-light clearfix pos-rlt">
                                    <span
                                      class="active pos-abt w-full h-full bg-black-opacity text-center"
                                    >
                                      <i class="glyphicon glyphicon-ok text-white m-t-xs"></i>
                                    </span>
                                    <b class="bg-dark header"></b>
                                    <b class="bg-dark header"></b>
                                    <b class="bg-light light-dark"></b>
                                  </span>
                                </label>

                                <label
                                  class="i-checks block m-b-sm"
                                  ng-click="
          app.settings.navbarHeaderColor='bg-info dker';
          app.settings.navbarCollapseColor='bg-info dker';
          app.settings.asideColor='bg-light dker b-r';
         "
                                  role="button"
                                  tabindex="0"
                                >
                                  <input
                                    type="radio"
                                    ng-model="app.settings.themeID"
                                    v-model="editData.theme"
                                    value="bg-info|bg-info|bg-light light-info"
                                    class="ng-pristine ng-untouched ng-valid"
                                    name="b"
                                    aria-checked="false"
                                    tabindex="-1"
                                    aria-invalid="false"
                                  />
                                  <span class="block bg-light clearfix pos-rlt">
                                    <span
                                      class="active pos-abt w-full h-full bg-black-opacity text-center"
                                    >
                                      <i class="glyphicon glyphicon-ok text-white m-t-xs"></i>
                                    </span>
                                    <b class="bg-info header"></b>
                                    <b class="bg-info header"></b>
                                    <b class="bg-light light-info"></b>
                                  </span>
                                </label>

                                <label
                                  class="i-checks block m-b-sm"
                                  ng-click="
          app.settings.navbarHeaderColor='bg-primary';
          app.settings.navbarCollapseColor='bg-primary';
          app.settings.asideColor='bg-dark';
         "
                                  role="button"
                                  tabindex="0"
                                >
                                  <input
                                    type="radio"
                                    ng-model="app.settings.themeID"
                                    v-model="editData.theme"
                                    value="bg-primary|bg-primary|bg-dark dark-primary"
                                    class="ng-pristine ng-untouched ng-valid"
                                    name="b"
                                    aria-checked="false"
                                    tabindex="-1"
                                    aria-invalid="false"
                                  />
                                  <span class="block bg-light clearfix pos-rlt">
                                    <span
                                      class="active pos-abt w-full h-full bg-black-opacity text-center"
                                    >
                                      <i class="glyphicon glyphicon-ok text-white m-t-xs"></i>
                                    </span>
                                    <b class="bg-primary header"></b>
                                    <b class="bg-primary header"></b>
                                    <b class="bg-dark dark-primary"></b>
                                  </span>
                                </label>

                                <label
                                  class="i-checks block m-b-sm"
                                  ng-click="
          app.settings.navbarHeaderColor='bg-info dker';
          app.settings.navbarCollapseColor='bg-info dk';
          app.settings.asideColor='bg-black';
         "
                                  role="button"
                                  tabindex="0"
                                >
                                  <input
                                    type="radio"
                                    ng-model="app.settings.themeID"
                                    v-model="editData.theme"
                                    value="bg-info|bg-info|bg-black black-info"
                                    class="ng-pristine ng-untouched ng-valid"
                                    name="b"
                                    aria-checked="false"
                                    tabindex="-1"
                                    aria-invalid="false"
                                  />
                                  <span class="block bg-light clearfix pos-rlt">
                                    <span
                                      class="active pos-abt w-full h-full bg-black-opacity text-center"
                                    >
                                      <i class="glyphicon glyphicon-ok text-white m-t-xs"></i>
                                    </span>
                                    <b class="bg-info header"></b>
                                    <b class="bg-info header"></b>
                                    <b class="bg-black black-info"></b>
                                  </span>
                                </label>

                                <label
                                  class="i-checks block m-b-sm"
                                  ng-click="
          app.settings.navbarHeaderColor='bg-success';
          app.settings.navbarCollapseColor='bg-success';
          app.settings.asideColor='bg-dark';
          "
                                  role="button"
                                  tabindex="0"
                                >
                                  <input
                                    type="radio"
                                    ng-model="app.settings.themeID"
                                    v-model="editData.theme"
                                    value="bg-success|bg-success|bg-dark dark-success"
                                    class="ng-pristine ng-untouched ng-valid"
                                    name="b"
                                    aria-checked="false"
                                    tabindex="-1"
                                    aria-invalid="false"
                                  />
                                  <span class="block bg-light clearfix pos-rlt">
                                    <span
                                      class="active pos-abt w-full h-full bg-black-opacity text-center"
                                    >
                                      <i class="glyphicon glyphicon-ok text-white m-t-xs"></i>
                                    </span>
                                    <b class="bg-success header"></b>
                                    <b class="bg-success header"></b>
                                    <b class="bg-dark dark-success"></b>
                                  </span>
                                </label>

                                <label class="i-checks block">
                                  <input
                                    type="radio"
                                    v-model="editData.theme"
                                    value="bg-danger|bg-danger|bg-dark dark-danger"
                                    class="ng-pristine ng-untouched ng-valid"
                                    name="b"
                                    aria-checked="false"
                                    tabindex="-1"
                                    aria-invalid="false"
                                  />
                                  <span class="block bg-light clearfix pos-rlt">
                                    <span
                                      class="active pos-abt w-full h-full bg-black-opacity text-center"
                                    >
                                      <i class="glyphicon glyphicon-ok text-white m-t-xs"></i>
                                    </span>
                                    <b class="bg-danger header"></b>
                                    <b class="bg-danger header"></b>
                                    <b class="bg-dark dark-danger"></b>
                                  </span>
                                </label>

                                <label class="i-checks block">
                                  <input
                                    type="radio"
                                    v-model="editData.theme"
                                    value="bg-danger|bg-danger|bg-light light-danger"
                                    class="ng-pristine ng-untouched ng-valid"
                                    name="b"
                                    aria-checked="false"
                                    tabindex="-1"
                                    aria-invalid="false"
                                  />
                                  <span class="block bg-light clearfix pos-rlt">
                                    <span
                                      class="active pos-abt w-full h-full bg-black-opacity text-center"
                                    >
                                      <i class="glyphicon glyphicon-ok text-white m-t-xs"></i>
                                    </span>
                                    <b class="bg-danger header"></b>
                                    <b class="bg-danger header"></b>
                                    <b class="bg-light light-danger"></b>
                                  </span>
                                </label>
                                <label class="i-checks block">
                                  <input
                                    type="radio"
                                    v-model="editData.theme"
                                    value="bg-success|bg-success|bg-light light-success"
                                    class="ng-pristine ng-untouched ng-valid"
                                    name="b"
                                    aria-checked="false"
                                    tabindex="-1"
                                    aria-invalid="false"
                                  />
                                  <span class="block bg-light clearfix pos-rlt">
                                    <span
                                      class="active pos-abt w-full h-full bg-black-opacity text-center"
                                    >
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
      <add-modal ref="addModal" v-on:refresh-data="querySearch"></add-modal>
    </div>
  </div>
</template>

<script>
import XLSX from "xlsx";
import pager from "vue-simple-pager";
import PullTo from "vue-pull-to";
import AddModal from "./addModal.vue";
export default {
  components: {
    "add-modal": AddModal,
    "bootstrap-modal": require("vue2-bootstrap-modal"),
    pager,
    PullTo
  },
  data() {
    console.log("this.$env.api.host:",this.$env.conf.api.host)

    return {
       options: {
        target: this.$env.conf.api.host + "/image/upload", //上传地址
        testChunks: false,
        withCredentials: true, //携带jwt
        singleFile: true //单文件上传
      },
      attrs: {
        accept: "image/*"
      },
      errorImg: "",
      statusText: {
        success: "上传成功",
        error: "出错了",
        uploading: "上传中",
        paused: "暂停中",
        waiting: "等待中"
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
        secret: "",
        wechat_status: "1",
        callbackurl: ""
      },
      editData: {},
      secrectData: { id: 0, secret: "" },
      importData: { id: 0, menus: [] },
      enableData: { id: null, status: null },
      pi: 1,
      ps: 10,
      totalPage: 0,
      pageAuth: [],
      auth: { addsys: false },

      errorTemplate: {
        920: "当前系统下面已存在菜单数据,不能导入"
      }
    };
  },
  props: ["path"],
  mounted() {
    this.$refs.main.style.height = document.documentElement.clientHeight + "px";
    this.query();
  },
  computed: {
    editLayout() {
      return this.editData.layout.split(" ");
    }
  },
  methods: {
    fileEditSuccess(rootFile, file, message, chunk) {
      console.log("fileEditSuccess:", message);
      let data = JSON.parse(message);
      this.editData.logo = data.data;
    },
    //上传失败事件
    fileError(rootFile, file, message, chunk) {
      this.$notify({
        title: "错误",
        message: "上传失败，请稍后再试",
        type: "error",
        offset: 50,
        duration: 2000
      });
    },
    Add() {
      this.$refs.addModal.setModal();
    },
    handleSizeChange(val) {
      this.ps = val;
      this.query();
    },
    goPage(val) {
      this.pi = val;
      this.query();
    },

    //点查询事件
    querySearch() {
      this.pi = 1;
      this.ps = 10;
      this.query();
    },

    query() {
      this.$http
        .post("/sys/index/getall", {
          pi: this.pi,
          ps: this.ps,
          name: this.sysname,
          status: this.selected
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
    openAdd() {
      this.$refs.addModal.open();
    },
    deleteById(id) {
      this.$confirm("确定执行此操作?是否继续?", "提示", {
        confirmButtonText: "确定",
        cancelButtonText: "取消",
        type: "warning"
      }).then(() => {
        this.$http
          .post("/sys/index/del", { id: id })
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

    cancel() {
      this.$refs.editModal.close();
    },
    secretCancel() {
      this.$refs.secretModal.close();
    },
    enable(id, status) {
      this.enableData.id = id;
      this.enableData.status = status;

      this.$confirm("确定执行此操作?是否继续?", "提示", {
        confirmButtonText: "确定",
        cancelButtonText: "取消",
        type: "warning"
      }).then(() => {
        this.$http
          .post("/sys/index/changestatus", {
            id: this.enableData.id,
            status: this.enableData.status
          })
          .then(res => {
            this.goPage(this.pi);
            this.$notify({
              title: "成功",
              message: "状态修改成功",
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

    //修改秘钥
    setSecret(id) {
      this.secrectData = {
        id: id,
        secret: ""
      };
      this.$refs.secretModal.open();
    },

    saveSecret() {
      let secretInfo = this.secrectData;
      let msg = this.checkSecretInfo(this.secrectData);
      if (msg) {
        this.$notify({
          title: "错误",
          message: msg,
          type: "error",
          offset: 50,
          duration: 2000
        });
        return;
      }
      this.$http
        .post("/sys/index/changesecret", secretInfo)
        .then(res => {
          this.$refs.secretModal.close();
          this.$notify({
            title: "成功",
            message: "秘钥修改成功",
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
      let edit = Object.assign({}, this.editData);
      this.editData.layout.forEach((item, index) => {
        str += item + " ";
      });
      edit.layout = str.trim();

      var msg = this.checkBeforeSave(edit);
      if (msg) {
        this.$notify({
          title: "错误",
          message: msg,
          type: "error",
          offset: 50,
          duration: 2000
        });
        return;
      }
      this.$http
        .post("/sys/index/edit", edit)
        .then(res => {
          this.$refs.editModal.close();
          this.goPage({ page: this.pi });
          this.$notify({
            title: "成功",
            message: "编辑成功",
            type: "success",
            offset: 50,
            duration: 2000
          });
        })
        .catch(err => {
          if (err.response && err.response.status == 911) {
            this.$notify({
              title: "失败",
              message: "系统名称或英文名称已存在",
              type: "error",
              offset: 50,
              duration: 2000
            });
            return
          }
          this.$notify({
            title: "错误",
            message: "网络错误,请稍后再试",
            type: "error",
            offset: 50,
            duration: 2000
          });
        });
    },
    manage(id, ident) {
      this.$emit("addTab", "菜单配置(" + ident + ")", "/sys/index/func?id=" + id);
    },
    //管理数据权限数据
    managePermission(id, ident) {
      this.$emit(
        "addTab",
        "数据规则配置(" + ident + ")",
        "/sys/index/datapermission?id=" + id
      );

      // this.$router.push({
      //   name: "datapermission",
      //   query: {
      //     id: id
      //   }
      // });
    },

    checkBeforeSave(editData) {
      if (!editData.name) {
        return "系统名称不能为空";
      }
      if (!editData.ident) {
        return "系统英文名称不能为空";
      }
      // if(!editData.secret) {
      //   return "secret不能为空";
      // }
      // if (!editData.logo) {
      //   return "logo图片必须上传";
      // }
      if (!editData.theme) {
        return "请选择主题样式";
      }
      if (!editData.layout) {
        return "请选择页面布局样式";
      }
      return "";
    },

    checkSecretInfo(secretInfo) {
      if (!secretInfo.secret) {
        return "秘钥不能为空";
      }
    },
    deletePic() {
      this.editData.logo = "";
    },

    //导出菜单
    exportMenu(id) {
      this.$http
        .post("/sys/index/menuexport", { id: id })
        .then(res => {
          var data = [
            [
              "id",
              "name",
              "parent",
              "level_id",
              "icon",
              "path",
              "enable",
              "sortrank",
              "is_open"
            ]
          ];
          res.forEach(element => {
            data.push([
              element["id"],
              element["name"],
              element["parent"],
              element["level_id"],
              element["icon"],
              element["path"],
              element["enable"],
              element["sortrank"],
              element["is_open"]
            ]);
          });
          var sheet = XLSX.utils.aoa_to_sheet(data);
          var name = "菜单" + this.$utility.dateFormat(new Date(),"yyyyMMddhhmm") + ".xlsx";
          this.openDownloadDialog(this.sheet2blob(sheet), name);
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

    fileChange(ev) {
      var fileName = ev.target.files[0]["name"];
      console.log(fileName);
      var str = fileName.split(".");
      if (str[1] != "xlsx" && str[1] != "xls") {
        this.$notify({
          title: "提示",
          message: "文件格式有误(只支持excel文件)",
          type: "error",
          offset: 50,
          duration: 2000
        });
        return false;
      }
      let f = ev.target.files[0],
        reader = new FileReader();
      reader.onload = e => {
        let data = e.target.result;
        let wb = XLSX.read(data, {
          type: "array"
        });
        let jsonData = XLSX.utils.sheet_to_json(wb.Sheets[wb.SheetNames[0]]);
        //console.log("jsonData: ", jsonData);
        this.importData.menus = jsonData;
      };
      reader.readAsArrayBuffer(f);
    },

    //保存要导入的菜单
    saveImportMenu() {
      console.log("菜单数据:", this.importData);
      if (this.importData.menus.length == 0) {
        this.$notify({
          title: "提示",
          message: "菜单数据为空",
          type: "error",
          offset: 50,
          duration: 2000
        });
        return false;
      }

      this.$http
        .post("/sys/index/menuimport", { data: JSON.stringify(this.importData) })
        .then(res => {
          this.$notify({
            title: "成功",
            message: "导入菜单成功",
            type: "success",
            offset: 50,
            duration: 2000
          });

          this.importClose();
        })
        .catch(err => {
          var msg = err;
          if (err.response) {
            msg =
              this.errorTemplate[err.response.status] || "网络错误,请稍后再试";
          }

          this.$notify({
            title: "错误",
            message: msg,
            type: "error",
            offset: 50,
            duration: 2000
          });
        });
    },

    //导入菜单
    importMenu(id) {
      this.importData.id = id;
      this.importData.menus = [];
      this.$refs.importModal.open();
    },

    importClose() {
      this.$refs.importModal.close();
      $("#upload_file").val(null);
    },

    sheet2blob(sheet, sheetName) {
      sheetName = sheetName || "sheet1";
      var workbook = {
        SheetNames: [sheetName],
        Sheets: {}
      };
      workbook.Sheets[sheetName] = sheet;
      // 生成excel的配置项
      var wopts = {
        bookType: "xlsx", // 要生成的文件类型
        bookSST: false, // 是否生成Shared String Table，官方解释是，如果开启生成速度会下降，但在低版本IOS设备上有更好的兼容性
        type: "binary"
      };
      var wbout = XLSX.write(workbook, wopts);
      var blob = new Blob([s2ab(wbout)], {
        type: "application/octet-stream"
      });
      // 字符串转ArrayBuffer
      function s2ab(s) {
        var buf = new ArrayBuffer(s.length);
        var view = new Uint8Array(buf);
        for (var i = 0; i != s.length; ++i) view[i] = s.charCodeAt(i) & 0xff;
        return buf;
      }
      return blob;
    },
    openDownloadDialog(url, saveName) {
      if (typeof url == "object" && url instanceof Blob) {
        url = URL.createObjectURL(url);
      }
      var aLink = document.createElement("a");
      aLink.href = url;
      aLink.download = saveName || "";
      var event;
      if (window.MouseEvent) event = new MouseEvent("click");
      else {
        event = document.createEvent("MouseEvents");
        event.initMouseEvent(
          "click",
          true,
          false,
          window,
          0,
          0,
          0,
          0,
          0,
          false,
          false,
          false,
          false,
          0,
          null
        );
      }
      aLink.dispatchEvent(event);
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
.page-pagination {
  padding: 10px 15px;
  text-align: right;
  margin-bottom: 50px;
}
</style>

