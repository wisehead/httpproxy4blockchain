// Copyright 2015 The Gorilla WebSocket Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// +build ignore

package main

import (
	"encoding/json"
	"flag"
	"log"
	"net/url"
	"os"
	"os/signal"
	"time"

	"github.com/gorilla/websocket"
)

// RPCRequest represents a JSON-RPC request object.
type RPCRequest struct {
	Method  string      `json:"method"`
	Params  interface{} `json:"params,omitempty"`
	ID      uint        `json:"id"` //chenhui
	JSONRPC string      `json:"jsonrpc"`
}

var addr = flag.String("addr", "localhost:8080", "http service address")

func main() {
	flag.Parse()
	log.SetFlags(0)

	interrupt := make(chan os.Signal, 1)
	signal.Notify(interrupt, os.Interrupt)

	u := url.URL{Scheme: "ws", Host: *addr, Path: "/echo"}
	log.Printf("connecting to %s", u.String())

	c, _, err := websocket.DefaultDialer.Dial(u.String(), nil)
	if err != nil {
		log.Fatal("dial:", err)
	}
	defer c.Close()

	done := make(chan struct{})

	go func() {
		defer close(done)
		for {
			_, message, err := c.ReadMessage()
			if err != nil {
				log.Println("read:", err)
				return
			}
			log.Printf("recv: %s", message)
		}
	}()

	//Message 1:source-state
	request := &RPCRequest{
		Method: "source-state",
		Params: map[string]interface{}{
			"key":     "mytest/4",
			"channel": "vvtrip",
		},
		ID:      0,
		JSONRPC: "2.0",
	}

	/*
		//Message 2:source-transactions
		request := &RPCRequest{
			Method: "source-transactions",
			Params: map[string]interface{}{
				"key":     "mytest/1",
				"channel": "vvtrip",
			},
			ID:      0,
			JSONRPC: "2.0",
		}
	*/
	/*
		//Message 3:source-transaction
			request := &RPCRequest{
				Method: "source-transaction",
				Params: map[string]interface{}{
					"key":     "mytest/1",
					"tx_id":   "b11a94dd1142559380d1a715da39b6899ed55511f7e23164a50159e4dad4f936",
					"channel": "vvtrip",
				},
				ID:      0,
				JSONRPC: "2.0",
			}
	*/
	/*type RPCRequest struct {
		Method  string      `json:"method"`
		Params  interface{} `json:"params,omitempty"`
		ID      uint        `json:"id"` //chenhui
		JSONRPC string      `json:"jsonrpc"`
	}*/

	//message 4: source-get-binary
	/*
		request := &RPCRequest{
			Method: "source-get-binary",
			Params: map[string]interface{}{
				"channel": "vvtrip",
				"key":     "test/my1.txt",
			},
			JSONRPC: "2.0",
			ID:      0,
		}
	*/
	message, _ := json.Marshal(request)
	//err = c.WriteMessage(websocket.TextMessage, []byte(message.String()))
	err = c.WriteMessage(websocket.TextMessage, []byte(message))
	if err != nil {
		log.Println("write:", err)
		return
	}

	//ticker := time.NewTicker(time.Second)
	//defer ticker.Stop()

	for {
		select {
		case <-done:
			return
			/*
				case t := <-ticker.C:
					err := c.WriteMessage(websocket.TextMessage, []byte(t.String()))
					if err != nil {
						log.Println("write:", err)
						return
					}
			*/
		case <-interrupt:
			log.Println("interrupt")

			// Cleanly close the connection by sending a close message and then
			// waiting (with timeout) for the server to close the connection.
			err := c.WriteMessage(websocket.CloseMessage, websocket.FormatCloseMessage(websocket.CloseNormalClosure, ""))
			if err != nil {
				log.Println("write close:", err)
				return
			}
			select {
			case <-done:
			case <-time.After(time.Second):
			}
			return
		}
	}

}
