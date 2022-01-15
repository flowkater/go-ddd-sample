package user_infrastructure

import (
	"github.com/flowkater/go-ddd-sample/src/domain/user_domain"
	"gorm.io/gorm"
)

type userReader struct {
	userRepository UserRepository
}

func NewUserReader(userRepository UserRepository) user_domain.UserReader {
	return &userReader{
		userRepository: userRepository,
	}
}

func (r *userReader) GetUserById(db *gorm.DB, id uint) (*user_domain.User, error) {
	return r.userRepository.GetUserById(db, id)
}
