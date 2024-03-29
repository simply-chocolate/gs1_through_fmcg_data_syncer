package utils

import (
	"gs1_syncer/fmcg_api_wrapper"
	"gs1_syncer/sap_api_wrapper"
)

func MapAllergens(baseItemData fmcg_api_wrapper.FmcgProductBodyBaseItem, itemData sap_api_wrapper.SapApiItemsData) fmcg_api_wrapper.FmcgProductBodyBaseItem {
	baseItemData.AllergenGluten = "AW"
	baseItemData.ContainmentLevelGluten = FormatContainmentLevelSapToFmcg(itemData.ContainmentLevelGluten)
	baseItemData.AllergenCrustacea = "AC"
	baseItemData.ContainmentLevelCrustacea = FormatContainmentLevelSapToFmcg(itemData.ContainmentLevelCrustacea)
	baseItemData.AllergenEgg = "AE"
	baseItemData.ContainmentLevelEgg = FormatContainmentLevelSapToFmcg(itemData.ContainmentLevelEgg)
	baseItemData.AllergenFish = "AF"
	baseItemData.ContainmentLevelFish = FormatContainmentLevelSapToFmcg(itemData.ContainmentLevelFish)
	baseItemData.AllergenPeanut = "AP"
	baseItemData.ContainmentLevelPeanut = FormatContainmentLevelSapToFmcg(itemData.ContainmentLevelPeanut)
	baseItemData.AllergenSoy = "AY"
	baseItemData.ContainmentLevelSoy = FormatContainmentLevelSapToFmcg(itemData.ContainmentLevelSoy)
	baseItemData.AllergenMilk = "AM"
	baseItemData.ContainmentLevelMilk = FormatContainmentLevelSapToFmcg(itemData.ContainmentLevelMilk)
	baseItemData.AllergenAlmonds = "SA"
	baseItemData.ContainmentLevelAlmonds = FormatContainmentLevelSapToFmcg(itemData.ContainmentLevelAlmonds)
	baseItemData.AllergenHazelnut = "SH"
	baseItemData.ContainmentLevelHazelnut = FormatContainmentLevelSapToFmcg(itemData.ContainmentLevelHazelnut)
	baseItemData.AllergenWalnut = "SW"
	baseItemData.ContainmentLevelWalnut = FormatContainmentLevelSapToFmcg(itemData.ContainmentLevelWalnut)
	baseItemData.AllergenCashew = "SC"
	baseItemData.ContainmentLevelCashew = FormatContainmentLevelSapToFmcg(itemData.ContainmentLevelCashew)
	baseItemData.AllergenPecan = "SP"
	baseItemData.ContainmentLevelPecan = FormatContainmentLevelSapToFmcg(itemData.ContainmentLevelPecan)
	baseItemData.AllergenBrazilNut = "SR"
	baseItemData.ContainmentLevelBrazilNut = FormatContainmentLevelSapToFmcg(itemData.ContainmentLevelBrazilNut)
	baseItemData.AllergenPistachio = "ST"
	baseItemData.ContainmentLevelPistachio = FormatContainmentLevelSapToFmcg(itemData.ContainmentLevelPistachio)
	baseItemData.AllergenQueenslandNut = "SM"
	baseItemData.ContainmentLevelQueenslandNut = FormatContainmentLevelSapToFmcg(itemData.ContainmentLevelQueenslandNut)
	baseItemData.AllergenCelery = "BC"
	baseItemData.ContainmentLevelCelery = FormatContainmentLevelSapToFmcg(itemData.ContainmentLevelCelery)
	baseItemData.AllergenMustard = "BM"
	baseItemData.ContainmentLevelMustard = FormatContainmentLevelSapToFmcg(itemData.ContainmentLevelMustard)
	baseItemData.AllergenSulfurDioxideAndSulfites = "AU"
	baseItemData.ContainmentLevelSulfurDioxideAndSulfites = FormatContainmentLevelSapToFmcg(itemData.ContainmentLevelSulfurDioxideAndSulfites)
	baseItemData.AllergenSesameSeeds = "AS"
	baseItemData.ContainmentLevelSesameSeeds = FormatContainmentLevelSapToFmcg(itemData.ContainmentLevelSesameSeeds)
	baseItemData.AllergenLupine = "NL"
	baseItemData.ContainmentLevelLupine = FormatContainmentLevelSapToFmcg(itemData.ContainmentLevelLupine)
	baseItemData.AllergenMollusks = "UM"
	baseItemData.ContainmentLevelMollusks = FormatContainmentLevelSapToFmcg(itemData.ContainmentLevelMollusks)

	containmentLevelNuts := "FREE_FROM"
	if baseItemData.ContainmentLevelPeanut == "MAY_CONTAIN" ||
		baseItemData.ContainmentLevelAlmonds == "MAY_CONTAIN" ||
		baseItemData.ContainmentLevelHazelnut == "MAY_CONTAIN" ||
		baseItemData.ContainmentLevelWalnut == "MAY_CONTAIN" ||
		baseItemData.ContainmentLevelCashew == "MAY_CONTAIN" ||
		baseItemData.ContainmentLevelPecan == "MAY_CONTAIN" ||
		baseItemData.ContainmentLevelBrazilNut == "MAY_CONTAIN" ||
		baseItemData.ContainmentLevelPistachio == "MAY_CONTAIN" ||
		baseItemData.ContainmentLevelQueenslandNut == "MAY_CONTAIN" {
		containmentLevelNuts = "MAY_CONTAIN"
	}
	if baseItemData.ContainmentLevelPeanut == "CONTAINS" ||
		baseItemData.ContainmentLevelAlmonds == "CONTAINS" ||
		baseItemData.ContainmentLevelHazelnut == "CONTAINS" ||
		baseItemData.ContainmentLevelWalnut == "CONTAINS" ||
		baseItemData.ContainmentLevelCashew == "CONTAINS" ||
		baseItemData.ContainmentLevelPecan == "CONTAINS" ||
		baseItemData.ContainmentLevelBrazilNut == "CONTAINS" ||
		baseItemData.ContainmentLevelPistachio == "CONTAINS" ||
		baseItemData.ContainmentLevelQueenslandNut == "CONTAINS" {
		containmentLevelNuts = "CONTAINS"
	}
	baseItemData.AllergenNuts = "AN"
	baseItemData.ContainmentLevelNuts = containmentLevelNuts

	return baseItemData
}

func MapNutritionalInformation(baseItemData fmcg_api_wrapper.FmcgProductBodyBaseItem, itemData sap_api_wrapper.SapApiItemsData) fmcg_api_wrapper.FmcgProductBodyBaseItem {
	baseItemData.PreparationState = "UNPREPARED"
	baseItemData.NutritionalReferenceValue = 100
	baseItemData.NutritionalReferenceUOM = "GRM"

	baseItemData.EnergyInKcal = itemData.EnergyInKcal
	baseItemData.EnergyInkJ = itemData.EnergyInkJ
	baseItemData.EnergyInKcalPrecision = "APPROXIMATELY"
	baseItemData.NutritionalFat = "FAT"
	baseItemData.NutritionalFatPrecision = "APPROXIMATELY"
	baseItemData.NutritionalFatValue = itemData.NutritionalFatValue
	baseItemData.NutritionalFatUOM = "GRM"

	baseItemData.NutritionalFattyAcids = "FASAT"
	baseItemData.NutritionalFattyAcidsPrecision = "APPROXIMATELY"
	baseItemData.NutritionalFattyAcidsValue = itemData.NutritionalFattyAcidsValue
	baseItemData.NutritionalFattyUOM = "GRM"

	baseItemData.NutritionalCarboHydrates = "CHOAVL"
	baseItemData.NutritionalCarboHydratesPrecision = "APPROXIMATELY"
	baseItemData.NutritionalCarboHydratesValue = itemData.NutritionalCarboHydratesValue
	baseItemData.NutritionalCarboHydratesUOM = "GRM"

	baseItemData.NutritionalSugar = "SUGAR-"
	baseItemData.NutritionalSugarPrecision = "APPROXIMATELY"
	baseItemData.NutritionalSugarValue = itemData.NutritionalSugarValue
	baseItemData.NutritionalSugarUOM = "GRM"

	baseItemData.NutritionalProtein = "PRO-"
	baseItemData.NutritionalProteinPrecision = "APPROXIMATELY"
	baseItemData.NutritionalProteinValue = itemData.NutritionalProteinValue
	baseItemData.NutritionalProteinUOM = "GRM"

	baseItemData.NutritionalSalt = "SALTEQ"
	baseItemData.NutritionalSaltPrecision = "APPROXIMATELY"
	baseItemData.NutritionalSaltValue = itemData.NutritionalSaltValue
	baseItemData.NutritionalSaltUOM = "GRM"

	baseItemData.ListOfIngredientsLanguageCodeDA = "da"
	// CAP THIS TO 5000 CHARACTERS
	if len(itemData.ListOfIngredientsDA) > 5000 {
		baseItemData.ListOfIngredientsDA = itemData.ListOfIngredientsDA[:5000]
	} else {
		baseItemData.ListOfIngredientsDA = itemData.ListOfIngredientsDA
	}

	return baseItemData
}

func MapLogisticalInformation(baseItemData fmcg_api_wrapper.FmcgProductBodyBaseItem, itemData sap_api_wrapper.SapApiItemsData) fmcg_api_wrapper.FmcgProductBodyBaseItem {
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
	baseItemData.PackagingType = itemData.PackagingType

	return baseItemData
}
