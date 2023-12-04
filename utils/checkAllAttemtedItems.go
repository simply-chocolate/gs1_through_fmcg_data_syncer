package utils

import (
	"fmt"
	"gs1_syncer/fmcg_api_wrapper"
	"gs1_syncer/sap_api_wrapper"
)

func UpdateAttemptedItemsStatus(AttemptedSapItemsData sap_api_wrapper.SapApiGetItemsDataResults) error {

	// Itteraate over the items and check if the status in FMCG has changed
	for _, item := range AttemptedSapItemsData.Value {

		// For each barcode, get the status from FMCG
		for _, barcode := range item.ItemBarCodeCollection {
			status, err := fmcg_api_wrapper.FMCGApiGetProductStatus(fmcg_api_wrapper.FMCGIdentifierData{
				TargetMarketCode: "208",
				GTIN:             barcode.Barcode,
			}, 0)
			if err != nil {
				return fmt.Errorf("error getting the status of the product%v: %v", item.ItemCode, err)
			} else {
				if status.Body.FmcgProductStatus == "NOT_FOUND" {
					// Product doesnt exist in FMCG yet
					continue
				} else {
					gs1resp := ""
					if status.Body.Gs1Status != "OK" {
						for _, gs1response := range status.Body.Gs1Response {
							gs1resp += gs1response
						}
						if gs1resp == "" {
							for _, validationError := range status.Body.ValidationErrors {
								gs1resp += fmt.Sprintf("FieldId: %v. Message: %v.\n", validationError.FieldId, validationError.Message)
							}
						}
						if gs1resp == "" {
							for _, sendStatus := range status.Body.SendStatusList {
								if sendStatus.SendStatus != "READY" {
									gs1resp += fmt.Sprintf("ProductId: %v. SendStatus: %v.\n", sendStatus.ProductId, sendStatus.SendStatus)
								}
							}
						}
					}
					sap_api_wrapper.SetGs1StatusAndResponse(item.ItemCode, status.Body.Gs1Status, gs1resp)
				}
			}
		}
	}

	return nil
}
