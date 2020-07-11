package db

import (
	"context"
	"go.mongodb.org/mongo-driver/mongo"
)

var (
	MongoClient *mongo.Client
	MongoCtx context.Context
)