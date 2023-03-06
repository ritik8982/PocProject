package config

import (
	"context"
	"fmt"
	"time"

	"github.com/redis/go-redis/v9"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var Db *mongo.Database

func CreateDatabase(dbName string) {
	//connecting with mongodb
	client, err := mongo.NewClient(options.Client().ApplyURI("mongodb://localhost:27017"))
	if err != nil {
		fmt.Println(err)
	}
	ctx, cancel := context.WithTimeout(context.Background(), 20*time.Second)
	defer cancel()
	err = client.Connect(ctx)
	if err != nil {
		fmt.Println("error in line 25", err)
	}
	Db = client.Database(dbName) //creating the db
	fmt.Println("database created")

}

var RedisDatabase *redis.Client

func CreateRedisDatabase() {
	RedisDatabase = redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "", // no password set
		DB:       0,  // use default DB
	})
}
