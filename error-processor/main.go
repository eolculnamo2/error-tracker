package main

import (
	"log"
	"context"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"

	"github.com/eolculnamo2/error-tracker/error-processor/app"
	"github.com/eolculnamo2/error-tracker/error-processor/app/db"
)

/*
Todo next:
1) Connect to MongoDB
2) Add routes to process errors
3) Add authenticated Routes to return errors 
4) Add CORS restrictions.. If a route that provides info, it should be strictly for dashbaord otherwise any url
*/

func main() {
	log.Println("Starting Error Processor")

	var newClientErr error
	db.MongoClient, newClientErr = mongo.NewClient(options.Client().ApplyURI("mongodb://rob:P%40##wordRob123@ds017195.mlab.com:17195/error-tracker?retryWrites=false"))
    if newClientErr != nil {
        log.Fatal(newClientErr)
    }
    db.MongoCtx = context.Background()
    err := db.MongoClient.Connect(db.MongoCtx)
    if err != nil {
        log.Fatal(err)
    }
    defer db.MongoClient.Disconnect(db.MongoCtx)
	app.StartRoutes()
}