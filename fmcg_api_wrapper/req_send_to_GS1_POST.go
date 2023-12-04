package fmcg_api_wrapper

import "fmt"

type FMCGIdentifierData struct {
	GTIN             string `json:"D8165"`
	TargetMarketCode string `json:"D8255"`
}

type FMCGSendToGS1PostResult struct {
	Status            int    `json:"status"`
	ProductId         string `json:"productId"`
	Result            string `json:"result"`
	HierakiStatusList []struct {
		ProductID  string `json:"productId"`
		SendStatus string `json:"sendStatus"`
	}
}

func FMCGSendToGS1(FMCGIdentifierData FMCGIdentifierData, count int) (string, error) {
	resp, err := GetFMCGApiBaseClient().
		//DevMode().
		R().
		EnableDump().
		SetSuccessResult(FMCGSendToGS1PostResult{}).
		SetErrorResult(FMCGSendToGS1PostResult{}).
		SetBody(map[string]interface{}{
			"D8165": FMCGIdentifierData.GTIN,
			"D8255": FMCGIdentifierData.TargetMarketCode,
		}).
		Post("sendToGS1")
	if err != nil {
		return "", err
	}

	if resp.IsErrorState() {
		if resp.StatusCode == 400 {
			response := resp.ErrorResult().(*FMCGSendToGS1PostResult)

			return response.Result, nil
		} else {
			fmt.Printf("[SEND2CS1OTHR]: resp is err statusCode: %v. Dump: %v\n", resp.StatusCode, resp.Dump())
		}
		return "", resp.Err
	}

	response := resp.SuccessResult().(*FMCGSendToGS1PostResult)

	// Check the result

	return response.Result, nil
}
