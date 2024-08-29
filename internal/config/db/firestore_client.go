package db

import (
	"context"
	"fmt"
	"os"

	"cloud.google.com/go/firestore"
	firebaseLib "firebase.google.com/go"
	"google.golang.org/api/option"
)

type FirebaseApp struct {
	client *firestore.Client
}

type FirebaseAppInitFunc func(ctx context.Context, sa option.ClientOption) (*firebaseLib.App, error)
type FirestoreInitFunc func(app *firebaseLib.App, ctx context.Context) (*firestore.Client, error)

func NewFirestoreClient(ctx context.Context, appInit FirebaseAppInitFunc, firestoreInit FirestoreInitFunc) (ConfigDatabaseClient, error) {
	credentialsFilePath := os.Getenv("FIREBASE_CREDENTIALS_CDP_PATH")
	if credentialsFilePath == "" {
		return nil, fmt.Errorf("FIREBASE_CREDENTIALS_CDP_PATH is not set")
	}

	sa := option.WithCredentialsFile(credentialsFilePath)
	app, err := appInit(ctx, sa)
	if err != nil {
		return nil, fmt.Errorf("error initializing app: %v", err)
	}

	client, err := firestoreInit(app, ctx)
	if err != nil {
		return nil, fmt.Errorf("error getting firestore client: %v", err)

	}
	return &FirebaseApp{client: client}, nil
}

var getDoc = func(f *FirebaseApp, ctx context.Context, colName string, docName string) (*firestore.DocumentSnapshot, error) {
	return f.client.Collection(colName).Doc(docName).Get(ctx)
}

func (f *FirebaseApp) GetDocFromCollection(ctx context.Context, colName string, docName string) (interface{}, error) {
	ref, err := getDoc(f, ctx, colName, docName)
	if err != nil {
		return nil, fmt.Errorf("error getting doc: %v", err)
	}
	return ref, nil
}

var closeFunc = func(f *FirebaseApp) error {
	return f.client.Close()
}

func (f *FirebaseApp) Close() error {
	return closeFunc(f)
}
