package api

import (
	"fmt"
	"lambda-func/database"
	"lambda-func/types"
)

type ApiHandler struct {
	dbStore database.UserStore
}

func NewApiHandler(dbStore database.UserStore) ApiHandler {
	return ApiHandler{
		dbStore: dbStore,
	}
}

func (api ApiHandler) RegisterUserHandler(event types.RegisterUser) error {
	if event.Username == "" || event.Password == "" {
		return fmt.Errorf("request has empty fields")
	}

	// does a user with this username already exist?
	userExists, err := api.dbStore.DoesUserExists(event.Username)

	if err != nil {
		return fmt.Errorf("there was an error checking if the user exists %w", err)
	}

	if userExists {
		return fmt.Errorf("a user with this username already exists")
	}

	// we know that the user does not exist

	err = api.dbStore.InsertUser(event)

	if err != nil {
		return fmt.Errorf("there was an error register the user %w", err)
	}

	return nil
}
