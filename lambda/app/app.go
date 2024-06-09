package app

import (
	"lambda-func/api"
	"lambda-func/database"
)

type App struct {
	ApiHandler api.ApiHandler
}

func NewApp() App{
	// we actually initialize the database here
	// gets passed to the api handler
	db := database.NewDynamoDBClient("userTable");
	apiHandler := api.NewApiHandler(db);

	return App {
		ApiHandler: apiHandler,
	}
}