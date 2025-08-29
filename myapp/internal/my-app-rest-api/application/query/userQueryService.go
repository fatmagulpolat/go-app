package query

import (
	"context"
	"errors"
	"myapp/internal/my-app-rest-api/application/repository"
	"myapp/internal/my-app-rest-api/domain"
)

type IUserQueryService interface {
	GetById(ctx context.Context, id string) (*domain.User, error)
	Get(ctx context.Context) ([]*domain.User, error)
}

type userQueryService struct {
	userRepository repository.IUserRepository
}

func NewUserQueryService(userRepository repository.IUserRepository) IUserQueryService {
	return &userQueryService{
		userRepository: userRepository,
	}
}

// GetById implements IUserQueryService.
func (u *userQueryService) GetById(ctx context.Context, id string) (*domain.User, error) {
	user, err := u.userRepository.GetById(ctx, id)
	if err != nil {
		return nil, err
	}

	if user == nil {
		return nil, errors.New("not found error")
	}

	return user, nil
}

// Get implements IUserQueryService.
func (u *userQueryService) Get(ctx context.Context) ([]*domain.User, error) {
	users, err := u.userRepository.Get(ctx)
	if err != nil {
		return nil, err
	}

	if users == nil {
		return nil, errors.New("not found users")
	}

	return users, nil
}
