package entity

import "go.mongodb.org/mongo-driver/v2/bson"

type Calculation struct {
	ID    bson.ObjectID `bson:"_id,omitempty"`
	Hash  string        `bson:"hash"`
	Done  bool          `bson:"done"`
	Value int           `bson:"value"`
}
