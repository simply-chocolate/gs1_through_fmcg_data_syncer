package fmcg_api_wrapper

import (
	"fmt"
	"gs1_syncer/teams_notifier"
)

type FmcgProductBodyBaseItem struct {
	GTIN     string `json:"D8165"` // Barcode with 0 in front
	DataType string `json:"D8164"` // [CORRECT, ...] // https://simplychocolate.fmcgproducts.dk/fmcg/pa/simplychocolate/pa.nsf/keyword.xsp?id=DocumentCommandHeaderCode.da
	//	General Information
	DataCarrierTypeCode                 string `json:"D8208"` // [EAN_13, ...] - https://simplychocolate.fmcgproducts.dk/fmcg/pa/simplychocolate/pa.nsf/keyword.xsp?id=DataCarrierTypeCodeList.da
	BrandName                           string `json:"D8211"`
	CountryOfOrigin                     string `json:"D8219"`   // [208, ...] - https://simplychocolate.fmcgproducts.dk/fmcg/pa/simplychocolate/pa.nsf/keyword.xsp?id=CountryCodeList.da
	ManufacturerGLN                     string `json:"D8242_1"` //
	BrandOwnerGLN                       string `json:"D8346"`   //
	GPCCategoryCode                     string `json:"D8245"`   // 10000045
	ImportClassificationValue           string `json:"D8253_1"` // Default 18069019 for Chocolate
	ImportClassificationType            string `json:"D8254_1"` // [INSTRASTAT, CUSTOMS_TARIFF_NUMBER...] - https://simplychocolate.fmcgproducts.dk/fmcg/pa/simplychocolate/pa.nsf/keyword.xsp?id=ImportClassificationTypeCodeList.da
	TargetMarketCode                    string `json:"D8255"`   // Default 208 for DK
	ItemCode                            string `json:"D8256"`
	ItemNameDA                          string `json:"D8258_1"` // The name of the item in the language specified in D8259_1
	ItemNameLanguageCodeDA              string `json:"D8259_1"` // [da, ...] - https://simplychocolate.fmcgproducts.dk/fmcg/pa/simplychocolate/pa.nsf/keyword.xsp?id=LanguageCodeList.da
	FunctionalProductNameDA             string `json:"D8313_1"`
	FunctionalProductNameLanguageCodeDA string `json:"D8121_1"`     // [da, ...] - https://simplychocolate.fmcgproducts.dk/fmcg/pa/simplychocolate/pa.nsf/keyword.xsp?id=LanguageCodeList.da
	RegulatedProductNameDA              string `json:"D8146_1"`     // Chokolade
	RegulatedProductNameLanguageCodeDA  string `json:"D8146Attr_1"` // [da, ...] - https://simplychocolate.fmcgproducts.dk/fmcg/pa/simplychocolate/pa.nsf/keyword.xsp?id=LanguageCodeList.da

	IsOrderingUnit                bool   `json:"D8271"`   // [TRUE, FALSE] (True for cases and displays, False for BASE_UNIT)
	UnitOfMeasure                 string `json:"D8276"`   // [BASE_UNIT_OR_EACH, CASE, PALLET, DISPLAY_SHIPPER] - https://simplychocolate.fmcgproducts.dk/fmcg/pa/simplychocolate/pa.nsf/keyword.xsp?id=TradeItemUnitDescriptorCodeList.da
	ShelfLifeFromArrivalInDays    int    `json:"D8283"`   //
	ShelfLifeFromProductionInDays int    `json:"D8284"`   //
	MaximumStorageTemp            int    `json:"D3599_1"` //
	MinimumStorageTemp            int    `json:"D3608_1"` //
	TemperatureType               string `json:"D3614_1"` // [STORAGE_HANDLING, ...] - https://simplychocolate.fmcgproducts.dk/fmcg/pa/simplychocolate/pa.nsf/keyword.xsp?id=TemperatureQualifierCodeList.da
	TemperatureOUM                string `json:"D8374_1"` // [CEL, ...] - https://simplychocolate.fmcgproducts.dk/fmcg/pa/simplychocolate/pa.nsf/keyword.xsp?id=TemperatureMeasurementUnitCodeList.da

	IsQuantityOrPriceVarying    bool   `json:"D8297"` // [TRUE, FALSE]
	DangerousContent            string `json:"D8030"` // [NOT_APPLICABLE, TRUE, FALSE, UNSPECIFIED] - https://simplychocolate.fmcgproducts.dk/fmcg/pa/simplychocolate/pa.nsf/keyword.xsp?id=NonBinaryLogicEnumerationCodeList.da
	RelevantForPriceComparison  string `json:"D8019"` // [NOT_APPLICABLE, TRUE, FALSE, UNSPECIFIED] - https://simplychocolate.fmcgproducts.dk/fmcg/pa/simplychocolate/pa.nsf/keyword.xsp?id=NonBinaryLogicEnumerationCodeList.da
	IsConsumerUnit              bool   `json:"D8216"` // [TRUE, FALSE]
	IsShippingUnit              bool   `json:"D8236"` // [TRUE, FALSE]
	IsPackagingMarkedReturnable bool   `json:"D8311"` // [TRUE, FALSE]
	OrganicTradeItemCodeList    int    `json:"D0798"` // [1, 5] - https://simplychocolate.fmcgproducts.dk/fmcg/pa/simplychocolate/pa.nsf/keyword.xsp?id=OrganicTradeItemCodeList.da
	EffectiveDateTime           string `json:"D8286"` // DateTime
	StartAvailabilityDateTime   string `json:"D8314"` // DateTime

	//	Logistical Information
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

	//	Allergens
	//	List of Allergens and their codes - https://simplychocolate.fmcgproducts.dk/fmcg/pa/simplychocolate/pa.nsf/keyword.xsp?id=AllergenTypeCodeList.da
	// 	List of Containment types - https://simplychocolate.fmcgproducts.dk/fmcg/pa/simplychocolate/pa.nsf/keyword.xsp?id=AllergenTypeCodeList.da
	AllergenGluten                           string `json:"D8166_1"`  // AW
	ContainmentLevelGluten                   string `json:"D8170_1"`  // [FREE_FROM, CONTAINS, MAY_CONTAIN]
	AllergenCrustacea                        string `json:"D8166_2"`  // AC (Krebsdyr)
	ContainmentLevelCrustacea                string `json:"D8170_2"`  // [FREE_FROM, CONTAINS, MAY_CONTAIN]
	AllergenEgg                              string `json:"D8166_3"`  // AE
	ContainmentLevelEgg                      string `json:"D8170_3"`  // [FREE_FROM, CONTAINS, MAY_CONTAIN]
	AllergenFish                             string `json:"D8166_4"`  // AF
	ContainmentLevelFish                     string `json:"D8170_4"`  // [FREE_FROM, CONTAINS, MAY_CONTAIN]
	AllergenPeanut                           string `json:"D8166_5"`  // AP
	ContainmentLevelPeanut                   string `json:"D8170_5"`  // [FREE_FROM, CONTAINS, MAY_CONTAIN]
	AllergenSoy                              string `json:"D8166_6"`  // AY
	ContainmentLevelSoy                      string `json:"D8170_6"`  // [FREE_FROM, CONTAINS, MAY_CONTAIN]
	AllergenMilk                             string `json:"D8166_7"`  // AM
	ContainmentLevelMilk                     string `json:"D8170_7"`  // [FREE_FROM, CONTAINS, MAY_CONTAIN]
	AllergenAlmonds                          string `json:"D8166_8"`  // SA
	ContainmentLevelAlmonds                  string `json:"D8170_8"`  // [FREE_FROM, CONTAINS, MAY_CONTAIN]
	AllergenHazelnut                         string `json:"D8166_9"`  // SH
	ContainmentLevelHazelnut                 string `json:"D8170_9"`  // [FREE_FROM, CONTAINS, MAY_CONTAIN]
	AllergenWalnut                           string `json:"D8166_10"` // SW
	ContainmentLevelWalnut                   string `json:"D8170_10"` // [FREE_FROM, CONTAINS, MAY_CONTAIN]
	AllergenCashew                           string `json:"D8166_11"` // SC
	ContainmentLevelCashew                   string `json:"D8170_11"` // [FREE_FROM, CONTAINS, MAY_CONTAIN]
	AllergenPecan                            string `json:"D8166_12"` // SP
	ContainmentLevelPecan                    string `json:"D8170_12"` // [FREE_FROM, CONTAINS, MAY_CONTAIN]
	AllergenBrazilNut                        string `json:"D8166_13"` // SR (Paranød)
	ContainmentLevelBrazilNut                string `json:"D8170_13"` // [FREE_FROM, CONTAINS, MAY_CONTAIN]
	AllergenPistachio                        string `json:"D8166_14"` // ST
	ContainmentLevelPistachio                string `json:"D8170_14"` // [FREE_FROM, CONTAINS, MAY_CONTAIN]
	AllergenQueenslandNut                    string `json:"D8166_15"` // SQ
	ContainmentLevelQueenslandNut            string `json:"D8170_15"` // [FREE_FROM, CONTAINS, MAY_CONTAIN]
	AllergenCelery                           string `json:"D8166_16"` // BC
	ContainmentLevelCelery                   string `json:"D8170_16"` // [FREE_FROM, CONTAINS, MAY_CONTAIN]
	AllergenMustard                          string `json:"D8166_17"` // BM
	ContainmentLevelMustard                  string `json:"D8170_17"` // [FREE_FROM, CONTAINS, MAY_CONTAIN]
	AllergenSulfurDioxideAndSulfites         string `json:"D8166_18"` // AU
	ContainmentLevelSulfurDioxideAndSulfites string `json:"D8170_18"` // [FREE_FROM, CONTAINS, MAY_CONTAIN]
	AllergenSesameSeeds                      string `json:"D8166_19"` // AS
	ContainmentLevelSesameSeeds              string `json:"D8170_19"` // [FREE_FROM, CONTAINS, MAY_CONTAIN]
	AllergenLupine                           string `json:"D8166_20"` // NL
	ContainmentLevelLupine                   string `json:"D8170_20"` // [FREE_FROM, CONTAINS, MAY_CONTAIN]
	AllergenMollusks                         string `json:"D8166_21"` // UM (Bløddyr)
	ContainmentLevelMollusks                 string `json:"D8170_21"` // [FREE_FROM, CONTAINS, MAY_CONTAIN]
	AllergenNuts                             string `json:"D8166_22"` // AN
	ContainmentLevelNuts                     string `json:"D8170_22"` // [FREE_FROM, CONTAINS, MAY_CONTAIN]

	// 	Nutritional Information
	EnergyInkJ                        string `json:"D8175-UNPREPARED"`   //
	EnergyInKcal                      string `json:"D8171-UNPREPARED"`   //
	EnergyInKcalPrecision             string `json:"D8172-UNPREPARED"`   //	[APPROXIMATELY, ...] - https://simplychocolate.fmcgproducts.dk/fmcg/pa/simplychocolate/pa.nsf/keyword.xsp?id=MeasurementPrecisionCodeList.da
	PreparationState                  string `json:"D8173"`              // [UNPREPARED, PREPARED]
	NutritionalReferenceValue         int    `json:"D8187-UNPREPARED"`   // 100
	NutritionalReferenceUOM           string `json:"D8188-UNPREPARED"`   // [GRM, ...] https://simplychocolate.fmcgproducts.dk/fmcg/pa/simplychocolate/pa.nsf/keyword.xsp?id=MeasurementUnitCodeList.da
	NutritionalFat                    string `json:"D8181-UNPREPARED_1"` // [FAT, ...] - https://simplychocolate.fmcgproducts.dk/fmcg/pa/simplychocolate/pa.nsf/keyword.xsp?id=NutrientTypeCodeList.da
	NutritionalFatPrecision           string `json:"D8182-UNPREPARED_1"` //	[APPROXIMATELY, ...] - https://simplychocolate.fmcgproducts.dk/fmcg/pa/simplychocolate/pa.nsf/keyword.xsp?id=MeasurementPrecisionCodeList.da
	NutritionalFatValue               string `json:"D8183-UNPREPARED_1"` //
	NutritionalFatUOM                 string `json:"D8184-UNPREPARED_1"` //	[GRM, ...] - https://simplychocolate.fmcgproducts.dk/fmcg/pa/simplychocolate/pa.nsf/keyword.xsp?id=MeasurementUnitCodeList.da
	NutritionalFattyAcids             string `json:"D8181-UNPREPARED_2"` // [FASAT, ...] - https://simplychocolate.fmcgproducts.dk/fmcg/pa/simplychocolate/pa.nsf/keyword.xsp?id=NutrientTypeCodeList.da
	NutritionalFattyAcidsPrecision    string `json:"D8182-UNPREPARED_2"` //	[APPROXIMATELY, ...] - https://simplychocolate.fmcgproducts.dk/fmcg/pa/simplychocolate/pa.nsf/keyword.xsp?id=MeasurementPrecisionCodeList.da
	NutritionalFattyAcidsValue        string `json:"D8183-UNPREPARED_2"` //
	NutritionalFattyUOM               string `json:"D8184-UNPREPARED_2"` // [GRM, ...] - https://simplychocolate.fmcgproducts.dk/fmcg/pa/simplychocolate/pa.nsf/keyword.xsp?id=MeasurementUnitCodeList.da
	NutritionalCarboHydrates          string `json:"D8181-UNPREPARED_3"` // [CHOAVL, ...] - https://simplychocolate.fmcgproducts.dk/fmcg/pa/simplychocolate/pa.nsf/keyword.xsp?id=NutrientTypeCodeList.da
	NutritionalCarboHydratesPrecision string `json:"D8182-UNPREPARED_3"` // [APPROXIMATELY, ...] - https://simplychocolate.fmcgproducts.dk/fmcg/pa/simplychocolate/pa.nsf/keyword.xsp?id=MeasurementPrecisionCodeList.da
	NutritionalCarboHydratesValue     string `json:"D8183-UNPREPARED_3"` //
	NutritionalCarboHydratesUOM       string `json:"D8184-UNPREPARED_3"` // [GRM, ...] - https://simplychocolate.fmcgproducts.dk/fmcg/pa/simplychocolate/pa.nsf/keyword.xsp?id=MeasurementUnitCodeList.da
	NutritionalSugar                  string `json:"D8181-UNPREPARED_4"` // [SUGAR-, ...] - https://simplychocolate.fmcgproducts.dk/fmcg/pa/simplychocolate/pa.nsf/keyword.xsp?id=NutrientTypeCodeList.da
	NutritionalSugarPrecision         string `json:"D8182-UNPREPARED_4"` // [APPROXIMATELY, ...] - https://simplychocolate.fmcgproducts.dk/fmcg/pa/simplychocolate/pa.nsf/keyword.xsp?id=MeasurementPrecisionCodeList.da
	NutritionalSugarValue             string `json:"D8183-UNPREPARED_4"` //
	NutritionalSugarUOM               string `json:"D8184-UNPREPARED_4"` // [GRM, ...] - https://simplychocolate.fmcgproducts.dk/fmcg/pa/simplychocolate/pa.nsf/keyword.xsp?id=MeasurementUnitCodeList.da
	NutritionalProtein                string `json:"D8181-UNPREPARED_5"` // [PRO-, ...] - https://simplychocolate.fmcgproducts.dk/fmcg/pa/simplychocolate/pa.nsf/keyword.xsp?id=NutrientTypeCodeList.da
	NutritionalProteinPrecision       string `json:"D8182-UNPREPARED_5"` // [APPROXIMATELY, ...] - https://simplychocolate.fmcgproducts.dk/fmcg/pa/simplychocolate/pa.nsf/keyword.xsp?id=MeasurementPrecisionCodeList.da
	NutritionalProteinValue           string `json:"D8183-UNPREPARED_5"` //
	NutritionalProteinUOM             string `json:"D8184-UNPREPARED_5"` // [GRM, ...] - https://simplychocolate.fmcgproducts.dk/fmcg/pa/simplychocolate/pa.nsf/keyword.xsp?id=MeasurementUnitCodeList.da
	NutritionalSalt                   string `json:"D8181-UNPREPARED_6"` // [SALTEQ, NACL, ...] - https://simplychocolate.fmcgproducts.dk/fmcg/pa/simplychocolate/pa.nsf/keyword.xsp?id=NutrientTypeCodeList.da
	NutritionalSaltPrecision          string `json:"D8182-UNPREPARED_6"` // [APPROXIMATELY, ...] - https://simplychocolate.fmcgproducts.dk/fmcg/pa/simplychocolate/pa.nsf/keyword.xsp?id=MeasurementPrecisionCodeList.da
	NutritionalSaltValue              string `json:"D8183-UNPREPARED_6"` //
	NutritionalSaltUOM                string `json:"D8184-UNPREPARED_6"` // [GRM, ...] - https://simplychocolate.fmcgproducts.dk/fmcg/pa/simplychocolate/pa.nsf/keyword.xsp?id=MeasurementUnitCodeList.da

	// SAP now has multiple languages, but we only use the "current list of ingredients" field, cause that is the one we know matches the physical product.
	ListOfIngredientsDA             string `json:"D8179_1"` //
	ListOfIngredientsLanguageCodeDA string `json:"D8180_1"` // da (must be non-capitalized)

	//	General Information
	StorageInformationLanguageCode1 string `json:"D0352Attr_1"` // Language code of the language the StorageInformation in D0352_1 is written in.
	StorageInformation              string `json:"D0352_1"`     // Default is "Tørt og ved max 21 grader"

	/* This might not be necessary as we can just upload the pictures directly into GS1.
	// Image Information
	ProductImageUrl        string `json:"D8350_1"` // Image of Product
	ProductImageType       string `json:"D8349_1"` // [PRODUCT_IMAGE, ...] - https://simplychocolate.fmcgproducts.dk/fmcg/pa/simplychocolate/pa.nsf/keyword.xsp?id=ReferencedFileTypeCodeList.da
	ProductImageFileName   string `json:"D8388_1"` // GTIN + Z1C1
	ProductImageFileFormat string `json:"D8387_1"` // png
	*/
}

type FmcgProductPostResult struct {
	StatusCode        int    `json:"status"`
	ProductId         string `json:"productId"`
	UnitType          string `json:"unitType"`
	LastModified      string `json:"lastModified"`
	FmcgProductStatus string `json:"fmcgProductsStatus"`
	Gs1Status         string `json:"gs1Status"`
	ValidationErrors  []struct {
		FieldId     string   `json:"fieldId"`
		FieldLabel  string   `json:"fieldLabel"`
		Message     string   `json:"message"`
		MessageType string   `json:"messageType"`
		RequiredBy  []string `json:"requiredBy"`
	} `json:"validationErrors"`
}

func FMCGApiPostBaseItem(ItemInfo FmcgProductBodyBaseItem) error {

	resp, err := GetFMCGApiBaseClient().
		//DevMode().
		R().
		EnableDump().
		SetSuccessResult(FmcgProductPostResult{}).
		SetErrorResult(FMCGSendToGS1PostResult{}).
		SetBody(ItemInfo).
		Post("")
	if err != nil {
		return err
	}

	if resp.IsErrorState() {
		if resp.StatusCode == 400 {
			response := resp.ErrorResult().(*FMCGSendToGS1PostResult)
			fmt.Printf("[POSTBU400]: status code 400. Errorcode: %v\n", response.Result)

			return nil
		} else {
			fmt.Printf("[POSTBUOTHR]: resp is err statusCode: %v. Dump: %v\n", resp.StatusCode, resp.Dump())
		}
		return resp.Err
	}

	response := resp.SuccessResult().(*FmcgProductPostResult)

	if len(response.ValidationErrors) != 0 {
		for _, validationError := range response.ValidationErrors {
			if validationError.FieldId == "D8271" {
				continue
			}
			err = teams_notifier.SendValidationErrorToTeams(ItemInfo.ItemCode,
				ItemInfo.GTIN,
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
		SendToGS1Data.GTIN = ItemInfo.GTIN
		SendToGS1Data.TargetMarketCode = ItemInfo.TargetMarketCode

		err = SendGTINToGS1(SendToGS1Data, ItemInfo.ItemCode)
		if err != nil {
			return fmt.Errorf("error sending product with GTIN:%v to GS1. \nError:%v", SendToGS1Data.GTIN, err)
		}
	}

	return nil
}
