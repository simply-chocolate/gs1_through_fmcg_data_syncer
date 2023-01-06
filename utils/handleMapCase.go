package utils

import (
	"fmt"
	"gs1_syncer/fmcg_api_wrapper"
	"gs1_syncer/sap_api_wrapper"
	"strconv"
)

func MapCaseData(caseData fmcg_api_wrapper.FmcgProductBodyCase, itemData sap_api_wrapper.SapApiItemsData, baseItemGTIN string) (fmcg_api_wrapper.FmcgProductBodyCase, error) {
	caseData.DataType = "CORRECT"
	caseData.DataCarrierTypeCode = "EAN_13"
	caseData.CountryOfOrigin = "208"

	if itemData.ManufacturerGLN == "" {
		caseData.ManufacturerGLN = "5790002336560"
	} else {
		caseData.ManufacturerGLN = itemData.ManufacturerGLN
	}

	if itemData.BrandOwnerGLN == "" {
		caseData.BrandOwnerGLN = "5790002336560"
	} else {
		caseData.BrandOwnerGLN = itemData.BrandOwnerGLN
	}

	caseData.BrandName = itemData.BrandName
	caseData.GPCCategoryCode = "10000045"
	caseData.TargetMarketCode = "208"
	caseData.ItemCode = itemData.ItemCode

	if len(itemData.ItemNameDA) > 35 {
		caseData.ItemNameDA = itemData.ItemNameDA[0:34]
	} else {
		caseData.ItemNameDA = itemData.ItemNameDA
	}

	caseData.ItemNameLanguageCodeDA = "da"
	caseData.FunctionalProductNameDA = itemData.FunctionalName
	caseData.FunctionalProductNameLanguageCodeDA = "da"
	caseData.UnitOfMeasure = "CASE"
	caseData.IsOrderingUnit = true
	caseData.IsPackageSalesReady = "FALSE" // TODO: Opret i SAP da barer jo teknisk set er, men andre ting er ikke.

	shelfLifeAsInt, err := strconv.Atoi(itemData.ShelfLifeFromArrivalInDays)
	if err != nil {
		return fmcg_api_wrapper.FmcgProductBodyCase{}, fmt.Errorf("error converting shelfLife to int. err: %v", err)
	}
	caseData.ShelfLifeFromArrivalInDays = shelfLifeAsInt

	caseData.ShelfLifeFromProductionInDays = itemData.ShelfLifeFromProductionInDays

	// Logistical Information
	caseData.UnitGTIN = baseItemGTIN
	caseData.UnitsPerCase = itemData.UnitsPerCase
	caseData.PackagingType = "BX" // TODO: Add fields to SAP
	caseData = MapLogisticalInformationCase(caseData, itemData)

	// Dates
	caseData.EffectiveDateTime, err = FormatSapDateToFMCHDate(itemData.EffectiveDateTime)
	if err != nil {
		return fmcg_api_wrapper.FmcgProductBodyCase{}, fmt.Errorf("error converting effectiveDateTime to FMCG Format. err: %v", err)
	}
	caseData.StartAvailabilityDateTime, err = FormatSapDateToFMCHDate(itemData.AvailabilityDateTime)
	if err != nil {
		return fmcg_api_wrapper.FmcgProductBodyCase{}, fmt.Errorf("error converting AvailabilityDateTime to FMCG Format. err: %v", err)
	}

	return caseData, nil
}
