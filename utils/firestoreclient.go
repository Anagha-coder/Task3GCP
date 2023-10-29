package utils

import (
	"context"

	"cloud.google.com/go/firestore"
	"google.golang.org/api/option"
)

func CreateFirestoreClient() (*firestore.Client, error) {
	ctx := context.Background()

	opt := option.WithCredentialsFile("C:/Users/adesa/Downloads/userServiceAcckey.json")

	client, err := firestore.NewClient(ctx, "task3gcp", opt)
	if err != nil {
		return nil, err
	}

	return client, nil
}
