package utils

import (
	"fmt"
	"gs1_syncer/fmcg_api_wrapper"
	"gs1_syncer/sap_api_wrapper"
	"strconv"
)

func MapMixCaseData(mixCaseData fmcg_api_wrapper.FmcgProductBodyMixCase, itemData sap_api_wrapper.SapApiItemsData) error {

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
	mixCaseData.FunctionalProductNameDA = itemData.FunctionalProductNameDA
	mixCaseData.FunctionalProductNameLanguageCodeDA = "da"

	mixCaseData.UnitOfMeasure = "DISPLAY_SHIPPER"
	mixCaseData.IsOrderingUnit = true
	mixCaseData.IsPackageSalesReady = "FALSE"

	shelfLifeAsInt, err := strconv.Atoi(itemData.ShelfLifeFromArrivalInDays)
	if err != nil {
		return fmt.Errorf("error converting shelfLife to int on mixCaseData. GTIN: %v err: %v", mixCaseData.GTIN, err)
	}
	mixCaseData.ShelfLifeFromArrivalInDays = shelfLifeAsInt

	mixCaseData.ShelfLifeFromProductionInDays = itemData.ShelfLifeFromProductionInDays

	// Logistical Information

	mixCaseData.PackagingType = itemData.PackagingType
	mixCaseData = MapLogisticalInformationMixCase(mixCaseData, itemData)

	// Dates
	mixCaseData.EffectiveDateTime, err = FormatSapDateToFMCGDate(itemData.EffectiveDateTime)
	if err != nil {
		return fmt.Errorf("error converting effectiveDateTime to FMCG Format. err: %v", err)
	}
	mixCaseData.StartAvailabilityDateTime, err = FormatSapDateToFMCGDate(itemData.AvailabilityDateTime)
	if err != nil {
		return fmt.Errorf("error converting AvailabilityDateTime to FMCG Format. err: %v", err)
	}

	// Base Units
	mixCaseContent, err := MapBaseUnitsForMixCase(mixCaseData, itemData)
	if err != nil {
		return fmt.Errorf("error mapping base units to FMCG Format. GTIN: %v\nerr: %v", mixCaseData.GTIN, err)
	}

	fmt.Printf("mixCaseContent: %v", mixCaseContent)

	err = fmcg_api_wrapper.FMCGApiPostMixCase(mixCaseData, mixCaseContent, 0)
	if err != nil {
		return fmt.Errorf("error posting the case to FMCG. GTIN: %v \nError: %v", mixCaseData.GTIN, err)
	}

	return nil
}
