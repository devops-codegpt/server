package main

import (
	"fmt"
	"github.com/devops-codegpt/server/app"
	"github.com/urfave/cli/v2"
	"os"
)

func main() {
	application := &cli.App{
		Name:  "CodeGpt",
		Usage: "CodeGpt API for AI ChatCodeOps Service",
		Flags: []cli.Flag{
			&cli.BoolFlag{
				Name:    "cors",
				Usage:   "Whether to allow cross-domain",
				EnvVars: []string{"CORS"},
			},
			&cli.StringFlag{
				Name:    "config-file",
				Usage:   "Config file",
				EnvVars: []string{"CONFIG_FILE"},
			},
			&cli.StringFlag{
				Name:    "address",
				Usage:   "Bind address for the API server.",
				EnvVars: []string{"ADDRESS"},
				Value:   ":8089",
			},
		},
		Description: `
CodeGpt is an AI ChatCodeOps Service.

The following services are supported:
- AI Chat
- AI code analysis
- AI Devops
`,
		UsageText: `server [options]`,
		Action: func(ctx *cli.Context) error {
			server, err := app.App(
				app.WithCors(ctx.Bool("cors")),
				app.WithConfigFile(ctx.String("config-file")))
			if err != nil {
				fmt.Println("Failed to initialize server")
				return err
			}
			// Start server
			return server.Start(ctx.String("address"))
		},
	}

	err := application.Run(os.Args)
	if err != nil {
		os.Exit(1)
	}
}
