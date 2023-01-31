package teams_notifier

import (
	"fmt"
	"os"

	goteamsnotify "github.com/atc0005/go-teams-notify/v2"
	"github.com/atc0005/go-teams-notify/v2/messagecard"
)

func SendValidationErrorToTeams(GTIN string, fieldId string, fieldLabel string, message string, messageType string, postBody string) error {
	fmt.Println("SendValidationErrorToTeams", GTIN, fieldId, fieldLabel, message, messageType, postBody)
	client := goteamsnotify.NewTeamsClient()
	webhook := os.Getenv("TEAMS_WEBHOOK_URL")

	card := messagecard.NewMessageCard()
	card.Title = "ValidationError"
	card.Text = fmt.Sprintf("Script ran into a validation Error.<BR/>"+"**GTIN**: %v<BR/>"+
		"**Field ID**: %s<BR/>"+
		"**Field Label**:%s<BR/>"+
		"**Message** %s<BR/>"+
		"**Message Type** %s <BR/>"+
		"**Post Body** %s<BR/>", GTIN, fieldId, fieldLabel, message, messageType, postBody)

	if err := client.Send(webhook, card); err != nil {
		return fmt.Errorf("SendValidationErrorToTeams failed to send the error. Error: %v", err)
	}
	return nil
}
