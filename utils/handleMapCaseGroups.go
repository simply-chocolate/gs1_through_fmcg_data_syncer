package utils

import (
	"gs1_syncer/fmcg_api_wrapper"
	"gs1_syncer/sap_api_wrapper"
)

func MapLogisticalInformationCase(caseData fmcg_api_wrapper.FmcgProductBodyCase, itemData sap_api_wrapper.SapApiItemsData) fmcg_api_wrapper.FmcgProductBodyCase {
	caseData.NetContent = itemData.CaseNetWeight
	caseData.NetContentUoM = "GRM"
	caseData.Height = itemData.CaseHeight
	caseData.HeightUOM = "MMT"
	caseData.Width = itemData.CaseWidth
	caseData.WidthUOM = "MMT"
	caseData.Depth = itemData.CaseDepth
	caseData.DepthUOM = "MMT"
	caseData.NetWeight = itemData.CaseNetWeight
	caseData.NetWeightUoM = "GRM"
	caseData.GrossWeight = itemData.CaseGrossWeight
	caseData.GrossWeightUoM = "GRM"
	caseData.PalletGrossWeight = itemData.PalletGrossWeight
	caseData.PalletGrossWeightUoM = "GRM"
	caseData.PalletHeight = itemData.PalletHeight
	caseData.PalletHeightUoM = "MMT"
	caseData.PalletWidth = itemData.PalletWidth
	caseData.PalletWidthUoM = "MMT"
	caseData.PalletDepth = itemData.PalletDepth
	caseData.PalletDepthUoM = "MMT"
	caseData.PalletLayerAmount = itemData.PalletLayerAmount
	caseData.PalletUnitsPerLayer = itemData.PalletUnitsPerLayer
	caseData.PalletSendingUnitAmount = itemData.PalletSendingUnitAmount

	return caseData
}
