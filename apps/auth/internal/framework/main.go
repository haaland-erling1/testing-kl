package framework

import (
	"context"

	"kloudlite.io/pkg/db"
	"kloudlite.io/pkg/errors"
)

func MakeFramework(cfg *Config) (fm *Framework, e error) {
	errors.HandleErr(&e)
	ctx := context.Background()
	mongoCli, e := db.MakeMongoClient(ctx, cfg.MongoDB.Uri, cfg.MongoDB.Db)

	fm = &Framework{
		start: func() error {
			return nil
		},
	}
	return
}
