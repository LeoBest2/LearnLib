import { createStore } from 'vuex'

export default createStore({
  state: {
    username: ''
  },
  mutations: {
    login(state, value){
      state.username = value
    },
    logoff(state){
      state.username = ''
    }
  },
  actions: {
  },
  modules: {
  }
})
