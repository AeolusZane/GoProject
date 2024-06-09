package main

import (
	"github.com/aws/aws-lambda-go/lambda"
	"lambda-func/app"
	"fmt"
)

type MyEvent struct {
	Username string `json:"username"` // take this field from the payload

}

// Take in payload and something with it
func HandleRequest(event MyEvent) (string, error){
	if event.Username == "" {
		return "", fmt.Errorf("username cannot be empty");
	}
	return fmt.Sprintf("Hello %s!", event.Username), nil;
}

func main(){
	myApp := app.NewApp();
	lambda.Start(myApp.ApiHandler.RegisterUserHandler)
}