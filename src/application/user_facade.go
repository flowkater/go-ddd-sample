package application

import (
	"context"
	"fmt"
	"strconv"
	"time"

	"github.com/flowkater/go-ddd-sample/src/application/dataloader"
	"github.com/flowkater/go-ddd-sample/src/domain/user_domain"
	"github.com/flowkater/go-ddd-sample/src/interfaces/model"
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

func (u *UserFacade) NewUserLoader(ctx context.Context) *dataloader.UserLoader {
	return dataloader.NewUserLoader(dataloader.UserLoaderConfig{
		Wait:     2 * time.Millisecond,
		MaxBatch: 100,
		Fetch: func(keys []uint) ([]*model.User, []error) {
			users := make([]*model.User, len(keys))
			errors := make([]error, len(keys))

			user_infos, err := u.userService.FindUsersByIds(ctx, keys)
			if err != nil {
				return nil, []error{err}
			}

			fmt.Println(len(user_infos))

			for i, key := range keys {

				users[i] = &model.User{ID: int(key), Name: "user " + strconv.Itoa(int(key))}
			}
			return users, errors
		},
	})
}
