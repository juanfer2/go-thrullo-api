package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	"fmt"

	"github.com/juanfer2/go-thrullo-api.git/src/graph/generated"
	"github.com/juanfer2/go-thrullo-api.git/src/models"
)

func (r *boardResolver) CreatedAt(ctx context.Context, obj *models.Board) (string, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *boardResolver) UpdateAt(ctx context.Context, obj *models.Board) (string, error) {
	panic(fmt.Errorf("not implemented"))
}

// Board returns generated.BoardResolver implementation.
func (r *Resolver) Board() generated.BoardResolver { return &boardResolver{r} }

type boardResolver struct{ *Resolver }
