package fmcg_api_wrapper

import "fmt"

type FMCGProductStatusResult struct {
	StatusCode        int    `json:"status"`
	ProductId         string `json:"productId"`
	LastModified      string `json:"lastModified"`
	FmcgProductStatus string `json:"fmcgProductsStatus"`
	Gs1Status         string `json:"gs1Status"`
	Gs1Response       string `json:"gs1Response"`
}

type FMCTProductStatusReturn struct {
	Body *FMCGProductStatusResult
}

func FMCGApiGetProductStatus(FMCGIdentifierData FMCGIdentifierData, count int) (FMCTProductStatusReturn, error) {
	resp, err := GetFMCGApiBaseClient().
		//DevMode().
		R().
		EnableDump().
		SetResult(FMCGProductStatusResult{}).
		Get(fmt.Sprintf("/status/%v.%v", FMCGIdentifierData.GTIN, FMCGIdentifierData.TargetMarketCode))
	if err != nil {
		return FMCTProductStatusReturn{}, err
	}

	if resp.IsError() {
		fmt.Printf("resp is err statusCode: %v. Dump: %v\n", resp.StatusCode, resp.Dump())
		return FMCTProductStatusReturn{}, resp.Err
	}

	return FMCTProductStatusReturn{
		Body: resp.Result().(*FMCGProductStatusResult),
	}, nil

}
