<template>
  <div class="swipercontiner">
    <div data-v-21e0ac8e="" id="bg" style="background: rgb(33, 150, 243);">
        <canvas data-v-21e0ac8e="" width="1840" height="454"></canvas>
        <canvas data-v-21e0ac8e="" width="1840" height="454"></canvas> 
        <canvas data-v-21e0ac8e="" width="1840" height="454"></canvas>
    </div>
    <div class="title">
        请选择要登入的系统
    </div>
    <div class="list">
        <ul v-for="(item, index) in systems" :key="index">
            <li class="everyone">
                <span class="icon">
                    <img v-if="item.logo !=''" :src="item.logo" />
                    <img v-if="item.logo==''" src="../../img/iocn_yh.png" >
                </span>
                <span class="text">
                    {{item.name}}
                </span>
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
                        if (current.callbackurl) {
                            current.callbackurl = JoinUrlParams(current.callbackurl, {code:this.code});
                        } else {
                            current.callbackurl = "javascript:return false";
                        }
                        if (current.name.length >= 9) {
                            current.name = current.name.substr(0,9);
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
.swipercontiner{ height:100%;}

body{font-family:"黑体";	background:#f5f5f5; font-size:12px; margin:0;padding:0;}

li{	list-style:none;}

.title{
	font-size: 60px;
    padding: 80px 0;
    text-align: center;
    font-weight: bold;
}
.list{
	width: 900px;
    margin: 0 auto;
}
.everyone{
    width: 45%;
    text-align: center;
    display: inline-grid;
    margin: 20px;
    float: left;
}
.list .icon{
	background-color: #fff;
	padding: 60px 0;
	border-top-left-radius: 10px;
    border-top-right-radius: 10px
}

.list .icon img {
    height: 90px;
    width:90px;
}

.list .text{
	font-size: 30px;
	color: #fff;
    padding: 30px 0;
     background-color: rgba(0,14,13,0.5);
    border-bottom-left-radius: 10px;
    border-bottom-right-radius: 10px
}

#bg {
    position: fixed;
    top: 0;
    left: 0;
    width: 100%;
    height: 100%;
    z-index: -1;
}

#bg canvas {
    position: absolute;
    top: 0;
    left: 0;
    width: 100%;
    height: 100%;
}

</style>