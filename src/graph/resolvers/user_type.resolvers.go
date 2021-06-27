package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	"fmt"

	"github.com/juanfer2/go-thrullo-api.git/src/graph/generated"
	"github.com/juanfer2/go-thrullo-api.git/src/models"
)

func (r *userResolver) CreatedAt(ctx context.Context, obj *models.User) (string, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *userResolver) UpdateAt(ctx context.Context, obj *models.User) (string, error) {
	panic(fmt.Errorf("not implemented"))
}

// User returns generated.UserResolver implementation.
func (r *Resolver) User() generated.UserResolver { return &userResolver{r} }

type userResolver struct{ *Resolver }
