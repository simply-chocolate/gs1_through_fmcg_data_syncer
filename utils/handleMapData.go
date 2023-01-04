package utils

import (
	"fmt"
	"gs1_syncer/fmcg_api_wrapper"
	"strconv"
	"strings"
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
		for _, ItemBarCodeCollection := range itemData.ItemBarCodeCollection {
			// Check which UoM and then Map for BaseUnit or for Case
			if ItemBarCodeCollection.UoMEntry == 1 {
				var baseItemData fmcg_api_wrapper.FmcgProductBodyBaseItem
				// TODO: Lav et check der finder ud af om dette GLN nummer allerede findes i GS1(FMCG Systemet) for at bestemme om D8164 skal være
				baseItemData.DataType = "CORRECT"

				// GEneral Information
				baseItemData.GTIN = "0" + ItemBarCodeCollection.Barcode
				baseItemData.DataCarrierTypeCode = "EAN_13"
				baseItemData.CountryOfOrigin = "208"
				baseItemData.ManufacturerGLN = "5790002336560"
				baseItemData.BrandName = itemData.BrandName
				baseItemData.GPCCategoryCode = "10000045"
				baseItemData.ImportClassificationValue = "18069019 00"
				baseItemData.ImportClassificationType = "CUSTOMS_TARIFF_NUMBER"
				baseItemData.TargetMarketCode = "208"
				baseItemData.ItemCode = itemData.ItemCode
				baseItemData.ItemNameDA = itemData.ItemNameDA
				baseItemData.ItemNameLanguageCodeDA = "da"
				baseItemData.FunctionalProductNameDA = itemData.FunctionalName
				baseItemData.FunctionalProductNameLanguageCodeDA = "da"
				baseItemData.RegulatedProductNameDA = "Chokolade"
				baseItemData.RegulatedProductNameLanguageCodeDA = "da"
				baseItemData.UnitOfMeasure = "BASE_UNIT_OR_EACH"
				baseItemData.IsOrderingUnit = false
				baseItemData.ShelfLifeFromArrivalInDays, err = strconv.Atoi(itemData.ShelfLifeFromArrivalInDays)
				if err != nil {
					return fmt.Errorf("error converting shelfLife to int. err: %v", err)
				}

				baseItemData.ShelfLifeFromProductionInDays = itemData.ShelfLifeFromProductionInDays
				baseItemData.IsQuantityOrPriceVarying = false
				baseItemData.DangerousContent = "FALSE"
				baseItemData.RelevantForPriceComparison = "FALSE"
				baseItemData.IsConsumerUnit = true
				baseItemData.IsShippingUnit = false
				baseItemData.IsPackagingMarkedReturnable = false
				var organicCode int
				if strings.Contains(itemData.ItemNameDA, "ØKO") {
					organicCode = 1
				} else {
					organicCode = 5
				}
				baseItemData.OrganicTradeItemCodeList = organicCode
				// Logistical Data
				baseItemData.NetContent = itemData.BaseUnitNetWeight
				baseItemData.NetContentUoM = "GRM"
				baseItemData.Height = itemData.BaseUnitHeight
				baseItemData.HeightUOM = "MMT"
				baseItemData.Width = itemData.BaseUnitWidth
				baseItemData.WidthUOM = "MMT"
				baseItemData.Depth = itemData.BaseUnitDepth
				baseItemData.DepthUOM = "MMT"
				baseItemData.NetWeight = itemData.BaseUnitNetWeight
				baseItemData.NetWeightUoM = "GRM"
				baseItemData.GrossWeight = itemData.BaseUnitGrossWeight
				baseItemData.GrossWeightUoM = "GRM"

				// TODO: These probably need some formatting

				baseItemData.EffectiveDateTime, err = FormatSapDateToFMCHDate(itemData.EffectiveDateTime)
				if err != nil {
					return fmt.Errorf("error converting effectiveDateTime to FMCG Format. err: %v", err)
				}
				baseItemData.StartAvailabilityDateTime, err = FormatSapDateToFMCHDate(itemData.AvailabilityDateTime)
				if err != nil {
					return fmt.Errorf("error converting AvailabilityDateTime to FMCG Format. err: %v", err)
				}

				// Allergen information
				baseItemData = MapAllergens(baseItemData, itemData)
				baseItemData = MapNutritionalInformation(baseItemData, itemData)

				err = fmcg_api_wrapper.FMCGApiPostBaseItem(baseItemData, 0)
				if err != nil {
					return fmt.Errorf("error postign the baseItem to FMCG. Err: %v", err)
				}

			} else if ItemBarCodeCollection.UoMEntry == 2 {
				var caseData fmcg_api_wrapper.FmcgProductBodyCase
				caseData.GTIN = "0" + ItemBarCodeCollection.Barcode
				caseData.UnitOfMeasure = "CASE"
			}
		}
	}

	return nil
}
