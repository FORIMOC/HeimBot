<template>
  <SideMenu></SideMenu>
  <HeaderView></HeaderView>
  <div class="key-users-main-container">
    <div class="key-users-option-wrapper">
      <input v-model="this.$store.state.input_key_user" type="text" placeholder="输入关键用户ID">
      <div class="btn-wrapper">
        <button @click="add_key_user">添加关键用户</button>
        <button @click="delete_key_user">删除关键用户</button>
      </div>
    </div>
    <div class="key-users-list">
      <KeyUserCard
          v-for="key_user in this.$store.state.key_user_list" :key="key_user.id"
          :key_user_config="key_user"
      ></KeyUserCard>
    </div>
  </div>
</template>

<script>
import axios from "axios";
// import qs from 'qs'
axios.defaults.baseURL = '/api';

import HeaderView from "@/components/HeaderView";
import SideMenu from "@/components/SideMenu";
import qs from "qs";
import KeyUserCard from "@/components/KeyUserCard";
export default {
  name: "KeyUserListView",
  components: {KeyUserCard, SideMenu, HeaderView},
  methods: {
    add_key_user() {
      axios.post('/api/key_user/add_key_user', qs.stringify({
        user_id:  this.$store.state.input_key_user,
      }), {
        headers: {
          'Authorization': this.$store.state.user.token
        }
      }).then(() => {
        this.$store.state.input_key_user=''
        location.reload()
      }).catch(error => {
        alert(error.response.data.msg)
      })
    },
    delete_key_user() {
      axios.post('/api/key_user/delete_key_user', qs.stringify({
        user_id:  this.$store.state.input_key_user,
      }), {
        headers: {
          'Authorization': this.$store.state.user.token
        }
      }).then(() => {
        this.$store.state.input_key_user=''
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
      // 获取关键用户列表
      axios.post("/api/key_user/get_key_user_list", {}, {
        headers: {
          'Authorization': this.$store.state.user.token
        }
      }).then(response => {
        this.$store.state.key_user_list = response.data.data.key_user_list
      }).catch(error => {
        alert(error.response.data.msg)
      })
    }
  }
}
</script>

<style scoped>
.key-users-main-container{
  margin: 0 50px 0 50px;
}
.key-users-option-wrapper{
  display: flex;
  justify-content: center;
}
.key-users-option-wrapper input{
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

.key-users-list{
  display: flex;
  flex-wrap: wrap;
  justify-content: space-around;
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