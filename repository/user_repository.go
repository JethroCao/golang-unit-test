package repository

import (
	"go.uber.org/zap"
	"unit-test/database"
	log2 "unit-test/log"
	"unit-test/model"
)

type UserRepository struct {
}

func NewUserRepository() *UserRepository {
	return &UserRepository{}
}

func (r *UserRepository) GetUserById(userId int) (*model.User, error) {
	log2.Logger.Info("UserRepository GetUserById", zap.Any("userId", userId))

	user := &model.User{}
	if err := database.Db.Where("id = ?", userId).First(user).Error; err != nil {
		log2.Logger.Error("GetUserById database error", zap.Error(err))
		return nil, err
	}

	return user, nil
}
