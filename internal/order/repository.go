package order

import (
	"context"
	"event-driven-sample/pkg/entity"
	"fmt"
	"go.mongodb.org/mongo-driver/v2/bson"
	"go.mongodb.org/mongo-driver/v2/mongo"
	"go.mongodb.org/mongo-driver/v2/mongo/options"
	"log"
)

type Repository struct {
	Client     *mongo.Client
	Collection *mongo.Collection
}

func NewRepository(host, database, username, password string) (*Repository, error) {
	// "mongodb://username:password@localhost:27017"
	connectionString := fmt.Sprintf("mongodb://%v:%v@%v", username, password, host)
	client, err := mongo.Connect(options.Client().ApplyURI(connectionString))
	if err != nil {
		return nil, err
	}
	collection := client.Database(database).Collection("calculations")
	return &Repository{client, collection}, nil
}

func (r Repository) Save(calculation entity.Calculation) error {
	_, err := r.Collection.InsertOne(context.TODO(), calculation)
	if err != nil {
		return err
	}
	return nil
}

func (r Repository) FindByHash(matrixHash string) error {
	filter := bson.D{{"hash", matrixHash}}
	r.Collection.FindOne(context.TODO(), filter)

	var result entity.Calculation
	err := r.Collection.FindOne(context.TODO(), filter).Decode(&result)
	if err != nil {
		return err
	}
	return nil
}

func (r Repository) Close() {
	if err := r.Client.Disconnect(context.TODO()); err != nil {
		log.Printf("failed to disconnect mongo client: %v", err)
	}
}
