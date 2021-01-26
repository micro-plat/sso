<template>

  <li :class="liClassVal">

    <span :class="spanClassVal" @click='open(model)'></span>
    <span title='选中'  >
         <input type="checkbox" @click.stop="checkNode(model)" :checked="model.checked" >
      </span>
    <a @mouseenter='enterFunc(model)' @mouseleave='leaveFunc(model)' @contextmenu.prevent='cxtmenufunc(model)'>
      <span :class="{loadSyncNode:model.loadNode==1}" v-if='model.loadNode==1'></span>
      <span :class='model.icon' v-else></span>
      <span v-show='ischeck' id="treeDemo_5_check" class="button chk"
            :class='{"checkbox_false_full":!model.ckbool,"checkbox_true_full":model.ckbool}' @click='ckFunc(model)'
            treenode_check=""></span>
      <span :class='enableClass' @click='Func(model)' v-if="model.level_id == 4"
            style="color: #23b7e5">{{model.name}}</span>
      <span :class='enableClass' @click='Func(model)' v-if="model.level_id < 4">{{model.name}}</span>
      <!--新增-->
      <span v-show='model.hover' v-if="model.parent < 4 || model.level_id <= 3" title='新增' class="button add"
            @click="addNode(model)"></span>
      <!--删除-->
      <span v-show='model.hover' title='删除' class="button remove" @click="delNode(model)"></span>
      <!--启用禁用-->
      <span v-show='model.hover' v-if="model.enable == 0" title='启用' class="button enable"
            @click="enable(model)"></span>
      <span v-show='model.hover' v-if="model.enable == 1" title='禁用' class="button disable"
            @click="enable(model)"></span>
      <!--上移-->
      <span v-show='model.hover' title='上移' class="button up" @click="upNode(model)"></span>
      <!--下移-->
      <span v-show='model.hover' title='下移' class="button down" @click="downNode(model)"></span>

      <span v-show='model.hover' v-if='model.parent == 0' title='添加顶级节点' class="button add_line"
            @click="addTopNode(model)"></span>
    </a>

    <ul v-show='model.isFolder'>
      <ztree-item v-for="(item,i) in model.children" :key='i' :callback='callback' :checkfunc="checkfunc" :expandfunc='expandfunc'
                  :cxtmenufunc='cxtmenufunc' :model.sync="item" :num.sync='i' root='1'
                  :nodes.sync='model.children.length' :ischeck='ischeck' :trees.sync='trees'></ztree-item>
    </ul>
  </li>
</template>

<script>
  export default {
    name: 'ztreeItem',
    data() {
      return {
        parentNodeModel: null,
        checkedNames: []
      }
    },
    props: {
      model: {
        type: Object,
        twoWay: true
      },
      num: {
        type: Number,
        twoWay: true
      },
      nodes: {
        type: Number,
        twoWay: true,
        default: 0
      },
      trees: {
        type: Array,
        twoWay: true,
        default: []
      },
      root: {
        type: String,
        twoWay: true
      },
      callback: {
        type: Function
      },
      checkfunc:{
        type:Function
      },
      expandfunc: {
        type: Function
      },
      cxtmenufunc: {
        type: Function
      },

      ischeck: {
        type: Boolean,
        twoWay: true,
        default: false
      }
    },
    update() {
      this.initTreeData();
    },
    mounted() {
    },
    methods: {
      checkNode(m) {
        console.log("checked")
        m.checked = !m.checked;
        if (typeof this.checkfunc == "function") {
          if (m.checked){
                this.checkfunc(m.id);
          }
        }
      },
      Func(m) {
        // 查找点击的子节点
        var recurFunc = (data, list) => {
          data.forEach((i) => {
            if (i.id == m.id) {
              i.clickNode = true;
              if (typeof this.callback == "function") {
                this.callback.call(null, m, list, this.trees);
              }
            } else {
              i.clickNode = false;
            }

            if (i.children) {
              recurFunc(i.children, i);
            }
          })
        }

        recurFunc(this.trees, this.trees);
      },
      ckFunc(m) {
        m.ckbool = !m.ckbool;

        // 查找复选框的所有子节点
        var recurFuncChild = (data) => {
          data.forEach((i) => {
            i.ckbool = m.ckbool;
            if (i.children) recurFuncChild(i.children);
          })
        }
        recurFuncChild(m.children);

        // 查找复选框的所有父节点
        var isFindRootBool = false, parentId = m.parentId;
        var recurFuncParent = (data, list) => {
          data.forEach((i) => {
            if (!isFindRootBool) {
              if (i.id == parentId && parentId > 0) {
                parentId = i.parentId;
                i.ckbool = m.ckbool;
                // 重新查找
                recurFuncParent(this.trees, this.trees);
              } else if (i.id == m.id && i.parentId == 0) {
                i.ckbool = m.ckbool;
                isFindRootBool = true;
              } else {
                recurFuncParent(i.children, i);
              }
            }
          })

        }
        recurFuncParent(this.trees, this.trees);
      },
      getParentNode(m, cb) {
        // 查找点击的子节点
        var recurFunc = (data, list) => {
          data.forEach((i) => {
            if (i.id == m.id) this.parentNodeModel = list;
            if (i.children) {
              (typeof cb == "function") && cb.call(i.children);
              recurFunc(i.children, i);
            }
          })
        }
        recurFunc(this.trees, this.trees);
      },
      open(m) {
        m.isExpand = !m.isExpand;

        if (typeof this.expandfunc == "function" && m.isExpand) {
          if (m.loadNode != 2) {
            this.expandfunc.call(null, m);
          } else {
            m.isFolder = !m.isFolder;
          }
        } else {
          m.isFolder = !m.isFolder;
        }
      },
      enterFunc(m) {
        m.hover = true;
        this.getParentNode(m, null);
      },
      leaveFunc(m) {
        m.hover = false;
      },
      // 新增节点
      addNode(nodeModel) {
        console.log(nodeModel)
        if (nodeModel.level_id >= 4) {
          return false
        }
        if (nodeModel) {
          var _nid = +new Date();
          nodeModel.children.push({
            id: _nid,
            parentId: nodeModel.id,
            parentLevel: nodeModel.level_id,
            name: "",
            path: "-",
            icon: "-",
            clickNode: false,
            isNew: true,
            ckbool: false,
            isCheck: this.ischeck,
            isFolder: false,
            isExpand: false,
            hover: false,
            loadNode: 0,
            children: []
          });
          nodeModel.isFolder = true;
          nodeModel.children.forEach((item, index) => {
            if (item.isNew == true) {
              this.Func(item)
            }
          })
        } else {
          return false
        }
      },
      addTopNode(nodeModel) {
        if (nodeModel) {
          var _nid = +new Date();
          this.parentNodeModel.push({
            id: _nid,
            parentId: 0,
            parentLevel: 0,
            name: "",
            path: "-",
            icon: "-",
            clickNode: false,
            isNew: true,
            ckbool: false,
            isCheck: this.ischeck,
            isFolder: false,
            isExpand: false,
            hover: false,
            loadNode: 0,
            children: []
          });
          nodeModel.isFolder = true;
          this.parentNodeModel.forEach((item, index) => {
            if (item.isNew == true) {
              this.Func(item)
            }
          })
        } else {

          return false
        }
      },
      // 删除节点
      delNode(nodeModel) {
        if (nodeModel) {
          this.$http.post("/sys/index/funcdel", {id: nodeModel.id})
            .then(res => {
              if (this.parentNodeModel.hasOwnProperty("children")) {
                this.parentNodeModel.children.splice(this.parentNodeModel.children.indexOf(nodeModel), 1);
              } else if (this.parentNodeModel instanceof Array) {
                // 第一级根节点处理
                this.parentNodeModel.splice(this.parentNodeModel.indexOf(nodeModel), 1);
              }
              nodeModel = null;
              this.$notify({
                title: '成功',
                message: '删除节点成功',
                type: 'success',
                offset: 50
              });
            })
            .catch(err => {
              if (err.response.status == 403) {
                this.$notify({
                  title: '错误',
                  message: '登录超时,请重新登录',
                  type: 'error',
                  offset: 50,
                  duration: 2000,
                  onClose: function () {
                    this.$router.push("/login");
                  }
                });
              } else {
                this.$notify({
                  title: '错误',
                  message: '网络错误,请稍后再试',
                  type: 'error',
                  offset: 50,
                  duration: 2000,
                });
              }
            });

        } else {
          return false;
        }
      },
      enable(nodeModel) {
        if (nodeModel) {
          let status = 1;
          if (nodeModel.enable == 1) {
            status = 0
          } 

          this.$http.post("/sys/index/funcchangestatus", {id: nodeModel.id, status: status})
            .then(res => {
              nodeModel.enable = status;
              this.$notify({
                title: '成功',
                message: '状态修改成功',
                type: 'success',
                offset: 50
              });
            })
            .catch(err => {
              this.$notify({
                title: '错误',
                message: '网络错误,请稍后再试',
                type: 'error',
                offset: 50,
                duration: 2000,
              });
            });
        } else {
          return false;
        }
      },

      upNode(nodeModel) {
        nodeModel.is_up = 2;
        this.$http.post("/sys/index/exchange",nodeModel).then(res =>{
            this.$notify({
              title: '成功',
              message: '上移成功',
            });
            if (typeof this.callback == "function") {
                nodeModel.isSort = true;
                this.callback.call(null,nodeModel);
            }
          })
      },
      downNode(nodeModel) {
        nodeModel.is_up = 1;
        this.$http.post("/sys/index/exchange",nodeModel).then(res =>{
            this.$notify({
              title: '成功',
              message: '下移成功',
            });
            if (typeof this.callback == "function") {
                nodeModel.isSort = true;
                this.callback.call(null,nodeModel);
            }
          })
      },

      //获取系统下面的菜单数据
      getSysMenu(sysId) {
        this.$http.get("/system/func", { id: sysId })
        .then(res => {
          if (res.length != 0) {
            this.model = res;
            console.log(this.model,"数据")
            return;
          }
          this.model.push({
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
      }
    },
    computed: {
      rootClass() {
        var strRootClass = "";

        // 根判断
        if (this.root == "0") {
          this.model.children = this.model.children || [];
          strRootClass =
            this.num == 0 && this.model.children.length == 0
              ? "roots_docu"
              : this.nodes == 1 ||
              (this.num == 0 && this.nodes != this.num + 1)
              ? "root_"
              : this.nodes == this.num + 1 ? "bottom_" : "center_";

          // 子树判断
        } else if (this.root == "1") {
          this.model.children = this.model.children || [];
          strRootClass =
            this.nodes > 1 &&
            this.model.children.length > 0 &&
            this.nodes != this.num + 1
              ? "center_"
              : (this.num == 0 && this.nodes > 1) ||
              this.nodes != this.num + 1
              ? "center_docu"
              : (this.nodes == 1 && this.num != 0) ||
              (this.nodes == this.num + 1 &&
                this.model.children.length > 0)
                ? "bottom_"
                : "bottom_docu";
        }

        return strRootClass;
      },
      // 是否有儿子节点
      isChildren() {
        return this.num + 1 != this.nodes;
      },
      // 展开/收起
      prefixClass() {
        var returnChar = "";
        if (this.rootClass.indexOf("docu") == -1) {
          if (this.model.isFolder) {
            returnChar = "open";
          } else {
            returnChar = "close";
          }
        }

        if (
          this.model.children.length == 0 &&
          this.rootClass.indexOf("docu") == -1
        ) {
          returnChar = "docu";
        }

        return returnChar;
      },
      liClassVal() {
        return "level" + this.num;
      },
      spanClassVal() {
        return (
          "button level" +
          this.num +
          " switch " +
          this.rootClass +
          this.prefixClass
        );
      },
      aClassVal() {
        return this.model.clickNode
          ? "level" + this.num + " curSelectedNode"
          : "level" + this.num;
      },
      enableClass() {
        if (this.model.enable == 1) {
          return "enablefont";
        } else {
          return "disablefont";
        }
      }
    }
  }
</script>

<style>
  @import "font-awesome/css/font-awesome.css";
</style>

