package utils

import (
	"fmt"
	"gs1_syncer/teams_notifier"
	"time"
)

// Calls the APIs and retrieves the information needed to handle the integration of data
func MapData() error {

	// Check if we need to update the status of the products in SAP
	AttemptedSapItemsData, err := GetAttemptedItemsFromSap()
	if err != nil {
		fmt.Println("Couldn't get attempted items from SAP. Sleeping 10 minutes")
		time.Sleep(10 * time.Minute)
		AttemptedSapItemsData, err = GetAttemptedItemsFromSap()
		if err != nil {
			return fmt.Errorf("error getting the attempted items from SAP: %v", err)
		}
	}
	err = UpdateAttemptedItemsStatus(AttemptedSapItemsData)
	if err != nil {
		teams_notifier.SendUnknownErrorToTeams(err)
	}

	// Then we get the items from SAP that has been requested to be updated
	SapItemsData, err := GetItemDataFromSap()
	if err != nil {
		fmt.Println("Couldn't get ItemData from SAP. Sleeping 10 minutes")
		time.Sleep(10 * time.Minute)
		SapItemsData, err = GetItemDataFromSap()
		if err != nil {
			return fmt.Errorf("error getting the ItemData from SAP: %v", err)
		}
	}

	mixDisplays := IterateProductsAndMapToFMCGFormat(SapItemsData)

	// Then if any of the items are mixCases, we do them afterwards
	err = IterateMixCasesAndMapToFMCGFormat(mixDisplays)
	if err != nil {
		teams_notifier.SendUnknownErrorToTeams(err)
	}

	return nil
}
