package user_domain

import "context"

type UserService interface {
	GetUserById(ctx context.Context, id uint) (*UserInfo, error)
}
