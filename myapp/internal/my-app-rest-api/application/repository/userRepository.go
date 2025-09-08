package repository

import (
	"context"
	"errors"
	"fmt"
	"myapp/configuration"
	"myapp/internal/my-app-rest-api/application/pkg/utils"
	"myapp/internal/my-app-rest-api/domain"
	"strings"
	"time"

	"github.com/couchbase/gocb/v2"
)

type IUserRepository interface {
	GetById(ctx context.Context, id string) (*domain.User, error)
	Get(ctx context.Context) ([]*domain.User, error)
	Upsert(ctx context.Context, user *domain.User) error
}

type UserRepository struct {
	AppCluster    *gocb.Cluster
	AppUserBucket *gocb.Bucket
	userList      []*domain.User // class'ın property'si
}

// func NewUserRepository() IUserRepository ile aslında interface implementasyonu yapılmış oldu
func NewUserRepository(cluster *gocb.Cluster) IUserRepository {
	return &UserRepository{
		userList:      utils.GetUserStub(),
		AppCluster:    cluster,
		AppUserBucket: cluster.Bucket("user"),
	}
}

func (u *UserRepository) GetById(ctx context.Context, id string) (*domain.User, error) {
	var user domain.User

	queryResult, err := u.AppUserBucket.DefaultCollection().Get(id, &gocb.GetOptions{Timeout: time.Second * 1})
	if err != nil {
		return nil, errors.New("there was an error while getting data by userId")
	}

	if err = queryResult.Content(&user); err != nil {
		return nil, errors.New("there was an error while getting data by userId")
	}
	return &user, nil
}

func (u *UserRepository) Get(ctx context.Context) ([]*domain.User, error) {

	queryStr := strings.ReplaceAll("SELECT u.* FROM `{{bucket}}` u", "{{bucket}}", configuration.AppUserBucket)
	queryResult, err := u.AppCluster.Query(queryStr, &gocb.QueryOptions{
		Timeout: 10 * time.Millisecond,
	})

	if err != nil {
		return nil, err
	}

	var users []*domain.User
	for queryResult.Next() {
		var user domain.User
		queryResult.Row(&user)
		users = append(users, &user)
	}

	return users, nil
}

func (u *UserRepository) Upsert(ctx context.Context, user *domain.User) error {
	_, err := u.AppUserBucket.DefaultCollection().Upsert(user.Id,
		user,
		&gocb.UpsertOptions{Context: ctx},
	)
	if err != nil {
		fmt.Printf("userRepository.Upsert ERROR : %#v\n", err.Error())
		return errors.New("INTERNAL SERVER ERROR")
	}
	return nil
}
