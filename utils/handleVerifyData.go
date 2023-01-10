package utils

import (
	"fmt"
	"gs1_syncer/sap_api_wrapper"
	"time"
)

// Takes in itemData and the list of Status Times and formats them into time.Time for comparison
func handleCheckIfSapUpdateTimeIsNewer(itemData sap_api_wrapper.SapApiItemsData, FMCGProductsStatusTimes map[string]time.Time, GTIN string) (bool, error) {
	LastUpdateTime, err := FormatSAPDateAndSAPTimetoTimeType(itemData.UpdateDate, itemData.UpdateTime)
	if err != nil {
		return false, fmt.Errorf("error formatting SAP update time into time.\n error:%v", err)
	}

	FMCGUpdateTime, timeExists := FMCGProductsStatusTimes[GTIN]
	if !timeExists {
		// This would mean that the item haven't been added to FMCG yet, so it should return true.
		return true, nil
	}

	fmt.Printf("now comparing times: %v and %v\n", FMCGUpdateTime, LastUpdateTime)
	fmt.Println(FMCGUpdateTime.After(LastUpdateTime))

	if FMCGUpdateTime.After(LastUpdateTime) {
		return false, nil
	}

	return true, nil
}
