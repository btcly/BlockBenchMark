package main

import (
	"blcokbenchmark/src/block/eth/ethrpc"
	"blcokbenchmark/src/block/fabric/fabricrpc"
	"blcokbenchmark/src/block/meepo/meeporpc"
	"blcokbenchmark/src/mysql"

	// "blcokbenchmark/src/mysql"
	"blcokbenchmark/src/redis"
	"blcokbenchmark/src/utils"
	"net"
	"time"
	"flag"

	// ctx "context"
	"github.com/spf13/viper"

	pb "blcokbenchmark/src/protocs"

	"golang.org/x/net/context"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
	"github.com/golang/glog"
)

// 全部变量
// 服务端uuid，唯一标识
var serverUUID string

type ETHEndpoint struct {
	rpcClient *ethrpc.RPCClient
	blockname string
}

type MeepoEndpoint struct {
	rpcClient *meeporpc.RPCClient
	blockname string
}

type FabricEndpoint struct {
	fabric_client *fabricrpc.FabricClient
	blockname     string
}

// GRPC
type BlockWorkLoadService struct {
	pb.UnimplementedWorkLoadServer
	Fabric_endpoint *FabricEndpoint
	Eth_endpoint    *ETHEndpoint
	Meepo_endpoint  *MeepoEndpoint
}

func (s BlockWorkLoadService) SendWorkLoad(c context.Context, req *pb.WorkLoadRequest) (*pb.WorkLoadResponse, error) {
	ret := new(pb.WorkLoadResponse)
	ret.Message = "hello, " + req.BlockchainName

	// glog.Info("########name:%s, id:%s, func:%s", req.BlockchainName, req.ChaincodeID, req.ChaincodeFunc)
	// go s.endpoint.dbconn.InsertBlockNameInfosRecord(req.BlockchainName, req.ChaincodeID)
	if s.Fabric_endpoint != nil && req.BlockchainName == s.Fabric_endpoint.blockname {
		s.Fabric_endpoint.EndpointWorkLoad(req.ChaincodeID, req.ChaincodeFunc, req.ClientUUID, req.Params)
	} else if s.Eth_endpoint != nil && req.BlockchainName == s.Eth_endpoint.blockname {
		s.Eth_endpoint.EndpointWorkLoad(req.ChaincodeID, req.ChaincodeFunc, req.ClientUUID, req.Params)
	} else if s.Meepo_endpoint != nil && req.BlockchainName == s.Meepo_endpoint.blockname {
		s.Meepo_endpoint.EndpointWorkLoad(req.ChaincodeID, req.ChaincodeFunc, req.ClientUUID, req.Params)
	} else {
		glog.Info("now not support blockchain [%s].\n", req.BlockchainName)
	}

	return ret, nil
}

//  每一类区块链分别三个函数，
// 初始化：			EndpointInitXXX
// 压测数据更新：	EndpointUpdate
// 接受客户端数据：	EndpointWorkLoad

// FABRIC
func EndpointInitFabric(nodeurl, blockname string) *FabricEndpoint {
	endpoint := &FabricEndpoint{}
	glog.Info("fabric initsdk")
	//初始化sdk
	sdk := fabricrpc.InitSDK()
	client := fabricrpc.InitFabricClient(sdk)

	endpoint.fabric_client = client
	endpoint.blockname = blockname
	return endpoint
}

func (endpoint *FabricEndpoint) EndpointUpdate() {
	trans_hexs := redis.Redisclient.GetBlockTransInfo(endpoint.blockname)
	for _, hex := range trans_hexs {
		timestamp_hex, block_height := endpoint.fabric_client.GetTranCountByTransID(hex)

		if timestamp_hex != -1 {
			// end_time := time.Now() //
			end_time := time.UnixMilli(timestamp_hex)
			// glog.Info("---->fabric trans hex:%s, timestamp_hex:%d, time:%v.\n", hex, timestamp_hex, end_time)
			redis.Redisclient.SetBlockUUIDSuccess(endpoint.blockname, hex, block_height, end_time)
		} else {
			redis.Redisclient.IncrBlockUUID(endpoint.blockname, hex)
		}
	}
}

func (endpoint *FabricEndpoint) EndpointWorkLoad(ChaincodeID, ChaincodeFunc, ClientUUID string, Params []string) {
	go func() {
		start_time := time.Now() // 秒
		hex := endpoint.fabric_client.ExecContract(ChaincodeID, ChaincodeFunc, Params)
		if len(hex) > 0 {
			// glog.Info("---->hex:%s\n", hex)
			// glog.Info("---->fabric trans hex:%s, time:%v.\n", hex, start_time.UnixMilli())
			redis.Redisclient.SetBlockQPS(endpoint.blockname, ChaincodeID, hex, ClientUUID, start_time, start_time)
		}
	}()
}

// ETH
func EndpointInitETH(nodeurl, blockname string) *ETHEndpoint {
	endpoint := &ETHEndpoint{}

	rpcClient := ethrpc.NewRPCClient(nodeurl, blockname)

	endpoint.blockname = blockname
	endpoint.rpcClient = rpcClient

	return endpoint
}

func (endpoint *ETHEndpoint) EndpointUpdate() {
	trans_hexs := redis.Redisclient.GetBlockTransInfo(endpoint.blockname)
	for _, hex := range trans_hexs {
		//  TODO 该时间返回为秒，系统设计单位为毫秒需要改正
		timestamp_hex, block_height := endpoint.rpcClient.GetTranCountByTransID(hex)
		if timestamp_hex != -1 {
			// glog.Info("--->EndpointUpdate hex:%s, timestampe:%d.\n", hex, timestamp_hex)
			// end_time := time.Now() //
			end_time := time.UnixMilli(timestamp_hex)
			redis.Redisclient.SetBlockUUIDSuccess(endpoint.blockname, hex, int64(block_height), end_time)
		} else {
			redis.Redisclient.IncrBlockUUID(endpoint.blockname, hex)
		}
	}
}

func (endpoint *ETHEndpoint) EndpointWorkLoad(ChaincodeID, ChaincodeFunc, ClientUUID string, Params []string) {
	go func() {
		start_time := time.Now()
		hex := endpoint.rpcClient.ExecContract(ChaincodeID, ChaincodeFunc, Params)
		if len(hex) > 0 {
			// glog.Info("--->EndpointWorkLoad hex:%s, start_time:%v, ChaincodeID:%s, ChaincodeFunc:%s, Params:%v.\n", hex, start_time, ChaincodeID, ChaincodeFunc, Params)
			// glog.Info("---->eth trans hex:%s, time:%v.\n", hex, start_time.UnixMilli())
			redis.Redisclient.SetBlockQPS(endpoint.blockname, ChaincodeID, hex, ClientUUID, start_time, start_time)
		}
	}()
}

// Meepo
func EndpointInitMeepo(nodeurl, blockname string) *MeepoEndpoint {
	endpoint := &MeepoEndpoint{}

	rpcClient := meeporpc.NewRPCClient(nodeurl, blockname)
	if rpcClient == nil {
		glog.Exit("conn failed.")
	}

	endpoint.blockname = blockname
	endpoint.rpcClient = rpcClient

	return endpoint
}

func (endpoint *MeepoEndpoint) EndpointUpdate() {
	trans_hexs := redis.Redisclient.GetBlockTransInfo(endpoint.blockname)
	for _, hex := range trans_hexs {
		//  TODO 该时间返回为秒，系统设计单位为毫秒需要改正
		timestamp_hex, block_height := endpoint.rpcClient.GetTranCountByTransID(hex)
		if timestamp_hex != -1 {
			// glog.Info("--->EndpointUpdate hex:%s, timestampe:%d.\n", hex, timestamp_hex)
			// end_time := time.Now() //
			end_time := time.UnixMilli(timestamp_hex)
			redis.Redisclient.SetBlockUUIDSuccess(endpoint.blockname, hex, int64(block_height), end_time)
		} else {
			redis.Redisclient.IncrBlockUUID(endpoint.blockname, hex)
		}
	}
}

func (endpoint *MeepoEndpoint) EndpointWorkLoad(ChaincodeID, ChaincodeFunc, ClientUUID string, Params []string) {
	go func() {
		start_time := time.Now()
		hex := endpoint.rpcClient.ExecContract(ChaincodeID, ChaincodeFunc, Params)
		if len(hex) > 0 {
			// glog.Info("--->EndpointWorkLoad hex:%s, start_time:%v, ChaincodeID:%s, ChaincodeFunc:%s, Params:%v.\n", hex, start_time, ChaincodeID, ChaincodeFunc, Params)
			// glog.Info("---->eth trans hex:%s, time:%v.\n", hex, start_time.UnixMilli())
			redis.Redisclient.SetBlockQPS(endpoint.blockname, ChaincodeID, hex, ClientUUID, start_time, start_time)
		}
	}()
}

// MAIN
func main() {
	// log.SetFlags(log.LstdFlags | log.Lshortfile)

	//初始化命令行参数
	flag.Parse()
	//退出时调用，确保日志写入文件中
	defer glog.Flush()

	serverUUID = utils.GetLocalUUID()

	config := viper.New()
	config.AddConfigPath("./conf/server")
	config.SetConfigName("config_server")
	config.SetConfigType("json")

	if err := config.ReadInConfig(); err != nil {
		glog.Exit(">>>>error, ", err)
	}

	//   初始化redis
	redis_url := config.GetString("redis.url")
	redis_passwd := config.GetString("redis.passwd")

	redis.Redisclient = redis.InitRedis(redis_url, redis_passwd, serverUUID)

	ticker := time.NewTicker(time.Second)
	defer ticker.Stop()

	// 初始化mysql
	mysql_endpoint := &mysql.MySQLendpoint{}
	mysql_endpoint.UserName = config.GetString("mysql.userName")
	mysql_endpoint.Password = config.GetString("mysql.password")
	mysql_endpoint.IpAddrees = config.GetString("mysql.ipAddrees")
	mysql_endpoint.Port = config.GetInt("mysql.port")
	mysql_endpoint.DbName = config.GetString("mysql.dbName")
	mysql_endpoint.Charset = config.GetString("mysql.charset")
	mysql.Mysql_endpoint = *mysql_endpoint
	mysql.InitMysql()
	///////////////////////////////////////////////////////////////////

	// 提前初始化服务接口
	block_service := &BlockWorkLoadService{}

	// 初始化区块链服务端
	nodes := config.Get("nodes")
	for _, value := range nodes.([]interface{}) {
		node := value.(map[string]interface{})

		glog.Info("############[%s]\n", node["name"])
		name := node["name"].(string)
		nodeurl := node["nodeurl"].(string)
		is_open := node["open"].(bool)

		redis.Redisclient.BlockNameMap[name] = is_open
		if is_open == false {
			continue
		}

		if name == "Fabric" {
			fabric_node := EndpointInitFabric(nodeurl, name)
			glog.Info("##### [%s] Init Successed!\n", name)
			block_service.Fabric_endpoint = fabric_node
			go func() {
				time.Sleep(5 * time.Second)
				for range ticker.C {
					fabric_node.EndpointUpdate()
				}
			}()
		} else if name == "ETHPersonal" {
			eth_node := EndpointInitETH(nodeurl, name)
			block_service.Eth_endpoint = eth_node
			glog.Info("##### [%s] Init Successed!\n", name)
			go func() {
				time.Sleep(5 * time.Second)
				for range ticker.C {
					eth_node.EndpointUpdate()
				}
			}()
		} else if name == "MeepoPersonal" {
			meepo_node := EndpointInitMeepo(nodeurl, name)
			block_service.Meepo_endpoint = meepo_node
			glog.Info("##### [%s] Init Successed!\n", name)
			go func() {
				time.Sleep(5 * time.Second)
				for range ticker.C {
					meepo_node.EndpointUpdate()
				}
			}()
		}
	}
	glog.Info("### UUID[%s]\n", serverUUID)

	//  定时更新redis数据
	go func() {
		time.Sleep(5 * time.Second)
		for range ticker.C {
			redis.Redisclient.ExeQPS()
		}
	}()

	//  初始化服务端socket
	sever_ip := config.GetString("sever_ip")
	lis, err := net.Listen("tcp", sever_ip)
	if err != nil {
		panic(err)
	}
	grpcServer := grpc.NewServer()
	pb.RegisterWorkLoadServer(grpcServer, block_service)
	reflection.Register(grpcServer)
	grpcServer.Serve(lis)

	select {}
}
