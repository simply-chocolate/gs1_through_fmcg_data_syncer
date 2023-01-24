package sap_api_wrapper

import (
	"fmt"
)

type SapApiGetMixCaseContentResult struct {
	Value    []SapApiMixCaseContent `json:"value"`
	NextLink string                 `json:"odata.nextLink"`
}

type SapApiMixCaseContent struct {
	ItemCode              string `json:"ItemCode"`
	ItemBarCodeCollection []struct {
		Barcode  string `json:"Barcode"`
		UoMEntry int    `json:"UoMEntry"`
	} `json:"ItemBarCodeCollection"`
	BarCodeForHF string `json:"U_BOYX_EAN"`
}

type SapApiGetMixCaseContentReturn struct {
	Body *SapApiGetMixCaseContentResult
}

func SapApiGetMixCaseContent(params SapApiQueryParams) (SapApiGetMixCaseContentReturn, error) {
	client, err := GetSapApiAuthClient()
	if err != nil {
		fmt.Println("Error getting an authenticaed client")
		return SapApiGetMixCaseContentReturn{}, err
	}

	resp, err := client.
		//DevMode().
		R().
		SetResult(SapApiGetMixCaseContentResult{}).
		SetQueryParams(params.AsReqParams()).
		Get("Items")
	if err != nil {
		fmt.Println(err)
		return SapApiGetMixCaseContentReturn{}, err
	}

	return SapApiGetMixCaseContentReturn{
		Body: resp.Result().(*SapApiGetMixCaseContentResult),
	}, nil

}

func SapApiGetMixCaseContent_AllPages(params SapApiQueryParams) (SapApiGetMixCaseContentReturn, error) {
	res := SapApiGetMixCaseContentResult{}
	for page := 0; ; page++ {
		params.Skip = page * 20

		getItemsRes, err := SapApiGetMixCaseContent(params)
		if err != nil {
			return SapApiGetMixCaseContentReturn{}, err
		}

		res.Value = append(res.Value, getItemsRes.Body.Value...)

		if getItemsRes.Body.NextLink == "" {
			break
		}
	}

	return SapApiGetMixCaseContentReturn{
		Body: &res,
	}, nil
}
