package resolvers

import (
	"context"

	"github.com/SatvikR/liproduce/pkg/database"
	"github.com/SatvikR/liproduce/pkg/database/entities"
	"github.com/SatvikR/liproduce/pkg/gen/model"
)

func CreateProduct(ctx *context.Context, input *model.NewProduct) (*entities.Product, error) {
	product := &entities.Product{
		OwnerId:     int32(input.OwnerID),
		ProductName: input.ProductName,
	}

	database.GetConnection().
		Model(product).
		Insert()

	return product, nil
}

func Products(ctx *context.Context) ([]*entities.Product, error) {
	var products []*entities.Product

	database.GetConnection().
		Model(&products).
		Relation("Owner").
		Select()

	return products, nil
}

func Product(ctx *context.Context, id int) (*entities.Product, error) {
	product := new(entities.Product)

	database.GetConnection().
		Model(product).
		Relation("Owner").
		Where("product.id = ?", id).
		Select()

	return product, nil
}
