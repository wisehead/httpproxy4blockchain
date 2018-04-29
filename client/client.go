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

//var addr = flag.String("addr", "localhost:8080", "http service address")
var addr = flag.String("addr", "127.0.0.1:8080", "http service address")

//101.201.36.66
//var addr = flag.String("addr", "101.201.36.66:8080", "http service address")

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

	/*
		//Message 1:source-state
		request := &RPCRequest{
			Method: "source-state",
			Params: map[string]interface{}{
				"key":     "mytest/6",
				"channel": "vvtrip",
			},
			ID:      0,
			JSONRPC: "2.0",
		}
	*/

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
				"key":     "mytest/5",
				"tx_id":   "5f61e09ca61b6ec3db86bbb21134c65d8419f3e32df334e802224b1dd1fcceaf",
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
	request := &RPCRequest{
		Method: "source-get-binary",
		Params: map[string]interface{}{
			"channel": "vvtrip",
			"key":     "test/my2.pic",
		},
		JSONRPC: "2.0",
		ID:      0,
	}

	//1.seed info
	qulityInfoX := make(map[string]interface{})
	qulityInfoX["unpolishedRiceRate"] = "84.1%"
	qulityInfoX["polishedRiceRate"] = "75.7%"
	qulityInfoX["headRiceRate"] = "66.8%"
	qulityInfoX["grainLength"] = "6.6mm"
	qulityInfoX["gelConsistency"] = "67.0%"
	qulityInfoX["tasteScore"] = "88-92"

	seedInfoX := make(map[string]interface{})
	seedInfoX["company"] = "黑龙江省五常市龙洋种子有限公司"
	seedInfoX["registeredNumber"] = "912301847345944427"
	seedInfoX["unifiedSocialCreditCode"] = "912301847345944427"
	seedInfoX["productInfo"] = "五优稻4号（稻花香2号）"
	seedInfoX["seedValidationNumber"] = "黑审稻2009005"
	seedInfoX["qulity"] = qulityInfoX

	//2.BiologicalOrganicFertilizerInfo
	chemicalCompositionInfoX := make(map[string]interface{})
	chemicalCompositionInfoX["enzyme"] = ">=8%"
	chemicalCompositionInfoX["organicMatter"] = ">=45%"
	chemicalCompositionInfoX["aminoAcid"] = ">=12%"
	chemicalCompositionInfoX["humicAcid"] = ">=10%"

	biologicalOrganicFertilizerInfoX := make(map[string]interface{})
	biologicalOrganicFertilizerInfoX["company"] = "哈尔滨恒丰源农业科技发展有限责任公司"
	biologicalOrganicFertilizerInfoX["unifiedSocialCreditCode"] = "912301030780694118"
	biologicalOrganicFertilizerInfoX["organizationCode"] = "078069411"
	biologicalOrganicFertilizerInfoX["registeredNumber"] = "230103100339348"
	biologicalOrganicFertilizerInfoX["productInfo"] = "恒丰源生物有机肥"
	biologicalOrganicFertilizerInfoX["chemicalComposition"] = chemicalCompositionInfoX

	/*
		3.有机认证信息
		基地有机证编号：227OP1600100
		加工有机证编号：227OP1600099
	*/
	organicAuthenticationInfoX := make(map[string]interface{})
	organicAuthenticationInfoX["organicEvidenceInBaseNum"] = "227OP1600100"
	organicAuthenticationInfoX["processingOrganicSyndromeNum"] = "227OP1600099"

	/*4.产品信息
	  大米检验报告单
	  序号	检验项目	计量单位	标准要求	实测值	单项结论
	  1	感官要求		应符合标准要求
	  2	加工精度		应符合标准要求
	  3	黄粒米	%	≤0.1
	  4	不完善粒	%	≤0.5
	  5	杂质总量	%	≤0.10
	  	糠    粉	%	≤0.02
	  	矿 物 质	%	0
	  6	碎米总量	%	≤10.0
	  	小 碎 米	%	≤0.2
	  7	水   份	%	≤15.5
	  8	垩白粒率	%	≤15
	  9	食味品质	分	>=85
	  10	直链淀粉	%	15-20
	  11	胶 稠 度	mm	>=70
	  12	霉 变 粒	%	≤2.0
	*/
	var productInfomationX []map[string]interface{}

	productItemX1 := make(map[string]interface{})
	productItemX1["id"] = "1"
	productItemX1["inspectionProject"] = "感官要求"
	productItemX1["measurementUnit"] = "%"
	productItemX1["standardRequirements"] = "应符合标准要求"
	productItemX1["measuredValue"] = "0"
	productItemX1["singleConclusion"] = "ok"
	productInfomationX = append(productInfomationX, productItemX1)

	productItemX2 := make(map[string]interface{})
	productItemX2["id"] = "2"
	productItemX2["inspectionProject"] = "加工精度"
	productItemX2["measurementUnit"] = "%"
	productItemX2["standardRequirements"] = "应符合标准要求"
	productItemX2["measuredValue"] = "0"
	productItemX2["singleConclusion"] = "ok"
	productInfomationX = append(productInfomationX, productItemX2)

	productItemX3 := make(map[string]interface{})
	productItemX3["id"] = "3"
	productItemX3["inspectionProject"] = "黄粒米"
	productItemX3["measurementUnit"] = "%"
	productItemX3["standardRequirements"] = "≤0.1"
	productItemX3["measuredValue"] = "0"
	productItemX3["singleConclusion"] = "ok"
	productInfomationX = append(productInfomationX, productItemX3)

	productItemX4 := make(map[string]interface{})
	productItemX4["id"] = "4"
	productItemX4["inspectionProject"] = "不完善粒"
	productItemX4["measurementUnit"] = "%"
	productItemX4["standardRequirements"] = "≤0.5"
	productItemX4["measuredValue"] = "0"
	productItemX4["singleConclusion"] = "ok"
	productInfomationX = append(productInfomationX, productItemX4)

	productItemX5 := make(map[string]interface{})
	productItemX5["id"] = "5"
	productItemX5["inspectionProject"] = "杂质总量"
	productItemX5["measurementUnit"] = "%"
	productItemX5["standardRequirements"] = "≤0.10"
	productItemX5["measuredValue"] = "0"
	productItemX5["singleConclusion"] = "ok"
	productInfomationX = append(productInfomationX, productItemX5)

	productItemX6 := make(map[string]interface{})
	productItemX6["id"] = "5"
	productItemX6["inspectionProject"] = "糠粉"
	productItemX6["measurementUnit"] = "%"
	productItemX6["standardRequirements"] = "≤0.02"
	productItemX6["measuredValue"] = "0"
	productItemX6["singleConclusion"] = "ok"
	productInfomationX = append(productInfomationX, productItemX6)

	productItemX7 := make(map[string]interface{})
	productItemX7["id"] = "5"
	productItemX7["inspectionProject"] = "矿物质"
	productItemX7["measurementUnit"] = "%"
	productItemX7["standardRequirements"] = "0"
	productItemX7["measuredValue"] = "0"
	productItemX7["singleConclusion"] = "ok"
	productInfomationX = append(productInfomationX, productItemX7)

	productItemX8 := make(map[string]interface{})
	productItemX8["id"] = "6"
	productItemX8["inspectionProject"] = "碎米总量"
	productItemX8["measurementUnit"] = "%"
	productItemX8["standardRequirements"] = "≤10.0"
	productItemX8["measuredValue"] = "0"
	productItemX8["singleConclusion"] = "ok"
	productInfomationX = append(productInfomationX, productItemX8)

	productItemX9 := make(map[string]interface{})
	productItemX9["id"] = "6"
	productItemX9["inspectionProject"] = "小碎米"
	productItemX9["measurementUnit"] = "%"
	productItemX9["standardRequirements"] = "≤0.2"
	productItemX9["measuredValue"] = "0"
	productItemX9["singleConclusion"] = "ok"
	productInfomationX = append(productInfomationX, productItemX9)

	productItemX10 := make(map[string]interface{})
	productItemX10["id"] = "7"
	productItemX10["inspectionProject"] = "水份"
	productItemX10["measurementUnit"] = "%"
	productItemX10["standardRequirements"] = "≤15.5"
	productItemX10["measuredValue"] = "0"
	productItemX10["singleConclusion"] = "ok"
	productInfomationX = append(productInfomationX, productItemX10)

	productItemX11 := make(map[string]interface{})
	productItemX11["id"] = "8"
	productItemX11["inspectionProject"] = "垩白粒率"
	productItemX11["measurementUnit"] = "%"
	productItemX11["standardRequirements"] = "≤15"
	productItemX11["measuredValue"] = "0"
	productItemX11["singleConclusion"] = "ok"
	productInfomationX = append(productInfomationX, productItemX11)

	productItemX12 := make(map[string]interface{})
	productItemX12["id"] = "9"
	productItemX12["inspectionProject"] = "食味品质"
	productItemX12["measurementUnit"] = "分"
	productItemX12["standardRequirements"] = ">=85"
	productItemX12["measuredValue"] = "100"
	productItemX12["singleConclusion"] = "ok"
	productInfomationX = append(productInfomationX, productItemX12)

	productItemX13 := make(map[string]interface{})
	productItemX13["id"] = "10"
	productItemX13["inspectionProject"] = "直链淀粉"
	productItemX13["measurementUnit"] = "%"
	productItemX13["standardRequirements"] = "15-20"
	productItemX13["measuredValue"] = "16"
	productItemX13["singleConclusion"] = "ok"
	productInfomationX = append(productInfomationX, productItemX13)

	productItemX14 := make(map[string]interface{})
	productItemX14["id"] = "11"
	productItemX14["inspectionProject"] = "胶稠度"
	productItemX14["measurementUnit"] = "mm"
	productItemX14["standardRequirements"] = ">=70"
	productItemX14["measuredValue"] = "80"
	productItemX14["singleConclusion"] = "ok"
	productInfomationX = append(productInfomationX, productItemX14)

	productItemX15 := make(map[string]interface{})
	productItemX15["id"] = "12"
	productItemX15["inspectionProject"] = "霉变粒"
	productItemX15["measurementUnit"] = "%"
	productItemX15["standardRequirements"] = "≤2.0"
	productItemX15["measuredValue"] = "0"
	productItemX15["singleConclusion"] = "ok"
	productInfomationX = append(productInfomationX, productItemX15)

	/*5.土壤检测报告（数据辛苦从图片中提取即可）*/

	var soilCheckReportX []map[string]interface{}

	soilCheckReportItemX1 := make(map[string]interface{})
	soilCheckReportItemX1["smapleId"] = "1"
	soilCheckReportItemX1["sampleInspectionProject"] = "PH"
	soilCheckReportItemX1["soilUnit"] = "-"
	soilCheckReportItemX1["soilIndex"] = "-"
	soilCheckReportItemX1["soilMeasuredData"] = "5.1"
	soilCheckReportItemX1["sampleNumDetectionLimit"] = "-"
	soilCheckReportItemX1["ln166872SingleConclusion"] = "-"
	soilCheckReportItemX1["ln166872DetectionBasis"] = "NY/T 1377-2007"
	soilCheckReportX = append(soilCheckReportX, soilCheckReportItemX1)

	soilCheckReportItemX2 := make(map[string]interface{})
	soilCheckReportItemX2["smapleId"] = "2"
	soilCheckReportItemX2["sampleInspectionProject"] = "阳离子交换量"
	soilCheckReportItemX2["soilUnit"] = "cmod(+)/kg"
	soilCheckReportItemX2["soilIndex"] = "-"
	soilCheckReportItemX2["soilMeasuredData"] = "54.6"
	soilCheckReportItemX2["sampleNumDetectionLimit"] = "-"
	soilCheckReportItemX2["ln166872SingleConclusion"] = "-"
	soilCheckReportItemX2["ln166872DetectionBasis"] = "NY/T 1243-1999"
	soilCheckReportX = append(soilCheckReportX, soilCheckReportItemX2)

	soilCheckReportItemX3 := make(map[string]interface{})
	soilCheckReportItemX3["smapleId"] = "3"
	soilCheckReportItemX3["sampleInspectionProject"] = "铅"
	soilCheckReportItemX3["soilUnit"] = "mg/kg"
	soilCheckReportItemX3["soilIndex"] = "<=250"
	soilCheckReportItemX3["soilMeasuredData"] = "15.9"
	soilCheckReportItemX3["sampleNumDetectionLimit"] = "-"
	soilCheckReportItemX3["ln166872SingleConclusion"] = "合格"
	soilCheckReportItemX3["ln166872DetectionBasis"] = "GB/T 17141-1997"
	soilCheckReportX = append(soilCheckReportX, soilCheckReportItemX3)

	soilCheckReportItemX4 := make(map[string]interface{})
	soilCheckReportItemX4["smapleId"] = "4"
	soilCheckReportItemX4["sampleInspectionProject"] = "砷"
	soilCheckReportItemX4["soilUnit"] = "mg/kg"
	soilCheckReportItemX4["soilIndex"] = "<=40"
	soilCheckReportItemX4["soilMeasuredData"] = "6.34"
	soilCheckReportItemX4["sampleNumDetectionLimit"] = "-"
	soilCheckReportItemX4["ln166872SingleConclusion"] = "合格"
	soilCheckReportItemX4["ln166872DetectionBasis"] = "GB/T 22105.2-2008"
	soilCheckReportX = append(soilCheckReportX, soilCheckReportItemX4)

	soilCheckReportItemX5 := make(map[string]interface{})
	soilCheckReportItemX5["smapleId"] = "5"
	soilCheckReportItemX5["sampleInspectionProject"] = "汞"
	soilCheckReportItemX5["soilUnit"] = "mg/kg"
	soilCheckReportItemX5["soilIndex"] = "<=0.30"
	soilCheckReportItemX5["soilMeasuredData"] = "0.069"
	soilCheckReportItemX5["sampleNumDetectionLimit"] = "-"
	soilCheckReportItemX5["ln166872SingleConclusion"] = "合格"
	soilCheckReportItemX5["ln166872DetectionBasis"] = "GB/T 22105.1-2008"
	soilCheckReportX = append(soilCheckReportX, soilCheckReportItemX5)

	soilCheckReportItemX6 := make(map[string]interface{})
	soilCheckReportItemX6["smapleId"] = "6"
	soilCheckReportItemX6["sampleInspectionProject"] = "铬"
	soilCheckReportItemX6["soilUnit"] = "mg/kg"
	soilCheckReportItemX6["soilIndex"] = "<=150"
	soilCheckReportItemX6["soilMeasuredData"] = "55"
	soilCheckReportItemX6["sampleNumDetectionLimit"] = "-"
	soilCheckReportItemX6["ln166872SingleConclusion"] = "合格"
	soilCheckReportItemX6["ln166872DetectionBasis"] = "HJ 491-2009"
	soilCheckReportX = append(soilCheckReportX, soilCheckReportItemX6)

	soilCheckReportItemX7 := make(map[string]interface{})
	soilCheckReportItemX7["smapleId"] = "7"
	soilCheckReportItemX7["sampleInspectionProject"] = "镉"
	soilCheckReportItemX7["soilUnit"] = "mg/kg"
	soilCheckReportItemX7["soilIndex"] = "<-0.30"
	soilCheckReportItemX7["soilMeasuredData"] = "0.21"
	soilCheckReportItemX7["sampleNumDetectionLimit"] = "-"
	soilCheckReportItemX7["ln166872SingleConclusion"] = "合格"
	soilCheckReportItemX7["ln166872DetectionBasis"] = "GB/T 17141-1997"
	soilCheckReportX = append(soilCheckReportX, soilCheckReportItemX7)

	soilCheckReportItemX8 := make(map[string]interface{})
	soilCheckReportItemX8["smapleId"] = "8"
	soilCheckReportItemX8["sampleInspectionProject"] = "六六六"
	soilCheckReportItemX8["soilUnit"] = "mg/kg"
	soilCheckReportItemX8["soilIndex"] = "<=0.50"
	soilCheckReportItemX8["soilMeasuredData"] = "未检出"
	soilCheckReportItemX8["sampleNumDetectionLimit"] = "0.8*10^(-4)"
	soilCheckReportItemX8["ln166872SingleConclusion"] = "合格"
	soilCheckReportItemX8["ln166872DetectionBasis"] = "GB/T 14550-2003"
	soilCheckReportX = append(soilCheckReportX, soilCheckReportItemX8)

	soilCheckReportItemX9 := make(map[string]interface{})
	soilCheckReportItemX9["smapleId"] = "9"
	soilCheckReportItemX9["sampleInspectionProject"] = "滴滴涕"
	soilCheckReportItemX9["soilUnit"] = "mg/kg"
	soilCheckReportItemX9["soilIndex"] = "<=0.50"
	soilCheckReportItemX9["soilMeasuredData"] = "未检出"
	soilCheckReportItemX9["sampleNumDetectionLimit"] = "4.87*10^(-3)"
	soilCheckReportItemX9["ln166872SingleConclusion"] = "合格"
	soilCheckReportItemX9["ln166872DetectionBasis"] = "GB/T 14550-2003"
	soilCheckReportX = append(soilCheckReportX, soilCheckReportItemX9)

	oneRecordX := make(map[string]interface{})
	oneRecordX["seedInfoData"] = seedInfoX
	oneRecordX["biologicalOrganicFertilizerInfoData"] = biologicalOrganicFertilizerInfoX
	oneRecordX["organicAuthenticationInfoData"] = organicAuthenticationInfoX
	oneRecordX["productInfomationData"] = productInfomationX
	oneRecordX["soilCheckReportData"] = soilCheckReportX

	//testRecord, _ := json.Marshal(oneRecordX)
	//log.Println("oneRecord is:", string(testRecord))
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
