package sap_api_wrapper

import "fmt"

type GS1StatusAndResponseBody struct {
	GS1Status   string `json:"U_CCF_GS1_Status"`
	GS1Response string `json:"U_CCF_GS1_Response"`
}

type GS1StatusAndResponseResult struct {
}

// Takes the Gs1Status and Gs1 Response and updates the item in SAP
func SetGs1StatusAndResponse(itemCode string, GS1Status string, GS1Response string) error {
	var body GS1StatusAndResponseBody
	body.GS1Status = GS1Status
	body.GS1Response = GS1Response

	client, err := GetSapApiAuthClient()
	if err != nil {
		fmt.Println("Error getting an authenticaed client")
		return err
	}

	resp, err := client.
		//DevMode().
		R().
		SetResult(SapApiPostLoginResult{}).
		SetHeader("Content-Type", "application/json").
		SetBody(body).
		Patch(fmt.Sprintf("Items('%v')", itemCode))
	if err != nil {
		return err
	}

	if resp.StatusCode != 204 {
		return fmt.Errorf("unexpected errorcode when patching the items endpoint. Itemcode:%v. StatusCode:%v", itemCode, resp.StatusCode)
	}

	return nil
}