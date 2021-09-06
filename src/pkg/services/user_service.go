package services

import (
	"go.mongodb.org/mongo-driver/mongo"
	"pkg"
	"pkg/configuration"
)

type UserService struct {
	config		configuration.Configuration
	dbClient	*mongo.Client
	dbService	root.DbService
}

func NewUserService(config configuration.Configuration, dbClient *mongo.Client, dbService root.DbService) *UserService {
	return &UserService{config, dbClient, dbService}
}

func (rcvr *UserService) CreateUser(u root.User) (root.User, error) {
	return root.User{}, nil
}