package mail

import (
	"fmt"
	"os"

	mailgun "github.com/mailgun/mailgun-go"
)

func TransferSlipUploadNotify(orderID string, slipImageUrl string) (string, error) {

	mg := mailgun.NewMailgun(DOMAIN, API_KEY, PUBLIC_API_KEY)
	message := mg.NewMessage(
		os.Getenv("EMAIL_SENDER"),
		"Transfer Slip Uploaded - "+orderID,
		"Transfer slip is uploaded.",
		ADMIN_EMAILS...,
	)
	html := fmt.Sprint("Customer's uploaded money transfer slip for order ID: ",
		orderID,
		"<br><br>",
		`<img src="`+slipImageUrl+`">`)
	message.SetHtml(html)
	_, id, err := mg.Send(message)
	return id, err
}
