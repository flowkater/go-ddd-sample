package mapper

import (
	"github.com/flowkater/go-ddd-sample/src/domain/user_domain"
	"github.com/flowkater/go-ddd-sample/src/interfaces/model"
)

func UserOf(userInfo *user_domain.UserInfo) *model.User {
	return &model.User{
		ID:   int(userInfo.ID),
		Name: userInfo.Name,
	}
}
