#!/bin/bash

# 用于统一生成初始信息
# 并统一复制到orderer和peer文件夹中

rm -rf channel-artifacts crypto-config
rm -rf ../peers/channel-artifacts ../peers/crypto-config
# 生成证书文件
./bin/cryptogen generate --config=./crypto-config.yaml
#   生成创世区块：
./bin/configtxgen -profile SampleMultiNodeEtcdRaft -channelID fabric-cluster-channel -outputBlock ./channel-artifacts/genesis.block
# 创建通道配置信息：
./bin/configtxgen -profile TwoOrgsChannel -outputCreateChannelTx ./channel-artifacts/mychannel.tx -channelID mychannel
#   为 Org1 定义锚节点：
./bin/configtxgen -profile TwoOrgsChannel -outputAnchorPeersUpdate ./channel-artifacts/Org1MSPanchors.tx -channelID mychannel -asOrg Org1MSP
#   为 Org2 定义锚节点：
./bin/configtxgen -profile TwoOrgsChannel -outputAnchorPeersUpdate ./channel-artifacts/Org2MSPanchors.tx -channelID mychannel -asOrg Org2MSP

cp channel-artifacts crypto-config ../peers/ -rf
