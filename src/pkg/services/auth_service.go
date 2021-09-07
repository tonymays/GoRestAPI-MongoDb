package services

import (
	"go.mongodb.org/mongo-driver/mongo"
	"pkg/configuration"
)

// ---- AuthService ----
type AuthService struct {
	config		configuration.Configuration
	dbClient	*mongo.Client
}

// ---- NewAuthService ----
func NewAuthService(config configuration.Configuration, dbClient *mongo.Client) *AuthService {
	return &AuthService{config, dbClient}
}