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
)

// Access is the resolver for the access field.
func (r *credentialResolver) Access(ctx context.Context, obj *entities.Credential) (model.GithubComKloudliteAPIAppsContainerRegistryInternalDomainEntitiesRepoAccess, error) {
	return model.GithubComKloudliteAPIAppsContainerRegistryInternalDomainEntitiesRepoAccess(obj.Access), nil
}

// CreationTime is the resolver for the creationTime field.
func (r *credentialResolver) CreationTime(ctx context.Context, obj *entities.Credential) (string, error) {
	if obj == nil {
		return "", errors.Newf("resource is nil")
	}

	return obj.CreationTime.Format(time.RFC3339), nil
}

// Expiration is the resolver for the expiration field.
func (r *credentialResolver) Expiration(ctx context.Context, obj *entities.Credential) (*model.GithubComKloudliteAPIAppsContainerRegistryInternalDomainEntitiesExpiration, error) {
	if obj == nil {
		return nil, errors.Newf("resource is nil")
	}

	return &model.GithubComKloudliteAPIAppsContainerRegistryInternalDomainEntitiesExpiration{
		Value: obj.Expiration.Value,
		Unit:  model.GithubComKloudliteAPIAppsContainerRegistryInternalDomainEntitiesExpirationUnit(obj.Expiration.Unit),
	}, nil
}

// UpdateTime is the resolver for the updateTime field.
func (r *credentialResolver) UpdateTime(ctx context.Context, obj *entities.Credential) (string, error) {
	if obj == nil {
		return "", errors.Newf("resource is nil")
	}

	return obj.UpdateTime.Format(time.RFC3339), nil
}

// Access is the resolver for the access field.
func (r *credentialInResolver) Access(ctx context.Context, obj *entities.Credential, data model.GithubComKloudliteAPIAppsContainerRegistryInternalDomainEntitiesRepoAccess) error {
	if obj == nil {
		return errors.Newf("resource is nil")
	}

	obj.Access = entities.RepoAccess(data)
	return nil
}

// Expiration is the resolver for the expiration field.
func (r *credentialInResolver) Expiration(ctx context.Context, obj *entities.Credential, data *model.GithubComKloudliteAPIAppsContainerRegistryInternalDomainEntitiesExpirationIn) error {
	if obj == nil {
		return errors.Newf("resource is nil")
	}

	obj.Expiration = entities.Expiration{
		Value: data.Value,
		Unit:  entities.ExpirationUnit(data.Unit),
	}
	return nil
}

// Credential returns generated.CredentialResolver implementation.
func (r *Resolver) Credential() generated.CredentialResolver { return &credentialResolver{r} }

// CredentialIn returns generated.CredentialInResolver implementation.
func (r *Resolver) CredentialIn() generated.CredentialInResolver { return &credentialInResolver{r} }

type credentialResolver struct{ *Resolver }
type credentialInResolver struct{ *Resolver }