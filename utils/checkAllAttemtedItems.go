package utils

import (
	"fmt"
	"gs1_syncer/fmcg_api_wrapper"
	"gs1_syncer/sap_api_wrapper"
)

func UpdateAttemptedItemsStatus(AttemptedSapItemsData sap_api_wrapper.SapApiGetItemsDataResults) error {

	for _, item := range AttemptedSapItemsData.Value {
		for _, barcode := range item.ItemBarCodeCollection {
			status, err := fmcg_api_wrapper.FMCGApiGetProductStatus(fmcg_api_wrapper.FMCGIdentifierData{
				TargetMarketCode: "208",
				GTIN:             barcode.Barcode,
			}, 0)
			if err != nil {
				return fmt.Errorf("error getting the status of the product%v: %v", item.ItemCode, err)
			} else {
				if status.Body.FmcgProductStatus == "NOT_FOUND" {
					continue
				} else {
					gs1resp := ""
					for _, response := range status.Body.Gs1Response {
						gs1resp += response
					}
					sap_api_wrapper.SetGs1StatusAndResponse(item.ItemCode, status.Body.Gs1Status, gs1resp)
				}
			}
		}
	}

	return nil
}
