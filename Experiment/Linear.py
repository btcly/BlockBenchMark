#encoding=utf-8
"""
@authot=gang wang
date:2023.11.27
"""
import numpy as np
import pandas as pd
from sklearn.metrics import mean_squared_error, r2_score, mean_absolute_error
from math import sqrt
# 读取四个Excel文件
file_names = ['defi11.xlsx', 'defi12.xlsx', 'defi13.xlsx', 'defi14.xlsx', 'defi15.xlsx', 'defi16.xlsx']
dfs = [pd.read_excel(file) for file in file_names]

# 合并数据框
merged_df = pd.concat(dfs, ignore_index=True)

# 将Timestamps列转换为日期时间格式
merged_df['Timestamp'] = pd.to_datetime(merged_df['Timestamp'], unit='s')  # 增加unit='s'来指定时间戳是以秒为单位

# 设置 'Timestamp' 列为索引
merged_df.set_index('Timestamp', inplace=True)

# 去重
merged_df = merged_df.drop_duplicates(subset=["Transaction Hash"], keep="first")

# 打印前五行
#print(merged_df.head())

# 重采样为每小时 ('1H') 并计数
resampled_df = merged_df.resample("1H").count()

# 使用真实数据
real_data = np.array(resampled_df['Transaction Hash'].values)

# 打印重采样后的前五行
#print(resampled_df.head())


# 生成三万个随机数  data 来源
random_numbers = real_data
print("原始数据",random_numbers)
# 准备数据
X_train = random_numbers[:-1]
print("X数据",X_train)
y_train = random_numbers[1:]
print('标签数据',y_train)
from sklearn.linear_model import LinearRegression

# 初始化线性回归模型
model = LinearRegression()
# 训练模型
model.fit(X_train.reshape(-1, 1), y_train)
# 预测第30001个随机数
prediction = model.predict([[random_numbers[-1]]])
print(prediction)
print("-----------------")
print(random_numbers[-1])
prediction2 = model.predict(prediction.reshape(-1,1))
print("预测30002个：",prediction2)
# 计算评估指标
y_pred = model.predict(X_train.reshape(-1, 1))
mse = mean_squared_error(y_train, y_pred)
r2 = r2_score(y_train, y_pred)
mae = mean_absolute_error(y_train, y_pred)
rmse = sqrt(mse)

print("Mean Squared Error (MSE):", mse)
print("R-squared (R2):", r2)
print("Mean Absolute Error (MAE):", mae)
print("Root Mean Squared Error (RMSE):", rmse)