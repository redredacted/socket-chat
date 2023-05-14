package service

import (
	"github.com/redredacted/socket-chat/repositoty/users"
)

type UserService struct {
	userRepository UserRepository
}

func NewUserService(userRepository UserRepository) *UserService {
	return &UserService{
		userRepository: userRepository,
	}
}
