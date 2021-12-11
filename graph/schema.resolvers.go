package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	"fmt"

	"github.com/iyiola-dev/go-graphql/graph/generated"
	"github.com/iyiola-dev/go-graphql/graph/model"
)

func (r *mutationResolver) CreatBook(ctx context.Context, input model.BookInput) (*model.Book, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *mutationResolver) DeleteBook(ctx context.Context, id int) (string, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *mutationResolver) UpdateBook(ctx context.Context, id int, input model.BookInput) (string, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *queryResolver) GetAllBooks(ctx context.Context) ([]*model.Book, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *queryResolver) GetOneBook(ctx context.Context, id int) (*model.Book, error) {
	panic(fmt.Errorf("not implemented"))
}

// Mutation returns generated.MutationResolver implementation.
func (r *Resolver) Mutation() generated.MutationResolver { return &mutationResolver{r} }

// Query returns generated.QueryResolver implementation.
func (r *Resolver) Query() generated.QueryResolver { return &queryResolver{r} }

type mutationResolver struct{ *Resolver }
type queryResolver struct{ *Resolver }
