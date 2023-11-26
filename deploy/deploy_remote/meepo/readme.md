# 介绍
此为Meepo远程安装目录，需要通过ansible-playbook将程序拷贝到其它机器上

# 文件介绍
Meepo相关的文件统一放到meepo_pack中

代码地址为：https://github.com/InPlusLab/Meepo

重要文件介绍：
```shell
./meepo_pack/conf：存放单个节点的配置，该文件名字和主机列表的名称保持一致
```
```shell
./meepo_pack/demo-meepo.json：meepo区块的配置
```


# 执行命令
因为使用相对路径，需要再当前目录下操作
```shell
ansible-playbook -i ansible_hosts_meepo.yaml ansible_deploy_meepo.yaml
```
# 主机列表
> 在使用ansible_hosts_fabric.yaml的时候请严格按照Fabric的集群进行更改。

## 本项目的Fabric结构
本项目搭建的Meepo包含四个节点
| ip地址       | 节点名称  |
| :----------- |:------------:|
|   192.168.93.141    |   node1  |
|   192.168.93.142    |   node2  |
|   192.168.93.143    |   node3  |
|   192.168.93.144    |   node4  |

## 主机列表配置
```shell
meepo:
  vars:
    workpath: "/opt/block/meepo" # 远程机器工作目录

  hosts:
    node1: # 远程节点名字，此名字和每个节点的配置相关
      ansible_host: 192.168.93.141
      ansible_user: meepo 
    node2:
      ansible_host: 192.168.93.142
      ansible_user: meepo
    node3:
      ansible_host: 192.168.93.143
      ansible_user: meepo
    node4:
      ansible_host: 192.168.93.144
      ansible_user: meepo




```

# 任务列表
具体的任务流程如下：
1. 安装服务：安装docker、docker-compose、firewalld等服务
2. 配置环境：开发端口、配置docker镜像等
4. 安装meepo节点
5. 添加enode节点：组成私有网络