package routes

import (
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"

	"poc/api/controller"
)

func CreateRoutesAndServer() {

	var e = echo.New()
	e.Use(middleware.Recover())

	//basic authentication
	e.Use(middleware.BasicAuth(func(userName string, password string, c echo.Context) (bool, error) {
		if userName == "ritik" && password == "ritik@leadSquared" {
			return true, nil
		} else {
			return false, nil
		}
	}))
	//apis
	e.GET("/get-all-leads", controller.GetAllLeads)
	e.GET("/get-lead", controller.GetLead)
	e.POST("/create", controller.CreateLead)
	e.PUT("/get-lead", controller.UpdateLead)

	//staring the server
	e.Start(":5775")
}
