package main

import (
	"context"
	"errors"
	"fmt"
	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
	"github.com/brendsanchez/go-project/aws"
	"github.com/brendsanchez/go-project/db"
	"github.com/brendsanchez/go-project/models"
	"os"
)

func main() {
	lambda.Start(execLambda)
}

func execLambda(ctx context.Context, event events.CognitoEventUserPoolsPostConfirmation) (events.CognitoEventUserPoolsPostConfirmation, error) {
	aws.Init()
	if !existSecret() {
		err := errors.New("error to find env SecretName")
		return event, err
	}

	var datos models.SignUp

	for row, att := range event.Request.UserAttributes {
		switch row {
		case "email":
			datos.UserEmail = att
		case "sub":
			datos.UserUUID = att
		}

	}
	err := db.ReadSecret()
	if err != nil {
		fmt.Print("fail to readSecret")
		return event, err
	}

	err = db.SignUp(datos)
	return event, err
}

func existSecret() bool {
	var exist bool
	_, exist = os.LookupEnv("SecretName")
	return exist
}
