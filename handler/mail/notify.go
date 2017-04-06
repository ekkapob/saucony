package mail

import (
	"bytes"
	"html/template"
	"os"

	"github.com/ekkapob/saucony/handler"
	"github.com/ekkapob/saucony/model"
	mailgun "github.com/mailgun/mailgun-go"
)

func OrderNotify(
	orderID string,
	cart model.Cart,
	customer model.Customer) (string, error) {

	domain := os.Getenv("MG_DOMAIN")
	apiKey := os.Getenv("MG_API_KEY")
	publicApiKey := os.Getenv("MG_PUBLIC_API_KEY")

	mg := mailgun.NewMailgun(domain, apiKey, publicApiKey)
	message := mg.NewMessage(
		"Saucony Thailand <contact@sauconythailand.com>",
		"Thank you for Your Order - "+orderID,
		"We've received your order.",
		"ekkapob@gmail.com",
	)

	html := orderEmailHtml(orderID, cart, customer)
	message.SetHtml(html)

	_, id, err := mg.Send(message)
	return id, err
}

type OrderTmplData struct {
	OrderID  string
	Cart     model.Cart
	Customer model.Customer
}

func orderEmailHtml(
	orderID string,
	cart model.Cart,
	customer model.Customer) string {

	var content bytes.Buffer
	t := template.New("email").Funcs(handler.BaseFuncMap())
	t = template.Must(t.ParseFiles("templates/emails/order.tmpl"))
	t.ExecuteTemplate(&content, "email", OrderTmplData{orderID, cart, customer})
	return content.String()
}