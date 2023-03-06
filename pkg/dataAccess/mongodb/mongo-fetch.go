package mongodb

import (
	"context"
	"fmt"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"

	"poc/pkg/models"
)

type Collection struct {
	LeadCollection *mongo.Collection
}

//func (db *Collection) CreateMongoCollection(nameOfCollection string) {
//	db.LeadCollection = config.Db.Collection(nameOfCollection) //creating the collection
//}

func (db *Collection) FindOne(filter interface{}) *mongo.SingleResult {
	return db.LeadCollection.FindOne(context.Background(), filter)
}

func (db *Collection) FindAll(filter interface{}, opts ...*options.FindOptions) (*mongo.Cursor, error) {
	return db.LeadCollection.Find(context.Background(), filter)
}

func (db *Collection) InsertOne(reqBody models.Lead) (*mongo.InsertOneResult, error) {
	return db.LeadCollection.InsertOne(context.Background(), reqBody)
}

func (db *Collection) UpdateOne(filter interface{}, update interface{}) (*mongo.UpdateResult, error) {
	return db.LeadCollection.UpdateOne(context.Background(), filter, update)
}

func (db *Collection) TotalDocument() (int, error) {
	if db.LeadCollection == nil {
		fmt.Println("the collection is not initialise")
	}
	len, err := db.LeadCollection.EstimatedDocumentCount(context.Background())
	ans := (int)(len)
	return ans, err
}
