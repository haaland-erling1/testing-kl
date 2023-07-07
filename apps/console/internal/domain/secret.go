package domain

import (
	"fmt"
	"time"

	"kloudlite.io/apps/console/internal/domain/entities"
	"kloudlite.io/pkg/repos"
	t "kloudlite.io/pkg/types"
)

// query

func (d *domain) ListSecrets(ctx ConsoleContext, namespace string, pq t.CursorPagination) (*repos.PaginatedRecord[*entities.Secret], error) {
	if err := d.canReadResourcesInWorkspace(ctx, namespace); err != nil {
		return nil, err
	}

	return d.secretRepo.FindPaginated(ctx, repos.Filter{
		"accountName":        ctx.AccountName,
		"clusterName":        ctx.ClusterName,
		"metadata.namespace": namespace,
	}, pq)
}

func (d *domain) findSecret(ctx ConsoleContext, namespace string, name string) (*entities.Secret, error) {
	exSecret, err := d.secretRepo.FindOne(ctx, repos.Filter{
		"accountName":        ctx.AccountName,
		"clusterName":        ctx.ClusterName,
		"metadata.namespace": namespace,
		"metadata.name":      name,
	})
	if err != nil {
		return nil, err
	}
	if exSecret == nil {
		return nil, fmt.Errorf("no secret with name=%s,namespace=%s found", name, namespace)
	}
	return exSecret, nil
}

func (d *domain) GetSecret(ctx ConsoleContext, namespace string, name string) (*entities.Secret, error) {
	if err := d.canReadResourcesInWorkspace(ctx, namespace); err != nil {
		return nil, err
	}
	return d.findSecret(ctx, namespace, name)
}

// mutations

func (d *domain) CreateSecret(ctx ConsoleContext, secret entities.Secret) (*entities.Secret, error) {
	if err := d.canMutateResourcesInWorkspace(ctx, secret.Namespace); err != nil {
		return nil, err
	}

	secret.EnsureGVK()
	if err := d.k8sExtendedClient.ValidateStruct(ctx, &secret.Secret); err != nil {
		return nil, err
	}

  secret.IncrementRecordVersion()
	secret.AccountName = ctx.AccountName
	secret.ClusterName = ctx.ClusterName
	secret.SyncStatus = t.GenSyncStatus(t.SyncActionApply, secret.RecordVersion)

	s, err := d.secretRepo.Create(ctx, &secret)
	if err != nil {
		if d.secretRepo.ErrAlreadyExists(err) {
			// TODO: better insights into error, when it is being caused by duplicated indexes
			return nil, err
		}
		return nil, err
	}

	if err := d.applyK8sResource(ctx, &s.Secret, s.RecordVersion); err != nil {
		return s, err
	}

	return s, nil
}

func (d *domain) UpdateSecret(ctx ConsoleContext, secret entities.Secret) (*entities.Secret, error) {
	if err := d.canMutateResourcesInWorkspace(ctx, secret.Namespace); err != nil {
		return nil, err
	}

	secret.EnsureGVK()
	if err := d.k8sExtendedClient.ValidateStruct(ctx, &secret.Secret); err != nil {
		return nil, err
	}

	exSecret, err := d.findSecret(ctx, secret.Namespace, secret.Name)
	if err != nil {
		return nil, err
	}

	if exSecret.Type != secret.Type {
		return nil, fmt.Errorf("updating secret.type is forbidden")
	}

  exSecret.IncrementRecordVersion()
	exSecret.ObjectMeta.Labels = secret.ObjectMeta.Labels
	exSecret.ObjectMeta.Annotations = secret.ObjectMeta.Annotations
	exSecret.Data = secret.Data
	exSecret.StringData = secret.StringData
	exSecret.SyncStatus = t.GenSyncStatus(t.SyncActionApply, exSecret.RecordVersion)

	upSecret, err := d.secretRepo.UpdateById(ctx, exSecret.Id, exSecret)
	if err != nil {
		return nil, err
	}

	if err := d.applyK8sResource(ctx, &upSecret.Secret, upSecret.RecordVersion); err != nil {
		return upSecret, err
	}

	return upSecret, nil
}

func (d *domain) DeleteSecret(ctx ConsoleContext, namespace string, name string) error {
	if err := d.canMutateResourcesInWorkspace(ctx, namespace); err != nil {
		return err
	}

	exSecret, err := d.findSecret(ctx, namespace, name)
	if err != nil {
		return err
	}

	exSecret.SyncStatus = t.GenSyncStatus(t.SyncActionDelete, exSecret.RecordVersion)
	if _, err := d.secretRepo.UpdateById(ctx, exSecret.Id, exSecret); err != nil {
		return err
	}

	return d.deleteK8sResource(ctx, &exSecret.Secret)
}

func (d *domain) OnDeleteSecretMessage(ctx ConsoleContext, secret entities.Secret) error {
	exSecret, err := d.findSecret(ctx, secret.Namespace, secret.Name)
	if err != nil {
		return err
	}

	if err := d.MatchRecordVersion(secret.Annotations, exSecret.RecordVersion); err != nil {
		return err
	}

	return d.secretRepo.DeleteById(ctx, exSecret.Id)
}

func (d *domain) OnUpdateSecretMessage(ctx ConsoleContext, secret entities.Secret) error {
	exSecret, err := d.findSecret(ctx, secret.Namespace, secret.Name)
	if err != nil {
		return err
	}

	annotatedVersion, err := d.parseRecordVersionFromAnnotations(secret.Annotations)
	if err != nil {
		return d.resyncK8sResource(ctx, exSecret.SyncStatus.Action, &exSecret.Secret, exSecret.RecordVersion)
	}

	if annotatedVersion != exSecret.RecordVersion {
		return d.resyncK8sResource(ctx, exSecret.SyncStatus.Action, &exSecret.Secret, exSecret.RecordVersion)
	}

	exSecret.CreationTimestamp = secret.CreationTimestamp
	exSecret.Labels = secret.Labels
	exSecret.Annotations = secret.Annotations

	exSecret.Status = secret.Status

	exSecret.SyncStatus.State = t.SyncStateReceivedUpdateFromAgent
	exSecret.SyncStatus.Error = nil
	exSecret.SyncStatus.LastSyncedAt = time.Now()

	_, err = d.secretRepo.UpdateById(ctx, exSecret.Id, exSecret)
	return err
}

func (d *domain) OnApplySecretError(ctx ConsoleContext, errMsg, namespace, name string) error {
	exSecret, err2 := d.findSecret(ctx, namespace, name)
	if err2 != nil {
		return err2
	}

	exSecret.SyncStatus.State = t.SyncStateErroredAtAgent
	exSecret.SyncStatus.LastSyncedAt = time.Now()
	exSecret.SyncStatus.Error = &errMsg

	_, err := d.secretRepo.UpdateById(ctx, exSecret.Id, exSecret)
	return err
}

func (d *domain) ResyncSecret(ctx ConsoleContext, namespace, name string) error {
	if err := d.canMutateResourcesInWorkspace(ctx, namespace); err != nil {
		return err
	}

	exSecret, err := d.findSecret(ctx, namespace, name)
	if err != nil {
		return err
	}

	return d.resyncK8sResource(ctx, exSecret.SyncStatus.Action, &exSecret.Secret, exSecret.RecordVersion)
}
