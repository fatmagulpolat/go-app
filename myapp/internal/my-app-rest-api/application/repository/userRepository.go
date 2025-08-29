package repository

import (
	"context"
	"myapp/internal/my-app-rest-api/application/pkg/utils"
	"myapp/internal/my-app-rest-api/domain"
)

type IUserRepository interface {
	GetById(ctx context.Context, id string) (*domain.User, error)
	Get(ctx context.Context) ([]*domain.User, error)
	Upsert(ctx context.Context, user *domain.User) error
}

type UserRepository struct {
	userList []*domain.User // class'ın property'si
}

// func NewUserRepository() IUserRepository ile aslında interface implementasyonu yapılmış oldu
func NewUserRepository() IUserRepository {
	return &UserRepository{
		userList: utils.GetUserStub(),
	}
}

func (u *UserRepository) GetById(ctx context.Context, id string) (*domain.User, error) {
	for _, user := range u.userList {
		if user.Id == id {
			return user, nil
		}
	}
	return nil, nil
}

func (u *UserRepository) Get(ctx context.Context) ([]*domain.User, error) {

	users := u.userList
	if users == nil {
		return make([]*domain.User, 0), nil
	}
	return u.userList, nil
}

func (u *UserRepository) Upsert(ctx context.Context, user *domain.User) error {
	u.userList = append(u.userList, user)
	return nil
}
