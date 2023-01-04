package sap_api_wrapper

import (
	"fmt"
)

type SapApiGetItemsDataResults struct {
	Value    []SapApiItemsData `json:"value"`
	NextLink string            `json:"odata.nextLink"`
}

/*
type ItemUnitOfMeasurement struct {
	UoMType    string `json:"UoMType"`
	UoMEntry   int    `json:"UoMEntry"`
	Length     int    `json:"Length1"`
	LengthUnit int    `json:"Length1Unit"` // For dimensions [1 = mm?, 2 = cm, 3 = m?]
	Width      int    `json:"Width1"`
	WidthUnit  int    `json:"Width1Unit"` // For dimensions [1 = mm?, 2 = cm, 3 = m?]
	Height     int    `json:"Height1"`
	HeightUnit int    `json:"Height1Unit"` // For dimensions [1 = mm?, 2 = cm, 3 = m?]
	Weight     int    `json:"Weight1"`
	WeightUnit int    `json:"Weight1Unit"` // For weight [1 = ?, 2 = g?, 3 = kg]
	// Maybe the units is 1 list, so 2 = cm and 3 = kg no matter what type of input?
}
*/

type SapApiItemsData struct {
	// General Information
	ItemBarCodeCollection []struct {
		Barcode  string `json:"Barcode"`
		UoMEntry int    `json:"UoMEntry"`
	} `json:"ItemBarCodeCollection"`
	ItemCode                      string `json:"ItemCode"`
	ItemNameDA                    string `json:"ItemName"`
	FunctionalName                string `json:"U_BOYX_Varebeskrivelse"` // TODO: Create a new field for this eventually
	ShelfLifeFromArrivalInDays    string `json:"U_BOYX_Holdbarhed_Kunde"`
	ShelfLifeFromProductionInDays int    `json:"U_BOYX_Holdbarhed"`
	AvailabilityDateTime          string `json:"U_CCF_LaunchDate"`
	EffectiveDateTime             string `json:"U_CCF_GS1_Ajour"`
	BrandName                     string `json:"U_BOYX_varemrk"`

	// TODO: Add "Økologimærke felt i SAP " https://simplychocolate.fmcgproducts.dk/fmcg/pa/simplychocolate/pa.nsf/keyword.xsp?id=OrganicTradeItemCodeList.da

	// Logistical Data // TODO: Use the actual UOM module in SAP for these instead of the homemade one from Bitpeople.
	//ItemUnitOfMeasurementCollection []ItemUnitOfMeasurement `json:"ItemUnitOfMeasurementCollection"`
	// We want to use the above instead, but the data is not in SAP yet, so we start with this.
	BaseUnitWidth       int `json:"U_BOYX_Bredde_e"`
	BaseUnitHeight      int `json:"U_BOYX_hoojde_e"`
	BaseUnitDepth       int `json:"U_BOYX_dybde_e"`
	BaseUnitNetWeight   int `json:"U_BOYX_netto_e"`
	BaseUnitGrossWeight int `json:"U_BOYX_brutto_e"`

	// Allergen containment information
	ContainmentLevelGluten    string `json:"U_BOYX_gluten"`
	ContainmentLevelCrustacea string `json:"U_BOYX_Krebsdyr"`
	ContainmentLevelEgg       string `json:"U_BOYX_aag"`
	ContainmentLevelFish      string `json:"U_BOYX_fisk"`
	ContainmentLevelPeanut    string `json:"U_BOYX_JN"`
	ContainmentLevelSoy       string `json:"U_BOYX_soja"`
	ContainmentLevelMilk      string `json:"U_BOYX_ML"`
	ContainmentLevelAlmonds   string `json:"U_BOYX_mandel"`
	ContainmentLevelHazelnut  string `json:"U_BOYX_hassel"`
	ContainmentLevelWalnut    string `json:"U_BOYX_val"`
	ContainmentLevelCashew    string `json:"U_BOYX_Cashe"`
	ContainmentLevelPecan     string `json:"U_BOYX_Pekan"`
	ContainmentLevelBrazilNut string `json:"U_BOYX_peka"`
	ContainmentLevelPistachio string `json:"U_BOYX_Pistacie"`

	// Claims
	GlutenFree  string `json:"U_BOYX_Gluten1"`
	LactoseFree string `json:"U_BOYX_Lactose"`
	Vegetarian  string `json:"U_BOYX_Vegetar"`
	Vegan       string `json:"U_BOYX_Vegan"`
	CowFree     string `json:"U_BOYX_Okse"`
	PigFree     string `json:"U_BOYX_gris"`
	GMOFree     string `json:"U_BOYX_GMO"`

	// Nutritional Information
	EnergyInkJ                    string `json:"U_BOYX_Energi"`
	EnergyInKcal                  string `json:"U_BOYX_Energik"`
	NutritionalFatValue           string `json:"U_BOYX_fedt"`
	NutritionalFattyAcidsValue    string `json:"U_BOYX_fedtsyre"`
	NutritionalCarboHydratesValue string `json:"U_BOYX_Kulhydrat"`
	NutritionalSugarValue         string `json:"U_BOYX_sukkerarter"`
	NutritionalProteinValue       string `json:"U_BOYX_Protein"`
	NutritionalSaltValue          string `json:"U_BOYX_salt"`
}

type SapApiGetItemsDataReturn struct {
	Body *SapApiGetItemsDataResults
}

func SapApiGetItemsData(params SapApiQueryParams) (SapApiGetItemsDataReturn, error) {
	client, err := GetSapApiAuthClient()
	if err != nil {
		fmt.Println("Error getting an authenticaed client")
		return SapApiGetItemsDataReturn{}, err
	}

	resp, err := client.
		//DevMode().
		R().
		SetResult(SapApiGetItemsDataResults{}).
		SetQueryParams(params.AsReqParams()).
		Get("Items")
	if err != nil {
		fmt.Println(err)
		return SapApiGetItemsDataReturn{}, err
	}

	return SapApiGetItemsDataReturn{
		Body: resp.Result().(*SapApiGetItemsDataResults),
	}, nil

}

func SapApiGetItemsData_AllPages(params SapApiQueryParams) (SapApiGetItemsDataReturn, error) {
	res := SapApiGetItemsDataResults{}
	for page := 0; ; page++ {
		params.Skip = page * 20

		getItemsRes, err := SapApiGetItemsData(params)
		if err != nil {
			return SapApiGetItemsDataReturn{}, err
		}

		res.Value = append(res.Value, getItemsRes.Body.Value...)

		if getItemsRes.Body.NextLink == "" {
			break
		}
	}

	return SapApiGetItemsDataReturn{
		Body: &res,
	}, nil
}
