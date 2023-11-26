#!/bin/bash

# 提交链码

# 可能存在序号错误

# 提交 SACC 链码定义，在 org1 或者 org2 上均可
peer lifecycle chaincode commit -o orderer1.example.com:7050 --channelID mychannel --name sacc --version 1.0 --sequence 1 --init-required --tls true --cafile ${CORE_PEER_ORDER_TLS_CERT_FILE} --peerAddresses peer0.org1.example.com:7051 --tlsRootCertFiles ${CORE_PEER0_ORG1_TLS_CA_CRT} --peerAddresses peer0.org2.example.com:7051 --tlsRootCertFiles ${CORE_PEER0_ORG2_TLS_CA_CRT}  # > chaincodemsg 2>&1 || true
# # 假设错误信息保存在变量error_msg中
# error_msg=`cat chaincodemsg`
# # 检查error_msg中是否包含"error"
# if [[ $error_msg == *"Error"* ]]; then
#     # 使用awk命令提取"new definition must be sequence x"中的x值
#     seqnum=$(echo "$error_msg" | awk -F 'new definition must be sequence ' '{print $2}')
    
#     # 输出提取到的值
#     echo "提取到的序列号为: $seqnum"
#     peer lifecycle chaincode commit -o orderer1.example.com:7050 --channelID mychannel --name sacc --version 1.0 --sequence $seqnum --init-required --tls true --cafile ${CORE_PEER_ORDER_TLS_CERT_FILE} --peerAddresses peer0.org1.example.com:7051 --tlsRootCertFiles ${CORE_PEER0_ORG1_TLS_CA_CRT} --peerAddresses peer0.org2.example.com:7051 --tlsRootCertFiles ${CORE_PEER0_ORG2_TLS_CA_CRT} # > /dev/null 2>&1 || true
# fi

# 提交 KVSTORE 链码定义，在 org1 或者 org2 上均可
peer lifecycle chaincode commit -o orderer1.example.com:7050 --channelID mychannel --name kvstore --version 1.0 --sequence 1 --init-required --tls true --cafile ${CORE_PEER_ORDER_TLS_CERT_FILE} --peerAddresses peer0.org1.example.com:7051 --tlsRootCertFiles ${CORE_PEER0_ORG1_TLS_CA_CRT} --peerAddresses peer0.org2.example.com:7051 --tlsRootCertFiles ${CORE_PEER0_ORG2_TLS_CA_CRT}  # > chaincodemsg 2>&1 || true
# 假设错误信息保存在变量error_msg中
# error_msg=`cat chaincodemsg`
# # 检查error_msg中是否包含"error"
# if [[ $error_msg == *"Error"* ]]; then
#     # 使用awk命令提取"new definition must be sequence x"中的x值
#     seqnum=$(echo "$error_msg" | awk -F 'new definition must be sequence ' '{print $2}')
    
#     # 输出提取到的值
#     echo "提取到的序列号为: $seqnum"
#     peer lifecycle chaincode commit -o orderer1.example.com:7050 --channelID mychannel --name kvstore --version 1.0 --sequence $seqnum --init-required --tls true --cafile ${CORE_PEER_ORDER_TLS_CERT_FILE} --peerAddresses peer0.org1.example.com:7051 --tlsRootCertFiles ${CORE_PEER0_ORG1_TLS_CA_CRT} --peerAddresses peer0.org2.example.com:7051 --tlsRootCertFiles ${CORE_PEER0_ORG2_TLS_CA_CRT} # > /dev/null 2>&1 || true
# fi

# 提交 SMALLBANK 链码定义，在 org1 或者 org2 上均可
peer lifecycle chaincode commit -o orderer1.example.com:7050 --channelID mychannel --name smallbank --version 1.0 --sequence 1 --init-required --tls true --cafile ${CORE_PEER_ORDER_TLS_CERT_FILE} --peerAddresses peer0.org1.example.com:7051 --tlsRootCertFiles ${CORE_PEER0_ORG1_TLS_CA_CRT} --peerAddresses peer0.org2.example.com:7051 --tlsRootCertFiles ${CORE_PEER0_ORG2_TLS_CA_CRT}  # > chaincodemsg 2>&1 || true
# # 假设错误信息保存在变量error_msg中
# error_msg=`cat chaincodemsg`
# # 检查error_msg中是否包含"error"
# if [[ $error_msg == *"Error"* ]]; then
#     # 使用awk命令提取"new definition must be sequence x"中的x值
#     seqnum=$(echo "$error_msg" | awk -F 'new definition must be sequence ' '{print $2}')
    
#     # 输出提取到的值
#     echo "提取到的序列号为: $seqnum"
#     peer lifecycle chaincode commit -o orderer1.example.com:7050 --channelID mychannel --name smallbank --version 1.0 --sequence $seqnum --init-required --tls true --cafile ${CORE_PEER_ORDER_TLS_CERT_FILE} --peerAddresses peer0.org1.example.com:7051 --tlsRootCertFiles ${CORE_PEER0_ORG1_TLS_CA_CRT} --peerAddresses peer0.org2.example.com:7051 --tlsRootCertFiles ${CORE_PEER0_ORG2_TLS_CA_CRT} # > /dev/null 2>&1 || true
# fi

# peer lifecycle chaincode querycommitted --channelID mychannel --name sacc --cafile ${CORE_PEER_ORDER_TLS_CERT_FILE}

peer chaincode invoke  -C mychannel -n sacc -o orderer1.example.com:7050 --tls true --cafile ${CORE_PEER_ORDER_TLS_CERT_FILE} --peerAddresses peer0.org1.example.com:7051 --tlsRootCertFiles ${CORE_PEER0_ORG1_TLS_CA_CRT} --peerAddresses peer0.org2.example.com:7051 --tlsRootCertFiles ${CORE_PEER0_ORG2_TLS_CA_CRT} --isInit -c '{"Args":["a","100"]}'
peer chaincode invoke  -C mychannel -n kvstore -o orderer1.example.com:7050 --tls true --cafile ${CORE_PEER_ORDER_TLS_CERT_FILE} --peerAddresses peer0.org1.example.com:7051 --tlsRootCertFiles ${CORE_PEER0_ORG1_TLS_CA_CRT} --peerAddresses peer0.org2.example.com:7051 --tlsRootCertFiles ${CORE_PEER0_ORG2_TLS_CA_CRT} --isInit -c '{"Args":[]}'
peer chaincode invoke  -C mychannel -n smallbank -o orderer1.example.com:7050 --tls true --cafile ${CORE_PEER_ORDER_TLS_CERT_FILE} --peerAddresses peer0.org1.example.com:7051 --tlsRootCertFiles ${CORE_PEER0_ORG1_TLS_CA_CRT} --peerAddresses peer0.org2.example.com:7051 --tlsRootCertFiles ${CORE_PEER0_ORG2_TLS_CA_CRT} --isInit -c '{"Args":[]}'


# 查询数据
# peer chaincode invoke -o orderer1.example.com:7050 --tls true --cafile ${CORE_PEER_ORDER_TLS_CERT_FILE} --peerAddresses peer0.org1.example.com:7051 --tlsRootCertFiles ${CORE_PEER0_ORG1_TLS_CA_CRT} --peerAddresses peer0.org2.example.com:7051 --tlsRootCertFiles ${CORE_PEER0_ORG2_TLS_CA_CRT} --channelID mychannel --name sacc -c '{"function":"get","Args":["a"]}' --waitForEvent 
