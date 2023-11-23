import {createStore} from 'vuex'

export default createStore({
  state: {
    // user
    user: {},
    // 群聊列表
    groups:{},
    input_group: '',
    // 群聊
    group_status_option: '',
    group: {},
    group_graphs_users: {},
    group_graphs_links: {},
    group_users: {},
    group_user_count: '',
    group_times: {},
    group_keywords: {},
    group_keyword_count: '',
    group_cache_messages: {},
    // 关键词
    keyword_list:{},
    input_keyword: '',
    /* 关键用户 */
    // 关键用户列表
    key_user_list:{},
    key_user_status_option: '',
    key_user: {},
    key_user_groups: {},
    key_user_group_count: '',
    key_user_times: {},
    key_user_keywords: {},
    key_user_keyword_count: '',
    key_user_cache_messages: {},
    // 新增关键用户ID输入框
    input_key_user: '',
    // 关键用户用户名输入框
    input_key_user_username: '',
    // 自动检索获得的推荐用户名(群备注)
    recommended_key_user_username: {},
    // 关键用户备注输入框
    input_key_user_remark: '',
    // 超群
    hyper_group_points: {},
    hyper_group_edges: {},
    max_group: '',
    min_group: '',
    max_edge: '',
    min_edge: '',
    max_gcd: '',
    min_gcd: '',
    // 诈骗消息日志
    fraud_messages: {},
    fraud_groups: {},
    // 时间序列
    time_series_results: {},
  },
  getters: {
  },
  mutations: {
    set_user(state, user) {
      state.user = user
    },
    del_user(state) {
      state.user = {}
    },
    del_recommended_username(state) {
      state.recommended_key_user_username = {}
    },
    del_remark(state) {
      state.input_key_user_remark = ''
    },
    del_key_user_username(state) {
      state.input_key_user_username = ''
    },
    time_regular(time) {
      return time
    },
  },
  actions: {
    token_missing_check() {
      return new Promise((resolve, reject) => {
        typeof this.$store.state.user.token === "undefined" ? resolve('response') : reject('error')
      })
    }
  },
  modules: {
  }
})
