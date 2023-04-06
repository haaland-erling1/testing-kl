package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.
// Code generated by github.com/99designs/gqlgen version v0.17.28

import (
	"context"

	"kloudlite.io/apps/infra/internal/app/graph/generated"
	"kloudlite.io/apps/infra/internal/domain"
	"kloudlite.io/apps/infra/internal/domain/entities"
)

// InfraCreateBYOCCluster is the resolver for the infra_createBYOCCluster field.
func (r *mutationResolver) InfraCreateBYOCCluster(ctx context.Context, cluster entities.BYOCCluster) (*entities.BYOCCluster, error) {
	return r.Domain.CreateBYOCCluster(toInfraContext(ctx), cluster)
}

// InfraUpdateBYOCCluster is the resolver for the infra_updateBYOCCluster field.
func (r *mutationResolver) InfraUpdateBYOCCluster(ctx context.Context, cluster entities.BYOCCluster) (*entities.BYOCCluster, error) {
	return r.Domain.UpdateBYOCCluster(toInfraContext(ctx), cluster)
}

// InfraDeleteBYOCCluster is the resolver for the infra_deleteBYOCCluster field.
func (r *mutationResolver) InfraDeleteBYOCCluster(ctx context.Context, name string) (bool, error) {
	err := r.Domain.DeleteBYOCCluster(toInfraContext(ctx), name)
	if err != nil {
		return false, err
	}
	return true, nil
}

// InfraCreateCluster is the resolver for the infra_createCluster field.
func (r *mutationResolver) InfraCreateCluster(ctx context.Context, cluster entities.Cluster) (*entities.Cluster, error) {
	return r.Domain.CreateCluster(toInfraContext(ctx), cluster)
}

// InfraUpdateCluster is the resolver for the infra_updateCluster field.
func (r *mutationResolver) InfraUpdateCluster(ctx context.Context, cluster entities.Cluster) (*entities.Cluster, error) {
	return r.Domain.UpdateCluster(toInfraContext(ctx), cluster)
}

// InfraDeleteCluster is the resolver for the infra_deleteCluster field.
func (r *mutationResolver) InfraDeleteCluster(ctx context.Context, name string) (bool, error) {
	if err := r.Domain.DeleteCluster(toInfraContext(ctx), name); err != nil {
		return false, err
	}
	return true, nil
}

// InfraCreateCloudProvider is the resolver for the infra_createCloudProvider field.
func (r *mutationResolver) InfraCreateCloudProvider(ctx context.Context, cloudProvider entities.CloudProvider, providerSecret entities.Secret) (*entities.CloudProvider, error) {
	return r.Domain.CreateCloudProvider(toInfraContext(ctx), cloudProvider, providerSecret)
}

// InfraUpdateCloudProvider is the resolver for the infra_updateCloudProvider field.
func (r *mutationResolver) InfraUpdateCloudProvider(ctx context.Context, cloudProvider entities.CloudProvider, providerSecret *entities.Secret) (*entities.CloudProvider, error) {
	return r.Domain.UpdateCloudProvider(toInfraContext(ctx), cloudProvider, providerSecret)
}

// InfraDeleteCloudProvider is the resolver for the infra_deleteCloudProvider field.
func (r *mutationResolver) InfraDeleteCloudProvider(ctx context.Context, name string) (bool, error) {
	if err := r.Domain.DeleteCloudProvider(toInfraContext(ctx), name); err != nil {
		return false, err
	}
	return true, nil
}

// InfraCreateEdge is the resolver for the infra_createEdge field.
func (r *mutationResolver) InfraCreateEdge(ctx context.Context, edge entities.Edge) (*entities.Edge, error) {
	return r.Domain.CreateEdge(toInfraContext(ctx), edge)
}

// InfraUpdateEdge is the resolver for the infra_updateEdge field.
func (r *mutationResolver) InfraUpdateEdge(ctx context.Context, edge entities.Edge) (*entities.Edge, error) {
	return r.Domain.UpdateEdge(toInfraContext(ctx), edge)
}

// InfraDeleteEdge is the resolver for the infra_deleteEdge field.
func (r *mutationResolver) InfraDeleteEdge(ctx context.Context, clusterName string, name string) (bool, error) {
	if err := r.Domain.DeleteEdge(toInfraContext(ctx), clusterName, name); err != nil {
		return false, err
	}
	return true, nil
}

// InfraDeleteWorkerNode is the resolver for the infra_deleteWorkerNode field.
func (r *mutationResolver) InfraDeleteWorkerNode(ctx context.Context, clusterName string, edgeName string, name string) (bool, error) {
	return r.Domain.DeleteWorkerNode(toInfraContext(ctx), clusterName, edgeName, name)
}

// InfraCheckNameAvailability is the resolver for the infra_checkNameAvailability field.
func (r *queryResolver) InfraCheckNameAvailability(ctx context.Context, resType domain.ResType, name string) (*domain.CheckNameAvailabilityOutput, error) {
	return r.Domain.CheckNameAvailability(toInfraContext(ctx), resType, name)
}

// InfraListBYOCClusters is the resolver for the infra_listBYOCClusters field.
func (r *queryResolver) InfraListBYOCClusters(ctx context.Context) ([]*entities.BYOCCluster, error) {
	clusters, err := r.Domain.ListBYOCClusters(toInfraContext(ctx))
	if clusters == nil {
		clusters = make([]*entities.BYOCCluster, 0)
	}
	return clusters, err
}

// InfraGetBYOCCluster is the resolver for the infra_getBYOCCluster field.
func (r *queryResolver) InfraGetBYOCCluster(ctx context.Context, name string) (*entities.BYOCCluster, error) {
	return r.Domain.GetBYOCCluster(toInfraContext(ctx), name)
}

// InfraListClusters is the resolver for the infra_listClusters field.
func (r *queryResolver) InfraListClusters(ctx context.Context) ([]*entities.Cluster, error) {
	cls, err := r.Domain.ListClusters(toInfraContext(ctx))
	if cls == nil {
		cls = make([]*entities.Cluster, 0)
	}
	return cls, err
}

// InfraGetCluster is the resolver for the infra_getCluster field.
func (r *queryResolver) InfraGetCluster(ctx context.Context, name string) (*entities.Cluster, error) {
	return r.Domain.GetCluster(toInfraContext(ctx), name)
}

// InfraListCloudProviders is the resolver for the infra_listCloudProviders field.
func (r *queryResolver) InfraListCloudProviders(ctx context.Context) ([]*entities.CloudProvider, error) {
	cp, err := r.Domain.ListCloudProviders(toInfraContext(ctx))
	if cp == nil {
		cp = make([]*entities.CloudProvider, 0)
	}
	return cp, err
}

// InfraGetCloudProvider is the resolver for the infra_getCloudProvider field.
func (r *queryResolver) InfraGetCloudProvider(ctx context.Context, name string) (*entities.CloudProvider, error) {
	return r.Domain.GetCloudProvider(toInfraContext(ctx), name)
}

// InfraListEdges is the resolver for the infra_listEdges field.
func (r *queryResolver) InfraListEdges(ctx context.Context, clusterName string, providerName *string) ([]*entities.Edge, error) {
	e, err := r.Domain.ListEdges(toInfraContext(ctx), clusterName, providerName)
	if e == nil {
		e = make([]*entities.Edge, 0)
	}
	return e, err
}

// InfraGetEdge is the resolver for the infra_getEdge field.
func (r *queryResolver) InfraGetEdge(ctx context.Context, clusterName string, name string) (*entities.Edge, error) {
	return r.Domain.GetEdge(toInfraContext(ctx), clusterName, name)
}

// InfraGetMasterNodes is the resolver for the infra_getMasterNodes field.
func (r *queryResolver) InfraGetMasterNodes(ctx context.Context, clusterName string) ([]*entities.MasterNode, error) {
	return r.Domain.GetMasterNodes(toInfraContext(ctx), clusterName)
}

// InfraGetWorkerNodes is the resolver for the infra_getWorkerNodes field.
func (r *queryResolver) InfraGetWorkerNodes(ctx context.Context, clusterName string, edgeName string) ([]*entities.WorkerNode, error) {
	return r.Domain.GetWorkerNodes(toInfraContext(ctx), clusterName, edgeName)
}

// InfraGetNodePools is the resolver for the infra_getNodePools field.
func (r *queryResolver) InfraGetNodePools(ctx context.Context, clusterName string, edgeName string) ([]*entities.NodePool, error) {
	return r.Domain.GetNodePools(toInfraContext(ctx), clusterName, edgeName)
}

// Mutation returns generated.MutationResolver implementation.
func (r *Resolver) Mutation() generated.MutationResolver { return &mutationResolver{r} }

// Query returns generated.QueryResolver implementation.
func (r *Resolver) Query() generated.QueryResolver { return &queryResolver{r} }

type mutationResolver struct{ *Resolver }
type queryResolver struct{ *Resolver }
