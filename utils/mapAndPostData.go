package utils

import (
	"fmt"
	"gs1_syncer/fmcg_api_wrapper"
	"gs1_syncer/sap_api_wrapper"
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

	var mixCases []sap_api_wrapper.SapApiItemsData

	for _, itemData := range SapItemsData.Value {
		var UnitGTIN string

		// If the item is a mixDisplay we need to append it to the mixCases list and handle it later
		if itemData.UoMGroupEntry == 42 || itemData.TypeOfProduct == "KampagneDisplay" {
			for _, ItemBarCodeCollection := range itemData.ItemBarCodeCollection {
				if ItemBarCodeCollection.UoMEntry == 1 {
					return fmt.Errorf("error: UoMEntry is 1 for a mixDisplay. GTIN: %v", ItemBarCodeCollection.Barcode)
				} else if ItemBarCodeCollection.UoMEntry == 2 {
					mixCases = append(mixCases, itemData)
				}
			}
		} else {
			for _, ItemBarCodeCollection := range itemData.ItemBarCodeCollection {
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
	}

	time.Sleep(5 * time.Minute)

	// First we go through each item and check the GS1 status and set it in SAP
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

	// Then we need to call the data again from SAP to get the updated GS1 status
	SapItemsData, err = GetItemDataFromSap()
	if err != nil {
		fmt.Println("Couldn't get Invoices from SAP. Sleeping 10 minutes")
		time.Sleep(10 * time.Minute)
		SapItemsData, err = GetItemDataFromSap()
		if err != nil {
			return fmt.Errorf("error getting the invoices from SAP: %v", err)
		}
	}

	// Then we go through each mixDisplay and map the data
	for _, itemData := range mixCases {
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

	time.Sleep(5 * time.Minute)

	// Then we go through each of the mixCases and check the GS1 status and set it in SAP
	for _, itemData := range mixCases {
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
