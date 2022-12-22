package sap_api_wrapper

import (
	"fmt"
	"net/http"
	"time"

	"github.com/imroc/req/v3"
)

var authCookiesCache []*http.Cookie
var cacheExpiresAt time.Time

func GetSapApiAuthClient() (*req.Client, error) {
	if authCookiesCache == nil || cacheExpiresAt.Before(time.Now()) {
		loginRes, err := SapApiPostLogin()
		if err != nil {
			fmt.Println(err)
			panic("Error getting an authenticated client from SAP. Please investigate")
		}

		authCookiesCache = loginRes.Cookies
		cacheExpiresAt = time.Now().Add((loginRes.Body.SessionTimeout - 1) * time.Minute)
	}

	return GetSapApiBaseClient().SetCommonCookies(authCookiesCache...), nil
}
