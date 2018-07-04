package handler

import (
	"encoding/json"
	"httpproxy4blockchain/jsonrpc"
	"httpproxy4blockchain/logger"
	"strings"
)

//in this part, we try to decouple the whole code by a route-controller structure;
//before this server running, all the controller would be written in the router by function init();
//when the client send a json, this server decode this json and decide which controller to process this message;

//我在Server的内部加入一层Router,通过Router对通过Socket发来的信息，通过我们设定的规则进行解析判断后，调用相关的Controller进行任务的分发处理。
//在这个过程中不仅Controller彼此独立，匹配规则和Controller之间也是相互独立的。

//ResultState is the result strunct.
type ResultState struct {
	State string `json:"state"`
}

//RPCResponseState is the strunct for source-state message response.
type RPCResponseState struct {
	JSONRPC string      `json:"jsonrpc"`
	Result  ResultState `json:"result,omitempty"`
	//Result map[string]interface{} `json:"result,omitempty"`
	//Result *json.RawMessage `json:"result,omitempty"`
	Error *RPCError `json:"error,omitempty"`
	ID    uint      `json:"id"`
}

//B_chenhui
//ResultState is the result strunct.
type ResultValue struct {
	Value string `json:"value"`
}

//RPCResponseState is the strunct for source-state message response.
type RPCResponseState2 struct {
	JSONRPC string      `json:"jsonrpc"`
	Result  ResultValue `json:"result,omitempty"`
	//Result map[string]interface{} `json:"result,omitempty"`
	//Result *json.RawMessage `json:"result,omitempty"`
	Error *RPCError `json:"error,omitempty"`
	ID    uint      `json:"id"`
}

//RPCResponseState is the strunct for source-state message response.
type RPCResponseState3 struct {
	JSONRPC string              `json:"jsonrpc"`
	Result  ReadCountResultType `json:"result,omitempty"`
	//Result map[string]interface{} `json:"result,omitempty"`
	//Result *json.RawMessage `json:"result,omitempty"`
	Error *RPCError `json:"error,omitempty"`
	ID    uint      `json:"id"`
}

//{\"value\":\"hahaha3\",\"readcount\":{\"readcount\":11,\"accesstime\":1529403532}}
type ReadCountInfoType struct {
	ReadCount  uint64 `json:"readCount"`
	AccessTime uint64 `json:"accessTime"`
}
type ReadCountResultType struct {
	Value     string            `json:"value"`
	ReadCount ReadCountInfoType `json:"readCount"`
}

//E_chenhui
type State_Resp_Pic_Msg struct {
	Pic string `json:"pic"`
}

//RPCResponseState is the strunct for source-state message response.
type RPCResponsePic struct {
	JSONRPC string             `json:"jsonrpc"`
	Result  State_Resp_Pic_Msg `json:"result,omitempty"`
	//Result map[string]interface{} `json:"result,omitempty"`
	//Result *json.RawMessage `json:"result,omitempty"`
	Error *RPCError `json:"error,omitempty"`
	ID    uint      `json:"id"`
}

//Timestamp is to describ the timestamp.
type Timestamp struct {
	Nanos   uint32 `json:"nanos"`
	Seconds uint32 `json:"seconds"`
}

//ResultTransaction is the struct of result of source-transactions message.
type ResultTransactions struct {
	Timestamp Timestamp `json:"timestamp"`
	Tx_id     string    `json:"tx_id"`
	Value     string    `json:"value"`
}

/*
{
    "jsonrpc": "2.0",
    "id": 0,
    "result": [
        {
            "timestamp": {
                "nanos": 970000000,
                "seconds": 1522853782
            },
            "tx_id": "f3c691ecde3fd667bb1afee96931aa17082f77e8b0e2eaeb71a97e4b26f594f7",
            "value": "1002"
        }
    ]
}
*/
//RPCResponseState is the strunct for source-state message response.
type RPCResponseTransactions struct {
	JSONRPC string               `json:"jsonrpc"`
	Result  []ResultTransactions `json:"result,omitempty"`
	//Result map[string]interface{} `json:"result,omitempty"`
	//Result *json.RawMessage `json:"result,omitempty"`
	Error *RPCError `json:"error,omitempty"`
	ID    uint      `json:"id"`
}

/*
{
    "jsonrpc": "2.0",
    "id": 0,
    "result": {
        "channel_id": "vvtrip",
        "data": [
            {
                "write": {
                    "is_delete": false,
                    "key": "mytest/1",
                    "value": "value1"
                }
            }
        ],
        "timestamp": "Sun Apr 15 2018 04:55:03 GMT+0000 (UTC)",
        "tx_id": "b11a94dd1142559380d1a715da39b6899ed55511f7e23164a50159e4dad4f936",
        "type": "ENDORSER_TRANSACTION"
    }
}*/
type RPCResponseTransaction struct {
	JSONRPC string            `json:"jsonrpc"`
	Result  ResultTransaction `json:"result,omitempty"`
	//Result map[string]interface{} `json:"result,omitempty"`
	//Result *json.RawMessage `json:"result,omitempty"`
	Error *RPCError `json:"error,omitempty"`
	ID    uint      `json:"id"`
}

type ResultTransaction struct {
	Channel_id string     `json:"channel_id"`
	Data       []DataItem `json:"data"`
	Timestamp  uint       `json:"timestamp"`
	Tx_id      string     `json:"tx_id"`
	Type       string     `json:"type"`
}
type DataItem struct {
	Write WriteData `json:"write"`
}
type WriteData struct {
	Is_delete bool   `json:"is_delete"`
	Key       string `json:"key"`
	Value     string `json:"value"`
}

// RPCError represents a JSON-RPC error object if an RPC error occurred.
//
// Code: holds the error code
//
// Message: holds a short error message
//
// Data: holds additional error data, may be nil
//
// See: http://www.jsonrpc.org/specification#error_object
type RPCError struct {
	Code    int         `json:"code"`
	Message string      `json:"message"`
	Data    interface{} `json:"data,omitempty"`
}

//MethodParams for JSON-RPC 2.0 parameters.
type MethodParams struct {
	Channel string `json:"channel"`
	Key     string `json:"key"`
	Tx_id   string `json:"tx_id"`
}

//Msg defined between app client and goproxy4blockchain
type Msg struct {
	Meta    map[string]interface{} `json:"meta"`
	Content jsonrpc.RPCRequest     `json:"content"`
}

/*{
    \"ID\": \"100\",
    \"jianyanxiangmu\": \"ganguanyaoqiu\",
    \"jiliangdanwei\": \"%\",
    \"biaozhunyaoqiu\": \"fuhebiaozhun\",
    \"pic1\": \"test/my1.txt\",
    \"pic2\": \"test/my2.pic\"
}"*/
type State_Resp_Msg struct {
	ID             string `json:"id"`
	Jianyanxiangmu string `json:"jianyanxiangmu"`
	Jiliangdanwei  string `json:"jiliangdanwei"`
	Biaozhunyaoqiu string `json:"biaozhunyaoqiu"`
	Pic1           string `json:"pic1"`
	Pic2           string `json:"pic2"`
}

/*
type Msg struct {
	Meta    map[string]interface{} `json:"meta"`
	Content interface{}            `json:"content"`
}
*/

//Controller is an interface, you can implement by yourself.
type Controller interface {
	Excute(message Msg) []byte
}

/*
//CheckError is to check whether there is an error, if so print it out.
func CheckError(err error) {
	if err != nil {
		//LogErr(os.Stderr, "Fatal error: %s", err.Error())
		logger.Error("Fatal error: %s", err.Error())
	}
}
*/

//sendJsonrpcRequest is to send request to block chain service.
func sendJsonrpcRequest(method string, key string, tx_id string) (*jsonrpc.RPCResponse, error) {
	var err error
	//rpcClient := jsonrpc.NewClient("https://www.ninechain.net/api/v2.1")
	rpcClient := jsonrpc.NewClient("https://testnet.ninechain.net/api/v2.1")
	if rpcClient == nil {
		logger.Error("sendJsonrpcRequest() pcClient is nil!")
		return nil, err
	}

	var rpcResp *jsonrpc.RPCResponse
	//B_chenhui
	if strings.Contains(method, "asset") {
		rpcResp, err = rpcClient.Call(method, &MethodParams{Channel: "assettestchannel", Key: key})
	} else {
		//E_chenhui
		if method == "source-transaction" {
			rpcResp, err = rpcClient.Call(method, &MethodParams{Channel: "notaryinfotestchannel", Key: key, Tx_id: tx_id})
		} else {
			rpcResp, err = rpcClient.Call(method, &MethodParams{Channel: "notaryinfotestchannel", Key: key})
		}
	} //chenhui

	if err != nil {
		logger.Error("sendJsonrpcRequest() err for rpcClient.Call:", err.Error())
		return nil, err
	}
	return rpcResp, nil
}

//B_chenhui
//sendJsonrpcRequest is to send request to block chain service.
func sendJsonrpcRequest4Asset(method string, message []byte) (*jsonrpc.RPCResponse, error) {
	var rpcRequest jsonrpc.RPCRequest
	err := json.Unmarshal(message, &rpcRequest)
	if err != nil {
		logger.Error("sendJsonrpcRequest4Asset() Unmarshal err,", err)
		return nil, err
	}

	//rpcRequest := entermsg.Content
	logger.Info("parsing the JSONRPC2.0 message from app client...")
	f := rpcRequest.Params

	//rpcClient := jsonrpc.NewClient("https://www.ninechain.net/api/v2.1")
	rpcClient := jsonrpc.NewClient("https://testnet.ninechain.net/api/v2.1")
	if rpcClient == nil {
		logger.Error("sendJsonrpcRequest4Asset() pcClient is nil!")
		return nil, err
	}

	var rpcResp *jsonrpc.RPCResponse
	//B_chenhui
	if strings.Contains(method, "asset") {
		//rpcResp, err = rpcClient.Call(method, &MethodParams{Channel: "assettestchannel", Key: key})
		rpcResp, err = rpcClient.Call(method, &f)
	}

	if err != nil {
		logger.Error("sendJsonrpcRequest4Asset() err for rpcClient.Call:", err.Error())
		return nil, err
	}
	return rpcResp, nil
}

//E_chenhui

//verifyStateMsg is to parse and verify the format of source-state message.
func verifyStateMsg(rpcResp *jsonrpc.RPCResponse) (bool, error) {

	mirrormsg, err := json.Marshal(rpcResp)
	if err != nil {
		logger.Error("verifyStateMsg() error:", err.Error())
		return false, err
	}
	logger.Info("verifyStateMsg() mirrormsg:", string(mirrormsg))

	var rpcRespState = new(RPCResponseState)
	err = json.Unmarshal(mirrormsg, &rpcRespState)
	if err != nil {
		logger.Error("verifyStateMsg() error:", err.Error())
		return false, err
	}

	id := rpcRespState.ID
	logger.Info("verifyStateMsg() rpcRespState.id:", id)
	jsonrpc := rpcRespState.JSONRPC
	logger.Info("verifyStateMsg() rpcRespState.jsonrpc:", jsonrpc)
	rpcresult := rpcRespState.Result
	state := rpcresult.State
	logger.Info("verifyStateMsg() rpcRespState.Result.state:", state)

	return true, nil
}

func verifyGetBinaryMsg(rpcResp *jsonrpc.RPCResponse) (bool, error) {

	mirrormsg, err := json.Marshal(rpcResp)
	if err != nil {
		logger.Error("verifyGetBinaryMsg() error:", err.Error())
		return false, err
	}
	logger.Info("verifyGetBinaryMsg() mirrormsg:", string(mirrormsg))

	var rpcRespPic = new(RPCResponsePic)
	err = json.Unmarshal(mirrormsg, &rpcRespPic)
	if err != nil {
		logger.Error("verifyGetBinaryMsg() error:", err.Error())
		return false, err
	}

	id := rpcRespPic.ID
	logger.Info("verifyGetBinaryMsg() rpcRespState.id:", id)
	jsonrpc := rpcRespPic.JSONRPC
	logger.Info("verifyGetBinaryMsg() rpcRespState.jsonrpc:", jsonrpc)
	rpcresult := rpcRespPic.Result
	pic := rpcresult.Pic
	logger.Info("verifyGetBinaryMsg() rpcRespState.Result.pic:", pic)

	return true, nil
}

/*type State_Resp_Msg struct {
	ID             string `json:"id"`
	Jianyanxiangmu string `json:"jianyanxiangmu"`
	Jiliangdanwei  string `json:"jiliangdanwei"`
	Biaozhunyaoqiu string `json:"biaozhunyaoqiu"`
	Pic1           string `json:"pic1"`
	Pic2           string `json:"pic2"`
}*/
func getPics(rpcResp *jsonrpc.RPCResponse) (string, string, error) {

	mirrormsg, err := json.Marshal(rpcResp)
	if err != nil {
		logger.Error("getPics() error:", err.Error())
		return "", "", err
	}

	logger.Info("getPics() mirrormsg:", string(mirrormsg))

	var rpcRespState = new(RPCResponseState)
	err = json.Unmarshal(mirrormsg, &rpcRespState)
	if err != nil {
		logger.Error("getPics() error:", err.Error())
		return "", "", err
	}

	id := rpcRespState.ID
	logger.Info("getPics() rpcRespState.id:", id)
	jsonrpc := rpcRespState.JSONRPC
	logger.Info("getPics() rpcRespState.jsonrpc:", jsonrpc)
	rpcresult := rpcRespState.Result
	state := rpcresult.State
	logger.Info("getPics() rpcRespState.Result.state:", state)
	var stateRespMsg = new(State_Resp_Msg)
	err = json.Unmarshal([]byte(state), &stateRespMsg)
	//don't check error here, since some source-state message might have no pictures.
	pic1 := stateRespMsg.Pic1
	logger.Info("getPic1() rpcRespState.pic1:", pic1)
	pic2 := stateRespMsg.Pic2
	logger.Info("getPic1() rpcRespState.pic2:", pic2)
	return pic1, pic2, nil
}

/*
func getPics2(rpcResp *jsonrpc.RPCResponse) (string, string, error) {

	mirrormsg, err := json.Marshal(rpcResp)
	if err != nil {
		logger.Error("getPics2() error:", err.Error())
		return "", "", err
	}

	logger.Info("getPics2() mirrormsg:", string(mirrormsg))

	var rpcRespState = new(RPCResponseState)
	err = json.Unmarshal(mirrormsg, &rpcRespState)
	if err != nil {
		logger.Error("getPics2() error:", err.Error())
		return "", "", err
	}

	id := rpcRespState.ID
	logger.Info("getPics2() rpcRespState.id:", id)
	jsonrpc := rpcRespState.JSONRPC
	logger.Info("getPics2() rpcRespState.jsonrpc:", jsonrpc)
	rpcresult := rpcRespState.Result
	state := rpcresult.State
	logger.Info("getPics2() rpcRespState.Result.state:", state)
	var stateRespMsg = new(State_Resp_Msg)
	err = json.Unmarshal([]byte(state), &stateRespMsg)
	//don't check error here, since some source-state message might have no pictures.
	pic1 := stateRespMsg.Pic1
	logger.Info("getPic2() rpcRespState.pic1:", pic1)
	pic2 := stateRespMsg.Pic2
	logger.Info("getPic2() rpcRespState.pic2:", pic2)
	return pic1, pic2, nil
}
*/
//verifyTransactionMsg is to parse and verify the format of source-state message.
func verifyTransactionMsg(rpcResp *jsonrpc.RPCResponse) (bool, error) {

	mirrormsg, err := json.Marshal(rpcResp)
	if err != nil {
		logger.Error("verifyTransactionMsg() error:", err.Error())
		return false, err
	}
	logger.Info("verifyTransactionMsg() mirrormsg:", string(mirrormsg))

	var rpcRespTx = new(RPCResponseTransaction)
	err = json.Unmarshal(mirrormsg, &rpcRespTx)
	if err != nil {
		logger.Error("verifyTransactionMsg() error:", err.Error())
		return false, err
	}
	id := rpcRespTx.ID
	logger.Info("verifyTransactionMsg() rpcRespTx.id:", id)
	jsonrpc := rpcRespTx.JSONRPC
	logger.Info("verifyTransactionMsg() rpcRespTx.jsonrpc:", jsonrpc)
	rpcresult := rpcRespTx.Result
	/*
		type ResultTransaction struct {
			Channel_id string    `json:"channel_id"`
			Data       string    `json:"data"`
			Timestamp  Timestamp `json:"timestamp"`
			Tx_id      string    `json:"tx_id"`
			Type       string    `json:"type"`
		}
		type DataItem struct {
			Write WriteData `json:"write"`
		}
		type WriteData struct {
			Is_delete bool   `json:"is_delete"`
			Key       string `json:"key"`
			Value     string `json:"value"`
		}
	*/
	channel_id := rpcresult.Channel_id
	logger.Info("verifyTransactionMsg() rpcRespTx.Result.channel_id:", channel_id)
	timestamp := rpcresult.Timestamp
	logger.Info("verifyTransactionMsg() rpcRespTx.Result.timestamp:", timestamp)
	tx_id := rpcresult.Tx_id
	logger.Info("verifyTransactionMsg() rpcRespTx.Result.tx_id:", tx_id)
	resulttype := rpcresult.Type
	logger.Info("verifyTransactionMsg() rpcRespTx.Result.resulttype:", resulttype)
	data := rpcresult.Data
	write := data[0].Write
	is_delete := write.Is_delete
	logger.Info("verifyTransactionMsg() rpcRespTx.Result.data.is_delete:", is_delete)
	key := write.Key
	logger.Info("verifyTransactionMsg() rpcRespTx.Result.data.key:", key)
	value := write.Value
	logger.Info("verifyTransactionMsg() rpcRespTx.Result.data.value:", value)
	return true, nil
}

//verifyTransactionsMsg is to parse and verify the format of source-state message.
func verifyTransactionsMsg(rpcResp *jsonrpc.RPCResponse) (bool, error) {

	mirrormsg, err := json.Marshal(rpcResp)
	if err != nil {
		logger.Error("verifyTransactionsMsg() error:", err.Error())
		return false, err
	}
	logger.Info("verifyTransactionMsgs() mirrormsg:", string(mirrormsg))

	var rpcRespTx = new(RPCResponseTransactions)
	err = json.Unmarshal(mirrormsg, &rpcRespTx)
	if err != nil {
		logger.Error("verifyTransactionsMsg() error:", err.Error())
		return false, err
	}

	id := rpcRespTx.ID
	logger.Info("verifyTransactionMsgs() rpcRespTx.id:", id)
	jsonrpc := rpcRespTx.JSONRPC
	logger.Info("verifyTransactionMsgs() rpcRespTx.jsonrpc:", jsonrpc)
	rpcresults := rpcRespTx.Result
	for i := range rpcresults {
		//表示遍历数组，而i表示的是数组的下标值，
		//result[i]表示获得第i个json对象即JSONObject
		//result[i]通过.字段名称即可获得指定字段的值
		tx_id := rpcresults[i].Tx_id
		logger.Info("verifyTransactionsMsg() rpcRespTx.Result.Tx_id:", tx_id)
		value := rpcresults[i].Value
		logger.Info("verifyTransactionsMsg() rpcRespTx.Result.Value:", value)
		timestamp := rpcresults[i].Timestamp
		nanos := timestamp.Nanos
		seconds := timestamp.Seconds
		logger.Info("verifyTransactionMsgs() rpcRespTx.Result.Timestamp.Seconds:", seconds, "Nanos:", nanos)
	}

	return true, nil
}

func verifyMsg(method string, rpcResp *jsonrpc.RPCResponse) (bool, error) {
	if method == "source-state" {
		return verifyStateMsg(rpcResp)
	}
	if method == "source-transactions" {
		return verifyTransactionsMsg(rpcResp)
	}
	if method == "source-transaction" {
		return verifyTransactionMsg(rpcResp)
	}
	if method == "source-get-binary" {
		return verifyGetBinaryMsg(rpcResp)
	}

	//let the client to handle the result error in jsonrpc.
	//won't handle the error here.
	return true, nil
	//return false, nil
}

func handle_big_message(respMsg []byte) ([]byte, error) {
	/*type State_Resp_Msg struct {
		ID             string `json:"id"`
		Jianyanxiangmu string `json:"jianyanxiangmu"`
		Jiliangdanwei  string `json:"jiliangdanwei"`
		Biaozhunyaoqiu string `json:"biaozhunyaoqiu"`
		Pic1           string `json:"pic1"`
		Pic2           string `json:"pic2"`
	}*/

	var rpcRespState = new(RPCResponseState)
	err := json.Unmarshal(respMsg, &rpcRespState)
	if err != nil {
		logger.Error("handle_big_message(): error.....")
		return nil, err
	}

	id := rpcRespState.ID
	logger.Info("handle_big_message() rpcRespState.id:", id)
	jsonrpc := rpcRespState.JSONRPC
	logger.Info("handle_big_message() rpcRespState.jsonrpc:", jsonrpc)
	rpcresult := rpcRespState.Result
	state := rpcresult.State
	logger.Info("handle_big_message() rpcRespState.Result.state:", state)
	var stateRespMsg = new(State_Resp_Msg)
	json.Unmarshal([]byte(state), &stateRespMsg)
	pic1 := stateRespMsg.Pic1
	logger.Info("handle_big_message() rpcRespState.pic1:", pic1)
	pic2 := stateRespMsg.Pic2
	logger.Info("handle_big_message() rpcRespState.pic2:", pic2)
	stateId := stateRespMsg.ID
	logger.Info("handle_big_message() rpcRespState.stateId:", stateId)
	jianyanxiangmu := stateRespMsg.Jianyanxiangmu
	logger.Info("handle_big_message() rpcRespState.jianyanxiangmu:", jianyanxiangmu)
	jiliangdanwei := stateRespMsg.Jiliangdanwei
	logger.Info("handle_big_message() rpcRespState.jiliangdanwei:", jiliangdanwei)
	biaozhunyaoqiu := stateRespMsg.Biaozhunyaoqiu
	logger.Info("handle_big_message() rpcRespState.biaozhunyaoqiu:", biaozhunyaoqiu)

	logger.Info("handle_big_message() pic1:", pic1)
	method := "source-get-binary"
	tx_id := ""
	rpcResp, err := sendJsonrpcRequest(method, pic1, tx_id)
	if err != nil {
		logger.Error("handle_big_message():sendJsonrpcRequest error.....")
		return nil, err
	}
	respMsg, err = json.Marshal(rpcResp)
	if err != nil {
		logger.Error("handle_big_message(): error.....")
		return nil, err
	}

	logger.Info("handle_big_message() echo the message:", string(respMsg))
	var rpcRespPic = new(RPCResponsePic)
	err = json.Unmarshal(respMsg, &rpcRespPic)
	if err != nil {
		logger.Error("handle_big_message(): error.....")
		return nil, err
	}

	rpcresult2 := rpcRespPic.Result
	pic1_obj := rpcresult2.Pic
	logger.Info("handle_big_message() rpcRespState.Result.pic1_obj:", pic1_obj)

	logger.Info("handle_big_message() pic2:", pic2)
	method = "source-get-binary"
	tx_id = ""
	rpcResp, err = sendJsonrpcRequest(method, pic2, tx_id)
	if err != nil {
		logger.Error("handle_big_message(): error.....")
		return nil, err
	}
	respMsg, err = json.Marshal(rpcResp)
	if err != nil {
		logger.Error("handle_big_message(): error.....")
		return nil, err
	}
	logger.Info("handle_big_message() echo the message:", string(respMsg))

	var rpcRespPic2 = new(RPCResponsePic)
	err = json.Unmarshal(respMsg, &rpcRespPic2)
	if err != nil {
		logger.Error("handle_big_message(): error.....")
		return nil, err
	}
	rpcresult3 := rpcRespPic2.Result
	pic2_obj := rpcresult3.Pic
	logger.Info("handle_big_message() rpcRespState.Result.pic2_obj:", pic2_obj)

	stateX := make(map[string]interface{})
	stateX["pic1"] = pic1_obj
	stateX["pic2"] = pic2_obj
	stateX["id"] = stateId
	stateX["jianyanxiangmu"] = jianyanxiangmu
	stateX["jiliangdanwei"] = jiliangdanwei
	stateX["biaozhunyaoqiu"] = biaozhunyaoqiu

	resultX := make(map[string]interface{})
	resultX["state"] = stateX

	response := make(map[string]interface{})
	response["id"] = 0
	response["jsonrpc"] = "2.0"
	response["result"] = resultX

	respMsg, err = json.Marshal(response)
	if err != nil {
		logger.Error("handle_big_message(): error.....")
		return nil, err
	}
	logger.Info("handle_big_message() echo the message:", string(respMsg))
	return respMsg, nil
}

//B_chenhui
func handle_suyuan_message(respMsg []byte, readcount uint64, accesstime uint64) ([]byte, error) {
	/*
			type BatchInformationType struct {
			BatchNumber                             string                                      `json:"batchNumber"`
			BatchOutput                             string                                      `json:"batchOutput"`
			SeedInfo                                SeedInfoType                                `json:"seedInfo"`                                //一、种子信息：
			BiologicalOrganicFertilizer             BiologicalOrganicFertilizerType             `json:"biologicalOrganicFertilizer"`             //二、生物有机肥
			OrganicCertificationOfBase              OrganicCertificationOfBaseType              `json:"organicCertificationOfBase"`              //三、基地有机认证
			InspectionReport                        InspectionReportType                        `json:"inspectionReport"`                        //四、检验报告
			PatentInfo                              PatentInfoType                              `json:"patentInfo"`                              //五、加工工艺专利技术证书
			DetectionReportOfEmbryoRateAndIntegrity DetectionReportOfEmbryoRateAndIntegrityType `json:"detectionReportOfEmbryoRateAndIntegrity"` //六、留胚率、完整度检测报告
			PositionInformation                     PositionInformationType                     `json:"positionInformation"`                     //七、位置信息
			ProductInspectionReport                 ProductInspectionReportType                 `json:"productInspectionReport"`                 //八、产品出厂检测报告
		}
	*/

	var rpcRespState = new(RPCResponseState2)
	err := json.Unmarshal(respMsg, &rpcRespState)
	if err != nil {
		logger.Error("handle_suyuan_message(): error.....")
		return nil, err
	}

	id := rpcRespState.ID
	logger.Info("handle_suyuan_message() rpcRespState.id:", id)
	jsonrpc := rpcRespState.JSONRPC
	logger.Info("handle_suyuan_message() rpcRespState.jsonrpc:", jsonrpc)
	rpcresult := rpcRespState.Result
	value := rpcresult.Value
	logger.Info("handle_suyuan_message() rpcRespState.Result.state:", value)
	var stateRespMsg = new(BatchInformationType)
	json.Unmarshal([]byte(value), &stateRespMsg)
	batchNumber := stateRespMsg.BatchNumber
	logger.Info("handle_suyuan_message() rpcRespState.batchNumber:", batchNumber)
	batchOutput := stateRespMsg.BatchOutput
	logger.Info("handle_suyuan_message() rpcRespState.batchOutput:", batchOutput)

	stateRespMsg.ReadCount = readcount
	stateRespMsg.AccessTime = accesstime

	//1.seedinfo
	seedInfo := stateRespMsg.SeedInfo
	grade := seedInfo.Grade
	logger.Info("handle_suyuan_message() rpcRespState.SeedInfo.grade:", grade)

	//2.BiologicalOrganicFertilizer

	//3.OrganicCertificationOfBase
	url_prefix := "/home/mengchun/go/src/httpproxy4blockchain"
	organicCertificationOfBase := stateRespMsg.OrganicCertificationOfBase
	pic1 := organicCertificationOfBase.PictureName
	logger.Info("handle_suyuan_message() rpcRespState.OrganicCertificationOfBase.PictureName:", pic1)

	//B_chenhui0704
	pic1xx := url_prefix + pic1
	//B_chenhui0704

	/*
		ff, _ := ioutil.ReadFile(pic1)          //我还是喜欢用这个快速读文件
		bufstore := make([]byte, 5000000)       //数据缓存
		base64.StdEncoding.Encode(bufstore, ff) // 文件转base64
		//_ = ioutil.WriteFile("output2/1.full.txt", bufstore, 0666) //直接写入到文件就ok完活了。
		index := bytes.IndexByte(bufstore, 0)
		rbyf_pn := bufstore[0:index]
		_ = ioutil.WriteFile("output2/1.jpg.txt", rbyf_pn, 0666) //直接写入到文件就ok完活了。
	*/

	stateRespMsg.OrganicCertificationOfBase.PictureName = pic1xx
	pic1x := organicCertificationOfBase.PictureName
	logger.Info("handle_suyuan_message() rpcRespState.OrganicCertificationOfBase.PictureName:", pic1x)
	//logger.Info("handle_suyuan_message() rpcRespState.OrganicCertificationOfBase.PictureName real:", stateRespMsg.OrganicCertificationOfBase.PictureName)

	//4.检测报告，InspectionReport
	inspectionReport := stateRespMsg.InspectionReport
	pic2 := inspectionReport.Picture1Cover
	logger.Info("handle_suyuan_message() rpcRespState.InspectionReport.Picture1Cover:", pic2)
	/*
		ff, _ = ioutil.ReadFile(pic2)           //我还是喜欢用这个快速读文件
		bufstore = make([]byte, 5000000)        //数据缓存
		base64.StdEncoding.Encode(bufstore, ff) // 文件转base64
		index = bytes.IndexByte(bufstore, 0)
		rbyf_pn = bufstore[0:index]
		_ = ioutil.WriteFile("output2/2.jpg.txt", rbyf_pn, 0666) //直接写入到文件就ok完活了。
	*/
	//B_chenhui0704
	pic2xx := url_prefix + pic2
	//B_chenhui0704

	stateRespMsg.InspectionReport.Picture1Cover = pic2xx
	//pic2x := inspectionReport.Picture1Cover
	//logger.Info("handle_suyuan_message() rpcRespState.inspectionReport.Picture1Cover:", pic2x)

	pic3 := inspectionReport.PictureBaseSoil
	logger.Info("handle_suyuan_message() rpcRespState.InspectionReport.PictureBaseSoil:", pic3)
	/*
		ff, _ = ioutil.ReadFile(pic3)           //我还是喜欢用这个快速读文件
		bufstore = make([]byte, 5000000)        //数据缓存
		base64.StdEncoding.Encode(bufstore, ff) // 文件转base64
		index = bytes.IndexByte(bufstore, 0)
		rbyf_pn = bufstore[0:index]
		_ = ioutil.WriteFile("output2/3.jpg.txt", rbyf_pn, 0666) //直接写入到文件就ok完活了。
	*/
	//B_chenhui0704
	pic3xx := url_prefix + pic3
	//B_chenhui0704
	stateRespMsg.InspectionReport.PictureBaseSoil = pic3xx
	//pic3x := inspectionReport.PictureBaseSoil
	//logger.Info("handle_suyuan_message() rpcRespState.inspectionReport.PictureBaseSoil:", pic3x)

	pic4 := inspectionReport.PictureIrrigatedWaterSource
	logger.Info("handle_suyuan_message() rpcRespState.InspectionReport.PictureIrrigatedWaterSource:", pic4)
	/*
		ff, _ = ioutil.ReadFile(pic4)           //我还是喜欢用这个快速读文件
		bufstore = make([]byte, 5000000)        //数据缓存
		base64.StdEncoding.Encode(bufstore, ff) // 文件转base64
		index = bytes.IndexByte(bufstore, 0)
		rbyf_pn = bufstore[0:index]
		_ = ioutil.WriteFile("output2/4.jpg.txt", rbyf_pn, 0666) //直接写入到文件就ok完活了。
	*/
	//B_chenhui0704
	pic4xx := url_prefix + pic4
	//B_chenhui0704
	stateRespMsg.InspectionReport.PictureIrrigatedWaterSource = pic4xx

	//5、加工工艺专利技术证书,PatentInfo
	patentInfo := stateRespMsg.PatentInfo
	patentInfoItems := patentInfo.Items
	for i := range patentInfoItems {
		pic5678 := patentInfoItems[i].PictureName
		logger.Info("handle_suyuan_message() rpcRespState.PatentInfo.PictureName:", pic5678)
		/*
			ff, _ = ioutil.ReadFile(pic5678)        //我还是喜欢用这个快速读文件
			bufstore = make([]byte, 5000000)        //数据缓存
			base64.StdEncoding.Encode(bufstore, ff) // 文件转base64
			index = bytes.IndexByte(bufstore, 0)
			rbyf_pn = bufstore[0:index]
			//_ = ioutil.WriteFile("output2/5.jpg.txt", rbyf_pn, 0666) //直接写入到文件就ok完活了。
		*/
		//B_chenhui0704
		pic5678xx := url_prefix + pic5678
		//B_chenhui0704
		stateRespMsg.PatentInfo.Items[i].PictureName = pic5678xx
	}

	//6、留胚率、完整度检测报告,DetectionReportOfEmbryoRateAndIntegrity
	detectionReportOfEmbryoRateAndIntegrity := stateRespMsg.DetectionReportOfEmbryoRateAndIntegrity
	pic9 := detectionReportOfEmbryoRateAndIntegrity.Picture1
	logger.Info("handle_suyuan_message() rpcRespState.DetectionReportOfEmbryoRateAndIntegrity.Picture1:", pic9)
	/*
		ff, _ = ioutil.ReadFile(pic9)           //我还是喜欢用这个快速读文件
		bufstore = make([]byte, 5000000)        //数据缓存
		base64.StdEncoding.Encode(bufstore, ff) // 文件转base64
		index = bytes.IndexByte(bufstore, 0)
		rbyf_pn = bufstore[0:index]
		_ = ioutil.WriteFile("output2/9.jpg.txt", rbyf_pn, 0666) //直接写入到文件就ok完活了。
	*/
	//B_chenhui0704
	pic9xx := url_prefix + pic9
	//B_chenhui0704
	stateRespMsg.DetectionReportOfEmbryoRateAndIntegrity.Picture1 = pic9xx

	pic10 := detectionReportOfEmbryoRateAndIntegrity.Picture2
	logger.Info("handle_suyuan_message() rpcRespState.DetectionReportOfEmbryoRateAndIntegrity.Picture2:", pic10)
	/*
		ff, _ = ioutil.ReadFile(pic10)          //我还是喜欢用这个快速读文件
		bufstore = make([]byte, 5000000)        //数据缓存
		base64.StdEncoding.Encode(bufstore, ff) // 文件转base64
		index = bytes.IndexByte(bufstore, 0)
		rbyf_pn = bufstore[0:index]
		_ = ioutil.WriteFile("output2/10.jpg.txt", rbyf_pn, 0666) //直接写入到文件就ok完活了。
	*/
	//B_chenhui0704
	pic10xx := url_prefix + pic10
	//B_chenhui0704
	stateRespMsg.DetectionReportOfEmbryoRateAndIntegrity.Picture2 = pic10xx

	pic11 := detectionReportOfEmbryoRateAndIntegrity.Picture3
	logger.Info("handle_suyuan_message() rpcRespState.DetectionReportOfEmbryoRateAndIntegrity.Picture3:", pic11)
	/*
		ff, _ = ioutil.ReadFile(pic11)          //我还是喜欢用这个快速读文件
		bufstore = make([]byte, 5000000)        //数据缓存
		base64.StdEncoding.Encode(bufstore, ff) // 文件转base64
		index = bytes.IndexByte(bufstore, 0)
		rbyf_pn = bufstore[0:index]
		_ = ioutil.WriteFile("output2/11.jpg.txt", rbyf_pn, 0666) //直接写入到文件就ok完活了。
	*/
	//B_chenhui0704
	pic11xx := url_prefix + pic11
	//B_chenhui0704
	stateRespMsg.DetectionReportOfEmbryoRateAndIntegrity.Picture3 = pic11xx

	pic12 := detectionReportOfEmbryoRateAndIntegrity.Picture4
	logger.Info("handle_suyuan_message() rpcRespState.DetectionReportOfEmbryoRateAndIntegrity.Picture4:", pic12)
	/*
		ff, _ = ioutil.ReadFile(pic12)          //我还是喜欢用这个快速读文件
		bufstore = make([]byte, 5000000)        //数据缓存
		base64.StdEncoding.Encode(bufstore, ff) // 文件转base64
		index = bytes.IndexByte(bufstore, 0)
		rbyf_pn = bufstore[0:index]
		_ = ioutil.WriteFile("output2/12.jpg.txt", rbyf_pn, 0666) //直接写入到文件就ok完活了。
	*/
	//B_chenhui0704
	pic12xx := url_prefix + pic12
	//B_chenhui0704
	stateRespMsg.DetectionReportOfEmbryoRateAndIntegrity.Picture4 = pic12xx

	//7.位置信息,PositionInformation
	positionInformation := stateRespMsg.PositionInformation
	//pic13 := positionInformation.Picture1PlantingBaseLocation
	pos1 := positionInformation.Position1PlantingBaseLocation
	logger.Info("handle_suyuan_message() rpcRespState.positionInformation.Picture1PlantingBaseLocation:", pos1)
	//ff, _ = ioutil.ReadFile(pic13)          //我还是喜欢用这个快速读文件
	//bufstore = make([]byte, 5000000)        //数据缓存
	//base64.StdEncoding.Encode(bufstore, ff) // 文件转base64
	//index = bytes.IndexByte(bufstore, 0)
	//rbyf_pn = bufstore[0:index]
	//_ = ioutil.WriteFile("output2/13.jpg.txt", rbyf_pn, 0666) //直接写入到文件就ok完活了。
	//stateRespMsg.PositionInformation.Picture1PlantingBaseLocation = string(rbyf_pn)

	//pic14 := positionInformation.Picture2StorageBaseLocation
	pos2 := positionInformation.Position2StorageBaseLocation
	logger.Info("handle_suyuan_message() rpcRespState.positionInformation.Position2StorageBaseLocation:", pos2)
	/*
		ff, _ = ioutil.ReadFile(pic14)          //我还是喜欢用这个快速读文件
		bufstore = make([]byte, 5000000)        //数据缓存
		base64.StdEncoding.Encode(bufstore, ff) // 文件转base64
		index = bytes.IndexByte(bufstore, 0)
		rbyf_pn = bufstore[0:index]
		_ = ioutil.WriteFile("output2/14.jpg.txt", rbyf_pn, 0666) //直接写入到文件就ok完活了。
		stateRespMsg.PositionInformation.Picture2StorageBaseLocation = string(rbyf_pn)
	*/

	//pic15 := positionInformation.Picture3ProcessingBaseLocation
	pos3 := positionInformation.Position3ProcessingBaseLocation
	logger.Info("handle_suyuan_message() rpcRespState.positionInformation.Picture3ProcessingBaseLocation:", pos3)
	/*
		ff, _ = ioutil.ReadFile(pic15)          //我还是喜欢用这个快速读文件
		bufstore = make([]byte, 5000000)        //数据缓存
		base64.StdEncoding.Encode(bufstore, ff) // 文件转base64
		index = bytes.IndexByte(bufstore, 0)
		rbyf_pn = bufstore[0:index]
		_ = ioutil.WriteFile("output2/15.jpg.txt", rbyf_pn, 0666) //直接写入到文件就ok完活了。
		stateRespMsg.PositionInformation.Picture3ProcessingBaseLocation = string(rbyf_pn)
	*/

	/*
			rpcresults := rpcRespTx.Result
		for i := range rpcresults {
			//表示遍历数组，而i表示的是数组的下标值，
			//result[i]表示获得第i个json对象即JSONObject
			//result[i]通过.字段名称即可获得指定字段的值
			tx_id := rpcresults[i].Tx_id
			logger.Info("verifyTransactionsMsg() rpcRespTx.Result.Tx_id:", tx_id)
			value := rpcresults[i].Value
			logger.Info("verifyTransactionsMsg() rpcRespTx.Result.Value:", value)
			timestamp := rpcresults[i].Timestamp
			nanos := timestamp.Nanos
			seconds := timestamp.Seconds
			logger.Info("verifyTransactionMsgs() rpcRespTx.Result.Timestamp.Seconds:", seconds, "Nanos:", nanos)
		}
	*/
	//pic4x := inspectionReport.PictureIrrigatedWaterSource
	//logger.Info("handle_suyuan_message() rpcRespState.inspectionReport.PictureIrrigatedWaterSource:", pic4x)
	/*
		method := "source-get-binary"
		tx_id := ""
		rpcResp, err := sendJsonrpcRequest(method, pic1, tx_id)
		if err != nil {
			logger.Error("handle_suyuan_message():sendJsonrpcRequest error.....")
			return nil, err
		}
		respMsg, err = json.Marshal(rpcResp)
		if err != nil {
			logger.Error("handle_suyuan_message(): error.....")
			return nil, err
		}

		logger.Info("handle_suyuan_message() echo the message:", string(respMsg))
		var rpcRespPic = new(RPCResponsePic)
		err = json.Unmarshal(respMsg, &rpcRespPic)
		if err != nil {
			logger.Error("handle_suyuan_message(): error.....")
			return nil, err
		}

		rpcresult2 := rpcRespPic.Result
		pic1_obj := rpcresult2.Pic
		logger.Info("handle_suyuan_message() rpcRespState.Result.pic1_obj:", pic1_obj)

		logger.Info("handle_suyuan_message() pic2:", pic2)
		method = "source-get-binary"
		tx_id = ""
		rpcResp, err = sendJsonrpcRequest(method, pic2, tx_id)
		if err != nil {
			logger.Error("handle_suyuan_message(): error.....")
			return nil, err
		}
		respMsg, err = json.Marshal(rpcResp)
		if err != nil {
			logger.Error("handle_suyuan_message(): error.....")
			return nil, err
		}
		logger.Info("handle_suyuan_message() echo the message:", string(respMsg))

		var rpcRespPic2 = new(RPCResponsePic)
		err = json.Unmarshal(respMsg, &rpcRespPic2)
		if err != nil {
			logger.Error("handle_suyuan_message(): error.....")
			return nil, err
		}
		rpcresult3 := rpcRespPic2.Result
		pic2_obj := rpcresult3.Pic
		logger.Info("handle_suyuan_message() rpcRespState.Result.pic2_obj:", pic2_obj)

		stateX := make(map[string]interface{})
		stateX["pic1"] = pic1_obj
		stateX["pic2"] = pic2_obj
		stateX["id"] = stateId
		stateX["jianyanxiangmu"] = jianyanxiangmu
		stateX["jiliangdanwei"] = jiliangdanwei
		stateX["biaozhunyaoqiu"] = biaozhunyaoqiu

		resultX := make(map[string]interface{})
		resultX["state"] = stateX

		response := make(map[string]interface{})
		response["id"] = 0
		response["jsonrpc"] = "2.0"
		response["result"] = resultX

		respMsg, err = json.Marshal(response)
		if err != nil {
			logger.Error("handle_big_message(): error.....")
			return nil, err
		}
		logger.Info("handle_big_message() echo the message:", string(respMsg))
	*/
	respMsg, err = json.Marshal(stateRespMsg)
	if err != nil {
		logger.Error("handle_suyuan_message(): error.....")
		return nil, err
	}
	logger.Info("handle_suyuan_message() echo the message:", string(respMsg))
	return respMsg, nil
}

//E_chenhui

//Excute is the function that each Controller needs to implement.
func Excute(message []byte) ([]byte, error) {
	//mirrormsg, err := json.Marshal(message)

	var rpcRequest jsonrpc.RPCRequest
	err := json.Unmarshal(message, &rpcRequest)
	if err != nil {
		logger.Error("Excute() Unmarshal err,", err)
		return nil, err
	}

	//rpcRequest := entermsg.Content
	logger.Info("parsing the JSONRPC2.0 message from app client...")

	method := rpcRequest.Method
	logger.Info("Excute() parsing Method:", method)
	f := rpcRequest.Params
	//B_chenhui
	var key string
	if strings.Contains(method, "source") {
		key = f.(map[string]interface{})["key"].(string)
		logger.Info("Excute() rpcRequest.Params.Key:", key)
	}
	//E_chenhui
	channel := f.(map[string]interface{})["channel"].(string)
	logger.Info("Excute() rpcRequest.Params.Channel:", channel)
	var tx_id string
	if method == "source-transaction" {
		tx_id = f.(map[string]interface{})["tx_id"].(string)
		logger.Info("Excute() rpcRequest.Params.tx_id:", tx_id)
	}

	//B_chenhui
	var rpcResp *jsonrpc.RPCResponse
	if strings.Contains(method, "source") {
		rpcResp, err = sendJsonrpcRequest(method, key, tx_id)
		if err != nil {
			logger.Error("Excute() sendJsonrpcRequest ,", err)
			return nil, err
		}
	} else {
		rpcResp, err = sendJsonrpcRequest4Asset(method, message)
		if err != nil {
			logger.Error("Excute() sendJsonrpcRequest4Asset ,", err)
			return nil, err
		}
	}
	//E_chenhui

	respMsg, err := json.Marshal(rpcResp)
	if err != nil {
		logger.Error("Excute() Marshal ,", err)
		return nil, err
	}

	logger.Info("Excute() echo the message:", string(respMsg))

	isok, err := verifyMsg(method, rpcResp)
	if err != nil || !isok {
		logger.Error("Excute() verifyMsg ,", err)
		return nil, err
	}
	//B_chenhui
	if strings.Contains(method, "source") {

		if len(key) == 21 {
			respMsg, err = handle_suyuan_message(respMsg, 0, 0)
			if err != nil {
				logger.Error("Excute() handle_suyuan_message error::", err)
				return nil, err
			}
			logger.Info("Excute() echo the message:", string(respMsg))
		}
		if len(key) == 31 {
			var rpcRespState = new(RPCResponseState3)
			err := json.Unmarshal(respMsg, &rpcRespState)
			if err != nil {
				logger.Error("Excute(): error.....")
				return nil, err
			}

			id := rpcRespState.ID
			logger.Info("Excute() rpcRespState.id:", id)
			jsonrpc := rpcRespState.JSONRPC
			logger.Info("Excute() rpcRespState.jsonrpc:", jsonrpc)
			rpcresult := rpcRespState.Result
			//value := rpcresult.Value
			logger.Info("Excute() rpcRespState.Result:", rpcresult)

			value2 := rpcresult.Value
			logger.Info("Excute() stateRespMsg.Value:", value2)
			readCountInfo := rpcresult.ReadCount
			accesstime := readCountInfo.AccessTime
			logger.Info("Excute() stateRespMsg.ReadCountInfo.AccessTime:", accesstime)
			readcount := readCountInfo.ReadCount
			logger.Info("Excute() stateRespMsg.ReadCountInfo.ReadCount:", readcount)

			batchKey := key[0:21]
			logger.Info("Excute() batchKey is:", batchKey)
			tx_id = ""
			rpcResp, err = sendJsonrpcRequest(method, batchKey, tx_id)
			if err != nil {
				logger.Error("Excute() Marshal ,", err)
				return nil, err
			}
			respMsg, err = json.Marshal(rpcResp)
			if err != nil {
				logger.Error("Excute() Marshal ,", err)
				return nil, err
			}
			respMsg, err = handle_suyuan_message(respMsg, readcount, accesstime)
			if err != nil {
				logger.Error("Excute() handle_suyuan_message error::", err)
				return nil, err
			}
			logger.Info("Excute() echo the message:", string(respMsg))
		}
	}
	//E_chenhui
	return respMsg, nil
}
