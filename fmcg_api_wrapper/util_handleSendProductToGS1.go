package fmcg_api_wrapper

import (
	"fmt"
	"gs1_syncer/sap_api_wrapper"
	"time"
)

func SendGTINToGS1(identifierData FMCGIdentifierData, itemCode string) error {
	result, err := FMCGSendToGS1(identifierData, 0)
	if err != nil {
		return fmt.Errorf("error sending product with GTIN:%v to GS1. \nError:%v", identifierData.GTIN, err)
	}

	// Wait a minute and see if the status has changed
	time.Sleep(1 * time.Minute)

	err = GetProductStatusAndSetStatusInSAP(identifierData, itemCode, result)
	if err != nil {
		return err
	}

	return nil
}

func GetProductStatusAndSetStatusInSAP(identifierData FMCGIdentifierData, itemCode string, gs1PostResult string) error {
	resp, err := FMCGApiGetProductStatus(identifierData, 0)
	if err != nil {
		return fmt.Errorf("error getting the GS1 status GTIN:%v from FMCG. \nError:%v", identifierData.GTIN, err)
	}

	gs1Resp := ""
	gs1Status := ""

	if resp.Body.FmcgProductStatus == "NOT_FOUND" {
		gs1Status = "NEVER_SENT"

	} else {
		gs1Status = resp.Body.Gs1Status

		if len(resp.Body.Gs1Response) == 0 {
			gs1Resp = gs1PostResult

		} else {
			for _, gs1Response := range resp.Body.Gs1Response {
				gs1Resp += gs1Response
			}
		}
	}

	err = sap_api_wrapper.SetGs1StatusAndResponse(itemCode, gs1Status, gs1Resp)
	if err != nil {
		return err
	}

	return nil
}
