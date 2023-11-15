package sap_api_wrapper

import (
	"fmt"
)

type SapApiGetItemsDataResults struct {
	Value    []SapApiItemsData `json:"value"`
	NextLink string            `json:"odata.nextLink"`
}

type SapApiItemsData struct {
	// Update Information
	UpdateDate string `json:"UpdateDate"`
	UpdateTime string `json:"UpdateTime"`
	GS1Status  string `json:"U_CCF_GS1_Status"`

	TypeOfProduct string `json:"U_CCF_Type"`    // If this is Equal to "Kampagne" then it should not have a BaseUnit ItemCode
	UoMGroupEntry int    `json:"UoMGroupEntry"` // If this is Equal to 42 then it's a Campaign Display
	// General Information
	ItemBarCodeCollection []struct {
		Barcode  string `json:"Barcode"`
		UoMEntry int    `json:"UoMEntry"`
	} `json:"ItemBarCodeCollection"`
	ItemCode   string `json:"ItemCode"`
	ItemNameDA string `json:"ItemName"`
	/*FunctionalName                string `json:"U_BOYX_Varebeskrivelse"`*/
	FunctionalProductNameDA       string `json:"U_CCF_Functional_Name"`
	ShelfLifeFromArrivalInDays    int    `json:"U_CCF_ShelfLifeArrival"`
	ShelfLifeFromProductionInDays int    `json:"U_BOYX_Holdbarhed"`
	MinimumStorageTemp            int    `json:"U_BOYX_Minimum"`
	MaximumStorageTemp            int    `json:"U_BOYX_Maximum"`
	AvailabilityDateTime          string `json:"U_CCF_LaunchDate"`
	EffectiveDateTime             string `json:"U_CCF_GS1_Ajour"`
	BrandName                     string `json:"U_BOYX_varemrk"`
	OrganicTradeItemCodeList      int    `json:"U_CCF_OrganicCode"`      // https://simplychocolate.fmcgproducts.dk/fmcg/pa/simplychocolate/pa.nsf/keyword.xsp?id=OrganicTradeItemCodeList.da
	ManufacturerGLN               string `json:"U_CCF_Manufacturer_GLN"` // "D8242" https://simplychocolate.fmcgproducts.dk/fmcg/pa/simplychocolate/pa.nsf/field.xsp?documentId=516f5dd5e012c0f7c12589090055223c&action=openDocument
	BrandOwnerGLN                 string `json:"U_CCF_BrandOwner_GLN"`   // "D8346" https://simplychocolate.fmcgproducts.dk/fmcg/pa/simplychocolate/pa.nsf/field.xsp?documentId=d23d8eda0e728e26c1258909005522e1&action=openDocument
	ProductImageUrl               string `json:"U_CCF_GS1_Z1C1_URL"`
	PackagingType                 string `json:"U_CCF_Packaging_Type"`
	IsSalesReady                  string `json:"U_CCF_IsCaseSalesReady"`
	// Logistical Data

	// BaseUnit
	BaseUnitWidth       int `json:"U_BOYX_Bredde_e"`
	BaseUnitHeight      int `json:"U_BOYX_hoojde_e"`
	BaseUnitDepth       int `json:"U_BOYX_dybde_e"`
	BaseUnitNetWeight   int `json:"U_BOYX_netto_e"`
	BaseUnitGrossWeight int `json:"U_BOYX_brutto_e"`

	// Case
	CaseWidth       int `json:"U_BOYX_bredde_k"`
	CaseHeight      int `json:"U_BOYX_hoojde_k"`
	CaseDepth       int `json:"U_BOYX_dybde_k"`
	CaseNetWeight   int `json:"U_BOYX_netto_k"`
	CaseGrossWeight int `json:"U_BOYX_brutto_k"`
	UnitsPerCase    int `json:"U_BOYX_antal_k"`

	// Pallet
	PalletWidth             int `json:"U_BOYX_bredde_p"`
	PalletHeight            int `json:"U_BOYX_hoojde_p"`
	PalletDepth             int `json:"U_BOYX_dybde_p"`
	PalletNetWeight         int `json:"U_BOYX_netto_p"`
	PalletGrossWeight       int `json:"U_BOYX_brutto_p"`
	LayersPerPallet         int `json:"U_BOYX_kolli_p"`
	PalletUnitsPerLayer     int `json:"U_BOYX_antalpalle_p"`
	PalletSendingUnitAmount int `json:"U_BOYX_kollipalle_p"`

	// Allergen containment information
	ContainmentLevelGluten                   string `json:"U_BOYX_gluten"`
	ContainmentLevelCrustacea                string `json:"U_BOYX_Krebsdyr"`
	ContainmentLevelEgg                      string `json:"U_BOYX_aag"`
	ContainmentLevelFish                     string `json:"U_BOYX_fisk"`
	ContainmentLevelPeanut                   string `json:"U_BOYX_JN"`
	ContainmentLevelSoy                      string `json:"U_BOYX_soja"`
	ContainmentLevelMilk                     string `json:"U_BOYX_ML"`
	ContainmentLevelAlmonds                  string `json:"U_BOYX_mandel"`
	ContainmentLevelHazelnut                 string `json:"U_BOYX_hassel"`
	ContainmentLevelWalnut                   string `json:"U_BOYX_val"`
	ContainmentLevelCashew                   string `json:"U_BOYX_Cashe"`
	ContainmentLevelPecan                    string `json:"U_BOYX_Pekan"`
	ContainmentLevelBrazilNut                string `json:"U_BOYX_peka"`
	ContainmentLevelPistachio                string `json:"U_BOYX_Pistacie"`
	ContainmentLevelQueenslandNut            string `json:"U_BOYX_Queensland"`
	ContainmentLevelCelery                   string `json:"U_BOYX_Selleri"`
	ContainmentLevelMustard                  string `json:"U_BOYX_Sennep"`
	ContainmentLevelSulfurDioxideAndSulfites string `json:"U_BOYX_Svovldioxid"`
	ContainmentLevelSesameSeeds              string `json:"U_BOYX_Sesam"`
	ContainmentLevelLupine                   string `json:"U_BOYX_Lupin"`
	ContainmentLevelMollusks                 string `json:"U_BOYX_BL"`

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
	ListOfIngredientsDA           string `json:"U_BOYX_varedel"`
}

type SapApiGetItemsDataReturn struct {
	Body *SapApiGetItemsDataResults
}

func SapApiGetItemsData(params SapApiQueryParams) (SapApiGetItemsDataReturn, error) {
	client, err := GetSapApiAuthClient()
	if err != nil {
		fmt.Println("Error getting an authenticated client")
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
