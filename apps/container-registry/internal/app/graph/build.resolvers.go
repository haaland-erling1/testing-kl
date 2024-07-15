package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.
// Code generated by github.com/99designs/gqlgen version v0.17.45

import (
	"context"
	"github.com/kloudlite/api/pkg/errors"
	"time"

	"github.com/kloudlite/api/apps/container-registry/internal/app/graph/generated"
	"github.com/kloudlite/api/apps/container-registry/internal/app/graph/model"
	"github.com/kloudlite/api/apps/container-registry/internal/domain/entities"
	fn "github.com/kloudlite/api/pkg/functions"
)

// CreationTime is the resolver for the creationTime field.
func (r *buildResolver) CreationTime(ctx context.Context, obj *entities.Build) (string, error) {
	if obj == nil {
		return "", errors.Newf("build is nil")
	}

	return obj.CreationTime.Format(time.RFC3339), nil
}

// ErrorMessages is the resolver for the errorMessages field.
func (r *buildResolver) ErrorMessages(ctx context.Context, obj *entities.Build) (map[string]interface{}, error) {
	if obj == nil {
		return nil, errors.Newf("build is nil")
	}

	var m map[string]any
	if err := fn.JsonConversion(obj.ErrorMessages, &m); err != nil {
		return nil, errors.NewE(err)
	}

	return m, nil
}

// Source is the resolver for the source field.
func (r *buildResolver) Source(ctx context.Context, obj *entities.Build) (*model.GithubComKloudliteAPIAppsContainerRegistryInternalDomainEntitiesGitSource, error) {
	if obj == nil {
		return nil, errors.Newf("build is nil")
	}
	var m model.GithubComKloudliteAPIAppsContainerRegistryInternalDomainEntitiesGitSource
	if err := fn.JsonConversion(obj.Source, &m); err != nil {
		return nil, err
	}

	return &m, nil
}

// Spec is the resolver for the spec field.
func (r *buildResolver) Spec(ctx context.Context, obj *entities.Build) (*model.GithubComKloudliteOperatorApisDistributionV1BuildRunSpec, error) {
	if obj == nil {
		return nil, errors.Newf("build is nil")
	}
	var m model.GithubComKloudliteOperatorApisDistributionV1BuildRunSpec
	if err := fn.JsonConversion(obj.Spec, &m); err != nil {
		return nil, err
	}
	return &m, nil
}

// Status is the resolver for the status field.
func (r *buildResolver) Status(ctx context.Context, obj *entities.Build) (model.GithubComKloudliteAPIAppsContainerRegistryInternalDomainEntitiesBuildStatus, error) {
	if obj == nil {
		return "", errors.Newf("build is nil")
	}

	return model.GithubComKloudliteAPIAppsContainerRegistryInternalDomainEntitiesBuildStatus(obj.Status), nil
}

// UpdateTime is the resolver for the updateTime field.
func (r *buildResolver) UpdateTime(ctx context.Context, obj *entities.Build) (string, error) {
	if obj == nil {
		return "", errors.Newf("build is nil")
	}

	return obj.UpdateTime.Format(time.RFC3339), nil
}

// Source is the resolver for the source field.
func (r *buildInResolver) Source(ctx context.Context, obj *entities.Build, data *model.GithubComKloudliteAPIAppsContainerRegistryInternalDomainEntitiesGitSourceIn) error {
	if obj == nil {
		return errors.Newf("build is nil")
	}

	return fn.JsonConversion(data, &obj.Source)
}

// Spec is the resolver for the spec field.
func (r *buildInResolver) Spec(ctx context.Context, obj *entities.Build, data *model.GithubComKloudliteOperatorApisDistributionV1BuildRunSpecIn) error {
	if obj == nil {
		return errors.Newf("build is nil")
	}

	return fn.JsonConversion(data, &obj.Spec)
}

// Build returns generated.BuildResolver implementation.
func (r *Resolver) Build() generated.BuildResolver { return &buildResolver{r} }

// BuildIn returns generated.BuildInResolver implementation.
func (r *Resolver) BuildIn() generated.BuildInResolver { return &buildInResolver{r} }

type buildResolver struct{ *Resolver }
type buildInResolver struct{ *Resolver }
