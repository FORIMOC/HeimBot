import numpy as np
from scipy.fftpack import fft, fftfreq
import matplotlib.pyplot as plt
from statsmodels.tsa.stattools import acf

# 生成示例数据
time = np.linspace(0, 1, 1024)  # 时间范围从0到1，总共取1000个点
frequency1 = 5  # 第一个频率成分为5Hz
frequency2 = 15  # 第二个频率成分为15Hz
amplitude1 = 2  # 第一个频率成分的振幅
amplitude2 = 1.5  # 第二个频率成分的振幅

# 生成包含两个频率成分的示例数据
data = amplitude1 * np.sin(2 * np.pi * frequency1 * time) \
       + amplitude2 * np.sin(2 * np.pi * frequency2 * time + np.pi / 4)

# 负值设为 0
data = np.clip(data, 0, None)

# 绘制原始数据图形
plt.plot(time, data)
plt.xlabel('Time')
plt.ylabel('Value')
plt.title('Original Data')
plt.show()

# 快速傅里叶变换计算前 k 个最强子序列

# 对数据应用快速傅里叶变换，将数据从时域转换到频域
fft_series = fft(data)
# 计算频谱的功率（取复数结果的绝对值）
# power[k] = |X[k]| = sqrt(A[k]^2 + B[k]^2)
power = np.abs(fft_series)
# 计算 fft_series 每个元素的采样频率
sample_freq = fftfreq(fft_series.size)

# 大于0的样本频率为真
pos_mask = np.where(sample_freq > 0)
# 过滤掉所有不大于0的频率，即仅包含正频率
freqs = sample_freq[pos_mask]
# 同样过滤掉所有不大于0的功率，即仅包含正功率
powers = power[pos_mask]

top_k = 3
# 查找 powers 数组中顶部 top_k 元素的索引
top_k_idxs = np.argpartition(powers, -top_k)[-top_k:]
# top_k 的功率
top_k_power = powers[top_k_idxs][::-1]
# top_k 的周期数量
top_k_period_num = (1 / freqs[top_k_idxs]).astype(int)[::-1]
# top_k 的频率
top_k_freq = len(data) / top_k_period_num
# top_k 的振幅
top_k_amplitude = top_k_power / (len(time) / 4)

# 计算相位差（默认为正弦函数，即相位差为0）
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

print(f"fft_phase_difference: {top_k_phase}")
print(f"top_k_power: {top_k_power}")
print(f"fft_period_num: {top_k_period_num}")
print(f"fft_freq: {top_k_freq}")
print(f"fft_amplitude: {top_k_amplitude}")
