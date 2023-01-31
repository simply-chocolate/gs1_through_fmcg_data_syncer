package main

import (
	"fmt"
	"gs1_syncer/sap_api_wrapper"
	"gs1_syncer/teams_notifier"
	"gs1_syncer/utils"
	"log"
	"time"

	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}
	fmt.Printf("%v Started the Script \n", time.Now().Format("2006-01-02 15:04:05"))

	err = utils.MapData()
	if err != nil {
		teams_notifier.SendUnknownErrorToTeams(err)
	}

	sap_api_wrapper.SapApiPostLogout()

	fmt.Printf("%v Success \n", time.Now().Format("2006-01-02 15:04:05"))

}
