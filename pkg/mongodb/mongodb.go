package mongodb

import (
	"context"
	"event-driven-sample/pkg/entity"
	"fmt"
	"go.mongodb.org/mongo-driver/v2/mongo"
	"go.mongodb.org/mongo-driver/v2/mongo/options"
	"log"
)

type MongoDB struct {
	Client   *mongo.Client
	Database string
}

func New(host, database, username, password string) (*MongoDB, error) {
	// "mongodb://username:password@localhost:27017"
	connectionString := fmt.Sprintf("mongodb://%v:%v@%v", username, password, host)
	client, err := mongo.Connect(options.Client().ApplyURI(connectionString))
	if err != nil {
		return nil, err
	}
	return &MongoDB{client, database}, nil
}

func (m MongoDB) Save(calculation entity.Calculation) error {
	collection := m.Client.Database(m.Database).Collection("calculations")
	_, err := collection.InsertOne(context.TODO(), calculation)
	if err != nil {
		return fmt.Errorf("failed to save: %w", err)
	}
	return nil
}

func (m MongoDB) Close() {
	if err := m.Client.Disconnect(context.TODO()); err != nil {
		log.Printf("failed to disconnect mongo client: %v", err)
	}
}
