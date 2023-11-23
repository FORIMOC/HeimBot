<template>
  <div class="login-wrapper">
    <img :src="xtrullor" width="400" height="360">
    <h1>Your Vision | Their Future</h1>
    <input id="login-userID" v-model="this.user.userID" type="text" placeholder="UserID">
    <input id="login-password" v-model="this.user.password" type="password" placeholder="Password">
    <div class="btn-wrapper">
      <button @click="register">Register</button>
      <button @click="go_to_login"><i class="fa-solid fa-right-from-bracket"></i>&nbsp;to Login</button>
    </div>
  </div>
</template>

<script>
import axios from "axios";
import qs from 'qs'
axios.defaults.baseURL = '/api';

import xtrullor from '/src/assets/xtrullor.jpg'
export default {
  name: "RegisterView",
  data() {
    return {
      xtrullor: xtrullor,
      user: {
        userID: "",
        password: "",
      }
    }
  },
  methods: {
    go_to_login() {
      this.$router.push('/login')
    },
    register() {
      const input_userID = document.getElementById('login-userID')
      const input_password = document.getElementById('login-password')
      axios.post('/api/auth/register', qs.stringify({
        user_id: this.user.userID,
        password: this.user.password
      })).then(response => {
        // 保存 user 信息
          this.$store.commit('set_user', response.data.data)
          this.$router.push('/login')
      }).catch(error => {
        alert(error.response.data.msg)
          input_userID.style = "color: #d63031;"
          input_password.style = "color: #d63031;"
      })
    }
  },
}
</script>

<style scoped>
.login-wrapper{
  min-height: 100vh;
  display: flex;
  flex-direction: column;
  justify-content: center;
  align-items: center;
  cursor: default;
}
.login-wrapper h1{
  font-family: Rubik, sans-serif;
  font-size: 5em;
  font-weight: 200;
}
.login-wrapper input{
  width: 1000px;
  font-family: Rubik, sans-serif;
  font-size: 2.5em;
  font-weight: 200;
  padding: 5px;
  text-align: center;
  border-radius: 8px;
  border: 1px solid rgba(0, 0, 0, 0.1);
  box-shadow: 0 10px 20px rgb(0 0 0 / 20%);
  margin: 30px;
  transition: all 0.2s ease-in-out;
}

.btn-wrapper{
  display: flex;
  justify-content: center;
}

.btn-wrapper button{
  border: none;
  border-radius: 6px;
  font-family: Rubik, sans-serif;
  font-size: 2.4em;
  font-weight: 200;
  padding: 12px 40px 12px 40px;
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