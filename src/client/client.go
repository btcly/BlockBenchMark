package main

import (
	"errors"
	"log"
	"math/rand"
	"reflect"
	"strconv"
	"strings"
	"time"
	"flag"

	// zmq "github.com/pebbe/zmq4"
	pb "blcokbenchmark/src/protocs"
	"blcokbenchmark/src/utils"

	"github.com/spf13/viper"
	"golang.org/x/net/context"
	"google.golang.org/grpc"

	// "google.golang.org/grpc/credentials/insecure"
	"google.golang.org/grpc/resolver"
	"google.golang.org/grpc/resolver/manual"
	"github.com/golang/glog"
)

type FuncAndParamslen struct {
	ChaincodeFunc string
	ParamsLen     int64
}

type ClientEndPoint struct {
	client_type    string
	server_ip      []string
	blockname      string
	chaincode      string
	functionParams []FuncAndParamslen
	qps            float64
	open           bool
	uuid           string
}

func toStringSlice(actual interface{}) ([]string, error) {
	var res []string
	value := reflect.ValueOf(actual)
	if value.Kind() != reflect.Slice && value.Kind() != reflect.Array {
		return nil, errors.New("parse error")
	}
	for i := 0; i < value.Len(); i++ {
		res = append(res, value.Index(i).Interface().(string))
	}
	return res, nil
}

func main() {
	//初始化命令行参数
    flag.Parse()
    //退出时调用，确保日志写入文件中
    defer glog.Flush()
	log.SetFlags(log.LstdFlags | log.Lshortfile)
	
	config := viper.New()
	config.AddConfigPath("./")
	config.SetConfigName("config_client")
	config.SetConfigType("json")
	rand.Seed(time.Now().UnixNano())

	if err := config.ReadInConfig(); err != nil {
		glog.Exit(">>>>error, ", err)
	}

	// redis_url := config.GetString("redis.url")
	// redis_passwd := config.GetString("redis.passwd")
	// redis_client := redis.InitRedis(redis_url, redis_passwd)

	uuid := utils.GetLocalUUID()

	server_ip, _ := toStringSlice(config.Get("sever_ip"))
	glog.Info("server_ip>>>>>", config.Get("sever_ip"))

	nodes := config.Get("items")
	for _, value := range nodes.([]interface{}) {
		node := value.(map[string]interface{})
		glog.Info("############[%s]\n############[%s]\n############[%s]\n############[%s]\n############[%s]\n\n", node["type"], node["sever_ip"], node["blockname"], node["qps"], node["open"])

		client_endpoint := &ClientEndPoint{}
		client_endpoint.client_type = node["type"].(string)
		client_endpoint.server_ip = server_ip
		client_endpoint.blockname = node["blockname"].(string)
		client_endpoint.chaincode = node["chaincode"].(string)

		functionParams, _ := toStringSlice(node["functionParams"])
		client_endpoint.functionParams = []FuncAndParamslen{}
		for _, item := range functionParams {
			fplen := FuncAndParamslen{}
			idx := strings.Index(item, ":") //  用于分割函数明和参数个数

			if idx == -1 {
				fplen.ChaincodeFunc = item
				fplen.ParamsLen = 0
			} else {
				fplen.ChaincodeFunc = item[:idx]
				fplen.ParamsLen, _ = strconv.ParseInt(item[idx+1:], 10, 64)
			}

			client_endpoint.functionParams = append(client_endpoint.functionParams, fplen)
		}
		// client_endpoint.functionParams, _ = toStringSlice(node["functionParams"])
		client_endpoint.qps = node["qps"].(float64)
		client_endpoint.open = node["open"].(bool)
		client_endpoint.uuid = uuid

		// //  统计每秒的qps和延迟
		// ticker := time.NewTicker(time.Second)
		// defer ticker.Stop()
		// go func() {
		// 	for range ticker.C {
		// 		redis_client.ExeQPS(client_endpoint.blockname, client_endpoint.chaincode)
		// 	}
		// }()

		go perf(client_endpoint)

	}

	for {
		time.Sleep(10 * time.Second)
		//此处等待，接受数据开始工作
	}
}
func perf(client_endpoint *ClientEndPoint) {
	if !client_endpoint.open {
		// glog.Info("############ [%s %s] is close!", client_endpoint.blockname, client_endpoint.chaincode)
		return
	} else {
		glog.Info("############ [%s %s] is open!", client_endpoint.blockname, client_endpoint.chaincode)

	}

	addrs := make([]resolver.Address, 0)
	for _, value := range client_endpoint.server_ip {

		addrs = append(addrs, resolver.Address{Addr: value})
	}
	glog.Info(">>>>>", addrs)

	r := manual.NewBuilderWithScheme("whatever")
	conn, err := grpc.Dial(
		r.Scheme()+":///test.server",
		grpc.WithInsecure(),
		grpc.WithResolvers(r),
		grpc.WithDefaultServiceConfig(`{"loadBalancingPolicy":"round_robin"}`))
	r.UpdateState(resolver.State{Addresses: addrs})
	// {Addr: "127.0.0.1:50051"}, {Addr: "127.0.0.1:7001"}}})

	if err != nil {
		glog.Exit("did not connect: %v", err)
	}
	defer conn.Close()

	client := pb.NewWorkLoadClient(conn)

	// conn, err := grpc.Dial(client_endpoint.server_ip, grpc.WithTransportCredentials(insecure.NewCredentials()))
	// if err != nil {
	// 	// handle error
	// 	panic(err)
	// }
	// defer conn.Close()

	// client := pb.NewWorkLoadClient(conn)

	glog.Info("############ UUID[%s]-QPS[%f] >>>>>>>>>>>!", client_endpoint.uuid, client_endpoint.qps)
	ticker := time.NewTicker(time.Microsecond * time.Duration(1000000/client_endpoint.qps))
	num := 0
	for {
		// ticker.Reset(time.Millisecond * time.Duration(1000/client_endpoint.qps)) // 这里复用了 timer
		select {
		case <-ticker.C:
			go func() {
				// 随机压测函数
				// create_time := time.Now().UnixNano() / 1000000 // 毫秒
				func_index := rand.Intn(len(client_endpoint.functionParams))
				params := []string{}
				for idx := int64(0); idx < client_endpoint.functionParams[func_index].ParamsLen; idx++ {
					params = append(params, strconv.Itoa(rand.Intn(1000)))
				}
				// glog.Info("function:%s, params:%v.\n", client_endpoint.functionParams[func_index].ChaincodeFunc, params)

				req := pb.WorkLoadRequest{
					BlockchainName: client_endpoint.blockname,
					ChaincodeID:    client_endpoint.chaincode,
					ChaincodeFunc:  client_endpoint.functionParams[func_index].ChaincodeFunc,
					ClientUUID:     client_endpoint.uuid,
					ParamsLen:      client_endpoint.functionParams[func_index].ParamsLen,
					Params:         params,
					// Params: []string{string(test_index), string(test_index), string(test_index)},
				}
				// glog.Info("############ [%s]-[%d]-[%s] >>>>>>>>>>>!", client_endpoint.blockname, create_time, req.Params)
				_, err := client.SendWorkLoad(context.Background(), &req)
				if err != nil {
					// glog.Info("client.SendWorkLoad error:", err)
					return
				}
				// glog.Info("get msg from server:[%v] \n", reply)
			}()
			num += 1
		}
	}
}
