package teams_notifier

import (
	"encoding/json"
	"fmt"
	"os"

	goteamsnotify "github.com/atc0005/go-teams-notify/v2"
)

func SendOrderErrorToTeams(orderNumber json.Number, orderId json.Number) {
	client := goteamsnotify.NewClient()
	webhook := os.Getenv("TEAMS_WEBHOOK_URL")

	card := goteamsnotify.NewMessageCard()
	card.Title = "Order Error"
	card.Text = fmt.Sprintf("Script failed to import an order.<BR/>"+"**Order number**: %v<BR/>"+"**Link**: https://simply-chocolate-copenhagen.myshopify.com/admin/orders/%v", orderNumber, orderId)

	if err := client.Send(webhook, card); err != nil {
		fmt.Println("SendOrderErrorToTeams failed to send the error.")
	}

}
