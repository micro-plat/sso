import Vue from 'vue'
import Vuex from 'vuex'
Vue.use(Vuex)


export default new Vuex.Store({
  state: {
    selectMenu: [],
    logincode: "",
  },
  getters: {
    getSelectMenu: state => {
      return state.selectMenu;
    },
    getLoginCode: state=>{
      return state.logincode;
    }
  }
})
