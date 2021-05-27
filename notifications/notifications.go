package notifications

import (
	"context"
	"encoding/json"
	"net/http"
	"nextnotes/utils"

	"firebase.google.com/go/messaging"
)

func Notifications(w http.ResponseWriter, r *http.Request) {
	utils.EnableCors(&w)
	app, err := utils.GetFirebaseApp()
	_ = app
	_ = err

	client, err := app.Messaging(context.Background())
	message := &messaging.Message{
		Notification: &messaging.Notification{
			Title: "Price drop",
			Body:  "5% off all electronics",
		},
	}
	_ = client
	_ = err

	response, err := client.Send(context.Background(), message)
	_ = response
	_ = err

	switch r.Method {
	case http.MethodGet:
		json.NewEncoder(w).Encode("")
		return
	case http.MethodPost:
		json.NewEncoder(w).Encode("")
		return
	case http.MethodOptions:
		return
	default:
		http.Error(w, "405 - Method Not Allowed", http.StatusMethodNotAllowed)
	}
}
