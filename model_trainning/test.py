import pandas as pd

# 读取 CSV 文件
df = pd.read_csv('data/train.csv')

# 分层抽样，抽样比例为 10%
sampled_df = df.groupby('category', group_keys=False).apply(lambda x: x.sample(frac=0.1))

# 保存抽样结果到新的 CSV 文件
sampled_df = df.shuf
sampled_df.to_csv('data/train_0.1.csv', index=False)

# 输出抽样结果的统计信息
print('抽样后的数据行数:', len(sampled_df))
print(sampled_df.head())
