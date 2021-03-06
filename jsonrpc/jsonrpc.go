// Package jsonrpc provides a JSON-RPC 2.0 client that sends JSON-RPC requests and receives JSON-RPC responses using HTTP.
package jsonrpc

import (
	"bytes"
	"crypto/md5"
	"encoding/base64"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"httpproxy4blockchain/logger"
	"io/ioutil"
	"math/rand"
	"net/http"
	"reflect"
	"strconv"
	"time"
)

const (
	jsonrpcVersion = "2.0"
	defaultID      = 0
)

// RPCClient sends JSON-RPC requests over HTTP to the provided JSON-RPC backend.
//
// RPCClient is created using the factory function NewClient().
type RPCClient interface {
	// Call is a very handy function to send a JSON-RPC request to the server endpoint.
	//
	// params can only be an array or an object, no primitive values.
	// So there are a few simple rules to notice:
	//
	// 1. no params: params field is omitted. e.g. Call("getinfo")
	//
	// 2. single params primitive value: value is wrapped in array. e.g. Call("getByID", 1423)
	//
	// 3. single params value array or object: value is unchanged. e.g. Call("storePerson", &Person{Name: "Alex"})
	//
	// 4. multiple params values: always wrapped in array. e.g. Call("setDetails", "Alex, 35, "Germany", true)
	//
	// Examples:
	//   Call("getinfo") -> {"method": "getinfo"}
	//   Call("getPersonId", 123) -> {"method": "getPersonId", "params": [123]}
	//   Call("setName", "Alex") -> {"method": "setName", "params": ["Alex"]}
	//   Call("setMale", true) -> {"method": "setMale", "params": [true]}
	//   Call("setNumbers", []int{1, 2, 3}) -> {"method": "setNumbers", "params": [1, 2, 3]}
	//   Call("setNumbers", 1, 2, 3) -> {"method": "setNumbers", "params": [1, 2, 3]}
	//   Call("savePerson", &Person{Name: "Alex", Age: 35}) -> {"method": "savePerson", "params": {"name": "Alex", "age": 35}}
	//   Call("setPersonDetails", "Alex", 35, "Germany") -> {"method": "setPersonDetails", "params": ["Alex", 35, "Germany"}}
	//
	// for more information, see the examples or the unit tests
	Call(method string, params ...interface{}) (*RPCResponse, error)

	// CallFor is a very handy function to send a JSON-RPC request to the server endpoint
	// and directly specify an object to store the response.
	//
	// out: will store the unmarshaled object, if request was successful.
	// should always be provided by references. can be nil even on success.
	//
	// method and params: see Call() function
	//
	// if the request was not successful or the rpc response returns an error,
	// error holds the error object. if it was an JSON-RPC error it can be casted
	// to *RPCError.
	//
	CallFor(out interface{}, method string, params ...interface{}) error
}

type BCRequest struct {
	Alg   string `json:"alg"`
	Data  string `json:"data,omitempty"`
	Nonce string `json:"nonce"` //chenhui
	Sign  string `json:"sign"`
}

// RPCRequest represents a JSON-RPC request object.
//
// Method: string containing the method to be invoked
//
// Params: can be nil. if not must be an json array or object
//
// ID: may always set to 1 for single requests. Should be unique for every request in one batch request.
//
// JSONRPC: must always be set to "2.0" for JSON-RPC version 2.0
//
// See: http://www.jsonrpc.org/specification#request_object
type RPCRequest struct {
	Method  string      `json:"method"`
	Params  interface{} `json:"params,omitempty"`
	ID      uint        `json:"id"` //chenhui
	JSONRPC string      `json:"jsonrpc"`
}

// RPCResponse represents a JSON-RPC response object.
//
// Result: holds the result of the rpc call if no error occurred, nil otherwise. can be nil even on success.
//
// Error: holds an RPCError object if an error occurred. must be nil on success.
//
// ID: may always be 1 for single requests. should be unique for every request in one batch request.
//
// JSONRPC: must always be set to "2.0" for JSON-RPC version 2.0
//
// See: http://www.jsonrpc.org/specification#response_object
type RPCResponse struct {
	JSONRPC string      `json:"jsonrpc"`
	Result  interface{} `json:"result,omitempty"`
	//Result map[string]interface{} `json:"result,omitempty"`
	//Result *json.RawMessage `json:"result,omitempty"`
	Error *RPCError `json:"error,omitempty"`
	ID    uint      `json:"id"`
}

type BCResponse struct {
	Data  string `json:"data"`
	Nonce string `json:"nonce"`
	//Result map[string]interface{} `json:"result,omitempty"`
	//Result *json.RawMessage `json:"result,omitempty"`
	//Error *RPCError `json:"error,omitempty"`
	Timestamp string `json:"timestamp"`
	Sign      string `json:"sign"`
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

// Error function is provided to be used as error object.
func (e *RPCError) Error() string {
	return strconv.Itoa(e.Code) + ":" + e.Message
}

// HTTPError represents a error that occurred on HTTP level.
//
// An error of type HTTPError is returned when a HTTP error occurred (status code)
// and the body could not be parsed to a valid RPCResponse object that holds a RPCError.
//
// Otherwise a RPCResponse object is returned with a RPCError field that is not nil.
type HTTPError struct {
	Code int
	err  error
}

// Error function is provided to be used as error object.
func (e *HTTPError) Error() string {
	return e.err.Error()
}

type rpcClient struct {
	endpoint      string
	httpClient    *http.Client
	customHeaders map[string]string
}

// RPCClientOpts can be provided to NewClientWithOpts() to change configuration of RPCClient.
//
// HTTPClient: provide a custom http.Client (e.g. to set a proxy, or tls options)
//
// CustomHeaders: provide custom headers, e.g. to set BasicAuth
type RPCClientOpts struct {
	HTTPClient    *http.Client
	CustomHeaders map[string]string
}

// NewClient returns a new RPCClient instance with default configuration.
//
// endpoint: JSON-RPC service URL to which JSON-RPC requests are sent.
//--
func NewClient(endpoint string) RPCClient {
	return NewClientWithOpts(endpoint, nil)
}

// NewClientWithOpts returns a new RPCClient instance with custom configuration.
//
// endpoint: JSON-RPC service URL to which JSON-RPC requests are sent.
//
// opts: RPCClientOpts provide custom configuration
//--
func NewClientWithOpts(endpoint string, opts *RPCClientOpts) RPCClient {
	rpcClient := &rpcClient{
		endpoint:      endpoint,
		httpClient:    &http.Client{},
		customHeaders: make(map[string]string),
	}

	if opts == nil {
		return rpcClient
	}

	if opts.HTTPClient != nil {
		rpcClient.httpClient = opts.HTTPClient
	}

	if opts.CustomHeaders != nil {
		for k, v := range opts.CustomHeaders {
			rpcClient.customHeaders[k] = v
		}
	}

	return rpcClient
}

//B_chenhui
const letterBytes = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"
const (
	letterIdxBits = 6                    // 6 bits to represent a letter index
	letterIdxMask = 1<<letterIdxBits - 1 // All 1-bits, as many as letterIdxBits
	letterIdxMax  = 63 / letterIdxBits   // # of letter indices fitting in 63 bits
)

var src = rand.NewSource(time.Now().UnixNano())

func RandStringBytesMaskImprSrc(n int) string {
	b := make([]byte, n)
	// A src.Int63() generates 63 random bits, enough for letterIdxMax characters!
	for i, cache, remain := n-1, src.Int63(), letterIdxMax; i >= 0; {
		if remain == 0 {
			cache, remain = src.Int63(), letterIdxMax
		}
		if idx := int(cache & letterIdxMask); idx < len(letterBytes) {
			b[i] = letterBytes[idx]
			i--
		}
		cache >>= letterIdxBits
		remain--
	}

	return string(b)
}

//E_chenhui
//--
func (client *rpcClient) Call(method string, params ...interface{}) (*RPCResponse, error) {
	/*
		md5{"params": {"channel": "vvtrip", "key": "mytest/1"}, "jsonrpc": "2.0", "id": 0, "method": "source-state"}
	*/
	rpcRequest := &RPCRequest{
		Params:  transformParams(params...),
		JSONRPC: "2.0",
		ID:      0,
		Method:  method,
	}

	body, err := json.Marshal(rpcRequest)
	if err != nil {
		return nil, err
	}

	logger.Info("Call() :rpcRequest is: ", string(body))

	str := strconv.Quote(string(body))
	logger.Info("Call() str:" + str)

	nonce := RandStringBytesMaskImprSrc(32)

	//data_orig := "md5" + string(body) + "R8n9eO3SVDTYbQrkZMw75vLisxBdNo6l" + "3072c26dedb17d5545e53099fced54d30e13ad7f98a0ca542a73549535540659"
	//data_orig := "md5" + string(body) + "R8n9eO3SVDTYbQrkZMw75vLisxBdNo6l" + "3072c26dedb17d5545e53099fced54d30e13ad7f98a0ca542a73549535540600"
	data_orig := "md5" + string(body) + nonce + "3072c26dedb17d5545e53099fced54d30e13ad7f98a0ca542a73549535540600"
	logger.Info("Call() :data_orig is: ", data_orig)

	h := md5.New()
	h.Write([]byte(data_orig)) // 需要加密的字符串为 123456
	cipherStr := h.Sum(nil)
	logger.Info("Call() sign is:", hex.EncodeToString(cipherStr)) // 输出加密结果

	bcRequest := &BCRequest{
		Data:  string(body),
		Nonce: "R8n9eO3SVDTYbQrkZMw75vLisxBdNo6l",
		Alg:   "md5",
		Sign:  hex.EncodeToString(cipherStr),
	}
	bcRequest.Nonce = nonce //chenhui
	return client.doCall(bcRequest, method)
}

func (client *rpcClient) CallFor(out interface{}, method string, params ...interface{}) error {
	rpcResponse, err := client.Call(method, params...)
	if err != nil {
		logger.Error("CallFor() ", err)
		return err
	}
	return rpcResponse.GetObject(out)
}

func (client *rpcClient) newRequest(req interface{}) (*http.Request, error) {
	body, err := json.Marshal(req)
	if err != nil {
		logger.Error("newRequest() ", err)
		return nil, err
	}

	logger.Info("newRequest() rpcClient :body is: ", string(body))
	request, err := http.NewRequest("POST", client.endpoint, bytes.NewReader(body))
	if err != nil {
		logger.Error("newRequest() ", err)
		return nil, err
	}

	request.Header.Set("Content-Type", "application/json")
	//request.Header.Set("X-Api-Key", "7a79d668d61993119516d7c898aa072bb971467752e3e7bb2751cc474080dbdb")
	request.Header.Set("X-Api-Key", "7a79d668d61993119516d7c898aa072bb971467752e3e7bb2751cc474080db00")

	// set default headers first, so that even content type and accept can be overwritten
	for k, v := range client.customHeaders {
		request.Header.Set(k, v)
	}

	return request, nil
}

//--
func (client *rpcClient) doCall(bcRequest *BCRequest, method string) (*RPCResponse, error) {

	httpRequest, err := client.newRequest(bcRequest)
	if err != nil || httpRequest == nil {
		return nil, err
	}
	httpResponse, err := client.httpClient.Do(httpRequest)
	if err != nil {
		logger.Error("doCall() :httpResponse is error")
		//return nil, fmt.Errorf("rpc call on %v: %v", httpRequest.URL.String(), err.Error())
		return nil, err
	}
	logger.Info("doCall() :httpResponse is:", httpResponse.Body)
	result, _ := ioutil.ReadAll(httpResponse.Body)
	logger.Info("doCall() response is:", string(result))

	defer httpResponse.Body.Close()

	var rpcResponse *RPCResponse
	var rpcResp = new(RPCResponse)
	var bcResponse *BCResponse
	bcResponse = new(BCResponse)

	err = json.Unmarshal(result, &bcResponse)
	logger.Info("doCall() bcResponse:", bcResponse)
	// it is possible that this result is a binary object.
	if err != nil {
		if method == "source-get-binary" {
			binary_obj := []byte(result)
			logger.Info("doCall() binary_obj is:", string(binary_obj))
			result := make(map[string]interface{})
			result["pic"] = base64.StdEncoding.EncodeToString(binary_obj)
			response := make(map[string]interface{})
			response["id"] = 0
			response["jsonrpc"] = "2.0"
			response["result"] = result
			b, _ := json.Marshal(response)
			err = json.Unmarshal(b, &rpcResp)
			if err != nil {
				logger.Error("doCall():")
				return nil, err
			}
			return rpcResp, nil

		} else {
			logger.Error("doCall() :json.Unmarshal error.....")
			return nil, err
		}
	}

	data := bcResponse.Data
	logger.Info("doCall() bcResponse.data:", data)
	nonce := bcResponse.Nonce
	logger.Info("doCall() bcResponse.nonce:", nonce)
	timestamp := bcResponse.Timestamp
	logger.Info("doCall() bcResponse.timestamp:", timestamp)
	sign := bcResponse.Sign
	logger.Info("doCall() bcResponse.sign:", sign)

	mirrormsg, err := json.Marshal(bcResponse)
	logger.Info("doCall() mirrormsg:", string(mirrormsg))

	err = json.Unmarshal([]byte(data), &rpcResp)
	logger.Info("doCall() rpcResp:", rpcResp)
	logger.Info("doCall() method is:", method)
	if err != nil {
		if method == "source-get-binary" {
			binary_obj := []byte(data)
			logger.Info("doCall() binary_obj is:", string(binary_obj))
		} else {
			logger.Error("doCall :json.Unmarshal error.....") //chenhui
			return nil, err
		}
	}

	rpcResponse = rpcResp
	logger.Info("doCall() rpcResponse:", rpcResponse)

	// parsing error
	if err != nil && err.Error() != "EOF" {
		logger.Error("doCall() :httpResponse parsing error.....") //chenhui
		// if we have some http error, return it
		if httpResponse.StatusCode >= 400 {
			return nil, &HTTPError{
				Code: httpResponse.StatusCode,
				//err:  fmt.Errorf("rpc call %v() on %v status code: %v. could not decode body to rpc response: %v", RPCRequest.Method, httpRequest.URL.String(), httpResponse.StatusCode, err.Error()),
				err: fmt.Errorf("rpc call on %v: %v", httpRequest.URL.String(), err.Error()),
			}
		}
		//fmt.Printf("rpc call %v() on %v status code: %v. could not decode body to rpc response: %v", RPCRequest.Method, httpRequest.URL.String(), httpResponse.StatusCode, err.Error()) //chenhui
		logger.Info("rpc call on %v: %v", httpRequest.URL.String(), err.Error())
		//return nil, fmt.Errorf("rpc call %v() on %v status code: %v. could not decode body to rpc response: %v", RPCRequest.Method, httpRequest.URL.String(), httpResponse.StatusCode, err.Error())
		return nil, fmt.Errorf("rpc call on %v: %v", httpRequest.URL.String(), err.Error())
	}

	// response body empty
	if rpcResponse == nil {
		logger.Info("doCall() :rpcResponse is null .....") //chenhui
		// if we have some http error, return it
		if httpResponse.StatusCode >= 400 {
			return nil, &HTTPError{
				Code: httpResponse.StatusCode,
				//err:  fmt.Errorf("rpc call %v() on %v status code: %v. rpc response missing", RPCRequest.Method, httpRequest.URL.String(), httpResponse.StatusCode),
				err: fmt.Errorf("rpc call on %v: %v", httpRequest.URL.String(), err.Error()),
			}
		}
		//return nil, fmt.Errorf("rpc call %v() on %v status code: %v. rpc response missing", RPCRequest.Method, httpRequest.URL.String(), httpResponse.StatusCode)
		return nil, fmt.Errorf("rpc call on %v: %v", httpRequest.URL.String(), err.Error())
	}

	return rpcResponse, nil
}

func transformParams(params ...interface{}) interface{} {
	var finalParams interface{}

	// if params was nil skip this and p stays nil
	if params != nil {
		switch len(params) {
		case 0: // no parameters were provided, do nothing so finalParam is nil and will be omitted
		case 1: // one param was provided, use it directly as is, or wrap primitive types in array
			if params[0] != nil {
				var typeOf reflect.Type

				// traverse until nil or not a pointer type
				for typeOf = reflect.TypeOf(params[0]); typeOf != nil && typeOf.Kind() == reflect.Ptr; typeOf = typeOf.Elem() {
				}

				if typeOf != nil {
					// now check if we can directly marshal the type or if it must be wrapped in an array
					switch typeOf.Kind() {
					// for these types we just do nothing, since value of p is already unwrapped from the array params
					case reflect.Struct:
						finalParams = params[0]
					case reflect.Array:
						finalParams = params[0]
					case reflect.Slice:
						finalParams = params[0]
					case reflect.Interface:
						finalParams = params[0]
					case reflect.Map:
						finalParams = params[0]
					default: // everything else must stay in an array (int, string, etc)
						finalParams = params
					}
				}
			} else {
				finalParams = params
			}
		default: // if more than one parameter was provided it should be treated as an array
			finalParams = params
		}
	}

	return finalParams
}

// GetInt converts the rpc response to an int64 and returns it.
//
// If result was not an integer an error is returned.
/*
func (RPCResponse *RPCResponse) GetInt() (int64, error) {
	val, ok := RPCResponse.Result.(json.Number)
	if !ok {
		return 0, fmt.Errorf("could not parse int64 from %s", RPCResponse.Result)
	}

	i, err := val.Int64()
	if err != nil {
		return 0, err
	}

	return i, nil
}
*/

// GetFloat converts the rpc response to float64 and returns it.
//
// If result was not an float64 an error is returned.
func (RPCResponse *RPCResponse) GetFloat() (float64, error) {
	val, ok := RPCResponse.Result.(json.Number)
	if !ok {
		return 0, fmt.Errorf("could not parse float64 from %s", RPCResponse.Result)
	}

	f, err := val.Float64()
	if err != nil {
		return 0, err
	}

	return f, nil
}

// GetBool converts the rpc response to a bool and returns it.
//
// If result was not a bool an error is returned.
func (RPCResponse *RPCResponse) GetBool() (bool, error) {
	val, ok := RPCResponse.Result.(bool)
	if !ok {
		return false, fmt.Errorf("could not parse bool from %s", RPCResponse.Result)
	}

	return val, nil
}

// GetString converts the rpc response to a string and returns it.
//
// If result was not a string an error is returned.
func (RPCResponse *RPCResponse) GetString() (string, error) {
	val, ok := RPCResponse.Result.(string)
	if !ok {
		return "", fmt.Errorf("could not parse string from %s", RPCResponse.Result)
	}

	return val, nil
}

// GetObject converts the rpc response to an arbitrary type.
//
// The function works as you would expect it from json.Unmarshal()
func (RPCResponse *RPCResponse) GetObject(toType interface{}) error {
	js, err := json.Marshal(RPCResponse.Result)
	if err != nil {
		return err
	}

	err = json.Unmarshal(js, toType)
	if err != nil {
		return err
	}

	return nil
}
