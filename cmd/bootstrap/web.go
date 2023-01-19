package bootstrap

import (
	"context"
	"fmt"

	"github.com/shankalpa12/config"
	"github.com/shankalpa12/pkg"
	"go.uber.org/fx"
)

func WebApplication() {
	fx.New(
		fx.Provide(
			config.New,
			pkg.NewHTTPServer,
		),

		fx.Invoke(startServer),
	).Run()
}

func startServer(lifecycle fx.Lifecycle, config config.Config, server pkg.HTTPServer) {
	fmt.Println("...........................Server Starting.....................................")
	lifecycle.Append(
		fx.Hook{
			OnStart: func(ctx context.Context) error {
				go func() {
					server.Start(config)
				}()
				return nil
			},
			OnStop: func(ctx context.Context) error {
				fmt.Println("There was Error While Starting the Server.. Stopping Application")
				return nil
			},
		},
	)
}
