package entity

type Calculation struct {
	Hash  string `bson:"hash"`
	Done  bool   `bson:"done"`
	Value int    `bson:"value"`
}
