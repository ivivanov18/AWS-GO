package api

import "lambda-func/database"

type ApiHandler struct {
	dbStore database.DynamoDbClient
}

func NewApiHandler(dbStore database.DynamoDbClient) ApiHandler {
	return ApiHandler{
		dbStore: dbStore,
	}
}
