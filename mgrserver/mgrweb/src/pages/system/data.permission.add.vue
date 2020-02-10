<template>
  <div ref="modal" class="modal fade background-darken" tabindex="-1" role="dialog" :class="{in:isOpen,show:isShow}"  @keyup.esc="close()">
    <div class="modal-dialog" style="width:600px;" role="document">
        <div class="modal-content">
            <div class="modal-header">
              <button type="button" class="close" data-dismiss="modal" aria-label="Close" @click="close()">&times;</button>
                <h4 class="modal-title" align="center">新增/修改系统</h4>
            </div>
            <div class="modal-body" >
                <el-form :model="addData"  :inline="true"  ref="addform1" label-position="right" label-width="100px">
                    <el-row :span="24">
                        <el-col :span="12">
                            <div class="form-group">
                            <label>系统名称</label>
                            <input class="form-control" placeholder="请输入系统名称"  name="name" v-model="addData.name"  type="text">
                            <div class="form-height text-danger"><span v-show="errors.first('name')">系统名称不能为空</span></div>
                        </div>
                        </el-col>
                        <el-col :span="12">
                            <div class="form-group" style="margin-left:10px;">
                                <label>系统英文名称</label>
                                <input class="form-control" placeholder="请输入系统英文名称" name="name-e" v-model="addData.ident"  type="text">
                                <div class="form-height text-danger"><span v-show="errors.first('name-e')">系统英文名称不能为空且为字母</span></div>
                            </div>
                        </el-col>
                    </el-row>
                    <div class="form-group">
                        <label>secret</label>
                        <input class="form-control" placeholder="请输系统签名所需的secret" v-model="addData.secret" name="secret"  type="text">
                        <div class="form-height text-danger"><span v-show="errors.first('secret')">secret不能为空</span> </div>
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
import {trimError} from '@/services/util.js'
export default {
    data(){
        return{
            addData: {
                name: "",
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

    computed:{
    //   headers(){
    //     return {
    //       "__jwt__": localStorage.getItem("__jwt__")
    //     }
    //   }
    },
    
    methods:{
        setModal(){
            this.Init()
        },
        Init() {
            this.open();
        },
        cancelSubmit(){
            this.close()
        },
        submit(){
            let str = "";
            this.addData.style.forEach((item, index) => {
              str += item + " ";
            });
            var msg = this.checkBeforeSave();
            if (msg) {
              this.$notify({
                title: '失败',
                message: msg,
                type: 'error',
                offset: 50,
                duration:2000,
                });
              return;
            }
            this.$http.post("/system/info/add", {
                name: this.addData.name,
                callbackurl: this.addData.callbackurl,
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
            this.isShow = false;
            this.$emit('refresh-data')
            this.$notify({
              title: '成功',
              message: '添加成功',
              type: 'success',
              offset: 50,
              duration:2000,
            });
            })
            .catch(err => {
              console.log(err);

            if (err.response.status == 911) {
                this.$notify({
                title: '失败',
                message: "系统名称或英文名称已存在",
                type: 'error',
                offset: 50,
                duration:2000,
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

        open() {
            if (this.isShow) {
                return
            }
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
                this.initAddData()
        },
        resetForm() {
            if (this.$refs["addform1"] != undefined) {
                this.$refs["addform1"].resetFields();
            }
           
        },
        checkBeforeSave() {
          if (!this.addData.name) {
            return "系统名称不能为空";
          }
          if(!this.addData.ident) {
            return "系统英文名称不能为空";
          }
          if(!this.addData.secret) {
            return "secret不能为空";
          }
          if(!this.addData.theme) {
            return "请选择主题样式";
          }
          if(!this.addData.style.length) {
            return "请选择页面布局样式";
          }
          return ""
        },
        initAddData(){
          this.addData.name = "";
          this.addData.logo = "";
          this.addData.theme = "";
          this.addData.style = [];
          this.addData.ident = "";
          this.addData.secret = "";
          this.addData.wechat_status = "1";
          this.addData.callbackurl = "";
        }
      }
    }
</script>