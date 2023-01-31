package utils

import (
	"fmt"
	"gs1_syncer/fmcg_api_wrapper"
	"gs1_syncer/sap_api_wrapper"
	"math"
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

	mixCaseData.PalletGrossWeight = itemData.PalletGrossWeight
	mixCaseData.PalletGrossWeightUoM = "GRM"
	mixCaseData.PalletHeight = itemData.PalletHeight
	mixCaseData.PalletHeightUoM = "MMT"
	mixCaseData.PalletWidth = itemData.PalletWidth
	mixCaseData.PalletWidthUoM = "MMT"
	mixCaseData.PalletDepth = itemData.PalletDepth
	mixCaseData.PalletDepthUoM = "MMT"
	mixCaseData.LayersPerPallet = itemData.LayersPerPallet
	mixCaseData.PalletUnitsPerLayer = itemData.PalletUnitsPerLayer
	mixCaseData.PalletSendingUnitAmount = itemData.PalletSendingUnitAmount

	return mixCaseData
}

// Maps for the Bar Cases with only 20 units in them
// We need to remove the -20 from the item code and then get that base item GTIN from that, and just say units per case is 20
func Map20BarsCaseDisplay(
	mixCaseData fmcg_api_wrapper.FmcgProductBodyMixCase,
	itemData sap_api_wrapper.SapApiItemsData,
	baseUnits []fmcg_api_wrapper.FMCGMixCaseContentBaseItem,
) ([]fmcg_api_wrapper.FMCGMixCaseContentBaseItem, error) {

	baseItemItemCode := itemData.ItemCode[0 : len(itemData.ItemCode)-3]

	baseItem, err := GetMixContentItemInfoFromSap(baseItemItemCode)
	if err != nil {
		return []fmcg_api_wrapper.FMCGMixCaseContentBaseItem{}, fmt.Errorf("error getting the base item from SAP. ItemCode:%v\n err:%v", itemData.ItemCode, err)
	}

	if len(baseItem.Value) == 0 {
		return []fmcg_api_wrapper.FMCGMixCaseContentBaseItem{}, fmt.Errorf("error getting the base item from SAP. len of baseItem.Value == 0 ItemCode:%v\n err:%v", itemData.ItemCode, err)
	}
	if len(baseItem.Value[0].ItemBarCodeCollection) == 0 {
		return []fmcg_api_wrapper.FMCGMixCaseContentBaseItem{}, fmt.Errorf("error getting the base item from SAP. len of baseItem.Value[0].ItemBarCodeCollection == 0 ItemCode:%v\n err:%v", itemData.ItemCode, err)
	}

	var baseUnit fmcg_api_wrapper.FMCGMixCaseContentBaseItem
	baseUnit.UnitGTINItem = "0" + baseItem.Value[0].ItemBarCodeCollection[0].Barcode
	baseUnit.UnitsPerCase = 20.0

	// TODO: Figure out why the script doesn't reach this with GTIN 05710885015642
	fmt.Println("Finished adding baseunits to mixcase. BaseUnit:", baseUnit.UnitGTINItem, "UnitsPerCase:", baseUnit.UnitsPerCase, "ItemCode:", itemData.ItemCode)

	baseUnits = append(baseUnits, baseUnit)
	fmt.Printf("UoMGroupEntry = 40. baseUnits: %v\n", baseUnits)

	return baseUnits, nil
}

// Retrieves the mix case content from SAP and checks if the content has an ItemBarCodeCollection
// If it does, it adds it to the baseUnits slice
// IF not, it checks if it's a HF item, and if it is, it retrieves the GTIN from the HF item field.
func MapMixCaseContent(
	mixCaseData fmcg_api_wrapper.FmcgProductBodyMixCase,
	itemData sap_api_wrapper.SapApiItemsData,
	baseUnits []fmcg_api_wrapper.FMCGMixCaseContentBaseItem,
) ([]fmcg_api_wrapper.FMCGMixCaseContentBaseItem, error) {

	SapMixCaseContent, err := GetMixCaseItemsFromSap(itemData.ItemCode)
	if err != nil {
		fmt.Println("Couldn't get Invoices from SAP. Sleeping 10 minutes")
		time.Sleep(10 * time.Minute)
		SapMixCaseContent, err = GetMixCaseItemsFromSap(itemData.ItemCode)
		if err != nil {
			return []fmcg_api_wrapper.FMCGMixCaseContentBaseItem{}, fmt.Errorf("error getting the invoices from SAP: %v", err)
		}
	}

	for _, contentItem := range SapMixCaseContent.Value {
		mixContentItemInfo, err := GetMixContentItemInfoFromSap(contentItem.ItemCode)
		if err != nil {
			return []fmcg_api_wrapper.FMCGMixCaseContentBaseItem{}, fmt.Errorf("error getting MixContentItemInfo from SAP. err:%v", err)
		}

		for _, contentItemData := range mixContentItemInfo.Value {
			// We check if length of ItemBarCodeCollection is 0, because if it is, we need to use a different field for the GTIN.
			if len(contentItemData.ItemBarCodeCollection) == 0 {
				// If the item is not a HF, we skip it
				if contentItemData.TypeOfProduct != "HF" {
					continue
				}
				// If the item is a HF, but the BarCodeForHF is empty, we return an error, because then we can't get the GTIN
				if contentItemData.BarCodeForHF == "" {
					return []fmcg_api_wrapper.FMCGMixCaseContentBaseItem{}, fmt.Errorf("error getting MixContentItemInfo from SAP. BarCodeForHF is empty. ItemCode:%v", contentItemData.ItemCode)
				}

				var identifierData fmcg_api_wrapper.FMCGIdentifierData
				identifierData.GTIN = contentItemData.BarCodeForHF
				identifierData.TargetMarketCode = "208"

				productStatus, err := fmcg_api_wrapper.FMCGApiGetProductStatus(identifierData, 0)
				if err != nil {
					return []fmcg_api_wrapper.FMCGMixCaseContentBaseItem{}, fmt.Errorf("error getting content item from FMCG. If the status is 404 the it has not been synced yet. ItemCode:%v BarCode:%v\n err:%v", contentItemData.ItemCode, contentItemData.BarCodeForHF, err)
				}
				if productStatus.Body.Gs1Status != "OK" {
					return []fmcg_api_wrapper.FMCGMixCaseContentBaseItem{}, fmt.Errorf(" ItemCode:%v BarCode:%v has not been synced all the way to GS1 with OK. Check GS1 Status: %v\n err:%v", contentItemData.ItemCode, contentItemData.BarCodeForHF, productStatus.Body.Gs1Status, err)
				}

				var baseUnit fmcg_api_wrapper.FMCGMixCaseContentBaseItem
				baseUnit.UnitGTINItem = "0" + contentItemData.BarCodeForHF
				baseUnit.UnitsPerCase = math.Floor(contentItem.Quantity * float64(itemData.UnitsPerCase))

				baseUnits = append(baseUnits, baseUnit)
			}

			// In case theres a mix of FV and HF on the mixcase, we go through the regular mixCase flow as well
			for _, barcodeCollection := range contentItemData.ItemBarCodeCollection {
				// We use the baseItem as thats the amount on the Bill of Materials
				if barcodeCollection.UoMEntry == 1 {
					var identifierData fmcg_api_wrapper.FMCGIdentifierData
					identifierData.GTIN = "0" + barcodeCollection.Barcode
					identifierData.TargetMarketCode = "208"

					productStatus, err := fmcg_api_wrapper.FMCGApiGetProductStatus(identifierData, 0)
					if err != nil {
						return []fmcg_api_wrapper.FMCGMixCaseContentBaseItem{}, fmt.Errorf("error getting content item from FMCG. If the status is 404 the it has not been synced yet. ItemCode:%v BarCode:%v\n err:%v", contentItemData.ItemCode, barcodeCollection.Barcode, err)
					}
					if productStatus.Body.Gs1Status != "OK" {
						return []fmcg_api_wrapper.FMCGMixCaseContentBaseItem{}, fmt.Errorf(" ItemCode:%v BarCode:%v has not been synced all the way to GS1 with OK. Check GS1 Status: %v\n err:%v", contentItemData.ItemCode, barcodeCollection.Barcode, productStatus.Body.Gs1Status, err)
					}

					var baseUnit fmcg_api_wrapper.FMCGMixCaseContentBaseItem
					baseUnit.UnitGTINItem = "0" + barcodeCollection.Barcode
					baseUnit.UnitsPerCase = contentItem.Quantity
					baseUnits = append(baseUnits, baseUnit)

				}
			}
		}
	}

	return baseUnits, nil
}

// Figures out which base units to use for the mix case
// Then gets the required data from SAP for each base unit
// Then maps the data to the mix case
// Returns a mix case with the base units mapped
func MapBaseUnitsForMixCase(
	mixCaseData fmcg_api_wrapper.FmcgProductBodyMixCase,
	itemData sap_api_wrapper.SapApiItemsData,
) ([]fmcg_api_wrapper.FMCGMixCaseContentBaseItem, error) {
	var baseUnits []fmcg_api_wrapper.FMCGMixCaseContentBaseItem
	var err error

	if itemData.UoMGroupEntry == 40 {
		baseUnits, err = Map20BarsCaseDisplay(
			mixCaseData,
			itemData,
			baseUnits,
		)
		if err != nil {
			return []fmcg_api_wrapper.FMCGMixCaseContentBaseItem{}, fmt.Errorf("error mapping 20 bars case display. ItemCode:%v\n err:%v", itemData.ItemCode, err)
		}
	} else {
		baseUnits, err = MapMixCaseContent(
			mixCaseData,
			itemData,
			baseUnits)
		if err != nil {
			return []fmcg_api_wrapper.FMCGMixCaseContentBaseItem{}, fmt.Errorf("error mapping mix case content. ItemCode:%v\n err:%v", itemData.ItemCode, err)
		}
	}

	return baseUnits, nil
}
