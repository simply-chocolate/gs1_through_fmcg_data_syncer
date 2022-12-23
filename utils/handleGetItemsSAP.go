package utils

import (
	"gs1_syncer/sap_api_wrapper"
	"strconv"
)

func GetItemsFromSap() (map[string]map[string]string, error) {
	resp, err := sap_api_wrapper.SapApiGetItems_AllPages(sap_api_wrapper.SapApiQueryParams{
		Select: []string{"ItemCode", "ItemBarCodeCollection"},
		Filter: "Valid eq 'Y' and SalesItem eq 'Y'",
	})
	if err != nil {
		return map[string]map[string]string{}, err
	}

	BarCodes := make(map[string]map[string]string)

	BarCodes["POSTKORT"] = map[string]string{
		"ItemCode": "121",
		"UoMEntry": "-1",
	}

	for _, item := range resp.Body.Value {
		if len(item.ItemBarCodeCollection) == 0 {
			continue
		}
		for _, barCodeCollection := range item.ItemBarCodeCollection {
			if _, exists := BarCodes[barCodeCollection.Barcode]; !exists {
				BarCodes[barCodeCollection.Barcode] = map[string]string{
					"ItemCode": item.ItemCode,
					"UoMEntry": strconv.Itoa(barCodeCollection.UoMEntry),
				}
			}
		}
	}

	return BarCodes, nil
}
