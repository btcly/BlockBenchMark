# Introduction
This is the parameter configuration for server-side operation

# server configuration
The configuration file path is：./server/config_server.json

Configuration file explanation
```shell
{
    "type" : "server", # Type of service
    "mysql" : { # Information about installing mysql
        "userName" : "root", # username
        "password" : "123456",  # password
        "ipAddrees" : "127.0.0.1", # Database IP address
        "port" : "3306", # Database port
        "dbName" : "block", # database name
        "charset" : "utf8"
    },
    "redis" : # redis information
    {
        "url" : "127.0.0.1:6379", # redis IP address and port
        "passwd" : "" # password
    },
    "sever_ip" :":5555", # The port number of the local service
    "nodes" : # Blockchain types supported by local services
    [
        {
            "name" : "ETHPersonal", # Name, server and client remain consistent
            "nodeurl" : "ws://192.168.93.137:8546", # Blockchain address
            "conf":"", # Configuration information
            "open" : true # Whether to turn on
        },
        {
            "name" : "MeepoPersonal",
            "nodeurl" : "http://192.168.93.141:8545",
            "conf":"",
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

# Blockchain configuration

## ETH && Meepo
File Directory：./eth/keystore

Store address private key information
## Fabric
certificate file：./fabric/crypto-config


fabric-sdk-go Configuration file：./fabric/fabric_config.yaml

