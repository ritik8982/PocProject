package controller

import (
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
	"go.mongodb.org/mongo-driver/bson"

	"poc/api/service"
)

// handler function for /get-lead route
func GetLead(c echo.Context) error {
	leadId, err := strconv.Atoi(c.QueryParam("leadId")) // accessing the query param and converting to int

	if err != nil {
		return err
	}
	//service layer call
	ans := service.FindOne(bson.M{"unique_id": leadId}, c.QueryParam("leadId")) //service layer will return the response
	return c.JSON(http.StatusOK, ans)
}
