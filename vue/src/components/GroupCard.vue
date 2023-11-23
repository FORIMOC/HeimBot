<template>
  <div class="group-card">
    <div class="card-head">
      <div class="imgBox">
        <img v-if="group_config.Status === '3'" src="../assets/global.png"/>
        <img v-if="group_config.Status === '2'" src="../assets/positive.png"/>
        <img v-if="group_config.Status === '1'" src="../assets/passive.png"/>
      </div>
      <div class="basic-info">
        <h1>{{ group_config.GroupName }}</h1>
        <h2>{{ group_config.GroupID }}</h2>
      </div>
    </div>
    <div class="card-body">
      <p>成员数: {{ group_config.MemberNum }}/{{ group_config.MaxMemberNum }}</p>
      <p>总消息数: {{ group_config.SumMsgNum }}  |  平均消息长度: {{ group_config.AvgMsgLen.toFixed(2) }}</p>
      <div class="btn-wrapper">
        <button @click="goto_group">查看详情</button>
      </div>
    </div>
  </div>
</template>

<script>
export default {
  name: "GroupCard",
  props: {
    group_config: Object,
  },
  methods: {
    goto_group() {
      this.$router.push('/group/'+this.group_config.GroupID)
    }
  }
}
</script>

<style scoped>
.group-card{
  margin: 20px;
  width: 350px;
  border-radius: 12px;
  border: 1px solid rgba(0, 0, 0, 0.1);
  box-shadow: 0 10px 20px rgb(0 0 0 / 20%);
  cursor: default;
}

.card-head{
  display: flex;
  width: 350px;
  height: 80px;
}

.imgBox{
  padding: 5px;
  position: relative;
  width: 80px;
  min-width: 80px;
  height: 80px;
  min-height: 80px;
  display: flex;
  top: -10px;
  left: -10px;
  border-radius: 12px;
  border: 1px solid rgba(0, 0, 0, 0.1);
  box-shadow: 0 10px 20px rgb(0 0 0 / 20%);
  background: white;
}

.imgBox img{
  width: 100%;
  height: 100%;
}

.basic-info{
  display: flex;
  flex-direction: column;
  justify-content: center;
  align-items: center;
  overflow: hidden;
  width: 100%;
}

.basic-info h1{
  font-family: Rubik, sans-serif;
  font-size: 1.5em;
  font-weight: 200;
}

.basic-info h2{
  font-family: Rubik, sans-serif;
  font-size: 1.2em;
  font-weight: 200;
}

.card-body{
}

.card-body p{
  margin: 10px;
  font-family: Rubik, sans-serif;
  font-weight: 200;
}

.btn-wrapper{
  display: flex;
  justify-content: center;
}

.btn-wrapper button{
  border: none;
  border-radius: 6px;
  font-family: Rubik, sans-serif;
  font-size: 1.2em;
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