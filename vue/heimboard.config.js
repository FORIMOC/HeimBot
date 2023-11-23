module.exports = {
  version: '3.0.0_build_20230809',
  author: '',
  server: {
    host: '127.0.0.1',
    port: 8740
  },
  fraud_detect: {
    model: 'bert-base-chinese预训练模型+自制数据集微调的文本分类模型',
    fraud_color: '#ff7f50',
    cq_color: '#b2bec3',
    qq_color: '#fad390'
  },
  time_series: {
    algorithm: 'FFT快速傅里叶+ACF自相关算法',
    period: '5天',
  }
};
