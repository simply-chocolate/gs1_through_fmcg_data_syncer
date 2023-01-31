package utils

import (
	"gs1_syncer/fmcg_api_wrapper"
	"gs1_syncer/sap_api_wrapper"
	"gs1_syncer/teams_notifier"
)

func IterateProductsAndMapToFMCGFormat(
	SapItemsData sap_api_wrapper.SapApiGetItemsDataResults,
) []sap_api_wrapper.SapApiItemsData {
	var mixDisplays []sap_api_wrapper.SapApiItemsData

	for _, itemData := range SapItemsData.Value {
		var UnitGTIN string
		var err error

		// If the item is a mixDisplay we need to append it to the mixDisplays list and handle it later
		if itemData.UoMGroupEntry == 42 || itemData.TypeOfProduct == "KampagneDisplay" {
			for _, ItemBarCodeCollection := range itemData.ItemBarCodeCollection {
				if ItemBarCodeCollection.UoMEntry == 1 {
					teams_notifier.SendMappingErrorToTeams(ItemBarCodeCollection.Barcode, "At appending a mixCase to the mixCase list", "UoMEntry is 1 for a mixDisplay")
					continue
				} else if ItemBarCodeCollection.UoMEntry == 2 {
					mixDisplays = append(mixDisplays, itemData)
				}
			}
		} else {
			if len(itemData.ItemBarCodeCollection) == 1 && itemData.ItemBarCodeCollection[0].UoMEntry == 2 {
				mixDisplays = append(mixDisplays, itemData)
			} else {
				for _, ItemBarCodeCollection := range itemData.ItemBarCodeCollection {

					if ItemBarCodeCollection.UoMEntry == 1 {
						UnitGTIN = "0" + ItemBarCodeCollection.Barcode
						var baseItemData fmcg_api_wrapper.FmcgProductBodyBaseItem
						baseItemData.GTIN = "0" + ItemBarCodeCollection.Barcode

						baseItemData, err = MapBaseItemData(baseItemData, itemData)
						if err != nil {
							teams_notifier.SendMappingErrorToTeams(baseItemData.GTIN, "At mapping base item data", err.Error())
							continue
						}

						err = fmcg_api_wrapper.FMCGApiPostBaseItem(baseItemData, 0)
						if err != nil {
							teams_notifier.SendMappingErrorToTeams(baseItemData.GTIN, "At posting base item data", err.Error())
							continue
						}

					} else if ItemBarCodeCollection.UoMEntry == 2 {
						var caseData fmcg_api_wrapper.FmcgProductBodyCase
						caseData.GTIN = "0" + ItemBarCodeCollection.Barcode

						caseData, err = MapCaseData(caseData, itemData, UnitGTIN)
						if err != nil {
							teams_notifier.SendMappingErrorToTeams(caseData.GTIN, "At mapping case data", err.Error())
							continue
						}

						err = fmcg_api_wrapper.FMCGApiPostCase(caseData, 0)
						if err != nil {
							teams_notifier.SendMappingErrorToTeams(caseData.GTIN, "At posting case data", err.Error())
							continue

						}
					}
				}
			}
		}
	}
	return mixDisplays
}
