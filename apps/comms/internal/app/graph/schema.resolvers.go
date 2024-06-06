package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.
// Code generated by github.com/99designs/gqlgen version v0.17.45

import (
	"context"
	"github.com/kloudlite/api/apps/comms/internal/app/graph/generated"
	"github.com/kloudlite/api/apps/comms/internal/app/graph/model"
	"github.com/kloudlite/api/apps/comms/internal/domain/entities"
	"github.com/kloudlite/api/apps/comms/types"
	"github.com/kloudlite/api/pkg/errors"
	fn "github.com/kloudlite/api/pkg/functions"
	"github.com/kloudlite/api/pkg/repos"
)

// CommsUpdateNotificationConfig is the resolver for the comms_updateNotificationConfig field.
func (r *mutationResolver) CommsUpdateNotificationConfig(ctx context.Context, config entities.NotificationConf) (*entities.NotificationConf, error) {
	cc, err := toCommsContext(ctx)
	if err != nil {
		return nil, errors.NewE(err)
	}

	return r.Domain.UpdateNotificationConfig(cc, config)
}

// CommsUpdateSubscriptionConfig is the resolver for the comms_updateSubscriptionConfig field.
func (r *mutationResolver) CommsUpdateSubscriptionConfig(ctx context.Context, config entities.Subscription, id repos.ID) (*entities.Subscription, error) {
	cc, err := toCommsContext(ctx)
	if err != nil {
		return nil, errors.NewE(err)
	}

	return r.Domain.UpdateSubscriptionConfig(cc, id, config)
}

// CommsMarkNotificationAsRead is the resolver for the comms_markNotificationAsRead field.
func (r *mutationResolver) CommsMarkNotificationAsRead(ctx context.Context, id repos.ID) (*types.Notification, error) {
	cc, err := toCommsContext(ctx)
	if err != nil {
		return nil, errors.NewE(err)
	}

	return r.Domain.MarkNotificationAsRead(cc, id)
}

// CommsListNotifications is the resolver for the comms_listNotifications field.
func (r *queryResolver) CommsListNotifications(ctx context.Context, pagination *repos.CursorPagination) (*model.NotificationPaginatedRecords, error) {
	cc, err := toCommsContext(ctx)
	if err != nil {
		return nil, errors.NewE(err)
	}

	pr, err := r.Domain.ListNotifications(cc, fn.DefaultIfNil(pagination, repos.DefaultCursorPagination))
	if err != nil {
		return nil, errors.NewE(err)
	}

	return fn.JsonConvertP[model.NotificationPaginatedRecords](pr)
}

// CommsGetNotificationConfig is the resolver for the comms_getNotificationConfig field.
func (r *queryResolver) CommsGetNotificationConfig(ctx context.Context) (*entities.NotificationConf, error) {
	cc, err := toCommsContext(ctx)
	if err != nil {
		return nil, errors.NewE(err)
	}

	return r.Domain.GetNotificationConfig(cc)
}

// CommsGetSubscriptionConfig is the resolver for the comms_getSubscriptionConfig field.
func (r *queryResolver) CommsGetSubscriptionConfig(ctx context.Context, id repos.ID) (*entities.Subscription, error) {
	cc, err := toCommsContext(ctx)
	if err != nil {
		return nil, errors.NewE(err)
	}

	return r.Domain.GetSubscriptionConfig(cc, id)
}

// Mutation returns generated.MutationResolver implementation.
func (r *Resolver) Mutation() generated.MutationResolver { return &mutationResolver{r} }

// Query returns generated.QueryResolver implementation.
func (r *Resolver) Query() generated.QueryResolver { return &queryResolver{r} }

type mutationResolver struct{ *Resolver }
type queryResolver struct{ *Resolver }
