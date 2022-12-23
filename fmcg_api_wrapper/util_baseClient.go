package fmcg_api_wrapper

import (
	"os"

	"github.com/imroc/req/v3"
)

func GetFMCGApiBaseClient() *req.Client {
	return req.C().
		SetBaseURL(os.Getenv("FMCG_ADDRESS")).
		SetCommonBasicAuth(os.Getenv("FMCG_USERNAME"), os.Getenv("FMCG_PASSWORD")).
		SetCommonHeader("Content-Type", "application/json")
}
