<template>
  <div ref="modal" class="modal fade background-darken" tabindex="-1" role="dialog" :class="{in:isOpen,show:isShow}" @click.self="close()" @keyup.esc="close()">
    <div class="modal-dialog" style="width:600px;" role="document">
        <div class="modal-content">
            <div class="modal-header">
              <button type="button" class="close" data-dismiss="modal" aria-label="Close" @click="close()">&times;</button>
                <h4 class="modal-title" align="center"添加系统</h4>
            </div>
           
            <div class="modal-body" >
                <el-form :model="addData"  :inline="true"  ref="addform1" label-position="right" label-width="100px">
                    <!-- <el-row>
                        <el-col :span="12">
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
                                <label>超时时长</label>
                                <input class="form-control" placeholder="请输入超时时常" v-validate="'required|numeric'" v-model="addData.time_out" name="time_out"  type="text">
                                <div class="form-height text-danger"><span v-show="errors.first('time_out')">超时时长不能为空且必须为数字</span> </div>
                            </div>
                        </el-col>
                        <el-col :span="12">
                            <div class="form-group">
                                <label>首页地址</label>
                                <input class="form-control" placeholder="请输入首页地址" v-validate="'required'" name="addr" v-model="addData.addr"  type="text">
                                <div class="form-height text-danger"><span v-show="errors.first('addr')">首页地址不能为空</span></div>
                            </div>
                            <div class="form-group">
                                <label>超时时长</label>
                                <input class="form-control" placeholder="请输入超时时常" v-validate="'required|numeric'" v-model="addData.time_out" name="time_out"  type="text">
                                <div class="form-height text-danger"><span v-show="errors.first('time_out')">超时时长不能为空且必须为数字</span> </div>
                            </div>
                        </el-col>
                    </el-row> -->
              
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
                <label>secret</label>
                <input class="form-control" placeholder="请输系统签名所需的secret" v-validate="'required'" v-model="addData.secret" name="secret"  type="text">
                <div class="form-height text-danger"><span v-show="errors.first('secret')">secret不能为空</span> </div>
              </div>

              <div class="form-group">
                <label>sso登录后回调子系统的地址(如:http://www.123.com/abc)</label>
                <input class="form-control" placeholder="请输入回调地址" v-model="addData.callbackurl" name="callbackurl"  type="text">
              </div>

              <div class="form-group">
                <!--<label>logo</label>-->
                <input class="form-control" placeholder="logo地址" v-validate="'required'" name="logo" v-model="addData.logo"  readonly>
                <div class="form-height text-danger"> <span v-show="errors.first('logo')">logo地址不能为空</span> </div>
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
                </el-form>
            </div>
            <div class="modal-footer">
              <slot name="footer">
                <button data-dismiss="modal" @click="cancelSubmit" class="btn btn-default">取消</button>
                <button type="button"  @click="submit('addform1')" class="btn m-b-xs w-xs btn-success">保存</button>
              </slot>
            </div>
    </div>
    <!-- /.modal-content -->
  </div>
  <!-- /.modal-dialog -->
</div>
</template>

<script>
export default {
    data(){
        return{
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
            addData: {
                name: "",
                //time_out: "",
                logo: "",
                theme: "",
                style: [],
                ident: "",
                secret:"",
                wechat_status:"1",
                callbackurl:""
            },
            isOpen: false,
            isShow: false,
            isSubmit: false,
            lastKnownBodyStyle: {
                overflow: 'auto'
            },

        }
    },
    
    methods:{
        setModal(){
            this.Init()
        },
        Init() {
            this.first_submit = false;
            this.second_submit = false;
            this.open();
        },
        cancelSubmit(){
            this.close()
        },
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
        submit(){
            let str = "";
        this.addData.style.forEach((item, index) => {
            str += item + " ";
        });
        this.$http.post("/sso/sys/manage", {
            name: this.addData.name,
            callbackurl: this.addData.callbackurl,
            //time_out: this.addData.time_out,
            logo: this.addData.logo,
            style: str,
            theme: this.addData.theme,
            ident: this.addData.ident,
            secret:this.addData.secret,
            wechat_status: this.addData.wechat_status,
        })
            .then(res => {
            this.resetForm()
            this.isSubmit = true
            this.isOpen = false
            this.$emit('refresh-data')
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

        open() {
            if (this.isShow) {
                return
            }
            console.log("编号",this.isOpen)
            this.isShow = true
            this.$nextTick(() => {
                this.isOpen = true
                this.$refs.modal.focus()
                this.lastKnownBodyStyle.overflow = document.body.style.overflow
                document.body.style.overflow = 'hidden'
            })
        },
        close() {
            if(!this.isSubmit){
                this.$confirm("是否退出添加系统 ？", '提示', {
                    cancelButtonText: '取消',
                    confirmButtonText: '确定',
                    type: 'warning'
                }).then(()=>{
                    this.isOpen = false
                    this.$nextTick(() => {
                        setTimeout(() => {
                        this.isShow = false
                        document.body.style.overflow = this.lastKnownBodyStyle.overflow
                        this.resetForm()
                        this.closeModal()
                        this.$emit('refresh-data')
                        }, 100)
                    })                
                }).catch(err=>{
                    console.log("err messagebox",err)
                })
            }else{
                    this.isOpen = false
                    this.$nextTick(() => {
                        setTimeout(() => {
                        this.isShow = false
                        document.body.style.overflow = this.lastKnownBodyStyle.overflow
                        this.closeModal()
                    }, 500)
                    })
            }

        },
        closeModal() {
                this.isSubmit = false
        },
        resetForm() {
            if (this.$refs["addform1"] != undefined) {
                this.$refs["addform1"].resetFields();
            }
           
        },
    }
    }
</script>


