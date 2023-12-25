# Introduction
This is the Geth remote installation directory. You need to copy the program to other machines through ansible-playbook.

# Document introduction
Geth related files are unified into geth_pack

Introduction to important files and directories：
```shell
./geth_pack/conf：Stores the mining address of each node. The file name corresponds to the name in the host list.
```
```shell
./geth_pack/geth.json：ETH private chain configuration file, please be sure to keep the networkid consistent with the started networkid
```

# Execute a command
Because relative paths are used, operations need to be done in the current directory.
```shell
ansible-playbook -i ansible_hosts_geth.yaml ansible_deploy_geth.yaml
```
# Host List
> When using ansible_hosts_geth.yaml, please make changes strictly in accordance with the Geth multi-node cluster.

## Fabric structure of this project
This project builds a four-node ETH private chain
| ip address       | node name  |
| :----------- |:------------:|
|   192.168.93.137    |   geth1  |
|   192.168.93.138    |   geth2  |
|   192.168.93.139    |   geth3  |
|   192.168.93.140    |   geth4  |

## Host list configuration
```shell
geth:
  vars:
    workpath: "/opt/block/geth" # Remote node working directory

  hosts:
    geth1: # Remote node name, this name is related to the mining address of each node
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

# Task List
The specific task process is as follows：
1. install service：Install docker, docker-compose, firewalld and other services
2. Configuration Environment：Development port, configure docker image, etc.
3. Install geth node
4. Add enode nodes: mainly to form a private network
