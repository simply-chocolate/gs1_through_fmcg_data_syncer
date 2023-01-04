package teams_notifier

import (
	"encoding/json"
	"fmt"
	"os"

	goteamsnotify "github.com/atc0005/go-teams-notify/v2"
)

func SendOrderSpecificRequestsReturnErrorToTeams(requestName string, requestType string, response string, responseBody string, api string, orderId json.Number) {
	client := goteamsnotify.NewClient()
	webhook := os.Getenv("TEAMS_WEBHOOK_URL")

	card := goteamsnotify.NewMessageCard()
	card.Title = "Request Error"
	card.Text = fmt.Sprintf("**API:** %s <BR/>"+
		"**Request Type:** %s<BR/>"+
		"**Request Name:** %s <BR/>"+
		"**Response **: %s <BR/>"+
		"**ResponseBody **: %s <BR/>"+
		"**Link**: https://simply-chocolate-copenhagen.myshopify.com/admin/orders/%v", api, requestName, requestType, response, responseBody, orderId)

	if err := client.Send(webhook, card); err != nil {
		fmt.Println("SendVatCodeErrorToTeams failed to send the error.")
	}

}
