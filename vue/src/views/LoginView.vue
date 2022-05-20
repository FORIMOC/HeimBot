<template>
  <div class="login-wrapper">
    <img :src="xtrullor" width="400" height="360">
    <h1>Your Vision | Their Future</h1>
    <input id="login-input" v-model="this.$store.state.token" type="text">
    <div class="btn-wrapper">
      <button @click="login">Join Us</button>
    </div>
  </div>
</template>

<script>
import axios from "axios";
import qs from 'qs'
axios.defaults.baseURL = '/api';

import xtrullor from '/src/assets/xtrullor.jpg'
export default {
  name: "LoginView",
  data() {
    return {
      xtrullor: xtrullor
    }
  },
  methods: {
    login() {
      const input = document.getElementById('login-input')
      axios.post('/api/auth/login', qs.stringify({
        token: this.$store.state.token
      })).then(() => {
        this.$router.push('/')
      }).catch(() => {
        input.style = "color: #d63031;"
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
  margin: 50px;
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