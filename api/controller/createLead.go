package controller

import (
	"net/http"

	"github.com/labstack/echo/v4"
	"gopkg.in/go-playground/validator.v9"

	"poc/api/service"
	"poc/pkg/models"
)

// handlerFunc for Create route
func CreateLead(c echo.Context) error {

	var reqBody models.Lead
	err := c.Bind(&reqBody) //whatever the data is coming, bind with reqBody
	if err != nil {
		return err
	}
	var v = validator.New()
	err2 := v.Struct(&reqBody) //checking validation
	if err2 != nil {
		return c.JSON(http.StatusNotFound, "validation match nhi huye")
	}

	//insert the data into the collection - make call to the service layer
	res, err4 := service.Create(reqBody) //service layer will return the response

	if err4 != nil {
		return err4
	}
	return c.JSON(http.StatusOK, res)
}
