# Introduction
This is the parameter configuration for client operation

# client configuration

```shell

{
    "sever_ip" : ["127.0.0.1:5555"], # IP to connect to the server, supports multiple
    "items" : # The types of blocks currently supported by the client
    [
        {
            "type" : "client",
            "blockname" : "Fabric", # The name of the stress test block
            "chaincode" : "sacc", # Contracts requiring stress testing
            "functionParams" : ["get:1", "set:2"], # The function of the stress testing contract, the combination meaning: function name and number of parameters
            "qps" : 100, # QPS of stress test
            "open" : true # Whether the stress test is enabled, true means enabled
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
