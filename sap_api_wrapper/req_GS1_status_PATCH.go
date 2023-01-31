package sap_api_wrapper

import (
	"fmt"
	"time"
)

type GS1StatusAndResponseBody struct {
	GS1Status      string `json:"U_CCF_GS1_Status"`
	GS1Response    string `json:"U_CCF_GS1_Response"`
	GS1FMCGStatus  string `json:"U_CCF_Sync_GS1"`
	FMCGUpdateTime string `json:"U_CCF_FMCG_Update_Time"`
}

type GS1StatusAndResponseResult struct {
}

// Takes the Gs1Status and Gs1 Response and updates the item in SAP
func SetGs1StatusAndResponse(itemCode string, GS1Status string, GS1Response string) error {
	var body GS1StatusAndResponseBody
	body.GS1Status = GS1Status
	body.GS1Response = GS1Response
	body.GS1FMCGStatus = "A"
	body.FMCGUpdateTime = time.Now().Format("2006-01-02")

	// If the product has not been validated by GS1, we set the FMCG status to "Y" so that it will be synced again
	if GS1Status != "OK" {
		body.GS1FMCGStatus = "Y"
	}

	client, err := GetSapApiAuthClient()
	if err != nil {
		fmt.Println("Error getting an authenticaed client")
		return err
	}

	resp, err := client.
		//DevMode().
		R().
		EnableDump().
		SetResult(SapApiPostLoginResult{}).
		SetHeader("Content-Type", "application/json").
		SetBody(body).
		Patch(fmt.Sprintf("Items('%v')", itemCode))
	if err != nil {
		return err
	}
	if resp.IsError() {
		fmt.Printf("resp is err statusCode: %v. Dump: %v\n", resp.StatusCode, resp.Dump())
		return resp.Err
	}

	if resp.StatusCode != 204 {
		return fmt.Errorf("unexpected errorcode when patching the items endpoint. Itemcode:%v. StatusCode:%v", itemCode, resp.StatusCode)
	}

	return nil
}
