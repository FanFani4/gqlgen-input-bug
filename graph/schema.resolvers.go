package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	"fmt"

	"github.com/FanFani4/gqlgen-input-bug/domain"
	"github.com/FanFani4/gqlgen-input-bug/graph/generated"
)

func (r *mutationResolver) Foo(ctx context.Context, in domain.TestInput) (*domain.TestType, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *queryResolver) Foo(ctx context.Context, in domain.TestInput) (*domain.TestType, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *testTypeResolver) Role(ctx context.Context, obj *domain.TestType) (string, error) {
	panic(fmt.Errorf("not implemented"))
}

// Mutation returns generated.MutationResolver implementation.
func (r *Resolver) Mutation() generated.MutationResolver { return &mutationResolver{r} }

// Query returns generated.QueryResolver implementation.
func (r *Resolver) Query() generated.QueryResolver { return &queryResolver{r} }

// TestType returns generated.TestTypeResolver implementation.
func (r *Resolver) TestType() generated.TestTypeResolver { return &testTypeResolver{r} }

type mutationResolver struct{ *Resolver }
type queryResolver struct{ *Resolver }
type testTypeResolver struct{ *Resolver }
