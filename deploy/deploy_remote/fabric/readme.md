# Introduction
This is the Fabric remote installation directory. You need to copy the program to other machines through ansible-playbook to automatically deploy the installation task and install the contract.

# Document introduction
Fabric related files are unified into fabric_pack

The binary version used by fabric is Fabric-2.2. The certificate file and channel file have been generated in advance according to the configuration file in the current directory.

Introduction to important files and directories：
```shell
./fabric_pack/orders/create_crypto_config.sh：Generate Fabric's channel and certificate files according to the configuration and synchronize them to the peers directory
```
```shell
./fabric_pack/orders/crypto-config和./fabric_pack/peers/crypto-config： Same as certificate file
```

```shell
./fabric_pack/peers/chaincode：Chaincode storage directory
```
```shell
./fabric_pack/peers/fabricdata/deploy：Environment configuration and chaincode installation script
```

# Execute a command
Because relative paths are used, operations need to be done in the current directory.
```shell
ansible-playbook -i ansible_hosts_fabric.yaml ansible_deploy_fabric.yaml
```
# Host list
> When using ansible_hosts_fabric.yaml, please make changes strictly in accordance with the Fabric cluster.

## Fabric structure of this project
Please refer to this project[Configuration file](./fabric_pack/orders/crypto-config.yaml)，Three orders are built（raft algorithm），two org，Two peers per organization
| ip address       |     Organization  | Node name  |Node domain name  |
| :----------- |:------------:|:------------:|:------------:|
|   192.168.93.153    |   order  |orderer1|orderer1.example.com|
|   192.168.93.154    |   order  |orderer2|orderer2.example.com|
|   192.168.93.155    |   order  |orderer3|orderer3.example.com|
|   192.168.93.156    |   org1  |peer0.org1|peer0.org1.example.com|
|   192.168.93.147    |   org1  |peer1.org1|peer1.org1.example.com|
|   192.168.93.148    |   org2  |peer0.org2|peer0.org2.example.com|
|   192.168.93.149    |   org2  |peer1.org2|peer1.org2.example.com|

## Host list configuration
```shell
fabric:
  vars:
    workpath: "/opt/block/fabric" # Remote machine working directory
  
  children:
    orders: # Set the order name according to the Fabric configuration file
      hosts:
        orderer1: # The name must be consistent with the prefix of the domain name
          ansible_host: 192.168.93.153
          ansible_user: test
        orderer2:
          ansible_host: 192.168.93.154
          ansible_user: test
        orderer3:
          ansible_host: 192.168.93.155
          ansible_user: test

    peers: # Set the peer name according to the Fabric configuration file
      hosts:
        peer0.org1: # The name must be consistent with the prefix of the domain name
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

    peertool: # Set up a separate peer for chain code packaging and channel creation. Generally, choose any one from the peer.
      hosts:
        peer0.org1:
          ansible_host: 192.168.93.156
          ansible_user: test

```

# task list
The specific task process is as follows：
1. install service：Install docker, docker-compose, firewalld and other services
2. Configuration Environment：Development port, configure docker image, etc.
3. Install node：Used for reporting machine status data
4. Install orders
5. Install peers
6. Install chaincodes
