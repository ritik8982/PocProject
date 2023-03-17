package main

import (
	"poc/api/routes"
	"poc/api/service"
	"poc/pkg/config"
)

func main() {
	//first we need to create database and then collection and then routes
	// From the controller we will call the service layer api and then from service layer we will call database layer

	config.CreateRedisDatabase()   // creating the redis database
	config.CreateDatabase("PocDb") //creating the mongo database

	service.InitialiseLogger()     // this will initialse the logger
	service.InitialiseCollection() // this function will create Collection instance and will initialse as well

	routes.CreateRoutesAndServer() // to make the routes and start Server

	defer routes.LogFile.Close()
}
