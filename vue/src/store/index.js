import { createStore } from 'vuex'

export default createStore({
  state: {
    // token
    token: '',
    // GroupView
    status_option: '',
    group: {},
    graphs: {},
    users: {},
    user_count: '',
    times: {},
    keywords: {},
    keyword_count: '',
    premessages: {},
    // 主界面
    groups:{},
  },
  getters: {
  },
  mutations: {
    time_regular(time) {
      return time
    }
  },
  actions: {
  },
  modules: {
  }
})
