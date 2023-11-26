# 介绍
此为Geth远程安装目录，需要通过ansible-playbook将程序拷贝到其它机器上

# 文件介绍
Geth相关的文件统一放到geth_pack中

重要文件和目录介绍：
```shell
./geth_pack/conf：存放每个节点挖矿的地址，文件名字和主机列表中名字对应
```
```shell
./geth_pack/geth.json：ETH私链配置文件，请务必保持networkid和启动的networkid保持一致
```

# 执行命令
因为使用相对路径，需要再当前目录下操作
```shell
ansible-playbook -i ansible_hosts_geth.yaml ansible_deploy_geth.yaml
```
# 主机列表
> 在使用ansible_hosts_geth.yaml的时候请严格按照Geth多节点集群进行更改。

## 本项目的Fabric结构
本项目搭建的为四节点ETH私链
| ip地址       | 节点名称  |
| :----------- |:------------:|
|   192.168.93.137    |   geth1  |
|   192.168.93.138    |   geth2  |
|   192.168.93.139    |   geth3  |
|   192.168.93.140    |   geth4  |

## 主机列表配置
```shell
geth:
  vars:
    workpath: "/opt/block/geth" # 远程节点工作目录

  hosts:
    geth1: # 远程节点名字，此名字和每个节点的挖矿地址相关
      ansible_host: 192.168.93.137
      ansible_user: geth 
    geth2:
      ansible_host: 192.168.93.138 
      ansible_user: geth 
    geth3:
      ansible_host: 192.168.93.139 
      ansible_user: geth
    geth4:
      ansible_host: 192.168.93.140 
      ansible_user: geth


```

# 任务列表
具体的任务流程如下：
1. 安装服务：安装docker、docker-compose、firewalld等服务
2. 配置环境：开发端口、配置docker镜像等
3. 安装geth节点
4. 添加enode节点：主要是组成私有网络