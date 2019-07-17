<template>
  <div class="main">
      <div class="center">请选择要登入的系统</div>
      <div class="everyone" v-for="(item, index) in systems" :key="index">
          <span class="syslogo">
              <a :href="item.callbackurl">
                <img v-if="item.logo !=''" :src="item.logo" />
                <img v-if="item.logo==''" src="../../assets/logo.png" />
             </a>
          </span>
          <span class="note">{{item.name}}</span>
      </div>    
  </div>
</template>

<script>
   import {JoinUrlParams} from '@/services/common'
  export default {
    name: 'chose',
    data () {
      return {
          systems:[],
          code:""
      }
    },

    mounted(){
      document.title = "选择系统";
      this.code = this.$route.query.code;

      this.searchSystemInfo();
    },

    methods:{
        searchSystemInfo() {
            this.$post("lg/user/system")
            .then(res =>{
                if (res != undefined && res.length > 0) {
                    res.forEach((current, index) =>{
                        if (current.callbackurl) {
                            current.callbackurl = JoinUrlParams(current.callbackurl, {code:this.code});
                        } else {
                            current.callbackurl = "javascript:return false";
                        }
                    })
                }
                this.systems = res;
                console.log(this.systems);
            }).catch(err => {
                this.$router.push({ path: '/login', query: { callback: "", sysid: 0 }});
            });
        }
    }
  }
</script>

<style>
    .everyone {
        border-top: 1px solid green;
        padding-top: 20px;
    }
    .main {
        padding-bottom: 20px;   
        display:grid;
        margin:0 auto;
        width: 900px;
        border: 1px solid black;
    }
    .center {
        text-align: center;
        font-size: 16px;
        margin-bottom: 20px;
    }
    .syslogo {
        width: 100px;
        height: 100px;
        margin: 10px 20px;
    }

    .syslogo img {
        width: 100px;
        height: 100px;
    }

    .note {
        margin-left: 50px;
        line-height: 100px;
    }

</style>
