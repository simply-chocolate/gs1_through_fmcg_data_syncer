package utils

func MapPackagingType(sapPackagingType string) string {
	switch sapPackagingType {
	case "1", "4", "5", "6":
		return "BX"
	case "2":
		return "WRP"
	case "3":
		return "BG"
	case "7":
		return "JR"
	case "8":
		return "CNG"
	case "JR":
		return "JR"
	case "BX":
		return "BX"
	case "WRP":
		return "WRP"
	default:
		return "BX"
	}
}
