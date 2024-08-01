package database

import (
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/dynamodb"
)

type DynamoDbClient struct {
	databaseStore *dynamodb.DynamoDB
}

func NewDynamoDbClient() DynamoDbClient {
	dbSession := session.Must(session.NewSession())
	db := dynamodb.New(dbSession)

	return DynamoDbClient{
		databaseStore: db,
	}
}
