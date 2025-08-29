package user

import (
	"context"
	"myapp/internal/my-app-rest-api/application/repository"
	"myapp/internal/my-app-rest-api/domain"

	"github.com/google/uuid"
)

type ICommandHandler interface {
	Save(ctx context.Context, command Command) error
}

type commandHandler struct {
	userRepository repository.IUserRepository
}

// Save implements ICommandHandler.
func (h *commandHandler) Save(ctx context.Context, command Command) error {

	if err := h.userRepository.Upsert(ctx, h.BuildEntity(command)); err != nil {
		return err
	}
	return nil
}

func NewCommandHandler(userRepository repository.IUserRepository) ICommandHandler {
	return &commandHandler{userRepository: userRepository}
}

func (h *commandHandler) BuildEntity(command Command) *domain.User {
	return &domain.User{
		Id:        uuid.NewString(),
		FirtsName: command.FirtsName,
		LastName:  command.LastName,
		Email:     command.Email,
		Age:       command.Age,
		Password:  command.Password,
	}
}
