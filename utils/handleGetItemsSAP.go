package utils

import (
	"gs1_syncer/sap_api_wrapper"
)

func GetItemDataFromSap() (sap_api_wrapper.SapApiGetItemsDataResults, error) {
	resp, err := sap_api_wrapper.SapApiGetItemsData_AllPages(sap_api_wrapper.SapApiQueryParams{
		Filter: "ItemCode eq '0021050001'",
		//Filter: "U_CCF_GS1_Sync eq Y",
	})
	if err != nil {
		return sap_api_wrapper.SapApiGetItemsDataResults{}, err
	}

	return *resp.Body, nil
}
