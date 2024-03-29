package service

import (
	"context"
	"encoding/json"
	"fmt"
	"strconv"

	"github.com/pkg/errors"
	"github.com/redis/go-redis/v9"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"

	"poc/pkg/dataAccess/mongodb"
	"poc/pkg/dataAccess/redisdb"
	"poc/pkg/logger"
	"poc/pkg/models"
)

// service layer is acting like a wrapper it will check ki data is present in redis or not and if not will bring the data mongo database and update the redis cache as well

// cache is redis database
var cache redisdb.Redis

// initialse logger
var Log *logger.MyLogger

func InitialiseLogger() {
	Log = logger.GetLogger()
}

// initialse collection
var mongoCollection *mongodb.Collection

func InitialiseCollection() {
	Log.Info("initialising mongoCollection in service layer")
	mongoCollection = mongodb.GetCollection()
}

// create lead
func Create(reqBody models.Lead) (*mongo.InsertOneResult, error) {

	//Log.Info("Create method in service layer has started")

	// unique id of the lead will be the length of the collection
	//ans, err := mongoCollection.TotalDocument()
	ans, err := mongoCollection.TotalDocument()
	if err != nil {
		// my log
		Log.Error("error in Create Service layer")
		return nil, err
	}
	reqBody.UniqueId = ans

	//calling the dataAccess layer and return the reponse sent by the dataAccess layer
	ans2, err2 := mongoCollection.InsertOne(reqBody)
	//Log.Info("Create method in service layer has ended")
	return ans2, err2
}

// get all lead
func FindAll(filter interface{}, opts ...*options.FindOptions) []models.Lead {

	var allLeads []models.Lead // this will contain all the leads
	var oneLead models.Lead

	//calling the dataAccess layer
	//findElementRes, err := dataAccess.Collection.FindAll(filter)
	findElementRes, err := mongoCollection.FindAll(filter)
	if err != nil {
		fmt.Println("error in mongo - fetch FindAll")
	}
	//dataAccess layer will return a cursor, we will iterate over it and will store it
	for findElementRes.Next(context.Background()) {
		err := findElementRes.Decode(&oneLead)
		if err != nil {
			fmt.Println(err)
		}
		allLeads = append(allLeads, oneLead)
	}
	return allLeads
}

// getLead
func FindOne(filter interface{}, key string) models.Lead {
	var ans models.Lead

	//checking the data in redisdb db
	val2, err2 := cache.Get(key).Result()

	if err2 == redis.Nil {
		//not present in redis database so add the key
		fmt.Println("FindOne , adding the key, in if")
		//checking if the leadId exists in db or not
		var findOneLead models.Lead

		//err2 := dataAccess.Collection.FindOne(bson.M{"unique_id": keyInt}).Decode(&findOneLead)
		err2 := mongoCollection.FindOne(filter).Decode(&findOneLead)

		if err2 != nil {
			fmt.Println("error in FindOne service layer")
		}

		findOneLeadv2, _ := json.Marshal(findOneLead) //converting to byte
		key := strconv.Itoa(findOneLead.UniqueId)

		err4 := cache.Set(key, findOneLeadv2, 0).Err()

		if err4 != nil {
			fmt.Println("not able to set the values in redisdb")
		}
		return findOneLead
	} else {
		// exists in redisdb db

		err := json.Unmarshal([]byte(val2), &ans)
		if err != nil {
			fmt.Println("error in unmarshalling")
		}
		return ans
	}

}

// update lead
func UpdateOne(reqBody models.Lead, key string) (*mongo.UpdateResult, error) {

	keyInt, _ := strconv.Atoi(key)
	oneLead := FindOne(bson.M{"unique_id": keyInt}, key)
	var empty models.Lead
	if oneLead == empty {
		fmt.Println("Lead doesn't exists")
		return nil, errors.New("Lead doesn't exists")
	}
	keyInt, err := strconv.Atoi(key) //leadId should be same
	reqBody.UniqueId = keyInt

	// fields which we want to update
	updateField := bson.M{"$set": bson.M{"first_name": reqBody.FirstName, "last_name": reqBody.LastName, "email": reqBody.Email, "phone_no": reqBody.PhoneNo, "company_name": reqBody.CompanyName, "country": reqBody.Country}}

	//updateFileRes, err := collection.UpdateOne(bson.M{"unique_id": findOneLead.UniqueId}, updateField)
	updateFileRes, err := mongoCollection.UpdateOne(bson.M{"unique_id": keyInt}, updateField)

	if err != nil {
		fmt.Println("error in updateFileRes")
		return nil, errors.New("error in updateFileRes")
	}

	reqBodyv2, _ := json.Marshal(reqBody) // convert it to byte so we can store it in redisdb

	//updating in redisdb
	err7 := cache.Set(key, reqBodyv2, 0).Err()

	if err7 == redis.Nil {
		fmt.Println("not able to set the values in redisdb")
	}
	return updateFileRes, nil

}
