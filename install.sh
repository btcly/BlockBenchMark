#!/bin/bash
sudo snap install solc
snap install go --classic
sudo apt install pkg-config protoc* redis -y

sudo apt install -y ansible

# 编译PB
export GO_PATH=~/go
export PATH=$PATH:/$GO_PATH/bin

go install google.golang.org/protobuf/cmd/protoc-gen-go@latest
go install google.golang.org/grpc/cmd/protoc-gen-go-grpc@latest

protoc -I=./protocs/*.proto --plugin=/usr/bin/protoc-gen-go --go_out=./protocs --plugin=~/go/bin/protoc-gen-go-grpc --go-grpc_out=./protocs

# 修改为上海时间
timedatectl
sudo timedatectl set-timezone Asia/Shanghai
timedatectl


# ETH测试网，获取测试币
# https://goerli-faucet.pk910.de/
# https://sepolia-faucet.pk910.de/

# goerli搭建
# git clone https://github.com/prysmaticlabs/prysm.git
# cd prysm && ./prysm.sh beacon-chain generate-auth-secret
# geth --goerli --http --http.addr "0.0.0.0" --http.api eth,net,engine,admin --gcmode archive  --authrpc.addr localhost --authrpc.port 8551 --authrpc.vhosts localhost --authrpc.jwtsecret /home/masnail/code/prysm/jwt.hex 
# ./prysm.sh beacon-chain --execution-endpoint=http://localhost:8551 --jwt-secret=./jwt.hex --accept-terms-of-use #--checkpoint-sync-url="https://goerli.infura.io/ws/v3/d2da845d8fd04ca898226e28fdba1812" --genesis-beacon-api-url="https://goerli.infura.io/ws/v3/d2da845d8fd04ca898226e28fdba1812"

# "endpoint" : "https://goerli.infura.io/v3/d2da845d8fd04ca898226e28fdba1812"
# https://goerli.infura.io/v3/a14330eaff5545fd860d10c24020956b
# "endpoint" : "wss://sepolia.infura.io/ws/v3/d2da845d8fd04ca898226e28fdba1812"


#fabric
# wget https://github.com/hyperledger/fabric/releases/download/v1.2.0/hyperledger-fabric-linux-amd64-1.2.0.tar.gz


sudo apt install python3-pip -y
pip3 install sqlalchemy pymysql -i https://pypi.tuna.tsinghua.edu.cn/simple
