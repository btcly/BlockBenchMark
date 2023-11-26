#!/bin/bash
# 对链码进行打包
peer channel create -o orderer1.example.com:7050 -c mychannel -f ./channel-artifacts/mychannel.tx --tls true --cafile ${CORE_PEER_ORDER_TLS_CERT_FILE}

#  本地环境中设置依赖包
cd ./chaincode/go/sacc && go mod tidy && go mod vendor
cd ../kvstore && go mod tidy && go mod vendor
cd ../smallbank && go mod tidy && go mod vendor
cd ../../../

# 安装链码
peer lifecycle chaincode package sacc.tar.gz --path ./chaincode/go/sacc --lang golang --label sacc_1
peer lifecycle chaincode package kvstore.tar.gz --path ./chaincode/go/kvstore --lang golang --label kvstore_1
peer lifecycle chaincode package smallbank.tar.gz --path ./chaincode/go/smallbank --lang golang --label smallbank_1
cp *.tar.gz  *.block fabricdata/package/. -rf