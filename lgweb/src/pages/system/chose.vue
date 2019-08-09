<template>
  <div class="swipercontiner">
    <div class="title">
        请选择要登入的系统
    </div>
    <div class="list">
        <ul v-for="(item, index) in systems" :key="index" @click="goto(item.indexurl)">
            <li class="everyone">
                <span class="icon">
                    <img v-if="item.logo !=''" :src="item.logo" />
                    <img v-if="item.logo==''" src="../../img/iocn_yh.png" >
                </span>
                <span  class="text">{{item.name}}</span>
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
          systems:[]
      }
    },

    mounted(){
      document.title = "选择系统";
      this.searchSystemInfo();
    },

    methods:{
        goto(url) {
            if (url) {
                this.$post("lg/user/code")
                .then(res => {
                    window.location.href = JoinUrlParams(url,{code:res.data});
                    })
                .catch(err =>{
                    this.$router.push({ path: '/login'});
                });
            } 
        },
        searchSystemInfo() {
            this.$post("lg/user/system")
            .then(res =>{
                if (res != undefined && res.length > 0) {
                    res.forEach((current, index) =>{
                        if (current.name.length >= 9) {
                            current.name = current.name.substr(0,9);
                        }
                    })
                }
                this.systems = res;
            }).catch(err => {
                this.$router.push({ path: '/login'});
            });
        }
    }
  }
</script>

<style>
.swipercontiner{ height:100%;}

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
.everyone{
    width: 28%;
    text-align: center;
    display: inline-grid;
    margin: 20px;
    float: left;
}
.list .icon{
	background-color: rgba(21, 171, 160, 0.5);
    padding-top: 20px;
    line-height: 120px;
	border-top-left-radius: 10px;
    border-top-right-radius: 10px;
}

.list .icon img {
  
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