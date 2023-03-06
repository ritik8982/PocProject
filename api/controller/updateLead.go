package controller

import (
	"net/http"

	"github.com/labstack/echo/v4"
	"gopkg.in/go-playground/validator.v9"

	"poc/api/service"
	"poc/pkg/models"
)

// handler function for /get-lead(update) route
func UpdateLead(c echo.Context) error {

	key := c.QueryParam("leadId")

	var reqBody models.Lead
	err := c.Bind(&reqBody) //binding the data(sent by user) with reqBody
	if err != nil {
		return err
	}

	var v = validator.New()
	err2 := v.Struct(reqBody) //checking validation
	if err2 != nil {
		return err2
	}

	//calling service layer
	ans, err3 := service.UpdateOne(reqBody, key) //service layer will return the response
	if err3 != nil {
		return err3
	}
	return c.JSON(http.StatusOK, ans)
}
