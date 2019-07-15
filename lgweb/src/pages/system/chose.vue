<template>
  <div class="main">
      <div class="center">请选择要登入的系统</div>
      <div>
          <ul>
            <li class="syslogo" v-for="(item, index) in systems" :key="index">
                <a :href="item.indexurl">
                    <img v-if="item.logo !=''" :src="item.logo" :title="item.name" />
                    <img v-if="item.logo==''" src="../../assets/logo.png" :title="item.name" />
                </a>
            </li>
        </ul>
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
                        current.indexurl = JoinUrlParams(current.indexurl, {code:this.code});
                        current.logo = '';
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
    .main {
        padding-bottom: 20px;   
        display:grid;
        margin:0 auto;
        width: 1500px;
        border: 1px solid black;
    }
    .center {
        text-align: center;
        font-size: 16px;
        margin-bottom: 20px;
    }
    ul,li {
        list-style: none;
        padding: 0px;
        margin: 0px;
    }
    .syslogo {
        float: left;
        width: 100px;
        height: 100px;
        margin: 10px 20px;
    }

    .syslogo img {
        width: 100px;
        height: 100px;
    }

</style>
