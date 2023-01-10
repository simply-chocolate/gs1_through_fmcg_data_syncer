package utils

import (
	"gs1_syncer/fmcg_api_wrapper"
	"time"
)

func GetAllProductsStatusFMCG() (map[string]time.Time, error) {
	resp, err := fmcg_api_wrapper.FMCGApiGetAllProductStatus(0)
	if err != nil {
		return map[string]time.Time{}, err
	}
	LastModifiedDates := map[string]time.Time{}

	for _, products := range resp.Body.Products {
		GTIN := products.ProductId[0:14]
		LastModifiedDate, err := FormatFMCGDateToTimeType(products.LastModified)
		if err != nil {
			return map[string]time.Time{}, err
		}
		LastModifiedDates[GTIN] = LastModifiedDate
	}
	return LastModifiedDates, nil
}
