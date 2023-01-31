package fmcg_api_wrapper

import "fmt"

type FMCGIdentifierData struct {
	GTIN             string `json:"D8165"`
	TargetMarketCode string `json:"D8255"`
}

type FMCGSendToGS1PostResult struct {
	Status    int    `json:"status"`
	ProductId string `json:"productId"`
	Result    string `json:"result"`
}

func FMCGSendToGS1(FMCGIdentifierData FMCGIdentifierData, count int) (string, error) {
	resp, err := GetFMCGApiBaseClient().
		//DevMode().
		R().
		EnableDump().
		SetResult(FMCGSendToGS1PostResult{}).
		SetBody(map[string]interface{}{
			"D8165": FMCGIdentifierData.GTIN,
			"D8255": FMCGIdentifierData.TargetMarketCode,
		}).
		Post("sendToGS1")
	if err != nil {
		return "", err
	}

	if resp.IsError() {
		fmt.Printf("resp is err statusCode: %v. Dump: %v\n", resp.StatusCode, resp.Dump())
		return "", resp.Err
	}

	response := resp.Result().(*FMCGSendToGS1PostResult)

	// Check the result

	return response.Result, nil
}
