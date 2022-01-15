package user_domain

import "gorm.io/gorm"

type UserReader interface {
	GetUserById(db *gorm.DB, id uint) (*User, error)
}
