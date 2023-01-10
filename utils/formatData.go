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
func FormatSapDateToFMCGDate(sapDate string) (string, error) {
	if sapDate == "" {
		sapDate = "2020-01-01"
	}

	formattedDate, err := time.Parse("2006-01-02", sapDate)
	if err != nil {
		return "", fmt.Errorf("error parsing date: %v to fmcg format. Error:%v", sapDate, err)
	}

	return formattedDate.Format("02-01-2006"), nil
}

// Formats a date string in from the FMCG API to time.Time format for comparison
func FormatFMCGDateToTimeType(fcmgDate string) (time.Time, error) {
	formattedDate, err := time.Parse("2006-01-02T15:04:05.000-07:00", fcmgDate)

	if err != nil {
		return time.Time{}, fmt.Errorf("error parsing date: %v to SAP format. Error:%v", fcmgDate, err)
	}

	return formattedDate, nil
}

func FormatSAPDateAndSAPTimetoTimeType(sapDate string, sapTime string) (time.Time, error) {
	concattedDateTime := fmt.Sprintf("%vT%v", sapDate, sapTime)
	formattedDate, err := time.Parse("2006-01-02T15:04:05", concattedDateTime)

	if err != nil {
		return time.Time{}, fmt.Errorf("error parsing date: %v to SAP format. Error:%v", concattedDateTime, err)
	}

	return formattedDate, nil
}
