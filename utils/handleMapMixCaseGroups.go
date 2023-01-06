package utils

import (
	"fmt"
	"gs1_syncer/fmcg_api_wrapper"
	"gs1_syncer/sap_api_wrapper"
	"time"
)

func MapLogisticalInformationMixCase(mixCaseData fmcg_api_wrapper.FmcgProductBodyMixCase, itemData sap_api_wrapper.SapApiItemsData) fmcg_api_wrapper.FmcgProductBodyMixCase {
	mixCaseData.NetContent = itemData.CaseNetWeight
	mixCaseData.NetContentUoM = "GRM"
	mixCaseData.Height = itemData.CaseHeight
	mixCaseData.HeightUOM = "MMT"
	mixCaseData.Width = itemData.CaseWidth
	mixCaseData.WidthUOM = "MMT"
	mixCaseData.Depth = itemData.CaseDepth
	mixCaseData.DepthUOM = "MMT"
	mixCaseData.NetWeight = itemData.CaseNetWeight
	mixCaseData.NetWeightUoM = "GRM"
	mixCaseData.GrossWeight = itemData.CaseGrossWeight
	mixCaseData.GrossWeightUoM = "GRM"
	return mixCaseData
}

func MapBaseUnitsForMixCase(mixCaseData fmcg_api_wrapper.FmcgProductBodyMixCase, itemData sap_api_wrapper.SapApiItemsData) (fmcg_api_wrapper.FmcgProductBodyMixCase, error) {
	SapMixCaseContent, err := GetMixCaseItemsFromSap(itemData.ItemCode)
	if err != nil {
		fmt.Println("Couldn't get Invoices from SAP. Sleeping 10 minutes")
		time.Sleep(10 * time.Minute)
		SapMixCaseContent, err = GetMixCaseItemsFromSap(itemData.ItemCode)
		if err != nil {
			return fmcg_api_wrapper.FmcgProductBodyMixCase{}, fmt.Errorf("error getting the invoices from SAP: %v", err)
		}
	}

	for _, contentItem := range SapMixCaseContent.Value {
		mixContentItemInfo, err := GetMixContentItemInfoFromSap(contentItem.ItemCode)
		if err != nil {
			return fmcg_api_wrapper.FmcgProductBodyMixCase{}, fmt.Errorf("error getting MixContentItemInfo from SAP. err:%v", err)
		}
		for _, contentItemData := range mixContentItemInfo.Value {
			for _, barcodeCollection := range contentItemData.ItemBarCodeCollection {
				if barcodeCollection.UoMEntry == 2 {
					// TODO: Add some logic to check if this Barcode exists within FMCG/GS1 before you do anything with it.
					fmt.Printf("We're gonna do something with this barcode: %v\n", barcodeCollection.Barcode)
				}
			}
		}

	}
	return mixCaseData, nil
}
