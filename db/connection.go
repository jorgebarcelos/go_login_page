package db

import (
	"context"
	"log"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func get_connection()(client *mongo.Client, ctx context.Context){
	client, err := mongo.NewClient(options.Client().ApplyURI("mongodb://localhost:27017"))
	if err!= nil {
		log.Fatal(err)
	}
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	err = client.Connect(ctx)
	if err!=nil{
		log.Fatal(err)
	}
	return
}