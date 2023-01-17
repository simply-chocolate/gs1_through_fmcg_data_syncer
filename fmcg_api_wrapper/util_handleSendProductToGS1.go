package fmcg_api_wrapper

import (
	"fmt"
	"gs1_syncer/sap_api_wrapper"
)

func SendGTINToGS1(identifierData FMCGIdentifierData, itemCode string) error {
	err := FMCGSendToGS1(identifierData, 0)
	if err != nil {
		return fmt.Errorf("error sending product with GTIN:%v to GS1. \nError:%v", identifierData.GTIN, err)
	}

	err = GetProductStatusAndSetStatusInSAP(identifierData, itemCode)
	if err != nil {
		return err
	}

	return nil
}

func GetProductStatusAndSetStatusInSAP(identifierData FMCGIdentifierData, itemCode string) error {
	resp, err := FMCGApiGetProductStatus(identifierData, 0)
	if err != nil {
		return fmt.Errorf("error getting the GS1 status GTIN:%v from FMCG. \nError:%v", identifierData.GTIN, err)
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
