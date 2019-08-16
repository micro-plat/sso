<template>
  <div class="container">
    <div class="title">
        请选择要登入的系统
    </div>
    <div class="list">
        <ul v-for="(item, index) in systems" :key="index" @click="goto(item.index_url)">
            <li class="item">
                <span class="icons">
                    <img v-if="item.logo !=''" :src="item.logo" />
                    <img v-if="item.logo==''" src="../../img/icon_yh.png" >
                </span>
                <span  class="text">{{item.name | subStr}}</span>
            </li>
        </ul>
    </div>
</div>
</template>

<script>
   import {JoinUrlParams, GetUrlHosts} from '@/services/common'
  export default {
    name: 'choose',
    data () {
      return {
          systems:[]
      }
    },
    mounted(){
      document.title = "选择系统";
      this.searchSystemInfo();
    },
    methods:{
        goto(url) {
            window.location.href = GetUrlHosts(url);
        },
        searchSystemInfo() {
            this.$post("/member/system/get")
            .then(res =>{
                this.systems = res;
            }).catch(err => {
                this.$router.push({ path: '/login'});
            });
        }
    },
    filters:{
        subStr: function(value) {
            if (value && value.length > 9) {  
                value = value.substring(0, 9) + "...";
            }
            return value;
        }
    }
  }
</script>

<style>
.container{ height:100%;}

body{font-family:"黑体";background:#f5f5f5; font-size:12px; margin:0;padding:0;}

li{	list-style:none;}

.title{
	font-size: 34px;
    padding: 80px 0;
    text-align: center;
    font-weight: bold;
    color:#fff;
}
.list{
	width: 1100px;
    margin: 0 auto;
}
.item{
    width: 28%;
    text-align: center;
    display: inline-grid;
    margin: 20px;
    float: left;
}
.list .icons{
	background-color: rgba(21, 171, 160, 0.5);
    padding-top: 20px ;
    line-height: 120px;
	border-top-left-radius: 10px;
    border-top-right-radius: 10px;
}

.list .text{
	font-size: 20px;
	color: #fff;
    padding: 22px 0;
    background-color: rgba(0,14,13,0.5);
    border-bottom-left-radius: 10px;
    border-bottom-right-radius: 10px;
}
.list:hover {
    cursor: pointer;
}
 body{
    background: url(../../img/background.png) ;
     background-size: cover;
} 

</style>