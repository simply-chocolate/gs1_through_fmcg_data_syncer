package teams_notifier

import (
	"encoding/json"
	"fmt"
	"os"

	goteamsnotify "github.com/atc0005/go-teams-notify/v2"
)

func SendLineItemVatCodeErrorToTeams(orderNumber json.Number, orderId json.Number, countryCode string, errorPlace string, barcode string, itemType string) {
	client := goteamsnotify.NewClient()
	webhook := os.Getenv("TEAMS_WEBHOOK_URL")

	card := goteamsnotify.NewMessageCard()
	card.Title = "VAT Error"
	card.Text = fmt.Sprintf("Script ran into a VAT Error.<BR/>"+"**Order number**: %v<BR/>"+
		"**Country code**: %s<BR/>"+
		"**Link**: https://simply-chocolate-copenhagen.myshopify.com/admin/orders/%v<BR/>"+
		"**Error Happened** %s<BR/>"+
		"**Barcode that failed** %s <BR/>"+
		"**Type of item** %s <BR/>", orderNumber, countryCode, orderId, errorPlace, barcode, itemType)

	if err := client.Send(webhook, card); err != nil {
		fmt.Println("SendVatCodeErrorToTeams failed to send the error.")
	}

}
