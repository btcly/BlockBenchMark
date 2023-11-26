#!/bin/bash

# 加入通道
peer channel join -b fabricdata/package/mychannel.block


# org1和org2的peer0更新锚点
if echo "$CORE_PEER_ADDRESS" | grep -q "peer0"; then
    orgnum=$(echo "$CORE_PEER_ADDRESS"|awk -F "[g.]" '{print $3}')
    peer channel update -o orderer1.example.com:7050 -c mychannel -f ./channel-artifacts/Org${orgnum}MSPanchors.tx --tls --cafile ${CORE_PEER_ORDER_TLS_CERT_FILE}
fi


#######################[ SACC ]#####################
# 安装链码
peer lifecycle chaincode install fabricdata/package/sacc.tar.gz

# 记录链码ID
sacc_id=$(peer lifecycle chaincode queryinstalled|grep -Eo 'sacc_1:[0-9a-z]+')


# 如果包含 peer0，则执行的命令
if echo "$CORE_PEER_ADDRESS" | grep -q "peer0"; then
    # 所有每个组织的peer中都要执行一次，几个组织执行几次
    peer lifecycle chaincode approveformyorg --channelID mychannel --name sacc --version 1.0 --init-required --package-id ${sacc_id} --sequence 1 --tls true --cafile  ${CORE_PEER_ORDER_TLS_CERT_FILE} # > chaincodemsg 2>&1 || true
    # 假设错误信息保存在变量error_msg中
    # error_msg=`cat chaincodemsg`
    

    # # 检查error_msg中是否包含"error"
    # if [[ $error_msg == *"Error"* ]]; then
    #     # 使用awk命令提取"new definition must be sequence x"中的x值
    #     seqnum=$(echo "$error_msg" | awk -F 'new definition must be sequence ' '{print $2}')
        
    #     # 输出提取到的值
    #     echo "提取到的序列号为: $seqnum"
    #     peer lifecycle chaincode approveformyorg --channelID mychannel --name sacc --version 1.0 --init-required --package-id ${sacc_id} --sequence $seqnum --tls true --cafile  ${CORE_PEER_ORDER_TLS_CERT_FILE} # > /dev/null 2>&1 || true
    # fi
fi
#peer lifecycle chaincode checkcommitreadiness --channelID mychannel --name sacc --version 1.0 --init-required --sequence 1 --tls true --cafile $CORE_PEER_ORDER_TLS_CERT_FILE --output json

# peer lifecycle chaincode queryinstalled
#######################[ SACC END]#####################

#######################[ KVSTORE ]#####################
# 安装链码
peer lifecycle chaincode install fabricdata/package/kvstore.tar.gz

# 记录链码ID
kvstore_id=$(peer lifecycle chaincode queryinstalled|grep -Eo 'kvstore_1:[0-9a-z]+')

# 如果包含 peer0，则执行的命令
if echo "$CORE_PEER_ADDRESS" | grep -q "peer0"; then
    # 所有每个组织的peer中都要执行一次，几个组织执行几次
    peer lifecycle chaincode approveformyorg --channelID mychannel --name kvstore --version 1.0 --init-required --package-id ${kvstore_id} --sequence 1 --tls true --cafile  ${CORE_PEER_ORDER_TLS_CERT_FILE} # > chaincodemsg 2>&1 || true
    # 假设错误信息保存在变量error_msg中
    # error_msg=`cat chaincodemsg`

    # # 检查error_msg中是否包含"error"
    # if [[ $error_msg == *"Error"* ]]; then
    #     # 使用awk命令提取"new definition must be sequence x"中的x值
    #     seqnum=$(echo "$error_msg" | awk -F 'new definition must be sequence ' '{print $2}')
        
    #     # 输出提取到的值
    #     echo "提取到的序列号为: $seqnum"
    #     peer lifecycle chaincode approveformyorg --channelID mychannel --name kvstore --version 1.0 --init-required --package-id ${kvstore_id} --sequence $seqnum --tls true --cafile  ${CORE_PEER_ORDER_TLS_CERT_FILE} # > /dev/null 2>&1 || true
    # fi

fi
#peer lifecycle chaincode checkcommitreadiness --channelID mychannel --name kvstore --version 1.0 --init-required --sequence 1 --tls true --cafile $CORE_PEER_ORDER_TLS_CERT_FILE --output json

# peer lifecycle chaincode queryinstalled
#######################[ KVSTORE END]#####################

#######################[ SMALLBANK ]#####################
# 安装链码
peer lifecycle chaincode install fabricdata/package/smallbank.tar.gz

# 记录链码ID
smallbank_id=$(peer lifecycle chaincode queryinstalled|grep -Eo 'smallbank_1:[0-9a-z]+')

# 如果包含 peer0，则执行的命令
if echo "$CORE_PEER_ADDRESS" | grep -q "peer0"; then
    # 所有每个组织的peer中都要执行一次，几个组织执行几次
    peer lifecycle chaincode approveformyorg --channelID mychannel --name smallbank --version 1.0 --init-required --package-id ${smallbank_id} --sequence 1 --tls true --cafile  ${CORE_PEER_ORDER_TLS_CERT_FILE} # > chaincodemsg 2>&1 || true
    # 假设错误信息保存在变量error_msg中
    # error_msg=`cat chaincodemsg` 
    # # 检查error_msg中是否包含"error"
    # if [[ $error_msg == *"Error"* ]]; then
    #     # 使用awk命令提取"new definition must be sequence x"中的x值
    #     seqnum=$(echo "$error_msg" | awk -F 'new definition must be sequence ' '{print $2}')

    #     # 输出提取到的值
    #     echo "提取到的序列号为: $seqnum"
    #     peer lifecycle chaincode approveformyorg --channelID mychannel --name smallbank --version 1.0 --init-required --package-id ${smallbank_id} --sequence $seqnum --tls true --cafile  ${CORE_PEER_ORDER_TLS_CERT_FILE} # > /dev/null 2>&1 || true
    # fi
fi
#peer lifecycle chaincode checkcommitreadiness --channelID mychannel --name smallbank --version 1.0 --init-required --sequence 1 --tls true --cafile $CORE_PEER_ORDER_TLS_CERT_FILE --output json

# peer lifecycle chaincode queryinstalled
#######################[ SMALLBANK END]#####################