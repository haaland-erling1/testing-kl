package app

import (
	"context"
	"fmt"
	"time"

	"github.com/kloudlite/operator/pkg/kubectl"
	"go.uber.org/fx"
	"gopkg.in/yaml.v2"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"kloudlite.io/apps/build-worker/internal/domain"
	"kloudlite.io/apps/build-worker/internal/env"
	"kloudlite.io/pkg/logging"
	"kloudlite.io/pkg/redpanda"
)

const (
	GithubEventHeader string = "X-Github-Event"
	GitlabEventHeader string = "X-Gitlab-Event"
)

type Obj struct {
	Metadata struct {
		Name string `json:"name" yaml:"name"`
	} `json:"metadata" yaml:"metadata"`
}

func fxInvokeProcessBuilds() fx.Option {
	return fx.Options(

		fx.Invoke(
			func(d domain.Domain, consumer redpanda.Consumer, logr logging.Logger, envs *env.Env, yamlClient kubectl.YAMLClient) {
				consumer.StartConsuming(
					func(msg []byte, _ time.Time, offset int64) error {

						l := logr.WithName("build-worker")

						ctx := context.TODO()

						fmt.Println("msg: ", string(msg))
						var obj Obj

						if err := yaml.Unmarshal(msg, &obj); err != nil {
							fmt.Println("err1: ", err)
							return err
						}

						_, err := yamlClient.Client().BatchV1().Jobs("kl-core").Get(ctx, obj.Metadata.Name, v1.GetOptions{})
						if err != nil {

							if rr, err := yamlClient.ApplyYAML(ctx, msg); err != nil {
								fmt.Println("err3: ", err)
								return err
							} else {
								l.Infof("created job: %s", rr)
							}
						}

						fmt.Println("end")
						return fmt.Errorf("not implemented")
					},
				)
			},
		),
	)
}
