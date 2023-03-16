package utils

import (
	"os"

	"github.com/joho/godotenv"
)

func LoadEnvs() {
	godotenv.Load()
	if os.Getenv("SAP_DB_NAME") == "" {
		panic("Error loading environment variable SAP_DB_NAME")
	}
	if os.Getenv("SAP_UN") == "" {
		panic("Error loading environment variable SAP_UN")
	}
	if os.Getenv("SAP_PW") == "" {
		panic("Error loading environment variable SAP_PW")
	}
	if os.Getenv("SAP_URL") == "" {
		panic("Error loading environment variable SAP_URL")
	}
	if os.Getenv("FMCG_ADDRESS") == "" {
		panic("Error loading environment variable FMCG_ADDRESS")
	}
	if os.Getenv("FMCG_USERNAME") == "" {
		panic("Error loading environment variable FMCG_USERNAME")
	}
	if os.Getenv("FMCG_PASSWORD") == "" {
		panic("Error loading environment variable FMCG_PASSWORD")
	}
	if os.Getenv("TEAMS_WEBHOOK_URL") == "" {
		panic("Error loading environment variable TEAMS_WEBHOOK_URL")
	}
}
