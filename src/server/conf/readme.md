# 介绍
此为server端运行的参数配置

# server配置
配置文件路径为：./server/config_server.json

配置文件讲解
```shell
{
    "type" : "server", # 服务的类型
    "mysql" : { # 安装mysql的信息
        "userName" : "root", # 用户名
        "password" : "123456",  # 密码
        "ipAddrees" : "127.0.0.1", # 数据库的IP地址
        "port" : "3306", # 数据库的端口
        "dbName" : "block", # 数据库的名称
        "charset" : "utf8"
    },
    "redis" : # redis的信息
    {
        "url" : "127.0.0.1:6379", # redis的IP地址和端口
        "passwd" : "" # 密码
    },
    "sever_ip" :":5555", # 本地服务的端口号
    "nodes" : # 本地服务支持的区块链种类
    [
        {
            "name" : "ETHPersonal", # 名称，server和client保持一致
            "nodeurl" : "ws://192.168.93.137:8546", # 区块链的地址
            "conf":"", # 配置信息
            "chainid":"1234", # chainid
            "open" : true # 是否开启
        },
        {
            "name" : "MeepoPersonal",
            "nodeurl" : "http://192.168.93.141:8545",
            "conf":"",
            "chainid":"1234",
            "open" : true
        },
        {
            "name" : "Fabric",
            "nodeurl" : "",
            "conf":"",
            "open" : true
        }
    ]
}

```

# 区块链配置

## ETH && Meepo
文件目录：./eth/keystore

存放地址私钥信息
## Fabric
证书文件：./fabric/crypto-config


fabric-sdk-go配置文件：./fabric/fabric_config.yaml

