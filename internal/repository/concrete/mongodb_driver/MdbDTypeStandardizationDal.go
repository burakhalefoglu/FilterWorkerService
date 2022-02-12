package mongodb_driver

import (
	"FilterWorkerService/internal/model"
	"FilterWorkerService/pkg/database/mongodb"
	"context"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

type mdbDTypeStandardizationDal struct {
	Client *mongo.Client
}

func MdbDTypeStandardizationDalConstructor() *mdbDTypeStandardizationDal {
	return &mdbDTypeStandardizationDal{Client: mongodb.GetMongodbClient()}
}

func (m *mdbDTypeStandardizationDal) Add(tableName string, data *model.TypeStandardizationModel) error {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	collection := m.Client.Database("MLDatabase").Collection(tableName)
	var _, err = collection.InsertOne(ctx, bson.D{
		{"Key", data.Key},
		{"Value", data.Value},
	})
	if err != nil {
		return err
	}
	return nil
}

func (m *mdbDTypeStandardizationDal) GetByKey(tableName string, key string) (*model.TypeStandardizationModel, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	collection := m.Client.Database("MLDatabase").Collection(tableName)
	var result = collection.FindOne(ctx, bson.D{{
		"Key", key,
	}})

	var model = model.TypeStandardizationModel{}
	if result.Err() != nil {
		return &model, result.Err()
	}
	var err = result.Decode(&model)
	if err != nil {
		return &model, err
	}
	return &model, nil
}

func (m *mdbDTypeStandardizationDal) GetAll(tableName string) (*[]model.TypeStandardizationModel, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	collection := m.Client.Database("MLDatabase").Collection(tableName)
	var cur, err = collection.Find(ctx, bson.M{})

	var models []model.TypeStandardizationModel
	if err != nil {
		return nil, err
	}
	cur.All(ctx, &models)
	return &models, nil
}

func (m *mdbDTypeStandardizationDal) GetMaxByValue(tableName string) (int16, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	collection := m.Client.Database("MLDatabase").Collection(tableName)

	filter := []bson.M{{
		"$group": bson.M{
			"_id": nil,
			"max": bson.M{"$max": "$Value"},
		}},
	}
	var result, err = collection.Aggregate(ctx, filter)
	if err != nil {
		return 0, err
	}

	var model = model.TypeStandardizationModel{}
	if result.Err() != nil {
		return 0, result.Err()
	}
	decodeErr := result.Decode(&model)
	if decodeErr != nil {
		return 0, decodeErr
	}
	return model.Value, nil
}
