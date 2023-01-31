package utils

import (
	"fmt"
	"gs1_syncer/sap_api_wrapper"
)

func GetItemDataFromSap() (sap_api_wrapper.SapApiGetItemsDataResults, error) {
	resp, err := sap_api_wrapper.SapApiGetItemsData_AllPages(sap_api_wrapper.SapApiQueryParams{
		Filter: "U_CCF_Sync_GS1 eq 'Y'",
	})
	if err != nil {
		return sap_api_wrapper.SapApiGetItemsDataResults{}, err
	}

	return *resp.Body, nil
}

func GetAttemptedItemsFromSap() (sap_api_wrapper.SapApiGetItemsDataResults, error) {
	resp, err := sap_api_wrapper.SapApiGetItemsData_AllPages(sap_api_wrapper.SapApiQueryParams{
		Filter: "U_CCF_GS1_Status ne 'OK'",
	})
	if err != nil {
		return sap_api_wrapper.SapApiGetItemsDataResults{}, err
	}

	return *resp.Body, nil
}

func GetMixCaseItemsFromSap(itemCode string) (sap_api_wrapper.SapApiGetMixCaseDataResult, error) {
	resp, err := sap_api_wrapper.SapApiGetMixCaseData_AllPages(sap_api_wrapper.SapApiQueryParams{
		FatherItemCode: fmt.Sprintf("'%s'", itemCode),
	})
	if err != nil {
		return sap_api_wrapper.SapApiGetMixCaseDataResult{}, err
	}

	return *resp.Body, nil
}

func GetMixContentItemInfoFromSap(itemCode string) (sap_api_wrapper.SapApiGetMixCaseContentResult, error) {
	resp, err := sap_api_wrapper.SapApiGetMixCaseContent_AllPages(sap_api_wrapper.SapApiQueryParams{
		Filter: fmt.Sprintf("ItemCode eq '%s'", itemCode),
	})
	if err != nil {
		return sap_api_wrapper.SapApiGetMixCaseContentResult{}, err
	}

	return *resp.Body, nil
}
