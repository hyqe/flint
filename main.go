package main

import (
	"log"
	"os"

	_ "embed"

	"github.com/hyqe/flint/internal/cache"
	"github.com/hyqe/flint/internal/handlers"
	"github.com/hyqe/graceful"
	cli "github.com/urfave/cli/v2"
)

const (
	AppName = "Flint"
)

var (
	//go:embed LICENSE
	License string

	//go:embed VERSION
	Version string

	//go:embed docs.md
	Docs string
)

func main() {
	cliApp := &cli.App{
		Name:        AppName,
		Description: "\n" + Docs,
		Copyright:   "\n" + License,
		Version:     Version,
		Flags: []cli.Flag{
			&cli.IntFlag{
				Name:    "port",
				Aliases: []string{"p"},
				EnvVars: []string{"FLINT_PORT"},
				Value:   2000,
			},
			&cli.BoolFlag{
				Name:    "verbose",
				EnvVars: []string{"FLINT_VERBOSE"},
				Value:   false,
			},
		},
	}
	cliApp.Action = action
	err := cliApp.Run(os.Args)
	if err != nil {
		log.Fatal(err)
	}
}

func action(c *cli.Context) error {
	verbose := c.Bool("verbose")
	port := c.Int("port")

	handler := &handlers.Flint{
		Cacher:  cache.New(),
		Verbose: verbose,
	}

	server := graceful.NewServer(
		graceful.WithHandler(handler),
		graceful.WithPort(port),
	)

	if verbose {
		log.Printf("starting Flint on port: %v\n", port)
	}
	return graceful.Run(server)
}
