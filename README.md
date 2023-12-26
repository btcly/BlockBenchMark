
# Introduction
This is a blockchain stress testing tool developed using Golang and supports EOV,OEV,EV and Sharding.
# Architecture Diagram
![image](https://github.com/btcly/BlockBenchMark/assets/51044388/7c19e4d4-7c0a-4ee4-99c1-d15148790a0b)



# Software Version
| name of software       |     version number  |
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
Use ansible-playbook to perform cluster management and automatic deployment through host lists and task lists. Before use, please perform password-free login and password-free sudo access to the control end and remote end in advance.

Deployment is mainly divided into [monitoring deployment](./deploy/deploy_local/readme.md) （mysql，prometheus，grafana ...）and [blockchain construction](./deploy/deploy_remote/readme.md)（Fabric，ETH private chain，Meepo）, which can quickly build a testing environment.

## Monitor
Installation command：
```shell
    cd deploy/deploy_local
    /bin/bash deploy.sh
```
The installation content is as follows：
* mysql：Used for stress measurement data storage, generally installed on the non-server side
* grafana：Used for data display, generally installed on the non-server side
* prometheus：Used to collect machine data and monitor machine status. It is generally installed on the non-server side.
* node_exporter：Installed on the side of the monitored machine, it reports machine status data to prometheus. Generally, it needs to be installed on all machines in the cluster.

### Grafama Configuration
After the installation is complete, the address is：http://XXXX:3000

The login user and password are：admin、admin

The data of grafana has been configured, and the path is：./grafana_data/grafana.db

> **Change the IP address of mysql and prometheus data sources**

## Fabric Build
Excute a command：
```shell
    cd deploy/deploy_remote/fabric
    ansible-playbook -i ansible_hosts_fabric.yaml ansible_deploy_fabric.yaml
```
> Please modify it as needed before using it [Host list](./deploy/deploy_remote/fabric/ansible_hosts_fabric.yaml)：Used to manage Fabric cluster node allocation，[task list](./deploy/deploy_remote/fabric/ansible_deploy_fabric.yaml)：Normally no changes are required. For host list, task list and Fabric construction details, please click [here](./deploy/deploy_remote/fabric/readme.md)。

## Geth Build
Excute a command：
```shell
    cd deploy/deploy_remote/geth
    ansible-playbook -i ansible_hosts_geth.yaml ansible_deploy_geth.yaml
```
> Please modify it as needed before using it [Host list](./deploy/deploy_remote/geth/ansible_hosts_geth.yaml)：Used to manage Geth cluster node allocation，[task list](./deploy/deploy_remote/geth/ansible_deploy_geth.yaml)Normally no changes are required。For host list, task list and Geth build details, please click [here](./deploy/deploy_remote/geth/readme.md)。

## Meepo Build
Excute a command：
```shell
    cd deploy/deploy_remote/meepo
    ansible-playbook -i ansible_hosts_meepo.yaml ansible_deploy_meepo.yaml
```
> Please modify it as needed before using it [Host list](./deploy/deploy_remote/meepo/ansible_hosts_meepo.yaml)：Used to manage Geth cluster node allocation，[task list](./deploy/deploy_remote/meepo/ansible_deploy_meepo.yaml)Normally no changes are required。For host list, task list and Geth build details, please click [here](./deploy/deploy_remote/meepo/readme.md)。

# Build And Run (Ubuntu)
The stress testing system is mainly composed of server and client, and supports distributed deployment.
The deployment of ETH, Fabric, and Meepo test environments is also part of the system.
> All tests and operations are based on Ubuntu22.04. For the golang version, please run it according to the version recommended above.
## Server
server compile：
```shell
    cd src/server
    go build server.go
```
server run：
```shell
    ./server -log_dir=log
```
> conf directory and server binary are required to function properly

Server operation requires reading the program running configuration and blockchain-related configuration. Please see the detailed explanation of the relevant configuration [here](./src/server/conf/readme.md)。

After building Fabric/ETH/Meepo, if you use a domain name to access, you generally need to change the mapping between domain name and IP. Please modify the file /etc/hosts
```shell
192.168.93.156 peer0.org1.example.com
192.168.93.146 peer1.org1.example.com
192.168.93.147 peer0.org2.example.com
192.168.93.148 peer1.org2.example.com
192.168.93.153 orderer1.example.com
192.168.93.154 orderer2.example.com
192.168.93.155 orderer3.example.com
```

## Client
client compile：
```shell
    cd src/client
    go build client.go
```
client run：
```shell
    ./client -log_dir=log
```
> Requires config_client.json file and client binary to run

For a detailed explanation of the configuration file running by Client, please see [here](./src/client/readme.md)

# Others 
> **For Fabric, Geth, and Meepo, please check the configuration file carefully.**
