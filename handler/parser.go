package handler

/*
一、种子信息：
*/
type QulityType struct {
	UnpolishedRiceRate string `json:"unpolishedRicePercentage"` //糙米率84.1%
	PolishedRiceRate   string `json:"polishedRiceRate"`         //精米率75.7%
	HeadRiceRate       string `json:"headRiceRate"`             //整精密率66.8%
	GrainLength        string `json:"grainLength"`              //粒长6.6mm
	GelConsistency     string `json:"gelConsistency"`           //胶稠度67.0%
	TasteScore         string `json:"tasteScore"`               //食味评分88-92
}

type QualityExecutionStandardType struct {
	Purity                string `json:"purity"`                //纯度99%
	Cleanliness           string `json:"cleanliness"`           //净度98.0%
	GerminationPercentage string `json:"germinationPercentage"` //发芽率85%
	WaterContent          string `json:"waterContent"`          //水分16.0%
}

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
	BreedsOwner                               string                       `json:"breedsOwner"`                               //14.	五优稻4号(稻花香2号)品种权人：五常市利元种子有根公司
	NetContent                                string                       `json:"netContent"`                                //15.	净含量：25kg
	SeedProducer                              string                       `json:"seedProducer"`                              //16.	种子生产经营者：五常市利元种子有限公司
	RegisteredAddress                         string                       `json:"registeredAddress"`                         //17.	注册地地址：黑龙江省五常市龙凤山镇乐园村汪家店屯
	Website                                   string                       `json:"website"`                                   //18.	网址:www.clxseed.com
}

/*二、生物有机肥*/
type ProductMarkingType struct {
	MainTechnicalIndicators string `json:"mainTechnicalIndicators"` //1)主要技术指标：有效活菌数≥0.2亿/g  N+P2O5+K2O≥15% 有机质≥20%
	ScientificAddition      string `json:"scientificAddition"`      //2)科学添加:腐植酸≥12% 氨基酸≥1% 1酵素酶≥0.5亿/g中微量元素≥18%有益微生物蛋白酶及生长促进未知因子(UGF)
}
type BiologicalOrganicFertilizerType struct {
	Brand                                        string             `json:"brand"`                                        //1.	品牌：归农稻花香有机米专用
	LicenseNumberOfGreenFoodProductionData       string             `json:"licenseNumberOfGreenFoodProductionData"`       //2.	绿色食品生产资料许可编号：LSSZ-01-1503080005（中国绿色食品协会）
	ProductionTechnology                         string             `json:"productionTechnology"`                         //3.	生产工艺技术：酵素生物技术、复合微生物肥
	RegistrationNumberOfTheMinistryOfAgriculture string             `json:"registrationNumberOfTheMinistryOfAgriculture"` //4.	农业部登记证号：微生物肥（2016）准字（1896号）
	StandardOfExecution                          string             `json:"standardOfExecution"`                          //5.	执行标准：NY/T798-2015
	ProductMarking                               ProductMarkingType `json:"productMarking"`                               //6.	产品标值：
	PackageNumber                                string             `json:"packageNumber"`                                //7.	包装编号G16号
	Manufacturer                                 string             `json:"manufacturer"`                                 //8.	制造商：哈尔滨绿洲之星生物科技有限公司
	Address                                      string             `json:"address"`                                      //9.	地址:哈尔滨市香坊区印染街1号
	Website                                      string             `json:"website"`                                      //10.	网址：www-hrblz.com
}

/*三、基地有机认证*/
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

/*四、检验报告*/
type InspectionReportType struct {
	SampleName                        string            `json:"SampleName"`                        //样品名称
	ModelSpecification                string            `json:"ModelSpecification"`                //型号规格
	SampleNumber                      string            `json:"SampleNumber"`                      //样品编号
	trademark                         string            `json:"trademark"`                         //商标
	InspectionUnit                    string            `json:"InspectionUnit"`                    //送检单位
	InspectionCategory                string            `json:"InspectionCategory"`                //检验类别
	ProductionUnit                    string            `json:"ProductionUnit"`                    //生产单位
	SampleGradeAndState               string            `json:"SampleGradeAndState"`               //样品等级、状态
	SamplingSite                      string            `json:"SamplingSite"`                      //抽样地点
	DateToSample                      string            `json:"DateToSample"`                      //到样日期
	SampleQuantity                    string            `json:"SampleQuantity"`                    //样品数量
	SampleMaker                       string            `json:"SampleMaker"`                       //送样者
	SamplingBase                      string            `json:"SamplingBase"`                      //抽样基数
	OriginalNumberOrDateOfProduction  string            `json:"OriginalNumberOrDateOfProduction"`  //原编号或生产日期
	InspectionBasis                   string            `json:"InspectionBasis"`                   //检验依据
	InspectionProject                 string            `json:"InspectionProject"`                 //检验项目
	MainInstrumentsUsed               string            `json:"MainInstrumentsUsed"`               //所用主要仪器
	ExperimentalEnvironmentConditions string            `json:"ExperimentalEnvironmentConditions"` //实验环境条件
	InspectionConclusion              string            `json:"InspectionConclusion"`              //检验结论
	Remarks                           string            `json:"Remarks"`                           //备注
	Approver                          string            `json:"Approver"`                          //批准人
	DateOfApproval                    string            `json:"DateOfApproval"`                    //批准日期
	Auditor                           string            `json:"Auditor"`                           //审核人
	DateOfAudit                       string            `json:"DateOfAudit"`                       //审核日期
	TabulatingPerson                  string            `json:"TabulatingPerson"`                  //制表人
	DateOfTabulation                  string            `json:"DateOfTabulation"`                  //制表日期
	Items                             []ProductItemType `json:"items"`                             //详细指标
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

/*五、加工工艺专利技术证书*/
type PatentInfoType struct {
	Items []PatentItemType `json:"items"`
}
type PatentItemType struct {
	PatentName              string `json:"patentName"`              //专利名称
	PatentCertificateNumber string `json:"PatentCertificateNumber"` //专利证书编号
	PictureName             string `json:"PictureName"`             //专利证书图片
}

/*5.土壤检测报告（数据辛苦从图片中提取即可）*/
type SoilCheckReport struct {
	Items []SoilCheckReportItem `json:"items"`
}

type SoilCheckReportItem struct {
	SmapleId                 string `json:"smapleId"`
	SampleInspectionProject  string `json:"sampleInspectionProject"`
	SoilUnit                 string `json:"soilUnit"`
	SoilIndex                string `json:"soilIndex"`
	SoilMeasuredData         string `json:"soilMeasuredData"`
	SampleNumDetectionLimit  string `json:"sampleNumDetectionLimit"`
	LN166872SingleConclusion string `json:"ln166872SingleConclusion"`
	LN166872DetectionBasis   string `json:"ln166872DetectionBasis"`
}

type OneRecord struct {
	SeedInfoData                        SeedInfo                        `json:"seedInfoData"`
	BiologicalOrganicFertilizerInfoData BiologicalOrganicFertilizerInfo `json:"biologicalOrganicFertilizerInfoData"`
	OrganicAuthenticationInfoData       OrganicAuthenticationInfo       `json:"organicAuthenticationInfoData"`
	ProductInfomationData               ProductInfomation               `json:"productInfomationData"`
	SoilCheckReportData                 SoilCheckReport                 `json:"soilCheckReportData"`
}
