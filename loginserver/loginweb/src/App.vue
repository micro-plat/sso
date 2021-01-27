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
      hasLoaded: true
    }
  },
  created(){ 
    this.getWebconfig()
  },
  methods:{
    async getWebconfig(){
     
    }
  }
}
</script>


 