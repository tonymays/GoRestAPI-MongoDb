package services

import (
	"pkg/data"
)

type UserService struct {
	data	data.Data
}

func NewUserService(data data.Data) *UserService {
	return &UserService{data}
}