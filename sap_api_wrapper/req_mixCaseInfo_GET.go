package sap_api_wrapper

import (
	"fmt"
	"time"
)

type SapApiGetMixCaseDataResult struct {
	Value    []SapApiMixCaseData `json:"value"`
	NextLink string              `json:"odata.nextLink"`
}

type SapApiMixCaseData struct {
	ItemCode string  `json:"ItemCode"`
	Quantity float64 `json:"Quantity"`
}

type SapApiGetMixCaseDataReturn struct {
	Body *SapApiGetMixCaseDataResult
}

func SapApiGetMixCaseData(params SapApiQueryParams) (SapApiGetMixCaseDataReturn, error) {
	for i := 0; i < 200; i++ {
		client, err := GetSapApiAuthClient()
		if err != nil {
			fmt.Println("Error getting an authenticaed client")
			return SapApiGetMixCaseDataReturn{}, err
		}

		resp, err := client.
			//DevMode().
			R().
			SetResult(SapApiGetMixCaseDataResult{}).
			SetQueryParams(params.AsReqParams()).
			Get("SQLQueries('CQ10002')/List")
		if err != nil {
			fmt.Println(err)
			return SapApiGetMixCaseDataReturn{}, err
		}

		if resp.IsError() {
			if resp.StatusCode != 403 {
				//fmt.Printf("Dumping SAP Error %v\n", resp.Dump())
				return SapApiGetMixCaseDataReturn{}, fmt.Errorf("error getting mixCaseContent from to sap. unexpected errorcode. StatusCode :%v Status: %v", resp.StatusCode, resp.Status)
			} else {
				time.Sleep(100 * time.Millisecond)
			}

		} else {
			return SapApiGetMixCaseDataReturn{
				Body: resp.Result().(*SapApiGetMixCaseDataResult),
			}, nil
		}
	}
	return SapApiGetMixCaseDataReturn{}, fmt.Errorf("error getting invoices from SAP. Tried 200 times and couldn't get through")
}

func SapApiGetMixCaseData_AllPages(params SapApiQueryParams) (SapApiGetMixCaseDataReturn, error) {
	res := SapApiGetMixCaseDataResult{}
	for page := 0; ; page++ {
		params.Skip = page * 20

		getItemsRes, err := SapApiGetMixCaseData(params)
		if err != nil {
			return SapApiGetMixCaseDataReturn{}, err
		}

		res.Value = append(res.Value, getItemsRes.Body.Value...)

		if getItemsRes.Body.NextLink == "" {
			break
		}
	}

	return SapApiGetMixCaseDataReturn{
		Body: &res,
	}, nil
}
