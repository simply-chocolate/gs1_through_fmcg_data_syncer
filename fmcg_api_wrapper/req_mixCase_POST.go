package fmcg_api_wrapper

import (
	"fmt"
	"gs1_syncer/teams_notifier"
)

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
	ShelfLifeFromArrivalInDays    int    `json:"D8283"` // SAP FIELD: U_CCF_ShelfLifeArrival
	ShelfLifeFromProductionInDays int    `json:"D8284"` // SAP FIELD: U_BOYX_Holdbarhed

	MaximumStorageTemp int    `json:"D3599_1"` //
	MinimumStorageTemp int    `json:"D3608_1"` //
	TemperatureType    string `json:"D3614_1"` // [STORAGE_HANDLING, ...] - https://simplychocolate.fmcgproducts.dk/fmcg/pa/simplychocolate/pa.nsf/keyword.xsp?id=TemperatureQualifierCodeList.da
	TemperatureOUM     string `json:"D8374_1"` // [CEL, ...] - https://simplychocolate.fmcgproducts.dk/fmcg/pa/simplychocolate/pa.nsf/keyword.xsp?id=TemperatureMeasurementUnitCodeList.da

	IsQuantityOrPriceVarying    bool   `json:"D8297"` // [TRUE, FALSE]
	DangerousContent            string `json:"D8030"` // [NOT_APPLICABLE, TRUE, FALSE, UNSPECIFIED] - https://simplychocolate.fmcgproducts.dk/fmcg/pa/simplychocolate/pa.nsf/keyword.xsp?id=NonBinaryLogicEnumerationCodeList.da
	RelevantForPriceComparison  string `json:"D8019"` // [NOT_APPLICABLE, TRUE, FALSE, UNSPECIFIED] - https://simplychocolate.fmcgproducts.dk/fmcg/pa/simplychocolate/pa.nsf/keyword.xsp?id=NonBinaryLogicEnumerationCodeList.da
	IsConsumerUnit              bool   `json:"D8216"` // [TRUE, FALSE]
	IsShippingUnit              bool   `json:"D8236"` // [TRUE, FALSE]
	IsPackagingMarkedReturnable bool   `json:"D8311"` // [TRUE, FALSE]
	IsPackageSalesReady         string `json:"D3111"` // [NOT_APPLICABLE, TRUE, FALSE, UNSPECIFIED] - https://simplychocolate.fmcgproducts.dk/fmcg/pa/simplychocolate/pa.nsf/keyword.xsp?id=NonBinaryLogicEnumerationCodeList.da
	EffectiveDateTime           string `json:"D8286"` // DateTime
	StartAvailabilityDateTime   string `json:"D8314"` // DateTime

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

	PalletGrossWeight    int    `json:"D8080"` //
	PalletGrossWeightUoM string `json:"D8081"` // [GRM, ...] - https://simplychocolate.fmcgproducts.dk/fmcg/pa/simplychocolate/pa.nsf/keyword.xsp?id=MeasurementUnitCodeList.da
	PalletHeight         int    `json:"D8083"` //
	PalletHeightUoM      string `json:"D8084"` // [MMT, ...] - https://simplychocolate.fmcgproducts.dk/fmcg/pa/simplychocolate/pa.nsf/keyword.xsp?id=MeasurementUnitCodeList.da
	PalletDepth          int    `json:"D8085"` //
	PalletDepthUoM       string `json:"D8086"` // [MMT, ...] - https://simplychocolate.fmcgproducts.dk/fmcg/pa/simplychocolate/pa.nsf/keyword.xsp?id=MeasurementUnitCodeList.da
	PalletWidth          int    `json:"D8087"` //
	PalletWidthUoM       string `json:"D8088"` // [MMT, ...] - https://simplychocolate.fmcgproducts.dk/fmcg/pa/simplychocolate/pa.nsf/keyword.xsp?id=MeasurementUnitCodeList.da

	LayersPerPallet         int `json:"D8079"`
	PalletSendingUnitAmount int `json:"D8078"`
	PalletUnitsPerLayer     int `json:"D3438"`
}

type FMCGMixCaseContentBaseItem struct {
	UnitsPerCase float64
	UnitGTINItem string
}

func FMCGApiPostMixCase(mixCaseInfo FmcgProductBodyMixCase, mixCaseContent []FMCGMixCaseContentBaseItem, count int) error {
	resp, err := GetFMCGApiBaseClient().
		//DevMode().
		R().
		EnableDump().
		SetSuccessResult(FmcgProductPostResult{}).
		SetErrorResult(FMCGSendToGS1PostResult{}).
		SetBody(mixCaseInfo).
		Post("")
	if err != nil {
		return err
	}

	if resp.IsErrorState() {
		if resp.StatusCode == 400 {
			response := resp.ErrorResult().(*FMCGSendToGS1PostResult)
			fmt.Printf("[POSTMXCS400]: status code 400. Errorcode: %v\n", response.Result)

			return nil
		} else {
			fmt.Printf("[POSTMXCSOTHR]: resp is err statusCode: %v. Dump: %v\n", resp.StatusCode, resp.Dump())
		}
		return resp.Err
	}

	// Check for validation errors before the body is sent
	headerResponse := resp.SuccessResult().(*FmcgProductPostResult)
	if len(headerResponse.ValidationErrors) != 0 {
		for _, validationError := range headerResponse.ValidationErrors {
			err = teams_notifier.SendValidationErrorToTeams(mixCaseInfo.GTIN,
				validationError.FieldId,
				validationError.FieldLabel,
				validationError.Message,
				validationError.MessageType,
				fmt.Sprintf("%v", mixCaseInfo),
			)
			if err != nil {
				return err
			}
		}
	}

	// Iterate the mixCaseContent and create a map with all the base items
	body := map[string]interface{}{
		"D8165": mixCaseInfo.GTIN,
		"D8255": mixCaseInfo.TargetMarketCode,
	}
	for i, baseItem := range mixCaseContent {
		body[fmt.Sprintf("D8270_%v", i+1)] = fmt.Sprintf("%v", baseItem.UnitsPerCase)
		body[fmt.Sprintf("D8249_%v", i+1)] = baseItem.UnitGTINItem
	}

	response, err := FMCGApiPostMixCaseContent(body, 0)
	if err != nil {
		return err
	}

	if len(response.ValidationErrors) != 0 {
		for _, validationError := range response.ValidationErrors {
			err = teams_notifier.SendValidationErrorToTeams(
				mixCaseInfo.ItemCode,
				mixCaseInfo.GTIN,
				validationError.FieldId,
				validationError.FieldLabel,
				validationError.Message,
				validationError.MessageType,
			)
			if err != nil {
				return err
			}
		}
	} else {
		var SendToGS1Data FMCGIdentifierData
		SendToGS1Data.GTIN = mixCaseInfo.GTIN
		SendToGS1Data.TargetMarketCode = mixCaseInfo.TargetMarketCode

		err = SendGTINToGS1(SendToGS1Data, mixCaseInfo.ItemCode)
		if err != nil {
			return fmt.Errorf("error sending product with GTIN:%v to GS1. \nError:%v", SendToGS1Data.GTIN, err)
		}
	}

	return nil
}

func FMCGApiPostMixCaseContent(body map[string]interface{}, count int) (*FmcgProductPostResult, error) {

	resp, err := GetFMCGApiBaseClient().
		//DevMode().
		R().
		EnableDump().
		SetSuccessResult(FmcgProductPostResult{}).
		SetErrorResult(FMCGSendToGS1PostResult{}).
		SetBody(body).
		Post("")
	if err != nil {
		return nil, err
	}

	if resp.IsErrorState() {
		if resp.StatusCode == 400 {
			response := resp.ErrorResult().(*FMCGSendToGS1PostResult)
			fmt.Printf("[234mkld12]: status code 400. Errorcode: %v", response.Result)
			return nil, nil

		} else {
			fmt.Printf("[1234kmlsd12]: resp is err statusCode: %v. Dump: %v\n", resp.StatusCode, resp.Dump())
		}
		return nil, fmt.Errorf("error posting mixCaseContentInfo: %v", body["D8165"])
	}

	response := resp.SuccessResult().(*FmcgProductPostResult)

	return response, nil
}
