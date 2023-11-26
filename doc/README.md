# 介绍
## 创新点：
1. 特定应用的负载的曲线无法模拟，提出了一种基于VAE的方法进行生成，扩展了负载。
2. 评价一个区块链评测框架的可复现性，时间段内的KL散度。
3. 正态分布（merge）或者傅里叶变换保持一致性。
## 负载如何设计
smallbank：https://github.com/maghbari/smallbank-benchmark/blob/master/contracts/smallbank/smallbank.go
工作负载和工作数据不同，一个是操作，一个是数据。
负载应该是模拟现实生活中的具体应用。 例如模拟银行业务逻辑的smallbank，主要包括这几个业务。
 real application workloads,
standard benchmark workloads and synthetic workloads.（真实、标准、合成）
### 第一个涉及单个key、value的操作，生成均匀的key，value（无冲突），但是存在偏度查询和写。例如生成了key-value 为(wang,10),(zhang,20),(李，30)，偏度的值在【-1，1】之间。

![image](https://user-images.githubusercontent.com/51044388/225776994-012c2f64-3c04-434e-902a-90d8c11d77dc.png)
### 第二个，关于典型的两个key值之间的状态变化，就是smallbank工作负载。偏度的设置是模拟（转账金额高，或者转账频繁的账户）。其次，账户的金额和转账金额可以设置为幂律分布。
### 第三多个典型的负载就是TPC-C工作负载。是由这五种事务组成
![image](https://user-images.githubusercontent.com/51044388/225888252-ccca731a-3a89-4cb3-b4a1-7049b57cfc09.png)
调节五种事务的比例。所有比例之和为1.


### 第四YCSB特点：由于YCSB具有较少的事务逻辑，因此它适用于通过避免事务处理的复杂影响来演示工作负载访问分布对性能的影响。
四个事务：插入、更新事务、插入事务、扫描事务。数据分布包括四种：均匀、正态、zipf、最近邻分布。ycsb的操作分布为读写比例。
## Caliper 对比实验的代码在github地址：
https://github.com/Seafoodair/neublockchain
## 实验结果
![image](https://user-images.githubusercontent.com/51044388/219957191-e180824e-b358-4b8b-b8c1-afd8e4aaf56c.png)
![image](https://user-images.githubusercontent.com/51044388/219957269-ce8582b9-3b2b-4188-8efb-91ab94373e4f.png)


## 区块链分类
从区块链的准入原则来区分，区块链可分为联盟链和公有链。其中公有链的代表是bitcoin和Ethereum。联盟链的代表是fabric。
从区块链的成块原则可分为：区块链可分为链式结构和图结构
从区块链存储状态可分为：kv类型和关系类型的。
综上所述，绘制了一个结构图
![image](https://user-images.githubusercontent.com/51044388/192081836-7eb7b07e-0a85-4927-8a0f-3154787dc0bd.png)
# 通用的区块链评测框架
涉及的工作主要有区块链进行性能压测，监控指标，瓶颈分析，可视化等工作
## 1.性能测试
性能测试包括负载测试、压力测试、失败测试、配置测试、并发测试、可靠性测试、可扩展测试。
负载测试：逐渐增加系统负载，测试系统性能的测试。（找出峰值）
压力测试：测试在什么负载的条件下，系统处于失效状态。（找到临界值）
可靠性测试：长时间下，系统的稳定情况。一般用平均故障间隔时间来衡量。
可扩展测试：区块链不断扩展新的节点，最多支持多少节点。
## 2.性能指标
性能指标包括吞吐量、延迟、并发数量、资源利用率、错误率、网络吞吐量、错误率、系统稳定性。&nbsp;
吞吐量：单位时间内成功上链的事务（一般是指每秒钟多少事务上链）
延迟：事务从发起到上链所需要的时间。
资源利用率：处理事务过程中，cpu、内存、硬盘、宽带等软硬件的使用情况。为了能够在达到peak的情况下找到最有的资源配置。
系统稳定性：通常是至少连续运行24h以上，区块链系统稳定运行的时间。
## 3.性能测试的通用过程
<div align=center><img src="https://user-images.githubusercontent.com/51044388/192430925-c1647cf5-4d4c-46d7-ba29-46ae0c8f2167.png"/></div>

### 1.需求分析阶段
#### 一、为用户提供性能分析结果报告
#### 二、为开发者一套提供标准的评测方法（包括评测指标），分析挖掘区块链系统可能存在的瓶颈。
#### 三、开发一套自动化部署+评测的通用软件。
### 2.测试计划
#### 一、测试各个区块链的TPS,QPS,延迟。
#### 二、测试区块链运行过程中，CPU、内存、硬盘的使用率。
#### 三、区块链参数的变化，对区块链性能的影响。


### 3.测试用例
#### 一、测试以太坊公链和联盟链（转账和查询）
![image](https://user-images.githubusercontent.com/51044388/192133483-fafd0e9b-b8dc-4988-818d-cb4116a63514.png)
![image](https://user-images.githubusercontent.com/51044388/192136814-b5e8754b-6f69-4b02-a0a5-d55f92cdb920.png)


#### 二、测试fabric
### 4.脚本编写
使用的go语言进行编写。
### 测试环境
见环境规格


## 5.类似评测框架
| 框架 | 语言 |
|:-----------|----------:|
| BCTMARK | python |
| Hyperbench | go |
| Caliper | node.js |
| BlockBench | c++ |
| Gormit | python |
## 6.目前代码完成情况
完成了fabric fastfabric fabricsharp fabric++ ethereum 等调试
跑通了hyperbench和caliper的测试环境
代码能够测试公链和私有链。仅限于以太坊和fabric
# 总体架构
![image](https://user-images.githubusercontent.com/51044388/192221014-ff50627a-327b-49e6-861f-cfdb33397efa.png)

## master详细设计
![image](https://user-images.githubusercontent.com/51044388/192418373-bb176134-61b5-440e-a058-30d6c635c764.png)

## 客户端与接口设计
![image](https://user-images.githubusercontent.com/51044388/192420906-44e0612d-cdce-4442-b47f-a1734c0331f8.png)

## 内部设计
对区块链的验证和共识下手

1.可分为共识开始
2.共识完成
3.验证开始
4.验证结束
方法：使用log 对其进行打标签。
另外一种方法是火焰图

## 火焰图
https://juejin.cn/post/6986437088337985572
![image](https://user-images.githubusercontent.com/51044388/194694166-c4bb3b65-12c5-45c9-8e22-daba0da69cc3.png)

瓶颈分析的方法总结：
前面提到的5种技术——趋势、相关、比较、消除和模式匹配——是进行软件性能瓶颈分析的一些指导原则。每一个性能瓶颈分析都是一项探索任务，在实践中它更像是艺术而不是科学。因此，结合所有这些技能，专注和毅力是成功的关键!
