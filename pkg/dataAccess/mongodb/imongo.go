package mongodb

import (
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"

	"poc/pkg/models"
)

type DbMethods interface {
	FindOne(interface{}) *mongo.SingleResult
	FindAll(interface{}, *options.FindOptions) (*mongo.Cursor, error)
	InsertOne(models.Lead) (*mongo.InsertOneResult, error)
	UpdateOne(interface{}, interface{}) (*mongo.UpdateResult, error)
	TotalDocument() (int64, error)
}
