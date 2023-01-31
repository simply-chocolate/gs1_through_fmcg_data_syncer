package teams_notifier

import (
	"fmt"
	"os"

	goteamsnotify "github.com/atc0005/go-teams-notify/v2"
	"github.com/atc0005/go-teams-notify/v2/messagecard"
)

func SendMappingErrorToTeams(GTIN string, errorPlace string, errorMessage string) error {
	client := goteamsnotify.NewTeamsClient()
	webhook := os.Getenv("TEAMS_WEBHOOK_URL")

	card := messagecard.NewMessageCard()
	card.Title = "MappingError"
	card.Text = fmt.Sprintf("Script ran into a mapping error.<BR/>"+
		"**GTIN**: %v<BR/>"+
		"**Error Place**: %v<BR/>"+
		"**Error Message**: %s<BR/>", GTIN, errorPlace, errorMessage)

	if err := client.Send(webhook, card); err != nil {
		return fmt.Errorf("SendValidationErrorToTeams failed to send the error. Error: %v", err)
	}
	return nil
}
