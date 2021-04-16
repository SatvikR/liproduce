package resolvers

import (
	"context"

	"github.com/SatvikR/liproduce/pkg/database"
	"github.com/SatvikR/liproduce/pkg/database/entities"
	"github.com/SatvikR/liproduce/pkg/gen/model"
	"github.com/SatvikR/liproduce/pkg/utils"
)

const (
	producerType   = "producer"
	restaurantType = "restaurant"
)

func CreateUser(ctx *context.Context, input *model.NewUser) (*model.UserResponse, error) {
	hashedPassword, err := utils.Hash(&input.Password)

	if err != nil {
		return &model.UserResponse{
			UID:    0,
			Errors: []string{"internal server error"},
		}, nil
	}

	if input.UserType != producerType && input.UserType != restaurantType {
		return &model.UserResponse{
			UID:    0,
			Errors: []string{"invalid user type"},
		}, nil
	}

	newUser := &entities.User{
		Username: input.Username,
		Email:    input.Email,
		UserType: input.UserType,
		Password: hashedPassword,
	}

	database.GetConnection().
		Model(newUser).
		Insert()
	return &model.UserResponse{
		UID:    int(newUser.Id),
		Errors: []string{},
	}, nil
}
