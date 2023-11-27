package meeporpc

import (
	// "bytes"
	// "encoding/json"
	// "errors"
	// "net/http"
	"blcokbenchmark/src/block/eth/contrace/kvstore"
	"blcokbenchmark/src/block/eth/contrace/smallbank"
	"blcokbenchmark/src/block/eth/contrace/store"
	"blcokbenchmark/src/redis"
	"bytes"
	"context"
	"crypto/ecdsa"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"math/big"
	"math/rand"
	"net/http"
	"os"
	"path/filepath"
	"strconv"
	"strings"
	"time"

	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/accounts/keystore"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/ethereum/go-ethereum/rpc"
	"github.com/golang/glog"
)

// TODO 减少日志的打印
// TODO 支持自定义chainid

type RPCClient struct {
	ethClient *ethclient.Client
	rpcClient *rpc.Client
	blockName string
	nodeUrl   string
}

type ContractTrans struct {
	fromAddress  string
	contractName string
	contractHex  string
	isCheck      bool
}

var (
	gkeystore_path   = "./conf/block/eth/keystore/"
	gkeystore_passwd = "123456"
	gchainid         = big.NewInt(123456) //  启动后从区块链上获取
	ggasLimit        = uint64(300000)
	ggasPrice        = big.NewInt(0)
)

// 初始化ETH
func NewRPCClient(nodeUrl, blockName string) *RPCClient {
	eth_client, err := ethclient.Dial(nodeUrl)
	if err != nil {
		glog.Info(fmt.Sprintf("conn nodeurl:%s, err:%s.\n", nodeUrl, err))
		return nil
	}

	rpc_client, _ := rpc.DialContext(context.Background(), nodeUrl)

	rand.Seed(time.Now().UnixNano())

	rpcClient := &RPCClient{}
	rpcClient.ethClient = eth_client
	rpcClient.rpcClient = rpc_client
	rpcClient.blockName = blockName
	rpcClient.nodeUrl = nodeUrl

	rpcClient.DoPost()
	glog.Info(fmt.Sprintf("Meepo gchainid=", gchainid))

	loadKeystore(rpcClient, blockName)

	return rpcClient
}

// 私钥转为string
func privateKeyToString(privateKey *ecdsa.PrivateKey) string {
	privateKeyBytes := privateKey.D.Bytes()
	privateKeyHex := hex.EncodeToString(privateKeyBytes)
	return "0x" + privateKeyHex
}

// string转为私钥
func stringToPrivateKey(privateKeyStr string) (*ecdsa.PrivateKey, error) {
	// 移除 "0x" 前缀
	privateKeyHex := privateKeyStr[2:]

	// 解码十六进制字符串
	privateKeyBytes, err := hex.DecodeString(privateKeyHex)
	if err != nil {
		glog.Info(fmt.Sprintf("DecodeString privatekey:%s, err:%s.\n", privateKeyHex, err))
		return nil, err
	}

	privateKey := new(ecdsa.PrivateKey)
	privateKey.D = new(big.Int).SetBytes(privateKeyBytes)
	privateKey.PublicKey.Curve = crypto.S256() // 替换为您使用的曲线
	privateKey.PublicKey.X, privateKey.PublicKey.Y = privateKey.PublicKey.Curve.ScalarBaseMult(privateKey.D.Bytes())

	return privateKey, nil
}

func loadKeystore(client *RPCClient, blockName string) {
	glog.Info(fmt.Sprintf("load keystore file from ", blockName))
	var files []string
	contractHexs := []ContractTrans{}
	err := filepath.Walk(gkeystore_path, func(path string, info os.FileInfo, err error) error {
		// glog.Info(fmt.Sprintf(":---->path:%s.\n", path))
		if path != gkeystore_path {
			files = append(files, path)
		}
		return nil
	})
	if err != nil {
		glog.Info(fmt.Sprintf("load keystore file failed, err:%s.\n", err))
		return
	}
	for _, file := range files {
		//  该路径是二进制执行的相对位置
		keyJson, err := ioutil.ReadFile(file)
		if err != nil {
			glog.Info(fmt.Sprintf("read file:%s, failed, err:%s.\n", file, err))
			continue
		}

		key, err := keystore.DecryptKey(keyJson, gkeystore_passwd)
		if err != nil {
			glog.Info(fmt.Sprintf("parse json failed, err:%s.\n", err))
			continue
		}

		// glog.Info(fmt.Sprintf("keystore key:%v.\n", key)

		publicKey := key.PrivateKey.Public()
		publicKeyECDSA, ok := publicKey.(*ecdsa.PublicKey)
		if !ok {
			glog.Info(fmt.Sprintf("error casting public key to ECDSA"))
			continue
		}

		fromAddress := crypto.PubkeyToAddress(*publicKeyECDSA)

		// 对fromAddr加锁
		trylock := redis.Redisclient.SetNonceLock(blockName, fromAddress.Hex())
		// 加锁失败
		if !trylock {
			continue
		}
		//  设置地址的私钥
		privateKeyBytes := privateKeyToString(key.PrivateKey)
		redis.Redisclient.SetPrivateInitToRedis(blockName, fromAddress.Hex(), string(privateKeyBytes))

		client.GetBlance(fromAddress.Hex())

		gasPrice, err := client.ethClient.SuggestGasPrice(context.Background())
		if err != nil {
			glog.Info(fmt.Sprintf("suggest gas price, err:", err))
			continue
		}

		glog.Info(fmt.Sprintf("----->load meepo address:%s.", fromAddress.Hex()))

		ggasPrice = big.NewInt(0).Add(gasPrice, big.NewInt(int64(rand.Intn(300000))))

		//  STORE
		storehex := client.DeployContractTransaction(fromAddress.Hex(), store.StoreBin, ggasLimit, ggasPrice)
		contractHexs = append(contractHexs, ContractTrans{
			fromAddress:  fromAddress.Hex(),
			contractName: "store",
			contractHex:  storehex,
			isCheck:      false,
		})

		//  SMALLBANK
		kvstorehex := client.DeployContractTransaction(fromAddress.Hex(), store.StoreBin, ggasLimit, ggasPrice)
		contractHexs = append(contractHexs, ContractTrans{
			fromAddress:  fromAddress.Hex(),
			contractName: "kvstore",
			contractHex:  kvstorehex,
			isCheck:      false,
		})

		//  KVSTORE
		smallbankhex := client.DeployContractTransaction(fromAddress.Hex(), store.StoreBin, ggasLimit, ggasPrice)
		contractHexs = append(contractHexs, ContractTrans{
			fromAddress:  fromAddress.Hex(),
			contractName: "smmallbank",
			contractHex:  smallbankhex,
			isCheck:      false,
		})

		// 对fromAddr解锁
		redis.Redisclient.SetNonceunLock(blockName, fromAddress.Hex())
	}

	//  检查安装的合约
	for i := 0; i < 5; i++ {
		for _, item := range contractHexs {
			if item.isCheck == false {
				conAddress := client.GetContractByTransaction(item.contractHex)
				if conAddress != "" {
					redis.Redisclient.SetContracAddresstToRedis(blockName, item.contractName, item.fromAddress, conAddress)
					item.isCheck = true
				}
			}
		}
		time.Sleep(2)
	}

	glog.Info(fmt.Sprintf("loading Meepo contract Successed.\n"))

}

// 统一对外提供合约调用的接口
func (client *RPCClient) ExecContract(ChaincodeID, ChaincodeFunc string, Params []string) string {
	contractaddr, fromaddr, _ := redis.Redisclient.GetRandomFieldContract(client.blockName, ChaincodeID)
	if contractaddr == "" || fromaddr == "" {
		glog.Info(fmt.Sprintf("GetRandomFieldContract not found address.\n"))
		return ""
	}
	var hex = ""
	// 随机选择选择一个地址的信息
	if ChaincodeID == "store" {
		hex = client.ContractTransaction(fromaddr, contractaddr, ChaincodeFunc, store.StoreABI, Params, ggasLimit, ggasPrice)
	} else if ChaincodeID == "kvstore" {
		hex = client.ContractTransaction(fromaddr, contractaddr, ChaincodeFunc, kvstore.KvstoreABI, Params, ggasLimit, ggasPrice)
	} else if ChaincodeID == "smallbank" {
		hex = client.ContractTransaction(fromaddr, contractaddr, ChaincodeFunc, smallbank.SmallbankABI, Params, ggasLimit, ggasPrice)
	} else {
		glog.Info(fmt.Sprintf("not support chaincideid [%s], please check.\n", ChaincodeID))
	}
	return hex
}

// /////////////////////////////////////// JSON RPC ////////////////////////////////////////

func (client *RPCClient) DoPost() {
	var networkid string
	client.rpcClient.Call(&networkid, "net_version")
	glog.Info(fmt.Sprintf("--->meepo networkid:", networkid))

	var chainid string
	client.rpcClient.Call(&chainid, "eth_chainId")
	glog.Info(fmt.Sprintf("--->meepo chainid:", chainid))
	if chainid != "" {
		gchainid.SetString(chainid[2:], 16)
	}

	var proto_version string
	client.rpcClient.Call(&proto_version, "eth_protocolVersion")
	glog.Info(fmt.Sprintf("--->meepo proto_version:", proto_version))

	if networkid == "" && proto_version == "" {
		glog.Exit("client failed, please check.\n")
	}
}

func (client *RPCClient) GetTranCountByTransID(transHex string) (int64, int64) {
	result := client.JsonRPCListCore("eth_getTransactionByHash", []interface{}{
		transHex,
	})
	if result == nil {
		glog.Info(fmt.Sprintf("getbalance error!"))
		return -1, -1
	}
	ret_trans := result.(map[string]interface{})
	if ret_trans == nil || ret_trans["blockNumber"] == nil {
		return -1, -1
	}

	blockNumber := ret_trans["blockNumber"]
	number, _ := strconv.ParseInt(blockNumber.(string)[2:], 16, 32)

	result = client.JsonRPCListCore("eth_getBlockByNumber", []interface{}{
		blockNumber,
		true,
	})

	if result == nil {
		return -1, -1
	}
	blockInfo, _ := result.(map[string]interface{})

	// UTC时间转为UTC+8
	blockTime, _ := strconv.ParseInt(blockInfo["timestamp"].(string)[2:], 16, 64)

	return blockTime * 1000, number

}

func (client *RPCClient) GetTranCountByBlockNumber(Number int64) (int64, int64) {
	blockNumber := strconv.FormatInt(Number, 16)
	blockNumber = "0x" + blockNumber
	result := client.JsonRPCListCore("eth_getBlockByNumber", []interface{}{
		blockNumber,
		true,
	})

	if result == nil {
		return -1, -1
	}
	blockInfo, _ := result.(map[string]interface{})

	// UTC时间转为UTC+8
	blockTime, _ := strconv.ParseInt(blockInfo["timestamp"].(string)[2:], 16, 64)

	trancount := len(blockInfo["transactions"].([]string))
	return blockTime, int64(trancount)
}

func (client *RPCClient) FindTranCountByTransID(transHex string) bool {

	result := client.JsonRPCListCore("eth_getTransactionByHash", []interface{}{
		transHex,
	})
	if result == nil {
		glog.Info(fmt.Sprintf("getbalance error!"))
		return false
	}
	ret_trans := result.(map[string]interface{})
	if ret_trans == nil || ret_trans["blockNumber"] == nil {
		return false
	}
	return true
}

func (client *RPCClient) GetBlance(fromAddress string) {
	result := client.JsonRPCListCore("eth_getBalance", []interface{}{
		fromAddress,
	})
	if result == nil {
		glog.Info(fmt.Sprintf(fromAddress, " getbalance error!"))
		return
	}
	res_str := result.(string)
	balance := new(big.Int)
	balance, _ = balance.SetString(res_str[2:], 16)
	basevalue := big.NewInt(1000000000000000000)
	balance = balance.Div(balance, basevalue)
	glog.Info(fmt.Sprintf("fromAddress:%s, balance:%d ETH.\n", fromAddress, balance))
}

func (client *RPCClient) DeployContractTransaction(fromAddress, contractByteCode string, gaslimit uint64, gasPrice *big.Int) string {

	// result := client.JsonRPCListCore("personal_unlockAccount", []interface{}{
	// 	fromAddress,
	// 	gkeystore_passwd,
	// 	nil,
	// })
	// glog.Info(fmt.Sprintf("Response:", result))
	result := client.JsonRPCMapCore("eth_sendTransaction", map[string]interface{}{

		"from":     fromAddress,
		"value":    "0x123", // 2441406250
		"data":     contractByteCode,
		"gas":      fmt.Sprintf("0x%x", gaslimit), // 30400
		"gasPrice": fmt.Sprintf("0x%x", gasPrice), // 10000000000000
	})

	// glog.Info(fmt.Sprintf("Response:", result))
	if result == nil {
		return ""
	}

	return result.(string)
}

func (client *RPCClient) ContractTransaction(fromAddress, contractAddress, funcName, contractABI string, Params []string, gaslimit uint64, gasPrice *big.Int) string {

	funcABI, _ := abi.JSON(strings.NewReader(contractABI))

	var data []byte
	if len(Params) > 0 {
		// 创建一个新的 interface{} 切片
		interfaceSlice := make([]interface{}, len(Params))

		// 遍历原始字符串切片，并将每个字符串转换为 interface{}
		for i, v := range Params {
			interfaceSlice[i] = v
		}

		data, _ = funcABI.Pack(funcName, interfaceSlice...)
	} else {
		data, _ = funcABI.Pack(funcName)

	}

	methodID := funcABI.Methods[funcName].ID

	input := append(methodID[:], data...)

	// result := client.JsonRPCListCore("personal_unlockAccount", []interface{}{
	// 	fromAddress,
	// 	gkeystore_passwd,
	// 	nil,
	// })
	// glog.Info(fmt.Sprintf("Response:", result))
	result := client.JsonRPCMapCore("eth_sendTransaction", map[string]interface{}{

		"from":     fromAddress,
		"to":       contractAddress,
		"value":    "0x123456789", // 2441406250
		"data":     fmt.Sprintf("0x%x", input),
		"gas":      fmt.Sprintf("0x%x", gaslimit), // 30400
		"gasPrice": fmt.Sprintf("0x%x", gasPrice), // 10000000000000
	})

	// glog.Info(fmt.Sprintf("Response:", result))
	if result == nil {
		return ""
	}

	return result.(string)
}

func (client *RPCClient) GetContractByTransaction(transHex string) string {
	result := client.JsonRPCListCore("eth_getTransactionReceipt", []interface{}{
		transHex,
	})
	if result == nil {
		return ""
	}
	ret_tran := result.(map[string]interface{})
	if ret_tran == nil || ret_tran["contractAddress"] == nil {
		return ""
	}
	return ret_tran["contractAddress"].(string)
}

func (client *RPCClient) JsonRPCMapCore(functionName string, params map[string]interface{}) interface{} {
	// 准备 JSON-RPC 请求
	payload := map[string]interface{}{
		"jsonrpc": "2.0",
		"method":  functionName,
		"params": []interface{}{
			params,
		},
		"id": 1,
	}

	// 序列化 JSON-RPC 请求数据
	requestBody, err := json.Marshal(payload)
	if err != nil {
		glog.Info(fmt.Sprintf("Error encoding JSON:", err))
		return nil
	}

	// 发送 HTTP POST 请求
	resp, err := http.Post(client.nodeUrl, "application/json", bytes.NewBuffer(requestBody))
	if err != nil {
		glog.Info(fmt.Sprintf("Error sending HTTP request:", err))
		return nil
	}
	defer resp.Body.Close()

	// 读取响应数据
	var result map[string]interface{}
	err = json.NewDecoder(resp.Body).Decode(&result)
	if err != nil {
		glog.Info(fmt.Sprintf("Error decoding JSON response:", err))
		return nil
	}

	value, exists := result["result"]
	if exists {
		return value
	} else {
		glog.Info(fmt.Sprintf("JsonRpc exec error, ", result))
		return nil
	}
}

func (client *RPCClient) JsonRPCListCore(functionName string, params []interface{}) interface{} {
	// 准备 JSON-RPC 请求
	payload := map[string]interface{}{
		"jsonrpc": "2.0",
		"method":  functionName,
		"params":  params,
		"id":      1,
	}

	// 序列化 JSON-RPC 请求数据
	requestBody, err := json.Marshal(payload)
	if err != nil {
		glog.Info(fmt.Sprintf("Error encoding JSON:", err))
		return nil
	}

	// 发送 HTTP POST 请求
	resp, err := http.Post(client.nodeUrl, "application/json", bytes.NewBuffer(requestBody))
	if err != nil {
		glog.Info(fmt.Sprintf("Error sending HTTP request:", err))
		return nil
	}
	defer resp.Body.Close()

	glog.Info(fmt.Sprintf("TEST Error decoding JSON response:", resp.Body))
	// 读取响应数据
	var result map[string]interface{}
	err = json.NewDecoder(resp.Body).Decode(&result)
	if err != nil {
		glog.Info(fmt.Sprintf("Error decoding JSON response:", err))
		return nil
	}

	value, exists := result["result"]
	if exists {
		return value
	} else {
		glog.Info(fmt.Sprintf("JsonRpc exec error, ", result))
		return nil
	}

}
