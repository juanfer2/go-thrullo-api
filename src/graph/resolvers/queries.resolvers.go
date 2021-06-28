package resolvers

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	"fmt"

	"github.com/juanfer2/go-thrullo-api.git/src/graph/generated"
	"github.com/juanfer2/go-thrullo-api.git/src/models"
)

func (r *queryResolver) Healt(ctx context.Context) (*string, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *queryResolver) ListBoard(ctx context.Context) ([]*models.Board, error) {
	boardRepository := repositories.BoardRepository{}
	boards := boardRepository.List()

	return boards, nil
}

// Query returns generated.QueryResolver implementation.
func (r *Resolver) Query() generated.QueryResolver { return &queryResolver{r} }

type queryResolver struct{ *Resolver }
