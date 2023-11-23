<template>
  <SideMenu></SideMenu>
  <HeaderView></HeaderView>
  <div class="group-option-wrapper">
    <input v-model="this.$store.state.input_group" type="text" placeholder="输入群聊ID">
    <div class="btn-wrapper">
      <button @click="delete_group_data">删除群聊</button>
    </div>
  </div>
  <div class="main">
    <GroupCard
        v-for="group in this.$store.state.groups" :key="group.id"
        :group_config="group"
    ></GroupCard>
  </div>
  <div class="div-line"></div>
  <div class="about">
    <h1><i class="fa-solid fa-database"></i>&nbsp;全局监听(3)</h1>
    <p>&nbsp;&nbsp;&nbsp;&nbsp;获取最多的群聊信息，所有消息、所有用户都将被记录</p>
    <p>&nbsp;&nbsp;&nbsp;&nbsp;消耗基本的服务器资源，以及无上限条数的数据库存储</p>
    <h1><i class="fa-solid fa-hard-drive"></i>&nbsp;控局监听(2)</h1>
    <p>&nbsp;&nbsp;&nbsp;&nbsp;较积极的监听模式，能够记录群聊的大体情况</p>
    <p>&nbsp;&nbsp;&nbsp;&nbsp;消耗基本的服务器资源，以及有上限条数的数据库存储</p>
    <h1><i class="fa-solid fa-floppy-disk"></i>&nbsp;摘要监听(1)</h1>
    <p>&nbsp;&nbsp;&nbsp;&nbsp;消极监听模式，仅将捕获到的关键词/模型输出进行报警</p>
    <p>&nbsp;&nbsp;&nbsp;&nbsp;只消耗基本的服务器资源，并更新群聊的基本信息</p>
  </div>
</template>

<script>
import axios from "axios";
axios.defaults.baseURL = '/api';

import global from '/src/assets/global.png'
import positive from '/src/assets/positive.png'
import passive from '/src/assets/passive.png'
import GroupCard from "@/components/GroupCard";
import SideMenu from "@/components/SideMenu";
import HeaderView from "@/components/HeaderView";
import qs from "qs";
export default {
  components: {HeaderView, SideMenu, GroupCard},
  data() {
    return {
      global: global,
      positive: positive,
      passive: passive,
    }
  },
  methods: {
    delete_group_data() {
      axios.post('/api/group/delete_group_data', qs.stringify({
        group_id:  this.$store.state.input_group,
      }), {
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
      // 获取群聊列表
      axios.post("/api/group/get_group_list", {}, {
        headers: {
          'Authorization': this.$store.state.user.token
        }
      }).then(response => {
            this.$store.state.groups = response.data.data.groups
          }).catch(error => {
        alert(error.response.data.msg)
      })
    }
  }
}
</script>

<style scoped>
.main{
  display: flex;
  flex-wrap: wrap;
  justify-content: space-around;
}

.group-option-wrapper{
  display: flex;
  justify-content: center;
}

.group-option-wrapper input{
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

.div-line{
  width: 90%;
  height: 2px;
  border-radius: 50%;
  background: rgba(0, 0, 0, 0.2);
  left: 5%;
  position: relative;
  margin: 20px;
}

.about{
  margin: 20px;
  display: flex;
  flex-direction: column;
}

.about h1{
  font-size: 2.5em;
  font-weight: 200;
  font-family: Rubik,sans-serif;
  color: #192a56;
}

.about p{
  font-size: 1.5em;
  font-weight: 200;
  font-family: Rubik,sans-serif;
}
</style>
