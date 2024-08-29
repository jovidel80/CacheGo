package main

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"cloud.google.com/go/firestore"
	firebaseLib "firebase.google.com/go"
	"github.com/joho/godotenv"
	"github.com/jovidel80/cacheGo/internal/config/db"
	"github.com/jovidel80/cacheGo/internal/server"
	"google.golang.org/api/option"
)

type InMemoryCacheDatabase struct {
}

var configDbClient db.ConfigDatabaseClient

func (i *InMemoryCacheDatabase) GetCacheByKey(key string) string {
	dsnap, err := configDbClient.GetDocFromCollection(context.Background(), "config", "proxy-cdp")
	if err != nil {
		log.Fatal(err)
	}
	doc, ok := dsnap.(*firestore.DocumentSnapshot)
	if !ok {
		log.Fatalf("Error: expected *firestore.DocumentSnapshot, got %T", dsnap)
	}

	jsonBytes, err := json.Marshal(doc.Data())
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(string(jsonBytes))
	return "cache1"
}

func main() {
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatalf("Error loading .env file")
	}

	ctx := context.Background()

	configDbClient, err = db.NewFirestoreClient(ctx, RealFirebaseAppInit, RealFirestoreInit)
	if err != nil {
		log.Fatal(err)
	}

	server := &server.CacheServer{Database: &InMemoryCacheDatabase{}}
	log.Fatal(http.ListenAndServe(":8080", server))
}

func RealFirebaseAppInit(ctx context.Context, sa option.ClientOption) (*firebaseLib.App, error) {
	return firebaseLib.NewApp(ctx, nil, sa)
}

// RealFirestoreInit es la función de inicialización real para Firestore Client
func RealFirestoreInit(app *firebaseLib.App, ctx context.Context) (*firestore.Client, error) {
	return app.Firestore(ctx)
}
