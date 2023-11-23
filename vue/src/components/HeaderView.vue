<template>
  <header>
    <h1>HeimBoard</h1>
    <div class="iconBox">
      <img :src="vision"/>
    </div>
    <h1>SecurityBot</h1>
    <div class="config">
      <p>{{ version }}</p>
      <p>{{ author }}</p>
    </div>
  </header>
</template>

<script>
import {version, author} from '../../heimboard.config'
import vision from '/src/assets/vision.png'
export default {
  name: "HeaderView",
  data() {
    return {
      vision: vision,
      version: version,
      author: author
    }
  },
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
  position: relative;
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
  width: 200px;
  height: 100%;
}

.config{
  position: absolute;
  top: 6%;
  right: 2%;
}

.config p{
  font-family: Rubik, sans-serif;
}
</style>