package redis

import (
	mysql "blcokbenchmark/src/mysql"
	"context"
	// "log"
	"math/rand"
	"strconv"
	"time"

	//  "reflect"

	// "github.com/redis/go-redis/v9"
	"github.com/go-redis/redis/v8"
	"github.com/golang/glog"
)

type RedisInstanceInfo struct {
	RedisClient  *redis.Client
	Ctx          context.Context
	ServerUUID   string
	BlockNameMap map[string]bool //  存放blockname和是否启动的标记，true：启动，false：关闭
}

var Redisclient *RedisInstanceInfo

// 统计思想：
// 1、使用,服务器的uuid作为域值，blockname和hex作为hash的key，存放redis数据
// 2、数据项:timestart，timeend，chaincode，vaild(0：无效，-1：失败，1：成功)
// 3、在数据存储结束（成功或者失败）后，每隔一段时间数据同步到数据库对外展示

//  redis hash  key说明
// blockname+hex: 存放当前hex的基本信息，valid是用来表示上链成功信息记录（-1，失败；1，成功），默认值为-1
// blockname+serveruuid: 表示当前服务上，当前区块链上执行的hex列表
// blockname+Contractname: 存放当前区块链上的合约的安装地址信息
// blockname+fromeAddress: 存放当前区块链上地址信息(nonce, 私钥等)

// 采用可变参数，组成前缀key
func createRedisPrex(vals ...string) string {
	key_name := ""
	for index, val := range vals {
		if index >= 1 {
			key_name += (":" + val)
		} else {
			key_name += val
		}
	}
	return key_name
}

// vaild(0：无效，-1：失败，1：成功)
// startTime  endTime: milliseconds
func (redisIns *RedisInstanceInfo) SetBlockQPS(blockName, chainCodeID, trans_hex, client_uuid string, startTime, endTime time.Time) {
	redis_key := createRedisPrex(blockName, trans_hex)
	uuid_redis_key := createRedisPrex(blockName, redisIns.ServerUUID)
	//  标记当前的hex是否上链，默认为1。0为成功  [1,500)正在查询  [500,...]失败
	redisIns.RedisClient.HSet(redisIns.Ctx, uuid_redis_key, redis_key, 1)
	// 此数据暂时存放到redis上，key值和数据库字段名保持一致
	redisIns.RedisClient.HMSet(redisIns.Ctx, redis_key, "trans_hex", trans_hex, "server_uuid", redisIns.ServerUUID, "client_uuid", client_uuid, "block_name", blockName,
		"chaincodeID", chainCodeID, "start_time", startTime.UnixMilli(), "end_time", endTime.UnixMilli(), "valid", -1)
	// glog.Info("--->redis blockname:%s, hex:%s.\n", blockName, trans_hex)
}

// 查找当前服务端上，运行指定区块链的tranhex
//
//	查找当前hex是否已经成功上链，value为非0表示失败或者未上链的，需要遍历
func (redisIns *RedisInstanceInfo) GetBlockTransInfo(block_name string) []string {
	not_in_block := []string{}
	uuid_redis_key := createRedisPrex(block_name, redisIns.ServerUUID)
	result, _ := redisIns.RedisClient.HGetAll(redisIns.Ctx, uuid_redis_key).Result()
	for redis_key, value := range result {
		// glog.Info("---->redis:%s, value:%s, %s.\n", redis_key, value, redis_key[:len(block_name)])
		if value != "0" && redis_key[:len(block_name)] == block_name {
			not_in_block = append(not_in_block, redis_key[len(block_name)+1:])
		}
	}

	return not_in_block
}

// ServerUUID 中key的value当前为0表示已经在区块中，超过500表示已经超时或者失败

// ServerUUID 指定key的value增1
func (redisIns *RedisInstanceInfo) IncrBlockUUID(blockName, trans_hex string) {
	redis_key := createRedisPrex(blockName, trans_hex)
	uuid_redis_key := createRedisPrex(blockName, redisIns.ServerUUID)
	redisIns.RedisClient.HIncrBy(redisIns.Ctx, uuid_redis_key, redis_key, 1)
}

// ServerUUID 指定key已经上链的处理
// trans_count：交易的个数。对于私立(ETH测试网)一般设置为1，对于公链是需要设置为区块的交易数
func (redisIns *RedisInstanceInfo) SetBlockUUIDSuccess(blockName, trans_hex string, block_height int64, end_time time.Time) {
	redis_key := createRedisPrex(blockName, trans_hex)
	uuid_redis_key := createRedisPrex(blockName, redisIns.ServerUUID)
	redisIns.RedisClient.HSet(redisIns.Ctx, uuid_redis_key, redis_key, 0)
	redisIns.RedisClient.HMSet(redisIns.Ctx, redis_key, "end_time", end_time.UnixMilli(), "valid", 1, "block_height", block_height, "trans_count", 1)
}

// redis数据同步到mysql
func (redisIns *RedisInstanceInfo) SyncRedisToSQL() {
	data := []mysql.TableInfo{}
	//  遍历当前服务上允许接受的区块链
	for key, open := range redisIns.BlockNameMap {
		if !open {
			//  该block未再此server上启动
			continue
		}
		uuid_redis_key := createRedisPrex(key, redisIns.ServerUUID)
		//  遍历当前区块下的hex
		result, _ := redisIns.RedisClient.HGetAll(redisIns.Ctx, uuid_redis_key).Result()
		// glog.Info("-->read redis key:%s, count:%d.\n", uuid_redis_key, len(result))
		for redis_key, value := range result {
			value_cnt, _ := strconv.ParseInt(value, 10, 64)
			//  该值默认为1，[1,500)这个不认为是失败，需要待检查
			//  TODO暂时用超时代替，后续修改为检测结果
			if value_cnt >= 1 && value_cnt < 100 {
				// glog.Info("---->redis:%s, value:%s.\n", redis_key, value)
				continue
			}
			tran_result, _ := redisIns.RedisClient.HGetAll(redisIns.Ctx, redis_key).Result()
			// glog.Info("---->redis key:%s, result:%s.\n", redis_key, tran_result)

			// 从redis中读取暂时存放的数据
			trans_hex := tran_result["trans_hex"]
			server_uuid := tran_result["server_uuid"]
			client_uuid := tran_result["client_uuid"]
			block_name := tran_result["block_name"]
			chaincodeID := tran_result["chaincodeID"]
			start_time, _ := strconv.ParseInt(tran_result["start_time"], 10, 64) // ms
			end_time, _ := strconv.ParseInt(tran_result["end_time"], 10, 64)     // ms
			valid, _ := strconv.ParseInt(tran_result["valid"], 10, 64)
			trans_count, _ := strconv.ParseInt(tran_result["trans_count"], 10, 64)
			block_height, _ := strconv.ParseInt(tran_result["block_height"], 10, 64)
			// write sql
			if len(tran_result) >= 1 && valid != 0 {
				redisIns.RedisClient.Unlink(redisIns.Ctx, redis_key)
				redisIns.RedisClient.HDel(redisIns.Ctx, redisIns.ServerUUID, redis_key)
				// glog.Info("--->redis HDEL uuid:%s, hex:%s, value:%s-%d, tran_result:%s.\n", redisIns.ServerUUID, redis_key, value, value_cnt, tran_result)

				data = append(data, mysql.TableInfo{
					Tran_hex:     trans_hex,
					Server_uuid:  server_uuid,
					Client_uuid:  client_uuid,
					Block_name:   block_name,
					ChaincodeID:  chaincodeID,
					Create_time:  start_time / 1000,
					Start_time:   start_time,
					End_time:     end_time,
					Valid:        valid,
					Trans_count:  trans_count,
					Block_height: block_height,
				})
			}
			if len(data) >= 10 {
				glog.Info("sync count[%d] data to mysql.", len(data))
				mysql.Dbconn.InsertBatchBlockInfos(data)
				data = data[:0]
			}
		}
	}
	if len(data) >= 1 {
		glog.Info("sync count[%d] data to mysql.", len(data))
		mysql.Dbconn.InsertBatchBlockInfos(data)
		data = data[:0]
	}

}

// 统计QPS
func (redisIns *RedisInstanceInfo) ExeQPS() {
	redisIns.SyncRedisToSQL()
}

// 启动设置私钥
func (redisIns *RedisInstanceInfo) SetPrivateInitToRedis(blockName, address, privateKey string) {
	redis_key := createRedisPrex(blockName, address)
	redisIns.RedisClient.HSet(redisIns.Ctx, redis_key, "privateKey", privateKey)
}

// 获取私钥
func (redisIns *RedisInstanceInfo) GetPrivateInitToRedis(blockName, address string) string {
	redis_key := createRedisPrex(blockName, address)
	privateKey, _ := redisIns.RedisClient.HGet(redisIns.Ctx, redis_key, "privateKey").Result()
	return privateKey
}

// 启动设置eth的nonce初始值，该值是作为旧值
func (redisIns *RedisInstanceInfo) SetNonceInitToRedis(blockName, address string, nonce int64) {
	redis_key := createRedisPrex(blockName, address)
	redisIns.RedisClient.HSet(redisIns.Ctx, redis_key, "nonce", nonce)
}

// 获取eth的nonce，并加1保证递增顺序
// 为了保证顺序性，使用递增函数HIncrBy后数值变为最新值Nonce
func (redisIns *RedisInstanceInfo) GetNonceFromRedis(blockName, fromaddr string) int64 {

	redis_lock_key := createRedisPrex(blockName, fromaddr, "locknonce")
	redis_key := createRedisPrex(blockName, fromaddr)
	for {
		trylock := redisIns.RedisClient.SetNX(redisIns.Ctx, redis_lock_key, 1, 5000*time.Millisecond)
		if trylock.Val() {
			nonce_now := redisIns.RedisClient.HIncrBy(redisIns.Ctx, redis_key, "nonce", 1)
			redisIns.RedisClient.Del(redisIns.Ctx, redis_lock_key)
			// glog.Info("--->get address:%s, nonce:%s.\n", redis_key, nonce_now.Err())
			if nonce_now.Err() == nil {
				nonce := nonce_now.Val()
				return nonce
			}
		}
		time.Sleep(1 * time.Millisecond)
	}

}

// 设置区块链上地址对应安装的合约
// key=blockname:Contractname
// field=ServerUUID:fromeaddress
// value=contractaddr
func (redisIns *RedisInstanceInfo) SetContracAddresstToRedis(blockName, contractName, fromeAddress, contractAddress string) {
	redis_key := createRedisPrex(blockName, contractName)

	redisIns.RedisClient.HSet(redisIns.Ctx, redis_key, fromeAddress, contractAddress)
}

// // 获取合约地址---TODO将于其它分支（feature/add_contract_redis）开发
// // 1、获取全部数据到本地
// // 2、使用setnx进行加锁，将当前区块名称+合约名称+地址作为key
// // 3、加锁失败就跳过，加锁成功就进行下面部分
// // 4、判断合约是否存在，存在则保留，不存在就重新生成
// // 5、上述操作完成后解锁并删除key
// // 6、如果需要创建新的合约就放到检测内存中，部署完成就放入就绪hash中(此在ethrpc中操作)
// func (redisIns *RedisInstanceInfo) CheckETHsContracAddresstToRedis(ethClient *ethclient.Client, blockName, contractName string) {
// 	redis_key := createRedisPrex(blockName, contractName)
// 	result, _ := redisIns.RedisClient.HGetAll(redisIns.Ctx, redis_key).Result()
// 	for fromaddr, contractaddr := range result {
// 		addr_rediskey := createRedisPrex(redis_key, fromaddr, "checklock")
// 		// 设置key的超时时间为5秒
// 		trylock := redisIns.RedisClient.SetNX(redisIns.Ctx, addr_rediskey, 1, 5000*time.Millisecond)
// 		if trylock.Val() {
// 			// 检查合约是否存在
// 			// 要查询的合约地址
// 			contractAddress := common.HexToAddress(contractaddr)
// 			// 查询合约代码
// 			code, _ := ethClient.CodeAt(context.Background(), contractAddress, nil)
// 			glog.Info("--->redis_key:%s, fromaddr:%s, contractaddr:%s, code:%s.\n", redis_key, fromaddr, contractaddr, code)
// 			if len(code) <= 0 {
// 				redisIns.RedisClient.HDel(redisIns.Ctx, redis_key, fromaddr)
// 			}

// 			glog.Info("---->check address fromaddr:%s, contract:%s.\n", fromaddr, contractaddr)
// 			redisIns.RedisClient.Del(redisIns.Ctx, addr_rediskey)
// 		}
// 	}
// }

// 设置nonce的时候加锁
func (redisIns *RedisInstanceInfo) SetNonceLock(blockName, fromeAddress string) bool {
	redis_key := createRedisPrex(blockName, fromeAddress, "lock")
	trylock := redisIns.RedisClient.SetNX(redisIns.Ctx, redis_key, 1, 5000*time.Millisecond)
	return trylock.Val()

}

// 设置nonce的时候解锁
func (redisIns *RedisInstanceInfo) SetNonceunLock(blockName, fromeAddress string) {
	redis_key := createRedisPrex(blockName, fromeAddress, "lock")
	redisIns.RedisClient.Del(redisIns.Ctx, redis_key)
}

func (redisIns *RedisInstanceInfo) GetRandomFieldContract(blockName, contractName string) (string, string, error) {
	redis_key := createRedisPrex(blockName, contractName)

	// 获取所有字段名
	fields, err := redisIns.RedisClient.HKeys(redisIns.Ctx, redis_key).Result()
	// glog.Info("--->redis_key:%s, fields:%s.\n", redis_key, fields)

	if err != nil || len(fields) <= 0 {
		glog.Info("GetRandomFieldContract getkeys err, redis_key:%s.\n", redis_key)
		return "", "", err
	}

	// 随机选择一个字段
	rand.Seed(time.Now().UnixNano())
	randomIndex := rand.Intn(len(fields))
	randomfromaddr := fields[randomIndex]

	// 获取随机字段的值
	randomContractValue, err := redisIns.RedisClient.HGet(redisIns.Ctx, redis_key, randomfromaddr).Result()
	if err != nil || randomContractValue == "" {
		glog.Info("GetRandomFieldContract getkeys err, redis_key:%s, filed:%s.\n", redis_key, randomfromaddr)
		return "", "", err
	}

	return randomContractValue, randomfromaddr, nil
}

// INIT
func InitRedis(url, passwd, serverUUID string) *RedisInstanceInfo {
	redisClient := redis.NewClient(&redis.Options{
		Addr:     url,
		Password: passwd, // no password set
		DB:       0,      // use default DB
	})

	redis_ins := &RedisInstanceInfo{}
	redis_ins.RedisClient = redisClient
	redis_ins.Ctx = context.Background()
	redis_ins.ServerUUID = serverUUID
	redis_ins.BlockNameMap = make(map[string]bool)
	return redis_ins
}
