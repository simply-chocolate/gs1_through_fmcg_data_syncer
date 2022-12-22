package fmcg_api_wrapper

type FmcgProductBody struct {
	GTIN string `json:"D8165"` // Barcode with 0 in front

	// Allergens
	// List of Allergens and their codes https://simplychocolate.fmcgproducts.dk/fmcg/pa/simplychocolate/pa.nsf/keyword.xsp?id=AllergenTypeCodeList.da
	// List of Containment types https://simplychocolate.fmcgproducts.dk/fmcg/pa/simplychocolate/pa.nsf/keyword.xsp?id=AllergenTypeCodeList.da
	// [FREE_FROM, CONTAINS, MAY_CONTAIN]
	//  hasselnød, valnød, Cashew, pekan, paranød, pistacienød, queenslandnød, selleri, sennep, Svovldioxid, lupin, bløddyr

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
	AllergenWalnut                           string `json:"D8166_10"` // Sw
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

	// Nutritional Information
	EnergyInKcal              string `json:"D8171-UNPREPARED"` //
	PreparationState          string `json:"D8173"`            // [UNPREPARED, PREPARED]
	EnergyInkJ                string `json:"D8175-UNPREPARED"` //
	NutritionalReferenceValue int    `json:"D8187-UNPREPARED"` // 100 g

	// General Information
	GPCCategoryCode           string `json:"D8245"`   // 10000045
	ImportClassificationValue string `json:"D8253_1"` // Default 18069019 00 for Chocolate //TODO: Spørg FMCG indtil disse to felter
	ImportClassificationType  string `json:"D8254_1"` // [CUSTOMS_TARIFF_NUMBER, ...] - https://simplychocolate.fmcgproducts.dk/fmcg/pa/simplychocolate/pa.nsf/keyword.xsp?id=ImportClassificationTypeCodeList.da
	TargetMarketCode          string `json:"D8255"`   // Default 208 for DK
	ItemCode                  string `json:"D8256"`
	ItemNameDA                string `json:"D8258_1"` // The name of the item in the language specified in D8259_1
	ItemNameLanguageCode1     string `json:"D8259_1"` // Language code of the language the ItemName in D8258_1 is written in.

	DangerousContent string `json:"D8030"` // [NOT_APPLICABLE, TRUE, FALSE, UNSPECIFIED]- https://simplychocolate.fmcgproducts.dk/fmcg/pa/simplychocolate/pa.nsf/keyword.xsp?id=NonBinaryLogicEnumerationCodeList.da

	//Logistical Information
	Height    int    `json:"D8263"` //
	HeightUOM string `json:"D8264"` // [MMT, CMT, MT, ...] - https://simplychocolate.fmcgproducts.dk/fmcg/pa/simplychocolate/pa.nsf/keyword.xsp?id=MeasurementUnitCodeList.daWidth
	Depth     int    `json:"D8265"` //
	DepthUOM  string `json:"D8266"` // [MMT, CMT, MT, ...] - https://simplychocolate.fmcgproducts.dk/fmcg/pa/simplychocolate/pa.nsf/keyword.xsp?id=MeasurementUnitCodeList.daWidth
	Width     int    `json:"D8267"` //
	WidthUOM  string `json:"D8268"` // [MMT, CMT, MT, ...] - https://simplychocolate.fmcgproducts.dk/fmcg/pa/simplychocolate/pa.nsf/keyword.xsp?id=MeasurementUnitCodeList.daWidth

	IsOrderingUnit           bool   `json:"D8271"` // [TRUE, FALSE] (True for cases and displays, False for BASE_UNIT)
	UnitOfMeasure            string `json:"D8276"` // [BASE_UNIT_OR_EACH, CASE, PALLET, DISPLAY_SHIPPER] - https://simplychocolate.fmcgproducts.dk/fmcg/pa/simplychocolate/pa.nsf/keyword.xsp?id=TradeItemUnitDescriptorCodeList.da
	IsQuantityOrPriceVarying bool   `json:"D8297"` // [TRUE, FALSE]

	StorageInformationLanguageCode1 string `json:"D0352Attr_1"` // Language code of the language the StorageInformation in D0352_1 is written in.
	StorageInformation              string `json:"D0352_1"`     // Default is "Tørt og ved max 21 grader"
}
