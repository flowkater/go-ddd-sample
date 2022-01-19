package user_infrastructure

import (
	"github.com/flowkater/go-ddd-sample/src/domain/user_domain"
	"gorm.io/gorm"
)

type UserRepository interface {
	GetUserById(db *gorm.DB, id uint) (*user_domain.User, error)
	FindUsersByIds(db *gorm.DB, ids []uint) ([]*user_domain.User, error)
}

type userRepository struct{}

func NewUserRepository() UserRepository {
	return &userRepository{}
}

func (r *userRepository) GetUserById(db *gorm.DB, id uint) (*user_domain.User, error) {
	var user *user_domain.User

	if err := db.Debug().Where("id = ?", id).First(&user).Error; err != nil {
		return nil, err
	}

	return user, nil
}

func (r *userRepository) FindUsersByIds(db *gorm.DB, ids []uint) ([]*user_domain.User, error) {
	var users []*user_domain.User

	if err := db.Debug().Where("id IN (?)", ids).Find(&users).Error; err != nil {
		return nil, err
	}

	return users, nil
}
