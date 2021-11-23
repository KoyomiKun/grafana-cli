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
	fromUrl string
	toUrl   string
	fromKey string
	toKey   string
)

func main() {
	currentDir, _ := os.Getwd()
	app := &cli.App{
		Name:        "grafana-cli",
		Description: "grafana cmd client",
		Flags: []cli.Flag{
			&cli.StringFlag{
				Name:        "config",
				Aliases:     []string{"c"},
				Value:       path.Join(currentDir, "/cfg/config.json"),
				Destination: &cfgPath,
			},
		},
		Commands: []*cli.Command{
			{
				Name: "migrate",
				Flags: []cli.Flag{
					&cli.StringFlag{
						Name:        "from",
						Required:    true,
						Destination: &fromUrl,
					},
					&cli.StringFlag{
						Name:        "to",
						Required:    true,
						Destination: &toUrl,
					},
					&cli.StringFlag{
						Name:        "from-key",
						Required:    true,
						Destination: &fromKey,
					},
					&cli.StringFlag{
						Name:        "to-key",
						Required:    true,
						Destination: &toKey,
					},
				},
				Action: func(c *cli.Context) error {
					return startMigrate()
				},
			},
			{
				Name: "get",
				Subcommands: []*cli.Command{
					{
						Name:    "dashboards",
						Aliases: []string{"dashboard", "db"},
						Action: func(c *cli.Context) error {
							return startGetDb()
						},
					},
					{
						Name:    "alerts",
						Aliases: []string{"alert"},
						Action: func(c *cli.Context) error {
							return startGetAlert()
						},
					},
					{
						Name:    "tags",
						Aliases: []string{"tag"},
						Action: func(c *cli.Context) error {
							return startGetTag()
						},
					},
				},
			},
			{
				Name: "update",
				Subcommands: []*cli.Command{
					{
						Name:    "tags",
						Aliases: []string{"tag"},
						Action: func(c *cli.Context) error {
							return startUpdateTag()
						},
					},
				},
			},
		},
	}
	app.Run(os.Args)
}

func startMigrate() error {

	// init client
	clientFrom := client.NewClient(
		5*time.Second,
		fromUrl,
		fromKey,
	)
	clientTo := client.NewClient(
		5*time.Second,
		toUrl,
		toKey,
	)

	// migrate folders
	fdmFrom := manager.NewFolderManager(clientFrom)
	fdmTo := manager.NewFolderManager(clientTo)

	fds, err := fdmFrom.List()
	if err != nil {
		log.Errorf("fail getting folders: %v", err)
		return err
	}
	for _, fd := range fds {
		err := fdmTo.CreateFolder(fd)
		if err != nil {
			log.Warnf("fail create folder %s: %v", fd.Title, err)
		}
	}

	// migrate dashboards
	dbmFrom := manager.NewDashboardManager(clientFrom)
	dbmTo := manager.NewDashboardManager(clientTo)

	dbs, err := dbmFrom.List()
	if err != nil {
		log.Errorf("fail listing dashboards: %v", err)
		return err
	}

	for _, db := range dbs {
		db.Dashboard["id"] = nil
		db.Dashboard["version"] = 0
		err := dbmTo.Update(db)
		if err != nil {
			log.Warnf("fail updating dashbaord: %v", err)
			continue
		}
	}

	return nil
	// save folders to new one
	// get all old dashboards
	// save them to new dashboards
	// get all old datasources
	// save
}

func startGetAlert() error {
	return nil
}

func startGetTag() error {
	return nil
}

func startGetDb() error {
	return nil
}

func startUpdateTag() error {

	// init config
	config, err := config.NewConfig(cfgPath)
	if err != nil {
		log.Errorf("Fail creating config in %s: %v", cfgPath, err)
		return err
	}

	// init client
	client := client.NewClient(
		10*time.Second,
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
