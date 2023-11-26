# 介绍
此为client端运行的参数配置

# client配置

```shell

{
    "sever_ip" : ["127.0.0.1:5555"], # 连接server的IP，支持多个
    "items" : # 当前client支持区块的种类
    [
        {
            "type" : "client",
            "blockname" : "Fabric", # 压测区块的名称
            "chaincode" : "sacc", # 需要压测的合约
            "functionParams" : ["get:1", "set:2"], # 压测合约的函数，组合含义：函数名称+参数个数
            "qps" : 100, # 压测的QPS
            "open" : true # 压测是否开启，true为开启
        },
        {
            "type" : "client",
            "blockname" : "Fabric",
            "chaincode" : "smallbank",
            "functionParams" : ["Query:1", "Almagate:2", "GetBalance:1", "UpdateBalance:2", "UpdateSaving:2", "SendPayment:3", "WriteCheck:2"],
            "qps" : 8,
            "open" : false
        },
        {
            "type" : "client",
            "blockname" : "Fabric",
            "chaincode" : "kvstore",
            "functionParams" : ["Write:2", "Del:1", "Read:1"],
            "qps" : 16,
            "open" : false
        },
        {
            "type" : "client",
            "blockname" : "ETHPersonal",
            "chaincode" : "store",
            "functionParams" : ["setItem:2", "getItem:1", "versionContract:0"],
            "qps" : 0.5,
            "open" : false
        },
        {
            "type" : "client",
            "blockname" : "ETHPersonal",
            "chaincode" : "kvstore",
            "functionParams" : ["get:1", "set:2"],
            "qps" : 1,
            "open" : false
        },
        {
            "type" : "client",
            "blockname" : "ETHPersonal",
            "chaincode" : "smallbank",
            "functionParams" : ["almagate:2","getBalance:1","updateBalance:2","updateSaving:2","sendPayment:3","writeCheck:2"],
            "qps" : 1,
            "open" : false
        },
        {
            "type" : "client",
            "blockname" : "MeepoPersonal",
            "chaincode" : "store",
            "functionParams" : ["setItem:2", "getItem:1", "versionContract:0"],
            "qps" : 1,
            "open" : false
        },
        {
            "type" : "client",
            "blockname" : "MeepoPersonal",
            "chaincode" : "kvstore",
            "functionParams" : ["get:1", "set:2"],
            "qps" : 1,
            "open" : false
        },
        {
            "type" : "client",
            "blockname" : "MeepoPersonal",
            "chaincode" : "smallbank",
            "functionParams" : ["almagate:2","getBalance:1","updateBalance:2","updateSaving:2","sendPayment:3","writeCheck:2"],
            "qps" : 1,
            "open" : false
        }
    ]
}
```