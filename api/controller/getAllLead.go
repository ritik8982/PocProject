package controller

import (
	"net/http"

	"github.com/labstack/echo/v4"
	"go.mongodb.org/mongo-driver/bson"

	"poc/api/service"
	"poc/pkg/models"
)

// handlerFunc for /get-all-leads route
func GetAllLeads(c echo.Context) error {

	var allLeads []models.Lead // this will store all the information related to lead

	elementFilter := bson.M{
		"unique_id": bson.M{"$exists": true},
	}
	// service layer call
	allLeads = service.FindAll(elementFilter) //service layer will return the response

	return c.JSON(http.StatusOK, allLeads)
}
