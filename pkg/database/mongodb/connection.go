package mongodb

import (
	"FilterWorkerService/internal/IoC"
	"context"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"os"
	"time"
)

func GetMongodbClient() *mongo.Client {
	var log = &IoC.Logger
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	var url = os.Getenv("MONGODB_CONN")

	client, err := mongo.Connect(ctx, options.Client().ApplyURI(url))
	if err != nil{
		(*log).SendPanicLog("MongoConnection", "ConnectMongodb", err)
		panic(err)
	}
	return client
}