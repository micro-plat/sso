//系统功能查询
<template>
  <div class="app-content-body fade-in-up ng-scope">
    <div class="fade-in-down ng-scope">
      <div class="wrapper wrapper-lg">
            <vue-ztree :list='ztreeDataSource' :func='nodeClick' :checkfunc="checkFunc" :expand='expandClick'
                       :contextmenu='contextmenuClick' :is-open='true'></vue-ztree>
      </div>
    </div>
    <bootstrap-modal ref="editModal1" :need-header="true" :need-footer="true" >
      <div slot="title">
        编辑功能
      </div>
      <div slot="body">
        <div class="panel panel-default">
          <div class="panel-body">
            <form role="form" class="ng-pristine ng-valid ng-submitted">
              <div class="form-group">
                <label>功能名称</label>
                <input class="form-control" placeholder="" v-validate="'required'" name="name1"
                       v-model="currentData.name" type="text">
                <div class="form-height text-danger"><span v-show="errors.first('name1')">功能名称不能为空</span></div>
                <input class="form-control" placeholder="" v-model="currentData.id" type="hidden">
              </div>
            </form>
          </div>
        </div>
      </div>
      <div slot="footer">
        <a class="btn btn-sm btn-danger" @click="editFunc">提交</a>
        <a class="btn btn-sm btn-primary" @click="cancel">取消</a>
      </div>
    </bootstrap-modal>
    <bootstrap-modal ref="editModal2" :need-header="true" :need-footer="true">
      <div slot="title">
        编辑功能
      </div>
      <div slot="body">
        <div class="panel panel-default">
          <div class="panel-body">
            <form role="form" class="ng-pristine ng-valid ng-submitted height-min">
              <div class="form-group">
                <label>功能名称</label>
                <input class="form-control" placeholder="" v-validate="'required'" name="name1"
                       v-model="currentData.name" type="text">
                <div class="form-height text-danger"><span v-show="errors.first('name1')">功能名称不能为空</span></div>
                <input class="form-control" placeholder="" v-model="currentData.id" type="hidden">
              </div>
              <div class="form-group">
                <label>是否展开</label>
                <div class="radio">
                  <label class="i-checks">
                    <input type="radio" name="is_open1" v-model="currentData.is_open"  value="1" class="ng-pristine ng-untouched ng-valid">
                    <i></i>
                    展开
                  </label>
                  <label class="i-checks">
                    <input type="radio" name="is_open1" v-model="currentData.is_open" value="0" class="ng-pristine ng-untouched ng-valid">
                    <i></i>
                    收起
                  </label>
                </div>
                {{ errors.first('is_open1') }}
              </div>
              <div class="form-group">
                <label>图标</label>
                <span class="fa-stack fa-lg">
                
                     <i :class="currentData.icon +' '+ currentData.color" v-show="!currentData.iconTemp"></i>
                    <i :class="currentData.iconTemp +' '+ currentData.color" v-show="currentData.iconTemp"></i>
                  
                </span>
                <input class="form-control" placeholder="" v-validate="'required'" name="icon"
                       v-model="currentData.iconTemp +' '+ currentData.color" type="hidden">
                <div class="form-height text-danger"><span v-show="errors.first('icon')">图标不能为空</span></div>
              </div>
              <div class="form-group">
                <div class="settings">
                  <!--<div class="container">-->
                    <div class="row">
                        <!--<div class="panel-body ng-scope">-->
                            <div class="col col-sm-12 col-xs-8">
                              <label class="i-checks inline m-b-sm" v-for="(v,k) in iconList" :key="k">
                                <input type="radio" name="b" :value="v" v-model="currentData.iconTemp"
                                       class="ng-pristine ng-untouched ng-valid" />
                                <span class="inline clearfix pos-rlt icon-height">
            <span class="active pos-abt w-full h-full bg-black-opacity text-center">
               <i class="glyphicon glyphicon-ok text-white m-t-xs"></i>
            </span>
            <span class="fa-lg icon-padding">
              <i :class="v" ></i>
            </span>
          </span>
                              </label>
                            </div>
                        <!--</div>-->
                    </div>
                  <!--</div>-->
                </div>
              </div>
              <div class="form-group">
                <label>颜色</label>
                <div class="settings">
                  <!--<div class="container">-->
                    <div class="row">
                        <!--<div class="panel-body ng-scope">-->
                            <div class="col col-sm-12 col-xs-8">
                              <label class="i-checks inline m-b-sm icon-color" v-for="(v,k) in colorList" :key="k">
                                <input type="radio" name="c" :value="v" v-model="currentData.color"
                                       class="ng-pristine ng-untouched ng-valid" />
                                <span class="inline clearfix pos-rlt icon-height">
            <span class="active pos-abt w-full h-full bg-black-opacity text-center">
               <i class="glyphicon glyphicon-ok text-white m-t-xs"></i>
            </span>
            <i :class="v"><i class="fa  fa-circle icon-color-fa"></i></i>
          </span>
                              </label>
                            </div>
                        <!--</div>-->
                    </div>
                  <!--</div>-->
                </div>
              </div>
            </form>
          </div>
        </div>
      </div>
      <div slot="footer">
        <a class="btn btn-sm btn-danger" @click="editFunc">提交</a>
        <a class="btn btn-sm btn-primary" @click="cancel">取消</a>
      </div>
    </bootstrap-modal>
    <bootstrap-modal ref="editModal3" :need-header="true" :need-footer="true">
      <div slot="title">
        编辑功能
      </div>
      <div slot="body">
        <div class="panel panel-default">
          <div class="panel-body">
            <form role="form" class="ng-pristine ng-valid ng-submitted height-min">
              <div class="form-group">
                <label>功能名称</label>
                <input class="form-control" placeholder="" v-validate="'required'" name="name2"
                       v-model="currentData.name" type="text">
                <div class="form-height text-danger"><span v-show="errors.first('name2')">功能名称不能为空</span></div>
                <input class="form-control" placeholder="" v-model="currentData.id" type="hidden">
              </div>
              <div class="form-group">
                <label>地址</label>
                <input class="form-control" placeholder="" v-validate="'required'" name="path"
                       v-model="currentData.path" type="text">
                <div class="form-height text-danger"><span v-show="errors.first('path')">地址不能为空</span></div>
              </div>
              <div class="form-group">
                <label>排序编号</label>
                <input class="form-control" placeholder="" v-validate="'required'" name="sortrank" 
                      v-model="currentData.sortrank" type="text">
                      <div class="form-height text-danger"><span v-show="errors.first('sortrank')">排序编号不能为空</span></div>
              </div>
              <div class="form-group">
                <label>图标</label>
                <span class="fa-stack fa-lg">
                     <i :class="currentData.icon +' '+ currentData.color" v-show="!currentData.iconTemp"></i>
                    <i :class="currentData.iconTemp +' '+ currentData.color" v-show="currentData.iconTemp"></i>
                </span>
                <input class="form-control" placeholder="" v-validate="'required'" name="icon"
                       v-model="currentData.iconTemp +' '+currentData.color" type="hidden">
                <div class="form-height text-danger"><span v-show="errors.first('icon')">图标不能为空</span></div>
              </div>
              <div class="form-group">
                <div class="settings">
                  <!--<div class="container">-->
                  <div class="row">
                    <!--<div class="panel-body ng-scope">-->
                    <div class="col col-sm-12 col-xs-8">
                      <label class="i-checks inline m-b-sm" v-for="(v,k) in iconList" :key="k">
                        <input type="radio" name="b" :value="v" v-model="currentData.iconTemp"
                               class="ng-pristine ng-untouched ng-valid" />
                        <span class="inline clearfix pos-rlt icon-height">
            <span class="active pos-abt w-full h-full bg-black-opacity text-center">
               <i class="glyphicon glyphicon-ok text-white m-t-xs"></i>
            </span>
             <span class="fa-lg icon-padding">
              <i :class="v"></i>
            </span>
          </span>
                      </label>
                    </div>
                    <!--</div>-->
                  </div>
                  <!--</div>-->
                </div>
              </div>
              <div class="form-group">
                <label>颜色</label>
                <div class="settings">
                  <!--<div class="container">-->
                  <div class="row">
                    <!--<div class="panel-body ng-scope">-->
                    <div class="col col-sm-12 col-xs-8">
                      <label class="i-checks inline m-b-sm" v-for="(v,k) in colorList" :key="k">
                        <input type="radio" name="c" :value="v" v-model="currentData.color"
                               class="ng-pristine ng-untouched ng-valid" />
                        <span class="inline clearfix pos-rlt">
            <span class="active pos-abt w-full h-full bg-black-opacity text-center">
               <i class="glyphicon glyphicon-ok text-white m-t-xs"></i>
            </span>
            <i :class="v"><i class="fa  fa-circle icon-color-fa"></i></i>
          </span>
                      </label>
                    </div>

                    <!--</div>-->


                  </div>
                  <!--</div>-->
                </div>
              </div>

            </form>
          </div>
        </div>
      </div>
      <div slot="footer">
        <a class="btn btn-sm btn-danger" @click="editFunc">提交</a>
        <a class="btn btn-sm btn-primary" @click="cancel">取消</a>
      </div>
    </bootstrap-modal>
    <bootstrap-modal ref="editModal4" :need-header="true" :need-footer="true">
      <div slot="title">
        添加页面权限
      </div>
      <div slot="body">
        <div class="panel panel-default">
          <div class="panel-body">
            <form role="form" class="ng-pristine ng-valid ng-submitted height-min">
              <div class="form-group">
                <label>权限名称</label>
                <input class="form-control" placeholder="" v-validate="'required'" name="name3"
                       v-model="currentData.name" type="text">
                <div class="form-height text-danger"><span v-show="errors.first('name4')">权限名称不能为空</span></div>
                <input class="form-control" placeholder="" v-model="currentData.id" type="hidden">
              </div>
              <div class="form-group">
                <label>标签（tag）</label>
                <input class="form-control" placeholder="" v-validate="'required'" name="path3"
                       v-model="currentData.path" type="text">
                <div class="form-height text-danger"><span v-show="errors.first('path4')">地址不能为空</span></div>
              </div>
            </form>
          </div>
        </div>
      </div>
      <div slot="footer">
        <a class="btn btn-sm btn-danger" @click="editFunc">提交</a>
        <a class="btn btn-sm btn-primary" @click="cancel">取消</a>
      </div>
    </bootstrap-modal>
    <bootstrap-modal ref="addModal1" :need-header="true" :need-footer="true">
      <div slot="title">
        添加功能
      </div>
      <div slot="body">
        <div class="panel panel-default">
          <div class="panel-body">
            <form role="form" class="ng-pristine ng-valid ng-submitted">
              <div class="form-group">
                <label>功能名称</label>
                <input class="form-control" placeholder="" v-validate="'required'" name="name3"
                       v-model="currentData.name" type="text">
                <div class="form-height text-danger"><span v-show="errors.first('name3')">功能名称不能为空</span></div>
                <input class="form-control" placeholder="" v-model="currentData.id" type="hidden">
              </div>
            </form>
          </div>
        </div>
      </div>
      <div slot="footer">
        <a class="btn btn-sm btn-danger" @click="addFunc">提交</a>
        <a class="btn btn-sm btn-primary" @click="cancel">取消</a>
      </div>
    </bootstrap-modal>
    <bootstrap-modal ref="addModal2" :need-header="true" :need-footer="true">
      <div slot="title">
        添加功能
      </div>
      <div slot="body">
        <div class="panel panel-default">
          <div class="panel-body">
            <form role="form" class="ng-pristine ng-valid ng-submitted height-min">
              <div class="form-group">
                <label>功能名称</label>
                <input class="form-control" placeholder="" v-validate="'required'" name="name"
                       v-model="currentData.name" type="text">
                {{ errors.first('name') }}
                <input class="form-control" placeholder="" v-model="currentData.id" type="hidden">
              </div>
              <div class="form-group">
                <label>是否展开</label>
                <div class="radio">
                  <label class="i-checks">
                    <input type="radio" name="is_open" v-model="currentData.is_open"  value="1" class="ng-pristine ng-untouched ng-valid">
                    <i></i>
                    展开
                  </label>
                  <label class="i-checks">
                    <input type="radio" name="is_open" v-model="currentData.is_open" value="0" class="ng-pristine ng-untouched ng-valid">
                    <i></i>
                    收起
                  </label>
                </div>
                {{ errors.first('is_open') }}
              </div>
              <div class="form-group">
                <label>图标</label>
                <span class="fa-stack fa-lg">
                   <i :class="currentData.iconTemp +' '+ currentData.color"></i>
                </span>
                <input class="form-control" placeholder="" v-validate="'required'" name="icon"
                       v-model="currentData.iconTemp +' '+currentData.color" type="hidden">
                <div class="form-height text-danger"><span v-show="errors.first('icon')">图标不能为空</span></div>
              </div>
              <div class="form-group">
                <div class="settings">
                  <!--<div class="container">-->
                  <div class="row">
                    <!--<div class="panel-body ng-scope">-->
                    <div class="col col-sm-12 col-xs-8">
                      <label class="i-checks inline m-b-sm" v-for="(v,k) in iconList" :key="k">
                        <input type="radio" name="c" :value="v" v-model="currentData.iconTemp"
                               class="ng-pristine ng-untouched ng-valid" />
                        <span class="inline clearfix pos-rlt icon-height">
            <span class="active pos-abt w-full h-full bg-black-opacity text-center">
               <i class="glyphicon glyphicon-ok text-white m-t-xs"></i>
            </span>
            <span class="fa-lg icon-padding">
              <i :class="v"></i>
            </span>
          </span>
                      </label>
                    </div>
                    <!--</div>-->
                  </div>
                  <!--</div>-->
                </div>
              </div>
              <div class="form-group">
                <label>颜色</label>
                <div class="settings">
                  <!--<div class="container">-->
                  <div class="row">
                    <!--<div class="panel-body ng-scope">-->
                    <div class="col col-sm-12 col-xs-8">
                      <label class="i-checks inline m-b-sm" v-for="(v,k) in colorList" :key="k">
                        <input type="radio" name="b" :value="v" v-model="currentData.color"
                               class="ng-pristine ng-untouched ng-valid" />
                        <span class="inline clearfix pos-rlt">
            <span class="active pos-abt w-full h-full bg-black-opacity text-center">
               <i class="glyphicon glyphicon-ok text-white m-t-xs"></i>
            </span>
            <i :class="v"><i class="fa fa-circle icon-color-fa"></i></i>
          </span>
                      </label>
                    </div>
                    <!--</div>-->
                  </div>
                  <!--</div>-->
                </div>
              </div>
            </form>
          </div>
        </div>
      </div>
      <div slot="footer">
        <a class="btn btn-sm btn-danger" @click="addFunc">提交</a>
        <a class="btn btn-sm btn-primary" @click="cancel">取消</a>
      </div>
    </bootstrap-modal>
    <bootstrap-modal ref="addModal3" :need-header="true" :need-footer="true">
      <div slot="title">
        添加功能
      </div>
      <div slot="body">
        <div class="panel panel-default">
          <div class="panel-body">
            <form role="form" class="ng-pristine ng-valid ng-submitted height-min">
              <div class="form-group">
                <label>功能名称</label>
                <input class="form-control" placeholder="" v-validate="'required'" name="name3"
                       v-model="currentData.name" type="text">
                <div class="form-height text-danger"><span v-show="errors.first('name3')">功能名称不能为空</span></div>
                <input class="form-control" placeholder="" v-model="currentData.id" type="hidden">
              </div>
              <div class="form-group">
                <label>地址</label>
                <input class="form-control" placeholder="" v-validate="'required'" name="path3"
                       v-model="currentData.path" type="text">
                <div class="form-height text-danger"><span v-show="errors.first('path3')">地址不能为空</span></div>
              </div>
              <!--<div class="form-group">
                <label>排序编号</label>
                <input class="form-control" placeholder="" v-validate="'required'" name="sortrank" 
                      v-model="currentData.sortrank" type="text">
                <div class="form-height text-danger"><span v-show="errors.first('sortrank')"></span></div>
              </div>-->
              <div class="form-group">
                <label>图标</label>
                <span class="fa-stack fa-lg">
                   <i :class="currentData.iconTemp +' '+ currentData.color"></i>
                </span>
                <input class="form-control" placeholder="" v-validate="'required'" name="icon"
                       v-model="currentData.iconTemp + ' ' + currentData.color" type="hidden">
                <div class="form-height text-danger"><span v-show="errors.first('icon')">图标不能为空</span></div>
              </div>
              <div class="form-group">
                <div class="settings">
                  <!--<div class="container">-->
                  <div class="row">
                    <!--<div class="panel-body ng-scope">-->
                    <div class="col col-sm-12 col-xs-8">
                      <label class="i-checks inline m-b-sm" v-for="(v,k) in iconList" :key="k">
                        <input type="radio" name="b" :value="v" v-model="currentData.iconTemp"
                               class="ng-pristine ng-untouched ng-valid" />
                        <span class="inline clearfix pos-rlt icon-height">
            <span class="active pos-abt w-full h-full bg-black-opacity text-center">
               <i class="glyphicon glyphicon-ok text-white m-t-xs"></i>
            </span>
             <span class="fa-lg icon-padding">
              <i :class="v"></i>
            </span>
          </span>
                      </label>
                    </div>
                    <!--</div>-->
                  </div>
                  <!--</div>-->
                </div>
              </div>
              <div class="form-group">
                <label>颜色</label>
                <div class="settings">
                  <!--<div class="container">-->
                  <div class="row">
                    <!--<div class="panel-body ng-scope">-->
                    <div class="col col-sm-12 col-xs-8">
                      <label class="i-checks inline m-b-sm" v-for="(v,k) in colorList" :key="k">
                        <input type="radio" name="c" :value="v" v-model="currentData.color"
                               class="ng-pristine ng-untouched ng-valid" />
                        <span class="inline clearfix pos-rlt">
            <span class="active pos-abt w-full h-full bg-black-opacity text-center">
               <i class="glyphicon glyphicon-ok text-white m-t-xs"></i>
            </span>
            <i :class="v"><i class="fa fa-circle icon-color-fa"></i></i>
          </span>
                      </label>
                    </div>
                    <!--</div>-->
                  </div>
                  <!--</div>-->
                </div>
              </div>
            </form>
          </div>
        </div>
      </div>
      <div slot="footer">
        <a class="btn btn-sm btn-danger" @click="addFunc">提交</a>
        <a class="btn btn-sm btn-primary" @click="cancel">取消</a>
      </div>
    </bootstrap-modal>
    <bootstrap-modal ref="addModal4" :need-header="true" :need-footer="true">
      <div slot="title">
        添加页面权限
      </div>
      <div slot="body">
        <div class="panel panel-default">
          <div class="panel-body">
            <form role="form" class="ng-pristine ng-valid ng-submitted height-min">
              <div class="form-group">
                <label>权限名称</label>
                <input class="form-control" placeholder="" v-validate="'required'" name="name3"
                       v-model="currentData.name" type="text">
                <div class="form-height text-danger"><span v-show="errors.first('name4')">权限名称不能为空</span></div>
                <input class="form-control" placeholder="" v-model="currentData.id" type="hidden">
              </div>
              <div class="form-group">
                <label>标签（tag）</label>
                <input class="form-control" placeholder="" v-validate="'required'" name="path3"
                       v-model="currentData.path" type="text">
                <div class="form-height text-danger"><span v-show="errors.first('path4')">地址不能为空</span></div>
              </div>
            </form>
          </div>
        </div>
      </div>
      <div slot="footer">
        <a class="btn btn-sm btn-danger" @click="addFunc">提交</a>
        <a class="btn btn-sm btn-primary" @click="cancel">取消</a>
      </div>
    </bootstrap-modal>
  </div>

</template>


<script>
import vueZtree from "@/components/vue-tree.vue";
import bootstrapModal from "vue2-bootstrap-modal"
export default {
  data() {
    return {
      iconList: [
        "fa fa-address-book",
        "fa fa-address-book-o",
        "fa fa-address-card",
        "fa fa-address-card-o",
        "fa fa-user-circle-o",
        "fa fa-user-circle",
        "fa fa-bell",
        "fa fa-bell-o",
        "fa fa-bell-o",
        "fa fa-commenting",
        "fa fa-commenting-o",
        "fa fa-check-circle-o",
        "fa fa-circle",
        "fa fa-envelope",
        "fa fa-envelope-open",
        "fa fa-exclamation-circle",
        "fa fa-exclamation-triangle",
        "fa fa-exchange",
        "fa fa-external-link",
        "fa fa-file-image-o",
        "fa fa-folder",
        "fa fa-folder-o",
        "fa fa-folder-open",
        "fa fa-folder-open-o",
        "fa fa-cog",
        "fa fa-users",
        "fa fa-reply",
        "fa fa-share",
        "fa fa-map-marker",
        "fa fa-microphone",
        "fa fa-microphone-slash",
        "fa fa-minus-circle",
        "fa fa-paper-plane",
        "fa fa-paper-plane-o",
        "fa fa-pencil-square-o",
        "fa fa-plus-circle",
        "fa fa-power-off",
        "fa fa-podcast",
        "fa fa-recycle",
        "fa fa-pencil-square-o",
        "fa fa-refresh",
        "fa fa-share-alt",
        "fa fa-share-alt-square",
        "fa fa-sitemap",
        "fa fa-star",
        "fa fa-star-half",
        "fa fa-sticky-note",
        "fa fa-sticky-note-o",
        "fa fa-street-view",
        "fa fa-tag",
        "fa fa-tags",
        "fa fa-television",
        "fa fa-times-circle",
        "fa fa-trash",
        "fa fa-user-circle-o",
        "fa fa-user-times",
        "fa fa-user-plus",
        "fa fa-file-text",
        "fa fa-file-image-o",
        "fa fa-check-square-o",
        "fa fa-check-square",
        "fa fa-pie-chart",
        "fa fa-area-chart",
        "fa fa-bar-chart",
        "fa fa-line-chart",
        "fa fa-arrow-circle-down",
        "fa fa-arrow-circle-left",
        "fa fa-arrow-circle-right",
        "fa fa-arrow-circle-up",
        "fa fa-chevron-circle-down",
        "fa fa-chevron-circle-left",
        "fa fa-chevron-circle-right",
        "fa fa-chevron-circle-up",
        "fa fa-angle-double-down",
        "fa fa-angle-double-left",
        "fa fa-angle-double-right",
        "fa fa-angle-double-up",
        "fa fa-play-circle",
        "fa fa-play-circle-o",
        "fa fa-pause",
        "fa fa-pause-circle",
        "fa fa-step-backward",
        "fa fa-step-forward",
        "fa fa-play",
        "fa fa-phone",
        "fa fa-search",
        "fa fa-bell-o",
        "fa fa-bell-slash-o"
      ],
      colorList: [
        "text-success",
        "text-primary",
        "text-info",
        "text-danger"
      ],
      currentData: {},
      id: null,
      showAddr: false,
      ztreeDataSource: [],
      obj: null
    };
  },
  components: {
    vueZtree,
    bootstrapModal
  },
  mounted() {
    this.initData();
  },
  methods: {
    initData() {
      let routerParams = this.$route.query;
      this.id = routerParams.id;
      
      this.$http.get("/system/func", { id: this.id })
        .then(res => {
          if (res.length != 0) {
            this.ztreeDataSource = res;
            return;
          }
          this.ztreeDataSource.push({
            name: "新节点",
            children: [],
            path: "-",
            icon: "-",
            isNew: true,
            parentId: 0,
            parentLevel: 0,
            level_id:1,
          })
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
    editFunc() {
      let data = this.currentData;
      let icon = data.icon;
      if (data.iconTemp != "") {
        icon = data.iconTemp + " " + data.color
      }
      this.$http.post("/system/func/edit", {
        id: data.id,
        name: data.name,
        sortrank:data.sortrank,
        icon: icon,
        path: data.path,
        is_open: data.is_open,
      })
      .then(res => {
        this.initData();
        this.cancel();
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
    addFunc() {
      this.$http.post("/system/func/add", {
        parentid: this.currentData.parentId,
        parentlevel: this.currentData.parentLevel,
        sysid: this.id,
        sortrank: this.currentData.sortrank,
        name: this.currentData.name,
        icon: this.currentData.iconTemp +" "+ this.currentData.color,
        path: this.currentData.path,
        is_open: this.currentData.is_open
      })
      .then(res => {
        this.initData();
        this.cancel();
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

    // 点击节点
    nodeClick: function(d) {
      /*排序 */
      if (d.isSort === true) {
        this.initData();
      } else {
        console.log(d);
        this.$set( d, "color", "");
        this.$set( d, "iconTemp", "");
        this.obj = d;
        if (d.isNew == true) {  //添加
          if (d.parentLevel == 3 || d.level_id == 4) {
            this.currentData = d;
            this.$refs.addModal4.open();
          }
          if (d.parentLevel == 2 || d.level_id == 3) {
            this.currentData = d;
            this.$refs.addModal3.open();
          }
          if (d.parentLevel == 1 || d.level_id == 2) {
            this.currentData = d;
            this.$refs.addModal2.open();
          }
          if (d.parentLevel == 0 || d.level_id == 1) {
            this.currentData = d;
            this.$refs.addModal1.open();
          }
        } else {  //编辑

          if (d.level_id == 4 || d.parentLevel == 3) {
            this.currentData = d;
            this.$refs.editModal4.open();
          }
          if (d.level_id == 3 || d.parentLevel == 2) {
            this.currentData = d;
            this.$refs.editModal3.open();
          }
          if (d.level_id == 2 || d.parentLevel == 1) {
            this.currentData = d;
            this.$refs.editModal2.open();
          }
          if (d.level_id == 1 || d.parentLevel == 0) {
            this.currentData = d;
            this.$refs.editModal1.open();
          }
        }
      }
      
    },
    //点击选中
    checkFunc(m) {
      console.log("m",m)
    },
    // 右击事件
    contextmenuClick: function(m) {},
    // 点击展开收起
    expandClick: function(m) {
      // 点击异步加载
      if (m.isExpand) {
        // 动态加载子节点, 模拟ajax请求数据
        // 请注意 id 不能重复哦。
        if (m.hasOwnProperty("children")) {
          m.loadNode = 1; // 正在加载节点

          setTimeout(() => {
            m.loadNode = 2; // 节点加载完毕
            m.isFolder = !m.isFolder;
          }, 500);
        }
      }
    },
    cancel() {
      this.$refs.editModal1.close();
      this.$refs.editModal2.close();
      this.$refs.editModal3.close();
      this.$refs.editModal4.close();
      this.$refs.addModal1.close();
      this.$refs.addModal2.close();
      this.$refs.addModal3.close();
      this.$refs.addModal4.close();
    }
  }
};
</script>
