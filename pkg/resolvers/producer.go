package resolvers

import (
	"context"

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
