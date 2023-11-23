import json

from flask import Flask, request
import torch
import torch.nn as nn
from transformers import BertTokenizer, BertModel
from torch.utils.data import DataLoader

import numpy as np
from scipy.fftpack import fft, fftfreq
from statsmodels.tsa.stattools import acf
from scipy.interpolate import interp1d

# 本地预训练模型路径
BERT_PATH = 'bert-base-chinese'
# 模型路径
MODEL_PATH = 'model/bert_9.pt'
# 设备
device = 'cpu'
# 分组大小
batch_size = 10
# 分词器
tokenizer = BertTokenizer.from_pretrained(BERT_PATH)

# 时间序列放缩长度
SCALE_REGULATE_LENGTH = 1024
# 时间序列放缩类型
# cubic 样条插值 linear 线性插值
SCALE_REGULATE_KIND = 'cubic'


# BERT 模型
class BertClassifier(nn.Module):
    def __init__(self, dropout=0.5):
        super(BertClassifier, self).__init__()
        self.bert = BertModel.from_pretrained(BERT_PATH)
        self.dropout = nn.Dropout(dropout)
        self.linear = nn.Linear(768, 2)
        self.relu = nn.ReLU()

    def forward(self, input_id, mask):
        _, pooled_output = self.bert(input_ids=input_id, attention_mask=mask, return_dict=False)
        dropout_output = self.dropout(pooled_output)
        linear_output = self.linear(dropout_output)
        final_layer = self.relu(linear_output)
        return final_layer


app = Flask(__name__)


# 从 POST 数据包接受 JSON 数据
# {"texts": ["text1", "text2"]}
# {"results": [0, 1], "status": "success"}
@app.route('/fraud_detect', methods=['POST'])
def fraud_detect():
    data = request.get_json()
    # 数据存在且是字典类型, 表示是合法的JSON, 可以继续处理
    if data is not None and isinstance(data, dict):
        texts = data.get('texts')
        # 如果存在 texts 字段
        if texts is not None:
            # 如果 texts 是字符串列表
            if isinstance(texts, list) and all(isinstance(item, str) for item in texts):
                # texts 长度大于 30 (过长)
                if len(texts) > 30:
                    resp = {"result": [], "status": "failed", "msg": "Too much request text"}
                    return json.dumps(resp)
            # 如果 texts 不是字符串列表
            else:
                resp = {"result": [], "status": "failed", "msg": "`texts` should be a list of strings"}
                return json.dumps(resp)
        else:
            # 不存在 texts 字段
            resp = {"result": [], "status": "failed", "msg": "The `texts` field is required"}
            return json.dumps(resp)
    else:
        # 数据不是合法的 JSON
        resp = {"result": [], "status": "failed", "msg": "A legal JSON which contains `texts` field is required"}
        return json.dumps(resp)

    # 模型读取
    model = BertClassifier()
    model.load_state_dict(torch.load(MODEL_PATH, map_location=torch.device(device)))
    model = model.to(device)
    model.eval()

    # 将输入的 texts 处理为 bert 的输入格式
    def preprocess_text(text):
        encoded_input = tokenizer.encode_plus(
            text,
            add_special_tokens=True,
            padding='max_length',
            max_length=256,
            return_tensors='pt',
            truncation=True
        )
        input_id = encoded_input['input_ids']
        mask = encoded_input['attention_mask']
        return input_id, mask

    # 批量预测
    def predict_batch(texts):
        input_ids = []
        masks = []
        for text in texts:
            input_id, mask = preprocess_text(text)
            input_ids.append(input_id)
            masks.append(mask)
        input_ids = torch.cat(input_ids, dim=0)
        masks = torch.cat(masks, dim=0)
        dataset = torch.utils.data.TensorDataset(input_ids, masks)
        dataloader = DataLoader(dataset, batch_size=batch_size)
        predictions = []
        for batch in dataloader:
            batch = tuple(t.to(device) for t in batch)
            input_id, mask = batch
            with torch.no_grad():
                outputs = model(input_id.to(device), mask.to(device))
                _, predicted_labels = torch.max(outputs, dim=1)
                predictions.extend(predicted_labels.cpu().numpy())
        return predictions

    result_list = predict_batch(texts)
    result_list = [int(x) for x in result_list]
    resp = {"result": result_list, "status": "success"}
    print(resp)
    return json.dumps(resp)


# 从 POST 数据包接受 JSON 数据
# {"time_series": [4, 1, 3], "top_k": 3}
# {"result": [{"amplitude": 32.509, "freq": 2.0, "acf": 0.473}, {"amplitude": 9.432, "freq": 4.0, "acf": -0.546}], "status": "success"}
@app.route('/time_series_detect', methods=['POST'])
def time_series_detect():
    data = request.get_json()
    # 数据存在且是字典类型, 表示是合法的JSON, 可以继续处理
    if data is not None and isinstance(data, dict):
        time_series = data.get('time_series')
        # 如果存在 time_series 字段
        if time_series is not None:
            # 如果 time_series 是整型列表
            if isinstance(time_series, list) and all(isinstance(item, int) for item in time_series):
                # time_series 长度小于等于 5 (过短)
                if len(time_series) <= 5:
                    resp = {"result": [], "status": "failed", "msg": "Time series too short"}
                    return json.dumps(resp)
                # time_series 长度大于 10000 (过长)
                if len(time_series) > 10000:
                    resp = {"result": [], "status": "failed", "msg": "Time series too long"}
                    return json.dumps(resp)
            # 如果 time_series 不是整型列表
            else:
                resp = {"result": [], "status": "failed", "msg": "`time_series` should be a list of integers"}
                return json.dumps(resp)
        else:
            # 不存在 time_series 字段
            resp = {"result": [], "status": "failed", "msg": "The `time_series` field is required"}
            return json.dumps(resp)
        top_k = data.get('top_k')
        # 如果存在 top_k 字段
        if top_k is not None:
            # 如果 top_k 不是整型
            if not isinstance(top_k, int):
                resp = {"result": [], "status": "failed", "msg": "`top_k` should be a integer"}
                return json.dumps(resp)
        # 不存在 top_k 字段
        else:
            resp = {"result": [], "status": "failed", "msg": "`top_k` field is required"}
            return json.dumps(resp)
    else:
        # 数据不是合法的 JSON
        resp = {"result": [], "status": "failed", "msg": "A legal JSON which contains `time_series` field is required"}
        return json.dumps(resp)

    # === 时间序列的放缩 ===

    new_time = np.linspace(0, 1, SCALE_REGULATE_LENGTH)

    # 使用 cubic 样条插值
    f = interp1d(np.linspace(0, 1, len(time_series)), time_series, kind=SCALE_REGULATE_KIND)
    time_series = f(new_time).tolist()
    time_series = [round(num, 3) for num in time_series]

    # === 快速傅里叶变换计算前 k 个最强子序列 ===

    # 对数据应用快速傅里叶变换，将数据从时域转换到频域
    fft_series = fft(time_series)
    # 计算频谱的功率（取复数结果的绝对值）
    # power[k] = |X[k]| = sqrt(A[k]^2 + B[k]^2)
    power = np.abs(fft_series)
    # 计算fft_series每个元素的采样频率
    sample_freq = fftfreq(fft_series.size)

    # 大于0的样本频率为真
    pos_mask = np.where(sample_freq > 0)
    # 过滤掉所有不大于0的频率，即仅包含正频率
    freqs = sample_freq[pos_mask]
    # 同样过滤掉所有不大于0的功率，即仅包含正功率
    powers = power[pos_mask]

    try:
        # 查找 powers 数组中顶部 top_k 元素的索引
        top_k_idxs = np.argpartition(powers, -top_k)[-top_k:]
    # top_k 过大或过小
    except:
        resp = {"result": [], "status": "failed", "msg": "`top_k` is too small or too large"}
        return json.dumps(resp)

    # top_k 的功率
    top_k_power = powers[top_k_idxs][::-1]
    # top_k 的振幅
    top_k_amplitude = top_k_power / (len(time_series) / 4)
    # top_k 的周期数量
    top_k_period_num = (1 / freqs[top_k_idxs]).astype(int)[::-1]
    # top_k 的频率
    top_k_freq = len(time_series) / top_k_period_num
    # top_k 的相位差（默认为正弦函数）
    top_k_phase = np.zeros(top_k)

    # 逐个处理每个最强子序列的频率
    for i in range(top_k):
        # 获取最强子序列的频率索引
        freq_idx = top_k_idxs[i]

        # 获取相位差
        top_k_phase[i] = np.angle(fft_series[freq_idx])

        # 限制相位差在 [0, 2pi] 范围内
        top_k_phase[i] = np.mod(top_k_phase[i], 2 * np.pi)
    top_k_phase = top_k_phase.tolist()

    # 计算前 k 个最强子序列对应的 ACF 自相关系数
    acf_scores = []
    for lag in top_k_period_num:
        acf_scores.append(acf(time_series, nlags=lag)[-1])

    # 返回结果
    top_k_amplitude = [round(num, 3) for num in top_k_amplitude]
    top_k_freq = [round(num, 3) for num in top_k_freq]
    top_k_phase = [round(num, 3) for num in top_k_phase]
    acf_scores = [round(num, 3) for num in acf_scores]
    result = []
    for i in range(top_k):
        result.append({
            "amplitude": top_k_amplitude[i],
            "freq": top_k_freq[i],
            "phase": top_k_phase[i],
            "acf": acf_scores[i],
        })
    resp = {"result": result, "status": "success"}
    return json.dumps(resp)


if __name__ == '__main__':
    app.run(host="0.0.0.0", port=8742)
