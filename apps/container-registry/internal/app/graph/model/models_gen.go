// Code generated by github.com/99designs/gqlgen, DO NOT EDIT.

package model

import (
	"fmt"
	"io"
	"strconv"

	"kloudlite.io/apps/container-registry/internal/domain/entities"
	"kloudlite.io/pkg/repos"
)

type CredentialEdge struct {
	Cursor string               `json:"cursor"`
	Node   *entities.Credential `json:"node"`
}

type CredentialPaginatedRecords struct {
	Edges      []*CredentialEdge `json:"edges"`
	PageInfo   *PageInfo         `json:"pageInfo"`
	TotalCount int               `json:"totalCount"`
}

type KloudliteIoAppsContainerRegistryInternalDomainEntitiesExpiration struct {
	Unit  KloudliteIoAppsContainerRegistryInternalDomainEntitiesExpirationUnit `json:"unit"`
	Value int                                                                  `json:"value"`
}

type KloudliteIoAppsContainerRegistryInternalDomainEntitiesExpirationIn struct {
	Unit  KloudliteIoAppsContainerRegistryInternalDomainEntitiesExpirationUnit `json:"unit"`
	Value int                                                                  `json:"value"`
}

type KloudliteIoAppsContainerRegistryInternalDomainEntitiesRepoReference struct {
	Digest    string `json:"digest"`
	MediaType string `json:"mediaType"`
	Size      int    `json:"size"`
}

type PageInfo struct {
	EndCursor       *string `json:"endCursor,omitempty"`
	HasNextPage     *bool   `json:"hasNextPage,omitempty"`
	HasPreviousPage *bool   `json:"hasPreviousPage,omitempty"`
	StartCursor     *string `json:"startCursor,omitempty"`
}

type RepositoryEdge struct {
	Cursor string               `json:"cursor"`
	Node   *entities.Repository `json:"node"`
}

type RepositoryPaginatedRecords struct {
	Edges      []*RepositoryEdge `json:"edges"`
	PageInfo   *PageInfo         `json:"pageInfo"`
	TotalCount int               `json:"totalCount"`
}

type SearchCreds struct {
	Text *repos.MatchFilter `json:"text,omitempty"`
}

type SearchRepos struct {
	Text *repos.MatchFilter `json:"text,omitempty"`
}

type TagEdge struct {
	Cursor string        `json:"cursor"`
	Node   *entities.Tag `json:"node"`
}

type TagPaginatedRecords struct {
	Edges      []*TagEdge `json:"edges"`
	PageInfo   *PageInfo  `json:"pageInfo"`
	TotalCount int        `json:"totalCount"`
}

type KloudliteIoAppsContainerRegistryInternalDomainEntitiesExpirationUnit string

const (
	KloudliteIoAppsContainerRegistryInternalDomainEntitiesExpirationUnitD KloudliteIoAppsContainerRegistryInternalDomainEntitiesExpirationUnit = "d"
	KloudliteIoAppsContainerRegistryInternalDomainEntitiesExpirationUnitH KloudliteIoAppsContainerRegistryInternalDomainEntitiesExpirationUnit = "h"
	KloudliteIoAppsContainerRegistryInternalDomainEntitiesExpirationUnitM KloudliteIoAppsContainerRegistryInternalDomainEntitiesExpirationUnit = "m"
	KloudliteIoAppsContainerRegistryInternalDomainEntitiesExpirationUnitW KloudliteIoAppsContainerRegistryInternalDomainEntitiesExpirationUnit = "w"
	KloudliteIoAppsContainerRegistryInternalDomainEntitiesExpirationUnitY KloudliteIoAppsContainerRegistryInternalDomainEntitiesExpirationUnit = "y"
)

var AllKloudliteIoAppsContainerRegistryInternalDomainEntitiesExpirationUnit = []KloudliteIoAppsContainerRegistryInternalDomainEntitiesExpirationUnit{
	KloudliteIoAppsContainerRegistryInternalDomainEntitiesExpirationUnitD,
	KloudliteIoAppsContainerRegistryInternalDomainEntitiesExpirationUnitH,
	KloudliteIoAppsContainerRegistryInternalDomainEntitiesExpirationUnitM,
	KloudliteIoAppsContainerRegistryInternalDomainEntitiesExpirationUnitW,
	KloudliteIoAppsContainerRegistryInternalDomainEntitiesExpirationUnitY,
}

func (e KloudliteIoAppsContainerRegistryInternalDomainEntitiesExpirationUnit) IsValid() bool {
	switch e {
	case KloudliteIoAppsContainerRegistryInternalDomainEntitiesExpirationUnitD, KloudliteIoAppsContainerRegistryInternalDomainEntitiesExpirationUnitH, KloudliteIoAppsContainerRegistryInternalDomainEntitiesExpirationUnitM, KloudliteIoAppsContainerRegistryInternalDomainEntitiesExpirationUnitW, KloudliteIoAppsContainerRegistryInternalDomainEntitiesExpirationUnitY:
		return true
	}
	return false
}

func (e KloudliteIoAppsContainerRegistryInternalDomainEntitiesExpirationUnit) String() string {
	return string(e)
}

func (e *KloudliteIoAppsContainerRegistryInternalDomainEntitiesExpirationUnit) UnmarshalGQL(v interface{}) error {
	str, ok := v.(string)
	if !ok {
		return fmt.Errorf("enums must be strings")
	}

	*e = KloudliteIoAppsContainerRegistryInternalDomainEntitiesExpirationUnit(str)
	if !e.IsValid() {
		return fmt.Errorf("%s is not a valid Kloudlite_io__apps__container___registry__internal__domain__entities_ExpirationUnit", str)
	}
	return nil
}

func (e KloudliteIoAppsContainerRegistryInternalDomainEntitiesExpirationUnit) MarshalGQL(w io.Writer) {
	fmt.Fprint(w, strconv.Quote(e.String()))
}

type KloudliteIoAppsContainerRegistryInternalDomainEntitiesRepoAccess string

const (
	KloudliteIoAppsContainerRegistryInternalDomainEntitiesRepoAccessRead      KloudliteIoAppsContainerRegistryInternalDomainEntitiesRepoAccess = "read"
	KloudliteIoAppsContainerRegistryInternalDomainEntitiesRepoAccessReadWrite KloudliteIoAppsContainerRegistryInternalDomainEntitiesRepoAccess = "read_write"
)

var AllKloudliteIoAppsContainerRegistryInternalDomainEntitiesRepoAccess = []KloudliteIoAppsContainerRegistryInternalDomainEntitiesRepoAccess{
	KloudliteIoAppsContainerRegistryInternalDomainEntitiesRepoAccessRead,
	KloudliteIoAppsContainerRegistryInternalDomainEntitiesRepoAccessReadWrite,
}

func (e KloudliteIoAppsContainerRegistryInternalDomainEntitiesRepoAccess) IsValid() bool {
	switch e {
	case KloudliteIoAppsContainerRegistryInternalDomainEntitiesRepoAccessRead, KloudliteIoAppsContainerRegistryInternalDomainEntitiesRepoAccessReadWrite:
		return true
	}
	return false
}

func (e KloudliteIoAppsContainerRegistryInternalDomainEntitiesRepoAccess) String() string {
	return string(e)
}

func (e *KloudliteIoAppsContainerRegistryInternalDomainEntitiesRepoAccess) UnmarshalGQL(v interface{}) error {
	str, ok := v.(string)
	if !ok {
		return fmt.Errorf("enums must be strings")
	}

	*e = KloudliteIoAppsContainerRegistryInternalDomainEntitiesRepoAccess(str)
	if !e.IsValid() {
		return fmt.Errorf("%s is not a valid Kloudlite_io__apps__container___registry__internal__domain__entities_RepoAccess", str)
	}
	return nil
}

func (e KloudliteIoAppsContainerRegistryInternalDomainEntitiesRepoAccess) MarshalGQL(w io.Writer) {
	fmt.Fprint(w, strconv.Quote(e.String()))
}
