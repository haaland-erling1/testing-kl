package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.
// Code generated by github.com/99designs/gqlgen version v0.17.45

import (
	"context"
	"fmt"

	"github.com/kloudlite/api/apps/comms/internal/app/graph/generated"
	"github.com/kloudlite/api/common"
)

// UserID is the resolver for the userId field.
func (r *github__com___kloudlite___api___common__CreatedOrUpdatedByResolver) UserID(ctx context.Context, obj *common.CreatedOrUpdatedBy) (string, error) {
	if obj == nil {
		return "", fmt.Errorf("obj is nil")
	}

	return string(obj.UserId), nil
}

// Github__com___kloudlite___api___common__CreatedOrUpdatedBy returns generated.Github__com___kloudlite___api___common__CreatedOrUpdatedByResolver implementation.
func (r *Resolver) Github__com___kloudlite___api___common__CreatedOrUpdatedBy() generated.Github__com___kloudlite___api___common__CreatedOrUpdatedByResolver {
	return &github__com___kloudlite___api___common__CreatedOrUpdatedByResolver{r}
}

type github__com___kloudlite___api___common__CreatedOrUpdatedByResolver struct{ *Resolver }
