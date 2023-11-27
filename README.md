
# 介绍
此为区块链压测工具，使用Golang开发，支持ETH、Fabric、Meepo、NeuChain（待完成）四类区块链。
# 架构图
<div align=center><img width="1000" height="600" src="https://github.com/masnail/BlockBenchMark/assets/51044388/87ac2cfa-05cd-4f3f-803c-ba901daab1d1"/></div>


# 软件版本
| 软件名称       |     版本号  |
| :----------- |------------:|
|   ubuntu    |   v22.04  |
|   Go          |   v1.17.13     |
|   HyperLedger Fabric |    v2.2.0|
|   Ethereum    |   v1.10.20 |
|   Meepo    |   v0.0.0（commit:7e84406856） |
|NeuChain |  v0.0.1 |
|ansible   |v2.10.8|
|ansible-playbook   |v2.10.8|
# Deploy
使用ansible-playbook通过主机列表和任务列表进行集群管理和自动部署，使用前请提前对控制端和远程端进行免密登录和sudo权限免密获取。

部署主要分为[监控部署](./deploy/deploy_local/readme.md)（mysql，prometheus，grafana等），[区块链搭建](./deploy/deploy_remote/readme.md)（Fabric，ETH私链，Meepo），可以快速搭建测试的环境。

## monitor
安装命令：
```shell
    cd deploy/deploy_local
    /bin/bash deploy.sh
```
安装内容如下：
* mysql：用于压测数据存储，一般安装在非server端
* grafana：用于数据展示，一般安装在非server端
* prometheus：用于收集机器数据，进行机器状态监控，一般安装在非server端
* node_exporter：安装在被监控机器侧，上报机器状态数据给prometheus，一般集群中的机器都需要安装

### Grafama配置
安装完成后地址为：http://XXXX:3000

登录用户和密码为：admin、admin

grafana的数据已经配置，路径为：./grafana_data/grafana.db

> **更改mysql和prometheus数据源的IP地址**

## Fabric 搭建
执行命令：
```shell
    cd deploy/deploy_remote/fabric
    ansible-playbook -i ansible_hosts_fabric.yaml ansible_deploy_fabric.yaml
```
> 使用的时候请根据需要先自行修改[主机列表](./deploy/deploy_remote/fabric/ansible_hosts_fabric.yaml)：用于管理Fabric集群节点分配，[任务列表](./deploy/deploy_remote/fabric/ansible_deploy_fabric.yaml)一般情况下不需要更改。主机列表、任务列表和Fabric搭建的详细信息，请点击[此处](./deploy/deploy_remote/fabric/readme.md)。

## Geth 搭建
执行命令：
```shell
    cd deploy/deploy_remote/geth
    ansible-playbook -i ansible_hosts_geth.yaml ansible_deploy_geth.yaml
```
> 使用的时候请根据需要先自行修改[主机列表](./deploy/deploy_remote/geth/ansible_hosts_geth.yaml)：用于管理Geth集群节点分配，[任务列表](./deploy/deploy_remote/geth/ansible_deploy_geth.yaml)一般情况下不需要更改。主机列表、任务列表和Geth搭建的详细信息，请点击[此处](./deploy/deploy_remote/geth/readme.md)。

## Meepo 搭建
执行命令：
```shell
    cd deploy/deploy_remote/meepo
    ansible-playbook -i ansible_hosts_meepo.yaml ansible_deploy_meepo.yaml
```
> 使用的时候请根据需要先自行修改[主机列表](./deploy/deploy_remote/meepo/ansible_hosts_meepo.yaml)：用于管理Geth集群节点分配，[任务列表](./deploy/deploy_remote/meepo/ansible_deploy_meepo.yaml)一般情况下不需要更改。主机列表、任务列表和Geth搭建的详细信息，请点击[此处](./deploy/deploy_remote/meepo/readme.md)。

# Build And Run (Ubuntu)
压测系统主要是由server和client组成，其支持分布式部署。
对于ETH、Fabric、Meepo测试环境部署也是该系统的一部分。
> 所有的测试和运行都是基于Ubuntu22.04进行，golang版本请按照上面建议的版本运行。
## Server
server编译：
```shell
    cd src/server
    go build server.go
```
server运行：
```shell
    ./server -log_dir=log
```
> 需要conf目录和server二进制才能正常运行

Server运行需要读取程序运行配置和区块链相关配置，相关配置具体讲解请见[此处](./src/server/conf/readme.md)。
## Client
client编译：
```shell
    cd src/client
    go build client.go
```
client运行：
```shell
    ./client -log_dir=log
```
> 需要config_client.json文件和client二进制才能运行

Client运行的配置文件具体讲解请见[此处](./src/client/readme.md)

# 其它 
> **对于Fabric、Geth、Meepo请仔细检查配置文件**

