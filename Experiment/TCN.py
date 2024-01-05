#encoding=utf-8
"""
@authot=gang wang
date:2023.11.27
"""


import numpy as np
import pandas as pd
from sklearn.metrics import mean_absolute_error, mean_squared_error, r2_score
from sklearn.model_selection import train_test_split
from tensorflow.keras.models import Sequential
from tensorflow.keras.layers import  Dense

# 读取四个Excel文件
file_names = ['defi11.xlsx', 'defi12.xlsx', 'defi13.xlsx', 'defi14.xlsx', 'defi15.xlsx', 'defi16.xlsx']
#file_names = ['SandboxGames.xlsx', 'SandboxGames11.xlsx', 'SandboxGames12.xlsx']
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

# 重采样为每小时 ('1H') 并计数
resampled_df = merged_df.resample("1H").count()

# 使用真实数据
real_data = np.array(resampled_df['Transaction Hash'].values)

# 准备时间序列数据
sequence_length = 10  # 设置时间序列长度
X, y = [], []

for i in range(len(real_data) - sequence_length):
    X.append(real_data[i:i + sequence_length])
    y.append(real_data[i + sequence_length])

X = np.array(X)
y = np.array(y)

# 划分数据集
X_train, X_temp, y_train, y_temp = train_test_split(X, y, test_size=0.3, random_state=42)
X_val, X_test, y_val, y_test = train_test_split(X_temp, y_temp, test_size=0.5, random_state=42)

# # 转换数据形状以符合 TCN 输入的要求
# X_train_tcn = X_train.reshape(X_train.shape[0], X_train.shape[1], 1)
# X_val_tcn = X_val.reshape(X_val.shape[0], X_val.shape[1], 1)
# X_test_tcn = X_test.reshape(X_test.shape[0], X_test.shape[1], 1)

# 构建 TCN 模型
import tensorflow as tf
from tensorflow.keras import layers
from tensorflow.keras import backend as K

import tensorflow as tf
from tensorflow.keras import layers

import tensorflow as tf
from tensorflow.keras import layers

# Remove the TCN class from the tcn_model function and place it outside

class TCN(layers.Layer):
    def __init__(self, nb_filters, kernel_size, nb_stacks, dilations, padding='causal', activation='relu', dropout_rate=0.0, use_skip_connections=True, return_sequences=True, input_shape=None):
        super(TCN, self).__init__(input_shape=input_shape)
        self.nb_filters = nb_filters
        self.kernel_size = kernel_size
        self.nb_stacks = nb_stacks
        self.dilations = dilations
        self.padding = padding
        self.activation = activation
        self.dropout_rate = dropout_rate
        self.use_skip_connections = use_skip_connections
        self.return_sequences = return_sequences

        self.conv_layers = []
        self.dropout_layers = []
        for s in range(nb_stacks):
            for d in dilations:
                self.conv_layers.append(layers.Conv1D(filters=nb_filters, kernel_size=kernel_size,
                                                      dilation_rate=d, padding=padding, activation=activation))
                self.dropout_layers.append(layers.Dropout(rate=dropout_rate))

        self.conv_layers = [item for sublist in zip(self.conv_layers, self.dropout_layers) for item in sublist]
        self.concat = layers.Concatenate(axis=-1)

    def call(self, inputs, training=None):
        x = inputs
        skip_connections = []
        for i in range(len(self.conv_layers) // 2):
            x = self.conv_layers[2 * i](x)
            x = self.conv_layers[2 * i + 1](x)
            if self.use_skip_connections:
                skip_connections.append(x)
        if self.use_skip_connections:
            x = self.concat(skip_connections)
        if not self.return_sequences:
            x = x[:, -1, :]
        return x

# Use the TCN layer in the model outside the function

def tcn_model(units, input_shape, nb_stacks, dilations, name="tcn"):
    model = Sequential()
    model.add(TCN(return_sequences=True, nb_filters=units, kernel_size=2, activation='relu', input_shape=input_shape, use_skip_connections=True, nb_stacks=nb_stacks, dilations=dilations))
    model.add(Dense(1))
    model.compile(optimizer='adam', loss='mse')
    return model

# Initialize TCN model
tcn_units = 64
nb_stacks = 1  # You can adjust this
dilations = [1, 2, 4, 8]  # You can adjust this
model = tcn_model(units=tcn_units, input_shape=(sequence_length, 1), nb_stacks=nb_stacks, dilations=dilations)


# Rest of the code remains the same
X_train_tcn = X_train.reshape(X_train.shape[0], X_train.shape[1], 1)
X_val_tcn = X_val.reshape(X_val.shape[0], X_val.shape[1], 1)
X_test_tcn = X_test.reshape(X_test.shape[0], X_test.shape[1], 1)

# Example reshaping for X_train, X_val, and X_test
X_train_reshaped = X_train.reshape((X_train.shape[0], -1))
X_val_reshaped = X_val.reshape((X_val.shape[0], -1))
X_test_reshaped = X_test.reshape((X_test.shape[0], -1))

# 训练模型
history = model.fit(X_train_tcn, y_train, epochs=10, batch_size=1, verbose=2, validation_data=(X_val_tcn, y_val))

# 评估模型
y_train_pred = model.predict(X_train_reshaped)
y_val_pred = model.predict(X_val_tcn)
y_test_pred = model.predict(X_test_tcn)
print(y_test_pred[:, -1, :])
print("-------------------------------")
print(y_train.shape)
# 计算评估指标
print(y_test.shape)
#print()
new_shape = y_test_pred.shape
reshaped_array = y_test_pred.reshape((new_shape[0], new_shape[1] * new_shape[2]))
final_reshaped_array = reshaped_array[:, :1]

#train_mse = mean_squared_error(y_train, y_train_pred)
#val_mse = mean_squared_error(y_val, y_val_pred)
test_mse = mean_squared_error(y_test, final_reshaped_array)

#train_mae = mean_absolute_error(y_train, y_train_pred)
#val_mae = mean_absolute_error(y_val, y_val_pred)
test_mae = mean_absolute_error(y_test, final_reshaped_array)
test_rmse = np.sqrt(test_mse)
#print("Training MSE:", train_mse)
#print("Validation MSE:", val_mse)
print("Test MSE:", test_mse)

#print("Training MAE:", train_mae)
#print("Validation MAE:", val_mae)
print("Test MAE:", test_mae)

# 计算 R2 指标
#r2_train = r2_score(y_train, y_train_pred)
#r2_val = r2_score(y_val, y_val_pred)
r2_test = r2_score(y_test, final_reshaped_array)
print("Test     RMSE:", test_rmse)
#print("Training R2:", r2_train)
#print("Validation R2:", r2_val)
print("Test R2:", r2_test)
