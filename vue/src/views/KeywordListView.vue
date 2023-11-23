<template>
  <SideMenu></SideMenu>
  <HeaderView></HeaderView>
  <div class="keywords-main-container">
    <div class="keywords-option-wrapper">
      <input v-model="this.$store.state.input_keyword" type="text" placeholder="输入关键词">
      <div class="btn-wrapper">
        <button @click="add_keyword">添加关键词</button>
        <button @click="delete_keyword">删除关键词</button>
        <button @click="regulate_keyword">重整关键词一致性</button>
      </div>
    </div>
    <div class="keywords-list">
      <div v-for="keyword in this.$store.state.keyword_list" :key="keyword.id">{{ keyword.Keyword }}</div>
    </div>
  </div>
</template>

<script>
import axios from "axios";
axios.defaults.baseURL = '/api';

import HeaderView from "@/components/HeaderView";
import qs from "qs";
import SideMenu from "@/components/SideMenu";
export default {
  name: "KeywordView",
  components: {SideMenu, HeaderView},
  methods: {
    add_keyword() {
      axios.post('/api/keyword/add_keyword', qs.stringify({
        keyword:  this.$store.state.input_keyword.trim(),
      }), {
        headers: {
          'Authorization': this.$store.state.user.token
        }
      }).then(() => {
        this.$store.state.input_keyword=''
        location.reload()
      }).catch(error => {
        alert(error.response.data.msg)
      })
    },
    delete_keyword() {
      axios.post('/api/keyword/delete_keyword', qs.stringify({
        keyword:  this.$store.state.input_keyword.trim(),
      }), {
        headers: {
          'Authorization': this.$store.state.user.token
        }
      }).then(() => {
        this.$store.state.input_keyword=''
        location.reload()
      }).catch(error => {
        alert(error.response.data.msg)
      })
    },
    regulate_keyword() {
      axios.post('/api/keyword/regulate_keyword', {}, {
        headers: {
          'Authorization': this.$store.state.user.token
        }
      }).then(() => {
        location.reload()
      }).catch(error => {
        alert(error.response.data.msg)
      })
    }
  },
  created() {
    // 登录验证
    if (typeof this.$store.state.user.token === "undefined") {
      this.$router.push('/login')
    }else {
      // 关键词列表
      axios.post('/api/keyword/list_keyword', {}, {
        headers: {
          'Authorization': this.$store.state.user.token
        }
      }).then(response => {
        this.$store.state.keyword_list=response.data.data.keyword_list
      }).catch(error => {
        alert(error.response.data.msg)
      })
    }
  }
}
</script>

<style scoped>
.keywords-main-container{
  margin: 0 50px 0 50px;
}
.keywords-option-wrapper{
  display: flex;
  justify-content: center;
}
.keywords-option-wrapper input{
  width: 400px;
  font-family: Rubik, sans-serif;
  font-size: 2.5em;
  font-weight: 200;
  padding: 5px;
  text-align: center;
  border-radius: 8px;
  border: 1px solid rgba(0, 0, 0, 0.1);
  box-shadow: 0 10px 20px rgb(0 0 0 / 20%);
  margin: 50px;
  transition: all 0.2s ease-in-out;
}

.keywords-list{
  display: flex;
  flex-wrap: wrap;
  justify-content: space-around;
  align-items: center;
  cursor: default;
}
.keywords-list>div{
  width: 200px;
  text-align: center;
  margin: 50px;
  padding: 10px;
  font-size: 2.5em;
  font-family: Rubik, sans-serif;
  font-weight: 200;
  overflow: hidden;
  word-break: break-all;
}
.keywords-list>div:before{
  content: '';
  position: absolute;
  z-index: -1;
  border: none;
  border-radius: 4px;
  height: 15px;
  transform: translateY(35px);
  background: rgba(79, 192, 141, 0.5);
  width: 170px;
}

.btn-wrapper{
  display: flex;
  justify-content: center;
  align-items: center;
}

.btn-wrapper button{
  border: none;
  border-radius: 6px;
  font-family: Rubik, sans-serif;
  font-size: 1.2em;
  max-height: 40px;
  font-weight: 200;
  padding: 6px 20px 6px 20px;
  margin: 10px;
  background: #c7ecee;
  cursor: pointer;
  position: relative;
  transition: all 0.5s ease-in-out;
}

.btn-wrapper button:hover{
  background: #192a56;
  color: #fff;
}
</style>