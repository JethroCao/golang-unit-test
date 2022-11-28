package service

import (
	"go.uber.org/zap"
	log2 "unit-test/log"
	"unit-test/model"
	"unit-test/repository"
)

type UserService struct {
	userRepository *repository.UserRepository
}

func NewUserService() *UserService {
	return &UserService{
		userRepository: repository.NewUserRepository(),
	}
}

func (s *UserService) GetUserById(userId int) (*model.User, error) {
	log2.Logger.Info("UserService GetUserById", zap.Any("userId", userId))

	user, err := s.userRepository.GetUserById(userId)
	if err != nil {
		log2.Logger.Error("UserService GetUserById, s.userRepository.GetUserById error", zap.Error(err))
		return nil, err
	}

	return user, nil
}
