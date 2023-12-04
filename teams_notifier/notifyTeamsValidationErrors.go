package teams_notifier

import (
	"fmt"
	"os"

	goteamsnotify "github.com/atc0005/go-teams-notify/v2"
	"github.com/atc0005/go-teams-notify/v2/messagecard"
)

func SendValidationErrorToTeams(ItemCode string, GTIN string, fieldId string, fieldLabel string, message string, messageType string) error {
	client := goteamsnotify.NewTeamsClient()
	webhook := os.Getenv("TEAMS_WEBHOOK_URL")

	card := messagecard.NewMessageCard()
	card.Title = "ValidationError"
	card.Text = fmt.Sprintf("Script ran into a validation Error.<BR/>"+
		"**ItemCode**: %v<BR/>"+
		"**GTIN**: %v<BR/>"+
		"**Field ID**: %s<BR/>"+
		"**Field Label**:%s<BR/>"+
		"**Message** %s<BR/>"+
		"**Message Type** %s <BR/>", ItemCode, GTIN, fieldId, fieldLabel, message, messageType)

	if err := client.Send(webhook, card); err != nil {
		return fmt.Errorf("SendValidationErrorToTeams failed to send the error. Error: %v", err)
	}
	return nil
}
