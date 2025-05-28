package shortening

import (
	"go.mongodb.org/mongo-driver/mongo"
)

type mgo struct {
	db *mongo.Database
}

func NewMongoDB(client *mongo.Client) *mgo {
	return &mgo{db: client.Database("url_shortener")}
}

func (m *mgo) col() *mongo.Collection {
	return m.db.Collection("shortenings")
}

// func (m *mgo) Put(ctx context.Context, shortening model.Shortening) (*model.Shortening, error) {
// 	const op = "shortening.mgo.Put"

// 	shortening.CreatedAt = time.Now().UTC()

// 	// 1. Check if the identifier already exists

// }
