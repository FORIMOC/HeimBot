<template>
  <HeaderView></HeaderView>
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
    <p>&nbsp;&nbsp;&nbsp;&nbsp;消耗基本的监听服务器资源，以及无上限的数据库资源</p>
    <h1><i class="fa-solid fa-hard-drive"></i>&nbsp;控局监听(2)</h1>
    <p>&nbsp;&nbsp;&nbsp;&nbsp;较积极的监听模式，能够把控群聊的大体局势</p>
    <p>&nbsp;&nbsp;&nbsp;&nbsp;消耗基本的监听服务器资源，以及有上限的数据库资源</p>
    <h1><i class="fa-solid fa-floppy-disk"></i>&nbsp;摘要监听(1)</h1>
    <p>&nbsp;&nbsp;&nbsp;&nbsp;消极监听模式，基本只将捕获到的关键词进行报警</p>
    <p>&nbsp;&nbsp;&nbsp;&nbsp;只消耗基本的监听服务器资源，不启用数据库</p>
  </div>
</template>

<script>
import axios from "axios";
axios.defaults.baseURL = '/api';

import global from '/src/assets/global.png'
import positive from '/src/assets/positive.png'
import passive from '/src/assets/passive.png'
import GroupCard from "@/components/GroupCard";
import HeaderView from "@/components/HeaderView";
import qs from "qs";
export default {
  components: {HeaderView, GroupCard},
  data() {
    return {
      global: global,
      positive: positive,
      passive: passive,
    }
  },
  created() {
    // 登录验证
    axios
        .post('/api/auth/login', qs.stringify({
          token: this.$store.state.token
        })).then(() => {
      // 获取群聊列表
      axios
          .post("/api/group/get_group_list")
          .then(response => {
            this.$store.state.groups = response.data.data.groups
          }).catch(error => {
            alert(error.response.data.msg)
          })

        }).catch(() => {
          this.$router.push('/login')
        })
  }
}
</script>

<style scoped>
.main{
  display: flex;
  flex-wrap: wrap;
  justify-content: space-around;
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
