package app

import (
	"context"
	"fmt"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"log"
	"time"
)

func NewMongoDatabase(env *Env) *mongo.Client {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	dbUser := env.DBUser
	dbPass := env.DBPass
	dbName := env.DBName

	serverAPI := options.ServerAPI(options.ServerAPIVersion1)
	optionsString := fmt.Sprintf("mongodb+srv://%s:%s@%s.cxodm.mongodb.net/?retryWrites=true&w=majority", dbUser, dbPass, dbName)
	mongoOptions := options.Client().ApplyURI(optionsString).SetServerAPIOptions(serverAPI)

	client, err := mongo.Connect(ctx, mongoOptions)
	if err != nil {
		panic(err)
	}

	err = client.Ping(ctx, nil)
	if err != nil {
		log.Fatal(err)
	}
	log.Println("Connected to MongoDB!")

	return client
}

func CloseMongoDBConnection(client *mongo.Client) {
	if client == nil {
		return
	}

	err := client.Disconnect(context.TODO())
	if err != nil {
		log.Fatal(err)
	}

	log.Println("Connection to MongoDB closed.")
}
