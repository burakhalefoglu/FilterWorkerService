package mongodb

import (
	"context"
	"log"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var collection *mongo.Collection 
var ctx = context.TODO()
var mongoClient *mongo.Client

func ConnectMongoDb() {
    clientOptions := options.Client().ApplyURI("mongodb://localhost:27017/")
    client, err := mongo.Connect(ctx, clientOptions)
    if err != nil {
        log.Fatal(err)
		return
    }
	err = client.Ping(ctx, nil)
	if err != nil {
	  log.Fatal(err)
	  return
	}
	mongoClient = client
	log.Printf("Connected to MongoDB")
}

func DeleteCollection(collectionName string, filter interface{}){
	
	collection = mongoClient.Database("Client").Collection(collectionName)

	res, err := collection.DeleteOne(ctx, filter)

	if(err != nil) {
		log.Println(err)
	}
	log.Println(res)
}

func GetCollection(collectionName string, filter interface{}){
	
	collection = mongoClient.Database("Client").Collection(collectionName)

	result := collection.FindOne(ctx, filter)

	if(result.Err() != nil) {
		log.Println(result.Err())
	}
	log.Println(result)
}

func AddCollection(collectionName string, data interface{}) ( m *mongo.InsertOneResult, errors error){
	
	collection = mongoClient.Database("Client").Collection(collectionName)

	res, err := collection.InsertOne(ctx, data)

	if(err != nil) {
		log.Println(err)
		return nil, err
	}
	log.Println(res)
	return res, nil
}

func UpdateCollection(collectionName string, filter interface{}, data interface{}) ( m *mongo.UpdateResult, errors error){
	
	collection = mongoClient.Database("Client").Collection(collectionName)

	res, err := collection.UpdateOne(ctx, filter, data)

	if(err != nil) {
		log.Println(err)
		return nil, err
	}
	log.Println(res)
	return res, nil
}
