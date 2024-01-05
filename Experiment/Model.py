#encodint=utf-8
"""
@author=wanggang
"""
import numpy as np
import pandas as pd
import tensorflow as tf
from tensorflow.keras.models import Sequential
from tensorflow.keras.layers import Conv1D, Bidirectional, GRU, Dense, Flatten, MultiHeadAttention
from sklearn.model_selection import train_test_split
from sklearn.metrics import accuracy_score, f1_score, recall_score, mean_absolute_error,mean_squared_error
from sklearn.metrics import r2_score

# 读取四个Excel文件
file_names = ['defi11.xlsx', 'defi12.xlsx', 'defi13.xlsx','defi14.xlsx','defi15.xlsx','defi16.xlsx']
dfs = [pd.read_excel(file) for file in file_names]

# 合并数据框
merged_df = pd.concat(dfs, ignore_index=True)

# 将Timestamps列转换为日期时间格式
merged_df['Timestamp'] = pd.to_datetime(merged_df['Timestamp'], unit='s')  # 增加unit='s'来指定时间戳是以秒为单位

# 设置 'Timestamp' 列为索引
merged_df.set_index('Timestamp', inplace=True)

# 去重
merged_df = merged_df.drop_duplicates(subset=["Transaction Hash"], keep="first")

# 打印前五行和后十行
#print(merged_df.shape)
#print(merged_df.head())
#print(merged_df.tail(10))

length = len(merged_df)
print("new_tabular_length", length)
merged_df = merged_df.sort_values(by="Timestamp") #sort
# 取最后1200分钟的数据
last_1200_minutes = merged_df.last('30000T')
# 对 'Timestamp' 列重采样为每分钟 ('1T') 并计数
resampled_df = last_1200_minutes.resample("1T").count() #这个是对的，取最后1200的显示：
# 使用真实数据
real_data = np.array(resampled_df['Transaction Hash'].values)
print(real_data)
# 准备数据，使用滑动窗口构建序列
sequence_length = 100
X, y = [], []

for i in range(len(real_data) - sequence_length):
    X.append(real_data[i:i + sequence_length])
    y.append(real_data[i + sequence_length])

X = np.array(X)
y = np.array(y)

# 划分数据集
X_train, X_temp, y_train, y_temp = train_test_split(X, y, test_size=0.3, random_state=42)
X_val, X_test, y_val, y_test = train_test_split(X_temp, y_temp, test_size=0.5, random_state=42)

# 转换数据形状以符合 Conv1D 和 GRU 输入的要求
X_train_tcn = X_train.reshape(X_train.shape[0], X_train.shape[1], 1)
X_val_tcn = X_val.reshape(X_val.shape[0], X_val.shape[1], 1)
X_test_tcn = X_test.reshape(X_test.shape[0], X_test.shape[1], 1)
# 定义 Multi-Head Self-Attention 模块
class MultiHeadSelfAttention(tf.keras.layers.Layer):
    def __init__(self, d_model, num_heads):
        super(MultiHeadSelfAttention, self).__init__()
        self.multi_head_attention = MultiHeadAttention(num_heads=num_heads, key_dim=d_model // num_heads)

    def call(self, inputs):
        return self.multi_head_attention(inputs, inputs)

# 构建 TCN + BiGRU + Multi-Head Self-Attention + Linear 模型
model = Sequential([
    Conv1D(filters=64, kernel_size=2, activation='relu', input_shape=(sequence_length, 1)),
    Bidirectional(GRU(32, activation='relu', return_sequences=True)),
    MultiHeadSelfAttention(d_model=32, num_heads=4),
    Flatten(),
    Dense(1)
])

# 编译模型
model.compile(optimizer='adam', loss='mean_squared_error')

# 训练模型
history = model.fit(X_train_tcn, y_train, epochs=20, batch_size=1, verbose=2, validation_data=(X_val_tcn, y_val))

# 评估模型
y_train_pred = model.predict(X_train_tcn)
y_val_pred = model.predict(X_val_tcn)
y_test_pred = model.predict(X_test_tcn)


# 计算评估指标
train_mse = mean_squared_error(y_train, y_train_pred)
val_mse = mean_squared_error(y_val, y_val_pred)
test_mse = mean_squared_error(y_test, y_test_pred)

train_mae = mean_absolute_error(y_train, y_train_pred)
val_mae = mean_absolute_error(y_val, y_val_pred)
test_mae = mean_absolute_error(y_test, y_test_pred)
# 计算 R方值
train_r2 = r2_score(y_train, y_train_pred)
val_r2 = r2_score(y_val, y_val_pred)
test_r2 = r2_score(y_test, y_test_pred)

# 计算 MSE 和 RMSE


train_rmse = np.sqrt(train_mse)
val_rmse = np.sqrt(val_mse)
test_rmse = np.sqrt(test_mse)
print("Training MSE:", train_mse)
print("Validation MSE:", val_mse)
print("Test MSE:", test_mse)

print("Training MAE:", train_mae)
print("Validation MAE:", val_mae)
print("Test MAE:", test_mae)
#import matplotlib.pyplot as plt
print("Training RMSE:", train_rmse)
print("Validation RMSE:", val_rmse)
print("Test RMSE:", test_rmse)
print("Training R2:", train_r2)
print("Validation R2:", val_r2)
print("Test R2:", test_r2)
import matplotlib.pyplot as plt

# 训练集、验证集和测试集的真实值和预测结果
y_train_pred = model.predict(X_train_tcn)
y_val_pred = model.predict(X_val_tcn)
y_test_pred = model.predict(X_test_tcn)

# 绘制训练集的真实值和预测结果
plt.figure(figsize=(12, 6))
plt.plot(y_train, label='Original Train Data', color='blue')
plt.plot(y_train_pred, label='Predicted Train Data', linestyle='dashed', color='orange')
plt.title('Training Set: Original vs. Predicted Data')
plt.xlabel('Time Steps')
plt.ylabel('Values')
plt.legend()
plt.show()

# 绘制验证集的真实值和预测结果
plt.figure(figsize=(12, 6))
plt.plot(y_val, label='Original Validation Data', color='green')
plt.plot(y_val_pred, label='Predicted Validation Data', linestyle='dashed', color='red')
plt.title('Validation Set: Original vs. Predicted Data')
plt.xlabel('Time Steps')
plt.ylabel('Values')
plt.legend()
plt.show()

# 绘制测试集的真实值和预测结果
plt.figure(figsize=(12, 6))
plt.plot(y_test, label='Original Test Data', color='purple')
plt.plot(y_test_pred, label='Predicted Test Data', linestyle='dashed', color='brown')
plt.title('Test Set: Original vs. Predicted Data')
plt.xlabel('Time Steps')
plt.ylabel('Values')
plt.legend()
plt.show()
