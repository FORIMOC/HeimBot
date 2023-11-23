<template>
  <SideMenu></SideMenu>
  <HeaderView></HeaderView>
  <div class="fraud-detect-main-container">
    <div class="model-state-box">
      <div class="model-state-pic-box">
        <img :src="passive"/>
      </div>
      <div class="model-state-word-box">
        <h1>神经网络模块认为以下消息涉诈</h1>
        <h2>输出来自: {{ model }}</h2>
      </div>
    </div>
    <!-- fraud_messages -->
    <div class="fraud-message-box">
      <p v-for="msg in this.$store.state.fraud_messages" :key="msg.id"  style="max-width: 80%">
        <span style="color: #74B9FF;">{{ msg.GroupName }}({{ msg.GroupID }}) | {{ msg.Username }}({{ msg.UserID }})</span>
        <span style="color: #6C5CE7;font-weight: 700;">> </span>
        <span v-html="msg.Content" style="display: inline;"></span>
        <span style="position: absolute;right: 20px;color: #7f8c8d;">
          {{ msg.CreatedAt.substr(0,10) + " " + msg.CreatedAt.substr(11,8) }}
        </span>
      </p>
    </div>
    <div class="div-line"></div>
      <div class="related-state-box">
        <div class="related-state-word-box">
          <h1>涉诈群聊引流日志</h1>
          <h2><i class="fa-solid fa-database"></i>&nbsp;来自涉诈群聊数据库的历史记录</h2>
        </div>
      </div>
    <!-- fraud_groups -->
    <div class="fraud-message-box">
      <p v-for="msg in this.$store.state.fraud_groups" :key="msg.id"  style="max-width: 80%">
        <span style="color: #ff7675;font-weight: 700;"><span style="color: #6C5CE7;font-weight: 300;">用户</span>&nbsp;{{ msg.SenderID }}</span>
        <span style="color: #6C5CE7;font-weight: 300;"> 引流 </span>
        <span style="color: #ff7675;font-weight: 700;"><span style="color: #6C5CE7;font-weight: 300;">群聊</span>&nbsp;{{ msg.GroupID }}</span>
        <span style="position: absolute;right: 20px;color: #7f8c8d;">
          {{ msg.CreatedAt.substr(0,10) + " " + msg.CreatedAt.substr(11,8) }}
        </span>
      </p>
    </div>
    <div class="div-line"></div>
    <div class="model-state-box">
      <div class="model-state-pic-box" style="border-radius: 12px;">
        <img :src="series"/>
      </div>
      <div class="model-state-word-box">
        <h1>时间序列模块认为以下用户行为离群</h1>
        <h2>输出来自: {{ time_series.algorithm }}</h2>
      </div>
    </div>
      <div v-for="result in this.$store.state.time_series_results" :key="result.id"  style="max-width: 80%" class="time-series-result-box">
        <div class="time-series-title-box">
          <p>群聊ID: {{ result.GroupID }} | 结束时间: {{ result.CreatedAt.substr(0,10) + " " + result.CreatedAt.substr(11,8) }} | 时长: {{ time_series.period }}</p>
        </div>
        <div class="div-line"></div>
        <div class="time-series-function-box">
          <div class="time-series-function-single">
            <div class="time-series-function-single-pic" style="width: 150px; height: 150px">
              <img :src="func"/>
            </div>
            <div class="time-series-function-content" style="font-size: 1.8em">
              <p><span style="font-weight: 700">ACF自相关系数: {{ result.ACF1 }}</span></p>
              <p>振幅: {{ result.Amplitude1 }}</p>
              <p>频率: {{ result.Freq1 }}</p>
              <p>相位: {{ result.Phase1 }}</p>
              <p>f(x)=<span style="color: #e74c3c">{{ result.Amplitude1 }}</span>sin(2𝝅*<span style="color: #8e44ad;">{{ result.Freq1 }}</span>*x+<span style="color: #16a085">{{ result.Phase1 }}</span>)</p>
            </div>
          </div>
          <div class="time-series-function-single">
            <div class="time-series-function-single-pic">
              <img :src="func"/>
            </div>
            <div class="time-series-function-content">
              <p><span style="font-weight: 700">ACF自相关系数: {{ result.ACF2 }}</span></p>
              <p>振幅: {{ result.Amplitude2 }}</p>
              <p>频率: {{ result.Freq2 }}</p>
              <p>相位: {{ result.Phase2 }}</p>
              <p>f(x)=<span style="color: #e74c3c">{{ result.Amplitude2 }}</span>sin(2𝝅*<span style="color: #8e44ad;">{{ result.Freq2 }}</span>*x+<span style="color: #16a085">{{ result.Phase2 }}</span>)</p>
            </div>
          </div>
          <div class="time-series-function-single">
            <div class="time-series-function-single-pic">
              <img :src="func"/>
            </div>
            <div class="time-series-function-content">
              <p><span style="font-weight: 700">ACF自相关系数: {{ result.ACF3 }}</span></p>
              <p>振幅: {{ result.Amplitude3 }}</p>
              <p>频率: {{ result.Freq3 }}</p>
              <p>相位: {{ result.Phase3 }}</p>
              <p>f(x)=<span style="color: #e74c3c">{{ result.Amplitude3 }}</span>sin(2𝝅*<span style="color: #8e44ad;">{{ result.Freq3 }}</span>*x+<span style="color: #16a085">{{ result.Phase3 }}</span>)</p>
            </div>
          </div>
        </div>
        <div class="div-line"></div>
        <div class="time-series-mse-box">
          <div class="time-series-mse-pic">
            <img :src="gauss"/>
          </div>
          <div class="time-series-mse-word-box">
            <p>MSE均值: {{ result.MSEMean.toFixed(2) }} | MSE方差: {{ result.MSEVariance.toFixed(2) }}</p>
            <p>MSE最大值: {{ result.OutliersMSE.toFixed(2) }} | 对应用户ID: <span class="target-highlight">{{ result.Outliers }}</span></p>
            <p>&nbsp;</p>
            <p>{{ result.OutliersMSE.toFixed(2) }}
              <span class="relation-highlight">
                <span v-if="result.OutliersMSE.toFixed(2) > (result.MSEMean + Math.sqrt(result.MSEVariance)).toFixed(2)">&gt;</span>
                <span v-else>&lt;</span>
              μ+σ&nbsp;&nbsp;
              </span>
               = {{ result.MSEMean.toFixed(2) }} + {{ Math.sqrt(result.MSEVariance).toFixed(2) }} = {{ (result.MSEMean + Math.sqrt(result.MSEVariance)).toFixed(2) }}</p>
            <p>{{ result.OutliersMSE.toFixed(2) }}
              <span class="relation-highlight">
                <span v-if="result.OutliersMSE.toFixed(2) > (result.MSEMean + 2 * Math.sqrt(result.MSEVariance)).toFixed(2)">&gt;</span>
                <span v-else>&lt;</span>
              μ+2σ
              </span>
               = {{ result.MSEMean.toFixed(2) }} + {{ (2 * Math.sqrt(result.MSEVariance)).toFixed(2) }} = {{ (result.MSEMean + 2 * Math.sqrt(result.MSEVariance)).toFixed(2) }}</p>
            <p>{{ result.OutliersMSE.toFixed(2) }}
              <span class="relation-highlight">
                <span v-if="result.OutliersMSE.toFixed(2) > (result.MSEMean + 3 * Math.sqrt(result.MSEVariance)).toFixed(2)">&gt;</span>
              <span v-else>&lt;</span>
              μ+3σ
              </span>
               = {{ result.MSEMean.toFixed(2) }} + {{ (3 * Math.sqrt(result.MSEVariance)).toFixed(2) }} = {{ (result.MSEMean + 3 * Math.sqrt(result.MSEVariance)).toFixed(2) }}</p>
          </div>
        </div>
    </div>

    <div style="height: 20px"></div>
  </div>
</template>

<script>
import {fraud_detect} from "../../heimboard.config"
import {time_series} from "../../heimboard.config"
import passive from "/src/assets/passive.png"
import series from "/src/assets/series.png"
import func from "/src/assets/func.png"
import gauss from "/src/assets/gauss.png"
import HeaderView from "@/components/HeaderView";
import SideMenu from "@/components/SideMenu";
import axios from "axios";
axios.defaults.baseURL = '/api';

export default {
  name: "FraudDetectView",
  components: {HeaderView, SideMenu},
  data()
  {
    return {
      passive: passive,
      series: series,
      func: func,
      gauss: gauss,
      model: fraud_detect.model,
      time_series: time_series
    }
  },
  created() {
    if (typeof this.$store.state.user.token === "undefined") {
      this.$router.push('/login')
    }else {
      // 关键词列表
      axios.post('/api/keyword/list_keyword', {}, {
        headers: {
          'Authorization': this.$store.state.user.token
        }
      }).then(response => {
        this.$store.state.keyword_list=response.data.data.keyword_list
      }).catch(error => {
        alert(error.response.data.msg)
      })

      // 获取诈骗信息日志
      axios.post("/api/fraud_detect/get_fraud_messages", {
      }, {
        headers: {
          'Authorization': this.$store.state.user.token
        }
      }).then(response => {
        this.$store.state.fraud_messages = response.data.data.fraud_messages
        // 将内容高亮处理
        var keywords = []
        for (let i = 0; i < this.$store.state.keyword_list.length; i++) {
          keywords.push(this.$store.state.keyword_list[i].Keyword)
        }
        // 将关键词转换为正则表达式的形式，使用 | 分隔每个关键词
        var regex_fraud = new RegExp(keywords.join("|"), "gi");
        var regex_cq = new RegExp("\\[CQ:.*?\\]", "gi");
        var regex_qq = new RegExp("\\d{8,10}", "gi");
        for (let i = 0; i < this.$store.state.fraud_messages.length; i++) {
          // 使用正则表达式进行匹配和替换
          this.$store.state.fraud_messages[i].Content = this.$store.state.fraud_messages[i].Content.replace(regex_cq, function(match) {
            return '<span style="background-color: '+fraud_detect.cq_color+';">' + match + '</span>';
          });
          // 对关键词的高亮只有当存在关键词时才进行
          if (keywords.length > 0){
            this.$store.state.fraud_messages[i].Content = this.$store.state.fraud_messages[i].Content.replace(regex_fraud, function(match) {
              return '<span style="background-color: '+fraud_detect.fraud_color+';">' + match + '</span>';
            });
          }
          this.$store.state.fraud_messages[i].Content = this.$store.state.fraud_messages[i].Content.replace(regex_qq, function(match) {
            return '<span style="background-color: '+fraud_detect.qq_color+';">' + match + '</span>';
          });
        }
      }).catch(error => {
        alert(error.response.data.msg)
        this.$router.push("/")
      })

      // 涉诈群聊日志获取
      axios.post('/api/fraud_detect/get_fraud_groups', {}, {
        headers: {
          'Authorization': this.$store.state.user.token
        }
      }).then(response => {
        this.$store.state.fraud_groups=response.data.data.fraud_groups
      }).catch(error => {
        alert(error.response.data.msg)
      })

      // 时间序列检测结果获取
      axios.post('/api/time_series/get_time_series_result', {}, {
        headers: {
          'Authorization': this.$store.state.user.token
        }
      }).then(response => {
        this.$store.state.time_series_results=response.data.data.result
      }).catch(error => {
        alert(error.response.data.msg)
      })
    }
  }
}
</script>

<style scoped>
.time-series-result-box{
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
.time-series-title-box{
  font-weight: 300;
  font-size: 1.2em;
  font-family: Rubik, sans-serif;
  text-align: center;
  color: #636e72;
}
.time-series-function-box{
  display: flex;
  flex-direction: row;
  justify-content: space-around;
}
.time-series-function-single{
  display: flex;
}
.time-series-function-single-pic{
  width: 50px;
  height: 50px;
}
.time-series-function-single-pic img{
  width: 100%;
  height: 100%;
}
.time-series-function-content{
  font-size: 1.2em;
  font-family: Rubik, sans-serif;
}
.time-series-mse-box{
  display: flex;
  justify-content: space-around;
  align-items: center;
}
.time-series-mse-pic{
  width: 200px;
  height: 200px;
}
.time-series-mse-pic img{
  width: 100%;
  height: 100%;
}
.time-series-mse-word-box{
  font-size: 1.8em;
  font-family: Rubik, sans-serif;
}
.target-highlight:before{
  content: '';
  position: absolute;
  z-index: -1;
  border: none;
  border-radius: 4px;
  height: 15px;
  transform: translateY(20px) translateX(-5px);
  background: #ff6b81;
  width: 100px;
}
.relation-highlight:before{
  content: '';
  position: absolute;
  z-index: -1;
  border: none;
  border-radius: 4px;
  height: 15px;
  transform: translateY(25px) translateX(-5px);
  background: rgba(79, 192, 141, 0.5);
  width: 115px;
}


.model-state-box{
  display: flex;
  justify-content: center;
  align-items: center;
}
.model-state-pic-box{
  width: 200px;
  height: 200px;
  border-radius: 50%;
  padding: 18px;
  box-shadow: 0 10px 20px rgba(0 0 0 / 20%);
  border: 1px solid rgba(0, 0, 0, 0.1);
  position: relative;
  right: -24px;
  background-color: #fff;
  z-index: 1;
}
.model-state-pic-box img{
  width: 100%;
  height: 100%;
}
.model-state-word-box{
  display: flex;
  flex-direction: column;
  justify-content: center;
  align-items: center;
  border-radius: 12px;
  padding: 12px 96px 12px 96px;
  box-shadow: 0 10px 20px rgba(0 0 0 / 20%);
  border: 1px solid rgba(0, 0, 0, 0.1);
  position: relative;
  left: -24px;
  max-width: 60%;
  word-break: break-all;
  cursor: default;
}
.model-state-word-box h1{
  font-family: Rubik,sans-serif;
  font-size: 3em;
  font-weight: 200;
}
.model-state-word-box h2{
  font-family: Rubik,sans-serif;
  font-size: 1.2em;
  font-weight: 200;
  color: #636e72;
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

.fraud-message-box{
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
.fraud-message-box p{
  font-size: 1.2em;
  font-weight: 200;
  font-family: Rubik, sans-serif;
}
.fraud-message-box p span:nth-child(3){
  display: inline-flex;
  width: 65%;
}

.related-state-box{
  display: flex;
  justify-content: center;
  align-items: center;
}
.related-state-word-box{
  display: flex;
  flex-direction: column;
  justify-content: center;
  align-items: center;
  border-radius: 12px;
  padding: 12px 96px 12px 96px;
  box-shadow: 0 10px 20px rgba(0 0 0 / 20%);
  border: 1px solid rgba(0, 0, 0, 0.1);
  position: relative;
  left: -24px;
  max-width: 60%;
  word-break: break-all;
  cursor: default;
}
.related-state-word-box h1{
  font-family: Rubik,sans-serif;
  font-size: 3em;
  font-weight: 200;
}
.related-state-word-box h2{
  font-family: Rubik,sans-serif;
  font-size: 1.2em;
  font-weight: 200;
  color: #636e72;
}
</style>