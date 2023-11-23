<template>
  <div class="side-menu">
    <div class="toggle-btn fas fa-chevron-right"></div>
    <a href="/">
      <i class="fa-solid fa-comments"></i>
      群聊列表
    </a>
    <a v-if="this.$store.state.user.token && this.$store.state.user.level !== 2" href="/keywords">
      <i class="fa-solid fa-key"></i>
      关键词列表
    </a>
    <a v-if="this.$store.state.user.token && this.$store.state.user.level !== 2" href="/key_users">
      <i class="fa-solid fa-user-secret"></i>
      关键用户列表
    </a>
    <a href="/hyper_group">
      <i class="fa-solid fa-circle-nodes"></i>
      超群视角
    </a>
    <a href="/fraud_detect">
      <i class="fa-solid fa-magnifying-glass-chart"></i>
      反诈监控
    </a>

    <div class="line"></div>

    <!--    未登录-->
    <a v-if="!this.$store.state.user.token" href="/login">
      <i class="fa-solid fa-right-to-bracket"></i>
      登录
    </a>
    <a v-if="!this.$store.state.user.token" href="/register">
      <i class="fa-solid fa-user-plus"></i>
      注册
    </a>

    <!-- 已登录 -->

    <a v-if="this.$store.state.user.token" href="#">
      <i class="fa-solid fa-user"></i>
      {{ this.$store.state.user.userID }}
    </a>
    <a v-if="this.$store.state.user.token && this.$store.state.user.level !== 2" href="#">
      <i class="fa-solid fa-users-gear"></i>
      <span v-if="this.$store.state.user.level === 0">&nbsp;全权</span>
      <span v-if="this.$store.state.user.level === 1">&nbsp;协助</span>
    </a>
    <a v-if="this.$store.state.user.token" href="#" @click="logout">
      <i class="fa-solid fa-power-off"></i>
      注销
    </a>
  </div>
</template>

<script>
export default {
  name: "SideMenu",
  mounted() {
    const menubtn = document.getElementsByClassName('toggle-btn')[0];
    const menu = document.getElementsByClassName('side-menu')[0];
    menubtn.addEventListener('click', function () {
      menu.classList.toggle('active');
      menubtn.classList.toggle('fa-chevron-left')
      menubtn.classList.toggle('fa-chevron-right')
    })
  },
  methods: {
    logout() {
      this.$store.commit('del_user')
      this.$router.push('/login')
    },
  }
}
</script>

<style scoped>
.side-menu{
  z-index: 100;
  position: fixed;
  bottom: 5%;
  left: -250px;
  width: 250px;
  background-color: #f1f1f1aa;
  transition: 0.3s linear;
}

.side-menu.active{
  left: 0;
}

.side-menu a{
  text-decoration: none;
  display: block;
  padding: 14px 10px;
  color: #333;
  font-size: 24px;
  transition: 0.4s linear;
  overflow: hidden;
}

.side-menu a i{
  width: 40px;
  font-size: 30px;
  text-align: center;
}

.side-menu a:hover{
  background-color: #333;
  color: #f1f1f1;
}

.toggle-btn{
  position: absolute;
  right: -40px;
  top: 0;
  width: 40px;
  height: 60px;
  background: #333;
  color: #f1f1f1;
  line-height: 60px;
  text-align: center;
  font-size: 20px;
  cursor: pointer;
}

.line{
  width: 100%;
  height: 5px;
  background: #333;
}
</style>