<template>
  <div id="vapp">
    <router-view v-if="hasLoaded"/>
  </div>
</template>

<script>
export default {
  name: 'App',
  data(){
    return{
      hasLoaded: false
    }
  },
  created(){ 
    this.getWebconfig()
  },
  methods:{
    async getWebconfig(){
      var that = this;
      await this.$env.load(async function(){
        var data = await that.$http.get("/system/webconfig");
        that.hasLoaded = true
        return data;
      });
    }
  }
}
</script>


 