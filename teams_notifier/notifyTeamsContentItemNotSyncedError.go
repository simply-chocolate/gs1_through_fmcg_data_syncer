package teams_notifier

import (
	"fmt"
	"os"

	goteamsnotify "github.com/atc0005/go-teams-notify/v2"
	"github.com/atc0005/go-teams-notify/v2/messagecard"
)

func SendContentItemNotSyncedErrorToTeams(ItemCodeFV string, ItemCodeHF string, BarcodeForHF string, GS1StatusContentITem string) error {
	client := goteamsnotify.NewTeamsClient()
	webhook := os.Getenv("TEAMS_WEBHOOK_URL")

	card := messagecard.NewMessageCard()
	card.Title = "Content Item Not Synced Error"
	card.Text = fmt.Sprintf("Script ran into an Error.<BR/>"+
		"**ItemCode for FV**: %s<BR/>"+
		"**HF ItemCode**:%s<BR/>"+
		"**Barcode for HF**:%s<BR/>"+
		"**GS1 Status in SAP**:%s<BR/>"+
		"The FV that the Barcode in the SAP field 'EAN nummer (Stk) on the HF references, has not yet been synced to SAP'<BR/>", ItemCodeFV, ItemCodeHF, BarcodeForHF, GS1StatusContentITem)

	if err := client.Send(webhook, card); err != nil {
		fmt.Printf("ItemcodeFV: %s  ItemCodeHF: %s\n", ItemCodeFV, ItemCodeHF)
		return fmt.Errorf("SendContentNotSyncedErrorToTeams failed to send the error. Error: %v\n", err)
	}
	return nil
}
