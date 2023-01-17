package utils

import (
	"fmt"
	"gs1_syncer/fmcg_api_wrapper"
	"time"
)

// Calls the APIs and retrieves the information needed to handle the integration of data
func MapData() error {

	SapItemsData, err := GetItemDataFromSap()
	if err != nil {
		fmt.Println("Couldn't get Invoices from SAP. Sleeping 10 minutes")
		time.Sleep(10 * time.Minute)
		SapItemsData, err = GetItemDataFromSap()
		if err != nil {
			return fmt.Errorf("error getting the invoices from SAP: %v", err)
		}
	}

	for _, itemData := range SapItemsData.Value {
		var UnitGTIN string

		if itemData.UoMGroupEntry == 42 || itemData.TypeOfProduct == "KampagneDisplay" {
			for _, ItemBarCodeCollection := range itemData.ItemBarCodeCollection {
				if ItemBarCodeCollection.UoMEntry == 1 {
					return fmt.Errorf("error: UoMEntry is 1 for a mixDisplay. GTIN: %v", ItemBarCodeCollection.Barcode)
				} else if ItemBarCodeCollection.UoMEntry == 2 {
					var mixCaseData fmcg_api_wrapper.FmcgProductBodyMixCase
					mixCaseData.GTIN = "0" + ItemBarCodeCollection.Barcode

					mixCaseData, mixCaseContent, err := MapMixCaseData(mixCaseData, itemData)
					if err != nil {
						return fmt.Errorf("error mapping the MixCase. GTIN: %v \nError: %v", mixCaseData.GTIN, err)
					}

					err = fmcg_api_wrapper.FMCGApiPostMixCase(mixCaseData, mixCaseContent, 0)
					if err != nil {
						return fmt.Errorf("error posting the case to FMCG. GTIN: %v \nError: %v", mixCaseData.GTIN, err)
					}
				}
			}
		}
		for _, ItemBarCodeCollection := range itemData.ItemBarCodeCollection {
			// Check which UoM and then Map for BaseUnit or for Case
			if ItemBarCodeCollection.UoMEntry == 1 {
				UnitGTIN = "0" + ItemBarCodeCollection.Barcode
				var baseItemData fmcg_api_wrapper.FmcgProductBodyBaseItem
				baseItemData.GTIN = "0" + ItemBarCodeCollection.Barcode

				baseItemData, err = MapBaseItemData(baseItemData, itemData)
				if err != nil {
					return fmt.Errorf("error mapping the baseItem. GTIN: %v \nError: %v", baseItemData.GTIN, err)
				}

				err = fmcg_api_wrapper.FMCGApiPostBaseItem(baseItemData, 0)
				if err != nil {
					return err
				}

			} else if ItemBarCodeCollection.UoMEntry == 2 {
				var caseData fmcg_api_wrapper.FmcgProductBodyCase
				caseData.GTIN = "0" + ItemBarCodeCollection.Barcode

				caseData, err = MapCaseData(caseData, itemData, UnitGTIN)
				if err != nil {
					return fmt.Errorf("error mapping the case. GTIN: %v \nError: %v", caseData.GTIN, err)
				}

				err = fmcg_api_wrapper.FMCGApiPostCase(caseData, 0)
				if err != nil {
					return fmt.Errorf("error posting the case to FMCG. GTIN: %v \nError: %v", caseData.GTIN, err)
				}
			}
		}
	}

	// time.Sleep(30 * time.Minute)

	for _, itemData := range SapItemsData.Value {
		for _, ItemBarCodeCollection := range itemData.ItemBarCodeCollection {
			var identifierData fmcg_api_wrapper.FMCGIdentifierData
			identifierData.GTIN = "0" + ItemBarCodeCollection.Barcode
			identifierData.TargetMarketCode = "208"

			err := fmcg_api_wrapper.GetProductStatusAndSetStatusInSAP(identifierData, itemData.ItemCode)
			if err != nil {
				return fmt.Errorf("error getting the product status from FMCG and setting it in SAP while running through everything. GTIN: %v \nError: %v", identifierData.GTIN, err)
			}
		}
	}

	return nil
}
