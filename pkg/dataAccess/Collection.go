package dataAccess

import (
	"fmt"

	"poc/pkg/config"
	"poc/pkg/dataAccess/mongodb"
)

// collection can be differeent that 's why using this function to create the collection
var Collection mongodb.Collection

func CreateLeadCollection() {
	fmt.Println("inside the createLeadCollection")
	Collection.LeadCollection = config.Db.Collection("Leads")
	if Collection.LeadCollection == nil {
		fmt.Println("not initialised in CreateLeadCollection")
	}
}
