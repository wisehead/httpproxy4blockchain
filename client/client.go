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

//Version v0.97
//asset-exchange: exchange message ok.. start testing asset test chain.

// RPCRequest represents a JSON-RPC request object.
type RPCRequest struct {
	Method  string      `json:"method"`
	Params  interface{} `json:"params,omitempty"`
	ID      uint        `json:"id"` //chenhui
	JSONRPC string      `json:"jsonrpc"`
}

//var addr = flag.String("addr", "localhost:8080", "http service address")
var addr = flag.String("addr", "127.0.0.1:8088", "http service address")

//101.201.36.66
//var addr = flag.String("addr", "101.201.36.66:8380", "http service address")

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

	//Message 11:asset-mine, asset test chain.
	/*
		var arrsX []map[string]interface{}
		arrItem1X := make(map[string]interface{})
		arrItem1X["user"] = "userC"
		arrItem1X["amount"] = 1000000000000
		arrsX = append(arrsX, arrItem1X)

		arrItem2X := make(map[string]interface{})
		arrItem2X["user"] = "userD"
		arrItem2X["amount"] = 2000000000000
		arrsX = append(arrsX, arrItem2X)

		params := make(map[string]interface{})
		params["channel"] = "assettestchannel"
		params["uid"] = "20180625155401001"
		params["arr"] = arrsX

		request := make(map[string]interface{})
		request["method"] = "asset-mine"
		request["id"] = 0
		request["jsonrpc"] = "2.0"
		request["params"] = params
	*/

	//message 12: asset-transfer: recycle
	/*
		var arrsX []map[string]interface{}
		arrItem1X := make(map[string]interface{})
		arrItem1X["user"] = "userC"
		arrItem1X["amount"] = 100000000
		arrsX = append(arrsX, arrItem1X)

		params := make(map[string]interface{})
		params["channel"] = "assettestchannel"
		params["uid"] = "20180625174801001"
		params["type"] = "recycle"
		params["arr"] = arrsX

		request := make(map[string]interface{})
		request["method"] = "asset-transfer"
		request["id"] = 0
		request["jsonrpc"] = "2.0"
		request["params"] = params
	*/

	/*
		//message 13: asset-transfer: buy
		var arrsX []map[string]interface{}
		arrItem1X := make(map[string]interface{})
		arrItem1X["user"] = "userC"
		arrItem1X["amount"] = 5000000000
		arrsX = append(arrsX, arrItem1X)

		params := make(map[string]interface{})
		params["channel"] = "assettestchannel"
		params["uid"] = "20180625180001001"
		params["type"] = "buy"
		params["arr"] = arrsX

		request := make(map[string]interface{})
		request["method"] = "asset-transfer"
		request["id"] = 0
		request["jsonrpc"] = "2.0"
		request["params"] = params
	*/

	/*
		//message 14: asset-transfer: inner
		var arrsX []map[string]interface{}
		arrItem1X := make(map[string]interface{})
		arrItem1X["to"] = "userC"
		arrItem1X["from"] = "userD"
		arrItem1X["amount"] = 50000000000
		arrsX = append(arrsX, arrItem1X)

		params := make(map[string]interface{})
		params["channel"] = "assettestchannel"
		params["uid"] = "20180625181401001"
		params["type"] = "inner"
		params["arr"] = arrsX

		request := make(map[string]interface{})
		request["method"] = "asset-transfer"
		request["id"] = 0
		request["jsonrpc"] = "2.0"
		request["params"] = params
	*/

	/*
		//message 15: asset-exchange: xsb2mgb
		var arrsX []map[string]interface{}
		arrItem1X := make(map[string]interface{})
		arrItem1X["user"] = "userC"
		arrItem1X["amount"] = 900000000
		arrsX = append(arrsX, arrItem1X)

		params := make(map[string]interface{})
		params["channel"] = "assettestchannel"
		params["uid"] = "20180625182101001"
		params["type"] = "xsb2mgb"
		params["arr"] = arrsX

		request := make(map[string]interface{})
		request["method"] = "asset-exchange"
		request["id"] = 0
		request["jsonrpc"] = "2.0"
		request["params"] = params
	*/
	/*
		//message 16: asset-exchange: xsb2mgb
		var arrsX []map[string]interface{}
		arrItem1X := make(map[string]interface{})
		arrItem1X["user"] = "userC"
		arrItem1X["amount"] = 1100000000
		arrsX = append(arrsX, arrItem1X)

		params := make(map[string]interface{})
		params["channel"] = "assettestchannel"
		params["uid"] = "20180625182601001"
		params["type"] = "mgb2xsb"
		params["arr"] = arrsX

		request := make(map[string]interface{})
		request["method"] = "asset-exchange"
		request["id"] = 0
		request["jsonrpc"] = "2.0"
		request["params"] = params
	*/

	//message 17: asset-exchange: exchange
	params := make(map[string]interface{})
	params["channel"] = "assettestchannel"
	params["uid"] = "20180625191701001"
	params["type"] = "exchange"
	params["amount"] = 300000000

	request := make(map[string]interface{})
	request["method"] = "asset-exchange"
	request["id"] = 0
	request["jsonrpc"] = "2.0"
	request["params"] = params

	/*
		//Message 6:source-state, suyuan test chain, package info
		request := &RPCRequest{
			Method: "source-state",
			Params: map[string]interface{}{
				"key":     "1130102150229180616010000000005",
				"channel": "notaryinfotestchannel",
			},
			ID:      0,
			JSONRPC: "2.0",
		}
	*/

	/*
		//Message 5:source-state, suyuan test chain, batch info
		request := &RPCRequest{
			Method: "source-state",
			Params: map[string]interface{}{
				"key":     "113010215022918061601",
				"channel": "notaryinfotestchannel",
			},
			ID:      0,
			JSONRPC: "2.0",
		}
	*/
	/*
		//Message 1:source-state
		request := &RPCRequest{
			Method: "source-state",
			Params: map[string]interface{}{
				"key":     "mytest/1",
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

	/*
		//message 4: source-get-binary
		request := &RPCRequest{
			Method: "source-get-binary",
			Params: map[string]interface{}{
				"channel": "vvtrip",
				//"key":     "test/my2.pic",
				"key": "test/my1.txt",
			},
			JSONRPC: "2.0",
			ID:      0,
		}
	*/
	//1.seed info
	qulityX := make(map[string]interface{})
	qulityX["unpolishedRiceRate"] = "84.1%"
	qulityX["polishedRiceRate"] = "75.7%"
	qulityX["headRiceRate"] = "66.8%"
	qulityX["grainLength"] = "6.6mm"
	qulityX["gelConsistency"] = "67.0%"
	qulityX["tasteScore"] = "88-92"

	qualityExecutionStandardX := make(map[string]interface{})
	qualityExecutionStandardX["standardName"] = "GB4404.1-2008"
	qualityExecutionStandardX["unpolishedRiceRate"] = "99.9%"
	qualityExecutionStandardX["polishedRiceRate"] = "98.0%"
	qualityExecutionStandardX["headRiceRate"] = "85%"
	qualityExecutionStandardX["grainLength"] = "16.0%"

	/*
			type SeedInfoType struct {
			Grade                                     string                       `json:"grade"`                                     //1.	国家保护品种
			CropSpecies                               string                       `json:"cropSpecies"`                               //2.	作物种类：稻
			SeedCategory                              string                       `json:"seedCategory"`                              //3.	种子类别：常规种 原种
			VarietyName                               string                       `json:"varietyName"`                               //4.	品种名称：五优稻4号
			VarietyApprovalNumber                     string                       `json:"varietyApprovalNumber"`                     //5.	品种审定编号：黑审稻2009005 吉审稻2016011
			VarietyRightsNumber                       string                       `json:"varietyRightsNumber"`                       //6.	品种权号：CNA20080376.X
			LicenseNumberOfSeedProductionAndOperation string                       `json:"licenseNumberOfSeedProductionAndOperation"` //7.	种子生产经营许可证编号：C(黑哈五)农种许字(2016)第6001号
			Qulity                                    QulityType                   `json:"qulity"`                                    //8.	品质
			NumberOfPlantQuarantineCertificates       string                       `json:"numberOfPlantQuarantineCertificates"`       //9.	植物检疫证书编号：2301842017001091
			DateOfDetection                           string                       `json:"dateOfDetection"`                           //10.	检测日期：2017年11月 批号：0191
			QualityExecutionStandard                  QualityExecutionStandardType `json:"qualityExecutionStandard"`                  //11.	质量执行标准：GB4404.1-2008
			InformationCode                           string                       `json:"informationCode"`                           //12.	信息代码：0125727100017593
			FirstBreeders                             string                       `json:"firstBreeders"`                             //13.	五优稻4号(稻花香2号)第一育种人：田永太
			BreedsOwner                               string                       `json:""`                               //14.	五优稻4号(稻花香2号)品种权人：五常市利元种子有根公司
			NetContent                                string                       `json:""`                                //15.	净含量：25kg
			SeedProducer                              string                       `json:""`                              //16.	种子生产经营者：五常市利元种子有限公司
			RegisteredAddress                         string                       `json:""`                         //17.	注册地地址：黑龙江省五常市龙凤山镇乐园村汪家店屯
			Website                                   string                       `json:""`                                   //18.	网址:www.clxseed.com
		}
	*/
	seedInfoX := make(map[string]interface{})
	seedInfoX["grade"] = "国家保护品种"
	seedInfoX["cropSpecies"] = "稻"
	seedInfoX["seedCategory"] = "常规种 原种"
	seedInfoX["varietyName"] = "五优稻4号"
	seedInfoX["varietyApprovalNumber"] = "黑审稻2009005 吉审稻2016011"
	seedInfoX["varietyRightsNumber"] = "CNA20080376.X"
	seedInfoX["licenseNumberOfSeedProductionAndOperation"] = "C(黑哈五)农种许字(2016)第6001号"
	seedInfoX["qulity"] = qulityX
	seedInfoX["numberOfPlantQuarantineCertificates"] = "2301842017001091"
	seedInfoX["dateOfDetection"] = "2017年11月 批号：0191"
	seedInfoX["qualityExecutionStandard"] = qualityExecutionStandardX
	seedInfoX["informationCode"] = "0125727100017593"
	seedInfoX["firstBreeders"] = "田永太"
	seedInfoX["breedsOwner"] = "五常市利元种子有根公司"
	seedInfoX["netContent"] = "25kg"
	seedInfoX["seedProducer"] = "五常市利元种子有限公司"
	seedInfoX["registeredAddress"] = "黑龙江省五常市龙凤山镇乐园村汪家店屯"
	seedInfoX["website"] = "www.clxseed.com"

	/*
		二、生物有机肥
			type ProductMarkingType struct {
				MainTechnicalIndicators string `json:"mainTechnicalIndicators"` //1)主要技术指标：有效活菌数≥0.2亿/g  N+P2O5+K2O≥15% 有机质≥20%
				ScientificAddition      string `json:"scientificAddition"`      //2)科学添加:腐植酸≥12% 氨基酸≥1% 1酵素酶≥0.5亿/g中微量元素≥18%有益微生物蛋白酶及生长促进未知因子(UGF)
			}
			type BiologicalOrganicFertilizerType struct {
				Brand                                        string             `json:""`                                        //1.	品牌：归农稻花香有机米专用
				LicenseNumberOfGreenFoodProductionData       string             `json:""`       //2.	绿色食品生产资料许可编号：LSSZ-01-1503080005（中国绿色食品协会）
				ProductionTechnology                         string             `json:""`                         //3.	生产工艺技术：酵素生物技术、复合微生物肥
				RegistrationNumberOfTheMinistryOfAgriculture string             `json:""` //4.	农业部登记证号：微生物肥（2016）准字（1896号）
				StandardOfExecution                          string             `json:""`                          //5.	执行标准：NY/T798-2015
				ProductMarking                               ProductMarkingType `json:""`                               //6.	产品标值：
				PackageNumber                                string             `json:""`                                //7.	包装编号G16号
				Manufacturer                                 string             `json:""`                                 //8.	制造商：哈尔滨绿洲之星生物科技有限公司
				Address                                      string             `json:""`                                      //9.	地址:哈尔滨市香坊区印染街1号
				Website                                      string             `json:""`                                      //10.	网址：www-hrblz.com
			}
	*/
	productMarkingX := make(map[string]interface{})
	productMarkingX["mainTechnicalIndicators"] = "有效活菌数>=0.2亿/g  N+P2O5+K2O>=15% 有机质>=20%"
	productMarkingX["scientificAddition"] = "腐植酸>=12% 氨基酸>=1% 1酵素酶>=0.5亿/g中微量元素>=18% 有益微生物蛋白酶及生长促进未知因子(UGF)"

	biologicalOrganicFertilizerX := make(map[string]interface{})
	biologicalOrganicFertilizerX["brand"] = "归农稻花香有机米专用"
	biologicalOrganicFertilizerX["licenseNumberOfGreenFoodProductionData"] = "LSSZ-01-1503080005（中国绿色食品协会）"
	biologicalOrganicFertilizerX["productionTechnology"] = "酵素生物技术、复合微生物肥"
	biologicalOrganicFertilizerX["registrationNumberOfTheMinistryOfAgriculture"] = "微生物肥（2016）准字（1896号"
	biologicalOrganicFertilizerX["standardOfExecution"] = "NY/T798-2015"
	biologicalOrganicFertilizerX["productMarking"] = productMarkingX
	biologicalOrganicFertilizerX["packageNumber"] = "G16号"
	biologicalOrganicFertilizerX["manufacturer"] = "哈尔滨绿洲之星生物科技有限公司"
	biologicalOrganicFertilizerX["address"] = "哈尔滨市香坊区印染街1号"
	biologicalOrganicFertilizerX["website"] = "www-hrblz.com"

	/*
		三、基地有机认证
		type CertificationBasisType struct {
			Item1 string `json:"item1"` // GB/T19630.1-2011有机产品：生产
			Item2 string `json:"item2"` // GB/T19630.3-2011有机产品：标识与销售
			Item3 string `json:"item3"` // GB/T19630.4-2011有机产品：管理体系
		}

		type OrganicCertificationOfBaseType struct {
			CertificateNumber  string                 `json:"certificateNumber"`  //1.	证书编号：227OP1600099
			CertificationBasis CertificationBasisType `json:"certificationBasis"` //2.	认证依据
			PictureName        string                 `json:"pictureName"`        //picture name，全局唯一，需要从库里再取一次
		}
	*/
	certificationBasisX := make(map[string]interface{})
	certificationBasisX["item1"] = "GB/T19630.1-2011有机产品：生产"
	certificationBasisX["item2"] = "GB/T19630.3-2011有机产品：标识与销售"
	certificationBasisX["item3"] = "GB/T19630.4-2011有机产品：管理体系"

	organicCertificationOfBaseX := make(map[string]interface{})
	organicCertificationOfBaseX["certificateNumber"] = "227OP1600099"
	organicCertificationOfBaseX["certificationBasis"] = certificationBasisX
	organicCertificationOfBaseX["pictureName"] = "suyuan1/1.youji.png"

	/*四、检验报告*/
	/*
		type InspectionReportType struct {
			SampleName                        string            `json:"sampleName"`                        //样品名称
			ModelSpecification                string            `json:"modelSpecification"`                //型号规格
			SampleNumber                      string            `json:"sampleNumber"`                      //样品编号
			Trademark                         string            `json:"trademark"`                         //商标
			InspectionUnit                    string            `json:"inspectionUnit"`                    //送检单位
			InspectionCategory                string            `json:"inspectionCategory"`                //检验类别
			ProductionUnit                    string            `json:"productionUnit"`                    //生产单位
			SampleGradeAndState               string            `json:"sampleGradeAndState"`               //样品等级、状态
			SamplingSite                      string            `json:"samplingSite"`                      //抽样地点
			DateToSample                      string            `json:"dateToSample"`                      //到样日期
			SampleQuantity                    string            `json:"sampleQuantity"`                    //样品数量
			SampleMaker                       string            `json:"sampleMaker"`                       //送样者
			SamplingBase                      string            `json:"samplingBase"`                      //抽样基数
			OriginalNumberOrDateOfProduction  string            `json:"originalNumberOrDateOfProduction"`  //原编号或生产日期
			InspectionBasis                   string            `json:"inspectionBasis"`                   //检验依据
			InspectionProject                 string            `json:"inspectionProject"`                 //检验项目
			MainInstrumentsUsed               string            `json:"mainInstrumentsUsed"`               //所用主要仪器
			ExperimentalEnvironmentConditions string            `json:"experimentalEnvironmentConditions"` //实验环境条件
			InspectionConclusion              string            `json:"inspectionConclusion"`              //检验结论
			Remarks                           string            `json:"remarks"`                           //备注
			Approver                          string            `json:"approver"`                          //批准人
			DateOfApproval                    string            `json:"dateOfApproval"`                    //批准日期
			Auditor                           string            `json:"auditor"`                           //审核人
			DateOfAudit                       string            `json:"dateOfAudit"`                       //审核日期
			TabulatingPerson                  string            `json:"tabulatingPerson"`                  //制表人
			DateOfTabulation                  string            `json:"dateOfTabulation"`                  //制表日期
			Items                             []ProductItemType `json:"items"`                             //详细指标
			Picture1Cover                     string            `json:"picture1Cover"`                     //封面图片
			PictureBaseSoil                   string            `json:"pictureBaseSoil"`                   //基地土壤图片
			PictureIrrigatedWaterSource       string            `json:"pictureBaseSoil"`                   //灌溉水源图片
		}

		type ProductItemType struct {
			InspectionProject    string `json:"inspectionProject"`    //检验项目
			MeasurementUnit      string `json:"measurementUnit"`      //单位
			StandardRequirements string `json:"standardRequirements"` //标准要求
			TestResults          string `json:"testResults"`          //检验结果
			MethodDetectionLimit string `json:"methodDetectionLimit"` //方法检出限
			SingleConclusion     string `json:"singleConclusion"`     //单项判定
			TestMethod           string `json:"testMethod"`           //检验方法
		}
	*/
	var productInfomationX []map[string]interface{}

	productItemX1 := make(map[string]interface{})
	productItemX1["inspectionProject"] = "色泽、气味"
	productItemX1["measurementUnit"] = "——"
	productItemX1["standardRequirements"] = "无异常色泽和气味"
	productItemX1["testResults"] = "正常"
	productItemX1["methodDetectionLimit"] = "——"
	productItemX1["singleConclusion"] = "合格"
	productItemX1["testMethod"] = "GB/T 5492-2008"
	productInfomationX = append(productInfomationX, productItemX1)

	productItemX2 := make(map[string]interface{})
	productItemX2["inspectionProject"] = "不完善粒"
	productItemX2["measurementUnit"] = "%"
	productItemX2["standardRequirements"] = "<=3.0"
	productItemX2["testResults"] = "0.1"
	productItemX2["methodDetectionLimit"] = "——"
	productItemX2["singleConclusion"] = "合格"
	productItemX2["testMethod"] = "GB/T 5494-2008"
	productInfomationX = append(productInfomationX, productItemX2)

	productItemX3 := make(map[string]interface{})
	productItemX3["inspectionProject"] = "杂质总量"
	productItemX3["measurementUnit"] = "%"
	productItemX3["standardRequirements"] = "<=0.25"
	productItemX3["testResults"] = "0.00"
	productItemX3["methodDetectionLimit"] = "——"
	productItemX3["singleConclusion"] = "合格"
	productItemX3["testMethod"] = "GB/T 5494-2008"
	productInfomationX = append(productInfomationX, productItemX3)

	productItemX4 := make(map[string]interface{})
	productItemX4["inspectionProject"] = "杂质糠粉"
	productItemX4["measurementUnit"] = "%"
	productItemX4["standardRequirements"] = "<=0.15"
	productItemX4["testResults"] = "0.00"
	productItemX4["methodDetectionLimit"] = "——"
	productItemX4["singleConclusion"] = "合格"
	productItemX4["testMethod"] = "GB/T 5494-2008"
	productInfomationX = append(productInfomationX, productItemX4)

	productItemX5 := make(map[string]interface{})
	productItemX5["inspectionProject"] = "杂质矿物质"
	productItemX5["measurementUnit"] = "%"
	productItemX5["standardRequirements"] = "<=0.02"
	productItemX5["testResults"] = "0.00"
	productItemX5["methodDetectionLimit"] = "——"
	productItemX5["singleConclusion"] = "合格"
	productItemX5["testMethod"] = "GB/T 5494-2008"
	productInfomationX = append(productInfomationX, productItemX5)

	productItemX6 := make(map[string]interface{})
	productItemX6["inspectionProject"] = "杂质带壳稗粒"
	productItemX6["measurementUnit"] = "粒/kg"
	productItemX6["standardRequirements"] = "<=3"
	productItemX6["testResults"] = "0"
	productItemX6["methodDetectionLimit"] = "——"
	productItemX6["singleConclusion"] = "合格"
	productItemX6["testMethod"] = "GB/T 5494-2008"
	productInfomationX = append(productInfomationX, productItemX6)

	productItemX7 := make(map[string]interface{})
	productItemX7["inspectionProject"] = "杂质稻谷粒"
	productItemX7["measurementUnit"] = "粒/kg"
	productItemX7["standardRequirements"] = "<=4"
	productItemX7["testResults"] = "0"
	productItemX7["methodDetectionLimit"] = "——"
	productItemX7["singleConclusion"] = "合格"
	productItemX7["testMethod"] = "GB/T 5494-2008"
	productInfomationX = append(productInfomationX, productItemX7)

	productItemX8 := make(map[string]interface{})
	productItemX8["inspectionProject"] = "碎米总量"
	productItemX8["measurementUnit"] = "%"
	productItemX8["standardRequirements"] = "<=7.5"
	productItemX8["testResults"] = "2.3"
	productItemX8["methodDetectionLimit"] = "——"
	productItemX8["singleConclusion"] = "合格"
	productItemX8["testMethod"] = "GB/T 5503-2009"
	productInfomationX = append(productInfomationX, productItemX8)

	productItemX9 := make(map[string]interface{})
	productItemX9["inspectionProject"] = "碎米中小碎米含量"
	productItemX9["measurementUnit"] = "%"
	productItemX9["standardRequirements"] = "<=0.5"
	productItemX9["testResults"] = "0.0"
	productItemX9["methodDetectionLimit"] = "——"
	productItemX9["singleConclusion"] = "合格"
	productItemX9["testMethod"] = "GB/T 5503-2009"
	productInfomationX = append(productInfomationX, productItemX9)

	productItemX10 := make(map[string]interface{})
	productItemX10["inspectionProject"] = "黄粒米"
	productItemX10["measurementUnit"] = "%"
	productItemX10["standardRequirements"] = "<=0.5"
	productItemX10["testResults"] = "0.0"
	productItemX10["methodDetectionLimit"] = "——"
	productItemX10["singleConclusion"] = "合格"
	productItemX10["testMethod"] = "GB/T 5496-1985"
	productInfomationX = append(productInfomationX, productItemX10)

	productItemX11 := make(map[string]interface{})
	productItemX11["inspectionProject"] = "互混"
	productItemX11["measurementUnit"] = "%"
	productItemX11["standardRequirements"] = "<=5.0"
	productItemX11["testResults"] = "0.0"
	productItemX11["methodDetectionLimit"] = "——"
	productItemX11["singleConclusion"] = "合格"
	productItemX11["testMethod"] = "GB/T 5493-2008"
	productInfomationX = append(productInfomationX, productItemX11)

	productItemX12 := make(map[string]interface{})
	productItemX12["inspectionProject"] = "水分"
	productItemX12["measurementUnit"] = "%"
	productItemX12["standardRequirements"] = "<=15.5"
	productItemX12["testResults"] = "15.0"
	productItemX12["methodDetectionLimit"] = "——"
	productItemX12["singleConclusion"] = "合格"
	productItemX12["testMethod"] = "GB/T 5497-1985"
	productInfomationX = append(productInfomationX, productItemX12)

	productItemX13 := make(map[string]interface{})
	productItemX13["inspectionProject"] = "直链淀粉含量（干基）"
	productItemX13["measurementUnit"] = "%"
	productItemX13["standardRequirements"] = "13.0-20.0"
	productItemX13["testResults"] = "17.89"
	productItemX13["methodDetectionLimit"] = "——"
	productItemX13["singleConclusion"] = "合格"
	productItemX13["testMethod"] = "NY/T 83-1988"
	productInfomationX = append(productInfomationX, productItemX13)

	productItemX14 := make(map[string]interface{})
	productItemX14["inspectionProject"] = "垩白度"
	productItemX14["measurementUnit"] = "%"
	productItemX14["standardRequirements"] = "<=5"
	productItemX14["testResults"] = "0.5"
	productItemX14["methodDetectionLimit"] = "——"
	productItemX14["singleConclusion"] = "合格"
	productItemX14["testMethod"] = "NY/T 83-1988"
	productInfomationX = append(productInfomationX, productItemX14)

	productItemX15 := make(map[string]interface{})
	productItemX15["inspectionProject"] = "无机砷"
	productItemX15["measurementUnit"] = "mg/kg"
	productItemX15["standardRequirements"] = "<=0.15"
	productItemX15["testResults"] = "0.057"
	productItemX15["methodDetectionLimit"] = "——"
	productItemX15["singleConclusion"] = "合格"
	productItemX15["testMethod"] = "GB 5009.11-2014"
	productInfomationX = append(productInfomationX, productItemX15)

	productItemX16 := make(map[string]interface{})
	productItemX16["inspectionProject"] = "总汞"
	productItemX16["measurementUnit"] = "mg/kg"
	productItemX16["standardRequirements"] = "<=0.01"
	productItemX16["testResults"] = "0.0062"
	productItemX16["methodDetectionLimit"] = "——"
	productItemX16["singleConclusion"] = "合格"
	productItemX16["testMethod"] = "GB 5009.17-2014"
	productInfomationX = append(productInfomationX, productItemX16)

	productItemX17 := make(map[string]interface{})
	productItemX17["inspectionProject"] = "磷化物"
	productItemX17["measurementUnit"] = "mg/kg"
	productItemX17["standardRequirements"] = "<=0.01"
	productItemX17["testResults"] = "未检出"
	productItemX17["methodDetectionLimit"] = "<0.01"
	productItemX17["singleConclusion"] = "合格"
	productItemX17["testMethod"] = "GB 5009.36-2003"
	productInfomationX = append(productInfomationX, productItemX17)

	productItemX18 := make(map[string]interface{})
	productItemX18["inspectionProject"] = "乐果"
	productItemX18["measurementUnit"] = "mg/kg"
	productItemX18["standardRequirements"] = "<=0.01"
	productItemX18["testResults"] = "未检出"
	productItemX18["methodDetectionLimit"] = "<0.01"
	productItemX18["singleConclusion"] = "合格"
	productItemX18["testMethod"] = "GB/T 5009.20-2003"
	productInfomationX = append(productInfomationX, productItemX18)

	productItemX19 := make(map[string]interface{})
	productItemX19["inspectionProject"] = "敌敌畏"
	productItemX19["measurementUnit"] = "mg/kg"
	productItemX19["standardRequirements"] = "<=0.01"
	productItemX19["testResults"] = "未检出"
	productItemX19["methodDetectionLimit"] = "<0.01"
	productItemX19["singleConclusion"] = "合格"
	productItemX19["testMethod"] = "GB/T 5009.20-2003"
	productInfomationX = append(productInfomationX, productItemX19)

	productItemX20 := make(map[string]interface{})
	productItemX20["inspectionProject"] = "马拉硫磷"
	productItemX20["measurementUnit"] = "mg/kg"
	productItemX20["standardRequirements"] = "<=0.01"
	productItemX20["testResults"] = "未检出"
	productItemX20["methodDetectionLimit"] = "<0.01"
	productItemX20["singleConclusion"] = "合格"
	productItemX20["testMethod"] = "GB/T 5009.20-2003"
	productInfomationX = append(productInfomationX, productItemX20)

	productItemX21 := make(map[string]interface{})
	productItemX21["inspectionProject"] = "杀螟硫磷"
	productItemX21["measurementUnit"] = "mg/kg"
	productItemX21["standardRequirements"] = "<=0.01"
	productItemX21["testResults"] = "未检出"
	productItemX21["methodDetectionLimit"] = "<0.01"
	productItemX21["singleConclusion"] = "合格"
	productItemX21["testMethod"] = "GB/T 5009.20-2003"
	productInfomationX = append(productInfomationX, productItemX21)

	productItemX22 := make(map[string]interface{})
	productItemX22["inspectionProject"] = "三唑磷"
	productItemX22["measurementUnit"] = "mg/kg"
	productItemX22["standardRequirements"] = "<=0.01"
	productItemX22["testResults"] = "未检出"
	productItemX22["methodDetectionLimit"] = "<0.00034"
	productItemX22["singleConclusion"] = "合格"
	productItemX22["testMethod"] = "GB/T 20770-2008"
	productInfomationX = append(productInfomationX, productItemX22)

	productItemX23 := make(map[string]interface{})
	productItemX23["inspectionProject"] = "克百威"
	productItemX23["measurementUnit"] = "mg/kg"
	productItemX23["standardRequirements"] = "<=0.01"
	productItemX23["testResults"] = "未检出"
	productItemX23["methodDetectionLimit"] = "<0.00653"
	productItemX23["singleConclusion"] = "合格"
	productItemX23["testMethod"] = "GB/T 20770-2008"
	productInfomationX = append(productInfomationX, productItemX23)

	productItemX24 := make(map[string]interface{})
	productItemX24["inspectionProject"] = "甲胺磷"
	productItemX24["measurementUnit"] = "mg/kg"
	productItemX24["standardRequirements"] = "<=0.01"
	productItemX24["testResults"] = "未检出"
	productItemX24["methodDetectionLimit"] = "<0.004"
	productItemX24["singleConclusion"] = "合格"
	productItemX24["testMethod"] = "GB/T 5009.103-2003"
	productInfomationX = append(productInfomationX, productItemX24)

	productItemX25 := make(map[string]interface{})
	productItemX25["inspectionProject"] = "杀虫双"
	productItemX25["measurementUnit"] = "mg/kg"
	productItemX25["standardRequirements"] = "<=0.01"
	productItemX25["testResults"] = "未检出"
	productItemX25["methodDetectionLimit"] = "<0.002"
	productItemX25["singleConclusion"] = "合格"
	productItemX25["testMethod"] = "GB/T 5009.114-2003"
	productInfomationX = append(productInfomationX, productItemX25)

	productItemX26 := make(map[string]interface{})
	productItemX26["inspectionProject"] = "溴氯菊酯"
	productItemX26["measurementUnit"] = "mg/kg"
	productItemX26["standardRequirements"] = "<=0.01"
	productItemX26["testResults"] = "未检出"
	productItemX26["methodDetectionLimit"] = "<0.00088"
	productItemX26["singleConclusion"] = "合格"
	productItemX26["testMethod"] = "GB/T 5009.110-2003"
	productInfomationX = append(productInfomationX, productItemX26)

	productItemX27 := make(map[string]interface{})
	productItemX27["inspectionProject"] = "水胺硫磷"
	productItemX27["measurementUnit"] = "mg/kg"
	productItemX27["standardRequirements"] = "<=0.01"
	productItemX27["testResults"] = "未检出"
	productItemX27["methodDetectionLimit"] = "<0.01"
	productItemX27["singleConclusion"] = "合格"
	productItemX27["testMethod"] = "GB/T 20770-2008"
	productInfomationX = append(productInfomationX, productItemX27)

	productItemX28 := make(map[string]interface{})
	productItemX28["inspectionProject"] = "稻瘟灵"
	productItemX28["measurementUnit"] = "mg/kg"
	productItemX28["standardRequirements"] = "<=0.01"
	productItemX28["testResults"] = "未检出"
	productItemX28["methodDetectionLimit"] = "<0.01"
	productItemX28["singleConclusion"] = "合格"
	productItemX28["testMethod"] = "GB/T 5009.155-2003"
	productInfomationX = append(productInfomationX, productItemX28)

	productItemX29 := make(map[string]interface{})
	productItemX29["inspectionProject"] = "三环唑"
	productItemX29["measurementUnit"] = "mg/kg"
	productItemX29["standardRequirements"] = "<=0.01"
	productItemX29["testResults"] = "未检出"
	productItemX29["methodDetectionLimit"] = "<0.01"
	productItemX29["singleConclusion"] = "合格"
	productItemX29["testMethod"] = "GB/T 5009.155-2003"
	productInfomationX = append(productInfomationX, productItemX29)

	productItemX30 := make(map[string]interface{})
	productItemX30["inspectionProject"] = "三环唑"
	productItemX30["measurementUnit"] = "mg/kg"
	productItemX30["standardRequirements"] = "<=0.01"
	productItemX30["testResults"] = "未检出"
	productItemX30["methodDetectionLimit"] = "<0.01"
	productItemX30["singleConclusion"] = "合格"
	productItemX30["testMethod"] = "GB/T 5009.115-2003"
	productInfomationX = append(productInfomationX, productItemX30)

	productItemX31 := make(map[string]interface{})
	productItemX31["inspectionProject"] = "丁草胺"
	productItemX31["measurementUnit"] = "mg/kg"
	productItemX31["standardRequirements"] = "<=0.01"
	productItemX31["testResults"] = "未检出"
	productItemX31["methodDetectionLimit"] = "<0.01"
	productItemX31["singleConclusion"] = "合格"
	productItemX31["testMethod"] = "GB/T 20770-2008"
	productInfomationX = append(productInfomationX, productItemX31)

	productItemX32 := make(map[string]interface{})
	productItemX32["inspectionProject"] = "铅"
	productItemX32["measurementUnit"] = "mg/kg"
	productItemX32["standardRequirements"] = "<=0.2"
	productItemX32["testResults"] = "0.04"
	productItemX32["methodDetectionLimit"] = "——"
	productItemX32["singleConclusion"] = "合格"
	productItemX32["testMethod"] = "GB/T 5009.12-2010"
	productInfomationX = append(productInfomationX, productItemX32)

	productItemX33 := make(map[string]interface{})
	productItemX33["inspectionProject"] = "镉"
	productItemX33["measurementUnit"] = "mg/kg"
	productItemX33["standardRequirements"] = "<=0.2"
	productItemX33["testResults"] = "0.013"
	productItemX33["methodDetectionLimit"] = "——"
	productItemX33["singleConclusion"] = "合格"
	productItemX33["testMethod"] = "GB/T 5009.15-2014"
	productInfomationX = append(productInfomationX, productItemX33)

	productItemX34 := make(map[string]interface{})
	productItemX34["inspectionProject"] = "吡虫啉"
	productItemX34["measurementUnit"] = "mg/kg"
	productItemX34["standardRequirements"] = "<=0.05"
	productItemX34["testResults"] = "未检出"
	productItemX34["methodDetectionLimit"] = "<0.011"
	productItemX34["singleConclusion"] = "合格"
	productItemX34["testMethod"] = "GB/T 20770-2008"
	productInfomationX = append(productInfomationX, productItemX34)

	productItemX35 := make(map[string]interface{})
	productItemX35["inspectionProject"] = "噻嗪酮"
	productItemX35["measurementUnit"] = "mg/kg"
	productItemX35["standardRequirements"] = "<=0.3"
	productItemX35["testResults"] = "未检出"
	productItemX35["methodDetectionLimit"] = "<0.01"
	productItemX35["singleConclusion"] = "合格"
	productItemX35["testMethod"] = "GB/T 5009.184-2003"
	productInfomationX = append(productInfomationX, productItemX35)

	productItemX36 := make(map[string]interface{})
	productItemX36["inspectionProject"] = "毒死蜱"
	productItemX36["measurementUnit"] = "mg/kg"
	productItemX36["standardRequirements"] = "<=0.1"
	productItemX36["testResults"] = "未检出"
	productItemX36["methodDetectionLimit"] = "<0.008"
	productItemX36["singleConclusion"] = "合格"
	productItemX36["testMethod"] = "GB/T 5009.145-2003"
	productInfomationX = append(productInfomationX, productItemX36)

	productItemX37 := make(map[string]interface{})
	productItemX37["inspectionProject"] = "黄曲霉毒素B1"
	productItemX37["measurementUnit"] = "ug/kg"
	productItemX37["standardRequirements"] = "<=5.0"
	productItemX37["testResults"] = "<5.0"
	productItemX37["methodDetectionLimit"] = "——"
	productItemX37["singleConclusion"] = "合格"
	productItemX37["testMethod"] = "GB/T 5009.22-2003"
	productInfomationX = append(productInfomationX, productItemX37)

	/*
		type InspectionReportType struct {
			SampleName                        string            `json:""`                        //样品名称
			ModelSpecification                string            `json:""`                //型号规格
			SampleNumber                      string            `json:""`                      //样品编号
			Trademark                         string            `json:""`                         //商标
			InspectionUnit                    string            `json:""`                    //送检单位
			InspectionCategory                string            `json:""`                //检验类别
			ProductionUnit                    string            `json:""`                    //生产单位
			SampleGradeAndState               string            `json:""`               //样品等级、状态
			SamplingSite                      string            `json:""`                      //抽样地点
			DateToSample                      string            `json:""`                      //到样日期
			SampleQuantity                    string            `json:""`                    //样品数量
			SampleMaker                       string            `json:""`                       //送样者
			SamplingBase                      string            `json:""`                      //抽样基数
			OriginalNumberOrDateOfProduction  string            `json:""`  //原编号或生产日期
			InspectionBasis                   string            `json:""`                   //检验依据
			InspectionProject                 string            `json:""`                 //检验项目
			MainInstrumentsUsed               string            `json:""`               //所用主要仪器
			ExperimentalEnvironmentConditions string            `json:""` //实验环境条件
			InspectionConclusion              string            `json:""`              //检验结论
			Remarks                           string            `json:""`                           //备注
			Approver                          string            `json:""`                          //批准人
			DateOfApproval                    string            `json:""`                    //批准日期
			Auditor                           string            `json:""`                           //审核人
			DateOfAudit                       string            `json:""`                       //审核日期
			TabulatingPerson                  string            `json:""`                  //制表人
			DateOfTabulation                  string            `json:""`                  //制表日期
			Items                             []ProductItemType `json:""`                             //详细指标
			Picture1Cover                     string            `json:""`                     //封面图片
			PictureBaseSoil                   string            `json:""`                   //基地土壤图片
			PictureIrrigatedWaterSource       string            `json:""`                   //灌溉水源图片
		}
	*/
	inspectionReportX := make(map[string]interface{})
	inspectionReportX["sampleName"] = "有机大米"
	inspectionReportX["modelSpecification"] = "————"
	inspectionReportX["sampleNumber"] = "2017C3152"
	inspectionReportX["trademark"] = "————"
	inspectionReportX["inspectionUnit"] = "五常市饭包儿水稻种植专业合作社"
	inspectionReportX["inspectionCategory"] = "委托"
	inspectionReportX["productionUnit"] = "————"
	inspectionReportX["sampleGradeAndState"] = "正常籽粒"
	inspectionReportX["samplingSite"] = "————"
	inspectionReportX["dateToSample"] = "2017-07-06"
	inspectionReportX["sampleQuantity"] = "5kg"
	inspectionReportX["sampleMaker"] = "韩剑东"
	inspectionReportX["samplingBase"] = "————"
	inspectionReportX["originalNumberOrDateOfProduction"] = "————"
	inspectionReportX["inspectionBasis"] = "NY/T 419-2014"
	inspectionReportX["inspectionProject"] = "见报告第2页"
	inspectionReportX["mainInstrumentsUsed"] = "气相色谱仪液相色谱质谱联用仪"
	inspectionReportX["experimentalEnvironmentConditions"] = "符合实验条件"
	inspectionReportX["inspectionConclusion"] = "该样品按NY/T 419-2014标准检验合格"
	inspectionReportX["remarks"] = "只对来样负责"
	inspectionReportX["approver"] = "张瑞英"
	inspectionReportX["dateOfApproval"] = "2017-07-25"
	inspectionReportX["auditor"] = "马永华"
	inspectionReportX["dateOfAudit"] = "2017-07-25"
	inspectionReportX["tabulatingPerson"] = "盛慧"
	inspectionReportX["dateOfTabulation"] = "2017-07-25"
	inspectionReportX["items"] = productInfomationX
	inspectionReportX["picture1Cover"] = "suyuan1/2.jianyan1.png"
	inspectionReportX["pictureBaseSoil"] = "suyuan1/3.jianyan2.png"
	inspectionReportX["pictureIrrigatedWaterSource"] = "suyuan1/4.jianyan3.png"

	/*五、加工工艺专利技术证书*/
	/*
		type PatentInfoType struct {
			Items []PatentItemType `json:"items"`
		}
		type PatentItemType struct {
			PatentName              string `json:"patentName"`              //专利名称
			PatentCertificateNumber string `json:"patentCertificateNumber"` //专利证书编号
			PictureName             string `json:"pictureName"`             //专利证书图片
		}
	*/
	var patentInfoX []map[string]interface{}

	patentItemX1 := make(map[string]interface{})
	patentItemX1["patentName"] = "稻谷碾白机"
	patentItemX1["patentCertificateNumber"] = "第1978111号"
	patentItemX1["pictureName"] = "suyuan1/5.zhuanli1.png"
	patentInfoX = append(patentInfoX, patentItemX1)

	patentItemX2 := make(map[string]interface{})
	patentItemX2["patentName"] = "胚芽米精磨机"
	patentItemX2["patentCertificateNumber"] = "第2619488号"
	patentItemX2["pictureName"] = "suyuan1/6.zhuanli2.png"
	patentInfoX = append(patentInfoX, patentItemX2)

	patentItemX3 := make(map[string]interface{})
	patentItemX3["patentName"] = "一种双室胚芽米碾米机"
	patentItemX3["patentCertificateNumber"] = "第6479912号"
	patentItemX3["pictureName"] = "suyuan1/7.zhuanli3.png"
	patentInfoX = append(patentInfoX, patentItemX3)

	patentItemX4 := make(map[string]interface{})
	patentItemX4["patentName"] = "一种米糠、谷壳粉碎装置"
	patentItemX4["patentCertificateNumber"] = "第6479920号"
	patentItemX4["pictureName"] = "suyuan1/8.zhuanli4.png"
	patentInfoX = append(patentInfoX, patentItemX4)

	/*六、留胚率、完整度检测报告*/
	/*
	   type DetectionReportOfEmbryoRateAndIntegrityType struct {
	   	Picture1 string `json:"picture1"` //图片1
	   	Picture2 string `json:"picture2"` //图片2
	   	Picture3 string `json:"picture3"` //图片3
	   	Picture4 string `json:"picture4"` //图片4
	   }
	*/
	detectionReportOfEmbryoRateAndIntegrityX := make(map[string]interface{})
	detectionReportOfEmbryoRateAndIntegrityX["picture1"] = "suyuan1/9.liupeilv1.png"
	detectionReportOfEmbryoRateAndIntegrityX["picture2"] = "suyuan1/10.liupeilv2.png"
	detectionReportOfEmbryoRateAndIntegrityX["picture3"] = "suyuan1/11.liupeilv3.png"
	detectionReportOfEmbryoRateAndIntegrityX["picture4"] = "suyuan1/12.liupeilv4.png"

	/*七、位置信息*/
	/*
		type PositionInformationType struct {
			Picture1PlantingBaseLocation   string `json:"picture1PlantingBaseLocation"`   //种植基地
			Picture2StorageBaseLocation    string `json:"picture2StorageBaseLocation"`    //仓储基地
			Picture3ProcessingBaseLocation string `json:"picture3ProcessingBaseLocation"` //加工基地
		}
	*/
	positionInformationX := make(map[string]interface{})
	positionInformationX["picture1PlantingBaseLocation"] = "suyuan1/13.weizhi1.png"
	positionInformationX["picture2StorageBaseLocation"] = "suyuan1/14.weizhi2.png"
	positionInformationX["picture3ProcessingBaseLocation"] = "suyuan1/15.weizhi3.png"

	/*八、产品出厂检测报告*/
	/*
		type ProductInspectionReportType struct {
			Picture1 string `json:"picture1"` //图片1
			Picture2 string `json:"picture2"` //图片2
		}
	*/
	productInspectionReportX := make(map[string]interface{})
	productInspectionReportX["picture1"] = "suyuan1/16.baogao1.png"
	productInspectionReportX["picture2"] = "suyuan1/17.baogao2.png"

	//Overall BatchInfo
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
	batchInformation := make(map[string]interface{})
	batchInformation["batchNumber"] = "第一批次"
	batchInformation["batchOutput"] = "10000斤（每斤一个最小包装单位）"
	batchInformation["seedInfo"] = seedInfoX
	batchInformation["biologicalOrganicFertilizer"] = biologicalOrganicFertilizerX
	batchInformation["organicCertificationOfBase"] = organicCertificationOfBaseX
	batchInformation["inspectionReport"] = inspectionReportX
	batchInformation["patentInfo"] = patentInfoX
	batchInformation["detectionReportOfEmbryoRateAndIntegrity"] = detectionReportOfEmbryoRateAndIntegrityX
	batchInformation["positionInformation"] = positionInformationX
	batchInformation["productInspectionReport"] = productInspectionReportX

	//testRecord, _ := json.Marshal(batchInformation)
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
