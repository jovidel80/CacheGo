package db

import (
	"context"
	"errors"
	"fmt"
	"os"
	"testing"

	"cloud.google.com/go/firestore"
	firebaseLib "firebase.google.com/go"
	"google.golang.org/api/option"
)

func StubFirebaseAppInit(ctx context.Context, sa option.ClientOption) (*firebaseLib.App, error) {
	return &firebaseLib.App{}, nil
}

// Stub para app.Firestore
func StubFirestoreInit(app *firebaseLib.App, ctx context.Context) (*firestore.Client, error) {
	return &firestore.Client{}, nil
}

func TestNewFirestoreClient_MissingEnvVar(t *testing.T) {
	os.Unsetenv("FIREBASE_CREDENTIALS_CDP_PATH")

	client, err := NewFirestoreClient(context.Background(), StubFirebaseAppInit, StubFirestoreInit)

	if err == nil {
		t.Fatal("Expected an error, but none was received")
	}

	if client != nil {
		t.Fatal("Did not expect a client, but one was received")
	}
}

// Prueba de éxito usando stubs
func TestNewFirestoreClient_Success(t *testing.T) {
	os.Setenv("FIREBASE_CREDENTIALS_CDP_PATH", "mock-path")

	client, err := NewFirestoreClient(context.Background(), StubFirebaseAppInit, StubFirestoreInit)
	if err != nil {
		t.Fatalf("Did not expect an error, but got: %v", err)
	}

	if client == nil {
		t.Fatal("Expected a client, but got nil")
	}
}

func TestNewFirestoreClient_ErrorInitializingApp(t *testing.T) {
	os.Setenv("FIREBASE_CREDENTIALS_CDP_PATH", "mock-path")

	badAppInit := func(ctx context.Context, sa option.ClientOption) (*firebaseLib.App, error) {
		return nil, errors.New("mock error initializing app")
	}

	_, err := NewFirestoreClient(context.Background(), badAppInit, StubFirestoreInit)
	fmt.Println(err)
	if err == nil {
		t.Fatal("Expected an error, but none was received")
	}
}

func TestNewFirestoreClient_ErrorGettingFirestoreClient(t *testing.T) {
	os.Setenv("FIREBASE_CREDENTIALS_CDP_PATH", "mock-path")

	badFirestoreInit := func(app *firebaseLib.App, ctx context.Context) (*firestore.Client, error) {
		return nil, errors.New("mock error getting firestore client")
	}

	_, err := NewFirestoreClient(context.Background(), StubFirebaseAppInit, badFirestoreInit)
	if err == nil {
		t.Fatal("Expected an error, but none was received")
	}
}

func SetClose() {
	closeFunc = func(_ *FirebaseApp) error {
		return nil
	}
}

func TestFirestoreClient_Close(t *testing.T) {
	os.Setenv("FIREBASE_CREDENTIALS_CDP_PATH", "mock-path")
	client, _ := NewFirestoreClient(context.Background(), StubFirebaseAppInit, StubFirestoreInit)
	SetClose()
	err := client.Close()

	if err != nil {
		t.Fatalf("Did not expect an error, but got: %v", err)
	}
}

func SetGetDocFromCollection(ctx context.Context, colName string, docName string) {
	getDoc = func(_ *FirebaseApp, _ context.Context, _ string, _ string) (*firestore.DocumentSnapshot, error) {
		return &firestore.DocumentSnapshot{}, nil
	}
}

func TestFirestoreClient_GetDocFromCollection(t *testing.T) {
	os.Setenv("FIREBASE_CREDENTIALS_CDP_PATH", "mock-path")
	client, _ := NewFirestoreClient(context.Background(), StubFirebaseAppInit, StubFirestoreInit)
	SetGetDocFromCollection(context.Background(), "test", "test")
	_, err := client.GetDocFromCollection(context.Background(), "test", "test")
	if err != nil {
		t.Fatalf("Did not expect an error, but got: %v", err)
	}
}
