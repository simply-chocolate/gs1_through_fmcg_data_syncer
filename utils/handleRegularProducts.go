package utils

import (
	"fmt"
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

		if itemData.UpdateDate <= itemData.LastSyncDate {
			continue
		}

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

			// If theres only 1 barcode and its a case, we need to handle it as a mixDisplay cause it a "20" version of a bar.
		} else if len(itemData.ItemBarCodeCollection) == 1 && itemData.ItemBarCodeCollection[0].UoMEntry == 2 {
			mixDisplays = append(mixDisplays, itemData)
		} else {
			for _, ItemBarCodeCollection := range itemData.ItemBarCodeCollection {
				if ItemBarCodeCollection.UoMEntry == 1 {
					UnitGTIN = "0" + ItemBarCodeCollection.Barcode
					var baseItemData fmcg_api_wrapper.FmcgProductBodyBaseItem
					baseItemData.GTIN = "0" + ItemBarCodeCollection.Barcode
					if baseItemData.GTIN == "" {
						teams_notifier.SendMappingErrorToTeams(itemData.ItemCode, "At mapping base item data", "GTIN is empty")
						continue
					}
					baseItemData, err = MapBaseItemData(baseItemData, itemData)
					if err != nil {
						if baseItemData.GTIN == "" {
							teams_notifier.SendMappingErrorToTeams(itemData.ItemCode, "At mapping base item data", err.Error())
						} else {
							teams_notifier.SendMappingErrorToTeams(baseItemData.GTIN, "At mapping base item data", err.Error())
						}
						fmt.Printf("Error at mapping item %v. error: %v", itemData.ItemCode, err)
						continue
					}

					err = fmcg_api_wrapper.FMCGApiPostBaseItem(baseItemData)
					if err != nil {
						if baseItemData.GTIN == "" {
							teams_notifier.SendMappingErrorToTeams(itemData.ItemCode, "At posting base item data", err.Error())
						} else {
							teams_notifier.SendMappingErrorToTeams(baseItemData.GTIN, "At posting base item data", err.Error())
						}
						continue
					}

				} else if ItemBarCodeCollection.UoMEntry == 2 {
					var caseData fmcg_api_wrapper.FmcgProductBodyCase
					caseData.GTIN = "0" + ItemBarCodeCollection.Barcode
					if caseData.GTIN == "" {
						teams_notifier.SendMappingErrorToTeams(itemData.ItemCode, "At mapping base item data", "GTIN is empty")
						continue
					}

					caseData, err = MapCaseData(caseData, itemData, UnitGTIN)
					if err != nil {
						if caseData.GTIN == "" {
							teams_notifier.SendMappingErrorToTeams(itemData.ItemCode, "At mapping case data", err.Error())
						} else {
							teams_notifier.SendMappingErrorToTeams(caseData.GTIN, "At mapping case data", err.Error())
						}
						continue
					}

					err = fmcg_api_wrapper.FMCGApiPostCase(caseData, 0)
					if err != nil {
						if caseData.GTIN == "" {
							teams_notifier.SendMappingErrorToTeams(itemData.ItemCode, "At mapping case data", err.Error())
						} else {
							teams_notifier.SendMappingErrorToTeams(caseData.GTIN, "At mapping case data", err.Error())
						}
						continue

					}

				}
			}
		}
	}

	return mixDisplays
}
