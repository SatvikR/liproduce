package resolvers

import (
	"context"
	"strings"

	"github.com/SatvikR/liproduce/pkg/database"
	"github.com/SatvikR/liproduce/pkg/database/entities"
	"github.com/SatvikR/liproduce/pkg/gen/model"
)

func Producers(ctx *context.Context) ([]*entities.Producer, error) {
	var producers []*entities.Producer

	database.GetConnection().
		Model(&producers).
		Relation("Products").
		Select()

	return producers, nil
}

func Producer(ctx *context.Context, id int) (*entities.Producer, error) {
	producer := new(entities.Producer)

	database.GetConnection().
		Model(producer).
		Relation("Products").
		Where("producer.id = ?", id).
		Select()

	return producer, nil
}

func CreateProducer(ctx *context.Context, input *model.NewProducer) (*model.ProducerResponse, error) {
	newProducer := &entities.Producer{
		ProducerName: input.ProducerName,
		CanDeliver:   input.CanDeliver,
	}

	_, err := database.GetConnection().
		Model(newProducer).
		Insert()

	if err != nil {
		if strings.Contains(err.Error(), "#23505") {
			errorMessage := "That producer name is already taken"

			return &model.ProducerResponse{
				Error: &errorMessage,
			}, nil
		}
	}

	return &model.ProducerResponse{
		Producer: newProducer,
	}, nil
}
