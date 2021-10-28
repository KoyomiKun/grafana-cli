package main

import (
	"os"
	"path"
	"time"

	"github.com/urfave/cli/v2"

	"github.com/KoyomiKun/grafana-cli/pkg/client"
	"github.com/KoyomiKun/grafana-cli/pkg/config"
	"github.com/KoyomiKun/grafana-cli/pkg/manager"
	"github.com/KoyomiKun/grafana-cli/utils/log"
)

var (
	cfgPath string
)

func main() {
	currentDir, _ := os.Getwd()
	app := &cli.App{
		Name:        "grafana-cli",
		Description: "grafana cmd client",
		Commands: []*cli.Command{
			{
				Name:        "alert",
				Description: "alert management.",
				Subcommands: []*cli.Command{
					{
						Name:        "tags",
						Description: "Update tags.",
						Flags: []cli.Flag{
							&cli.StringFlag{
								Name:        "config",
								Aliases:     []string{"c"},
								Value:       path.Join(currentDir, "/cfg/config.json"),
								Destination: &cfgPath,
							},
						},
						Action: func(c *cli.Context) error {
							return startTag()
						},
					},
				},
			},
		},
	}
	app.Run(os.Args)
}

func startTag() error {

	// init config
	config, err := config.NewConfig(cfgPath)
	if err != nil {
		log.Errorf("Fail creating config in %s: %v", cfgPath, err)
		return err
	}

	// init client
	client := client.NewClient(
		2*time.Second,
		config.BaseUrl,
		config.APIKey,
	)

	// init tagManager
	tagManager := manager.NewTagManager(
		client,
		config.Alerts,
	)

	// update tag
	tagManager.Update()
	return nil

}
