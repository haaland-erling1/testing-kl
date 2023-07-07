// Code generated by github.com/99designs/gqlgen, DO NOT EDIT.

package model

import (
	"fmt"
	"io"
	"strconv"

	"kloudlite.io/apps/infra/internal/domain/entities"
)

type BYOCClusterEdge struct {
	Cursor string                `json:"cursor"`
	Node   *entities.BYOCCluster `json:"node"`
}

type BYOCClusterPaginatedRecords struct {
	Edges      []*BYOCClusterEdge `json:"edges"`
	PageInfo   *PageInfo          `json:"pageInfo"`
	TotalCount int                `json:"totalCount"`
}

type CloudProviderEdge struct {
	Cursor string                  `json:"cursor"`
	Node   *entities.CloudProvider `json:"node"`
}

type CloudProviderPaginatedRecords struct {
	Edges      []*CloudProviderEdge `json:"edges"`
	PageInfo   *PageInfo            `json:"pageInfo"`
	TotalCount int                  `json:"totalCount"`
}

type ClusterEdge struct {
	Cursor string            `json:"cursor"`
	Node   *entities.Cluster `json:"node"`
}

type ClusterPaginatedRecords struct {
	Edges      []*ClusterEdge `json:"edges"`
	PageInfo   *PageInfo      `json:"pageInfo"`
	TotalCount int            `json:"totalCount"`
}

type EdgeEdge struct {
	Cursor string         `json:"cursor"`
	Node   *entities.Edge `json:"node"`
}

type EdgePaginatedRecords struct {
	Edges      []*EdgeEdge `json:"edges"`
	PageInfo   *PageInfo   `json:"pageInfo"`
	TotalCount int         `json:"totalCount"`
}

type GithubComKloudliteClusterOperatorApisCmgrV1ClusterSpec struct {
	AccountName  string `json:"accountName"`
	Config       string `json:"config"`
	Count        int    `json:"count"`
	Provider     string `json:"provider"`
	ProviderName string `json:"providerName"`
	Region       string `json:"region"`
}

type GithubComKloudliteClusterOperatorApisCmgrV1ClusterSpecIn struct {
	AccountName  string `json:"accountName"`
	Config       string `json:"config"`
	Count        int    `json:"count"`
	Provider     string `json:"provider"`
	ProviderName string `json:"providerName"`
	Region       string `json:"region"`
}

type GithubComKloudliteClusterOperatorApisCmgrV1MasterNodeSpec struct {
	AccountName  string `json:"accountName"`
	ClusterName  string `json:"clusterName"`
	Config       string `json:"config"`
	Provider     string `json:"provider"`
	ProviderName string `json:"providerName"`
	Region       string `json:"region"`
}

type GithubComKloudliteClusterOperatorApisCmgrV1MasterNodeSpecIn struct {
	AccountName  string `json:"accountName"`
	ClusterName  string `json:"clusterName"`
	Config       string `json:"config"`
	Provider     string `json:"provider"`
	ProviderName string `json:"providerName"`
	Region       string `json:"region"`
}

type GithubComKloudliteClusterOperatorApisInfraV1CloudProviderSpec struct {
	AccountName    string                                                                       `json:"accountName"`
	DisplayName    string                                                                       `json:"display_name"`
	Provider       string                                                                       `json:"provider"`
	ProviderSecret *GithubComKloudliteClusterOperatorApisInfraV1CloudProviderSpecProviderSecret `json:"providerSecret"`
}

type GithubComKloudliteClusterOperatorApisInfraV1CloudProviderSpecIn struct {
	AccountName    string                                                                         `json:"accountName"`
	DisplayName    string                                                                         `json:"display_name"`
	Provider       string                                                                         `json:"provider"`
	ProviderSecret *GithubComKloudliteClusterOperatorApisInfraV1CloudProviderSpecProviderSecretIn `json:"providerSecret"`
}

type GithubComKloudliteClusterOperatorApisInfraV1CloudProviderSpecProviderSecret struct {
	Name      string `json:"name"`
	Namespace string `json:"namespace"`
}

type GithubComKloudliteClusterOperatorApisInfraV1CloudProviderSpecProviderSecretIn struct {
	Name      string `json:"name"`
	Namespace string `json:"namespace"`
}

type GithubComKloudliteClusterOperatorApisInfraV1EdgeSpec struct {
	AccountName  string                                                       `json:"accountName"`
	ClusterName  string                                                       `json:"clusterName"`
	Pools        []*GithubComKloudliteClusterOperatorApisInfraV1EdgeSpecPools `json:"pools,omitempty"`
	Provider     *string                                                      `json:"provider,omitempty"`
	ProviderName string                                                       `json:"providerName"`
	Region       string                                                       `json:"region"`
}

type GithubComKloudliteClusterOperatorApisInfraV1EdgeSpecIn struct {
	AccountName  string                                                         `json:"accountName"`
	ClusterName  string                                                         `json:"clusterName"`
	Pools        []*GithubComKloudliteClusterOperatorApisInfraV1EdgeSpecPoolsIn `json:"pools,omitempty"`
	Provider     *string                                                        `json:"provider,omitempty"`
	ProviderName string                                                         `json:"providerName"`
	Region       string                                                         `json:"region"`
}

type GithubComKloudliteClusterOperatorApisInfraV1EdgeSpecPools struct {
	Config string `json:"config"`
	Max    *int   `json:"max,omitempty"`
	Min    *int   `json:"min,omitempty"`
	Name   string `json:"name"`
}

type GithubComKloudliteClusterOperatorApisInfraV1EdgeSpecPoolsIn struct {
	Config string `json:"config"`
	Max    *int   `json:"max,omitempty"`
	Min    *int   `json:"min,omitempty"`
	Name   string `json:"name"`
}

type GithubComKloudliteClusterOperatorApisInfraV1NodePoolSpec struct {
	AccountName  string `json:"accountName"`
	ClusterName  string `json:"clusterName"`
	Config       string `json:"config"`
	EdgeName     string `json:"edgeName"`
	Max          *int   `json:"max,omitempty"`
	Min          *int   `json:"min,omitempty"`
	Provider     string `json:"provider"`
	ProviderName string `json:"providerName"`
	Region       string `json:"region"`
}

type GithubComKloudliteClusterOperatorApisInfraV1NodePoolSpecIn struct {
	AccountName  string `json:"accountName"`
	ClusterName  string `json:"clusterName"`
	Config       string `json:"config"`
	EdgeName     string `json:"edgeName"`
	Max          *int   `json:"max,omitempty"`
	Min          *int   `json:"min,omitempty"`
	Provider     string `json:"provider"`
	ProviderName string `json:"providerName"`
	Region       string `json:"region"`
}

type GithubComKloudliteClusterOperatorApisInfraV1WorkerNodeSpec struct {
	AccountName  string `json:"accountName"`
	ClusterName  string `json:"clusterName"`
	Config       string `json:"config"`
	EdgeName     string `json:"edgeName"`
	NodeIndex    *int   `json:"nodeIndex,omitempty"`
	Pool         string `json:"pool"`
	Provider     string `json:"provider"`
	ProviderName string `json:"providerName"`
	Region       string `json:"region"`
	Stateful     *bool  `json:"stateful,omitempty"`
}

type GithubComKloudliteClusterOperatorApisInfraV1WorkerNodeSpecIn struct {
	AccountName  string `json:"accountName"`
	ClusterName  string `json:"clusterName"`
	Config       string `json:"config"`
	EdgeName     string `json:"edgeName"`
	NodeIndex    *int   `json:"nodeIndex,omitempty"`
	Pool         string `json:"pool"`
	Provider     string `json:"provider"`
	ProviderName string `json:"providerName"`
	Region       string `json:"region"`
	Stateful     *bool  `json:"stateful,omitempty"`
}

type GithubComKloudliteOperatorApisClustersV1BYOCSpec struct {
	AccountName        string    `json:"accountName"`
	DisplayName        *string   `json:"displayName,omitempty"`
	IncomingKafkaTopic string    `json:"incomingKafkaTopic"`
	IngressClasses     []*string `json:"ingressClasses,omitempty"`
	Provider           string    `json:"provider"`
	PublicIps          []*string `json:"publicIps,omitempty"`
	Region             string    `json:"region"`
	StorageClasses     []*string `json:"storageClasses,omitempty"`
}

type GithubComKloudliteOperatorApisClustersV1BYOCSpecIn struct {
	AccountName        string    `json:"accountName"`
	DisplayName        *string   `json:"displayName,omitempty"`
	IncomingKafkaTopic string    `json:"incomingKafkaTopic"`
	IngressClasses     []*string `json:"ingressClasses,omitempty"`
	Provider           string    `json:"provider"`
	PublicIps          []*string `json:"publicIps,omitempty"`
	Region             string    `json:"region"`
	StorageClasses     []*string `json:"storageClasses,omitempty"`
}

type GithubComKloudliteOperatorPkgOperatorCheck struct {
	Generation *int    `json:"generation,omitempty"`
	Message    *string `json:"message,omitempty"`
	Status     bool    `json:"status"`
}

type GithubComKloudliteOperatorPkgOperatorResourceRef struct {
	APIVersion *string `json:"apiVersion,omitempty"`
	Kind       *string `json:"kind,omitempty"`
	Name       string  `json:"name"`
	Namespace  string  `json:"namespace"`
}

type GithubComKloudliteOperatorPkgRawJSONRawJSON struct {
	RawMessage interface{} `json:"RawMessage,omitempty"`
}

type KloudliteIoAppsInfraInternalDomainEntitiesHelmStatusVal struct {
	IsReady *bool  `json:"isReady,omitempty"`
	Message string `json:"message"`
}

type MasterNodeEdge struct {
	Cursor string               `json:"cursor"`
	Node   *entities.MasterNode `json:"node"`
}

type MasterNodePaginatedRecords struct {
	Edges      []*MasterNodeEdge `json:"edges"`
	PageInfo   *PageInfo         `json:"pageInfo"`
	TotalCount int               `json:"totalCount"`
}

type NodePoolEdge struct {
	Cursor string             `json:"cursor"`
	Node   *entities.NodePool `json:"node"`
}

type NodePoolPaginatedRecords struct {
	Edges      []*NodePoolEdge `json:"edges"`
	PageInfo   *PageInfo       `json:"pageInfo"`
	TotalCount int             `json:"totalCount"`
}

type PageInfo struct {
	EndCursor       *string `json:"endCursor,omitempty"`
	HasNextPage     bool    `json:"hasNextPage"`
	HasPreviousPage bool    `json:"hasPreviousPage"`
	StartCursor     *string `json:"startCursor,omitempty"`
}

type SecretEdge struct {
	Cursor string           `json:"cursor"`
	Node   *entities.Secret `json:"node"`
}

type SecretPaginatedRecords struct {
	Edges      []*SecretEdge `json:"edges"`
	PageInfo   *PageInfo     `json:"pageInfo"`
	TotalCount int           `json:"totalCount"`
}

type WorkerNodeEdge struct {
	Cursor string               `json:"cursor"`
	Node   *entities.WorkerNode `json:"node"`
}

type WorkerNodePaginatedRecords struct {
	Edges      []*WorkerNodeEdge `json:"edges"`
	PageInfo   *PageInfo         `json:"pageInfo"`
	TotalCount int               `json:"totalCount"`
}

type KloudliteIoPkgTypesSyncStatusAction string

const (
	KloudliteIoPkgTypesSyncStatusActionApply  KloudliteIoPkgTypesSyncStatusAction = "APPLY"
	KloudliteIoPkgTypesSyncStatusActionDelete KloudliteIoPkgTypesSyncStatusAction = "DELETE"
)

var AllKloudliteIoPkgTypesSyncStatusAction = []KloudliteIoPkgTypesSyncStatusAction{
	KloudliteIoPkgTypesSyncStatusActionApply,
	KloudliteIoPkgTypesSyncStatusActionDelete,
}

func (e KloudliteIoPkgTypesSyncStatusAction) IsValid() bool {
	switch e {
	case KloudliteIoPkgTypesSyncStatusActionApply, KloudliteIoPkgTypesSyncStatusActionDelete:
		return true
	}
	return false
}

func (e KloudliteIoPkgTypesSyncStatusAction) String() string {
	return string(e)
}

func (e *KloudliteIoPkgTypesSyncStatusAction) UnmarshalGQL(v interface{}) error {
	str, ok := v.(string)
	if !ok {
		return fmt.Errorf("enums must be strings")
	}

	*e = KloudliteIoPkgTypesSyncStatusAction(str)
	if !e.IsValid() {
		return fmt.Errorf("%s is not a valid Kloudlite_io__pkg__types_SyncStatusAction", str)
	}
	return nil
}

func (e KloudliteIoPkgTypesSyncStatusAction) MarshalGQL(w io.Writer) {
	fmt.Fprint(w, strconv.Quote(e.String()))
}

type KloudliteIoPkgTypesSyncStatusState string

const (
	KloudliteIoPkgTypesSyncStatusStateAppliedAtAgent          KloudliteIoPkgTypesSyncStatusState = "APPLIED_AT_AGENT"
	KloudliteIoPkgTypesSyncStatusStateErroredAtAgent          KloudliteIoPkgTypesSyncStatusState = "ERRORED_AT_AGENT"
	KloudliteIoPkgTypesSyncStatusStateIDLe                    KloudliteIoPkgTypesSyncStatusState = "IDLE"
	KloudliteIoPkgTypesSyncStatusStateInQueue                 KloudliteIoPkgTypesSyncStatusState = "IN_QUEUE"
	KloudliteIoPkgTypesSyncStatusStateReceivedUpdateFromAgent KloudliteIoPkgTypesSyncStatusState = "RECEIVED_UPDATE_FROM_AGENT"
)

var AllKloudliteIoPkgTypesSyncStatusState = []KloudliteIoPkgTypesSyncStatusState{
	KloudliteIoPkgTypesSyncStatusStateAppliedAtAgent,
	KloudliteIoPkgTypesSyncStatusStateErroredAtAgent,
	KloudliteIoPkgTypesSyncStatusStateIDLe,
	KloudliteIoPkgTypesSyncStatusStateInQueue,
	KloudliteIoPkgTypesSyncStatusStateReceivedUpdateFromAgent,
}

func (e KloudliteIoPkgTypesSyncStatusState) IsValid() bool {
	switch e {
	case KloudliteIoPkgTypesSyncStatusStateAppliedAtAgent, KloudliteIoPkgTypesSyncStatusStateErroredAtAgent, KloudliteIoPkgTypesSyncStatusStateIDLe, KloudliteIoPkgTypesSyncStatusStateInQueue, KloudliteIoPkgTypesSyncStatusStateReceivedUpdateFromAgent:
		return true
	}
	return false
}

func (e KloudliteIoPkgTypesSyncStatusState) String() string {
	return string(e)
}

func (e *KloudliteIoPkgTypesSyncStatusState) UnmarshalGQL(v interface{}) error {
	str, ok := v.(string)
	if !ok {
		return fmt.Errorf("enums must be strings")
	}

	*e = KloudliteIoPkgTypesSyncStatusState(str)
	if !e.IsValid() {
		return fmt.Errorf("%s is not a valid Kloudlite_io__pkg__types_SyncStatusState", str)
	}
	return nil
}

func (e KloudliteIoPkgTypesSyncStatusState) MarshalGQL(w io.Writer) {
	fmt.Fprint(w, strconv.Quote(e.String()))
}

type PaginationSortOrder string

const (
	PaginationSortOrderAsc  PaginationSortOrder = "ASC"
	PaginationSortOrderDesc PaginationSortOrder = "DESC"
)

var AllPaginationSortOrder = []PaginationSortOrder{
	PaginationSortOrderAsc,
	PaginationSortOrderDesc,
}

func (e PaginationSortOrder) IsValid() bool {
	switch e {
	case PaginationSortOrderAsc, PaginationSortOrderDesc:
		return true
	}
	return false
}

func (e PaginationSortOrder) String() string {
	return string(e)
}

func (e *PaginationSortOrder) UnmarshalGQL(v interface{}) error {
	str, ok := v.(string)
	if !ok {
		return fmt.Errorf("enums must be strings")
	}

	*e = PaginationSortOrder(str)
	if !e.IsValid() {
		return fmt.Errorf("%s is not a valid PaginationSortOrder", str)
	}
	return nil
}

func (e PaginationSortOrder) MarshalGQL(w io.Writer) {
	fmt.Fprint(w, strconv.Quote(e.String()))
}
