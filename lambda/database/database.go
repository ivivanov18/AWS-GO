package database

import (
	"fmt"
	"lambda-func/types"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/dynamodb"
)

const (
	TABLE_NAME = "userTable"
)

type UserStore interface {
	DoesUserExist(username string) (bool, error)
	InsertUser(user types.RegisterUser) error
}

// Implements the UserStore interface
// Methods: DoesUserExist and InsertUser
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

func (client DynamoDbClient) DoesUserExist(username string) (bool, error) {
	result, err := client.databaseStore.GetItem(&dynamodb.GetItemInput{
		TableName: aws.String(TABLE_NAME),
		Key: map[string]*dynamodb.AttributeValue{
			"username": {
				S: aws.String(username),
			},
		},
	})

	if err != nil {
		return true, err
	}

	if result.Item == nil {
		return false, nil
	}

	return true, nil
}

func (client DynamoDbClient) InsertUser(user types.RegisterUser) error {
	item := &dynamodb.PutItemInput{
		TableName: aws.String(TABLE_NAME),
		Item: map[string]*dynamodb.AttributeValue{
			"username": {S: aws.String(user.Username)},
			"password": {S: aws.String(user.Password)},
		},
	}

	_, err := client.databaseStore.PutItem(item)

	if err != nil {
		return err
	}

	return nil
}

func (client DynamoDbClient) GetUser(username string) (types.User, error) {
	var user types.User

	result, err := client.databaseStore.GetItem(&dynamodb.GetItemInput{
		TableName: aws.String(TABLE_NAME),
		Key: map[string]*dynamodb.AttributeValue(
			"username": {
				S: aws.String(username),
			},
		),
	})

	if err != nil {
		return user, err
	}

	if result.Item == nil {
		return user, fmt.Errorf("user not found")
	} 

	err = dynamodb.dynanodbattribute.UnmarshalMap(result.Item, &user)
	if err != nil {
		return user, err
	} 
	return user, nil
}