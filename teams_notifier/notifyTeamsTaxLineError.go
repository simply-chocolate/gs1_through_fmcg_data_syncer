package teams_notifier

import (
	"encoding/json"
	"fmt"
	"os"

	goteamsnotify "github.com/atc0005/go-teams-notify/v2"
)

func SendTaxLineErrorToTeams(orderNumber json.Number, orderId json.Number, productId json.Number) {
	client := goteamsnotify.NewClient()
	webhook := os.Getenv("TEAMS_WEBHOOK_URL")

	card := goteamsnotify.NewMessageCard()
	card.Title = "TaxLine"
	card.Text = fmt.Sprintf("Script failed to import an order because there was a problem with taxes on an item.<BR/>"+
		"**Order number**: %v<BR/>"+
		"**Order Link**: https://simply-chocolate-copenhagen.myshopify.com/admin/orders/%v<BR/>"+
		"**Product Link**: https://simply-chocolate-copenhagen.myshopify.com/admin/products/%v", orderNumber, orderId, productId)

	if err := client.Send(webhook, card); err != nil {
		fmt.Println("SendOrderErrorToTeams failed to send the error.")
	}

}
