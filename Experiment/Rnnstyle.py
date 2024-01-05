# encoding=utf-8
"""
@authot=gang wang
date:2023.11.27
"""
import numpy as np
import pandas as pd
from sklearn.metrics import mean_absolute_error, mean_squared_error, r2_score
from sklearn.model_selection import train_test_split
from tensorflow.keras.models import Sequential
from tensorflow.keras.layers import SimpleRNN, Dense

# 读取四个Excel文件
# 1 file_names = ['defi11.xlsx', 'defi12.xlsx', 'defi13.xlsx', 'defi14.xlsx', 'defi15.xlsx', 'defi16.xlsx']
file_names = ['SandboxGames.xlsx', 'SandboxGames11.xlsx', 'SandboxGames12.xlsx']
#file_names = ['ethereum_transactions11.xlsx', 'ethereum_transactions011.xlsx', 'ethereum_transactions12.xlsx', 'ethereum_transactions013.xlsx']
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
# print(merged_df.head())

# 重采样为每小时 ('1H') 并计数
resampled_df = merged_df.resample("1H").count()

# 使用真实数据
real_data = np.array(resampled_df['Transaction Hash'].values)

# 打印重采样后的前五行
# print(resampled_df.head())

random_numbers = real_data
# 准备时间序列数据
sequence_length = 10  # 设置时间序列长度
X, y = [], []

for i in range(len(random_numbers) - sequence_length):
    X.append(random_numbers[i:i + sequence_length])
    y.append(random_numbers[i + sequence_length])

X = np.array(X)
y = np.array(y)

# 划分数据集
X_train, X_temp, y_train, y_temp = train_test_split(X, y, test_size=0.2, random_state=42)
X_val, X_test, y_val, y_test = train_test_split(X_temp, y_temp, test_size=0.5, random_state=42)

# 构建 RNN 模型
model = Sequential()
model.add(SimpleRNN(50, activation='relu', input_shape=(sequence_length, 1)))
model.add(Dense(1))

model.compile(optimizer='adam', loss='mse')
X_train_reshaped = X_train.reshape((X_train.shape[0], X_train.shape[1], 1))

# 训练模型
model.fit(X_train_reshaped, y_train, epochs=10, batch_size=32, validation_split=0.2)

# 计算训练集的评价指标
y_train_pred = model.predict(X_train_reshaped)
mae_train = mean_absolute_error(y_train, y_train_pred)
mse_train = mean_squared_error(y_train, y_train_pred)
rmse_train = np.sqrt(mse_train)
r2_train = r2_score(y_train, y_train_pred)

print("Train MAE:", mae_train)
print("Train MSE:", mse_train)
print("Train RMSE:", rmse_train)
print("Train R2:", r2_train)

# 计算验证集的评价指标
X_val_reshaped = X_val.reshape((X_val.shape[0], X_val.shape[1], 1))
y_val_pred = model.predict(X_val_reshaped)

mae_val = mean_absolute_error(y_val, y_val_pred)
mse_val = mean_squared_error(y_val, y_val_pred)
rmse_val = np.sqrt(mse_val)
r2_val = r2_score(y_val, y_val_pred)

print("Validation MAE:", mae_val)
print("Validation MSE:", mse_val)
print("Validation RMSE:", rmse_val)
print("Validation R2:", r2_val)

# 计算测试集的评价指标
X_test_reshaped = X_test.reshape((X_test.shape[0], X_test.shape[1], 1))
y_test_pred = model.predict(X_test_reshaped)

mae_test = mean_absolute_error(y_test, y_test_pred)
mse_test = mean_squared_error(y_test, y_test_pred)
rmse_test = np.sqrt(mse_test)
r2_test = r2_score(y_test, y_test_pred)

print("Test MAE:", mae_test)
print("Test MSE:", mse_test)
print("Test RMSE:", rmse_test)
print("Test R2:", r2_test)
