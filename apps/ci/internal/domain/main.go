package domain

import (
	"context"
	"encoding/json"
	"fmt"
	"kloudlite.io/constants"
	"math"
	"strings"
	"time"

	"kloudlite.io/grpc-interfaces/kloudlite.io/rpc/finance"
	"kloudlite.io/pkg/config"
	fn "kloudlite.io/pkg/functions"
	"kloudlite.io/pkg/redpanda"

	"github.com/google/go-github/v43/github"
	"github.com/xanzy/go-gitlab"
	"go.uber.org/fx"
	"golang.org/x/oauth2"
	"kloudlite.io/grpc-interfaces/kloudlite.io/rpc/auth"
	"kloudlite.io/pkg/errors"
	"kloudlite.io/pkg/harbor"
	"kloudlite.io/pkg/logging"
	"kloudlite.io/pkg/repos"
	t "kloudlite.io/pkg/types"
)

type HarborHost string

type domainI struct {
	pipelineRepo    repos.DbRepo[*Pipeline]
	pipelineRunRepo repos.DbRepo[*PipelineRun]
	authClient      auth.AuthClient
	github          Github
	gitlab          Gitlab
	harborAccRepo   repos.DbRepo[*HarborAccount]
	harborHost      HarborHost
	logger          logging.Logger
	harborCli       *harbor.Client
	gitRepoHookRepo repos.DbRepo[*GitRepositoryHook]

	gitlabWebhookUrl string
	githubWebhookUrl string
	env              *Env
	producer         redpanda.Producer
	financeClient    finance.FinanceClient
	pipelineTemplate PipelineTemplate
}

func (d *domainI) ListPipelinesByGitInfo(ctx context.Context, url, provider, branch string) ([]*Pipeline, error) {
	return d.pipelineRepo.Find(
		ctx, repos.Query{
			Filter: repos.Filter{"git_repo_url": url, "git_provider": provider, "git_branch": branch},
		},
	)
}

func (d *domainI) createNewPipelineRun(ctx context.Context, pipeline *Pipeline) (*PipelineRun, error) {
	if pipeline == nil {
		return nil, errors.New("pipeline is nil")
	}

	return d.pipelineRunRepo.Create(
		ctx, &PipelineRun{
			// BaseEntity: repos.BaseEntity{
			// 	Id: d.pipelineRunRepo.NewId(),
			// },
			PipelineID:       pipeline.Id,
			CreationTime:     time.Now(),
			Success:          false,
			Message:          "not-started-yet",
			State:            PipelineStateIdle,
			GitProvider:      pipeline.GitProvider,
			GitBranch:        pipeline.GitBranch,
			GitRepo:          pipeline.GitRepoUrl,
			Build:            pipeline.Build,
			Run:              pipeline.Run,
			DockerBuildInput: &pipeline.DockerBuildInput,
			ArtifactRef:      pipeline.ArtifactRef,
		},
	)
}

func (d *domainI) CreateNewPipelineRun(ctx context.Context, pipelineId repos.ID) (*PipelineRun, error) {
	pipeline, err := d.pipelineRepo.FindById(ctx, pipelineId)
	if err != nil {
		return nil, err
	}
	return d.createNewPipelineRun(ctx, pipeline)
}

func (d *domainI) UpdatePipelineRunStatus(ctx context.Context, pStatus PipelineRunStatus) error {
	prun, err := d.pipelineRunRepo.FindById(ctx, repos.ID(pStatus.PipelineRunId))
	if err != nil {
		return err
	}

	prun.StartTime = &pStatus.StartTime
	prun.EndTime = pStatus.EndTime
	prun.Success = pStatus.Success
	prun.Message = pStatus.Message

	prun.State = func() PipelineState {
		if pStatus.EndTime == nil {
			return PipelineStateInProgress
		}
		if (pStatus.EndTime).Sub(pStatus.StartTime) > 0 {
			if pStatus.Success {
				return PipelineStatueSuccess
			}
			return PipelineStateError
		}
		return PipelineStateIdle
	}()

	if _, err = d.pipelineRunRepo.UpdateById(ctx, repos.ID(pStatus.PipelineRunId), prun); err != nil {
		return err
	}
	return nil
}

func (d *domainI) ListPipelineRuns(ctx context.Context, pipelineId repos.ID) ([]*PipelineRun, error) {
	pr, err := d.pipelineRunRepo.Find(ctx, repos.Query{Filter: repos.Filter{"pipeline_id": pipelineId}})
	if err != nil {
		return nil, err
	}
	return pr, nil
}

func (d *domainI) GetPipelineRun(ctx context.Context, pipelineRunId repos.ID) (*PipelineRun, error) {
	return d.pipelineRunRepo.FindById(ctx, pipelineRunId)
}

type Env struct {
	GithubWebhookUrl string `env:"GITHUB_WEBHOOK_URL" required:"true"`
	GitlabWebhookUrl string `env:"GITLAB_WEBHOOK_URL" required:"true"`

	GithubWebhookAuthzSecret string `env:"GITHUB_WEBHOOK_AUTHZ_SECRET" required:"true"`
	GitlabWebhookAuthzSecret string `env:"GITLAB_WEBHOOK_AUTHZ_SECRET" required:"true"`
	KlHookTriggerAuthzSecret string `env:"KL_HOOK_TRIGGER_AUTHZ_SECRET" required:"true"`
}

const (
	GitlabProvider string = "gitlab"
	GithubProvider string = "github"
)

func (d *domainI) HarborImageSearch(ctx context.Context, accountId repos.ID, q string, pagination *t.Pagination) ([]harbor.Repository, error) {
	return d.harborCli.SearchRepositories(
		ctx, accountId, q, harbor.ListOptions{
			Page: func() int64 {
				if pagination == nil {
					return 1
				}
				return int64(pagination.Page)
			}(),
			PageSize: func() int64 {
				if pagination == nil {
					return 20
				}
				return int64(pagination.PerPage)
			}(),
		},
	)
}

func (d *domainI) HarborImageTags(ctx context.Context, imageName string, pagination *t.Pagination) ([]harbor.ImageTag, error) {
	n := strings.SplitN(imageName, "/", 2)
	projectName := n[0]
	resourceName := n[1]

	return d.harborCli.ListTags(
		ctx, projectName, resourceName, harbor.ListTagsOpts{
			ListOptions: harbor.ListOptions{
				Page: func() int64 {
					if pagination == nil {
						return 1
					}
					return int64(pagination.Page)
				}(),
				PageSize: func() int64 {
					if pagination == nil {
						return 10
					}
					return int64(pagination.PerPage)
				}(),
			},
		},
	)
}

func (d *domainI) GetAppPipelines(ctx context.Context, appId repos.ID) ([]*Pipeline, error) {
	find, err := d.pipelineRepo.Find(
		ctx, repos.Query{
			Filter: repos.Filter{
				"app_id": appId,
			},
		},
	)
	if err != nil {
		return nil, err
	}
	return find, nil
}

type GitWebhookPayload struct {
	GitProvider string `json:"git_provider,omitempty"`
	RepoUrl     string `json:"repo_url,omitempty"`
	GitBranch   string `json:"git_branch,omitempty"`
	CommitHash  string `json:"commit_hash,omitempty"`
}

func getBranchFromRef(gitRef string) string {
	sp := strings.SplitN(gitRef, "/", 3)
	if len(sp) == 3 {
		return sp[2]
	}
	return ""
}

type ErrEventNotSupported struct {
	err error
}

func (e *ErrEventNotSupported) Error() string {
	return e.err.Error()
}

func (d *domainI) ParseGithubHook(eventType string, hookBody []byte) (*GitWebhookPayload, error) {
	hook, err := github.ParseWebHook(eventType, hookBody)
	if err != nil {
		d.logger.Infof("bad webhook body, dropping message ...")
		return nil, err
	}

	switch h := hook.(type) {
	case *github.PushEvent:
		{
			payload := GitWebhookPayload{
				GitProvider: constants.ProviderGithub,
				RepoUrl:     *h.Repo.HTMLURL,
				GitBranch:   getBranchFromRef(h.GetRef()),
				CommitHash:  h.GetAfter()[:int(math.Min(10, float64(len(h.GetAfter()))))],
			}
			return &payload, nil
		}
	default:
		return nil, &ErrEventNotSupported{err: fmt.Errorf("event type (%s), currently not supported", eventType)}
	}
}

func (d *domainI) ParseGitlabHook(eventType string, hookBody []byte) (*GitWebhookPayload, error) {
	hook, err := gitlab.ParseWebhook(gitlab.EventType(eventType), hookBody)
	if err != nil {
		return nil, errors.NewEf(err, "could not parse webhook body for eventType=%s", eventType)
	}
	switch h := hook.(type) {
	case *gitlab.PushEvent:
		{
			payload := &GitWebhookPayload{
				GitProvider: constants.ProviderGitlab,
				RepoUrl:     h.Repository.GitHTTPURL,
				GitBranch:   getBranchFromRef(h.Ref),
				CommitHash:  h.CheckoutSHA[:int(math.Min(10, float64(len(h.CheckoutSHA))))],
			}
			return payload, nil
		}
	default:
		return nil, &ErrEventNotSupported{err: fmt.Errorf("event type (%s) currently not supported", eventType)}
	}
}

func (d *domainI) gitlabPullToken(ctx context.Context, accTokenReq *auth.GetAccessTokenRequest) (string, error) {
	accessToken, err := d.authClient.GetAccessToken(ctx, accTokenReq)
	if err != nil || accessToken == nil {
		return "", errors.NewEf(err, "could not get gitlab access token")
	}

	accToken := AccessToken{
		UserId:   repos.ID(accessToken.UserId),
		Email:    accessToken.Email,
		Provider: accessToken.Provider,
		Token: &oauth2.Token{
			AccessToken:  accessToken.OauthToken.AccessToken,
			TokenType:    accessToken.OauthToken.TokenType,
			RefreshToken: accessToken.OauthToken.RefreshToken,
			Expiry:       time.UnixMilli(accessToken.OauthToken.Expiry),
		},
		Data: nil,
	}
	repoToken, err := d.gitlab.RepoToken(ctx, &accToken)
	if err != nil || repoToken == nil {
		return "", errors.NewEf(err, "could not get repoToken")
	}
	return repoToken.AccessToken, nil
}

func (d *domainI) GitlabPullToken(ctx context.Context, tokenId repos.ID) (string, error) {
	return d.gitlabPullToken(ctx, &auth.GetAccessTokenRequest{TokenId: string(tokenId)})
}

func (d *domainI) ListPipelines(ctx context.Context, projectId repos.ID) ([]*Pipeline, error) {
	find, err := d.pipelineRepo.Find(
		ctx, repos.Query{
			Filter: repos.Filter{
				"project_id": projectId,
			},
		},
	)
	if err != nil {
		return nil, err
	}
	return find, nil
}

func (d *domainI) GetPipelineRunParams(ctx context.Context, p *Pipeline, pRun *PipelineRun) (*TektonVars, error) {
	if p == nil || pRun == nil {
		return nil, errors.Newf("pipeline and pipeline-run must be non-nil")
	}

	pullToken, err := func() (string, error) {
		if p.GitProvider == GithubProvider {
			return d.github.GetInstallationToken(ctx, p.GitRepoUrl)
		}
		if p.GitProvider == GitlabProvider {
			if p.AccessTokenId == "" {
				return "", errors.Newf("pipeline has no associated gitlab access-token information")
			}
			return d.gitlabPullToken(ctx, &auth.GetAccessTokenRequest{TokenId: string(p.AccessTokenId)})
		}
		return "", errors.Newf("unknown git provider")
	}()
	if err != nil {
		return nil, err
	}

	return &TektonVars{
		PipelineId:    p.Id,
		PipelineRunId: pRun.Id,
		AccountId:     p.AccountId,
		GitRepo:       p.GitRepoUrl,
		GitUser: func() string {
			if p.GitProvider == "github" {
				return "x-access-token"
			}
			if p.GitProvider == "gitlab" {
				return "oauth2"
			}
			return ""
		}(),
		GitPassword:             pullToken,
		GitBranch:               p.GitBranch,
		GitCommitHash:           "",
		IsDockerBuild:           p.DockerBuildInput.DockerFile != "",
		DockerFile:              &p.DockerBuildInput.DockerFile,
		DockerContextDir:        &p.DockerBuildInput.ContextDir,
		DockerBuildArgs:         &p.DockerBuildInput.BuildArgs,
		BuildBaseImage:          fn.DefaultIfNil(p.Build).BaseImage,
		BuildCmd:                fn.DefaultIfNil(p.Build).Cmd,
		BuildOutputDir:          fn.DefaultIfNil(p.Build).OutputDir,
		RunBaseImage:            fn.DefaultIfNil(p.Run).BaseImage,
		RunCmd:                  fn.DefaultIfNil(p.Run).Cmd,
		TaskNamespace:           p.ProjectName,
		ArtifactDockerImageName: fmt.Sprintf("%s/%s/%s", d.harborHost, p.AccountId, p.ArtifactRef.DockerImageName),
		ArtifactDockerImageTag:  p.ArtifactRef.DockerImageTag,
	}, nil
}

func (d *domainI) GitlabListGroups(ctx context.Context, userId repos.ID, query *string, pagination *t.Pagination) (any, error) {
	token, err := d.getAccessTokenByUserId(ctx, "gitlab", userId)
	if err != nil {
		return nil, err
	}
	return d.gitlab.ListGroups(ctx, token, query, pagination)
}

func (d *domainI) GitlabListRepos(ctx context.Context, userId repos.ID, gid string, query *string, pagination *t.Pagination) (any, error) {
	token, err := d.getAccessTokenByUserId(ctx, "gitlab", userId)
	if err != nil {
		return nil, err
	}
	return d.gitlab.ListRepos(ctx, token, gid, query, pagination)
}

func (d *domainI) GitlabListBranches(ctx context.Context, userId repos.ID, repoId string, query *string, pagination *t.Pagination) (any, error) {
	token, err := d.getAccessTokenByUserId(ctx, "gitlab", userId)
	if err != nil {
		return nil, err
	}
	return d.gitlab.ListBranches(ctx, token, repoId, query, pagination)
}

func (d *domainI) GitlabAddWebhook(ctx context.Context, userId repos.ID, repoId string, pipelineId repos.ID) (repos.ID, error) {
	grHook, err := d.gitRepoHookRepo.FindOne(ctx, repos.Filter{"httpUrl": d.gitlabWebhookUrl})
	if err != nil {
		return "", err
	}

	token, err := d.getAccessTokenByUserId(ctx, "gitlab", userId)
	if err != nil {
		return "", err
	}

	if grHook != nil {
		exists, _ := d.gitlab.CheckWebhookExists(ctx, token, repoId, grHook.GitlabWebhookId)
		if exists {
			return grHook.Id, nil
		}
		if err := d.gitRepoHookRepo.DeleteById(ctx, grHook.Id); err != nil {
			return "", err
		}
	}

	webhookId, err := d.gitlab.AddWebhook(ctx, token, repoId, string(pipelineId))
	if err != nil {
		return "", err
	}

	grHook, err = d.gitRepoHookRepo.Create(
		ctx, &GitRepositoryHook{
			HttpUrl:         d.gitlabWebhookUrl,
			GitProvider:     GitlabProvider,
			GitlabWebhookId: webhookId,
		},
	)

	if err != nil {
		return "", err
	}
	return grHook.Id, nil
}

func (d *domainI) SaveUserAcc(ctx context.Context, acc *HarborAccount) error {
	_, err := d.harborAccRepo.Create(ctx, acc)
	if err != nil {
		return errors.NewEf(err, "[dbRepo] failed to create harbor account")
	}
	return nil
}

func (d *domainI) getAccessTokenByTokenId(ctx context.Context, tokenId repos.ID) (*AccessToken, error) {
	accTokenOut, err := d.authClient.GetAccessToken(
		ctx, &auth.GetAccessTokenRequest{
			TokenId: string(tokenId),
		},
	)
	if err != nil {
		return nil, errors.NewEf(err, "finding accessToken")
	}
	return &AccessToken{
		Id:       repos.ID(accTokenOut.Id),
		UserId:   repos.ID(accTokenOut.UserId),
		Email:    accTokenOut.Email,
		Provider: accTokenOut.Provider,
		Token: &oauth2.Token{
			AccessToken:  accTokenOut.OauthToken.AccessToken,
			TokenType:    accTokenOut.OauthToken.TokenType,
			RefreshToken: accTokenOut.OauthToken.RefreshToken,
			Expiry:       time.UnixMilli(accTokenOut.OauthToken.Expiry),
		},
	}, err
}

func (d *domainI) getAccessTokenByUserId(ctx context.Context, provider string, userId repos.ID) (*AccessToken, error) {
	accTokenOut, err := d.authClient.GetAccessToken(
		ctx, &auth.GetAccessTokenRequest{
			UserId:   string(userId),
			Provider: provider,
		},
	)
	if err != nil {
		return nil, errors.NewEf(err, "finding accessToken")
	}
	return &AccessToken{
		Id:       repos.ID(accTokenOut.Id),
		UserId:   repos.ID(accTokenOut.UserId),
		Email:    accTokenOut.Email,
		Provider: accTokenOut.Provider,
		Token: &oauth2.Token{
			AccessToken:  accTokenOut.OauthToken.AccessToken,
			TokenType:    accTokenOut.OauthToken.TokenType,
			RefreshToken: accTokenOut.OauthToken.RefreshToken,
			Expiry:       time.UnixMilli(accTokenOut.OauthToken.Expiry),
		},
	}, err
}

func (d *domainI) GithubInstallationToken(ctx context.Context, repoUrl string) (string, error) {
	return d.github.GetInstallationToken(ctx, repoUrl)
}

func (d *domainI) GithubListBranches(ctx context.Context, userId repos.ID, repoUrl string, pagination *t.Pagination) (any, error) {
	token, err := d.getAccessTokenByUserId(ctx, "github", userId)
	if err != nil {
		return "", err
	}
	return d.github.ListBranches(ctx, token, repoUrl, pagination)
}

func (d *domainI) GithubAddWebhook(ctx context.Context, userId repos.ID, repoUrl string) (repos.ID, error) {
	grHook, err := d.gitRepoHookRepo.Upsert(
		ctx, repos.Filter{"httpUrl": repoUrl}, &GitRepositoryHook{
			HttpUrl:     repoUrl,
			GitProvider: GithubProvider,
		},
	)
	if err != nil {
		return "", err
	}

	return grHook.Id, nil

	// grHook, err := d.gitRepoHookRepo.FindOne(ctx, repos.Filter{"httpUrl": d.githubWebhookUrl})
	// if err != nil {
	// 	return "", err
	// }

	// token, err := d.getAccessTokenByUserId(ctx, "github", userId)
	// if err != nil {
	// 	return "", err
	// }

	// if grHook != nil {
	// 	exists, _ := d.github.CheckWebhookExists(ctx, token, repoUrl, grHook.GithubWebhookId)
	// 	if exists {
	// 		return grHook.Id, nil
	// 	}
	// 	if err := d.gitRepoHookRepo.DeleteById(ctx, grHook.Id); err != nil {
	// 		return "", err
	// 	}
	// }

	// webhookId, err := d.github.AddWebhook(ctx, token, repoUrl, d.githubWebhookUrl)
	// if err != nil {
	// 	return "", err
	// }

	// grHook, err = d.gitRepoHookRepo.Create(
	// 	ctx, &GitRepositoryHook{
	// 		HttpUrl:     d.githubWebhookUrl,
	// 		GitProvider: GithubProvider,
	// 		// GithubWebhookId: webhookId,
	// 	},
	// )

	// if err != nil {
	// 	return "", err
	// }
	// return grHook.Id, nil
}

func (d *domainI) GithubSearchRepos(ctx context.Context, userId repos.ID, q, org string, pagination *t.Pagination) (any, error) {
	token, err := d.getAccessTokenByUserId(ctx, "github", userId)
	if err != nil {
		return nil, err
	}
	if err != nil {
		return nil, errors.NewEf(err, "while finding accessToken")
	}
	return d.github.SearchRepos(ctx, token, q, org, pagination)
}

func (d *domainI) GithubListRepos(ctx context.Context, userId repos.ID, installationId int64, pagination *t.Pagination) (any, error) {
	token, err := d.getAccessTokenByUserId(ctx, "github", userId)
	if err != nil {
		return nil, err
	}
	return d.github.ListRepos(ctx, token, installationId, pagination)
}

func (d *domainI) GithubListInstallations(ctx context.Context, userId repos.ID, pagination *t.Pagination) (any, error) {
	token, err := d.getAccessTokenByUserId(ctx, "github", userId)
	if err != nil {
		return nil, err
	}
	i, err := d.github.ListInstallations(ctx, token, pagination)
	if err != nil {
		return nil, err
	}
	return i, nil
}

func (d *domainI) CreatePipeline(ctx context.Context, userId repos.ID, pipeline Pipeline) (*Pipeline, error) {
	exP, err := d.pipelineRepo.FindOne(
		ctx, repos.Filter{
			"project_id":     pipeline.ProjectId,
			"app_id":         pipeline.AppId,
			"container_name": pipeline.ContainerName,
		},
	)
	if err != nil {
		return nil, err
	}

	if exP != nil {
		pipeline.Id = exP.Id
	} else {
		pipeline.Id = d.pipelineRepo.NewId()
	}

	// latestCommit := ""
	// 		return nil, err
	// 	}
	//
	// 	pipeline.AccessTokenId = token.Id
	// 	hookId, err := d.GithubAddWebhook(ctx, userId, pipeline.GitRepoUrl)
	// 	if err != nil {
	// 		return nil, err
	// 	}
	// 	pipeline.WebhookId = hookId
	//
	// 	// pipeline.WebhookId = hookId
	// 	commit, err := d.github.GetLatestCommit(ctx, token, pipeline.GitRepoUrl, pipeline.GitBranch)
	// 	if err != nil {
	// 		return nil, err
	// 	}
	// 	latestCommit = commit
	// }

	if pipeline.GitProvider == constants.ProviderGitlab {
		token, err := d.getAccessTokenByUserId(ctx, pipeline.GitProvider, userId)
		if err != nil {
			return nil, err
		}
		pipeline.AccessTokenId = token.Id
		hookId, err := d.GitlabAddWebhook(ctx, userId, d.gitlab.GetRepoId(pipeline.GitRepoUrl), pipeline.Id)
		if err != nil {
			return nil, err
		}
		pipeline.WebhookId = hookId

		// commit, err := d.gitlab.GetLatestCommit(ctx, token, pipeline.GitRepoUrl, pipeline.GitBranch)
		// if err != nil {
		// 	return nil, err
		// }
		// latestCommit = commit
	}

	p, err := d.pipelineRepo.Upsert(ctx, repos.Filter{"id": pipeline.Id}, &pipeline)
	if err != nil {
		return nil, err
	}
	if err := d.TriggerPipeline(ctx, userId, p.Id); err != nil {
		return nil, errors.NewEf(err, "failed to trigger pipeline run")
	}
	return p, nil
}

func (d *domainI) DeletePipeline(ctx context.Context, pipelineId repos.ID) error {
	return d.pipelineRepo.DeleteById(ctx, pipelineId)

	// if pipeline.GitProvider == common.ProviderGithub {
	//	token, err := d.getAccessTokenByUserId(ctx, "github", userId)
	//	if err != nil {
	//		return err
	//	}
	//	//if err := d.github.DeleteWebhook(ctx, token, pipeline.GitRepoUrl, *pipeline.WebhookId); err != nil {
	//	//	return err
	//	//}
	// }
	//
	// if pipeline.GitProvider == common.ProviderGitlab {
	//	token, err := d.getAccessTokenByUserId(ctx, "gitlab", userId)
	//	if err != nil {
	//		return err
	//	}
	//	//if err := d.gitlab.DeleteWebhook(ctx, token, pipeline.GitRepoUrl, *pipeline.WebhookId); err != nil {
	//	//	return err
	//	//}
	// }
}

func (d *domainI) TriggerPipeline(ctx context.Context, userId repos.ID, pipelineId repos.ID) error {
	// TODO: (nxtcoder17) Can this userId, trigger pipeline ?

	pipeline, err := d.pipelineRepo.FindById(ctx, pipelineId)
	if err != nil {
		return err
	}

	cluster, err := d.financeClient.GetAttachedCluster(ctx, &finance.GetAttachedClusterIn{AccountId: pipeline.AccountId})
	if err != nil {
		return err
	}

	pRun, err := d.createNewPipelineRun(ctx, pipeline)
	if err != nil {
		return err
	}

	pRunParams, err := d.GetPipelineRunParams(ctx, pipeline, pRun)
	if err != nil {
		return err
	}

	b, err := d.pipelineTemplate.RenderPipelineRun([]*TektonVars{pRunParams})
	if err != nil {
		return err
	}

	agentMsgBytes, err := json.Marshal(map[string]any{"action": "create", "yamls": b.Bytes()})
	if err != nil {
		return err
	}

	ctx, cancelFn := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancelFn()
	topicName := cluster.ClusterId + "-incoming"

	pMsg, err := d.producer.Produce(ctx, topicName, pRunParams.GitRepo, agentMsgBytes)
	if err != nil {
		d.logger.Errorf(err, "error processing message, could not pipeline output into topic=%s", topicName)
		return err
	}
	d.logger.Infof("processed git webhook, pipelined output into to topic=%s, offset=%d", pMsg.Topic, pMsg.Offset)
	return nil
}

func (d *domainI) GetPipeline(ctx context.Context, pipelineId repos.ID) (*Pipeline, error) {
	id, err := d.pipelineRepo.FindById(ctx, pipelineId)
	if err != nil {
		return nil, err
	}
	return id, nil
}

var Module = fx.Module(
	"domain",

	fx.Provide(config.LoadEnv[Env]()),
	fx.Provide(
		func(
			pipelineRepo repos.DbRepo[*Pipeline],
			pipelineRunRepo repos.DbRepo[*PipelineRun],
			authClient auth.AuthClient,
			gitlab Gitlab,
			github Github,
			harborHost HarborHost,
			logger logging.Logger,
			harborCli *harbor.Client,
			gitRepoHookRepo repos.DbRepo[*GitRepositoryHook],
			producer redpanda.Producer,
			financeClient finance.FinanceClient,
			pt PipelineTemplate,
			env *Env,
		) Domain {
			return &domainI{
				env:              env,
				authClient:       authClient,
				pipelineRepo:     pipelineRepo,
				pipelineRunRepo:  pipelineRunRepo,
				gitlab:           gitlab,
				github:           github,
				harborHost:       harborHost,
				logger:           logger.WithName("[ci]:domain/main.go"),
				harborCli:        harborCli,
				gitRepoHookRepo:  gitRepoHookRepo,
				gitlabWebhookUrl: env.GitlabWebhookUrl,
				githubWebhookUrl: env.GithubWebhookUrl,
				financeClient:    financeClient,
				producer:         producer,
				pipelineTemplate: pt,
			}
		},
	),
)
