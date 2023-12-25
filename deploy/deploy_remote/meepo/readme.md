# Introduction
This is the Meepo remote installation directory. You need to copy the program to other machines through ansible-playbook.

# Document introduction
Meepo related files are unified into meepo_pack

The code address is：https://github.com/InPlusLab/Meepo

Important document introduction：
```shell
./meepo_pack/conf：Stores the configuration of a single node. The name of this file must be consistent with the name of the host list.
```
```shell
./meepo_pack/demo-meepo.json：meepo's Block configuration
```


# Execute commands
Because relative paths are used, operations need to be done in the current directory.
```shell
ansible-playbook -i ansible_hosts_meepo.yaml ansible_deploy_meepo.yaml
```
# Host List
> When using ansible_hosts_fabric.yaml, please make changes strictly in accordance with the Fabric cluster.

## Fabric structure of this project
Meepo built in this project contains four nodes
| ip address       | node name  |
| :----------- |:------------:|
|   192.168.93.141    |   node1  |
|   192.168.93.142    |   node2  |
|   192.168.93.143    |   node3  |
|   192.168.93.144    |   node4  |

## Host list configuration
```shell
meepo:
  vars:
    workpath: "/opt/block/meepo" # Remote machine working directory

  hosts:
    node1: # Remote node name, this name is related to the configuration of each node
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

# Task List
The specific task process is as follows：
1. Install service：Install docker, docker-compose, firewalld and other services
2. Configuration Environment：Development port, configure docker image, etc.
4. Install meepo node
5. Add enode node：Form a private network
