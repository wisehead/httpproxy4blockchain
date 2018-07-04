package handler

/*
批次信息
第一批次，该批次总生产量5000份
批次及产量：第一批次，10000斤（每斤一个最小包装单位）
*/
type BatchInformationType struct {
	BatchNumber                             string                                      `json:"batchNumber"`
	BatchOutput                             string                                      `json:"batchOutput"`
	ReadCount                               uint64                                      `json:"readCount"`
	AccessTime                              uint64                                      `json:"accessTime"`
	TasteScore                              string                                      `json:"tasteScore"`
	SeedInfo                                SeedInfoType                                `json:"seedInfo"`                                //一、种子信息：
	BiologicalOrganicFertilizer             BiologicalOrganicFertilizerType             `json:"biologicalOrganicFertilizer"`             //二、生物有机肥
	OrganicCertificationOfBase              OrganicCertificationOfBaseType              `json:"organicCertificationOfBase"`              //三、基地有机认证
	InspectionReport                        InspectionReportType                        `json:"inspectionReport"`                        //四、检验报告
	PatentInfo                              PatentInfoType                              `json:"patentInfo"`                              //五、加工工艺专利技术证书
	DetectionReportOfEmbryoRateAndIntegrity DetectionReportOfEmbryoRateAndIntegrityType `json:"detectionReportOfEmbryoRateAndIntegrity"` //六、留胚率、完整度检测报告
	PositionInformation                     PositionInformationType                     `json:"positionInformation"`                     //七、位置信息
	ProductInspectionReport                 ProductInspectionReportType                 `json:"productInspectionReport"`                 //八、产品出厂检测报告
}

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
	StandardName          string `json:"standardName"`          //纯度99%
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
	PictureIrrigatedWaterSource       string            `json:"pictureIrrigatedWaterSource"`       //灌溉水源图片
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
	PatentCertificateNumber string `json:"patentCertificateNumber"` //专利证书编号
	PictureName             string `json:"pictureName"`             //专利证书图片
}

/*六、留胚率、完整度检测报告*/
type DetectionReportOfEmbryoRateAndIntegrityType struct {
	Picture1 string `json:"picture1"` //图片1
	Picture2 string `json:"picture2"` //图片2
	Picture3 string `json:"picture3"` //图片3
	Picture4 string `json:"picture4"` //图片4
}

/*七、位置信息*/
type PositionInformationType struct {
	Position1PlantingBaseLocation   string `json:"position1PlantingBaseLocation"`   //种植基地
	Position2StorageBaseLocation    string `json:"position2StorageBaseLocation"`    //仓储基地
	Position3ProcessingBaseLocation string `json:"position3ProcessingBaseLocation"` //加工基地
}

/*八、产品出厂检测报告*/
type ProductInspectionReportType struct {
	Picture1 string `json:"picture1"` //图片1
	Picture2 string `json:"picture2"` //图片2
}
