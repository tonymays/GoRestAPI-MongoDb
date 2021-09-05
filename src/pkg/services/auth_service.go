package services

import (
	"pkg/data"
)

type AuthService struct {
	data	data.Data
}

func NewAuthService(data data.Data) *UserService {
	return &UserService{data}
}