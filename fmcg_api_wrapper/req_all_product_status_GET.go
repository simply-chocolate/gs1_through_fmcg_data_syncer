package fmcg_api_wrapper

import "fmt"

type FMCGAllProductsStatusResults struct {
	StatusCode int                 `json:"status"`
	Products   []FMCGProductStatus `json:"products"`
}

type FMCGProductStatus struct {
	ProductId         string   `json:"productId"`
	LastModified      string   `json:"lastModified"`
	FmcgProductStatus string   `json:"fmcgProductsStatus"`
	Gs1Status         string   `json:"gs1Status"`
	Gs1Response       []string `json:"gs1Response"`
}

type FMCGAllProductsStatusReturn struct {
	Body *FMCGAllProductsStatusResults
}

func FMCGApiGetAllProductStatus(count int) (FMCGAllProductsStatusReturn, error) {
	resp, err := GetFMCGApiBaseClient().
		//DevMode().
		R().
		EnableDump().
		SetResult(FMCGAllProductsStatusResults{}).
		Get("/products/status")
	if err != nil {
		return FMCGAllProductsStatusReturn{}, err
	}

	if resp.IsError() {
		fmt.Printf("resp is err statusCode: %v. Dump: %v\n", resp.StatusCode, resp.Dump())
		return FMCGAllProductsStatusReturn{}, resp.Err
	}

	return FMCGAllProductsStatusReturn{
		Body: resp.Result().(*FMCGAllProductsStatusResults),
	}, nil

}
