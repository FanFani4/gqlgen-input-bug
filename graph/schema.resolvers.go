package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	"fmt"

	"github.com/FanFani4/gqlgen-input-bug/domain"
	"github.com/FanFani4/gqlgen-input-bug/graph/generated"
	"github.com/FanFani4/gqlgen-input-bug/graph/model"
)

func (r *mutationResolver) Foo(ctx context.Context, in domain.TestInput) (*domain.TestType, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *queryResolver) Foo(ctx context.Context, in domain.TestInput) (*domain.TestType, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *testTypeResolver) Role(ctx context.Context, obj *domain.TestType) (model.UserRole, error) {
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

// !!! WARNING !!!
// The code below was going to be deleted when updating resolvers. It has been copied here so you have
// one last chance to move it out of harms way if you want. There are two reasons this happens:
//  - When renaming or deleting a resolver the old code will be put in here. You can safely delete
//    it when you're done.
//  - You have helper methods in this file. Move them out to keep these resolver files clean.
func (r *testInputResolver) Role(ctx context.Context, obj *domain.TestInput, role model.UserRole) error {
	panic(fmt.Errorf("not implemented"))
}
func (r *Resolver) TestInput() generated.TestInputResolver { return &testInputResolver{r} }

type testInputResolver struct{ *Resolver }
