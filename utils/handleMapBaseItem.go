package utils

import (
	"fmt"
	"gs1_syncer/fmcg_api_wrapper"
	"gs1_syncer/sap_api_wrapper"
	"strconv"
	"strings"
)

func MapBaseItemData(baseItemData fmcg_api_wrapper.FmcgProductBodyBaseItem, itemData sap_api_wrapper.SapApiItemsData) (fmcg_api_wrapper.FmcgProductBodyBaseItem, error) {

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

	shelfLifeAsInt, err := strconv.Atoi(itemData.ShelfLifeFromArrivalInDays)
	if err != nil {
		return fmcg_api_wrapper.FmcgProductBodyBaseItem{}, fmt.Errorf("error converting shelfLife to int on baseItem. GTIN: %v err: %v", baseItemData.GTIN, err)
	}
	baseItemData.ShelfLifeFromArrivalInDays = shelfLifeAsInt

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
	baseItemData.PackagingType = "BX" // TODO: Add fields to SAP

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
