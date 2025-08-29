package utils

import "myapp/internal/my-app-rest-api/domain"

func GetUserStub() []*domain.User {
	return []*domain.User{
		{
			Id:        "1",
			FirtsName: "Fatmagül",
			LastName:  "Polat",
			Email:     "polat.ftmgl@gmail.com",
			Password:  "1234",
			Age:       29,
		},
		{
			Id:        "2",
			FirtsName: "Kadir",
			LastName:  "Aslan",
			Email:     "kadir.aslan@gmail.com",
			Password:  "1234",
			Age:       27,
		},
		{
			Id:        "1",
			FirtsName: "Ayşe",
			LastName:  "Polat",
			Email:     "ayşe.polat@gmail.com",
			Password:  "1234",
			Age:       15,
		},
	}
}
