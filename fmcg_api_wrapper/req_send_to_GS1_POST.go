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

	// TODO: Implementer at den sender et kald til GS1 efter 10 minutter og tjekker
	// https://simplychocolate.fmcgproducts.dk/api/status/05710885039013.208 for at tjekke "gs1Status". Hvis den er = "FAILED"så skal vi lave en teams besked det bruger "gs1Response".
	// TODO: Vi skal bruge selvsamme endpoint til at identificere hvor vidt et produkt er i GS1, og bruge "gs1LastSendDate" til at sammenligne med Updatedate i SAP API
	// Hvis gs1LastSendDate er ældre end updatedate skal vi ikke gøre noget (vi kan måske bruge denne dato som filter)
	// Når den er gået igennem efter 10 minutter skal GS1_Status felt i SAP rettes til OK
	// Der skal lave et felt i SAP der kan indeholde den fejlbesked vi får fra GS1
	// Feltet i SAP skal bare hedde "gs1Status" og "gs1ResponsE"

	response := resp.Result().(*FmcgProductPostResult)
	for _, validationError := range response.ValidationErrors {
		fmt.Println("fieldId:", validationError.FieldId)
		fmt.Println("fieldLabel:", validationError.FieldLabel)
		fmt.Println("message:", validationError.Message)
		fmt.Println("messageType:", validationError.MessageType)
		fmt.Println("________________")
	}

	return nil
}
