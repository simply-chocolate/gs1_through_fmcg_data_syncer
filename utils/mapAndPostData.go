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

	FMCGProductsStatusTimes, err := GetAllProductsStatusFMCG()
	if err != nil {
		return fmt.Errorf("error getting the product statuses from FMCG. err: %v", err)
	}

	for _, itemData := range SapItemsData.Value {
		var UnitGTIN string

		// Check if UoMGroupEntry = 42 if yes we need to call a custom Query the tells us which products and EANS are contained in this product
		if itemData.UoMGroupEntry == 42 || itemData.TypeOfProduct == "Campaign Display" {
			continue
			/*
				shouldBeProcessed, err := handleCheckIfSapUpdateTimeIsNewer(itemData, FMCGProductsStatusTimes[ItemBarCodeCollection.Barcode])
				if err != nil {
					return fmt.Errorf("error checking if mixDisplay should be processed by comparing update times for GTIN:%v\n error:%v", ItemBarCodeCollection.Barcode, err)
				}

				if !shouldBeProcessed {
					fmt.Println("This mixDisplay has a more recent update in FMCG system, so it will not be processed.")
					continue
				}

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
			*/
		}

		for _, ItemBarCodeCollection := range itemData.ItemBarCodeCollection {
			// Check which UoM and then Map for BaseUnit or for Case
			if ItemBarCodeCollection.UoMEntry == 1 {
				UnitGTIN = "0" + ItemBarCodeCollection.Barcode
				var baseItemData fmcg_api_wrapper.FmcgProductBodyBaseItem
				baseItemData.GTIN = "0" + ItemBarCodeCollection.Barcode

				shouldBeProcessed, err := handleCheckIfSapUpdateTimeIsNewer(itemData, FMCGProductsStatusTimes, baseItemData.GTIN)
				if err != nil {
					return fmt.Errorf("error checking if baseItem should be processed by comparing update times for GTIN:%v\n error:%v", ItemBarCodeCollection.Barcode, err)
				}

				if !shouldBeProcessed {
					continue
				}

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
				caseData.GTIN = "0" + ItemBarCodeCollection.Barcode

				shouldBeProcessed, err := handleCheckIfSapUpdateTimeIsNewer(itemData, FMCGProductsStatusTimes, caseData.GTIN)
				if err != nil {
					return fmt.Errorf("error checking if case should be processed by comparing update times for GTIN:%v\n error:%v", ItemBarCodeCollection.Barcode, err)
				}

				if !shouldBeProcessed {
					continue
				}

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

	return nil
}
