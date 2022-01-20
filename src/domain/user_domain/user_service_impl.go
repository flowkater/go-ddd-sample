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

func (s *userService) FindUsersByIds(ctx context.Context, ids []uint) ([]*UserInfo, error) {
	db := config.DBWithContext(ctx)

	users, err := s.userReader.FindUsersByIds(db, ids)
	if err != nil {
		return nil, err
	}

	var userInfos []*UserInfo
	for _, user := range users {
		userInfos = append(userInfos, &UserInfo{
			ID:   user.ID,
			Name: user.Name,
		})
	}

	return userInfos, nil
}

func (s *userService) UserByIdLoader(ctx context.Context) UserLoader {
	db := config.DBWithContext(ctx)

	return s.userReader.UserByIdLoader(db)
}
