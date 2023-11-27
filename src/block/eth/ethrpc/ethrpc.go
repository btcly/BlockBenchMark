package ethrpc

import (
	// "bytes"
	// "encoding/json"
	// "errors"
	// "net/http"
	"blcokbenchmark/src/block/eth/contrace/kvstore"
	"blcokbenchmark/src/block/eth/contrace/smallbank"
	"blcokbenchmark/src/block/eth/contrace/store"
	"blcokbenchmark/src/redis"
	"context"
	"crypto/ecdsa"
	"encoding/hex"
	"io/ioutil"
	"math/big"
	"math/rand"
	"os"
	"path/filepath"
	"strconv"
	"time"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/accounts/keystore"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
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
}

var (
	gkeystore_path   = "./conf/block/eth/keystore/"
	gkeystore_passwd = "123456"
	gchainid         = big.NewInt(1234)
	ggasLimit        = uint64(300000)
)

// 初始化ETH
func NewRPCClient(nodeUrl, blockName string) *RPCClient {

	eth_client, err := ethclient.Dial(nodeUrl)
	if err != nil {
		glog.Info("conn nodeurl:%s, err:%s.\n", nodeUrl, err)
		return nil
	}

	rpc_client, _ := rpc.DialContext(context.Background(), nodeUrl)

	rand.Seed(time.Now().UnixNano())

	rpcClient := &RPCClient{}
	rpcClient.ethClient = eth_client
	rpcClient.rpcClient = rpc_client
	rpcClient.blockName = blockName

	rpcClient.DoPost()

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
		glog.Info("DecodeString privatekey:%s, err:%s.\n", privateKeyHex, err)
		return nil, err
	}

	privateKey := new(ecdsa.PrivateKey)
	privateKey.D = new(big.Int).SetBytes(privateKeyBytes)
	privateKey.PublicKey.Curve = crypto.S256() // 替换为您使用的曲线
	privateKey.PublicKey.X, privateKey.PublicKey.Y = privateKey.PublicKey.Curve.ScalarBaseMult(privateKey.D.Bytes())

	return privateKey, nil
}

func loadKeystore(client *RPCClient, blockName string) {
	var files []string
	err := filepath.Walk(gkeystore_path, func(path string, info os.FileInfo, err error) error {
		// glog.Info(":---->path:%s.\n", path)
		if path != gkeystore_path {
			files = append(files, path)
		}
		return nil
	})
	if err != nil {
		glog.Info("load keystore file failed, err:%s.\n", err)
		return
	}
	for _, file := range files {

		//  该路径是二进制执行的相对位置
		keyJson, err := ioutil.ReadFile(file)
		if err != nil {
			glog.Info("read file:%s, failed, err:%s.\n", file, err)
			continue
		}

		key, err := keystore.DecryptKey(keyJson, gkeystore_passwd)
		if err != nil {
			glog.Info("parse json failed, err:%s.\n", err)
			continue
		}

		// glog.Info("keystore key:%v.\n", key)

		publicKey := key.PrivateKey.Public()
		publicKeyECDSA, ok := publicKey.(*ecdsa.PublicKey)
		if !ok {
			glog.Info("error casting public key to ECDSA")
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

		// glog.Info("--->HEX :%s,privateKeyBytes:%s.\n", hex.EncodeToString(key.PrivateKey.D.Bytes()), privateKeyBytes)

		nonce_from_block, err := client.ethClient.PendingNonceAt(context.Background(), fromAddress)
		if err != nil {
			glog.Info("pending failed, err:", err)
			continue
		}
		//  设置nonce初始化值,除去0外，以最大的数值为准
		nonce_from_redis := redis.Redisclient.GetNonceFromRedis(blockName, fromAddress.Hex())
		nonce := nonce_from_redis
		if nonce_from_block != 0 && nonce_from_redis < int64(nonce_from_block) {
			nonce = int64(nonce_from_block)
		}
		redis.Redisclient.SetNonceInitToRedis(blockName, fromAddress.Hex(), int64(nonce))

		gasPrice, err := client.ethClient.SuggestGasPrice(context.Background())
		if err != nil {
			glog.Info("suggest gas price, err:", err)
			continue
		}

		auth, err := bind.NewKeyedTransactorWithChainID(key.PrivateKey, gchainid)
		if err != nil {
			glog.Info("transactor gchainid[%d], err:%v.\n", gchainid, err)
			continue
		}
		auth.Nonce = big.NewInt(int64(nonce))
		auth.Value = big.NewInt(0) // in wei
		auth.GasLimit = ggasLimit  // in units
		auth.GasPrice = big.NewInt(0).Add(gasPrice, big.NewInt(int64(rand.Intn(300000))))

		//  STORE
		storeConAddress, storehex, _, err := store.DeployStore(auth, client.ethClient, "1.0.0") // 部署合约
		glog.Info("----->eth address:%s, storehex:%s, nonce:%d.\n", fromAddress.Hex(), storehex.Hash().String(), auth.Nonce)
		if err != nil {
			glog.Info("deploy DeployStore err:", err)
			continue
		}
		redis.Redisclient.SetContracAddresstToRedis(blockName, "store", fromAddress.Hex(), storeConAddress.Hex())

		//  SMALLBANK
		//  从redis获取
		auth.Nonce = big.NewInt(redis.Redisclient.GetNonceFromRedis(blockName, fromAddress.Hex()))
		// glog.Info("----->eth address:%s, nonce:%d.\n", fromAddress.Hex(), auth.Nonce)
		smallbankConAddress, _, _, err := smallbank.DeploySmallbank(auth, client.ethClient) // 部署合约
		if err != nil {
			glog.Info("deploy DeploySmallbank err:", err)
			continue
		}
		redis.Redisclient.SetContracAddresstToRedis(blockName, "kvstore", fromAddress.Hex(), smallbankConAddress.Hex())

		//  KVSTORE
		//  从redis获取
		auth.Nonce = big.NewInt(redis.Redisclient.GetNonceFromRedis(blockName, fromAddress.Hex()))
		// glog.Info("----->eth address:%s, nonce:%d.\n", fromAddress.Hex(), auth.Nonce)
		kvstoreConAddress, _, _, err := kvstore.DeployKvstore(auth, client.ethClient) // 部署合约
		if err != nil {
			glog.Info("deploy DeployKvstore err:", err)
			continue
		}
		redis.Redisclient.SetContracAddresstToRedis(blockName, "smallbank", fromAddress.Hex(), kvstoreConAddress.Hex())

		// 对fromAddr解锁
		redis.Redisclient.SetNonceunLock(blockName, fromAddress.Hex())
	}

	glog.Info("loading ETH contract Successed.\n")

}

func (client *RPCClient) GetAuthInfo(contractName string) (*bind.TransactOpts, string) {
	contractaddr, fromaddr, _ := redis.Redisclient.GetRandomFieldContract(client.blockName, contractName)
	if contractaddr == "" || fromaddr == "" {
		glog.Info("GetRandomFieldContract not found address.\n")
		return nil, ""
	}
	privateKeyStr := redis.Redisclient.GetPrivateInitToRedis(client.blockName, fromaddr)

	privateKey, _ := stringToPrivateKey(privateKeyStr)

	gasPrice, err := client.ethClient.SuggestGasPrice(context.Background())
	if err != nil {
		glog.Info("suggest gas price, err:%s.\n", err)
		return nil, ""
	}

	auth, err := bind.NewKeyedTransactorWithChainID(privateKey, gchainid)
	if err != nil {
		glog.Info("transactor gchainid[%d], err:%v.\n", gchainid, err)
		return nil, ""
	}

	//  从redis获取
	auth.Nonce = big.NewInt(redis.Redisclient.GetNonceFromRedis(client.blockName, fromaddr))
	auth.Value = big.NewInt(0) // in wei
	auth.GasLimit = ggasLimit  // in units
	auth.GasPrice = big.NewInt(0).Add(gasPrice, big.NewInt(int64(rand.Intn(300000))))

	// glog.Info("--->contractName:%s, contractaddr:%s, fromaddr:%s, auth:%d.\n", contractName, contractaddr, fromaddr, auth.Nonce)
	return auth, contractaddr
}

// 统一对外提供合约调用的接口
func (client *RPCClient) ExecContract(ChaincodeID, ChaincodeFunc string, Params []string) string {
	auth, contractaddr := client.GetAuthInfo(ChaincodeID)
	if auth == nil {
		glog.Info("now get ChaincodeID[%s] error, please check!\n", ChaincodeID)
		return ""
	}
	contractAddr := common.HexToAddress(contractaddr)
	var hex = ""
	//  随机选择选择一个地址的信息
	if ChaincodeID == "store" {
		hex = client.exeStoreFunction(ChaincodeFunc, Params, auth, contractAddr)
	} else if ChaincodeID == "kvstore" {
		hex = client.exeKVStoreFunction(ChaincodeFunc, Params, auth, contractAddr)
	} else if ChaincodeID == "smallbank" {
		hex = client.exeSmallbankFunction(ChaincodeFunc, Params, auth, contractAddr)
	} else {
		glog.Info("not support chaincideid [%s], please check.\n", ChaincodeID)
	}
	return hex
}

func (client *RPCClient) exeStoreFunction(funcname string, Params []string, auth *bind.TransactOpts, contractAddress common.Address) string {
	storenew, _ := store.NewStore(contractAddress, client.ethClient)
	storeRawcont := store.StoreRaw{Contract: storenew}

	var err error
	var storecTx *types.Transaction
	if funcname == "setItem" {
		storecTx, err = storeRawcont.Transact(auth, funcname, Params[0], Params[1])
	} else if funcname == "getItem" {
		storecTx, err = storeRawcont.Transact(auth, funcname, Params[0])
	} else if funcname == "versionContract" {
		storecTx, err = storeRawcont.Transact(auth, funcname)
	} else {
		glog.Info("contract Store not support funcname [%s].\n", funcname)
		return ""
	}
	if err != nil {
		glog.Info("Transact addres[%s-%d], err:%s", contractAddress, auth.Nonce, err)
		return ""
	}
	// glog.Info("txid : %v\n", storecTx.Hash().String())
	return storecTx.Hash().String()
}

func (client *RPCClient) exeKVStoreFunction(funcname string, Params []string, auth *bind.TransactOpts, contractAddress common.Address) string {

	kvstorenew, _ := kvstore.NewKvstore(contractAddress, client.ethClient)
	kvstoreRawcont := kvstore.KvstoreRaw{Contract: kvstorenew}

	var err error
	var storecTx *types.Transaction
	if funcname == "set" {
		storecTx, err = kvstoreRawcont.Transact(auth, funcname, Params[0], Params[1])
	} else if funcname == "get" {
		storecTx, err = kvstoreRawcont.Transact(auth, funcname, Params[0])
	} else {
		glog.Info("contract KVStore not support funcname [%s].\n", funcname)
		return ""
	}
	if err != nil {
		glog.Info("Transact addres[%s-%d], err:%s", contractAddress, auth.Nonce, err)
		return ""
	}
	// glog.Info("txid : %v\n", storecTx.Hash().String())
	return storecTx.Hash().String()
}

func (client *RPCClient) exeSmallbankFunction(funcname string, Params []string, auth *bind.TransactOpts, contractAddress common.Address) string {
	smallbanknew, _ := smallbank.NewSmallbank(contractAddress, client.ethClient)
	smallbankRawcont := smallbank.SmallbankRaw{Contract: smallbanknew}

	var err error
	var storecTx *types.Transaction
	if funcname == "almagate" {
		storecTx, err = smallbankRawcont.Transact(auth, funcname, Params[0], Params[1])
	} else if funcname == "getBalance" {
		storecTx, err = smallbankRawcont.Transact(auth, funcname, Params[0])
	} else if funcname == "updateBalance" {
		arg1, _ := new(big.Int).SetString(Params[1], 10)
		storecTx, err = smallbankRawcont.Transact(auth, funcname, Params[0], arg1)
	} else if funcname == "updateSaving" {
		arg1, _ := new(big.Int).SetString(Params[1], 10)
		storecTx, err = smallbankRawcont.Transact(auth, funcname, Params[0], arg1)
	} else if funcname == "sendPayment" {
		arg2, _ := new(big.Int).SetString(Params[2], 10)
		storecTx, err = smallbankRawcont.Transact(auth, funcname, Params[0], Params[1], arg2)
	} else if funcname == "writeCheck" {
		arg1, _ := new(big.Int).SetString(Params[1], 10)
		storecTx, err = smallbankRawcont.Transact(auth, funcname, Params[0], arg1)
	} else {
		glog.Info("contract Smallbank not support funcname [%s].\n", funcname)
		return ""
	}

	if err != nil {
		glog.Info("Transact addres[%s-%d], err:%s", contractAddress, auth.Nonce, err)
		return ""
	}
	// glog.Info("txid : %v\n", storecTx.Hash().String())
	return storecTx.Hash().String()
}

func (client *RPCClient) DoPost() {
	var networkid string
	client.rpcClient.Call(&networkid, "net_version")
	glog.Info("--->networkid:", networkid)

	var proto_version string
	client.rpcClient.Call(&proto_version, "eth_protocolVersion")
	glog.Info("--->proto_version:", proto_version)

	if networkid == "" && proto_version == "" {
		glog.Exit("client failed, please check.\n")
	}
}

func (client *RPCClient) GetBlance(addr string) {
	var reply interface{}

	err := client.rpcClient.Call(&reply, "eth_getBalance", addr, "latest") //第一个是用来存放回复数据的格式，第二个是请求方法
	if err != nil {
		glog.Info("err:", err)
	}
	//这里得到的还是16进制的需要做个进制转换成10进制
	n := new(big.Int)
	n, _ = n.SetString(reply.(string)[2:], 16)
	basevalue := big.NewInt(1000000000000000000)
	content := addr + " ---> " + (n.Div(n, basevalue)).String() + "ETH" + "\n"
	glog.Info(content)
}

func (client *RPCClient) GetTranCountByTransID(transHex string) (int64, int64) {
	var reply interface{}
	err := client.rpcClient.Call(&reply, "eth_getTransactionByHash", transHex) //第一个是用来存放回复数据的格式，第二个是请求方法
	if err != nil {
		glog.Info("err:", err)
		return -1, -1
	}
	if reply == nil {
		return -1, -1
	}
	res, _ := reply.(map[string]interface{})

	if _, ok := res["blockNumber"]; !ok {
		// 不存在
		glog.Info("res not find key blockNumber")
		return -1, -1
	}
	blockNumber := res["blockNumber"]
	if blockNumber == nil {
		return -1, -1
	}

	// restest, isPending, _ := client.ethClient.TransactionByHash(context.Background(), common.HexToHash(transHex))
	// glog.Info("---->eth_getTransactionByHash, res[%s].\n", res)
	// glog.Info("---->TransactionByHash, hex[%s], result:%v, time:%v.\n", transHex, isPending, restest.Time())

	number, _ := strconv.ParseInt(blockNumber.(string)[2:], 16, 32)
	err = client.rpcClient.Call(&reply, "eth_getBlockByNumber", blockNumber, true) //第一个是用来存放回复数据的格式，第二个是请求方法
	if err != nil {
		glog.Info("err:", err)
		return -1, -1
	}
	blockInfo, _ := reply.(map[string]interface{})
	if blockInfo == nil {
		glog.Info("err:", err)
		return -1, -1
	}

	if _, ok := blockInfo["timestamp"]; !ok {
		// 不存在
		glog.Info("blockInfo not find key timestamp")
		return -1, -1
	}
	if blockInfo["timestamp"] == nil {
		return -1, -1
	}
	blcokTime, _ := strconv.ParseInt(blockInfo["timestamp"].(string)[2:], 16, 64)
	// glog.Info("------>time:%s,%d\n", blockInfo["timestamp"], blcokTime)

	// // 获取区块头
	// block, err := client.ethClient.BlockByHash(context.Background(), common.HexToHash(res["blockHash"].(string)))
	// if err != nil {
	// 	glog.Info("---->address:%s, err:%s", res["blockHash"].(string), err)
	// 	return -1, -1
	// }

	// // 取出时间戳
	// timestamp := block.Header().Time

	// fmt.Printf("---->block number:%d, blockTime:%v.\n", block.Header().Number, int64(timestamp))

	return blcokTime * 1000, number
}

func (client *RPCClient) GetTranCountByBlockNumber(Number int64) (int64, int64) {
	var reply interface{}
	blockNumber := strconv.FormatInt(Number, 16)
	blockNumber = "0x" + blockNumber
	err := client.rpcClient.Call(&reply, "eth_getBlockByNumber", blockNumber, true) //第一个是用来存放回复数据的格式，第二个是请求方法
	if err != nil {
		glog.Info("err:", err)
		return -1, -1
	}
	blockInfo, _ := reply.(map[string]interface{})
	// glog.Info("------>time:", blockInfo["timestamp"])
	// UTC时间转为UTC+8
	blockTime, _ := strconv.ParseInt(blockInfo["timestamp"].(string)[2:], 16, 64)

	err = client.rpcClient.Call(&reply, "eth_getBlockTransactionCountByNumber", blockNumber) //第一个是用来存放回复数据的格式，第二个是请求方法
	if err != nil {
		glog.Info("err:", err)
		return -1, -1
	}

	trancount, _ := strconv.ParseInt(reply.(string)[2:], 16, 32)
	return blockTime, trancount
}

func (client *RPCClient) FindTranCountByTransID(transHex string) bool {
	var reply interface{}
	err := client.rpcClient.Call(&reply, "eth_getTransactionByHash", transHex) //第一个是用来存放回复数据的格式，第二个是请求方法
	if err != nil {
		glog.Info("err:", err)
		return false
	}

	blockInfo, _ := reply.(map[string]interface{})

	if blockInfo["blockNumber"] != nil {
		return true
	}
	return false
}

// func (cinstance *ContractInstance) GetContractItem(key string) string {
// 	storeCallNew, err := store.NewStoreCaller(cinstance.address, cinstance.client)
// 	bind_call := &bind.CallOpts{Pending: true}
// 	value, err := storeCallNew.GetItem(bind_call, key)
// 	if err != nil {
// 		glog.Info(err)
// 	}
// 	// glog.Info("value:", value)
// 	return value
// }

// func (cinstance *ContractInstance) GetContractVersion() string {
// 	storeCallNew, err := store.NewStoreCaller(cinstance.address, cinstance.client)
// 	bind_call := &bind.CallOpts{Pending: true}
// 	version, err := storeCallNew.VersionContract(bind_call)
// 	if err != nil {
// 		glog.Info(err)
// 	}
// 	// glog.Info("version:", version)
// 	return version
// }
