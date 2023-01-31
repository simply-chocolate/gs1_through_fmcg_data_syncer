package utils

import (
	"fmt"
	"gs1_syncer/fmcg_api_wrapper"
	"gs1_syncer/sap_api_wrapper"
)

func IterateMixCasesAndMapToFMCGFormat(
	mixDisplays []sap_api_wrapper.SapApiItemsData,
) error {
	for _, itemData := range mixDisplays {
		for _, ItemBarCodeCollection := range itemData.ItemBarCodeCollection {
			if ItemBarCodeCollection.UoMEntry == 1 {
				return fmt.Errorf("error: UoMEntry is 1 for a mixDisplay. GTIN: %v", ItemBarCodeCollection.Barcode)
			} else if ItemBarCodeCollection.UoMEntry == 2 {
				var mixCaseData fmcg_api_wrapper.FmcgProductBodyMixCase
				mixCaseData.GTIN = "0" + ItemBarCodeCollection.Barcode

				err := MapMixCaseData(mixCaseData, itemData)
				if err != nil {
					// TODO: Teams notifier this
					return fmt.Errorf("error mapping the MixCase. GTIN: %v \nError: %v", mixCaseData.GTIN, err)
				}
			}
		}
	}
	return nil
}
