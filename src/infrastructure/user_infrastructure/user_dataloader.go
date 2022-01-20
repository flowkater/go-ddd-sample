package user_infrastructure

import (
	"time"

	"github.com/flowkater/go-ddd-sample/src/domain/user_domain"
	"gorm.io/gorm"
)

func newUserDataloader(
	db *gorm.DB,
	userRepository UserRepository,
) user_domain.UserLoader {
	return *user_domain.NewUserLoader(user_domain.UserLoaderConfig{
		MaxBatch: 100,
		Wait:     1 * time.Millisecond,
		Fetch: func(ids []uint) ([]*user_domain.User, []error) {
			errors := make([]error, len(ids))

			records, err := userRepository.FindUsersByIds(db, ids)
			if err != nil {
				return nil, []error{err}
			}

			userByID := map[uint]*user_domain.User{}
			for _, record := range records {
				userByID[record.ID] = record
			}

			users := make([]*user_domain.User, len(ids))

			for i, id := range ids {
				users[i] = userByID[id]
			}

			return users, errors
		},
	})
}
