package fmcg_api_wrapper

import (
	"fmt"
	"gs1_syncer/sap_api_wrapper"
	"time"
)

func SendGTINToGS1(identiferData FMCGIdentifierData, itemCode string) error {
	err := FMCGSendToGS1(identiferData, 0)
	if err != nil {
		return fmt.Errorf("error sending product with GTIN:%v to GS1. \nError:%v", identiferData.GTIN, err)
	}

	fmt.Printf("Just posted the Product with GTIN: %v to GS1. time now is: %v\n", identiferData.GTIN, time.Now())

	resp, err := FMCGApiGetProductStatus(identiferData, 0)
	if err != nil {
		return fmt.Errorf("error getting the GS1 status GTIN:%v from FMCG. \nError:%v", identiferData.GTIN, err)
	}

	gs1Resp := ""
	for _, gs1Response := range resp.Body.Gs1Response {
		gs1Resp += gs1Response
	}

	err = sap_api_wrapper.SetGs1StatusAndResponse(itemCode, resp.Body.Gs1Status, gs1Resp)
	if err != nil {
		return err
	}

	return nil
}
