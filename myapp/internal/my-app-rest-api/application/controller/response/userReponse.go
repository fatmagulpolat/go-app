package response

import (
	"myapp/internal/my-app-rest-api/domain"
)

type UserResponse struct {
	Id        string `json :"id"`
	FirstName string `json:"firstName"`
	LastName  string `json:"lastName"`
	Email     string `json:"email"`
	Age       int    `json:"age"`
}

func ToUserResponse(user *domain.User) UserResponse {

	return UserResponse{
		Id:        user.Id,
		FirstName: user.FirtsName,
		LastName:  user.LastName,
		Email:     user.Email,
		Age:       user.Age,
	}
}

func ToUserResponseList(users []*domain.User) []UserResponse {
	userDtoList := make([]UserResponse, 0)

	for _, user := range users {
		userDto := ToUserResponse(user)
		userDtoList = append(userDtoList, userDto)
	}
	return userDtoList
}
