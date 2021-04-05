package gen

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"

	"github.com/SatvikR/liproduce/pkg/database/entities"
	"github.com/SatvikR/liproduce/pkg/gen/generated"
	"github.com/SatvikR/liproduce/pkg/gen/model"
	"github.com/SatvikR/liproduce/pkg/resolvers"
)

func (r *mutationResolver) CreateProducer(ctx context.Context, input model.NewProducer) (*model.ProducerResponse, error) {
	return resolvers.CreateProducer(&ctx, &input)
}

func (r *mutationResolver) CreateProduct(ctx context.Context, input model.NewProduct) (*entities.Product, error) {
	return resolvers.CreateProduct(&ctx, &input)
}

func (r *productResolver) CreatedAt(ctx context.Context, obj *entities.Product) (string, error) {
	return obj.CreatedAt.Format("2006-01-02 15:04:05"), nil
}

func (r *queryResolver) Producers(ctx context.Context) ([]*entities.Producer, error) {
	return resolvers.Producers(&ctx)
}

func (r *queryResolver) Producer(ctx context.Context, id int) (*entities.Producer, error) {
	return resolvers.Producer(&ctx, id)
}

func (r *queryResolver) Products(ctx context.Context) ([]*entities.Product, error) {
	return resolvers.Products(&ctx)
}

func (r *queryResolver) Product(ctx context.Context, id int) (*entities.Product, error) {
	return resolvers.Product(&ctx, id)
}

// Mutation returns generated.MutationResolver implementation.
func (r *Resolver) Mutation() generated.MutationResolver { return &mutationResolver{r} }

// Product returns generated.ProductResolver implementation.
func (r *Resolver) Product() generated.ProductResolver { return &productResolver{r} }

// Query returns generated.QueryResolver implementation.
func (r *Resolver) Query() generated.QueryResolver { return &queryResolver{r} }

type mutationResolver struct{ *Resolver }
type productResolver struct{ *Resolver }
type queryResolver struct{ *Resolver }
