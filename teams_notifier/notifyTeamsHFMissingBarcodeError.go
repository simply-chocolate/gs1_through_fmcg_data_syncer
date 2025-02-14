package teams_notifier

import (
	"fmt"
	"os"

	goteamsnotify "github.com/atc0005/go-teams-notify/v2"
	"github.com/atc0005/go-teams-notify/v2/messagecard"
)

func SendHFMissingBarcodeErrorToTeams(ItemCodeFV string, ItemCodeHF string) error {
	client := goteamsnotify.NewTeamsClient()
	webhook := os.Getenv("TEAMS_WEBHOOK_URL")

	card := messagecard.NewMessageCard()
	card.Title = "HF Missing Barcode Error"
	card.Text = fmt.Sprintf("Script ran into an error.<BR/>"+
		"**ItemCode for FV**: %s<BR/>"+
		"**HF ItemCode**:%s<BR/>"+
		"The HF item is missing a barcode in the SAP field 'EAN nummer (Stk)'<BR/>", ItemCodeFV, ItemCodeHF)

	if err := client.Send(webhook, card); err != nil {
		return fmt.Errorf("SendHFMissingBarcodeErrorToTeams failed to send the error. Error: %v\n", err)
	}
	return nil
}
