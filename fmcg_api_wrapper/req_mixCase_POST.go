package fmcg_api_wrapper

import "fmt"

type FmcgProductBodyMixCase struct {
	GTIN     string `json:"D8165"` // Barcode with 0 in front
	DataType string `json:"D8164"` // [CORRECT, ...] // https://simplychocolate.fmcgproducts.dk/fmcg/pa/simplychocolate/pa.nsf/keyword.xsp?id=DocumentCommandHeaderCode.da

	// General Information
	DataCarrierTypeCode                 string `json:"D8208"` // [EAN_13, ...] - https://simplychocolate.fmcgproducts.dk/fmcg/pa/simplychocolate/pa.nsf/keyword.xsp?id=DataCarrierTypeCodeList.da
	BrandName                           string `json:"D8211"`
	CountryOfOrigin                     string `json:"D8219"`   // [208, ...] - https://simplychocolate.fmcgproducts.dk/fmcg/pa/simplychocolate/pa.nsf/keyword.xsp?id=CountryCodeList.da
	ManufacturerGLN                     string `json:"D8242_1"` //
	BrandOwnerGLN                       string `json:"D8346"`   //
	GPCCategoryCode                     string `json:"D8245"`   // 10000045
	TargetMarketCode                    string `json:"D8255"`   // Default 208 for DK
	ItemCode                            string `json:"D8256"`
	ItemNameDA                          string `json:"D8258_1"` // The name of the item in the language specified in D8259_1
	ItemNameLanguageCodeDA              string `json:"D8259_1"` // [da, ...] - https://simplychocolate.fmcgproducts.dk/fmcg/pa/simplychocolate/pa.nsf/keyword.xsp?id=LanguageCodeList.da
	FunctionalProductNameDA             string `json:"D8313_1"`
	FunctionalProductNameLanguageCodeDA string `json:"D8121_1"` // [da, ...] - https://simplychocolate.fmcgproducts.dk/fmcg/pa/simplychocolate/pa.nsf/keyword.xsp?id=LanguageCodeList.da

	IsOrderingUnit                bool   `json:"D8271"` // [TRUE, FALSE] (True for cases and displays, False for BASE_UNIT)
	UnitOfMeasure                 string `json:"D8276"` // [BASE_UNIT_OR_EACH, CASE, PALLET, DISPLAY_SHIPPER] - https://simplychocolate.fmcgproducts.dk/fmcg/pa/simplychocolate/pa.nsf/keyword.xsp?id=TradeItemUnitDescriptorCodeList.da
	ShelfLifeFromArrivalInDays    int    `json:"D8283"` // SAP FIELD: U_BOYX_Holdbarhed_Kunde
	ShelfLifeFromProductionInDays int    `json:"D8284"` // SAP FIELD: U_BOYX_Holdbarhed
	IsQuantityOrPriceVarying      bool   `json:"D8297"` // [TRUE, FALSE]
	DangerousContent              string `json:"D8030"` // [NOT_APPLICABLE, TRUE, FALSE, UNSPECIFIED] - https://simplychocolate.fmcgproducts.dk/fmcg/pa/simplychocolate/pa.nsf/keyword.xsp?id=NonBinaryLogicEnumerationCodeList.da
	RelevantForPriceComparison    string `json:"D8019"` // [NOT_APPLICABLE, TRUE, FALSE, UNSPECIFIED] - https://simplychocolate.fmcgproducts.dk/fmcg/pa/simplychocolate/pa.nsf/keyword.xsp?id=NonBinaryLogicEnumerationCodeList.da
	IsConsumerUnit                bool   `json:"D8216"` // [TRUE, FALSE]
	IsShippingUnit                bool   `json:"D8236"` // [TRUE, FALSE]
	IsPackagingMarkedReturnable   bool   `json:"D8311"` // [TRUE, FALSE]
	IsPackageSalesReady           string `json:"D3111"` // [NOT_APPLICABLE, TRUE, FALSE, UNSPECIFIED] - https://simplychocolate.fmcgproducts.dk/fmcg/pa/simplychocolate/pa.nsf/keyword.xsp?id=NonBinaryLogicEnumerationCodeList.da
	EffectiveDateTime             string `json:"D8286"` // DateTime
	StartAvailabilityDateTime     string `json:"D8314"` // DateTime

	// Logistical Information
	NetWeight      int    `json:"D8068"`
	NetWeightUoM   string `json:"D8069"`   // [GRM, ...] - https://simplychocolate.fmcgproducts.dk/fmcg/pa/simplychocolate/pa.nsf/keyword.xsp?id=MeasurementUnitCodeList.da
	NetContent     int    `json:"D8217_1"` //
	NetContentUoM  string `json:"D8218_1"` // [GRM, ...] - https://simplychocolate.fmcgproducts.dk/fmcg/pa/simplychocolate/pa.nsf/keyword.xsp?id=MeasurementUnitCodeList.da
	GrossWeight    int    `json:"D8246"`
	GrossWeightUoM string `json:"D8247"`   // [GRM, ...] - https://simplychocolate.fmcgproducts.dk/fmcg/pa/simplychocolate/pa.nsf/keyword.xsp?id=MeasurementUnitCodeList.da
	Height         int    `json:"D8263"`   //
	HeightUOM      string `json:"D8264"`   // [MMT, ...] - https://simplychocolate.fmcgproducts.dk/fmcg/pa/simplychocolate/pa.nsf/keyword.xsp?id=MeasurementUnitCodeList.daWidth
	Depth          int    `json:"D8265"`   //
	DepthUOM       string `json:"D8266"`   // [MMT, ...] - https://simplychocolate.fmcgproducts.dk/fmcg/pa/simplychocolate/pa.nsf/keyword.xsp?id=MeasurementUnitCodeList.daWidth
	Width          int    `json:"D8267"`   //
	WidthUOM       string `json:"D8268"`   // [MMT, ...] - https://simplychocolate.fmcgproducts.dk/fmcg/pa/simplychocolate/pa.nsf/keyword.xsp?id=MeasurementUnitCodeList.daWidth
	PackagingType  string `json:"D8275_1"` // [WRP, BX, JR] - https://simplychocolate.fmcgproducts.dk/fmcg/pa/simplychocolate/pa.nsf/keyword.xsp?id=PackageTypeCodeList.da

	UnitsPerCaseItem1  int    `json:"D8270_1"`
	UnitGTINItem1      string `json:"D8249_1"`
	UnitsPerCaseItem2  int    `json:"D8270_2"`
	UnitGTINItem2      string `json:"D8249_2"`
	UnitsPerCaseItem3  int    `json:"D8270_3"`
	UnitGTINItem3      string `json:"D8249_3"`
	UnitsPerCaseItem4  int    `json:"D8270_4"`
	UnitGTINItem4      string `json:"D8249_4"`
	UnitsPerCaseItem5  int    `json:"D8270_5"`
	UnitGTINItem5      string `json:"D8249_5"`
	UnitsPerCaseItem6  int    `json:"D8270_6"`
	UnitGTINItem6      string `json:"D8249_6"`
	UnitsPerCaseItem7  int    `json:"D8270_7"`
	UnitGTINItem7      string `json:"D8249_7"`
	UnitsPerCaseItem8  int    `json:"D8270_8"`
	UnitGTINItem8      string `json:"D8249_8"`
	UnitsPerCaseItem9  int    `json:"D8270_9"`
	UnitGTINItem9      string `json:"D8249_9"`
	UnitsPerCaseItem10 int    `json:"D8270_10"`
	UnitGTINItem10     string `json:"D8249_10"`
	UnitsPerCaseItem11 int    `json:"D8270_11"`
	UnitGTINItem11     string `json:"D8249_11"`
	UnitsPerCaseItem12 int    `json:"D8270_12"`
	UnitGTINItem12     string `json:"D8249_12"`
	UnitsPerCaseItem13 int    `json:"D8270_13"`
	UnitGTINItem13     string `json:"D8249_13"`
	UnitsPerCaseItem14 int    `json:"D8270_14"`
	UnitGTINItem14     string `json:"D8249_14"`
	UnitsPerCaseItem15 int    `json:"D8270_15"`
	UnitGTINItem15     string `json:"D8249_15"`
}

func FMCGApiPostMixCase(caseInfo FmcgProductBodyMixCase, count int) error {
	resp, err := GetFMCGApiBaseClient().
		//DevMode().
		R().
		EnableDump().
		SetResult(FmcgProductPostResult{}).
		SetBody(caseInfo).
		Post("")
	if err != nil {
		return err
	}

	if resp.IsError() {
		fmt.Printf("resp is err statusCode: %v. Dump: %v\n", resp.StatusCode, resp.Dump())
		return resp.Err
	}

	response := resp.Result().(*FmcgProductPostResult)
	for _, validationError := range response.ValidationErrors {
		fmt.Println("fieldId:", validationError.FieldId)
		fmt.Println("fieldLabel:", validationError.FieldLabel)
		fmt.Println("message:", validationError.Message)
		fmt.Println("messageType:", validationError.MessageType)
		fmt.Println("________________")
	}

	return nil
}