package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.
// Code generated by github.com/99designs/gqlgen version v0.17.45

import (
	"context"
	"github.com/kloudlite/api/apps/container-registry/internal/app/graph/generated"
	"github.com/kloudlite/api/apps/container-registry/internal/app/graph/model"
	"github.com/kloudlite/api/apps/container-registry/internal/domain/entities"
	"github.com/kloudlite/api/pkg/errors"
	fn "github.com/kloudlite/api/pkg/functions"
)

// Repositories is the resolver for the repositories field.
func (r *githubSearchRepositoryResolver) Repositories(ctx context.Context, obj *entities.GithubSearchRepository) ([]*model.GithubComKloudliteAPIAppsContainerRegistryInternalDomainEntitiesGithubRepository, error) {
	if obj == nil {
		return nil, errors.Newf("Repositories: obj is nil")
	}

	var repositories []*model.GithubComKloudliteAPIAppsContainerRegistryInternalDomainEntitiesGithubRepository
	if err := fn.JsonConversion(obj.Repositories, &repositories); err != nil {
		return nil, err
	}

	return repositories, nil
}

// GithubSearchRepository returns generated.GithubSearchRepositoryResolver implementation.
func (r *Resolver) GithubSearchRepository() generated.GithubSearchRepositoryResolver {
	return &githubSearchRepositoryResolver{r}
}

type githubSearchRepositoryResolver struct{ *Resolver }
