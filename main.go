package main

import (
	"log"
	"os"
	"strings"

	_ "embed"

	"github.com/hyqe/flint/internal/handlers"
	"github.com/hyqe/flint/internal/storage"
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
			&cli.StringFlag{
				Name:    "storage",
				EnvVars: []string{"FLINT_STORAGE"},
				Usage:   "a directory to store values to disk",
				Value:   "",
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
	storagePath := c.String("storage")

	var store storage.Storage = &storage.Memory{}
	if strings.TrimSpace(storagePath) != "" {
		if _, err := os.Stat(storagePath); os.IsNotExist(err) {
			if err := os.MkdirAll(storagePath, os.ModePerm); err != nil {
				log.Fatal(err)
			}
		}
		store = &storage.FS{
			Path: storagePath,
		}
	}

	handler := &handlers.Flint{
		Storage: store,
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
