package main

import (
	"poc/api/routes"
	"poc/pkg/config"
	"poc/pkg/dataAccess"
)

func main() {
	//first we need to create database and then collection and then routes
	config.CreateRedisDatabase()
	config.CreateDatabase("PocDb")
	dataAccess.CreateLeadCollection()
	routes.CreateRoutesAndServer() // to make the routes and start Server

	// to create the mongodb database

}
