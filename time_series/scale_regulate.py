import numpy as np
from scipy.fftpack import fft
import matplotlib.pyplot as plt
from statsmodels.tsa.stattools import acf
from scipy.interpolate import interp1d

# 生成示例数据
time = np.linspace(0, 1, 128)  # 时间范围从0到1，总共取1000个点
frequency1 = 5  # 第一个频率成分为5Hz
frequency2 = 15  # 第二个频率成分为15Hz
amplitude1 = 1  # 第一个频率成分的振幅
amplitude2 = 0.5  # 第二个频率成分的振幅

# 生成包含两个频率成分的示例数据
data = amplitude1 * np.sin(2 * np.pi * frequency1 * time)\
       + amplitude2 * np.sin(2 * np.pi * frequency2 * time + np.pi/4)

# 负值设为 0
data = np.clip(data, 0, None)

# 放缩到1024长度
new_length = 1024
new_time = np.linspace(0, 1, new_length)

# 使用 cubic 样条插值
f = interp1d(time, data, kind='cubic')
new_data = f(new_time)

plt.plot(new_time, new_data)
plt.xlabel('Time')
plt.ylabel('Amplitude')
plt.title('Rescaled Time Series')
plt.show()
