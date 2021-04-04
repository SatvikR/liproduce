package resolvers

import (
	"context"

	"github.com/SatvikR/liproduce/graph/model"
	"github.com/SatvikR/liproduce/pkg/database"
	"github.com/SatvikR/liproduce/pkg/database/entities"
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

func CreateProducer(ctx *context.Context, input *model.NewProducer) (*entities.Producer, error) {
	newProducer := &entities.Producer{
		ProducerName: input.ProducerName,
		CanDeliver:   input.CanDeliver,
	}

	database.GetConnection().
		Model(newProducer).
		Insert()

	return newProducer, nil
}
