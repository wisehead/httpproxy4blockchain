// +build ignore
package main

// Copyright 2015 The Gorilla WebSocket Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

import (
	"encoding/json"
	"flag"
	"html/template"
	"httpproxy4blockchain/handler"
	"httpproxy4blockchain/jsonrpc"
	"httpproxy4blockchain/logger"
	"log"
	"net/http"
	"os"

	"github.com/gorilla/websocket"
)

//Version v0.85
//15 pics ok.

//var addr = flag.String("addr", "localhost:8080", "http service address")
var addr = flag.String("addr", "127.0.0.1:8088", "http service address")

var upgrader = websocket.Upgrader{} // use default options

//handleMsg is to handle the message from app client
func handleMsg(c *websocket.Conn, messageType int, postdata []byte) error {
	var rpcRequest jsonrpc.RPCRequest
	err := json.Unmarshal(postdata, &rpcRequest)
	if err != nil {
		logger.Error("Unmarshal:", err)
		return err
	}

	jsonrpc := rpcRequest.JSONRPC
	logger.Info("xxx rpcRequest.jsonrpc:", jsonrpc)
	method := rpcRequest.Method
	logger.Info("xxx rpcRequest.Method:", method)

	f := rpcRequest.Params
	key := f.(map[string]interface{})["key"].(string)
	logger.Info("rpcRequest.Params.Key:", key)
	channel := f.(map[string]interface{})["channel"].(string)
	logger.Info("rpcRequest.Params.Channel:", channel)
	rpcResp, err := handler.Excute(postdata)

	//chenhui
	if err != nil || rpcResp == nil {
		errorX := make(map[string]interface{})
		errorX["code"] = 0
		errorX["message"] = err.Error()

		response := make(map[string]interface{})
		response["id"] = 0
		response["jsonrpc"] = "2.0"
		response["error"] = errorX
		rpcResp, err = json.Marshal(response)
		if err != nil {
			logger.Error("handleMsg() Marshal :", err)
			return err
		}
	}

	err = c.WriteMessage(messageType, rpcResp)
	if err != nil {
		logger.Error("handleMsg() write:", err)
		return err
	}
	return nil
}

func echo(w http.ResponseWriter, r *http.Request) {
	c, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		logger.Error("upgrade:", err)
		return
	}
	defer c.Close()
	for {
		mt, message, err := c.ReadMessage()
		if err != nil {
			logger.Error("read:", err)
			break
		}
		logger.Info("recv: %s", message)
		//err = c.WriteMessage(mt, message)
		err = handleMsg(c, mt, message)
		if err != nil {
			logger.Error("handle:", err)
			break
		}
	}
}

func home(w http.ResponseWriter, r *http.Request) {
	homeTemplate.Execute(w, "ws://"+r.Host+"/echo")
}

const logPath = "./logpath/proxy.log"

var verbose = flag.Bool("verbose", false, "print info level logs to stdout")

func main() {
	flag.Parse()

	lf, err := os.OpenFile(logPath, os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0660)
	if err != nil {
		logger.Fatalf("Failed to open log file: %v", err)
	}
	defer lf.Close()

	defer logger.Init("LoggerProxy", *verbose, true, lf).Close()

	logger.Info("starting Blockchain proxy...!logPath is: ", logPath)

	log.SetFlags(0)
	http.HandleFunc("/echo", echo)
	http.HandleFunc("/", home)
	log.Fatal(http.ListenAndServe(*addr, nil))
}

var homeTemplate = template.Must(template.New("").Parse(`
<!DOCTYPE html>
<html>
<head>
<meta charset="utf-8">
<script>  
window.addEventListener("load", function(evt) {

    var output = document.getElementById("output");
    var input = document.getElementById("input");
    var ws;

    var print = function(message) {
        var d = document.createElement("div");
        d.innerHTML = message;
        output.appendChild(d);
    };

    document.getElementById("open").onclick = function(evt) {
        if (ws) {
            return false;
        }
        ws = new WebSocket("{{.}}");
        ws.onopen = function(evt) {
            print("OPEN");
        }
        ws.onclose = function(evt) {
            print("CLOSE");
            ws = null;
        }
        ws.onmessage = function(evt) {
            print("RESPONSE: " + evt.data);
        }
        ws.onerror = function(evt) {
            print("ERROR: " + evt.data);
        }
        return false;
    };

    document.getElementById("send").onclick = function(evt) {
        if (!ws) {
            return false;
        }
        print("SEND: " + input.value);
        ws.send(input.value);
        return false;
    };

    document.getElementById("close").onclick = function(evt) {
        if (!ws) {
            return false;
        }
        ws.close();
        return false;
    };

});
</script>
</head>
<body>
<table>
<tr><td valign="top" width="50%">
<p>Click "Open" to create a connection to the server, 
"Send" to send a message to the server and "Close" to close the connection. 
You can change the message and send multiple times.
<p>
<form>
<button id="open">Open</button>
<button id="close">Close</button>
<p><input id="input" type="text" value="Hello world!">
<button id="send">Send</button>
</form>
</td><td valign="top" width="50%">
<div id="output"></div>
</td></tr></table>
</body>
</html>
`))
