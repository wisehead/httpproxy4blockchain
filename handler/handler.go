package handler

import (
	"encoding/json"
	"httpproxy4blockchain/jsonrpc"
	"httpproxy4blockchain/logger"
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
	if method == "source-transaction" {
		rpcResp, err = rpcClient.Call(method, &MethodParams{Channel: "notaryinfotestchannel", Key: key, Tx_id: tx_id})
	} else {
		rpcResp, err = rpcClient.Call(method, &MethodParams{Channel: "notaryinfotestchannel", Key: key})
	}

	if err != nil {
		logger.Error("sendJsonrpcRequest() err for rpcClient.Call:", err.Error())
		return nil, err
	}
	return rpcResp, nil
}

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
		logger.Error("handle_big_message(): error.....") //chenhui
		return nil, err
	}

	id := rpcRespState.ID
	//logger.Info("xxx Excute() rpcRespState.id:", id)
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
		logger.Error("handle_big_message(): error.....") //chenhui
		return nil, err
	}
	logger.Info("Excute() echo the message:", string(respMsg))

	var rpcRespPic2 = new(RPCResponsePic)
	err = json.Unmarshal(respMsg, &rpcRespPic2)
	if err != nil {
		logger.Error("handle_big_message(): error.....") //chenhui
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
	key := f.(map[string]interface{})["key"].(string)
	logger.Info("Excute() rpcRequest.Params.Key:", key)
	channel := f.(map[string]interface{})["channel"].(string)
	logger.Info("Excute() rpcRequest.Params.Channel:", channel)
	var tx_id string
	if method == "source-transaction" {
		tx_id = f.(map[string]interface{})["tx_id"].(string)
		logger.Info("Excute() rpcRequest.Params.tx_id:", tx_id)
	}

	rpcResp, err := sendJsonrpcRequest(method, key, tx_id)
	if err != nil {
		logger.Error("Excute() sendJsonrpcRequest ,", err)
		return nil, err
	}

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

	//check whether there is a picture in the record.
	var pic1 string
	var pic2 string
	if method == "source-state" {
		pic1, pic2, err = getPics(rpcResp)
		if err != nil {
			logger.Error("Excute() getPics error::", err)
			return nil, err
		}
	}

	respMsg, err = json.Marshal(rpcResp)
	if err != nil {
		logger.Error("Excute() Marshal error::", err)
		return nil, err
	}

	logger.Info("Excute() echo the message:", string(respMsg))
	if method == "source-state" {
		if pic1 != "" || pic2 != "" {
			respMsg, err = handle_big_message(respMsg)
			if err != nil {
				logger.Error("Excute() handle_big_message error::", err)
				return nil, err
			}
			logger.Info("Excute() echo the message:", string(respMsg))
		}
	}

	return respMsg, nil
}
