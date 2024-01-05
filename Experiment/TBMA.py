import tensorflow as tf
from tensorflow.keras.models import Model
from tensorflow.keras.layers import Input, Dense, Conv1D, Bidirectional, GRU, Flatten
from tensorflow.keras.layers import MultiHeadAttention, LayerNormalization
import numpy as np
from sklearn.metrics import mean_squared_error, mean_absolute_error, r2_score
import pandas as pd
# TCN Block
class TCNBlock(tf.keras.layers.Layer):
    def __init__(self, filters, kernel_size, dilation_rate):
        super(TCNBlock, self).__init__()
        self.conv1d = Conv1D(filters, kernel_size, dilation_rate=dilation_rate, padding='same', activation='relu')
        self.norm = LayerNormalization()

    def call(self, x):
        return self.norm(self.conv1d(x))

# Building the Model
def build_model(input_shape, hidden_units, num_heads):
    inputs = Input(shape=input_shape)

    # TCN layer
    tcn = TCNBlock(hidden_units, kernel_size=3, dilation_rate=1)(inputs)

    # BiGRU layer
    bigru = Bidirectional(GRU(hidden_units, return_sequences=True))(tcn)

    # Multi-Head Attention layer
    attn_output = MultiHeadAttention(num_heads=num_heads, key_dim=hidden_units)(bigru, bigru)

    # Flatten and Dense layer for output
    flatten = Flatten()(attn_output)
    outputs = Dense(1)(flatten)

    return Model(inputs=inputs, outputs=outputs)
if __name__ == '__main__':
    # 读取四个Excel文件
    # file_names = ['defi11.xlsx', 'defi12.xlsx', 'defi13.xlsx', 'defi14.xlsx', 'defi15.xlsx', 'defi16.xlsx']
    #file_names = ['SandboxGames.xlsx', 'SandboxGames11.xlsx', 'SandboxGames12.xlsx']
    file_names = ['ethereum_transactions11.xlsx', 'ethereum_transactions011.xlsx', 'ethereum_transactions12.xlsx', 'ethereum_transactions013.xlsx']
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
    print("_________________________________++++++++++++++")
    print(real_data.shape)

    # Model parameters
    num_series = 1  # Number of features in the time series
    seq_len = 10  # Length of the time series
    hidden_units = 64  # Hidden units for layers
    num_heads = 4  # Number of attention heads

    # Create the model
    model = build_model((seq_len, num_series), hidden_units, num_heads)

    # Print model summary
    model.summary()



    # 重新塑造数据
    num_samples = len(real_data) - seq_len + 1
    reshaped_data = np.zeros((num_samples, seq_len, num_series))

    for i in range(num_samples):
        reshaped_data[i, :, 0] = real_data[i:i + seq_len]

    # 打印一下重新塑造后的数据的形状
    print(reshaped_data.shape)
    # 获取相应数量的 y 数据
    y_data = real_data[seq_len - 1:num_samples + seq_len - 1, np.newaxis]
    batch_size=32


    #example_data = np.random.rand(batch_size, seq_len, num_series).astype(np.float32)
    example_data=reshaped_data
    print("++++++++++++",example_data.shape)
    model.compile(optimizer='adam', loss='mse')
    model.fit(example_data, y_data, epochs=10, batch_size=batch_size)
    # Model prediction
    prediction = model.predict(example_data)
    print(prediction)
    print(prediction.shape)  # Should be (batch_size, 1) representing the prediction for each series

    # 假设 y_true 是真实值，y_pred 是预测值
    #y_true = np.random.rand(batch_size, 1)
    y_true= y_data
    y_pred = model.predict(example_data)

    # 计算 MSE
    mse = mean_squared_error(y_true, y_pred)
    print("Mean Squared Error (MSE):", mse)

    # 计算 MAE
    mae = mean_absolute_error(y_true, y_pred)
    print("Mean Absolute Error (MAE):", mae)

    # 计算 RMSE
    rmse = np.sqrt(mse)
    print("Root Mean Squared Error (RMSE):", rmse)

    # 计算 R²
    r2 = r2_score(y_true, y_pred)
    print("R² Score:", r2)



