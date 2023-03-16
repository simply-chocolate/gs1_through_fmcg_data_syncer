package main

import (
	"fmt"
	"gs1_syncer/sap_api_wrapper"
	"gs1_syncer/teams_notifier"
	"gs1_syncer/utils"
	"time"

	gocron "github.com/go-co-op/gocron"
)

func main() {
	utils.LoadEnvs()
	fmt.Println("Starting the Script V2")
	fmt.Println(time.Now())

	s := gocron.NewScheduler(time.UTC)
	_, _ = s.Cron("03 11 * * *").SingletonMode().Do(func() {
		fmt.Printf("%v Started the Script V2 \n", time.Now().Format("2006-01-02 15:04:05"))

		err := utils.MapData()
		if err != nil {
			teams_notifier.SendUnknownErrorToTeams(err)
		}

		sap_api_wrapper.SapApiPostLogout()
		fmt.Printf("%v Success \n", time.Now().Format("2006-01-02 15:04:05"))
	})
	s.StartBlocking()
}
