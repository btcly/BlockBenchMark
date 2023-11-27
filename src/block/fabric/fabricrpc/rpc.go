package fabricrpc

import (
	// context1 "context"
	// "log"

	_ "io/ioutil"

	// "time"
	// "os"blcokbenchmark/src/block/fabric

	proto "github.com/golang/protobuf/proto"
	"github.com/hyperledger/fabric-protos-go/common"
	"github.com/hyperledger/fabric-protos-go/peer"

	// "blcokbenchmark/src/block/github.com/fabric-sdk-go/pkg/client/channel"
	// "blcokbenchmark/src/block/github.com/fabric-sdk-go/pkg/client/ledger"
	// "blcokbenchmark/src/block/github.com/fabric-sdk-go/pkg/common/errors/retry"
	// "blcokbenchmark/src/block/github.com/fabric-sdk-go/pkg/common/providers/fab"
	// "blcokbenchmark/src/block/github.com/fabric-sdk-go/pkg/core/config"
	// "blcokbenchmark/src/block/github.com/fabric-sdk-go/pkg/fabsdk"

	"github.com/hyperledger/fabric-sdk-go/pkg/client/channel"
	"github.com/hyperledger/fabric-sdk-go/pkg/client/ledger"
	"github.com/hyperledger/fabric-sdk-go/pkg/common/errors/retry"
	"github.com/hyperledger/fabric-sdk-go/pkg/common/providers/fab"
	"github.com/hyperledger/fabric-sdk-go/pkg/core/config"
	"github.com/hyperledger/fabric-sdk-go/pkg/fabsdk"
	"github.com/golang/glog"
)

type FabricClient struct {
	channelclient *channel.Client
	ledgerclient  *ledger.Client
}

// init the sdk
func InitSDK() *fabsdk.FabricSDK {
	//// Initialize the SDK with the configuration file
	configProvider := config.FromFile("./conf/block/fabric/fabric_config.yaml")
	sdk, err := fabsdk.New(configProvider)
	if err != nil {
		glog.Info("failed to create sdk: %s", err)
	}
	return sdk
}

func InitFabricClient(sdk *fabsdk.FabricSDK) *FabricClient {
	ccp := sdk.ChannelContext("mychannel", fabsdk.WithUser("User1"), fabsdk.WithOrg("Org1"))
	sacc_client := &FabricClient{}
	sacc_client.channelclient, _ = channel.New(ccp)
	sacc_client.ledgerclient, _ = ledger.New(ccp)
	return sacc_client
}

// 统一对外提供链码调用的接口
func (client *FabricClient) ExecContract(ChaincodeID, ChaincodeFunc string, Params []string) string {
	hex, _ := client.exeChaincode(ChaincodeID, ChaincodeFunc, Params)
	return hex
}

// Fabric通用调用函数
func (client *FabricClient) exeChaincode(chainCode string, chainCodeFunc string, args []string) (string, error) {
	createCarArgs := make([][]byte, 0)
	for _, value := range args {
		createCarArgs = append(createCarArgs, []byte(value))
	}
	// createCarArgs := [][]byte{[]byte("key"), []byte("value")}
	// createCarArgstest := [][]byte{[]byte("key"), []byte("value")}
	// glog.Info("#########now:%s.\n", createCarArgs)
	resp1, err := client.channelclient.Execute(channel.Request{ChaincodeID: chainCode, Fcn: chainCodeFunc, Args: createCarArgs}, channel.WithRetry(retry.DefaultChannelOpts))
	// createCarArgs := [][]byte{[]byte("CAR2")}
	// resp1, err := channalClient.Execute(channel.Request{ChaincodeID: "fabcar", Fcn: "QueryCar", Args: createCarArgs}, channel.WithRetry(retry.DefaultChannelOpts))
	if err != nil {
		// glog.Info("ExeChaincode:", err)
		return "", err
	}
	transHex := string(resp1.TransactionID)

	// timestamp := time.Unix(pbtran.GetHeader().GetChannelHeader().Timestamp.Seconds, 0).Format("2006-01-02 15:04:05")
	// timestamp := pbtran.GetHeader().String()
	// payload_by := resp1.Responses

	// payload := common.Payload{}
	// proto.Unmarshal(payload_by, &payload)
	// channelHeader := common.ChannelHeader{}
	// proto.Unmarshal(payload.GetHeader().GetChannelHeader(), &channelHeader)

	// timestamp := int64(channelHeader.Timestamp.Seconds)*1000 + int64(channelHeader.Timestamp.Nanos)/1000

	// glog.Info(string(resp1.Payload))
	return transHex, err
}

func (client *FabricClient) FindTranCountByTransID(transHex string) int32 {
	// lclient, _ := ledger.New(ctx)
	pbtran, err := client.ledgerclient.QueryTransaction(fab.TransactionID(transHex))
	if err != nil {
		glog.Info("transid:%s failed, err:%s", transHex, err)
		return 0
	}

	// glog.Info(">>>###number:%d", pbtran.GetValidationCode())
	return pbtran.GetValidationCode()
}

func (client *FabricClient) GetTranCountByTransID(transHex string) (int64 /* 毫秒时间戳 */, int64) {
	// lclient, _ := ledger.New(ctx)
	if len(transHex) == 0 {
		return -1, -1
	}
	pbtran, err := client.ledgerclient.QueryTransaction(fab.TransactionID(transHex))
	if err != nil {
		glog.Info("QueryTransaction transid:%s failed, err:%s", transHex, err)
		return -1, -1
	}

	if peer.TxValidationCode(pbtran.ValidationCode) != peer.TxValidationCode_VALID {
		return -1, -1
	}

	payload_by := pbtran.GetTransactionEnvelope().GetPayload()

	payload := common.Payload{}
	proto.Unmarshal(payload_by, &payload)
	channelHeader := common.ChannelHeader{}
	proto.Unmarshal(payload.GetHeader().GetChannelHeader(), &channelHeader)

	// 交易上块时间单位毫秒
	timestamp := int64(channelHeader.Timestamp.Seconds)*1000 + int64(channelHeader.Timestamp.Nanos)/1000000
	// glog.Info("---test,hex:%s, time:%s, calc time:%d.", transHex, channelHeader.Timestamp, timestamp)

	// block
	pbblock, err := client.ledgerclient.QueryBlockByTxID(fab.TransactionID(transHex))
	if err != nil {
		glog.Info("QueryBlockByTxID transid:%s failed, err:%s", transHex, err)
		return -1, -1
	}
	number := int64(pbblock.GetHeader().GetNumber())
	return timestamp, number
}

func (client *FabricClient) GetTranCountByBlockNumber(blockNumber int64) (int64, int64) {
	// lclient, _ := ledger.New(ctx)
	pbblock, err := client.ledgerclient.QueryBlock(uint64(blockNumber))
	if err != nil {
		glog.Info("blocknumber:%d failed, err:%s", blockNumber, err)
		return -1, -1
	}

	envelope := common.Envelope{}
	proto.Unmarshal(pbblock.GetData().GetData()[0], &envelope)

	payload := common.Payload{}
	proto.Unmarshal(envelope.GetPayload(), &payload)
	channelHeader := common.ChannelHeader{}
	proto.Unmarshal(payload.GetHeader().GetChannelHeader(), &channelHeader)

	transaction := peer.Transaction{}
	proto.Unmarshal(payload.GetData(), &transaction)
	timestamp := int64(channelHeader.Timestamp.Seconds)*1000 + int64(channelHeader.Timestamp.Nanos)/1000000

	// glog.Info(">>>###timestamp:%d\n", timestamp)
	return timestamp, int64(len(transaction.GetActions()))
}
