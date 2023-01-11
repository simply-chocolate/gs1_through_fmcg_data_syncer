package fmcg_api_wrapper

import "fmt"

type FMCGIdentifierData struct {
	GTIN             string `json:"D8165"`
	TargetMarketCode string `json:"D8255"`
}

func FMCGSendToGS1(FMCGIdentifierData FMCGIdentifierData, count int) error {
	resp, err := GetFMCGApiBaseClient().
		//DevMode().
		R().
		EnableDump().
		SetResult(FmcgProductPostResult{}).
		SetBody(map[string]interface{}{
			"D8165": FMCGIdentifierData.GTIN,
			"D8255": FMCGIdentifierData.TargetMarketCode,
		}).
		Post("sendToGS1")
	if err != nil {
		return err
	}

	if resp.IsError() {
		fmt.Printf("resp is err statusCode: %v. Dump: %v\n", resp.StatusCode, resp.Dump())
		return resp.Err
	}

	return nil
}
