package user_domain

import (
	"context"

	"github.com/flowkater/go-ddd-sample/src/config"
)

type userService struct {
	userReader UserReader
}

func NewUserService(userReader UserReader) UserService {
	return &userService{
		userReader: userReader,
	}
}

func (s *userService) GetUserById(ctx context.Context, id uint) (*UserInfo, error) {
	db := config.DBWithContext(ctx)

	user, err := s.userReader.GetUserById(db, id)
	if err != nil {
		return nil, err
	}

	return &UserInfo{
		ID:   user.ID,
		Name: user.Name,
	}, nil
}
