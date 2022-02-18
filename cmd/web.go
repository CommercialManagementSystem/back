package main

import (
	"context"
	"os"

	"github.com/CommercialManagementSystem/back/internal"
	"github.com/CommercialManagementSystem/back/pkg/logger"
	"github.com/urfave/cli/v2"
)

// VERSION 版本号
var VERSION = "0.0.1"

func main() {
	logger.SetVersion(VERSION)
	ctx := logger.NewTagContext(context.Background(), "__main__")

	app := cli.NewApp()
	app.Name = "insider"
	app.Version = VERSION
	app.Usage = "QNXG insider"

	app.Flags = []cli.Flag{
		&cli.StringFlag{
			Name:        "mode",
			Aliases:     []string{"m"},
			Usage:       "run mode",
			DefaultText: "dev",
		},
	}

	app.Action = func(c *cli.Context) error {
		return internal.Run(
			ctx,
			internal.SetMode(c.String("mode")),
			internal.SetVersion(VERSION),
		)
	}

	err := app.Run(os.Args)
	if err != nil {
		logger.WithContext(ctx).Errorf(err.Error())
	}
}
