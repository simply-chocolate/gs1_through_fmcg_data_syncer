package utils

import (
	"fmt"
	"time"
)

func FormatContainmentLevelSapToFmcg(sapContainmentLevel string) string {
	switch sapContainmentLevel {
	case "Free from":
		return "FREE_FROM"
	case "May contain traces of":
		return "MAY_CONTAIN"
	case "In product":
		return "CONTAINS"
	default:
		return "ERROR"
	}
}

// Formats a date string in YYYY-mm-DD format to DD-mm-YYYY
func FormatSapDateToFMCHDate(sapDate string) (string, error) {
	if sapDate == "" {
		sapDate = "2020-01-01"
	}

	newDate, err := time.Parse("2006-01-02", sapDate)
	if err != nil {
		return "", fmt.Errorf("error parsing date: %v to fmcg format. Error:%v", sapDate, err)
	}

	return newDate.Format("02-01-2006"), nil
}
