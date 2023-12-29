import matplotlib.pyplot as plt
import numpy as np

pipeline_y = np.array([4.3135, 8.6732, 12.9430, 17.2875, 21.5635])
online_y = np.array([29.6454, 59.2825, 89.0560, 118.2692, 147.8090])

# plt.title('Bar', fontsize=20)
plt.xlabel('Workloads', fontsize=20)
plt.ylabel('Time', fontsize=20)

x = np.arange(len(pipeline_y))  # 产生均匀数组，长度等同于apples

plt.bar(x - 0.2,  # 横轴数据
        pipeline_y,  # 纵轴数据
        0.4,  # 柱体宽度
        color='#62B197',
        label='pipeline')
plt.bar(x + 0.2,  # 横轴数据
        online_y,  # 纵轴数据
        0.4,  # 柱体宽度
        color='#E18E6D', label='online', alpha=0.75)

plt.xticks(x, [2, 4, 6, 8, 10])
plt.xticks(fontproperties='Times New Roman', size=14)
plt.yticks(fontproperties='Times New Roman', size=14)
# 设置坐标轴的粗细
ax = plt.gca();  # 获得坐标轴的句柄
ax.spines['bottom'].set_linewidth(2);
ax.spines['left'].set_linewidth(2);
ax.spines['right'].set_linewidth(2);
ax.spines['top'].set_linewidth(2);
plt.legend()
plt.show()
