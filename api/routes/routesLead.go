package routes

import (
	"os"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"

	"poc/api/controller"
)

var LogFile, _ = os.Create("logfile.txt")

func CreateRoutesAndServer() {

	var e = echo.New()

	//echo.Logger()
	e.Use(middleware.Recover())

	//basic authentication
	e.Use(middleware.BasicAuth(func(userName string, password string, c echo.Context) (bool, error) {
		if userName == "ritik" && password == "ritik@leadSquared" {
			return true, nil
		} else {
			return false, nil
		}
	}))

	e.Use(middleware.Logger()) // this will print  jo bhi request aayi hai

	// jo bhi request(route ke leye) call kar rahe wo logFile me jake save ho jayegi lekin jese hi tm apna server band kargoe all the entries khatam ho jayegi
	e.Use(middleware.LoggerWithConfig(middleware.LoggerConfig{
		Format: "method=${method}, uri=${uri}, status=${status} ,time=${time} \n ",
		Output: LogFile,
	}))

	//apis
	e.GET("/get-all-leads", controller.GetAllLeads)
	e.GET("/get-lead", controller.GetLead)
	e.POST("/create", controller.CreateLead)
	e.PUT("/get-lead", controller.UpdateLead)

	//staring the server
	e.Start(":5775")
}
