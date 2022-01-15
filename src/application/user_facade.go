package application

import (
	"context"

	"github.com/flowkater/go-ddd-sample/src/domain/user_domain"
)

type UserFacade struct {
	userService user_domain.UserService
}

func NewUserFacade(userService user_domain.UserService) UserFacade {
	return UserFacade{
		userService: userService,
	}
}

func (u *UserFacade) GetUserById(ctx context.Context, id uint) (*user_domain.UserInfo, error) {
	return u.userService.GetUserById(ctx, id)
}
