<template>
  <SideMenu></SideMenu>
  <HeaderView></HeaderView>
  <div class="hyper-group-main-container">
    <div class="btn-wrapper">
      <button @click="calculate_hyper_group">生成超群网络</button>
    </div>
    <div class="group-list">
      <div class="group-card" v-for="group in this.$store.state.groups" :key="group.id">
        <div class="card-head">
          <div class="imgBox">
            <img v-if="group.Status === '3'" src="../assets/global.png"/>
            <img v-if="group.Status === '2'" src="../assets/positive.png"/>
            <img v-if="group.Status === '1'" src="../assets/passive.png"/>
          </div>
          <div class="basic-info">
            <h1>{{ group.GroupName }}</h1>
            <h2>{{ group.GroupID }}</h2>
          </div>
        </div>
        <div class="card-body">
          <p>成员数: {{ group.MemberNum }}/{{ group.MaxMemberNum }}</p>
          <p>总消息数: {{ group.SumMsgNum }}  |  平均消息长度: {{ group.AvgMsgLen.toFixed(2) }}<input type="checkbox"></p>
        </div>
      </div>
    </div>

    <div class="graph-wrapper">
      <div id="relation-chart" style="width: 35%; height: 400px; margin: 20px"></div>
      <div id="scatter-chart" style="width: 35%; height: 400px;margin: 20px"></div>
    </div>

  </div>
  <div class="div-line"></div>
  <div class="about">
    <h1><i class="fa-solid fa-scroll"></i>&nbsp;超群网络说明</h1>
    <p>&nbsp;&nbsp;&nbsp;&nbsp;点的提示值为群聊的人数，点的大小由其决定</p>
    <p>&nbsp;&nbsp;&nbsp;&nbsp;边的提示值为重合人数，边的宽度由亲和度决定</p>
    <h1><i class="fa-solid fa-calculator"></i>&nbsp;群聊亲和度计算公式</h1>
    <p>&nbsp;&nbsp;&nbsp;&nbsp;群1人数设为x1，群2人数设为x2，两群重合人数设为y</p>
    <p>&nbsp;&nbsp;&nbsp;&nbsp;Affinity=y^3/(x1*x2)</p>
    <p>&nbsp;&nbsp;&nbsp;&nbsp;</p>
    <p>&nbsp;&nbsp;&nbsp;&nbsp;</p>
  </div>
</template>

<script>
import axios from "axios";
import qs from 'qs'
import * as echarts from 'echarts'
axios.defaults.baseURL = '/api';
const $ = require('jquery')

import HeaderView from "@/components/HeaderView";
import SideMenu from "@/components/SideMenu";
export default {
  name: "HyperGroupView",
  components: {SideMenu, HeaderView},
  methods:{
    calculate_hyper_group() {
      var selected_groups = []
      $('.group-card').each(function () {
        if ($(this).children('.card-body').children('p').children('input[type=checkbox]').prop('checked')){
          selected_groups.push($(this).children('.card-head').children('.basic-info').children('h2').html())
        }
      })
      axios.post('/api/hyper_group/calculate_hyper_group', qs.stringify({
        selected_groups: selected_groups.join()
      }), {
        headers: {
          'Authorization': this.$store.state.user.token
        }
      }).then(response => {
        this.$store.state.hyper_group_points = response.data.data.hyper_group_points
        this.$store.state.hyper_group_edges = response.data.data.hyper_group_edges
        this.$store.state.max_group = response.data.data.max_group
        this.$store.state.min_group = response.data.data.min_group
        this.$store.state.max_edge = response.data.data.max_edge
        this.$store.state.min_edge = response.data.data.min_edge
        this.$store.state.max_gcd = response.data.data.max_gcd
        this.$store.state.min_gcd = response.data.data.min_gcd
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
  },
  mounted() {
    // 生成超群网络关系图
    // 点信息
    let relation_data = []
    var category
    var symbolSize
    var group_diff = this.$store.state.max_group-this.$store.state.min_group
    for (let i=0;i<this.$store.state.hyper_group_points.length;i++) {
      if (this.$store.state.hyper_group_points[i].MemberNum-this.$store.state.min_group<=Math.floor(group_diff*0.25)) {
        category = 3
        symbolSize = 30
      }else if (this.$store.state.hyper_group_points[i].MemberNum-this.$store.state.min_group<=Math.floor(group_diff*0.5)) {
        category = 2
        symbolSize = 40
      }else if (this.$store.state.hyper_group_points[i].MemberNum-this.$store.state.min_group<=Math.floor(group_diff*0.75)) {
        category = 1
        symbolSize = 50
      }else {
        category = 0
        symbolSize = 60
      }
      relation_data.push({
        name: this.$store.state.hyper_group_points[i].GroupName,
        des: this.$store.state.hyper_group_points[i].GroupName+'['+this.$store.state.hyper_group_points[i].MemberNum+']'+'('+this.$store.state.hyper_group_points[i].GroupID+')',
        category: category,
        symbolSize: symbolSize,
      })
    }
    // 边信息
    let relation_links = []
    var width
    var edge_diff = this.$store.state.max_edge-this.$store.state.min_edge
    for (let i=0;i<this.$store.state.hyper_group_edges.length;i++) {
      if (this.$store.state.hyper_group_edges[i].Affinity<=Math.floor(edge_diff*0.25)) {
        width = 2
      }else if (this.$store.state.hyper_group_edges[i].Affinity<=Math.floor(edge_diff*0.5)) {
        width = 5
      }else if (this.$store.state.hyper_group_edges[i].Affinity<=Math.floor(edge_diff*0.75)) {
        width = 8
      }else {
        width = 12
      }
      relation_links.push({
        name: this.$store.state.hyper_group_edges[i].GroupName1+' - '+this.$store.state.hyper_group_edges[i].GroupName2,
        source: this.$store.state.hyper_group_edges[i].GroupName1,
        target: this.$store.state.hyper_group_edges[i].GroupName2,
        des: this.$store.state.hyper_group_edges[i].GroupName1+': '+this.$store.state.hyper_group_edges[i].Value1.toFixed(4)+'->|'+this.$store.state.hyper_group_edges[i].Gcd+'|<-'+this.$store.state.hyper_group_edges[i].Value2.toFixed(4)+' :'+this.$store.state.hyper_group_edges[i].GroupName2+'('+this.$store.state.hyper_group_edges[i].Affinity.toFixed(4)+')',
        value: this.$store.state.hyper_group_edges[i].Gcd,
        lineStyle: {
          width: width,
        }
      })
    }
    // option
    let relation_chart = echarts.init(document.getElementById("relation-chart"));
    let categories = [];
    for (let i = 0; i <= 3; i++) {
      categories[i] = {
        name: '类别' + i
      };
    }
    let relation_option = {
      // 图的标题
      title: {
        text: '超群网络',
        subtext: '(最大(小)点: '+this.$store.state.max_group+'/'+this.$store.state.min_group+' | 最大(小)边: '+this.$store.state.max_gcd+'/'+this.$store.state.min_gcd+')',
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

    // 超群散点图
    let scatter_chart = echarts.init(document.getElementById("scatter-chart"));
    // 散点数据
    let scatter_points = []
    for (let i=0;i<this.$store.state.hyper_group_points.length;i++) {
      scatter_points.push([
        this.$store.state.hyper_group_points[i].MemberNum,
        this.$store.state.hyper_group_points[i].SumMsgNum,
        this.$store.state.hyper_group_points[i].GroupName+'('+this.$store.state.hyper_group_points[i].MemberNum+','+this.$store.state.hyper_group_points[i].SumMsgNum+')',
      ])
    }
    // option
    let scatter_option = {
      title: {
        text: '超群发言频率与人数散点图',
        left: 'center'
      },
      xAxis: {
        name: '群聊人数',
      },
      yAxis: {
        name: '群聊发言频率',
      },
      tooltip: {
        formatter: function (x) {
          return x.data[2];
        }
      },
      series: [
        {
          symbolSize: 15,
          data: scatter_points,
          type: 'scatter'
        }
      ]
    };
    scatter_chart.setOption(scatter_option)
  }
}
</script>

<style scoped>
.group-list{
  display: flex;
  flex-wrap: wrap;
  justify-content: space-around;
}

.graph-wrapper{
  display: flex;
  justify-content: space-around;
  align-items: center;
}

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


input[type="checkbox"]{
  position: relative;
  left: 20px;
  width: 24px;
  height: 24px;
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