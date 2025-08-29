package request

import "myapp/internal/my-app-rest-api/application/handler/user"

type UserCreateRequest struct {
	FirtsName string `json:"firstName" validate:"required"`
	LastName  string `json:"lastName" validate:"required"`
	Email     string `json:"email" validate:"required"`
	Age       int    `json:"age" validate:"required"`
}

func (req *UserCreateRequest) ToCommand() user.Command {
	return user.Command{
		FirtsName: req.FirtsName,
		LastName:  req.LastName,
		Email:     req.Email,
		Age:       req.Age,
	}
}
