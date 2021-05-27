package utils

import (
	"context"
	"log"
	"net/http"

	"cloud.google.com/go/firestore"
	firebase "firebase.google.com/go"
	"google.golang.org/api/option"
)

func GetFirebaseApp(ctx context.Context) (*firebase.App, error) {
	sa := option.WithCredentialsFile("credentials.json")
	app, err := firebase.NewApp(ctx, nil, sa)
	return app, err
}

func GetClientFirestore(ctx context.Context) *firestore.Client {
	app, err := GetFirebaseApp(ctx)

	if err != nil {
		log.Fatalln(err)
	}

	client, err := app.Firestore(ctx)
	if err != nil {
		log.Fatalln(err)
	}
	return client
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
