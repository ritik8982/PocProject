package config

//config is database layer
import (
	"context"
	"fmt"
	"time"

	"github.com/redis/go-redis/v9"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var Db *mongo.Database

// this function will create the redis database
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

// this will represent RedisDatabase
var RedisDatabase *redis.Client

// this function will initialse the redis database instance
func CreateRedisDatabase() {
	RedisDatabase = redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "", // no password set
		DB:       0,  // use default DB
	})
}
