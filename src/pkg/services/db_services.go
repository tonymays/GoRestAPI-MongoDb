package services

import (
	"go.mongodb.org/mongo-driver/mongo"
	"pkg/configuration"
)

type DbService struct {
	config				configuration.Configuration
	dbClient			*mongo.Client
	userCollection		*mongo.Collection
}

func NewDbService(config configuration.Configuration, dbClient *mongo.Client) *DbService {
	userCollection := dbClient.Database(config.DbName).Collection("users")
	return &DbService{config, dbClient, userCollection}
}
