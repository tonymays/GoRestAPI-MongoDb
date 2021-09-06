package services

import (
	"go.mongodb.org/mongo-driver/mongo"
	"pkg"
	"pkg/configuration"
)

type AuthService struct {
	config		configuration.Configuration
	dbClient	*mongo.Client
	dbService	root.DbService
}

func NewAuthService(config configuration.Configuration, dbClient *mongo.Client, dbService root.DbService) *AuthService {
	return &AuthService{config, dbClient, dbService}
}