# 介绍
此为Fabric远程安装目录，需要通过ansible-playbook将程序拷贝到其它机器上自动安装任务部署并安装合约

# 文件介绍
Fabric相关的文件统一放到fabric_pack中

fabric使用的二进制版本为Fabric-2.2，当前目录中已经按照配置文件提前生成了证书文件和通道文件

重要文件和目录介绍：
```shell
./fabric_pack/orders/create_crypto_config.sh：根据配置生成Fabric的通道和证书文件，并同步到peers目录
```
```shell
./fabric_pack/orders/crypto-config和./fabric_pack/peers/crypto-config： 同为证书文件
```

```shell
./fabric_pack/peers/chaincode：链码存放目录
```
```shell
./fabric_pack/peers/fabricdata/deploy：环境配置和链码安装脚本
```

# 执行命令
因为使用相对路径，需要再当前目录下操作
```shell
ansible-playbook -i ansible_hosts_fabric.yaml ansible_deploy_fabric.yaml
```
# 主机列表
> 在使用ansible_hosts_fabric.yaml的时候请严格按照Fabric的集群进行更改。

## 本项目的Fabric结构
本项目请参考[配置文件](./fabric_pack/orders/crypto-config.yaml)，搭建的是三个order（raft算法），两个org，每个组织两个peer
| ip地址       |     所属组织  | 节点名称  |节点域名  |
| :----------- |:------------:|:------------:|:------------:|
|   192.168.93.153    |   order  |orderer1|orderer1.example.com|
|   192.168.93.154    |   order  |orderer2|orderer2.example.com|
|   192.168.93.155    |   order  |orderer3|orderer3.example.com|
|   192.168.93.156    |   org1  |peer0.org1|peer0.org1.example.com|
|   192.168.93.147    |   org1  |peer1.org1|peer1.org1.example.com|
|   192.168.93.148    |   org2  |peer0.org2|peer0.org2.example.com|
|   192.168.93.149    |   org2  |peer1.org2|peer1.org2.example.com|

## 主机列表配置
```shell
fabric:
  vars:
    workpath: "/opt/block/fabric" # 远程机器工作目录
  
  children:
    orders: # 根据Fabric的配置文件设置order的名字
      hosts:
        orderer1: # 该名字和域名的前缀保持一致
          ansible_host: 192.168.93.153
          ansible_user: test
        orderer2:
          ansible_host: 192.168.93.154
          ansible_user: test
        orderer3:
          ansible_host: 192.168.93.155
          ansible_user: test

    peers: # 根据Fabric的配置文件设置peer的名字
      hosts:
        peer0.org1: # 该名字和域名的前缀保持一致
          ansible_host: 192.168.93.156
          ansible_user: test

        peer1.org1:
          ansible_host: 192.168.93.146
          ansible_user: test

        peer0.org2:
          ansible_host: 192.168.93.147
          ansible_user: test

        peer1.org2:
          ansible_host: 192.168.93.148
          ansible_user: test

    peertool: # 单独设置peer用于链码的打包和通道创建，一般从peer中任意选一个
      hosts:
        peer0.org1:
          ansible_host: 192.168.93.156
          ansible_user: test

```

# 任务列表
具体的任务流程如下：
1. 安装服务：安装docker、docker-compose、firewalld等服务
2. 配置环境：开发端口、配置docker镜像等
3. 安装node：用于机器状态数据上报
4. 安装orders
5. 安装peers
6. 安装链码