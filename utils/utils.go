package utils

import (
	"context"
	"log"
	"net/http"

	"cloud.google.com/go/firestore"
	firebase "firebase.google.com/go"
	"google.golang.org/api/option"
)

func GetFirebaseApp() (*firebase.App, error) {
	ctx := context.Background()
	sa := option.WithCredentialsFile("credentials.json")
	app, err := firebase.NewApp(ctx, nil, sa)
	return app, err
}

func GetClientFirestore() (context.Context, *firestore.Client) {
	ctx := context.Background()
	app, err := GetFirebaseApp()

	if err != nil {
		log.Fatalln(err)
	}

	client, err := app.Firestore(ctx)
	if err != nil {
		log.Fatalln(err)
	}
	return ctx, client
}

func EnableCors(w *http.ResponseWriter) {
	(*w).Header().Set(
		"Access-Control-Allow-Origin", "*",
	)
	(*w).Header().Set(
		"Access-Control-Allow-Methods",
		"POST, GET, OPTIONS, PUT, DELETE",
	)
	(*w).Header().Set(
		"Access-Control-Allow-Headers",
		"Accept, Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization",
	)
}
