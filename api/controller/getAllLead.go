package controller

import (
	"net/http"

	"github.com/labstack/echo/v4"
	"go.mongodb.org/mongo-driver/bson"

	"poc/api/service"
	"poc/pkg/models"
)

func GetAllLeads(c echo.Context) error {

	var allLeads []models.Lead // this will store all the information related to lead

	elementFilter := bson.M{
		"unique_id": bson.M{"$exists": true},
	}
	//findElementRes, err := Col_of_Leads.Find(context.Background(), elementFilter) //this will return cursor to the first element and will loop it and store it in slice of lead and will display
	//findElementRes, err := collection.FindAll(elementFilter)
	allLeads = service.FindAll(elementFilter)

	return c.JSON(http.StatusOK, allLeads)
}
