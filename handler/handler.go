package handler

import (
	"encoding/json"
	"httpproxy4blockchain/goproxy4blockchain/utils"
	"httpproxy4blockchain/jsonrpc"
	"log"
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

//Timestamp is to describ the timestamp.
type Timestamp struct {
	Nanos   uint32 `json:"nanos"`
	Seconds uint32 `json:"seconds"`
}

//ResultTransaction is the struct of result of source-transactions message.
type ResultTransaction struct {
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
type RPCResponseTransaction struct {
	JSONRPC string              `json:"jsonrpc"`
	Result  []ResultTransaction `json:"result,omitempty"`
	//Result map[string]interface{} `json:"result,omitempty"`
	//Result *json.RawMessage `json:"result,omitempty"`
	Error *RPCError `json:"error,omitempty"`
	ID    uint      `json:"id"`
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
}

//Msg defined between app client and goproxy4blockchain
type Msg struct {
	Meta    map[string]interface{} `json:"meta"`
	Content jsonrpc.RPCRequest     `json:"content"`
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

//sendJsonrpcRequest is to send request to block chain service.
func sendJsonrpcRequest(method string, key string) (*jsonrpc.RPCResponse, error) {
	var err error
	//rpcClient := jsonrpc.NewClient("http://my-rpc-service:8080/rpc")
	rpcClient := jsonrpc.NewClient("https://www.ninechain.net/api/v2.1")
	if rpcClient == nil {
		log.Println("rxxx sendJsonrpcRequest() pcClient is nil!")
		return nil, err
	}
	rpcResp, err := rpcClient.Call(method, &MethodParams{Channel: "vvtrip", Key: key})
	if err != nil {
		log.Println("xxx err for rpcClient.Call:", err.Error())
		return nil, err
	}
	//id := rpcResp.ID
	//log.Println("xxx sendJsonrpcRequest() rpcResp.id:", id)
	//jsonrpc := rpcResp.JSONRPC
	//log.Println("xxx sendJsonrpcRequest() rpcResp.jsonrpc:", jsonrpc)
	//rpcresult := rpcResp.Result
	//fix-me:
	//need serveral parser functions to parse different result structs.
	//tx_id := rpcresult["tx_id"].(string)
	//log.Println("xxx sendJsonrpcRequest() rpcResp.Result.tx_id:", tx_id)
	return rpcResp, nil
}

//verifyStateMsg is to parse and verify the format of source-state message.
func verifyStateMsg(rpcResp *jsonrpc.RPCResponse) (bool, error) {

	mirrormsg, err := json.Marshal(rpcResp)
	utils.CheckError(err)
	log.Println("xxx verifyStateMsg() mirrormsg:", string(mirrormsg))

	var rpcRespState = new(RPCResponseState)
	json.Unmarshal(mirrormsg, &rpcRespState)

	id := rpcRespState.ID
	log.Println("xxx verifyStateMsg() rpcRespState.id:", id)
	jsonrpc := rpcRespState.JSONRPC
	log.Println("xxx verifyStateMsg() rpcRespState.jsonrpc:", jsonrpc)
	rpcresult := rpcRespState.Result
	state := rpcresult.State
	log.Println("xxx verifyStateMsg() rpcRespState.Result.state:", state)

	return true, nil
}

//verifyTransactionMsg is to parse and verify the format of source-state message.
func verifyTransactionMsg(rpcResp *jsonrpc.RPCResponse) (bool, error) {

	mirrormsg, err := json.Marshal(rpcResp)
	utils.CheckError(err)
	log.Println("xxx verifyTransactionMsg() mirrormsg:", string(mirrormsg))

	var rpcRespTx = new(RPCResponseTransaction)
	json.Unmarshal(mirrormsg, &rpcRespTx)

	id := rpcRespTx.ID
	log.Println("xxx verifyTransactionMsg() rpcRespTx.id:", id)
	jsonrpc := rpcRespTx.JSONRPC
	log.Println("xxx verifyTransactionMsg() rpcRespTx.jsonrpc:", jsonrpc)
	rpcresults := rpcRespTx.Result
	for i := range rpcresults {
		//表示遍历数组，而i表示的是数组的下标值，
		//result[i]表示获得第i个json对象即JSONObject
		//result[i]通过.字段名称即可获得指定字段的值
		tx_id := rpcresults[i].Tx_id
		log.Println("xxx verifyTransactionMsg() rpcRespTx.Result.Tx_id:", tx_id)
		value := rpcresults[i].Value
		log.Println("xxx verifyTransactionMsg() rpcRespTx.Result.Value:", value)
		timestamp := rpcresults[i].Timestamp
		nanos := timestamp.Nanos
		seconds := timestamp.Seconds
		log.Println("xxx verifyTransactionMsg() rpcRespTx.Result.Timestamp.Seconds:", seconds, "Nanos:", nanos)
	}

	return true, nil
}

func verifyMsg(method string, rpcResp *jsonrpc.RPCResponse) (bool, error) {
	if method == "source-state" {
		return verifyStateMsg(rpcResp)
	} else {
		if method == "source-transactions" {
			return verifyTransactionMsg(rpcResp)
		}
	}
	return false, nil
}

//Excute is the function that each Controller needs to implement.
func Excute(message []byte) []byte {
	//mirrormsg, err := json.Marshal(message)

	var rpcRequest jsonrpc.RPCRequest
	err := json.Unmarshal(message, &rpcRequest)
	if err != nil {
		log.Println(err)
	}

	//rpcRequest := entermsg.Content
	log.Println("xxx parsing the JSONRPC2.0 message from app client...")
	/*
		id := rpcRequest.ID
		Log("xxx rpcRequest.id:", id)
		jsonrpc := rpcRequest.JSONRPC
		Log("xxx rpcRequest.jsonrpc:", jsonrpc)
	*/
	method := rpcRequest.Method
	log.Println("xxx Excute() parsing Method:", method)

	f := rpcRequest.Params
	key := f.(map[string]interface{})["key"].(string)
	log.Println("rpcRequest.Params.Key:", key)
	channel := f.(map[string]interface{})["channel"].(string)
	log.Println("rpcRequest.Params.Channel:", channel)

	rpcResp, err := sendJsonrpcRequest(method, key)

	isok, err := verifyMsg(method, rpcResp)
	if isok {
		respMsg, err := json.Marshal(rpcResp)
		log.Println("echo the message:", string(respMsg))
		utils.CheckError(err)
		return respMsg
	}
	return nil
}
