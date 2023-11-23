<template>
  <SideMenu></SideMenu>
  <HeaderView></HeaderView>
  <div class="key-user-main-container">
    <div class="graph-wrapper">
      <div class="info-box">
        <div class="status-box">
          <div class="img-box">
            <img v-if="this.$store.state.key_user.Status === '1'" :src="positive"/>
            <img v-if="this.$store.state.key_user.Status === '0'" :src="passive"/>
          </div>
          <h1 v-if="this.$store.state.key_user.Status === '1'">监听启用</h1>
          <h1 v-if="this.$store.state.key_user.Status === '0'">监听关闭</h1>
          <select v-model="this.$store.state.key_user_status_option">
            <option>0</option>
            <option>1</option>
          </select>
          <div class="btn-wrapper">
            <button @click="update_status">更改状态</button>
          </div>
          <div class="btn-wrapper">
            <button id="edit_info">编辑信息</button>
          </div>
        </div>
        <div class="word-box">
          <h1>{{ this.$store.state.key_user.Username }}</h1>
          <h2>{{ this.$store.state.key_user.UserID }}</h2>
          <p>&nbsp;</p>
          <p>总发言数: {{ this.$store.state.key_user.SumMsgNum }}</p>
          <p>&nbsp;</p>
          <p>平均发言长度: {{ this.$store.state.key_user.AvgMsgLen.toFixed(2) }}</p>
          <p>&nbsp;</p>
          <p>{{ this.$store.state.key_user.Remark }}</p>
        </div>
      </div>
      <div id="freq_chart" v-if="this.$store.state.key_user.Status === '1'" style="width: 30%; height: 400px"></div>
    </div>

    <div v-if="this.$store.state.key_user.Status === '1'" class="graph-wrapper">
      <div id="time_chart" style="width: 1400px; height: 400px"></div>
    </div>

    <div v-if="this.$store.state.key_user.Status === '1'" class="graph-wrapper">
      <div id="keywords_chart" style="width: 1400px; height: 400px"></div>
    </div>

    <!-- cache_message -->
    <div v-if="this.$store.state.key_user.Status === '1'" class="cache-message-box">
      <p v-for="msg in this.$store.state.key_user_cache_messages" :key="msg.id" style="max-width: 80%">
        <span style="color: #74B9FF;">{{ msg.GroupName }}({{ msg.GroupID }}) | {{ msg.Username }}({{ msg.UserID }})</span>
        <span style="color: #6C5CE7;font-weight: 700;">> </span>
        <span>{{ msg.Content }}</span>
        <span style="position: absolute;right: 20px;color: #7f8c8d;">
          {{ msg.CreatedAt.substr(0,10) + " " + msg.CreatedAt.substr(11,8) }}
        </span>
      </p>
    </div>
    <div style="height: 20px"></div>
  </div>

<!--  编辑信息界面-->
  <div id="loginmask"></div>
  <div class="box">
    <h1>编辑关键用户信息</h1>
    <input type="text" placeholder="用户名" v-model="this.$store.state.input_key_user_username">
    <input type="submit" @click="search_username" value="自动获取用户名">
    <div style="width: 100%;display: flex;justify-content: center">
      <textarea placeholder="备注" v-model="this.$store.state.input_key_user_remark"></textarea>
      <div class="recommended-username">
        <div v-for="(value, key) in this.$store.state.recommended_key_user_username" :key="key">
          <span v-if="value.length>0">{{ key }} | {{ value }}</span>
        </div>
      </div>
    </div>
    <input type="submit" @click="update_info" value="提交">
  </div>
</template>

<script>
import HeaderView from "@/components/HeaderView";
import SideMenu from "@/components/SideMenu";
import axios from "axios";
import qs from "qs";
import positive from "@/assets/positive.png";
import passive from "@/assets/passive.png";
import * as echarts from "echarts";
const $ = require('jquery')
axios.defaults.baseURL = '/api';

$(document).ready(function () {
  $('#edit_info').click(function () {
    $('#loginmask').fadeIn()
    $('.box').slideDown()
  })

  $('#a').click(function () {
    $('#loginmask').fadeOut()
    $('.box').slideUp()
  })

  $('#loginmask').click(function () {
    $('#loginmask').fadeOut()
    $('.box').slideUp()
  })
})

export default {
  name: "KeyUserView",
  components: {SideMenu, HeaderView},
  data() {
    return {
      positive: positive,
      passive: passive,
    }
  },
  methods: {
    update_status() {
      axios.post('/api/key_user/update_key_user_status', qs.stringify({
        user_id: this.$route.params.id,
        status: this.$store.state.key_user_status_option
      }), {
        headers: {
          'Authorization': this.$store.state.user.token
        }
      }).then(() => {
        location.reload()
      }).catch(error => {
        alert(error.response.data.msg)
      })
    },
    search_username() {
      $('.recommended-username').css("display", "block").css("border", "2px solid #eb4d4b")
      axios.post('/api/key_user/search_username', qs.stringify({
        user_id: this.$route.params.id,
      }), {
        headers: {
          'Authorization': this.$store.state.user.token
        }
      }).then(response => {
        this.$store.state.input_key_user_username = response.data.data.nickname
        this.$store.state.recommended_key_user_username = response.data.data.card
      }).catch(error => {
        alert(error.response.data.msg)
      })
    },
    update_info() {
      axios.post('/api/key_user/set_key_user_info', qs.stringify({
        user_id: this.$route.params.id,
        username: this.$store.state.input_key_user_username,
        remark: this.$store.state.input_key_user_remark,
      }), {
        headers: {
          'Authorization': this.$store.state.user.token
        }
      }).then(() => {
        this.$store.commit('del_recommended_username')
        this.$store.commit('del_remark')
        this.$store.commit('del_key_user_username')
        location.reload()
      }).catch(error => {
        alert(error.response.data.msg)
      })
    },
  },
  created() {
    if (typeof this.$store.state.user.token === "undefined") {
      this.$router.push('/login')
    }else {
      // 获取关键用户基本信息
      axios.post("/api/key_user/get_key_user_info", qs.stringify({
        user_id: this.$route.params.id,
      }), {
        headers: {
          'Authorization': this.$store.state.user.token
        }
      }).then(response => {
        this.$store.state.key_user = response.data.data.key_user
        this.$store.state.key_user_status_option = response.data.data.key_user.Status
      }).catch(error => {
        alert(error.response.data.msg)
        this.$router.push("/")
      })

      // 获取关键用户发言群聊
      axios.post("/api/key_user/get_key_user_groups", qs.stringify({
        user_id: this.$route.params.id,
      }), {
        headers: {
          'Authorization': this.$store.state.user.token
        }
      }).then(response => {
        this.$store.state.key_user_groups = response.data.data.key_user_groups
        this.$store.state.key_user_group_count = response.data.data.count
      })

      // 获取关键用户时间分布信息
      axios.post("/api/key_user/get_key_user_times", qs.stringify({
        user_id: this.$route.params.id,
      }), {
        headers: {
          'Authorization': this.$store.state.user.token
        }
      }).then(response => {
        this.$store.state.key_user_times = response.data.data.key_user_times
      })

      // 获取关键用户关键词数据
      axios.post("/api/key_user/get_key_user_keywords", qs.stringify({
        user_id: this.$route.params.id,
      }), {
        headers: {
          'Authorization': this.$store.state.user.token
        }
      }).then(response => {
        this.$store.state.key_user_keywords = response.data.data.keywords
        this.$store.state.key_user_keyword_count = response.data.data.count
      })

      // 获取关键用户缓存消息
      axios.post("/api/key_user/get_key_user_cache_messages", qs.stringify({
        user_id: this.$route.params.id,
      }), {
        headers: {
          'Authorization': this.$store.state.user.token
        }
      }).then(response => {
        this.$store.state.key_user_cache_messages = response.data.data.key_user_cache_messages
      })
    }
  },
  mounted() {
    if (this.$store.state.key_user.Status === "1") {
      // 生成发言频率图
      let freq_data = []
      for (let i=0;i<this.$store.state.key_user_groups.length;i++) {
        freq_data.push({
          value: this.$store.state.key_user_groups[i].Freq,
          name: this.$store.state.key_user_groups[i].GroupName+'('+this.$store.state.key_user_groups[i].GroupID+')'
        })
      }
      // option
      let freq_chart=echarts.init(document.getElementById('freq_chart'));
      let freq_option = {
        title: {
          text: '前'+this.$store.state.key_user_groups.length+'发言群聊',
          subtext: '(共在' + this.$store.state.key_user_group_count + '个群有过发言记录)',
          left: 'center'
        },
        tooltip: {
          trigger: 'item',
        },
        toolbox: {
          show: true,
          feature: {
            mark: {
              show: true
            },
            restore: {
              show: true
            },
            saveAsImage: {
              show: true
            }
          }
        },
        legend:{
          orient: 'vertical',
          left: 'left'
        },
        series:[
          {
            name: '发言频率',
            type: 'pie',
            radius: '50%',
            data: freq_data,
            emphasis: {
              itemStyle: {
                shadowBlur: 10,
                shadowOffsetX: 0,
                shadowColor: 'rgba(0,0,0,0.5)'
              }
            }
          }
        ]
      }
      freq_chart.setOption(freq_option)
      // 生成消息时间分布图
      let time_series = []
      for (let i=0;i<this.$store.state.key_user_times.length;i++) {
        let series = {
          name: this.$store.state.key_user_times[i].Month+'-'+this.$store.state.key_user_times[i].Day,
          type: 'line',
          smooth: true,
          data: []
        }
        series.data.push(this.$store.state.key_user_times[i].Hour1)
        series.data.push(this.$store.state.key_user_times[i].Hour2)
        series.data.push(this.$store.state.key_user_times[i].Hour3)
        series.data.push(this.$store.state.key_user_times[i].Hour4)
        series.data.push(this.$store.state.key_user_times[i].Hour5)
        series.data.push(this.$store.state.key_user_times[i].Hour6)
        series.data.push(this.$store.state.key_user_times[i].Hour7)
        series.data.push(this.$store.state.key_user_times[i].Hour8)
        series.data.push(this.$store.state.key_user_times[i].Hour9)
        series.data.push(this.$store.state.key_user_times[i].Hour10)
        series.data.push(this.$store.state.key_user_times[i].Hour11)
        series.data.push(this.$store.state.key_user_times[i].Hour12)
        time_series.push(series)
      }
      // option
      let time_chart=echarts.init(document.getElementById('time_chart'))
      let time_option = {
        title: {
          text: '消息频率时间分布',
          subtext: '(近'+this.$store.state.key_user_times.length+'天的记录)',
          left: 'center'
        },
        tooltip: {
          trigger: 'axis',
        },
        toolbox: {
          show: true,
          feature: {
            mark: {
              show: true
            },
            restore: {
              show: true
            },
            saveAsImage: {
              show: true
            }
          }
        },
        legend:{
          orient: 'vertical',
          left: 'left'
        },
        xAxis: {
          type: 'category',
          data: [
            '0:00-2:00','2:00-4:00','4:00-6:00',
            '6:00-8:00','8:00-10:00','10:00-12:00',
            '12:00-14:00','14:00-16:00','16:00-18:00',
            '18:00-20:00','20:00-22:00','22:00-24:00'
          ]
        },
        yAxis: {
          type: 'value'
        },
        series: time_series
      }
      time_chart.setOption(time_option)

      // 生成关键词统计柱状图
      // keyword 信息
      let keywords_data = []
      let keywords_freq = []
      for (let i=0;i<this.$store.state.key_user_keywords.length;i++) {
        keywords_data.push(this.$store.state.key_user_keywords[i].Keyword)
        keywords_freq.push(this.$store.state.key_user_keywords[i].Freq)
      }
      // option
      let keywords_chart=echarts.init(document.getElementById('keywords_chart'))
      let keywords_option = {
        title: {
          text: '关键词频率统计',
          subtext: '('+this.$store.state.key_user_keyword_count+'中的'+this.$store.state.key_user_keywords.length+'个关键词)',
          left: 'center'
        },
        tooltip: {
          trigger: 'item',
        },
        toolbox: {
          show: true,
          feature: {
            mark: {
              show: true
            },
            restore: {
              show: true
            },
            saveAsImage: {
              show: true
            }
          }
        },
        xAxis: {
          type: 'category',
          data: keywords_data
        },
        yAxis: {
          type: 'value'
        },
        series: [
          {
            data: keywords_freq,
            type: 'bar'
          }
        ]
      }
      keywords_chart.setOption(keywords_option)
    }
  }
}
</script>

<style scoped>
::-webkit-scrollbar {
  display: none; /* Chrome Safari */
}

#loginmask{
  width: 100%;
  height: 100vh;
  position: fixed;
  left: 0;
  top: 0;
  background-color: #2c3e50;
  z-index: 1;
  opacity: 0.8;
  display: none;
}

.recommended-username{
  background: none;
  display: none;
  margin: 20px auto;
  border: 2px solid #3498db;
  padding: 14px 10px;
  width: 40%;
  height: 300px;
  outline: none;
  color: white;
  border-radius: 24px;
  transition: 0.25s ease-in-out all;
  font-size: 1.5em;
  font-weight: 500;
  font-family: Rubik, sans-serif;
  overflow: auto;
}

.box{
  width: 60%;
  padding: 40px;
  position: absolute;
  top: 50%;
  left: 50%;
  transform: translate(-50%,-50%);
  background: #191919;
  text-align: center;
  border-radius: 24px;
  z-index: 2;
  font-family: Rubik, sans-serif;
  display: none;
  flex-direction: column;
}

.box h1{
  color: white;
  font-weight: 500;
}

.box input[type = "text"]{
  background: none;
  display: block;
  margin: 20px auto;
  text-align: center;
  border: 2px solid #3498db;
  padding: 14px 10px;
  width: 200px;
  outline: none;
  color: white;
  border-radius: 24px;
  transition: 0.25s;
  font-size: 1.5em;
  font-weight: 500;
  font-family: Rubik, sans-serif;
}

.box input[type = "text"]:focus{
  width: 300px;
  border-color: #2ecc71;
}

.box textarea{
  background: none;
  display: block;
  margin: 20px auto;
  text-align: center;
  border: 2px solid #3498db;
  padding: 14px 10px;
  width: 40%;
  height: 300px;
  outline: none;
  color: white;
  border-radius: 24px;
  transition: 0.25s ease-in-out all;
  font-size: 1.5em;
  font-weight: 500;
  font-family: Rubik, sans-serif;
}

.box textarea:focus{
  width: 50%;
  border-color: #2ecc71;
}

.box input[type = "submit"]{
  background: none;
  display: block;
  margin: 20px auto;
  text-align: center;
  border: 2px solid #2ecc71;
  padding: 14px 40px;
  outline: none;
  color: white;
  border-radius: 24px;
  transition: 0.25s;
  cursor: pointer;
  font-size: 1em;
  font-weight: 500;
  font-family: Rubik, sans-serif;
}

.box input[type = "submit"]:hover{
  background: #2ecc71;
}

.info-box{
  width: 30%;
  max-height: 60vh;
  box-shadow: 0 10px 20px rgba(0 0 0 / 20%);
  border: 1px solid rgba(0, 0, 0, 0.1);
  border-radius: 12px;
  margin: 0 30px 30px 0;
  display: flex;
  cursor: default;
}
.status-box{
  width: 200px;
  text-align: center;
  box-shadow: 0 10px 20px rgba(0 0 0 / 20%);
  border: 1px solid rgba(0, 0, 0, 0.1);
  border-radius: 12px;
  height: 430px;
  position: relative;
  left: -40px;
  top: -40px;
  background: white;
}
.status-box select{
  width: 100px;
  text-align: center;
  height: 30px;
  border: 1px solid rgba(0, 0, 0, 0.1);
  border-radius: 4px;
  font-size: 1.2em;
  font-family: Rubik, sans-serif;
  font-weight: 200;
  margin: 10px;
}
.status-box select option{
  width: 100px;
  text-align: center;
  height: 30px;
  border: 1px solid rgba(0, 0, 0, 0.1);
  border-radius: 4px;
  font-size: 1.2em;
  font-family: Rubik, sans-serif;
  font-weight: 200;
  transition: all 0.5s ease-in-out;
}
.img-box{
  padding: 5px;
  width: 200px;
  height: 200px;
}
.img-box img{
  width: 100%;
  height: 100%;
}
.status-box h1{
  font-family: Rubik, sans-serif;
  font-weight: 200;
  font-size: 2.5em;
  color: #192a56;
}
.word-box{
  word-break: break-all;
  margin: 20px 20px 20px 0;
  overflow: auto;
  width: 100%;
}
.word-box h1{
  font-family: Rubik, sans-serif;
  font-weight: 200;
  font-size: 2em;
  text-align: center;
}
.word-box h2{
  font-family: Rubik, sans-serif;
  font-weight: 200;
  font-size: 1.5em;
  text-align: center;
}
.word-box p{
  font-family: Rubik, sans-serif;
  font-weight: 200;
  font-size: 1.2em;
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

.graph-wrapper{
  display: flex;
  justify-content: center;
  align-items: center;
}
.cache-message-box{
  border-radius: 12px;
  padding: 24px;
  margin: 24px;
  box-shadow: 0 10px 20px rgba(0 0 0 / 20%);
  border: 1px solid rgba(0, 0, 0, 0.1);
  word-break: break-all;
  width: 70%;
  left: 15%;
  position: relative;
  cursor: default;
}
.cache-message-box p{
  font-size: 1.2em;
  font-weight: 200;
  font-family: Rubik, sans-serif;
}
.cache-message-box p span:nth-child(3){
  display: inline-flex;
  width: 65%;
}
</style>