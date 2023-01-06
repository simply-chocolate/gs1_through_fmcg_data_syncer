package utils

import (
	"fmt"
	"gs1_syncer/fmcg_api_wrapper"
	"time"
)

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

		// Check if UoMGroupEntry = 42 if yes we need to call a custom Query the tells us which products and EANS are contained in this product
		if itemData.UoMGroupEntry == 42 {
			var mixCaseData fmcg_api_wrapper.FmcgProductBodyMixCase

			for _, ItemBarCodeCollection := range itemData.ItemBarCodeCollection {
				if ItemBarCodeCollection.Barcode != "" {
					mixCaseData.GTIN = "0" + ItemBarCodeCollection.Barcode
				}
			}

			mixCaseData, err = MapMixCaseData(mixCaseData, itemData)
			if err != nil {
				return fmt.Errorf("error mapping the MixCase. GTIN: %v \nError: %v", mixCaseData.GTIN, err)
			}

			err = fmcg_api_wrapper.FMCGApiPostMixCase(mixCaseData, 0)
			if err != nil {
				return fmt.Errorf("error posting the case to FMCG. GTIN: %v \nError: %v", mixCaseData.GTIN, err)
			}

			return nil
		}

		for _, ItemBarCodeCollection := range itemData.ItemBarCodeCollection {
			// Check which UoM and then Map for BaseUnit or for Case
			if ItemBarCodeCollection.UoMEntry == 1 {
				UnitGTIN = "0" + ItemBarCodeCollection.Barcode

				var baseItemData fmcg_api_wrapper.FmcgProductBodyBaseItem
				baseItemData.GTIN = ItemBarCodeCollection.Barcode
				baseItemData, err = MapBaseItemData(baseItemData, itemData)
				if err != nil {
					return fmt.Errorf("error mapping the baseItem. GTIN: %v \nError: %v", baseItemData.GTIN, err)
				}

				err = fmcg_api_wrapper.FMCGApiPostBaseItem(baseItemData, 0)
				if err != nil {
					return fmt.Errorf("error posting the baseItem to FMCG. GTIN: %v \nError: %v", baseItemData.GTIN, err)
				}

			} else if ItemBarCodeCollection.UoMEntry == 2 {
				var caseData fmcg_api_wrapper.FmcgProductBodyCase
				caseData.GTIN = ItemBarCodeCollection.Barcode

				caseData, err := MapCaseData(caseData, itemData, UnitGTIN)
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

	return nil
}
