package utils

import (
	"fmt"
	"gs1_syncer/fmcg_api_wrapper"
	"gs1_syncer/sap_api_wrapper"
	"strconv"
)

func MapMixCaseData(mixCaseData fmcg_api_wrapper.FmcgProductBodyMixCase, itemData sap_api_wrapper.SapApiItemsData) (fmcg_api_wrapper.FmcgProductBodyMixCase, error) {

	mixCaseData.DataType = "CORRECT"
	mixCaseData.DataCarrierTypeCode = "EAN_13"
	mixCaseData.CountryOfOrigin = "208"

	if itemData.ManufacturerGLN == "" {
		mixCaseData.ManufacturerGLN = "5790002336560"
	} else {
		mixCaseData.ManufacturerGLN = itemData.ManufacturerGLN
	}

	if itemData.BrandOwnerGLN == "" {
		mixCaseData.BrandOwnerGLN = "5790002336560"
	} else {
		mixCaseData.BrandOwnerGLN = itemData.BrandOwnerGLN
	}

	mixCaseData.BrandName = itemData.BrandName
	mixCaseData.GPCCategoryCode = "10000045"
	mixCaseData.TargetMarketCode = "208"
	mixCaseData.ItemCode = itemData.ItemCode

	if len(itemData.ItemNameDA) > 35 {
		mixCaseData.ItemNameDA = itemData.ItemNameDA[0:34]
	} else {
		mixCaseData.ItemNameDA = itemData.ItemNameDA
	}

	mixCaseData.ItemNameLanguageCodeDA = "da"
	mixCaseData.FunctionalProductNameDA = itemData.FunctionalName
	mixCaseData.FunctionalProductNameLanguageCodeDA = "da"
	mixCaseData.UnitOfMeasure = "CASE"
	mixCaseData.IsOrderingUnit = true
	mixCaseData.IsPackageSalesReady = "FALSE" // TODO: Opret i SAP da barer jo teknisk set er, men andre ting er ikke.

	shelfLifeAsInt, err := strconv.Atoi(itemData.ShelfLifeFromArrivalInDays)
	if err != nil {
		return fmcg_api_wrapper.FmcgProductBodyMixCase{}, fmt.Errorf("error converting shelfLife to int. err: %v", err)
	}
	mixCaseData.ShelfLifeFromArrivalInDays = shelfLifeAsInt

	mixCaseData.ShelfLifeFromProductionInDays = itemData.ShelfLifeFromProductionInDays

	// Logistical Information
	//mixCaseData.UnitGTIN = baseItemGTIN
	//mixCaseData.UnitsPerCase = itemData.UnitsPerCase
	mixCaseData.PackagingType = "BX" // TODO: Add fields to SAP
	mixCaseData = MapLogisticalInformationMixCase(mixCaseData, itemData)

	// Dates
	mixCaseData.EffectiveDateTime, err = FormatSapDateToFMCHDate(itemData.EffectiveDateTime)
	if err != nil {
		return fmcg_api_wrapper.FmcgProductBodyMixCase{}, fmt.Errorf("error converting effectiveDateTime to FMCG Format. err: %v", err)
	}
	mixCaseData.StartAvailabilityDateTime, err = FormatSapDateToFMCHDate(itemData.AvailabilityDateTime)
	if err != nil {
		return fmcg_api_wrapper.FmcgProductBodyMixCase{}, fmt.Errorf("error converting AvailabilityDateTime to FMCG Format. err: %v", err)
	}

	// Base Units
	mixCaseData, err = MapBaseUnitsForMixCase(mixCaseData, itemData)
	if err != nil {
		return fmcg_api_wrapper.FmcgProductBodyMixCase{}, fmt.Errorf("error mapping baseunits to FMCG Format. GTIN: %v\nerr: %v", mixCaseData.GTIN, err)
	}

	return mixCaseData, err
}
