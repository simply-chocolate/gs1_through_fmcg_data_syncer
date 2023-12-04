package fmcg_api_wrapper

import "fmt"

type FMCGProductStatusResult struct {
	StatusCode        int      `json:"status"`
	ProductId         string   `json:"productId"`
	LastModified      string   `json:"lastModified"`
	FmcgProductStatus string   `json:"fmcgProductsStatus"`
	Gs1Status         string   `json:"gs1Status"`
	Gs1Response       []string `json:"gs1Response"`
	Gs1Warnings       []string `json:"gs1Warnings"`
	ValidationErrors  []struct {
		FieldId     string   `json:"fieldId"`
		FieldLabel  string   `json:"fieldLabel"`
		Message     string   `json:"message"`
		MessageType string   `json:"messageType"`
		RequiredBy  []string `json:"requiredBy"`
	} `json:"validationErrors"`
	SendStatusList []struct {
		ProductId  string `json:"productId"`
		SendStatus string `json:"sendStatus"`
	} `json:"sendStatusList"`
}

type FMCTProductStatusReturn struct {
	Body *FMCGProductStatusResult
}

func FMCGApiGetProductStatus(FMCGIdentifierData FMCGIdentifierData, count int) (FMCTProductStatusReturn, error) {
	resp, err := GetFMCGApiBaseClient().
		//DevMode().
		R().
		EnableDump().
		SetSuccessResult(FMCGProductStatusResult{}).
		Get(fmt.Sprintf("/status/%v.%v", FMCGIdentifierData.GTIN, FMCGIdentifierData.TargetMarketCode))
	if err != nil {
		return FMCTProductStatusReturn{}, err
	}

	if resp.IsErrorState() {
		if resp.StatusCode == 404 {
			return FMCTProductStatusReturn{
				Body: &FMCGProductStatusResult{
					FmcgProductStatus: "NOT_FOUND",
				},
			}, nil
		}
		fmt.Printf("resp is err statusCode: %v. Dump: %v\n", resp.StatusCode, resp.Dump())
		return FMCTProductStatusReturn{}, resp.Err
	}

	return FMCTProductStatusReturn{
		Body: resp.SuccessResult().(*FMCGProductStatusResult),
	}, nil

}
