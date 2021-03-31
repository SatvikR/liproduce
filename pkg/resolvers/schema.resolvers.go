package resolvers

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	"fmt"

	"github.com/SatvikR/liproduce/pkg/database/entities"
	"github.com/SatvikR/liproduce/pkg/generated"
)

func (r *productResolver) CreatedAt(ctx context.Context, obj *entities.Product) (string, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *queryResolver) Producers(ctx context.Context) ([]*entities.Producer, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *queryResolver) Producer(ctx context.Context, id int) (*entities.Producer, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *queryResolver) Products(ctx context.Context) ([]*entities.Product, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *queryResolver) Product(ctx context.Context, id int) (*entities.Product, error) {
	panic(fmt.Errorf("not implemented"))
}

// Product returns generated.ProductResolver implementation.
func (r *Resolver) Product() generated.ProductResolver { return &productResolver{r} }

// Query returns generated.QueryResolver implementation.
func (r *Resolver) Query() generated.QueryResolver { return &queryResolver{r} }

type productResolver struct{ *Resolver }
type queryResolver struct{ *Resolver }
