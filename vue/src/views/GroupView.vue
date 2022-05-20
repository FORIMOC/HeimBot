<template>
  <HeaderView></HeaderView>
  <div class="group-main-container">
    <div class="graph_wrapper">
      <div class="info-box">
        <div class="status-box">
          <div class="img-box">
            <img v-if="this.$store.state.group.Status === '3'" :src="global"/>
            <img v-if="this.$store.state.group.Status === '2'" :src="positive"/>
            <img v-if="this.$store.state.group.Status === '1'" :src="passive"/>
          </div>
          <h1 v-if="this.$store.state.group.Status === '3'">全局监听</h1>
          <h1 v-if="this.$store.state.group.Status === '2'">控局监听</h1>
          <h1 v-if="this.$store.state.group.Status === '1'">摘要监听</h1>
          <select v-model="this.$store.state.status_option">
            <option>1</option>
            <option>2</option>
            <option>3</option>
          </select>
          <div class="btn-wrapper">
            <button @click="update_status">更改设置</button>
          </div>
        </div>
        <div class="word-box">
          <h1>{{ this.$store.state.group.GroupName }}</h1>
          <h2>{{ this.$store.state.group.GroupID }}</h2>
          <p>&nbsp;</p>
          <p>成员数: {{ this.$store.state.group.Member }}/{{ this.$store.state.group.MaxMember }}</p>
          <p>总消息数: {{ this.$store.state.group.SumMsg }}  |  平均消息长度: {{ this.$store.state.group.AvgMsg.toFixed(2) }}</p>
        </div>
      </div>
      <div id="relation_chart" v-if="this.$store.state.group.Status !== '1'" style="width: 30%; height: 400px; margin: 20px"></div>
      <div id="freq_chart" v-if="this.$store.state.group.Status !== '1'" style="width: 30%; height: 400px"></div>
    </div>

    <div v-if="this.$store.state.group.Status === '3'" class="graph_wrapper">
      <div id="time_chart" style="width: 1400px; height: 400px"></div>
    </div>

    <div v-if="this.$store.state.group.Status === '3'" class="graph_wrapper">
      <div id="keywords_chart" style="width: 1400px; height: 400px"></div>
    </div>

    <!--    premessage-->
    <div v-if="this.$store.state.group.Status !== '1'" class="premessage-box">
      <p v-for="msg in this.$store.state.premessages" :key="msg.id">
        <span style="color: #74B9FF;">{{ msg.Username }}({{ msg.UserID }})</span>
        <span style="color: #6C5CE7;font-weight: 700;">> </span>
        <span>{{ msg.Content }}</span>
        <span style="position: absolute;right: 20px;color: #7f8c8d;">
          {{ msg.CreatedAt.substr(0,10) + " " + msg.CreatedAt.substr(11,8) }}
        </span>
      </p>
    </div>
    <div style="height: 20px"></div>
  </div>
</template>

<script>
import axios from "axios";
import qs from 'qs'
import * as echarts from 'echarts'
axios.defaults.baseURL = '/api';

import global from '/src/assets/global.png'
import positive from '/src/assets/positive.png'
import passive from '/src/assets/passive.png'
import HeaderView from "@/components/HeaderView";
export default {
  name: "GroupView",
  components: {HeaderView},
  data() {
    return {
      global: global,
      positive: positive,
      passive: passive,
    }
  },
  methods: {
    update_status() {
      axios.post('/api/group/update_status', qs.stringify({
        groupID:  this.$store.state.group.GroupID,
        status: this.$store.state.status_option
      })).then(() => {
        location.reload()
      }).catch(error => {
        alert(error.data.response)
      })
    }
  },
  created() {
    // 登录验证
    axios.post('/api/auth/login', qs.stringify({
      token: this.$store.state.token
    })).then(() => {

      // 群聊基础信息
      axios.post("/api/group/get_group_info", qs.stringify({
        groupID: this.$route.params.id,
      })).then(response => {
        this.$store.state.group = response.data.data.group
        this.$store.state.status_option = response.data.data.group.Status
      }).catch(() => {
        alert("不存在此群聊!")
        this.$router.push("/")
      })
      if (this.$store.state.group.Status !== '1') {
        // 群聊 graphs 信息
        axios.post('/api/group/get_group_graphs', qs.stringify({
          groupID: this.$route.params.id,
        })).then(response => {
          this.$store.state.graphs = response.data.data.graphs
        })
        // 群聊 users 信息
        axios.post('/api/group/get_users_list', qs.stringify({
          groupID: this.$route.params.id,
        })).then(response => {
          this.$store.state.users = response.data.data.users
          this.$store.state.user_count = response.data.data.count
        })
        // 群聊 premessage 信息
        axios.post("/api/group/get_group_premessages", qs.stringify({
          groupID: this.$route.params.id
        })).then(response => {
          this.$store.state.premessages = response.data.data.premessages
        })
      }
      if (this.$store.state.group.Status === '3') {
        // 群聊 time 信息
        axios.post('/api/group/get_group_time', qs.stringify({
          groupID: this.$route.params.id,
        })).then(response => {
          this.$store.state.times = response.data.data.times
        })
        // 群聊 keywords 信息
        axios.post('/api/group/get_group_keywords', qs.stringify({
          groupID: this.$route.params.id,
        })).then(response => {
          this.$store.state.keywords = response.data.data.keywords
          this.$store.state.keyword_count = response.data.data.count
        })
      }

    }).catch(() => {
      this.$router.push('/login')
    })
  },
  mounted() {
    // 生成社交关系图
    // 最大点和边
    let max_user_value = 0
    let max_link_value = 0
    for (let i=0;i<this.$store.state.users.length;i++) {
      max_user_value=this.$store.state.users[i].Value>max_user_value?this.$store.state.users[i].Value:max_user_value
    }
    for (let i=0;i<this.$store.state.graphs.length;i++) {
      max_link_value=this.$store.state.graphs[i].Value1+this.$store.state.graphs[i].Value2>max_link_value?this.$store.state.graphs[i].Value1+this.$store.state.graphs[i].Value2:max_link_value
    }
    // 点信息
    let relation_data = []
    var category
    var symbolSize
    for (let i=0;i<this.$store.state.user_count;i++) {
      if (this.$store.state.users[i].Value<=Math.floor(max_user_value*0.25)) {
        category = 3
        symbolSize = 30
      }else if (this.$store.state.users[i].Value<=Math.floor(max_user_value*0.5)) {
        category = 2
        symbolSize = 40
      }else if (this.$store.state.users[i].Value<=Math.floor(max_user_value*0.75)) {
        category = 1
        symbolSize = 50
      }else {
        category = 0
        symbolSize = 60
      }
      relation_data.push({
        name: this.$store.state.users[i].Username,
        des: this.$store.state.users[i].Username+'['+this.$store.state.users[i].Value.toFixed(1)+']'+'('+this.$store.state.users[i].UserID+')',
        category: category,
        symbolSize: symbolSize,
      })
    }
    // 边信息
    let relation_links = []
    var width
    for (let i=0;i<this.$store.state.graphs.length;i++) {
      if (this.$store.state.graphs[i].Value1+this.$store.state.graphs[i].Value2<=Math.floor(max_link_value*0.25)) {
        width = 2
      }else if (this.$store.state.graphs[i].Value1+this.$store.state.graphs[i].Value2<=Math.floor(max_link_value*0.5)) {
        width = 5
      }else if (this.$store.state.graphs[i].Value1+this.$store.state.graphs[i].Value2<=Math.floor(max_link_value*0.75)) {
        width = 8
      }else {
        width = 12
      }
      relation_links.push({
        name: this.$store.state.graphs[i].Username1+' - '+this.$store.state.graphs[i].Username2,
        source: this.$store.state.graphs[i].Username1,
        target: this.$store.state.graphs[i].Username2,
        des: this.$store.state.graphs[i].Username1+': '+this.$store.state.graphs[i].Value1.toFixed(1)+'->|<-'+this.$store.state.graphs[i].Value2.toFixed(1)+' :'+this.$store.state.graphs[i].Username2,
        value: (this.$store.state.graphs[i].Value1+this.$store.state.graphs[i].Value2).toFixed(1),
        lineStyle: {
          width: width,
        }
      })
    }
    // option
    let relation_chart = echarts.init(document.getElementById("relation_chart"));
    let categories = [];
    for (let i = 0; i <= 3; i++) {
      categories[i] = {
        name: '类别' + i
      };
    }
    let relation_option = {
      // 图的标题
      title: {
        text: '社交网络',
        subtext: '(最大点: '+max_user_value.toFixed(1)+' | 最大边: '+max_link_value.toFixed(1)+')',
        left: 'center'
      },
      // 提示框的配置
      tooltip: {
        formatter: function (x) {
          return x.data.des;
        }
      },
      // 工具箱
      toolbox: {
        // 显示工具箱
        show: true,
        feature: {
          mark: {
            show: true
          },
          // 还原
          restore: {
            show: true
          },
          // 保存为图片
          saveAsImage: {
            show: true
          }
        }
      },
      legend: [{
        // selectedMode: 'single',
        data: categories.map(function (a) {
          return a.name;
        }),
        orient: 'vertical',
        left: 'left'
      }],
      series: [{
        type: 'graph', // 类型:关系图
        layout: 'force', //图的布局，类型为力导图
        symbolSize: 20, // 调整节点的大小
        roam: true, // 是否开启鼠标缩放和平移漫游。默认不开启。如果只想要开启缩放或者平移,可以设置成 'scale' 或者 'move'。设置成 true 为都开启
        edgeSymbolSize: [2, 10],
        emphasis: {
          scale: true,
          focus: 'adjacency',
        },
        edgeLabel: {
          show: true,
          textStyle: {
            fontSize: 15
          },
          formatter: function (x) {
            return x.data.value;
          }
        },
        force: {
          repulsion: 5000,
          edgeLength: [30, 80]
        },
        draggable: true,
        lineStyle: {
          width: 2
        },
        label: {
          show: true,
          position: 'top'
        },
        data: relation_data,
        links: relation_links,
        categories: categories,
      }]
    }
    relation_chart.setOption(relation_option)

    // 生成发言频率图
    let freq_data = []
    for (let i=0;i<(5>this.$store.state.user_count?this.$store.state.user_count:5);i++) {
      freq_data.push({
        value: this.$store.state.users[i].Freq,
        name: this.$store.state.users[i].Username+'('+this.$store.state.users[i].UserID+')'
      })
    }
    // option
    let freq_chart=echarts.init(document.getElementById('freq_chart'));
    let freq_option = {
      title: {
        text: '前5活跃用户',
        subtext: '(共' + this.$store.state.user_count + '用户有过发言记录)',
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
    for (let i=0;i<this.$store.state.times.length;i++) {
      let series = {
        name: this.$store.state.times[i].Month+'-'+this.$store.state.times[i].Day,
        type: 'line',
        smooth: true,
        data: []
      }
        series.data.push(this.$store.state.times[i].Hour1)
        series.data.push(this.$store.state.times[i].Hour2)
        series.data.push(this.$store.state.times[i].Hour3)
        series.data.push(this.$store.state.times[i].Hour4)
        series.data.push(this.$store.state.times[i].Hour5)
        series.data.push(this.$store.state.times[i].Hour6)
        series.data.push(this.$store.state.times[i].Hour7)
        series.data.push(this.$store.state.times[i].Hour8)
        series.data.push(this.$store.state.times[i].Hour9)
        series.data.push(this.$store.state.times[i].Hour10)
        series.data.push(this.$store.state.times[i].Hour11)
        series.data.push(this.$store.state.times[i].Hour12)
      time_series.push(series)
    }
    // option
    let time_chart=echarts.init(document.getElementById('time_chart'))
    let time_option = {
      title: {
        text: '消息频率时间分布',
        subtext: '(近3天的记录)',
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
    for (let i=0;i<this.$store.state.keywords.length;i++) {
      keywords_data.push(this.$store.state.keywords[i].Keyword)
      keywords_freq.push(this.$store.state.keywords[i].Freq)
    }
    // option
    let keywords_chart=echarts.init(document.getElementById('keywords_chart'))
    let keywords_option = {
      title: {
        text: '关键词频率统计',
        subtext: '(共捕获到'+this.$store.state.keyword_count+'个关键词)',
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
</script>

<style scoped>
.info-box{
  width: 30%;
  height: 300px;
  box-shadow: 0 10px 20px rgba(0 0 0 / 20%);
  border: 1px solid rgba(0, 0, 0, 0.1);
  border-radius: 12px;
  margin-right: 30px;
  display: flex;
  cursor: default;
}
.status-box{
  width: 200px;
  text-align: center;
  box-shadow: 0 10px 20px rgba(0 0 0 / 20%);
  border: 1px solid rgba(0, 0, 0, 0.1);
  border-radius: 12px;
  height: 370px;
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

.graph_wrapper{
  display: flex;
  justify-content: center;
  align-items: center;
}
.premessage-box{
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
.premessage-box p{
  font-size: 1.2em;
  font-weight: 200;
  font-family: Rubik, sans-serif;
}
.premessage-box p span:nth-child(3){
  display: inline-flex;
  width: 65%;
}
</style>