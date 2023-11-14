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

	fmt.Printf("%v: Started the Script\n", time.Now().UTC().Format("2006-01-02 15:04:05"))
	err := utils.MapData()
	if err != nil {
		teams_notifier.SendUnknownErrorToTeams(err)
	}
	fmt.Printf("%v: Success\n", time.Now().UTC().Format("2006-01-02 15:04:05"))
	fmt.Printf("%v: Started the Cron Scheduler\n", time.Now().UTC().Format("2006-01-02 15:04:05"))

	s := gocron.NewScheduler(time.UTC)
	_, _ = s.Cron("0 * * * *").SingletonMode().Do(func() {
		fmt.Printf("%v: Started the Script\n", time.Now().UTC().Format("2006-01-02 15:04:05"))

		err := utils.MapData()
		if err != nil {
			teams_notifier.SendUnknownErrorToTeams(err)
		}

		sap_api_wrapper.SapApiPostLogout()
		fmt.Printf("%v: Success\n", time.Now().UTC().Format("2006-01-02 15:04:05"))
	})
	s.StartBlocking()
}
