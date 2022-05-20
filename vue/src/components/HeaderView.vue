<template>
  <header>
    <h1>HeimBoard</h1>
    <div class="iconBox">
      <img :src="vision"/>
    </div>
    <h1>MonitorBot</h1>
    <div class="menu-left">
      <div class="btn-wrapper">
        <button @click="goto_index">群聊列表</button>
      </div>
    </div>
    <div class="menu-right">
      <div class="btn-wrapper">
        <button>关键用户列表</button>
      </div>
      <div class="btn-wrapper">
        <button @click="goto_keywords">关键词列表</button>
      </div>
    </div>
  </header>
</template>

<script>
import vision from '/src/assets/vision.png'
export default {
  name: "HeaderView",
  data() {
    return {
      vision: vision
    }
  },
  methods: {
    goto_index() {
      this.$router.push('/')
    },
    goto_keywords() {
      this.$router.push('/keywords')
    }
  }
}
</script>

<style scoped>
header{
  width: 100%;
  height: 220px;
  display: flex;
  justify-content: center;
  align-items: center;
  cursor: default;
}

header h1{
  font-family: Rubik, sans-serif;
  font-size: 4em;
  font-weight: 200;
  color: #192a56;
}
.iconBox{
  width: 200px;
  height: 200px;
}

.iconBox img{
  width: 100%;
  height: 100%;
}

.menu-left{
  position: absolute;
  left: 40px;
  display: flex;
}

.menu-right{
  position: absolute;
  right: 40px;
  display: flex;
}

.btn-wrapper{
  display: flex;
  justify-content: center;
}

.btn-wrapper button{
  border: none;
  border-radius: 6px;
  font-family: Rubik, sans-serif;
  font-size: 1.8em;
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