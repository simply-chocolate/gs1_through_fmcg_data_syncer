package teams_notifier

import (
	"fmt"
	"os"

	goteamsnotify "github.com/atc0005/go-teams-notify/v2"
	"github.com/atc0005/go-teams-notify/v2/messagecard"
)

func SendProductStatusErrorToTeams(GTIN string, errorPlace string, errorMessage string) error {
	client := goteamsnotify.NewTeamsClient()
	webhook := os.Getenv("TEAMS_WEBHOOK_URL")

	card := messagecard.NewMessageCard()
	card.Title = "GS1 Status error"
	card.Text = fmt.Sprintf("Script ran into an error setting the GS1 status .<BR/>"+
		"**GTIN**: %v<BR/>"+
		"**Error Place**: %v<BR/>"+
		"**Error Message**: %s<BR/>", GTIN, errorPlace, errorMessage)

	if err := client.Send(webhook, card); err != nil {
		fmt.Printf("GTIN: %v  ErrorPlace: %s ErrorMessage: %s\n", GTIN, errorPlace, errorMessage)
		return fmt.Errorf("SendValidationErrorToTeams failed to send the error. Error: %v\n", err)
	}
	return nil
}
