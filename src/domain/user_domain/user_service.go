package user_domain

import "context"

type UserService interface {
	GetUserById(ctx context.Context, id uint) (*UserInfo, error)
	FindUsersByIds(ctx context.Context, ids []uint) ([]*UserInfo, error)
	UserByIdLoader(ctx context.Context) UserLoader
}
