package utils

import (
	"fmt"
	"gs1_syncer/fmcg_api_wrapper"
	"gs1_syncer/sap_api_wrapper"
)

func MapBaseItemData(baseItemData fmcg_api_wrapper.FmcgProductBodyBaseItem, itemData sap_api_wrapper.SapApiItemsData) (fmcg_api_wrapper.FmcgProductBodyBaseItem, error) {
	var err error
	// TODO: Lav et check der finder ud af om dette GLN nummer allerede findes i GS1(FMCG Systemet) for at bestemme om D8164 skal være
	baseItemData.DataType = "CORRECT"

	// General Information
	baseItemData.DataCarrierTypeCode = "EAN_13"
	baseItemData.CountryOfOrigin = "208"

	if itemData.ManufacturerGLN == "" {
		baseItemData.ManufacturerGLN = "5790002336560"
	} else {
		baseItemData.ManufacturerGLN = itemData.ManufacturerGLN
	}

	if itemData.BrandOwnerGLN == "" {
		baseItemData.BrandOwnerGLN = "5790002336560"
	} else {
		baseItemData.BrandOwnerGLN = itemData.BrandOwnerGLN
	}

	baseItemData.BrandName = itemData.BrandName
	baseItemData.GPCCategoryCode = "10000045"
	baseItemData.ImportClassificationValue = "18069019"
	baseItemData.ImportClassificationType = "INTRASTAT"
	baseItemData.TargetMarketCode = "208"
	baseItemData.ItemCode = itemData.ItemCode

	if len(itemData.ItemNameDA) > 35 {
		baseItemData.ItemNameDA = itemData.ItemNameDA[0:34]
	} else {
		baseItemData.ItemNameDA = itemData.ItemNameDA
	}

	baseItemData.ItemNameLanguageCodeDA = "da"
	baseItemData.FunctionalProductNameDA = itemData.FunctionalProductNameDA
	baseItemData.FunctionalProductNameLanguageCodeDA = "da"
	baseItemData.RegulatedProductNameDA = "Chokolade"
	baseItemData.RegulatedProductNameLanguageCodeDA = "da"
	baseItemData.UnitOfMeasure = "BASE_UNIT_OR_EACH"
	baseItemData.IsOrderingUnit = false
	baseItemData.MaximumStorageTemp = itemData.MaximumStorageTemp
	baseItemData.MinimumStorageTemp = itemData.MinimumStorageTemp
	baseItemData.TemperatureType = "STORAGE_HANDLING"
	baseItemData.TemperatureOUM = "CEL"

	baseItemData.ShelfLifeFromProductionInDays = itemData.ShelfLifeFromProductionInDays
	shelfLifeAsInt := itemData.ShelfLifeFromArrivalInDays
	if shelfLifeAsInt == 0 {
		shelfLifeAsInt = int(float64(baseItemData.ShelfLifeFromProductionInDays) * 0.75)
	}
	baseItemData.ShelfLifeFromArrivalInDays = shelfLifeAsInt

	baseItemData.IsQuantityOrPriceVarying = false
	baseItemData.DangerousContent = "FALSE"
	baseItemData.RelevantForPriceComparison = "FALSE"
	baseItemData.IsConsumerUnit = true
	baseItemData.IsShippingUnit = false
	baseItemData.IsPackagingMarkedReturnable = false
	baseItemData.OrganicTradeItemCodeList = itemData.OrganicTradeItemCodeList

	// Dates
	baseItemData.EffectiveDateTime, err = FormatSapDateToFMCGDate(itemData.EffectiveDateTime)
	if err != nil {
		return fmcg_api_wrapper.FmcgProductBodyBaseItem{}, fmt.Errorf("error converting effectiveDateTime to FMCG Format. err: %v", err)
	}
	baseItemData.StartAvailabilityDateTime, err = FormatSapDateToFMCGDate(itemData.AvailabilityDateTime)
	if err != nil {
		return fmcg_api_wrapper.FmcgProductBodyBaseItem{}, fmt.Errorf("error converting AvailabilityDateTime to FMCG Format. err: %v", err)
	}

	// Logistical information
	baseItemData = MapLogisticalInformation(baseItemData, itemData)
	// Allergen information
	baseItemData = MapAllergens(baseItemData, itemData)
	// Nutritional information
	baseItemData = MapNutritionalInformation(baseItemData, itemData)

	return baseItemData, nil
}
