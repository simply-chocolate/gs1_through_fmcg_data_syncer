package sap_api_wrapper

import (
	"encoding/json"
	"fmt"
	"time"
)

type SapApiGetInvoicesResult struct {
	Value []struct {
		DocNum     json.Number `json:"DocNum"`
		BarCode    string      `json:"CodeBars"`
		OrderRef   string      `json:"NumAtCard"`
		ItemCode   string      `json:"ItemCode"`
		StockPrice float64     `json:"StockPrice"`
		Weight     float64     `json:"Weight1"`
	} `json:"value"`

	NextLink string `json:"odata.nextLink"`
}

type SapApiGetInvoicesReturn struct {
	Body *SapApiGetInvoicesResult
}

func SapApiGetInvoices(params SapApiQueryParams) (SapApiGetInvoicesReturn, error) {
	for i := 0; i < 200; i++ {
		client, err := GetSapApiAuthClient()
		if err != nil {
			fmt.Println("Error getting an authenticaed client")
			return SapApiGetInvoicesReturn{}, err
		}

		resp, err := client.
			//DevMode().
			R().
			SetResult(SapApiGetInvoicesResult{}).
			SetQueryParams(params.AsReqParams()).
			Get("SQLQueries('CQ10001')/List")
		if err != nil {
			return SapApiGetInvoicesReturn{}, err
		}

		if resp.IsError() {
			if resp.StatusCode != 403 {
				fmt.Printf("Dumping SAP Error %v\n", resp.Dump())
				return SapApiGetInvoicesReturn{}, fmt.Errorf("error getting invoices from to sap. unexpected errorcode. StatusCode :%v Status: %v", resp.StatusCode, resp.Status)
			} else {
				time.Sleep(100 * time.Millisecond)
			}

		} else {
			return SapApiGetInvoicesReturn{
				Body: resp.Result().(*SapApiGetInvoicesResult),
			}, nil
		}
	}
	return SapApiGetInvoicesReturn{}, fmt.Errorf("error getting invoices from SAP. Tried 200 times and couldn't get through")
}

func SapApiGetInvoices_AllPages(params SapApiQueryParams) (SapApiGetInvoicesReturn, error) {
	res := SapApiGetInvoicesResult{}
	for page := 0; ; page++ {
		params.Skip = page * 20

		getInvoicesRes, err := SapApiGetInvoices(params)

		if err != nil {
			return SapApiGetInvoicesReturn{}, err
		}

		res.Value = append(res.Value, getInvoicesRes.Body.Value...)

		if getInvoicesRes.Body.NextLink == "" {
			break
		}
	}

	return SapApiGetInvoicesReturn{
		Body: &res,
	}, nil
}
