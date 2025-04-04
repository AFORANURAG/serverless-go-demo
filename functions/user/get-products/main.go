package main

import (
	"context"
	"os"

	domain "github.com/aws-samples/serverless-go-demo/domain/products"
	handlers "github.com/aws-samples/serverless-go-demo/handlers/products"
	"github.com/aws-samples/serverless-go-demo/store"

	"github.com/aws/aws-lambda-go/lambda"
)

func main() {
	tableName, ok := os.LookupEnv("TABLE")
	if !ok {
		panic("Need TABLE environment variable")
	}

	dynamodb := store.NewDynamoDBStore(context.TODO(), tableName)
	domain := domain.NewProductsDomain(dynamodb)
	handler := handlers.NewAPIGatewayV2Handler(domain)
	lambda.Start(handler.AllHandler)
}
